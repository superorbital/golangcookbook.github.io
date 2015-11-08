package main

import (
	"fmt"
	"os"
	"strconv"
)

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	str := "123"
	/*
		The first argument to ParseInt is the string to be converted.
		The second argument is the base to which the number is to be converted.
		The third argument is the number of bits to which the number should fit.
	*/
	i, err := strconv.ParseInt(str, 10, 64)
	handleError(err)
	fmt.Printf("%T %d\n", i, i)
	str = "invalidString123"
	_, err = strconv.ParseInt(str, 10, 64)
	handleError(err)
}
