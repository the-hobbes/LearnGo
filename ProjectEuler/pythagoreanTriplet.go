package main

import (
	"fmt"
	"os"
)


func findTriplet(tripletSum int) ([]int) {
	// given a number, find the pythagorean triplet the sum of which equals 
	// that number. 
	return []int{-1, -1, -1}
}

func test() {
	result := findTriplet(12)
	expectedResult := []int{3, 4, 5}

	for i := 0; i < len(expectedResult); i++ {
		if result[i] != expectedResult[i] {
			fmt.Fprintf(os.Stderr, 
				"***Test failed*** \n Wanted: %i, Received: %i \n", 
				expectedResult, result)
		    os.Exit(1)
		}
	}
	fmt.Println("***Test Passed***")
}

func main() {
	test()
}