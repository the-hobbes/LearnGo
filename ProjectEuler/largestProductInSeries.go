package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strconv"
)

var INPUT string // input global var

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func readInFile() string {
	dat, err := ioutil.ReadFile("dat")
    check(err)

    return string(dat)
}

func nDigitProduct(start, size int) (int, []int){
	digits := make([]int, size)
	counter := 0
	for i := 1; i <= len(digits); i++ {
		fmt.Println(counter, i)
		result, _ := strconv.Atoi(INPUT[ (start + counter) : (start + i)])
		digits = append(digits, result)
		counter = counter + 1
	}

	product := digits[0]
	for j := 1; j <= len(digits); j++ {
		product = product * digits[j]
	}

	return product, digits
}

func largestProduct(numAdjacentDigits int) (product int, adjacentDigits []int) {
	max := 0
	n := len(INPUT) - (numAdjacentDigits - 1)
	for i := 0; i < n; i++ {
		product, adjacentDigits = nDigitProduct(i, numAdjacentDigits)
		if product > max {
			max = product
		}
	}

	return product, adjacentDigits
}

func test() {
	productResult, adjacentDigitsResult := largestProduct(4)
	expectedProductResult := 5832
	expectedAdjacentDigitsResult := []int{9, 9, 8, 9}

	// test product
	if productResult != expectedProductResult {
		fmt.Fprintf(os.Stderr, 
			"***Test failed*** \n Wanted: %i, Received: %i \n", 
			expectedProductResult, productResult)
	    os.Exit(1)
	}
	// test adjacent digits
	for i := 0; i < len(expectedAdjacentDigitsResult); i++ {
		if adjacentDigitsResult[i] != expectedAdjacentDigitsResult[i] {
			fmt.Fprintf(os.Stderr, 
				"***Test failed*** \n Wanted: %i, Received: %i \n", 
				expectedAdjacentDigitsResult, adjacentDigitsResult)
		    os.Exit(1)
		}
	}
	
	fmt.Println("***Test Passed***")
}

func main() {
	INPUT = readInFile()
	test()
}