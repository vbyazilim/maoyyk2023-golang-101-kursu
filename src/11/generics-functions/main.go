package main

import "fmt"

func sum[T int | float64](a T, b T) T {
	return a + b
}

func main() {
	fmt.Println(sum(1, 2))     // 3
	fmt.Println(sum(1.2, 2.3)) // 3.5
}
