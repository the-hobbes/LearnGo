package main

import (
	"fmt"
	"os"
)

func calculateSquareOfSums(naturalNumbers int) (square int) {
	sum := 0
	for i := 0; i <= naturalNumbers; i++ {
		sum = sum + i
	}
	square = sum * sum
	return square
}

func calculateSumOfSquares(naturalNumbers int) (sum int) {
	for i := 0; i <= naturalNumbers; i++ {
		sum = sum + (i * i)
	}
	return sum
}

func sumSquareDifference(naturalNumbers int) (diff int) {
	sumOfSquares := calculateSumOfSquares(naturalNumbers)
	squareOfSums := calculateSquareOfSums(naturalNumbers)
	diff = squareOfSums - sumOfSquares
	return diff
}

func test() {
	result := sumSquareDifference(10)
	expectedResult := 2640
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

	naturalNumbers := 100
	result := sumSquareDifference(naturalNumbers)
	fmt.Println(result)
}