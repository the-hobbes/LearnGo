package main

// https://projecteuler.net/problem=14

import (
	"fmt"
	"math"
	"os"
)

func isEven(n int) bool {
	if math.Mod(float64(n), 2) == 0 {
		return true
	}
	return false
}

func generateChain(start int) int {
	current := start
	length := 1
	for {
		if isEven(current) {
			current = current / 2
		} else {
			current = (current * 3) + 1
		}
		
		length++
		if current == 1 {
			break
		}
	}
	return length
}

func longestChain() {
	// generate many chains, looking for the longest
	return
}

func test() {
	result := generateChain(13)
	expectedResult := 10

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
}