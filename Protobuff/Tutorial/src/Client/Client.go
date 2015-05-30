package main

// Go TCP Client.
// This client needs to:
// 1) take 2 command line arguments, the csv filename from which we need to 
//    read data, and the destination ip address/port number of the TCP server to
//    which we need to send data.
// 2) retrieve the data to send from the csv file
// 3) send the data to the destination TCP server
// Followed this guide: 
//  http://www.minaandrawos.com/2014/05/27/practical-guide-protocol-buffers-protobuf-go-golang/

import (
  "flag"
  "fmt"
  "os"
  "net"
  "io"
  "strconv"
  "encoding/csv"
  "ProtobufTest"
  "github.com/golang/protobuf/proto"
)

const (
  CLIENT_NAME = "GoClient"
  CLIENT_ID = 2
  CLIENT_DESCRIPTION = "This is a Go Protobuf client!!"
)

type Headers []string

func (h Headers) getHeaderIndex(headername string) int {
  // map each header name from the csv file to a column number.
  // itemid=0 itemname=1 itemvalue=2 itemType=3
  if len(headername) > 2 {
    for index, s := range h {
      if s == headername {
        return index
      }
    }
  }
  return -1 
}

func checkError(err error) {
  // if there is an actual error, the error message will be printed, and then 
  // the application will exit. Otherwise, nothing happens.
  if err != nil {
    fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
    os.Exit(1)
  }
}

func retrieveDataFromFile(fname *string)([]byte, error) {
  // takes a pointer to a string for the source filename, and returns an array
  // of bytes representing the serialized protobuff extracted from the csv file.
  // Also returns an error type, either containing nothing or the error that 
  // occurred while serializing the data.
  file, err := os.Open(*fname)
  checkError(err)
  defer file.Close() // means whenever the function exits, close the file object

  // use csvreader to read in the headers (first line)
  csvreader := csv.NewReader(file)
  var hdrs Headers // make an instance of the Headers type, an array of strings
  hdrs, err = csvreader.Read() // set hdrs equal to the first line in the csv
  checkError(err)
  ITEMIDINDEX := hdrs.getHeaderIndex("itemid")
  ITEMNAMEINDEX := hdrs.getHeaderIndex("itemname")
  ITEMVALUEINDEX := hdrs.getHeaderIndex("itemvalue")
  ITEMTYPEINDEX := hdrs.getHeaderIndex("itemType") 

  // initialize the protobuf that will contain the message, then populate it.
  // The keyword “new” will create a pointer that points to the 
  //  ProtobufTest.TestMessage type.
  ProtoMessage :=  new(ProtobufTest.TestMessage)
  ProtoMessage.ClientName = proto.String(CLIENT_NAME)
  ProtoMessage.ClientId = proto.Int32(CLIENT_ID)
  ProtoMessage.Description = proto.String(CLIENT_DESCRIPTION)

  // loop through the rest of the csv file, extract the data and pass it to the
  //  ProtoMessage pointer.
  // Remember, the message we are trying to send contains a list of 
  //  “messageitems” at which we store item id, item name, item value and item 
  //  type.
  for { // go's "while True", essentially
    record, err := csvreader.Read()
    if (err != io.EOF) {
      checkError(err)
    } else {
      break // this is the end of the file.
    }
    // create the MsgItem
    testMessageItem := new(ProtobufTest.TestMessage_MsgItem)
    // Strconv package is used to convert the strings we pull in from the csv 
    //  file into ints. Eg in 1,FirstItemName,222,0, we want to convert 3 of the
    //  fields into integers (1, 222, and 0).
    // http://golang.org/pkg/strconv/#Atoi

    // id
    itemid, err := strconv.Atoi(record[ITEMIDINDEX])
    checkError(err)
    // Go’s int type is a 64 bit int (int64), however the int values are int32 
    //  in the proto file. This is why another conversion to int32 is needed
    //  before passing the value to the protobuf struct.
    testMessageItem.Id = proto.Int32(int32(itemid))
    // itemName
    testMessageItem.ItemName = &record[ITEMNAMEINDEX]
    // itemValue
    itemvalue, err := strconv.Atoi(record[ITEMVALUEINDEX])
    checkError(err)
    testMessageItem.ItemValue = proto.Int32(int32(itemvalue))
    //  itemType
    itemtype,err := strconv.Atoi(record[ITEMTYPEINDEX])
    checkError(err)
    iType := ProtobufTest.TestMessage_ItemType(itemtype)
    testMessageItem.ItemType = &iType // experiment with passing by value

    // add another message item to the protobuff
    ProtoMessage.Messageitems = append(ProtoMessage.Messageitems, 
                                       testMessageItem)
    fmt.Println(record)
  }

  // Now we have constructed the protobuf message, and need to serialize it to
  //  send it across the wire. This is done with proto.Marshal().
  // fmt.Println(ProtoMessage.Messageitems)
  return proto.Marshal(ProtoMessage)
}

func sendDataToDest(data []byte, dest *string) {
  // takes dest as well as the array of serialized bytes then send them to a Go 
  //  TCP protobuf server.
  conn, err := net.Dial("tcp", *dest) // create tcp connection to dest address
  checkError(err)
  n, err := conn.Write(data) // write protobuf to that tcp connection
  checkError(err)
  fmt.Println("Sent " + strconv.Itoa(n) + " bytes")
}

func main() {
  // create a filename input flag with a default value
  filename := flag.String("f", "input.csv", "Enter the filename to read from")
  // create an ip address flag with a default value
  dest := flag.String("d", "127.0.0.1:2110", "Enter the destination address")
  // read the flags, supplied or default value
  flag.Parse()

  // try to read the data from the given file
  data, err := retrieveDataFromFile(filename)
  // see if there was an error with the data retrieval
  checkError(err)
  // send the read data to the destination location given
  sendDataToDest(data, dest)
}
