---
title: Reversing an Array
layout: default
---

## {{ page.title }}
{:.recipe}

How do I reverse an array?
{:.question}

Reversing an array is fairly straightforward in Go, due to multiple return values.  Simply loop through the first half of the array, swapping each element in turn with its mirror counterpart:

``` go
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
```

There's an even more terse implementation which makes use of the `,` operator inside the Go for loop:

``` go
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
```

Note that both of these `reverse` functions take slices, which are passed by value.  This is good for performance, but means we're actually modifying the passed in array, leading to unexpected results.  In other words, the `reverse` method could be written like such:

```
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
```

Now `reverse` doesn't return another reference to the slice, which makes it much more clear that it's modifying the slice in-place.  If we want to return a modified copy of the slice, we must do so manually:

```
package main

import "fmt"

func reverse(numbers []int) []int {
	newNumbers := make([]int, len(numbers))
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		newNumbers[i], newNumbers[j] = numbers[j], numbers[i]
	}
	return newNumbers
}

func main() {
	array := []int{1, 2, 3, 4, 5}

	fmt.Printf("%v\n", reverse(array))
	fmt.Printf("%v\n", array)
}
```

Running the above shows that the original array is untouched:

```
$ go run test.go
[5 4 0 2 1]
[1 2 3 4 5]
```

