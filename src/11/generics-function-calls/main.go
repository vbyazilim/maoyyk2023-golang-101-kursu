package main

import "fmt"

type (
	mapperFunc func(int) int
	numbers    []int
)

func numMutator(values numbers, fn mapperFunc) numbers {
	result := make(numbers, len(values))
	for i, v := range values {
		result[i] = fn(v)
	}
	return result
}

func main() {
	input := []int{1, 2, 3, 4, 5}
	fn := func(n int) int {
		return n * 2
	}

	fmt.Println(numMutator(input, fn))
}
