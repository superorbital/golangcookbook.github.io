package main

import (
	"fmt"
	"strings"
)

func ToSentence(words []string, andOrOr string) string {
	l := len(words)
	wordsForSentence := make([]string, l)
	copy(wordsForSentence, words)
	wordsForSentence[l-1] = andOrOr + " " + wordsForSentence[l-1]
	return strings.Join(wordsForSentence, ", ")
}

func main() {
	a := []string{"this", "that", "the other"}
	fmt.Println(strings.Join(a, ", "))
	fmt.Println(ToSentence(a, "or"))
}
