package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	str := "123.123"
	/*
		The first argument to ParseInt is the string to be converted.
		The second argument is the number of bits to which the number should fit.
	*/
	i, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%T %f\n", i, i)
}
