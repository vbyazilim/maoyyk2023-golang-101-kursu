package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		heartBeat := time.Tick(1 * time.Second)      // bize channel döner
		otherHeartBeat := time.Tick(5 * time.Second) // bize channel döner

		for {
			select {
			case <-heartBeat:
				fmt.Println("-> heartBeat (every second)")
			case <-otherHeartBeat:
				fmt.Println("-> otherHeartBeat (every 5 seconds)")
			}
		}
	}()

	// sonsuz döngüde "tick"
	for {
		fmt.Println("tick")
		time.Sleep(100 * time.Millisecond)
	}
}
