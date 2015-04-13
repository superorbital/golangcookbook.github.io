package main

import (
	"fmt"
	"strings"
)

func main() {
	words := strings.Fields("This, that, and the other.")
	for i, word := range words {
		fmt.Println(i, " => ", word)
	}
}
