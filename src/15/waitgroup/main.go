package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	urls := []string{
		"https://httpbin.org/delay/2",
		"https://kamp.linux.org.tr/2023-yaz/",
		"https://github.com/",
		"https://ugur.ozyilmazel.com/",
		"https://vigo.io",
		"https://fooo-fake-nonurl-xxxxxxxxx.com.tr",
	}

	wg.Add(len(urls))

	for _, url := range urls {
		// wg.Add(1) // <-- bu şekilde de olabilirdi...
		go func(url string) {
			res, err := http.Get(url) // nolint
			if err == nil {
				fmt.Println(url, res.Status)
			} else {
				fmt.Println(url, err)
			}

			wg.Done()
		}(url) // <- dışarıdaki url’i goroutine’ne geçiyoruz! yani her goroutine’e doğru değer!
	}

	wg.Wait()
}

// https://fooo-fake-nonurl-xxxxxxxxx.com.tr Get "https://fooo-fake-nonurl-xxxxxxxxx.com.tr": dial tcp: lookup fooo-fake-nonurl-xxxxxxxxx.com.tr: no such host
// https://vigo.io 200 OK
// https://ugur.ozyilmazel.com/ 200 OK
// https://github.com/ 200 OK
// https://kamp.linux.org.tr/2023-yaz/ 523
// https://httpbin.org/delay/2 200 OK
