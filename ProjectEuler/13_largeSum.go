package main

// https://projecteuler.net/problem=13

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
	"sort"
)

var INPUT string // input global var

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func castToIntArray(target string) []int {
	// used to convert a string of integers to an array of actual integers
	// i, _ := strconv.Atoi(scanner.Text())
	splitString := strings.Split(target, "")
	arr := make([]int, len(splitString))
	for i := 0; i < len(splitString); i++ {
		s, err := strconv.Atoi(splitString[i])
		check(err)
		arr[i] = s
	}
	return arr
}

func largeSum() {
	// use a scanner to read in a file, line by line
	file, err := os.Open("largeSumInput")
    check(err)
    defer file.Close()
    totalSum := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
	    // fmt.Println(scanner.Text())
	    // cast string to int array
	    arr := castToIntArray(scanner.Text())
	    // sort the array
	    sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	    fmt.Println(arr)
	    // sum the 11 most significant digits (easy w/sorting)
	    sum := 0
	    for i := 0; i < 11; i++ {
	    	sum = sum + arr[i]
	    }
	    totalSum = totalSum + sum
	    fmt.Println(totalSum)
	    // http://math.stackexchange.com/questions/184397/10-most-significant-digits-of-the-sum-of-a-100-50-digit-numbers
	}

	if err := scanner.Err(); err != nil {
	    log.Fatal(err)
	}
}

func main() {
	largeSum()
}