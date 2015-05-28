package ProtobufTest

// Go TCP Client
// This client needs to:
// 1) take 2 command line arguments, the csv filename from which we need to 
//    read data, and the destination ip address/port number of the TCP server to
//    which we need to send data.
// 2) retrieve the data to send from the csv file
// 3) send the data to the destination TCP server

import (
  "flag"
  "fmt"
  "os"
  "encoding/csv"
)

func checkError(err error) {
  // if there is an actual error; the error message will be printed, and then 
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
  defer file.Close() // whenever this function exits, close the file object
  csvreader := csv.NewReader(file)
}

func sendDataToDest(data []byte, dest *string) {
}

func main() {
  // create a filename input flag with a default value
  filename := flag.String("f", "input.csv", "Enter the filename to read from")
  // create an ip address flag with a default value
  destination := flag.String("d", "127.0.0.1:2110", "Enter the destination address")
  // read the flags, supplied or default value
  flag.Parse()

  // try to read the data from the given file
  data, err := retrieveDataFromFile(filename)
  // see if there was an error with the data retrieval
  checkError(err)
  // send the read data to the destination location given
  sendDatatoDest(data, dest)
}
