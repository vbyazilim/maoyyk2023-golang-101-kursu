// nolint:dupword
package main

import "fmt"

func main() {
	ch := make(chan int, 4)     // Kapasitesi 4 olan buffered bir channel
	done := make(chan struct{}) // sinyalizasyon için kullanılacak bir channel, buna done channel pattern denir

	// goroutine ateşliyoruz...
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 4
		// ch <- 5 // bloklar!
		close(ch) // channel'ı kapat, artık yazılamaz
	}()

	go func() {
		// buffered channel'dan oku, yöntem 1
		// for i := 0; i < cap(ch); i++ {
		// 	fmt.Println(<-ch)
		// }

		// buffered channel'dan oku, yöntem 2
		// fmt.Println(<-ch)
		// fmt.Println(<-ch)
		// fmt.Println(<-ch)
		// fmt.Println(<-ch)
		// fmt.Println(<-ch) chan tipinint zero-value'su yani 0 gelir...
		// buffer 4 idi. biz 5.yi okumak istedik.

		// buffered channel'dan oku, yöntem 3
		for d := range ch {
			fmt.Println(d)
		}

		done <- struct{}{}
	}()

	<-done
}

// 1
// 2
// 3
// 4
