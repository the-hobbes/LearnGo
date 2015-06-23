package main

import (
	"fmt"
	"os"
)

func sumSlice(input []int) int {
	total := 0
	for _, value := range input {
		total += value
	}
	return total
}

func subsetSum(candidates []int, target int, partial []int) []int {
	s := sumSlice(partial)
	if s == target {
		fmt.Println("Target found, numbers are:")
		fmt.Println(partial)
	}
	if s >= target {
		return
	}
	for i := 0; i < len(candidates); i++ {
		n := candidates[i]
		remaining := numbers[i+1:] // Golang slicing?
		subsetSum(remaining, target, partial + [n]) // append instead of + [n]
	}

	return []int{-1, -1, -1}
}

func findTriplet(tripletSum int) ([]int) {
	// given a number, find the pythagorean triplet the sum of which equals 
	// that number. note that a < b < c.
	var partial []int
	candidates := []int{1,2,3,4,5,6,7,8,9,10}
	result := subsetSum(candidates, tripletSum, partial)
	
	return result
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