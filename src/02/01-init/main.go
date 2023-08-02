package main

import "fmt"

func main() {
	fmt.Println("main fonksiyonundan merhaba!")
}

func init() {
	fmt.Println("init fonksiyonu çağırıldı")
}

// init fonksiyonu çağırıldı
// main fonksiyonundan merhaba!
