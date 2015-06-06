package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := "[one][two][three]"
	words := regexp.MustCompile("a*").Split(s, -1)
	if words != nil {
		for i, word := range words {
			fmt.Println(i, " => ", word)
		}
	}
}
