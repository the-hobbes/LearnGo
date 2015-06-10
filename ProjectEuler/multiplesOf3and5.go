package main

import (
		"fmt"
		"math"
)


func sumNumbers(total int) {

	results := make([]float64, 1)

	for i := 0; i < total; i++ {
		y := float64(i)
		
		if math.Mod(y, 3) == 0 {
			results = append(results, y)
		} else if math.Mod(y, 5) == 0 {
			results = append(results, y)
		}
	}
	// fmt.Println(results)
	sum := 0
	for i := 0; i < len(results); i++ {
		sum = sum + int(results[i])
	}
	fmt.Println(sum)
}

func main() {
	// Find the sum of all the multiples of 3 or 5 below 1000
	sumNumbers(1000)
}
