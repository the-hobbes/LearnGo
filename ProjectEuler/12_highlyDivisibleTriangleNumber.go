package main

// https://projecteuler.net/problem=12

import (
	"fmt"
	"os"
	"math"
)

func triangleNumberGenerator() (chan int) {
	// Return triangle numbers in order (via a channel), which are created
	// by adding the natural numbers.
	channel := make(chan int)
	go func() {
		for current := 1; ; current++ {
			nextTriangle := 0
			for i := current; i > 0; i-- {
				nextTriangle = nextTriangle + i
			}
			channel <- nextTriangle
		}
	}()
	return channel
}

func calculateDivisors(targetDivisors int) int {
	// Pull a triangle number from triangeNumberGenerators().
	// Factor it.
	// If the number of factors it has is > targetDivisors, return it.
	// TODO(pheven): Implement this. I havent solved 12 yet!
	generator := triangleNumberGenerator()
	result :=  -1
	for {
		x := <-generator
		fmt.Println("Current Triangle:", x)
		factorsOfx := 0
		for i := 1; i <= x; i++ {
			// if x is evenly divisible by i, i is a factor of x.
			if math.Mod(float64(x), float64(i)) == 0 {
				factorsOfx++
			}
		}
		if factorsOfx >= targetDivisors {
			result = x
			break
		}
	}
	return result
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
	test()
	result := calculateDivisors(500)
	fmt.Println(result)
}