package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	timeout := time.Now().Add(3 * 1000 * time.Millisecond) // 3sn

	ctx, cancel := context.WithDeadline(context.Background(), timeout)
	defer cancel()

LOOP:
	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("1sn!")
		case <-ctx.Done():
			fmt.Println("WithDeadline", ctx.Err())
			break LOOP
		}
	}

	fmt.Println("exit")
}

// 1sn!
// 1sn!
// 1sn!
// WithDeadline context deadline exceeded
// exit
