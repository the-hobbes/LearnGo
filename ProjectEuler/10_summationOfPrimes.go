// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
// Find the sum of all the primes below two million.

package main

import (
	"fmt"
	"os"
	"math/big"
)

func primeGenerator(upperBound int64) (chan int64) {
	// generates prime numbers on demand and fills a channel w/them
	channel := make(chan int64)
	go func() {
		for currentPrime := 2; ; currentPrime++ {
			i :=(big.NewInt(int64(currentPrime)))
			if i.ProbablyPrime(4) == true {
				channel <- int64(currentPrime)
			}
		}
	}()
	return channel
}

func sumPrimes(upperBound int64) int64 {
	// sums all primes found up to the upperBound
	var runningTotal int64
	generator := primeGenerator(upperBound)
	for i := int64(0); i < upperBound; i++ {
		foundPrime := <-generator
		if foundPrime > upperBound {
			break
		}
		runningTotal = runningTotal + foundPrime
	}
	return runningTotal
}

func test() {
	upperBound := int64(10)
	result := sumPrimes(upperBound)
	expectedResult := int64(17)
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
	r := sumPrimes(2000000)
	fmt.Println(r)
}
