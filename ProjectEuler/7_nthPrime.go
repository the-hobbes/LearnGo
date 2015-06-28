package main

import (
	"fmt"
	"os"
	"math/big"
)

func calculateNthPrime(n int) (nthPrime int64) {
	counter := int64(2) // start at 2
	numberOfPrimes := 0
	for {
		i := big.NewInt(counter)
		// http://stackoverflow.com/questions/21398396/how-do-you-use-golangs-probablyprime
		if i.ProbablyPrime(4) == true {
			numberOfPrimes = numberOfPrimes + 1
		}
		if numberOfPrimes >= n{
			nthPrime = counter
			break
		}
		counter = counter + 1
	}
	return nthPrime
}

func test() {
	result := calculateNthPrime(6) // 6th prime?
	expectedResult := int64(13)
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

	// 10001st prime?
	n := 10001
	result := calculateNthPrime(n)
	fmt.Println(result)
}
