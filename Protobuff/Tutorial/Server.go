package ProtobufTest

// Go TCP Server.
// This server needs to:
// 1) Listen to the TCP socket expected to receive the data, in our case it’s IP
//  address 127.0.01 and port number 2110 (see main function in Client.go).
// 2) Read the protobuf message and extract the values from it.
// 3) Write the extracted values to a CSV file.
// 4) Ensure that no issues will occur if multiple clients send data at the same
//  time to the port.

// Lets also use different goroutines for the main tasks in the program. We will
//  have:
//  - a goroutine responsible for extracting values from the protobuf message
//  - another that writes the extracted values to a CSV file

import (
  "fmt"
  "os"
  "strconv"
  "encoding/csv"
  "ProtobufTest"
  "github.com/golang/protobuf/proto"
)

// TODO: make checkError package wide, so we don't write it twice.
func checkError(err error) {
  // if there is an actual error, the error message will be printed, and then 
  // the application will exit. Otherwise, nothing happens.
  if err != nil {
    fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
    os.Exit(1)
  }
}

func handleProtoClient(conn net.Conn, c chan *ProtobufTest.TestMessage) {
  // Extract the serialized data received via TCP then convert it to the 
  //  protobuf message type, which is then passed to our go channel so that the 
  //  writeValuesTofile() routine can use it.
  fmt.Println("** Connection established. **")
  defer conn.Close() // close the connection when the function exits

  // Create a data buffer of type byte slice with capacity of 4096.
  // Note that "data" here is a Go "slice". Slices are the closest thing to 
  //  variable sized arrays in Go. They are initialized with the "make" keyword
  //  and require an initial expected capacity. They can be extended using 
  //  functions like append() or copy().
  data := make([]byte, 4096)
  // Read the data waiting on the connection and put it in the data buffer
  n, err := conn.Read(data)
  checkError(err)
  fmt.Println("** Decoding protobuf message... **")

  // Create an struct pointer of type ProtobufTest.TestMessage struct
  protodata := new(ProtobufTest.TestMessage) // see line 59 of ProtoTest.pb.go

  // Convert the data retrieved into the ProtobufTest.TestMessage struct type.
  err := proto.Unmarshal(data[0:n], protodata)
  checkError(err)

  // Push the protobuf message into a channel
  c <- protodata
}

func writeValuesTofile(datatowrite *ProtobufTest.TestMessage) {
  // Take the values retrieved by another thread from the protobuf message then 
  //  write them out to a CSV file.
  // This should only act when we retrieve a protobuf message, so we will rely
  //  on a channel that contains such messages. In other words, this goroutine 
  //  will lock as long as there are no protobuf messages in the channel, 
  //  otherwise it will take the values from the channel and proceed to write
  //  them to a csv file.

  // Get the client information from the protobuf we've been provided
  ClientName := datatowrite.GetClientName()
  // string from int: remember, the extra int() is to convert a 32 bit int to
  //  the default 64 bit int, so we can then pass it into Itoa() (which needs a
  //  64 bit integer) and convert it to a string.
  ClientID := strconv.Itoa(int(datatowrite.GetClientId())) 
  ClientDescription := datatowrite.GetClientDescription()

  // Get the message items (a list) from that protobuf
  items := datatowrite.GetMessageItems()

  // Now write the data we've gotten to a CSV file
  fmt.Println("** Writing data to csv file... **")

  //Open file for writes, if the file does not exist then create it
  file, err := os.OpenFile("CSVValues.csv", 
                           os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
  checkError(err)
  defer file.Close() // make sure the file gets closed on function exit
  writer := csv.NewWriter(file)
  // go through the list of message items, insert them into a string array, then
  //  write them to the CSV file.
  for _, item := range items {
    // pull out data
    ID := strconv.Itoa(int(item.GetId()))
    ItemName := item.GetItemName()
    ItemValue := strconv.Itoa(int(item.GetItemValue()))
    ItemType := strconv.Itoa(int(item.GetItemType()))
    // create an array for each item
    record := []string{ ClientID, 
                        ClientName, 
                        ClientDescription, 
                        ID, 
                        ItemName, 
                        ItemValue,
                        ItemType,
    }
    // writes a single CSV record to "record" along with any necessary quoting
    writer.Write(record) 
    fmt.Println(record)
  } 
  writer.Flush() // Flush writes any buffered data to the underlying io.Writer
  fmt.Println("** Done writing csv file. **")
}


func main() {
  fmt.Println("** Started Go TCP Protobuf Server **")
  // make a channel to send/receive protobufs
  c := make(chan *ProtobufTest.TestMessage)

  // Spawn a goroutine to listen to that channel and read data when it is there.
  go func() { // this inline function is called a Go "function literal"
    for {
      message := <- c // read the channel
      writeValuesTofile(message) // call the csv writer function
    }
  }()

  // Spawn a second goroutine to listen to the address and tcp port we want the
  //  protobuf messages to come in on. Whenever data is received, this new
  //  goroutine converts it to a protobuf struct of the desired type.
  listener, err := net.Listen("tcp", "127.0.0.1:2110")
  checkError(err)
  for {
    // if the err is nil then the data is available
    if conn, err := listener.Accept(); err == nil {
      go handleProtoClient(conn, c) // connection object and channel
    } else {
      continue // keep listening
    }
  }
}