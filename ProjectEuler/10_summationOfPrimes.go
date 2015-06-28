// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
// Find the sum of all the primes below two million.

package main

import (
	"fmt"
	"os"
	"math/big"
)

func primeGenerator() (chan int64) {
	// generates a prime number, on demand and in order
	channel := make(chan int64)
	
	go func() {
		for currentPrime := int64(2); ; currentPrime++ {
			i := big.NewInt(currentPrime)
			if i.ProbablyPrime(4) == true {
				channel <- currentPrime
				break
			}
		}
	}()
	return channel
}

func sumPrimes(n int) int {
	return -1
}

func test() {
	result := sumPrimes(10)
	expectedResult := 17
	if result != expectedResult {
		fmt.Fprintf(os.Stderr, 
			"***Test failed*** \n Wanted: %i, Received: %i \n", 
			expectedResult, result)
		os.Exit(1)
	}
	fmt.Println("***Test Passed***")
}

func main() {
	// test()
	for i := 0; i < 10; i++ {
		res := <-primeGenerator() // TODO fix this
		fmt.Println(res)
	}
	
}
