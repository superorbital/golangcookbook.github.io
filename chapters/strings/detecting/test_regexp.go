package main

import (
	"fmt"
	"regexp"
)

func main() {
	matched, err := regexp.MatchString("Wor.*", "Hello World")
	fmt.Println("Matched:", matched, "Error:", err)
	matched, err = regexp.MatchString("Good.*", "Hello World")
	fmt.Println("Matched:", matched, "Error:", err)
	matched, err = regexp.MatchString("(bad regexp", "Hello World")
	fmt.Println("Matched:", matched, "Error:", err)
}
