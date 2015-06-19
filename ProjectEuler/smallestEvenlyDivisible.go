package main

import(
	"fmt"
	"os"
	"math"
)

func isEventlyDisvisible(number, rangeStart, rangeStop int) bool {
	// test if the number passed in is evenly divisible by all the numbers in the range
	for i := rangeStart; i <= rangeStop; i++ {
		result := math.Remainder(float64(number), float64(i))
		if result != 0 { return false }
	}
	return true
}

func smallestEvenlyDivisible(maxNumber, rangeStart, rangeStop int) int{
	for i := maxNumber; i > 0; i-- {
		divisible := isEventlyDisvisible(i, rangeStart, rangeStop)
		if divisible == true {
			return i
		}
	}
	return -1
}

func test(maxNumber int) {
	result := smallestEvenlyDivisible(maxNumber, 1, 10)
	expectedResult := 2520
	if result != expectedResult {
		fmt.Fprintf(os.Stderr, 
			"***Test failed*** \n Wanted: %i, Received: %i \n", 
			expectedResult, result)
		os.Exit(1)
	}
	fmt.Println("***Test Passed***")
}

func main() {
	maxNumber := 5000  // start at some arbitrary high number
	start := 1
	stop := 20
	test(maxNumber) // known basecase works?
	res := smallestEvenlyDivisible(maxNumber, start, stop)
	for {
		if res != -1 { break }
		maxNumber = maxNumber * 2 // double each time we fail to fine a solution
		res = smallestEvenlyDivisible(maxNumber, start, stop)
		fmt.Println(maxNumber) // how big of a number are we getting to?
	}
	fmt.Println("Solution is: ")
	fmt.Println(res)
}