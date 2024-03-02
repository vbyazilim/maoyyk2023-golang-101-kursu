package main

import "fmt"

// SchoolNumber is a custom type definition uses int.
type SchoolNumber int

func sum[T ~int](a T, b T) T {
	return a + b
}

func main() {
	n1 := SchoolNumber(1)
	n2 := SchoolNumber(2)
	fmt.Println(sum(n1, n2))
}
