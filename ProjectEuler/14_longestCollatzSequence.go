package main

// https://projecteuler.net/problem=14

import (
	"fmt"
	"math"
	"os"
)

func isEven(n int) bool {
	// test if a number is even
	if math.Mod(float64(n), 2) == 0 {
		return true
	}
	return false
}

func generateChain(start int) int {
	// generate a single chain and return its length
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

func longestChain(start int) {
	// generate many chains, looking for the longest
	chainRoot := 0
	longest := 0
	for i := start; i > 0; i-- {
		l := generateChain(i)
		// fmt.Println(i, ":", l)
		if l > longest {
			longest = l
			chainRoot = i
		}
	}
	fmt.Println("Chain root:", chainRoot, "length:", longest)
}

func test() {
	result := generateChain(13) 
	expectedResult := 10 // the length of chain starting from 13

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
	longestChain(1000000)
}