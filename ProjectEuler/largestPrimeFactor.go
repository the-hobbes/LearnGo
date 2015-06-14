package main

import (
	"fmt"
	"os"
)

func maxValue(arr []int) int {
	// find max value in array
	max := arr[0] // assume first element is biggest at first
	for _, value := range arr {
		if value > max {
			max = value
		}
	}
	return max
}

func fermatFactor(n int) int {
	// calculate the largest prime factor of a given number.
	// http://stackoverflow.com/questions/24166478/efficient-ways-of-finding-the-largest-prime-factor-of-a-number/24169277#24169277
	wheel := []int{1, 2, 2, 4, 2, 4, 2, 4, 6, 2, 6}
	w := 0
	f := 2
	fs  := make([]int, 3)
	for {
		if f * f > n{break}
		for {
			if n % f != 0 {break}
			fs = append(fs, f)
			n = n / f
		}
		f = f + wheel[w]
		w = w + 1
		if w == 11 {
			w = 3
		}
	}
	if n > 1 {
		fs = append(fs, n)
	}

	largestFactor := maxValue(fs)
	return largestFactor
}

func test() {
	result := fermatFactor(13195)
	expectedResult := 29
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

	input := 600851475143
	result := fermatFactor(input)
	fmt.Println(result)
}