package main

import "fmt"

type number interface {
	int | int8 | int16 | float32 | float64
}

func sum[T number](a T, b T) T {
	return a + b
}

func main() {
	fmt.Println(sum(int8(10), int8(2)))   // 12
	fmt.Println(sum(int16(10), int16(2))) // 12
	fmt.Println(sum(1, 2))                // 3
	fmt.Println(sum(1.2, 2.3))            // 3.5
}
