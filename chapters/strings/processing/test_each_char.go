package main

import "fmt"

func main() {
	for i, c := range "abc" {
		fmt.Println(i, " => ", string(c))
	}
}
