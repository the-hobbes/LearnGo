package main

// https://projecteuler.net/problem=13

import (
	"fmt"
	"io/ioutil"
)

var INPUT string // input global var

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func readInFile() string {
	// TODO(pheven): we will want to process the file line by line, so
	// use the following to dothat instead of reading the whole thing in 
	// at once: http://stackoverflow.com/questions/8757389/reading-file-line-by-line-in-go
	dat, err := ioutil.ReadFile("largeSumInput")
    check(err)

    return string(dat)
}

func largeSum() {
	return
}

func main() {
	r := largeSum()
	fmt.Println(r)
}