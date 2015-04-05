package main

import (
	"fmt"
	"strings"
)

func main() {
	a := []string{"this", "that", "the other"}
	fmt.Println(strings.Join(a, ", "))
}
