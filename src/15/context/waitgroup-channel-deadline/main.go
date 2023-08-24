package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	// 300 milisaniyeden büyük işlemler cancel olacak

	defer cancel() // kullanılan kaynakları temizle (cleanup)

	var wg sync.WaitGroup

	workersAmount := 10 // goroutine sayısı
	taskQueue := make(chan int)

	// 10 tane goroutine kullanarak üretilen (produce) mesajları işleyeceğiz
	// random sürede çalışan http isteğini simüle etmeye çalışıyoruz.
	for i := 0; i < workersAmount; i++ {
		wg.Add(1)

		go func(workerID int) {
			defer wg.Done()

			for {
				select {
				case msg, ok := <-taskQueue:
					if !ok {
						fmt.Println("(closed) - workerID", workerID)
						return
					}
					longRunningHTTPOperation(ctx, workerID, msg)
				case <-ctx.Done():
					fmt.Println("---> (timeout) - workerID", workerID)
					return
				}
			}
		}(i)
	}

loop:
	// producer, 100 tane mesaj üret, kuyruğa ekle
	for i := 0; i < 100; i++ {
		select {
		case taskQueue <- i:
		case <-ctx.Done():
			fmt.Printf("---> (timeout/cancel) mesaj: %d\n", i)
			break loop
		}
	}

	close(taskQueue) //
	wg.Wait()        //

	fmt.Println("bitti")
}

func getRndDurationInMillisecond(n int64) time.Duration {
	randomInt, _ := rand.Int(rand.Reader, big.NewInt(n))
	rndInt := int(randomInt.Int64()) + 1
	return time.Duration(rndInt) * time.Millisecond
}

func longRunningHTTPOperation(ctx context.Context, workerID int, msg int) {
	dur := getRndDurationInMillisecond(1000)
	fmt.Println("-> (sending ?) - workerID", workerID, "mesaj", msg, "süre", dur)
	sleepWithContext(ctx, dur)
	fmt.Println("(sent) - workerID", workerID, "mesaj", msg, "süre", dur)
}

func sleepWithContext(ctx context.Context, d time.Duration) {
	select {
	case <-ctx.Done():
		return
	case <-time.After(d):
		return
	}
}
