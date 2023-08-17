package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 1) // aynı slice capacity gibi, kapasitesi 1 olan channel
	ch <- "hello world"        // gönder, bloklama yok
	fmt.Println(<-ch)          // anında görüntü

	fmt.Println("bitti")
}
