package main

import (
	"fmt"
	"math"
	"os"
)

func isSquare(n float64) (bool) {
	return true
}

func fermatFactor(n float64) (float64) {
	a := math.Ceil(math.Sqrt(n))
	b := a * a - n
	isSquare := isSquare(b)
	for { // while true...
		a++
		b = a * a - n
		if isSquare {
			break
		}
	}
	return a - math.Sqrt(b)
}

func test(testInput float64) {
	result := fermatFactor(testInput)
	expectedResult := float64(29)
	if result != expectedResult {
		fmt.Fprintf(os.Stderr, "***Test failed*** \n Wanted: %f, Received: %f \n", expectedResult, result)
	    os.Exit(1)
	}
}

func main() {
	// calculate the largest prime factor of a given number.
	testInput := float64(13195)
	test(testInput)

	input := float64(600851475143)
	result := fermatFactor(input)
	fmt.Println(result)
}