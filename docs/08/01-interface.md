# Bölüm 08/01: `interface`

Go’nun en önemli konuları nedir diye sorsanız cevap olarak iki konu var derim:

1. Interface
2. Concurrency

Bir işin nasıl yapılacağını, bir davranışı belirleyen şeydir `interface`.
Yazdığımız kodu test edebilmek için `interface`’lere ihtiyaç duyarız.

Hemen küçük bir örnek yapalım. **type definition** konusunu gördük;

```go
package main

import "fmt"

// Status represents custom status.
type Status string

// Custom status codes.
const (
	StatusOK    Status = "OK"
	StatusERROR Status = "ERROR"
)

func main() {
	fmt.Println(StatusOK) // OK
}
```

`fmt.Print` familyası, verilen şeyi print edeceği zaman şuna bakar, acaba
gelen argümanın tipi, `Stringer` `interface`’ini **satisfy** ediyor mu? Yani;
pas edilen tip ne ise, `Stringer`’da tanımlanan davranışlardan `String()`
metotuna sahip mi?

```go
type Stringer interface {
	String() string
}
```

Herhangi bir tipin `Stringer`’ı tatmin edebilmesi için, mutlaka `String()`
diye bir metotu olmalı ver geriye `string` dönmeli. Davranıştan kastedilen şey
bu. Bu bakımdan `interface` tanımı yapılırken mutlaka ilgili `interface`’in
adı `er` eki ile biter (İngilizce).

```go
type error interface {
	Error() string
}
```

Her kim ki `Error()` diye metotu olup geriye `string` döner, o artık `error`
olarak kullanılabilir.

```go
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}
```

Örneğe geri dönelim, şimdi bizim tipimiz de `Stringer`’ı tatmin etsin;

https://go.dev/play/p/Qxkj_AxtdbA

```go
package main

import "fmt"

// Status represents custom status.
type Status string

func (o Status) String() string {
	return "Status is: " + string(o)
}

// Custom status codes.
const (
	StatusOK    Status = "OK"
	StatusERROR Status = "ERROR"
)

func main() {
	fmt.Println(StatusOK) // Status is: OK
}
```

---

## Empty Interface

Rob Pike ne demişti?

> Empty interface (interface{}) says nothing

`Reader` diye bir `interface` ismi duyduğumuzda aklımıza şu gelmeli:

> Hmmm, demek ki `Read` diye bir metotu var

ya da `ReadCloser` diye bir `interface` ismi duyduğumuzda;

> Kesin `Read` ve `Close` diye metotları var

Go geleneği olarak, `FooBarer` -> `ReadCloser`, `WriteCloser` gibi ifade
edilir. Birazdan göreceğiz, aynı struct’lar gibi `interface`’lerde bir biri
içine gömülebiliyor:

```go
// FooBarBazer :)
type ReadSeekCloser interface {
	Reader
	Seeker
	Closer
}
```

Dolayısıyla, `Reader` şu şekilde olsa;

```go
type Reader interface {}
```

Bundan ne anlarız? Hiçbir metotu olmayan, **boş bir interface**! Bize
söylediği bir şey var mı? bir metot? davranış? **yok**... Peki hemen şu soru
gelmeli, acaba hangi tipler bu interface’i tatmin (satisfy) edebilir? Cevap
tüm tipler, sonradan tanımlı tipler, her şey! Neden? Bu interface’in bize
sunduğu, bizim implemente etmemiz (geliştirmemzi, yazmamız) gereken **hiç bir
metotu** yok!

Boş `interface` somut bir tip değildir (concrete type)

Go versiyon `1.18` ile hayatımıza `any` diye bir tip girdi, aslında;

```go
type any = interface{}
```

**empty interface** için **syntactic sugar** (yani küçük bir kolaylık,
güzellik); istersek halen `interface{}` şeklinde de kullanabiliriz.

Şimdi `any` ve `interface{}` kullanarak `greet` fonksiyonunu tekrar yazalım:

https://go.dev/play/p/86jGL4zv6_T

```go
package main

import "fmt"

// t aslında tipsiz :) ne gelirse onun için sorun yok
func greet(t interface{}) string {
	return fmt.Sprintf("merhaba! %v", t) // t'in value presentation'ı
}

// t aslında tipsiz :) ne gelirse onun için sorun yok
func greetAny(t any) string {
	return fmt.Sprintf("any - merhaba! %v", t)
}

func main() {
	fmt.Println(greet("hello")) // string
	// merhaba! hello

	fmt.Println(greet(1)) // integer
	// merhaba! 1

	fmt.Println(greet(3.14)) // float
	// merhaba! 3.14

	u := struct {
		name string
	}{
		"vigo",
	}
	fmt.Println(greet(u)) // anonymous struct
	// merhaba! {vigo}

	fmt.Println(greet([]string{"hello"})) // string slice
	// merhaba! [hello]

	fmt.Println(greet(nil)) // nil
	// merhaba! <nil>

	fmt.Println(greetAny("hello")) // string
	// any - merhaba! hello

	fmt.Println(greetAny(1)) // integer
	// any - merhaba! 1

	fmt.Println(greetAny(3.14)) // float
	// any - merhaba! 3.14

	u2 := struct {
		name string
	}{
		"vigo",
	}
	fmt.Println(greetAny(u2)) // anonymous struct
	// any - merhaba! {vigo}

	fmt.Println(greetAny([]string{"hello"})) // string slice
	// any - merhaba! [hello]

	fmt.Println(greetAny(nil)) // nil
	// any - merhaba! <nil>
}
```

Şimdi akıllarda şu soru var:

> Madem böyle bir şansımız var, neden sürekli tip tanımlıyoruz?

Aslında ufak bir hile yapıyoruz. `fmt.Sprintf` içeride `reflect` paketini
kullanarak bir dizi kontroller yapıyor, acaba gelen argüman `string`’e
benziyor mu? ya da `integer`’a uygun mu? bir süre kontrolden geçiyor ve onun
sonucunda çıktıyı görüyoruz. Bu aslında çok maliyetli bir işlem:

## Tip Kontrol Mekanizması

https://go.dev/play/p/tILifiOH34j

```go
package main

import (
	"fmt"
	"io"
	"strings"
)

type (
	customInt  int
	fakeString string
)

func printByType(t interface{}) {
	// .(type) sadece switch statement içinde çalışır.
	switch j := t.(type) {
	case nil:
		fmt.Println(j, " bu nil")
	case int:
		fmt.Println(j, " bu int")
	case customInt:
		fmt.Println(j, " bu customInt")
	case io.Reader:
		fmt.Println(j, " bu io.Reader")
	case string:
		fmt.Println(j, " bu string")
	case bool, rune:
		fmt.Println(j, " bu bool ya da rune")
	default:
		fmt.Printf("%v fikrim yok: %[1]T\n", j)
	}
}

func main() {
	printByType(nil)                        // <nil>  bu nil
	printByType(1)                          // 1  bu int
	printByType(3.14)                       // 3.14 fikrim yok: float64
	printByType("hello")                    // hello  bu string
	printByType(true)                       // true  bu bool ya da rune
	printByType('a')                        // 97  bu bool ya da rune
	printByType(customInt(5))               // 5  bu customInt
	printByType(strings.NewReader("hello")) // &{hello 0 -1}  bu io.Reader
	printByType(fakeString("hello"))        // hello fikrim yok: main.fakeString
}
```

Şu örneğe bakalım:

https://go.dev/play/p/KjajVlJQ-h4

```go
package main

import "fmt"

func main() {
	var i any // interface{}
	fmt.Printf("i: %v, %[1]T\n", i)
	if i == nil {
		fmt.Println("i = nil (nil)")
	}

	i = 1
	fmt.Printf("i: %v, %[1]T\n", i)
	if i == 1 {
		fmt.Println("i = 1 (int)")
	}

	i = "hello"
	fmt.Printf("i: %v, %[1]T\n", i)
	if i == "hello" {
		fmt.Println("i = hello (string)")
	}

	i = 3.14
	fmt.Printf("i: %v, %[1]T\n", i)
	if i == 3.14 {
		fmt.Println("i = 3.14 (float64)")
	}

	if _, ok := i.(string); ok {
		fmt.Println("i string'e cast olur")
	}
	if _, ok := i.(int); ok {
		fmt.Println("i int'e cast olur")
	}
	if _, ok := i.(float64); ok {
		fmt.Println("i float64'e cast olur")
	}
}
```

Ancak `interface{}` ya da `any` olan bir tip’i `.(TYPE)` yöntemiyle başka
`TYPE`’a uygun mu değil mi diye bakabiliriz; buna **type assertion** denir.

---

## Satisfying Interface

Yani interface’i tatmin etmek, onun bize söylediği metotlara sahip olan bir
tip üretmek. Ortada bir interface varsa ve biz onu mutlu etmek istiyorsak mutlaka bize
söylediği tüm davranışlarını bizim de yapmamız (implemente etmemiz) gerekir.

Örneğin insanlar konuşabilir, robotlar da. Eğer insanların ve robotların
`Talk()` diye bir metotu olursa aralarında konuşabilirler? Bizim için
konuşacak şeyin insan ya da robot olmasının bir önemi yok, önemli olan tek şey
`Talk()` metotu olması.

Go’daki tek **abstract type**’dır (soyut tür) `interface`. Bu da şu demek,
direkt olarak kullanılamazlar ve sadece arayüz görevini üstlenirler.

Yani elektrik prizi, kendisine takılan şeyin televizyon mu? cep telefonu şarj
aletimi olduğunu bilmemesi gibi...

https://go.dev/play/p/y-nDQb82Xqi

```go
package main

import "fmt"

// Positiver defines an interface for positive things.
type Positiver interface {
	Positive() bool
}

// Numero is a custom type definition uses int.
type Numero int

// Positive is a method for satisfying Positiver interface.
func (n Numero) Positive() bool {
	return n > 0
}

// Person is a custom type definition uses string.
type Person string

// Positive is a method for satisfying Positiver interface.
func (n Person) Positive() bool {
	return true
}

// isPositive accepts an interface which satisfies Positiver interface.
func isPositive(n Positiver) bool {
	return n.Positive()
}

func main() {
	n := Numero(5)
	h := Person("vigo")

	fmt.Println(n, isPositive(n)) // 5 true
	fmt.Println(h, isPositive(h)) // vigo true
}
```

Uzun lafın kısası, `isPositive` fonksiyonu, `Positiver` interface’ini satisfy
eden herhangi bir tipi input olarak alabilir:

    +----------+                 +
    |         /                / |
    |       /  <- method     /   |  object
    |       \                \   |
    |         \                \ |
    +----------+                 +
      interface

Örneğimizi bir tık daha geliştirelim:

https://go.dev/play/p/p-J6Dcev_P8

```go
package main

import (
	"fmt"
	"strconv"
)

// Positiver defines an interface for positive things.
type Positiver interface {
	Positive() bool
}

// Numero is is a custom type definition uses int.
type Numero int

// Positive is a method for satisfying Positiver interface.
func (n Numero) Positive() bool {
	return n > 0
}

func (n Numero) String() string {
	return "my value is: " + strconv.Itoa(int(n))
}

// Person is a custom type definition uses string.
type Person string

// Positive is a method for satisfying Positiver interface.
func (n Person) Positive() bool {
	return true
}

// isPositive accepts an interface which satisfiys Positiver interface.
func isPositive(n Positiver) bool {
	return n.Positive()
}

func main() {
	n := Numero(5)
	h := Person("vigo")

	fmt.Println(n, isPositive(n)) // my value is: 5 true
	fmt.Println(h, isPositive(h)) // vigo true
}
```

Yine built-in gelen `Formatter` interface’ini mutlu edelim:

```go
type Formatter interface {
	Format(f State, verb rune)
}
```

Şimdi kendi tipimiz için özel bir gösterim ekleyelim:

https://go.dev/play/p/zjfW_cXoa-M

```go
package main

import (
	"fmt"
	"strconv"
)

// Numero is a custom type definition uses int.
type Numero int

// String implements Stringer interface.
func (n Numero) String() string {
	return strconv.Itoa(int(n))
}

// Format implements Formatter interface.
func (n Numero) Format(f fmt.State, verb rune) {
	val := n.String() // get string version
	if verb == 81 {   // check if Q passed
		val = "\"" + val + "\"" // add quotes
	}
	fmt.Fprint(f, val)
}

func main() {
	a := Numero(1)

	fmt.Printf("number is: %Q\n", a) // number is: "1"
}
```

Düşünün ki bir fonksiyon yazmak istiyorsunuz, fonksiyon bir kısım işler yapıp
çıktıyı bir dosyaya yazacak. Test yaparken dosya yerine buffer’a yazmak
istiyorsunuz. Bunun için iki farklı fonksiyon mu yazmak gerek?

- `func OutputToFile`
- `func OutputToBuffer`

Yapmamız gereken, `Writer` interface’ini sağlayan herhangi bir tip’i argüman
olarak almak;

```go
type Writer interface {
    Write([]byte) (int, error)
}
```

Yani;

```go
func OutputTo(w io.Writer, . . . ) { . . . }
```

HTTP server mı lazım? Bunun için de bir `interface` var:

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

Gereken tek şey `ServeHTTP` metotu olan bir tip, `http.ResponseWriter`ve
`http.*Request` alması yeterli. İşte go’yu güçlü kılan en büyük özellik bu!

Bu konu ile ilgili güzel bir [makale][01]

```go
type home struct {}

func (h *home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("This is my home page"))
}
```

Bir interface birden fazla interface’den türeyebilir; ayni `ReadWriter` daki
gibi:

```go
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}
```

Örneğin:

https://go.dev/play/p/0pNQqYRrdDW

```go
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// readStream accepts io.Reader.
// this is the: "accept interface as function argument" techinque!
func readStream(r io.Reader) (string, error) {
	b := make([]byte, 1024) // 1024 bytes of storage

	n, err := r.Read(b)
	if err != nil {
		return "", fmt.Errorf("read stream error: %w", err)
	}
	// report read bytes
	return fmt.Sprintf("read %d bytes: %s (%v)", n, string(b), b[:n]), nil
}

func main() {
	s, err := readStream(strings.NewReader("abcde"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("byte stream read ->", s)
	fmt.Println()

	// read from file
	// run this in bash before running the code!
	// $ echo "hello" > /tmp/foo

	f, err := os.Open("/tmp/foo")
	if err != nil {
		log.Fatal(err)
	}

	s, err = readStream(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("file read ->", s)
}
```

Hep aklımızda Rob Pike’ın şu sözü çınlamalı:

> The bigger the interface, the weaker the abstraction.

Yani interface ne kadar büyük olursa o kadar çok metotu implemente etmek
zorunda kalacağız. Mümkün mertebe metot sayısını küçük tutalım. Ne yazık ki
bunu ne kadar bilim düşünsek de, bir rest-api geliştirdiğinizde interface’ler
ne yazık ki çok büyüyor...

Mesela şöyle bir `interface` var:

```go
type MegaCharger interface {
    Foo1(int) int
    Foo2(string) error
    Foo3(bool) bool
    Foo4(float64) (int64, error)
    Foo5([]string) []int
    Foo6()
}
```

Tüm metotları bizim de ilgili tipimizde yazmamız gerekecek! Şöyle de bir
espiri var:

> Strongest abstraction is the empty interface!

Yani en süper soyutlama hiç metotu olmayan boş interface’lerdir!

Acaba metotlarını yazdığınız, implemente ettiğiniz interface’i gerçekten tam
olarak implemente edebildiniz mi? Bunun için **Compile Time Proof** taktiğini
kullanırız:

https://go.dev/play/p/LKFzNddmJ0z

```go
package main

import (
	"fmt"
	"io"
)

// DemoRW is a fake ReadWriter
type DemoRW struct{}

func (d DemoRW) Read(p []byte) (n int, err error) {
	return 1, nil
}

func (d DemoRW) Write(p []byte) (n int, err error) {
	return 1, nil
}

var (
	// hepsi aynı işi yapar
	_ io.ReadWriter = (*DemoRW)(nil) // compile time proof, doesn't allocate
	_ io.ReadWriter = &DemoRW{}      // compile time proof, doesn't allocate
	_ io.ReadWriter = new(DemoRW)    // compile time proof, doesn't allocate
)

func checkInterfaceIsReadWriter(v any) bool {
	_, ok := v.(io.ReadWriter)
	return ok
}

func main() {
	drw := &DemoRW{}

	fmt.Println(checkInterfaceIsReadWriter(drw)) // true
	fmt.Println(drw.Read([]byte("hello")))       // <nil>
}
```

Unutmayalım!

- Interface’ler kesişme noktalarıdır
- Interface’ler bağımlılıkları kırabilir bozabilirler, iyi belgelendirilmiş olmalıdırlar

Bazı faydalı linkler:

- https://www.youtube.com/watch?v=PfQFjOwGGks
- https://www.youtube.com/watch?v=ak97oH0D6fI
- http://golang.org/s/using-guru


[01]: https://lets-go.alexedwards.net/sample/02.09-the-http-handler-interface.html