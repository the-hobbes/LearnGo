package main

import (
	"fmt"
	"os"
)

func sumSlice(input []int) int {
	// sum the numbers in a given int slice
	total := 0
	for _, value := range input {
		total += value
	}
	return total
}

func generatePermutations(numbers []int, target int, partial []int) {
	// generate all 3-number permutations of the numbers list.
	// if the sum of a 3-number permutation == the target, 
	// return the 3-number permutation.
	s := sumSlice(partial)
	if s == target {
		fmt.Println("Target found, numbers are:")
		fmt.Println(partial)
	}

	for i := 0; i < len(numbers); i++ {
		n := numbers[i]
		remaining := numbers[i+1:] // Golang slicing, i+1 to the end

		// This is a horrible hack, necessary because of how slices
		// work in Go. Explanation at the bottom of this file.

		// make a new array of a size one larger than the current size
		// of "partial"
		newPartial := make([]int, len(partial) + 1)
		// copy all of the original "partial" into the new array
		copy(newPartial[0:len(partial)], partial[:])
		// set the last element of the new array to be n
		newPartial[len(newPartial)-1] = n
		generatePermutations(remaining, target, newPartial)
	}
}

func findTriplet(tripletSum int) ([]int) {
	// kind of a wrapper for permutation function
	// given a number, find the pythagorean triplet such that its 
	// sum equals the given number. note that a < b < c.
	var partial []int
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// result := generatePermutations(numbers, tripletSum, partial)
	generatePermutations(numbers, tripletSum, partial)
	return nil
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
	// test()
	findTriplet(12)
}