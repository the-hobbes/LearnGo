// Solution to the pythagorean triplet problem using channels instead of brute 
// forcing via recursion, as my previous solution did. This program creates a 
// generator and starts from the bottom, returning when it has found a solution.

package main 

import "fmt"

// struct to represent a triplet. this could also be done using []int slices
type triplet struct {
	a, b, c int
}

func (t triplet) isTriplet() bool {
	// check to see if the triplet satisfies the pythagorean triplet reqs
	return t.a + t.b + t.c == 1000 && t.a * t.a + t.b * t.b == t.c * t.c	
}

func tripletGenerator() chan triplet {
	// return a channel that kind of works like a python generator. Each time
	// it is called, it dumps a new triplet onto the channel that is then 
	// consumed and checked. The channel starts at the values of 3,2,1
	channel := make(chan triplet)
	go func() { // generate combinations of 3 numbers
		for c := 3; ; c++ {
			for b := 2; b < c; b++ {
				for a := 1; a < b; a++ { // make sure a < b
					channel <- triplet{a, b, c}
				}
			}
		}
	}()
	return channel
}

func main() {
	generator := tripletGenerator()
	for { // infinite loop to keep searching all combinations of 3 integers
		triplet := <-generator // consume a result from the generator
		if triplet.isTriplet() { // meets requirements?
			fmt.Println(triplet.a * triplet.b * triplet.c)
			break // stop iterating
		}
	}
}