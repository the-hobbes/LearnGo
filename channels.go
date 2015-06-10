package main

import (
  "fmt"
  "time"
)

func pinger(c chan string) { // use channel bidirectionally (send or receive)
  for {
    // <- is used to send and receive messages on a channel
    c <- "ping" // means send "ping"
  }
}

func ponger(c chan <- string) { // only able to send on channel
  for {
    c <- "pong"
  } 
}

func printer(c <- chan string) { // only able to receive on channel
  for {
    msg := <- c // means receive a message and store it in msg
    fmt.Println(msg)
    time.Sleep(time.Second * 1)
  }
}

func regularPingPong(c chan string) {
  // print ping pong forever via three separate goroutines, coordinated through 
  //  a channel.
  go pinger(c)
  go ponger(c)
  go printer(c)

  // pingpong until user hits enter
  var input string
  fmt.Scanln(&input)
}

func main() {
  // make a channel called c that will pass strings
  var c chan string = make(chan string)

  regularPingPong(c)


}
