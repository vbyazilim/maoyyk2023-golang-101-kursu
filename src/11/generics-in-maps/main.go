package main

import "fmt"

// GenericMap represents generic map type.
type GenericMap[K comparable, V int | string] map[K]V

func main() {
	m := GenericMap[string, int]{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	fmt.Printf("%v\n", m) // map[one:1 three:3 two:2]
}
