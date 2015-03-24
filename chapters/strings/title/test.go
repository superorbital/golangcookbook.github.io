package main

import (
	"strings"
	"unicode"
)

func properTitle(input string) string {
	words := strings.Fields(input)
	smallwords := " a an on the to "

	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") {
			words[index] = word
		} else {
			words[index][0] = uint8(unicode.ToTitle(rune(words[index][0])))
		}
	}
	return strings.Join(words, " ")
}

func main() {
	s := "welcome to the dollhouse!"
	println(s, " => ", strings.Title(s))
	println(s, " => ", properTitle(s))
}
