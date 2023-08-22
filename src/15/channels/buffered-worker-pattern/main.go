package main

import "fmt"

const (
	workers   = 10 // üretenler
	consumers = 20 // tüketenler
)

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// go worker(jobs, results)

	// go worker(jobs, results)
	// go worker(jobs, results)
	// go worker(jobs, results)

	// 10 tane worker tetikliyoruz
	for i := 0; i < workers; i++ {
		go worker(jobs, results)
	}

	// 100 kapasitesi var buffered channel
	for i := 0; i < cap(jobs); i++ {
		jobs <- i
	}
	close(jobs)

	// 20'li 20'li tüketiyoruz
	for i := 0; i < consumers; i++ {
		fmt.Println(<-results)
	}
	close(results)
}

// jobs: send only channel
// results: receive only channel
func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- task(n)
	}
}

func task(n int) int {
	pow := n * n
	fmt.Println("n=", n, ",n * n=", pow)
	return pow
}
