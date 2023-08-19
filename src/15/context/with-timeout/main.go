package main

import (
	"context"
	"fmt"
	"time"
)

const timeout = 1 * time.Millisecond // 1 mili saniye

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	select {
	case <-time.After(1 * time.Second): // time.After geriye channel döner
		fmt.Println("1 saniye sonra...")
	case <-ctx.Done():
		fmt.Println("timeout!!!", ctx.Err()) // context deadline exceeded
	}
}

// timeout!!! context deadline exceeded
