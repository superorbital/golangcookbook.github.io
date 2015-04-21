package main

import (
	"fmt"
	"strings"
)

func main() {
	removePunctuation := func(r rune) rune {
		if strings.ContainsRune(".,:;", r) {
			return -1
		} else {
			return r
		}
	}

	s := "This, that, and the other."
	s = strings.Map(removePunctuation, s)
	words := strings.Fields(s)
	for i, word := range words {
		fmt.Println(i, " => ", word)
	}
}
