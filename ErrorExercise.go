package main

import (
  "fmt"
)

type ErrNegativeSqrt struct {
  badNumber float64
}

func (e *ErrNegativeSqrt) Error() string {
  return fmt.Sprint("cannot Sqrt negative number: ", e.badNumber) 
}

func Sqrt(x float64) (float64, error) {
  if x < 0 {
    return 0, &ErrNegativeSqrt{x}
  }
  z := 1.0
  for i := 0; i < 10; i += 1 {
    z -= ((z * z) - x) / (2 * z)
  }
  return z, nil
}

func main() {
  fmt.Println(Sqrt(2))
  fmt.Println(Sqrt(-2))
}

