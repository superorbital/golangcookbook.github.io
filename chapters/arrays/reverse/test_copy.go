package main

import "fmt"

func reverse(numbers []int) []int {
	newNumbers := make([]int, 0, len(numbers))
	for i := len(numbers)-1; i >= 0; i-- {
		newNumbers = append(newNumbers, numbers[i])
	}
	return newNumbers
}

func main() {
	array := []int{1, 2, 3, 4, 5}

	fmt.Printf("%v\n", reverse(array))
	fmt.Printf("%v\n", array)
}
