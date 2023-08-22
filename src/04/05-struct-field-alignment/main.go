package main

import "fmt"

// Bad represents bad struct field alignment for training purpose.
type Bad struct {
	Field1 bool
	Field2 int64
	Field3 bool
	Field4 float64
	Field5 []bool
}

func main() {
	b := Bad{}

	fmt.Printf("%+v\n", b)
}
