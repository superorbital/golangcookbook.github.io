package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "root:*:0:0:System Administrator:/root:/bin/sh"
	words := strings.Split(s, ":")
	for i, word := range words {
		fmt.Println(i, " => ", word)
	}
}
