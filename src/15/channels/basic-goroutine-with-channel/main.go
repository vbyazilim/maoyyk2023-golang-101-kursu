package main

import (
	"fmt"
)

func main() {
	ch := make(chan bool) // channel holds bool type!

	go func() {
		fmt.Println("hello from goroutine!")
		ch <- true // writing to ch channel
	}()

	<-ch // channel'dan veri gelene kadar blok!, reading from ch channel
	fmt.Println("exit!")
}
