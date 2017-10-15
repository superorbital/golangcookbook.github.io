package main

import "fmt"

func reverse(numbers []int) []int {
	newNumbers := make([]int, len(numbers))
	for i, j := 0, len(numbers)-1; i <= j; i, j = i+1, j-1 {
		newNumbers[i], newNumbers[j] = numbers[j], numbers[i]
	}
	return newNumbers
}

func main() {
	array := []int{1, 2, 3, 4, 5}

	fmt.Printf("%v\n", reverse(array))
	fmt.Printf("%v\n", array)
}
