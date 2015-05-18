
// In the terminal, execute the command `go run http1.go`, 
// then open the browser and visit `http://localhost:8000`

package main

import (
  "io"
  "net/http"
  "fmt"
)

func hello(w http.ResponseWriter, r *http.Request) {
  // arguments:
  // type `http.ResponseWriter` and its corresponding response stream, which is 
  //  actually an interface type.
  // The second is `*http.Request` and its corresponding HTTP request.

  // `io.WriteString` is a helper function to let you write a string into a 
  //    given writable stream. In Go, we call it the `io.Writer` interface.
  //  So, we write "hello world" into the w writeable stream.
  io.WriteString(w, "Hello World!")

  fmt.Println(r) // print the http request
}

func main() {
  // https://gowalker.org/net/http#HandleFunc
  http.HandleFunc("/", hello)
  // https://gowalker.org/net/http#ListenAndServe
  http.ListenAndServe(":8000", nil)
}

