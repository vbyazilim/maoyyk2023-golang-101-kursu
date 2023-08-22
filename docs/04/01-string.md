# Bölüm 04/01: Veri Tipleri

## Strings

İçinde **Unicode** karakterler bulunan karakterler dizisidir. Tip tanımlaması
yaparken `string` anahtar kelimesi ile ifade edilir. 2 şekilde tanımlanabilir;

1. Çift Tırnak içinde: `"merhaba"`
1. Back-tick <code>\`</code> içinde: <code>\`merhaba\`</code> (**Raw String**)

https://go.dev/play/p/CqjxRuGy1ki

```go
package main

import "fmt"

func main() {
	var username string = "vigo"
	email := "foo@bar.com"
	var password string

	fmt.Println("username", username) // username vigo
	fmt.Println("email", email)       // email foo@bar.com
	fmt.Println("password", password) // password
}
```

Keza;

https://go.dev/play/p/5LIGjBHkVD9

```go
package main

import "fmt"

func main() {
	jsonRaw := `{"success": true}`

	fmt.Println("jsonRaw", jsonRaw) // jsonRaw {"success": true}
	fmt.Printf("%T\n", jsonRaw)     // string
}
```

**Raw String**’lerde back-slash yani `\` işlenmez. String’in saf halini taşır.
Çift tırnak içinde yazdıklarımıza da **Interpreted String** denir ve bu durumda
back-slash ve diğer görünmeyen karakterler `\r\n\t` gibi hepsi işlenir.

```go
`abc`                // "abc"

`\n
\n`                  // "\\n\n\\n"

"\n"

"\""                 // `"`

"Hello, world!\n"

"日本語"

"\u65e5本\U00008a9e"

// hepsi aynı çıktıyı verir...
"日本語"                                 // Unicode text girdi, çıktı: 日本語
`日本語`                                 // Unicode raw girdi, çıktı: 日本語
"\u65e5\u672c\u8a9e"                    // Unicode girdi, çıktı: 日本語
"\U000065e5\U0000672c\U00008a9e"        // Unicode girdi, çıktı: 日本語
"\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e"  // bytes girdi, çıktı: 日本語
```

Unicode desteği olduğu için aşağıdaki gibi kod çalışır;

```go
message := "Hava 42\u00B0 derece!"   // \u00B0 = °
fmt.Println(message)                 // Hava 42° derece!
```

Şunu tekrar hatırlayalım:

- `byte` aslında `uint8` (alias)
- `rune` aslında `int32` (alias)

https://go.dev/play/p/chFT6db0pbJ

```go
package main

import "fmt"

func main() {
	s := "message"

	name := "uğur"

	fmt.Println(s[0], s[1]) // 109 101

	fmt.Println(name[0], name[1], name[2], name[3], name[4]) // 117 196 159 117 114
	// name neden 5 tane karaktere sahip? 4 değil mi?

	fmt.Printf("%T %[1]v\n", name)         // string uğur
	fmt.Printf("%T %[1]v\n", []rune(name)) // []int32 [117 287 117 114] - 287 ?
	fmt.Printf("%T %[1]v\n", []byte(name)) // []uint8 [117 196 159 117 114]
}
```

Unicode desteği olduğu için aslında `ğ` byte olarak iki karekterden oluşuyor:
`196` ve `159`. Bu değerlik onluk (decimal) değerleri. Default olarak Unicode
olduğundan aslında karakter dizisindeki her eleman için `-2147483648` ile
`2147483647` arası bir değer olabiliyor. Bu bakımdan `[]int32` slice’ındaki
`287` sayısı `ğ` yi ifade ediyor.

> Eğer kod içinde string uzunluğu ile ilgili bir iş yapacaksanız bunu hep 
hatırlayın

String’ler **immutable** yani değeri değiştirilemez karakterler serisi /
koleksiyonudur. Bu ne anlama geliyor?

```go
package main

import "fmt"

func main() {
	name := "uğur"

	fmt.Printf("%v %[1]T", name[0]) // 117 uint8

	name[0] = 'x' // cannot assign to name[0] (value of type byte)
}
```

Neticede bir koleksiyon olduğu için indeks numarası ile (slice konusunda göreceğiz)
içindeki karakterlere erişilebilir:

https://go.dev/play/p/qg19JixH5bJ

```go
package main

import "fmt"

func main() {
	message := "Hava 42\u00B0 derece!"

	fmt.Println(message)

	for index := range message {
		// %c : karakter
		// %d : digit/sayı
		// %x : hexadecimal/16'lık sistem
		fmt.Printf("%c | %d | $%x\n", message[index], message[index], message[index])
	}
}

// H | 72 | $48
// a | 97 | $61
// v | 118 | $76
// a | 97 | $61
//   | 32 | $20
// 4 | 52 | $34
// 2 | 50 | $32
// Â | 194 | $c2
//   | 32 | $20
// d | 100 | $64
// e | 101 | $65
// r | 114 | $72
// e | 101 | $65
// c | 99 | $63
// e | 101 | $65
// ! | 33 | $21
```

Yine slice konusunda işleyeceğimiz kesme/biçme işlerini de yapabiliriz:

https://go.dev/play/p/pcfKTcOLsXd

```go
package main

import "fmt"

func main() {
	s := "hello world"
	//    0123456789x

	fmt.Println(s)      // hello world
	fmt.Println(s[:5])  // hello - 0'dan 5'e kadar, 5 hariç
	fmt.Println(s[6:])  // world - 6'dan sona kadar
	fmt.Println(s[2:5]) // llo   - 2'den 5'e kadar, 5 hariç

	// fmt.Println(s[:-1]) invalid slice index -1 (index must be non-negative)
}
```

**String Concatenation** yani metinleri birbirleriyle toplamak da mümkün;

https://go.dev/play/p/vIectChuJKs

```go
package main

import "fmt"

func main() {
	hello := "Hello"
	world := "World"
	
	message := hello + " " + world
	fmt.Printf("%s\n", message) // Hello World
}
```
