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
	input := []float64{1.2, 2.3, 3.4, 4.5, 5.6}
	fn := func(n float64) float64 {
		return n * 2
	}

	fmt.Println(numMutator(input, fn))
	// [2.4 4.6 6.8 9 11.2]
}
