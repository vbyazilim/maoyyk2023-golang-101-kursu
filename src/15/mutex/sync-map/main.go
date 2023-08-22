package main

import (
	"fmt"
	"sync"
)

var (
	m  sync.Map
	wg sync.WaitGroup
)

func main() {
	// 10 tane goroutine kullanarak key:i, value: i
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			m.Store(i, i)
		}(i)
	}

	wg.Wait() // goroutine'lerin işini bitirmesini bekle

	m.Store("foo", "bar") // manual olarak key ekle

	// value, ok syntactic sugar, eklediğin key'i oku

	if v, ok := m.Load("foo"); ok {
		fmt.Println("foo ->", v)
	}

	// goroutine'lerle için doldurduğun map'ten değerleri geri oku.
	for i := 0; i < 10; i++ {
		if v, ok := m.Load(i); ok {
			fmt.Printf("%d -> %v\n", i, v)
		}
	}

	fmt.Println("bitti")
}
