package main

import (
	"fmt"
	"os"
	"strconv"
)

func reverseInteger(n int) int {
	var result string
	s := strconv.Itoa(n) // cast int to string

	for _,v := range s {
		result = string(v) + result
	}

	i, err := strconv.Atoi(result) // cast string to int
    if err != nil {
        // handle error
        fmt.Println(err)
        os.Exit(2)
    }
	return i
}

func isPalindrome(n int) bool {
	reversed := reverseInteger(n)
	if n == reversed {
		return true
	}
	return false
}

func findPalindrome(digits int) int {
	var max1, max2 int
	if digits == 2 {
		max1 = 99
		max2 = 99		
	} else {
		max1 = 999
		max2 = 999
	}

	for i := max1; i > 0; i-- {
		for j := max2; j > 0; j-- {
			product := i * j
			if isPalindrome(product) {
				return product
			}
		}
	}
	return -1
}

func test() {
	expectedResult := 9009
	result := findPalindrome(2)
	if result != expectedResult {
		fmt.Fprintf(os.Stderr, 
			"***Test failed*** \n Wanted: %d, Received: %d \n", 
			expectedResult, result)
	    os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "***Test Passed*** \n Wanted: %d, Received %d \n",
		expectedResult, result)
}

func main() {
	// Find the largest palindrome made from the product of two 3-digit numbers
	test()
	palindrome := findPalindrome(3) // really run for problem
	fmt.Println(palindrome)
}