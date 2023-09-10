# Bölüm 14/04: Test

## Benchmarking

Yazdığımız kod ne kadar hızlı çalışıyor? Bunu ölçmek için yaptığımız işleme
**Benchmarking** deniyor. Go, yine standart kütüphanesinden gelen araçlar
yardığımıyla bu analizi yapmamızı sağlıyor.

Test için `*testing.T` kullanırken, benchmarking için `*testing.B` kullanıyor
olacağız. Şimdi örnek paketimize bakalım.

[Örnek](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/14/test-benchmarking)

Bir string’i terse çevirmek için iki fonksiyonumuz var. Biri `Reverse` diğeri
de `ReverseVigo`. İkisi de kendilerine verilen string’i terse çeviriyor. Arada
ufak nüans farkları var. `Reverse` fonksiyonunu go’nun core developer’lardan
sevgili [Russ Cox](https://github.com/rsc) diğerini ise ben yazdım :)

```go
package stringutils

// Reverse reverses given string
// by Russ Cox - https://groups.google.com/g/golang-nuts/c/oPuBaYJ17t4/m/PCmhdAyrNVkJ
func Reverse(s string) string {
	r := make([]rune, len(s))

	n := 0
	for _, c := range s {
		r[n] = c
		n++
	}

	r = r[0:n]
	for i := 0; i < n/2; i++ {
		r[i], r[n-1-i] = r[n-1-i], r[i]
	}

	return string(r)
}

// ReverseVigo reverses given string too! a little buggy!
func ReverseVigo(s string) string {
	ss := make([]rune, len(s))

	for i, c := range s {
		ss[len(s)-1-i] = c
	}
	return string(ss)
}
```

Şimdi bu iki fonksiyonun performansını ölçelim:

```bash
$ go test -bench . -run none -benchtime 3s github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-benchmarking/stringutils
goos: darwin
goarch: arm64
pkg: github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-benchmarking/stringutils
BenchmarkReverse-10        	24825723	       128.4 ns/op
BenchmarkReverseVigo-10    	26858864	       128.6 ns/op
PASS
ok  	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-benchmarking/stringutils	7.511s
```

Komutun açıklaması:

- `-run none`: testlerden adında **none** geçenleri çalıştır, adı **none**
  olan test olmadığı için otomatik olarak testler çalışmadan direk benchmark’a
  geç!
- `-bench .`: aslında `-bench <regex>` ve `.` **any** anlamında, tüm
  `Benchmark<XXX>` fonksiyonları çalıştır
- `-benchtime 3s`: default çalışma süresi `1s` (saniye), biz `3s` saniye çalıştırmasını istedik.
- `Reverse` fonksiyonunu `24.825.723` kere çalışmış
- `ReverseVigo` fonksiyonunu `26.858.864` kere çalışmış

`Reverse` fonksiyonunu, call başına (operasyon) **128.4**
nano saniye tüketmiş. Yani **nano seconds per operation**. Benim yaptığım da
fena değil, bir tık yavaş kalmış, **128.6** nano saniye sürmüş bir fonksiyon
çağrımı.

Benchmark testi yaparken, testi yaptığınız bilgisayarın minimum kaynak
tüketmesini sağlayın. İlave process (çalışan gereksiz uygulamaları kapatın)
çalışmasın, 2. monitör bağlı olmasın, ek hard-disk bağlı olmasın, hatta
mümkünse internet bağlantısı bile olmasın!

```go
var gs string

func BenchmarkReverse(b *testing.B) {
	var s string

	for i := 0; i < b.N; i++ { // b.N'in alacağı değer, o an ki resource tüketimine bağlı
		s = stringutils.Reverse("aklındaysa kapında!") // mutlaka assignment yapmalıyız aksi halde bu kısım çalışmamız olur!
	}
	gs = s // mutlaka loop dışında assignment yapılmalı.
}

func BenchmarkReverseVigo(b *testing.B) {
	var s string

	for i := 0; i < b.N; i++ {
		s = stringutils.ReverseVigo("aklındaysa kapında!")
	}
	gs = s
}
```

Yani benchmark testlerinde, mutlaka test edeceğimiz fonksiyonun return
value’sunu, loop içinde bir değişkene atamalı ve mutlaka loop dışında da
assign edip kullanmalıyız. Yani iki kere (loop içinde ve dışında) bu atamayı
yapmamız gerekiyor!

Aksi halde; mesela;

```go
func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ { // b.N'in alacağı değer, o an ki resource tüketimine bağlı
		// stringutils.Reverse("aklındaysa kapında!")
         // _ = stringutils.Reverse("aklındaysa kapında!")
         // gibi...
	}
}
```

olsa, bu hiçbir anlamı olmayan, milyon kere loop yapan bir döngü olur sadece...

Şimdi bu fonksiyonların hafıza tüketimi durumuna bakalım:

```bash
$ go test -bench . -run none -benchtime 3s -benchmem github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-benchmarking/stringutils
goos: darwin
goarch: arm64
pkg: github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-benchmarking/stringutils
BenchmarkReverse-10        	24623949	       126.4 ns/op	     120 B/op	       2 allocs/op
BenchmarkReverseVigo-10    	28291470	       127.9 ns/op	     128 B/op	       2 allocs/op
PASS
ok  	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-benchmarking/stringutils	8.219s
```

- Yeni bir argüman `-benchmem` ekledik.
- `Reverse` fonksiyonu operasyon başına **120 bytes** tüketmiş ve HEAP’e 2 obje kaçmış!
- `ReverseVigo` fonksiyonu operasyon başına **128 bytes** tüketmiş ve HEAP’e 2 obje kaçmış!

String’leri toplamak performanslı bir yöntem;

```go
s := "hello"
s += " world"
```

gibi yapılması `fmt.Sprintf`’de daha performanslı demiştim:

```go
package stringutils_test

import (
	"fmt"
	"testing"
)

func BenchmarkSprintConcat(b *testing.B) {
	b.Run("sprint", benchSprint) // sub test gibi, sub benchmark!
	b.Run("concat", benchConcat)
}

func benchSprint(b *testing.B) {
	var s string

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = fmt.Sprint("hello") // nolint:gosimple
	}

	gs = s // bunu yapmazsak allocation
}

func benchConcat(b *testing.B) {
	var s string

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = "hello" + "world"
	}

	gs = s
}
```

Çalıştıralım;

```bash
$ go test -bench . -run none -benchtime 3s -benchmem github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-benchmarking/stringutils
goos: darwin
goarch: arm64
pkg: github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-benchmarking/stringutils
BenchmarkSprintConcat/sprint-10         	108456262	        33.05 ns/op	       5 B/op	       1 allocs/op
BenchmarkSprintConcat/concat-10         	1000000000	         0.3113 ns/op	       0 B/op	       0 allocs/op
BenchmarkReverse-10                     	28682458	       125.8 ns/op	     120 B/op	       2 allocs/op
BenchmarkReverseVigo-10                 	28101408	       128.2 ns/op	     128 B/op	       2 allocs/op
PASS
ok  	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-benchmarking/stringutils	14.461s
```

- `fmt.Sprint()` 3sn içinde **108.456.262** kere çalışmış... 1 operasyon 33.05
  ns, operasyon başına **5 bytes** ve operasyon başına 1 adet escape!
- Diğeri 3sn içinde **1.000.000.000** kere çalışmış :))))))) - 1 operasyon 0
  nano saniyeye yakın, 0 bytes allocation, 0 escape!

Yani **string concat** işi akıllara zarar verecek derecede hızlı, ayak izi
bırakmıyor ve **HEAP**’e çıkmadan **STACK**’de kalıyor!

Benchmark testleri de aynı diğer testleri çağırdığımız gibi çalışabiliyor:

```bash
$ go test -run none -bench /sprint -benchtime 3s -benchmem github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-benchmarking/stringutils
goos: darwin
goarch: arm64
pkg: github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-benchmarking/stringutils
BenchmarkSprintConcat/sprint-10         	107537920	        33.23 ns/op	       5 B/op	       1 allocs/op
PASS
ok  	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-benchmarking/stringutils	6.983s

$ go test -run none -bench /concat -benchtime 3s -benchmem github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-benchmarking/stringutils
goos: darwin
goarch: arm64
pkg: github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-benchmarking/stringutils
BenchmarkSprintConcat/concat-10         	1000000000	         0.3154 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-benchmarking/stringutils	0.518s
```

---

## Escape Analysis

Kodumuz içinde **HEAP**’e kaçanları nasıl buluruz? Bu işleme 
**Escape Analysis** yani kaçış analizi deniyor. Hemen örneğe bakalım:

[Örnek](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/14/test-escape-analysis)

https://go.dev/play/p/umeDnd8Iwdz

```go
package main

import "fmt"

// User holds user data.
type User struct {
	Email    string
	FullName string
}

// NewUserAsValue creates new User instance as value semantics.
func NewUserAsValue(email, fullName string) User {
	return User{email, fullName}
}

// NewUserAsPointer creates new User instance as pointer semantics.
func NewUserAsPointer(email, fullName string) *User {
	return &User{email, fullName}
}

func main() {
	u1 := NewUserAsValue("vigo@me.com", "Uğur Özy")
	u2 := NewUserAsPointer("vigo@me.com", "Uğur Özy")

	fmt.Println("u1 -> value", u1)   // u1 -> value {vigo@me.com Uğur Özy}
	fmt.Println("u2 -> pointer", u2) // u2 -> pointer &{vigo@me.com Uğur Özy}
}
```

İki fonksiyon var. Biri **value semantics** diğeri de **pointer semantics**
ile, biri değeri dönerken, diğeri de pointer’ı dönüyor. Kodun her zaman
**STACK**’de kalması bizim için iyi çünkü hız, performans açısından önemli.

Ne zaman **HEAP**’e geçerse, o zaman **garbage collection** devreye giriyor ve
bu performans/hız kaybına sebep oluyor.

Şimdi acaba HEAP’e kaçan bir şey var mı? diye bakıyoruz. Bunun için `-gcflags`
parametresini kullanarak kodu `build` ediyoruz. `-m` ise 
**print optimization decisions** anlamında.

Şimdi bu analizi yapabilmek için kodu derleyelim:

```bash
$ cd src/14/test-escape-analysis/
$ go build -gcflags="-m"
# escapedemo
./main.go:12:6: can inline NewUserAsValue
./main.go:17:6: can inline NewUserAsPointer
./main.go:22:22: inlining call to NewUserAsValue
./main.go:23:24: inlining call to NewUserAsPointer
./main.go:25:13: inlining call to fmt.Println
./main.go:26:13: inlining call to fmt.Println
./main.go:12:21: leaking param: email to result ~r0 level=0
./main.go:12:28: leaking param: fullName to result ~r0 level=0
./main.go:17:23: leaking param: email
./main.go:17:30: leaking param: fullName
./main.go:18:9: &User{...} escapes to heap        <--------------------- HEAP
./main.go:23:24: &User{...} escapes to heap       <--------------------- HEAP
./main.go:25:13: ... argument does not escape
./main.go:25:14: "u1 -> value" escapes to heap    <--------------------- HEAP
./main.go:25:29: u1 escapes to heap               <--------------------- HEAP
./main.go:26:13: ... argument does not escape
./main.go:26:14: "u2 -> pointer" escapes to heap  <--------------------- HEAP
```

Daha da detaylı analiz için;

```bash
$ go build -gcflags="-m -m" # daha detaylı hali...
$ go build -gcflags="-m -S" # hem escape analysis + assembly çıktısı
```

`-gcflags` parametrelerini;

```bash
$ go tool compile -help
```

ile görebiliriz.

---

## Memory ve CPU Profiling Temelleri

Şimdi https://github.com/vigo/stringutils-demo projesinde memory ve cpu
profiling yapalım.

```bash
$ cd /path/to/works
$ git clone git@github.com:vigo/stringutils-demo.git
$ cd stringutils-demo/
$ go test -run none -bench . -benchtime 3s -benchmem -memprofile m.out -cpuprofile=c.out
```

- `-memprofile` ile memory profile’ın çıktısını
- `-cpuprofile` ile cpu profile’ın çıktısını

alacağımız dosyaları belirtiyoruz. Çıktıların ikisi de **binary** dosyalar.

Eğer bilgisayarınızda `Graphviz` kurulu değilse mutlaka kurun;
`brew install graphviz` ile kurabilir.

```bash
$ go tool pprof stringutils-demo.test c.out
File: stringutils-demo.test
Type: cpu
Time: Aug 16, 2023 at 2:57pm (+03)
Duration: 3.53s, Total samples = 3.14s (89.03%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) 
```

`help` ile komutları görüntüleriz. Şimdi `list Reverse`:

    (pprof) list Reverse
    Total: 3.14s
    ROUTINE ======================== github.com/vigo/stringutils-demo.Reverse in /Users/vigo/Development/vigo/golang/libs/stringutils-demo/stringutils.go
          10ms      550ms (flat, cum) 17.52% of Total
             .          .     15:func Reverse(s string) (string, error) {
             .       50ms     16:	if !utf8.ValidString(s) {
             .          .     17:		return s, ErrInvalidUTF8
             .          .     18:	}
             .       80ms     19:	r := []rune(s)
             .          .     20:	lr := len(r)
             .       90ms     21:	ss := make([]rune, lr)
             .          .     22:
             .          .     23:	for i := 0; i < lr; i++ {
          10ms       10ms     24:		ss[lr-1-i] = r[i]
             .          .     25:	}
             .          .     26:
             .      320ms     27:	return string(ss), nil
             .          .     28:}
    ROUTINE ======================== github.com/vigo/stringutils-demo_test.BenchmarkReverse in /Users/vigo/Development/vigo/golang/libs/stringutils-demo/stringutils_test.go
             0      550ms (flat, cum) 17.52% of Total
             .          .     79:func BenchmarkReverse(b *testing.B) {
             .          .     80:	var s string
             .          .     81:	b.ResetTimer()
             .          .     82:	for i := 0; i < b.N; i++ {
             .      550ms     83:		s, _ = stringutils.Reverse("merhaba dünya!")
             .          .     84:	}
             .          .     85:
             .          .     86:	gs = s
             .          .     87:}
             .          .     88:
    (pprof) 

Şimdi görselleştirelim:

    (pprof) web

`svg` [dosyası](diagrams/pprof001.svg) oluşturur.

![pproof](diagrams/pprof001.svg)

Şimdi **memory** açısından bakalım:

```bash
$ go tool pprof -alloc_space stringutils-demo.test m.out
File: stringutils-demo.test
Type: alloc_space
Time: Aug 16, 2023 at 2:57pm (+03)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) list Reverse
Total: 2.32GB
ROUTINE ======================== github.com/vigo/stringutils-demo.Reverse in /Users/vigo/Development/vigo/golang/libs/stringutils-demo/stringutils.go
    2.32GB     2.32GB (flat, cum) 99.91% of Total
         .          .     15:func Reverse(s string) (string, error) {
         .          .     16:	if !utf8.ValidString(s) {
         .          .     17:		return s, ErrInvalidUTF8
         .          .     18:	}
         .          .     19:	r := []rune(s)
         .          .     20:	lr := len(r)
    1.69GB     1.69GB     21:	ss := make([]rune, lr)
         .          .     22:
         .          .     23:	for i := 0; i < lr; i++ {
         .          .     24:		ss[lr-1-i] = r[i]
         .          .     25:	}
         .          .     26:
  640.51MB   640.51MB     27:	return string(ss), nil
         .          .     28:}
ROUTINE ======================== github.com/vigo/stringutils-demo_test.BenchmarkReverse in /Users/vigo/Development/vigo/golang/libs/stringutils-demo/stringutils_test.go
         0     2.32GB (flat, cum) 99.91% of Total
         .          .     79:func BenchmarkReverse(b *testing.B) {
         .          .     80:	var s string
         .          .     81:	b.ResetTimer()
         .          .     82:	for i := 0; i < b.N; i++ {
         .     2.32GB     83:		s, _ = stringutils.Reverse("merhaba dünya!")
         .          .     84:	}
         .          .     85:
         .          .     86:	gs = s
         .          .     87:}
         .          .     88:
(pprof) 
```

`2.32GB` allocation var;

- `1.69GB`’ı `ss := make([]rune, lr)` ile
- `640.51MB`’ı `return string(ss)` ile

olmuş. Yani şu basit `Reverse` fonksiyonu hafıza tüketen bir canavar!

---

## Kaynaklar

- https://github.com/google/pprof/blob/master/doc/README.md
- https://mayurwadekar2.medium.com/escape-analysis-in-golang-ee40a1c064c1
- https://medium.com/a-journey-with-go/go-introduction-to-the-escape-analysis-f7610174e890
- https://faun.pub/golang-escape-analysis-reduce-pressure-on-gc-6bde1891d625
