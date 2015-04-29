// http://blog.golang.org/defer-panic-and-recover
package main

import "fmt"

func main() {
  fmt.Println("counting")

  for i := 0; i < 10; i++ {
    defer fmt.Println(i) // prints out in lifo order
  }

  fmt.Println("done")
}
