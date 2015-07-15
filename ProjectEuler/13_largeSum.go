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

func getFirst11(target string) int {
	// get the 11 most significant digits in a string, and 
	// convert that string to an integer. 
	s := target[0:12]
	i, err := strconv.Atoi(s)
	check(err)
	return i
}

func largeSum() string {
	// use a scanner to read in a file, line by line
	file, err := os.Open("largeSumInput")
    check(err)
    defer file.Close()
    sum := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
	    // cast string to int array
	    eleven := getFirst11(scanner.Text())
	    // sum each of the 11 most significant digits of each line
	    sum = sum + eleven
	    // http://math.stackexchange.com/questions/184397/10-most-significant-digits-of-the-sum-of-a-100-50-digit-numbers
	}

	if err := scanner.Err(); err != nil {
	    log.Fatal(err)
	}
	// as per the problem, we only want the top 10 digits
	strsum := strconv.Itoa(sum)
	return strsum[0:10]
}

func main() {
	result := largeSum()
	fmt.Println(result)
}