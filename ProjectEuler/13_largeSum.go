package main

// https://projecteuler.net/problem=13

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strconv"
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
}

func largeSum() {
	// use a scanner to read in a file, line by line
	file, err := os.Open("largeSumInput")
    check(err)
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
	    // fmt.Println(scanner.Text())
	    // cast string to int
	    arr := castToIntArray(scanner.Text())
	    // sort the arr
	    // sum the 11 most significant digits (easy w/sorting)
	    // http://math.stackexchange.com/questions/184397/10-most-significant-digits-of-the-sum-of-a-100-50-digit-numbers
	}

	if err := scanner.Err(); err != nil {
	    log.Fatal(err)
	}
}

func main() {
	largeSum()
}