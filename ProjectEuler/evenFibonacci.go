package main

import (
	"fmt"
	"math"
)

func fibbonacciClosure() func() (int, int) {
	a, b := 0, 1
	return func() (int, int) {
		a, b = b, a+b
		return a, b
	}
}

func generateFibs(upperBound int) int {
	var total int
	fib := fibbonacciClosure()
	for i := 0; i < upperBound; i++ {
		res, b := fib()
		if math.Mod(float64(res), 2) == 0 {
			total = total + res
		}
		i = b
	}
	return total
}


func main() {
	total := generateFibs(4000000)
	fmt.Println(total)
}