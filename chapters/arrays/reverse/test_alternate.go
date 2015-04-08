package main

import "fmt"

func reverse(numbers []int) {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
}

func main() {
	array := []int{1, 2, 3, 4, 5}
	reverse(array)

	fmt.Printf("%v\n", array)
}
