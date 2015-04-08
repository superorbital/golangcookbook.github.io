package main

import "fmt"

func reverse(numbers []int) []int {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func main() {
	fmt.Printf("%v\n", reverse([]int{1, 2, 3, 4, 5}))
	fmt.Printf("%v\n", reverse([]int{1, 2, 3, 4}))
}
