// demonstrate goroutines.
// the scanln is used so that the program won't exit too quickly to see the
// output. 

package main

import (
  "fmt"
  "time"
  "math/rand"
)

func f(n int) {
  for i := 0; i <10; i++ {
    fmt.Println(n, ":", i)
  }
}

// example of goroutine

func oneGoroutine() {
  go f(0)
  var input string
  fmt.Scanln(&input)
  fmt.Println("------")
}

// example of running 10 goroutines at once

func tenGoroutines() {
  for i:=0; i<10; i++ {
    go f(i)
  }
  var input string
  fmt.Scanln(&input)
  fmt.Println("------")
}

// demonstrate simultaneous go routines

func timedF(n int) {
  for i:=0; i<10; i++ {
    fmt.Println(n, ":", i)
    amt := time.Duration(rand.Intn(250))
    time.Sleep(time.Millisecond * amt)
  }
}

func demonstrateSimultaneousGoroutines() {
  for i:=0; i< 10; i++ {
    go timedF(i)
  }
  var input string
  fmt.Scanln(&input)
}

func main() {
  //oneGoroutine()
  // tenGoroutines()
  demonstrateSimultaneousGoroutines()
}

