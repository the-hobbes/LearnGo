// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
// Find the sum of all the primes below two million.

package main

import (
	"fmt"
	"os"
	"math/big"
)

func primeGenerator(upperBound int64) (chan int64) {
	// generates a prime number, on demand and in order
	channel := make(chan int64)
	// currentPrime := int64(2)
	go func() {
		// for currentPrime := int64(2); ; currentPrime++{
		// 	i := big.NewInt(currentPrime)
		// 	if i.ProbablyPrime(4) == true {
		// 		channel <- currentPrime
		// 		break
		// 	}
		// }
		for currentPrime := 2; ; currentPrime++ {
			i :=(big.NewInt(int64(currentPrime)))
			if i.ProbablyPrime(4) == true {
				channel <- int64(currentPrime)
			}
		}
	}()
	return channel
}

func sumPrimes(upperBound int64) int {
	return -1
}

func test() {
	upperBound := int64(10)
	result := sumPrimes(upperBound)
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
	upperBound := int64(10)
	generator := primeGenerator(upperBound)
	for i := 0; i < 10; i++ {
		res := <-generator
		if res > upperBound {
			break
		}
		fmt.Println(int(res))
	}
	
}
