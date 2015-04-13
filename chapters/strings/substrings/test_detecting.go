package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Assume", "Ass"))

	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("Hello", "aeiou&y"))

	fmt.Println(strings.HasPrefix("Hello World", "Hello"))
	fmt.Println(strings.HasSuffix("Hello World", "World"))
}
