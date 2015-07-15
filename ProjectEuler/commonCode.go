package common

import (
	"strings"
	"strconv"
)

func castToIntArray(target string) []int {
	// used to convert a string of integers to an array of actual integers
	splitString := strings.Split(target, "")
	arr := make([]int, len(splitString))
	for i := 0; i < len(splitString); i++ {
		s, err := strconv.Atoi(splitString[i])
		check(err)
		arr[i] = s
	}
	return arr
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}