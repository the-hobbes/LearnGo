package main

import (
  "golang.org/x/tour/wc"
  "fmt"
  "strings"
)

func WordCount(s string) map[string]int {
  m := make(map[string]int)
  fields := strings.Fields(s)
  for i := range fields {
    key := fields[i]
    elem, ok := m[key]
    if ok {
      // element exists. Update its count
      elem += 1
      m[key] = elem
    } else {
      // element doesn't exist. create it and set count to 1
      elem = 1
      m[key] = elem
    }
  }

  
  return m
}

func main() {
  wc.Test(WordCount)
  result := WordCount("a a b c")
  fmt.Println(result)
}

