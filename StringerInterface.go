// Person is the struct
// We want to be able to print a Person by calling fmt.Println
// the 'fmt' package looks to the 'Stringer' interface to know how to print things
// In order for fmt to know how to print a Person, we have to give the Stringer 
//  interface a method it knows how to work with. The stringer iterface looks like:
// type Stringer interface {
//    String() string
// }
// Therefore, we have to implement the String() method on our Person struct. This way,
// when fmt looks needs to print Person, it will see that Person automatically implements
// the stringer interface because we've defined a String method that tells it how to print.

package main

import "fmt"

type Person struct {
  Name string
  Age  int
}

func (p Person) String() string {
  return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
  a := Person{"Arthur Dent", 42}
  z := Person{"Zaphod Beeblebrox", 9001}
  fmt.Println(a, z)
}

