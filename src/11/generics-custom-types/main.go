package main

import "fmt"

// SchoolNumber is a type alias.
type SchoolNumber int

func sum[T ~int](a T, b T) T {
	return a + b
}

func main() {
	n1 := SchoolNumber(1)
	n2 := SchoolNumber(2)
	fmt.Println(sum(n1, n2))
}
