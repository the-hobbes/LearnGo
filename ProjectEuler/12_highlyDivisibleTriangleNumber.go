package main

// https://projecteuler.net/problem=12

import (
	"fmt"
	"os"
)

func triangleNumberGenerator() (chan int) {
	// Return triangle numbers in order (via a channel), which are created
	// by adding the natural numbers.
	channel := make(chan int)
	go func() {
		for start := 1; ; start++ {
			nextTriangle := start + start + 1
			channel <- nextTriangle
		}
	}()
	return channel
}

func calculateDivisors(targetDivisors int) (triangleNumber int) {
	// Pull a triangle number from triangeNumberGenerators().
	// Factor it.
	// If the number of factors it has is > targetDivisors, return it.
	return
}

func test() {
	result := calculateDivisors(5)
	expectedResult := 28

	if result != expectedResult {
		fmt.Fprintf(os.Stderr, 
			"***Test failed*** \n Wanted: %i, Received: %i \n", 
			expectedResult, result)
	    os.Exit(1)
	}
	
	fmt.Println("***Test Passed***")
}

func main() {
	// test()
	// result := calculateDivisors(500)
	// fmt.Println(result)
	generator := triangleNumberGenerator()
	for i := 0; i < 10; i++ {
		x := <-generator
		fmt.Println(x)
	}
}