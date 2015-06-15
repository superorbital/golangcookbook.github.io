package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := "[one][two][three]"
	matches := regexp.MustCompile(`\[(.*?)\]`).FindAllStringSubmatch(s, -1)
	if matches == nil {
		fmt.Println("No matches found.")
		return
	}

	for i, match := range matches {
		full := match[0]
		submatches := match[1:len(match)]
		fmt.Printf("%v => \"%v\" from \"%v\"\n", i, submatches[0], full)
	}
}
