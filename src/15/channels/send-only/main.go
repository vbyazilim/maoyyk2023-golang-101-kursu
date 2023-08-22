package main

import (
	"fmt"
)

func returnReceiveOnly() <-chan int {
	c := make(chan int)

	go func() {
		defer close(c)

		// bu fonksiyondan dönen channel'ı receive
		// eden her kimse, en fazla 100 tane sayı
		// receive edebilir.

		// loop bitiminde defer ile channel kapandığı
		// için, 100+ zero-value => 0 gelir...
		for i := 0; i < 100; i++ {
			c <- i
		}
	}()

	return c
}

func main() {
	r := returnReceiveOnly() // returns receive-only channel

	// r <- 10 // invalid operation: cannot send to receive-only type <-chan int

	for i := 0; i < 200; i++ {
		fmt.Println(<-r)
	}
}

// read/write-only channels are distinct types, the compiler can use its existing
// type-checking mechanisms to ensure the caller does not try to write stuff
// into a channel it has no business writing to.
