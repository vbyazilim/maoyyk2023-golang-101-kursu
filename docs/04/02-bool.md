# Bölüm 04/01: Veri Tipleri

## Booleans

`true` ve `false` değerleri için kullanılır. `1 bit integer` olarak ifade edilir.

https://go.dev/play/p/I7zJoDNQDpI

```go
package main

import "fmt"

func main() {
	var result bool
	
	fmt.Printf("%t\n", result) // false, initial value
	
	if 2 > 1 {
		result = true // evet, 2 büyüktür 1
	}

	fmt.Printf("%v\n", result) // value’su: true
	fmt.Printf("%t\n", result) // boolean olarak value’su: true
}
```

Özetle `true` ve `false` aslında birer sabittir. Mantıksal karşılaştırma yapmak
için `&&` ve `||` kullanılır. Eşitlik için `==`, eşit değildir için `!=`, olumsuzluk
yani **NOT** için `!` kullanılır;

https://go.dev/play/p/8JYEBnj86FU

```go
package main

import "fmt"

func main() {
	fmt.Printf("true && true -> %t\n", true && true)
	fmt.Printf("true && false -> %t\n", true && false)
	fmt.Printf("false && false -> %t\n", false && false)
	fmt.Printf("false && true -> %t\n", false && true)

	fmt.Printf("true || true -> %t\n", true || true)
	fmt.Printf("true || false -> %t\n", true || false)
	fmt.Printf("false || false -> %t\n", false || false)
	fmt.Printf("false || true -> %t\n", false || true)

	fmt.Printf("!true -> %t\n", !true)
	fmt.Printf("!false -> %t\n", !false)
}
```