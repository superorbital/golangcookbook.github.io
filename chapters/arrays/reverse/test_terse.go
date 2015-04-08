package main

import "fmt"

func reverse(numbers []int) []int {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func main() {
	fmt.Printf("%v\n", reverse([]int{1, 2, 3, 4, 5}))
	fmt.Printf("%v\n", reverse([]int{1, 2, 3, 4}))
}
