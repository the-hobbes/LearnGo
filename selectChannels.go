// this program prints "from 1" every 2 seconds and "from 2" every 3 seconds. 
//  select picks the first channel that is ready and receives from it, or sends to
//  it. If none of the channels are ready, the statement blocks until one becomes
//  available. if More that one of the channels are ready then the select statement
//  randomly picks which one to receive or send on. 
// You get the blocking and the random selection for free with the select 
//  statement. 

package main

import (
  "fmt"
  "time"
)

func main() {
  c1 := make(chan string, 1) // string channel w/a buffer of 1 message
  c2 := make(chan string, 1)

  go func() {
    for {
      c1 <- "from 1"
      time.Sleep(time.Second * 2)
    }
  }()

  go func() {
    for {
      c2 <- "from 2"
      time.Sleep(time.Second * 3) // stagger the sleep
    }
  }()

  go func() {
    for {
      // pick which channel to print
      select { 
      case msg1 := <- c1:
        fmt.Println("Message 1", msg1)
      case msg2 := <- c2:
        fmt.Println("Message 2", msg2)
      case <- time.After(time.Second):
        // time.After creates a channel and after the given duration will send 
        //  the current time on it. 
        fmt.Println("timeout")
      default:
        // the default case happens immediately if none of the channels are
        //  ready.
        fmt.Println("Nothing ready.")
      }
    }
  }()

  // exit by hitting enter
  var input string 
  fmt.Scanln(&input)
}