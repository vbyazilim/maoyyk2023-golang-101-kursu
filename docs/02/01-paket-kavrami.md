# Bölüm 02/01: Golang Uygulamasına Genel Bakış

## Paket Kavramı

Go paketleri, kodu gruplama birimleri ve projenizi düzenlemenize yardımcı
olur. “Paket” olarak adından da anlaşılacağı gibi, tek bir birim olarak bir
veya daha fazla kaynak kod dosyasını “paket” halinde paketlemenizi sağlar. Go
paketleri, kendi veya diğer paketleri kodunuzda kullanmanıza izin vererek
yeniden kullanılabilirlik sağlar.

Go için özel bir dosya ve paket adı bulunur: `main.go` ve `package main`. Eğer
bir go projesi altında `main.go` dosyası var ise mutlaka o dosyanın paket
adı da `main` olur. Bu uygulamaya giriş yeridir. Eğer orada bir go uygulaması
var ise o uygulamanın giriş kapısı `main.go` olur.

Dizin yapısına göre bir projede birden fazla `main.go` olabilir. Dosyanın
durduğu yere göre, **sadece bir tane** main paketi olur!

Şimdi klasik **Hello World** uygulaması yapalım:

```go
package main // paket deklerasyonu

import "fmt" // koda dışarıdan dahil edilen başka bir paket

// kodun esas kısım, çalışma ilk buradan başlayacak!
func main() {
	fmt.Println("Hello World")
}
```

---

##  `main` Paketi

```go
package main

func main(){
}
```

Programın başladığı yeri ifade eder. Her zaman `main` fonksiyonundan ile
başlar. Golang modüler bir yapıya sahip olduğu için farklı farklı
fonksiyonları farklı dosyalara ya da paketlere koymak mümkündür. 

---

## `init` Fonksiyonu

Tüm kaynak kod (*içinde go kodu bulunan her dosya*) dosyalarının kendine ait
bir `init` fonksiyonu olabilir. Derleme esnasında go compiler, sırasıyla;

1. Tüm değişken/sabit tanımlamalarını derler
1. Tanımlananın değişkenlerin/sabitleri `initialize` eder
1. `import` edilen tüm paketlerin devreye alır ve gereken `initialize`
   işlerini yapar
1. `init` fonksiyonunu çağırır
1. `main` fonksiyonunu çağırır

https://go.dev/play/p/i_JqpCRI4nh

```go
package main

import "fmt"

func main() {
	fmt.Println("main fonksiyonundan merhaba!")
}

func init(){
	fmt.Println("init fonksiyonu çağırıldı")
}

// init fonksiyonu çağırıldı
// main fonksiyonundan merhaba!
```

---

## Paket Kapsamı (package scope)

Go’da hiçbir şey **global** olarak tanımlanamaz, tüm değişkenler, sabitler,
fonksiyonlar mutlaka paket kapsamı içindedir. Yani yazılan her şey mutlaka
bir pakete aittir. Kod paketler içinde yaşar.

```go
package main

import "fmt" // fmt paketi yüklendi

func main() {
	fmt.Println("Merhaba") // fmt paketinden Println fonksiyonu çağırıldı
}
```

Fonksiyonun adı `Println` ve ilk harfi büyük harf `P`. Bunun sebebi, `fmt`
paketindeki `Println` fonksiyonu dış dünyaya açık, yani başka bir paketten
`import` edilip kullanılabilir (bu örnekte başka paket bizim main paketi).
Bu duruma `Println` fonksiyonun **Exportable** olduğunu gösterir.

Eğer bir fonksiyon, değişken, sabit ya da bir tip adı büyük harf ile başlıyorsa
bu diğer paketler tarafından kullanıldığının işaretidir.

Örneğin şöyle bir proje/dizin yapısı olsa:

    .
    ├── codeutils
    │   └── codeutils.go
    ├── go.mod
    └── main.go

ve `main.go`:

```go
package main

import (
	"demo/codeutils"
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	codeutils.PrintString("Hello World")
}
```

`codeutils.go` de şöyle olsa;

```go
package codeutils

import "fmt"

// PrintString prints given string.
func PrintString(s string) {
	printString(s)
}

func printString(s string) {
	fmt.Println(s)
}
```

`PrintString` fonksiyonu `codeutils` paketinde bulunan **Exportable** bir
fonksiyon diyebiliriz. Dikkat ettiyseniz `printString` ise küçük harfle başlıyor
ve sadece `codeutils` paketi içinde kullanılabilen **Unexportable** ya da `private`
bir fonksiyon olarak tanımlı. Yani `main.go` içinden `codeutils.printString`
yapsanız kodu derleyemezsiniz, hata alırsınız.


