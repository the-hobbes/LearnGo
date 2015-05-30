package ProtobufTest

import (
  "fmt"
  "os"
)

func checkError(err error) {
  // if there is an actual error, the error message will be printed, and then 
  // the application will exit. Otherwise, nothing happens.
  if err != nil {
    fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
    os.Exit(1)
  }
}
