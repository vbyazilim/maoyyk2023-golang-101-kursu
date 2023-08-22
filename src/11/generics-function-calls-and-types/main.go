package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type (
	mapperFunc[T any] func(T) T
	numbers[T any]    []T
)

func numMutator[T constraints.Ordered](values numbers[T], fn mapperFunc[T]) []T {
	result := make([]T, len(values))
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
