package main

import (
	"fmt"
	"os"
	"io/ioutil"
	// "strconv"
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

func largestProduct(numAdjacentDigits int) (product int, adjacentDigits []int) {
	numAdjacentDigits = -1
	product = -1
	adjacentDigits = make([]int,1)
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