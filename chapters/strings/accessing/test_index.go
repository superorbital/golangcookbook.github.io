package main

import "fmt"

func main() {
	s := "Hello World"
	fmt.Println(s[:5])           // the first 5 characters
	fmt.Println(s[:len(s)-6])    // the same, but counting back from the end
	fmt.Println(s[1:])           // all except the first character
	fmt.Println(s[len(s)-5:])    // the last 5 characters
	fmt.Println(s[1 : len(s)-1]) // all except the first and last
}
