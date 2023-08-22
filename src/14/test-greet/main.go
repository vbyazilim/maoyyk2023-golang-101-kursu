package main

import (
	"fmt"

	"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet"
)

func main() {
	fmt.Println(greet.SayHi())       // hi everybody!
	fmt.Println(greet.SayHi("uğur")) // hi uğur!
	fmt.Println(greet.SayHi("uğur", "erhan"))
	// hi uğur!
	// hi erhan!
}
