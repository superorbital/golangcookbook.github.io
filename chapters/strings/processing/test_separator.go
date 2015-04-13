package main

import (
	"fmt"
	"strings"
)

func main() {
	words := strings.Split("one:two:three", ":")
	for i, word := range words {
		fmt.Println(i, " => ", word)
	}
}
