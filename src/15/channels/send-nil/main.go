package main

/*
Package main implements channel merge and demonstrates setting channel to nil

	https://medium.com/justforfunc/why-are-there-nil-channels-in-go-9877cc0b2308

	Original code is authored by: Francesc Campoy

*/

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)

		// a ya da b channel'ı açık olduğu sürece
		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if !ok {
					fmt.Println("a is done")
					a = nil
					continue
				}
				c <- v
			case v, ok := <-b:
				if !ok {
					fmt.Println("b is done")
					b = nil
					continue
				}
				c <- v
			}
		}
	}()
	return c
}

func produceChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			randomInt, _ := rand.Int(rand.Reader, big.NewInt(1000))

			// sanki bir işlem oluyormuş gibi...
			time.Sleep(time.Duration(int(randomInt.Int64())+1) * time.Millisecond)
		}
		close(c)
	}()
	return c
}

func main() {
	a := produceChan(1, 3, 5, 7)
	b := produceChan(2, 4, 6, 8)

	c := merge(a, b)

	for v := range c {
		fmt.Println(v)
	}
}
