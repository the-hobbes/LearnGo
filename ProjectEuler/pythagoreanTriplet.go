package main

import (
	"fmt"
	"os"
)
var globalResult []int

func sumSlice(input []int) int {
	// sum the numbers in a given int slice
	total := 0
	for _, value := range input {
		total += value
	}
	return total
}

func productSlice(input []int) int {
	// find the product of the values in a slice
	total := 1
	for _, value := range input {
		total = total * value
	}
	return total
}

func generatePermutations(numbers []int, target int, partial []int) {
	// generate all 3-number permutations of the numbers list.
	// if the sum of a 3-number permutation == the target, 
	// return the 3-number permutation.
	s := sumSlice(partial)

	if s == target { // check to see if we've found our triplet
		if len(partial) == 3 {
			a := partial[0] 
			b := partial[1]
			c := partial[2]
			if ( a < b && b < c) {
				if (a*a + b*b == c*c) {
					globalResult = partial // we've found it
				}
			}
		}
	}

	for i := 0; i < len(numbers); i++ {
		n := numbers[i]
		remaining := numbers[i+1:] // Golang slicing, i+1 to the end

		// This is a horrible hack, necessary because of how slices
		// work in Go- append keeps a pointer to the same slice, so I needed
		// to copy the slice here to get a truly new one. 

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
	// make a slice and fill it, using tripletSum arbitrarily
	numbers := make([]int, tripletSum)
	for i := 0; i < len(numbers); i++ {
		numbers[i] = i
	}
	generatePermutations(numbers, tripletSum, partial)
	return globalResult // euuuugh
}

func test() {
	resultSlice := findTriplet(12)
	expectedResultSlice := []int{3, 4, 5}

	for i := 0; i < len(expectedResultSlice); i++ {
		if resultSlice[i] != expectedResultSlice[i] {
			fmt.Fprintf(os.Stderr, 
				"***Test failed*** \n Wanted: %i, Received: %i \n", 
				expectedResultSlice, resultSlice)
		    os.Exit(1)
		}
	}

	resultProduct := productSlice(resultSlice)
	expectedResultProduct := 60
	if resultProduct != expectedResultProduct {
		fmt.Fprintf(os.Stderr, 
			"***Test failed*** \n Wanted: %i, Received: %i \n", 
			expectedResultProduct, resultProduct)
	    os.Exit(1)
	}

	globalResult = nil // reset global
	fmt.Println("***Test Passed***")
}

func main() {
	// test()
	sliceResult := findTriplet(1000)
	fmt.Println(sliceResult)
	product := productSlice(sliceResult)
	fmt.Println(product)
}