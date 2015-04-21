package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "This, that, and the other."
	replacer := strings.NewReplacer(",", "", ".", "", ";", "")
	s = replacer.Replace(s)
	words := strings.Fields(s)
	for i, word := range words {
		fmt.Println(i, " => ", word)
	}
}
