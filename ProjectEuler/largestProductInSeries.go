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
	var digits []int
	counter := 0
	
	for i := 1; i <= size; i++ {
		result, _ := strconv.Atoi(INPUT[ (start + counter) : (start + i)])
		digits = append(digits, result)
		counter = counter + 1
	}

	product := digits[0]
	for j := 1; j < size; j++ {
		product = product * digits[j]
	}

	return product, digits
}

func largestProduct(size int) (max int, adjacentDigits []int) {
	max = 0
	n := len(INPUT) - (size - 1)
	for i := 0; i < n; i++ {
		product, digitSlice := nDigitProduct(i, size)
		if product > max {
			max = product
			adjacentDigits = digitSlice
		}
	}

	return max, adjacentDigits
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

	n := 13
	productResult, adjacentDigitsResult := largestProduct(n)
	fmt.Println(productResult, adjacentDigitsResult)
}