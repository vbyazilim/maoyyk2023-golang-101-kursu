# Bölüm 03/03: Dil Kuralları

## Değişkenler

İçinde değer saklayan, depolayan ve değiştirilebilir olan şeylerdir. Veriye
ulaşmak için kullanılan referanstır. Go, değişkenin tipine göre içinde yazan
değeri algılar. 2 tür tanımla şekli vardır;

1. **Long Variable Declaration** : Uzun değişken tanımlama. `var` anahtar
kelimesi ile kullanılır.
1. **Short Variable Declaration** : Kısa değişken tanımlama. `:=` ile kullanılır.

```go
var x = 5       // x’in değeri 5 ve tipi: dynamic type int
var y int = 5   // y’nin değeri 5 ve tipi: static type int
var z int       // z’nin değeri 0 ve tipi: static type int

var i, j int = 1, 2 // i ve j’nin değerleri 1 ve 2, tipi static type int
```

Aynı sabitlerde olduğu gibi, eğer tip tanımı yapmazsak go bunu kendi çözmeye
çalışır. `var z int` durumunda `z`’nin tipini (mimariye göre 32bit ya da
64bit) **integer** olarak çözer ve değişkeni **initialize** ederken hemen
zero-value’sunu (her tipin bir sıfır değeri bulunur) atar. Örnekteki durumda
`z`’nin zero-value’su `0` olur.

Peki `var s string` olsa `s`’in değeri ne olurdu? Boş string yani `""`. Çünki
string’lerin zero-value’su boş string olur. Özellikle **zero-value**
kelimesini vurguluyoruz çünkü ileriki konularda `reflection`’a girdiğimizde
`IsZero` göreceğiz ve ne olduğunu daha da iyi anlayacağız.

Baz tiplerin **zero-value** değerleri;

```go
var a int                // 0
var b float32            // 0
var c complex64          // (0+0i)
var d string             // ""
var e bool               // false
var f byte               // 0
var g []int              // [] bu içinde integerların olabileceği bir array
var h struct{}           // {}
var i map[string]string  // map[]
var j func()             // nil
```

---

### `fmt`

Örnekten önce hızlıca `fmt` paketinden gelen `Print` familyasına bakalım.
Örneklerde sıkça `fmt.Print`, `fmt.Println`, `fmt.Printf` göreceğiz. Adından
da anlaşıldığı gibi bu fonksiyonlar, standart çıktıya (stdout) bilgi göndermek
için, daha da basit bir tanımla ekrana yazı yazdırmak için kullandığımız
fonksiyonlardır.

---

#### `fmt.Print`

```go
func Print(a ...any) (n int, err error)
```

Array ve Slice konusunda **variadics** kavramını işlerken bu `...`’yı detaylı
göreceğiz ama kısaca `Print` fonksiyonu `any` tipinde `n` tane parametre alabilir;

```go
fmt.Print("hello", "world", 1, 2, []string{"foo"})
// helloworld1 2 [foo]
```

`string` dışındakilerin arasına bir boşluk karakteri (space) ekler çıktıya.
Satır sonuna otomatik olarak yeni satır `\n` (new line) karakteri **eklemez**.

Tip güvenliği (type safety) ve tip tanımlamanın bu kadar katı olduğu bir dilde
dikkat ettiyseniz farklı farklı tipleri bu fonksiyona gönderebildik.

#### `fmt.Println`

Neredeyse `Print` ile aynı fakat bu kez parametreler arasına otomatik boşluk
koyar ve satır sonuna otomatik `\n` ekler:

```go
fmt.Println("hello", "world", 1, 2, []string{"foo"})
// hello world 1 2 [foo]
```

#### `fmt.Printf`

Çok sık kullanacağımız, metin formatlama, yer düzenleme (string interpolation)
gibi işlerde bize kolaylıklar sağlar. `%<VERB>` yani `%` işareti ve fiil alır,
satır sonuna otomatik olarak yeni satır `\n` (new line) karakteri **eklemez**;

```go
fmt.Printf("merhaba %s\n", "dünya")
// merhaba dünya
```

`%s` bir fiil’dir ve **the uninterpreted bytes of the string or slice** yani
yorumlanmamış string ya da slice (liste kesiti) byte’larını temsil eder.

Biz örneklerde genelde;

- `%v` : değeri (value)
- `%+v` : struct değerlerinde alan adlarıyla görme
- `%#v` : **Go-syntax representation of the value** yani değerin go
  tarafındaki kod görüntüsü
- `%T` : **Go-syntax representation of the type of the value** yani değerin go
  tarafındaki tip görüntüsü
- `%d` : 10’luk sayı sistemindeki (decimal) sayılar

Varsayılan format her zaman `%v`. Daha fazla [detay için tıklayın][01].

---

Şimdi bir değişken tanımlayıp değerini değiştirelim;

https://go.dev/play/p/-HD4UqpJ-8E

```go
package main

import "fmt"

func main() {
	var a int
	fmt.Printf("%d\n", a) // 0
	
	a = 5
	fmt.Printf("%d\n", a) // 5

	a = 100
	fmt.Printf("%d\n", a) // 100
}
```

Dikkat ettiyseniz önce `var` ile değişkeni tanımladık. Bu esnada go hafızada
bu değişken için bir alan rezerve etti. Ne kadarlık bir alan? `int`’in
ihtiyacı olduğu kadar:

https://go.dev/play/p/wDuT-QE31Vv

```go
package main

import (
	"fmt"
	"unsafe"  // bu kütüphane sayesinde `Sizeof` kullanıyoruz, sonuç byte cinsinden
)

func main() {
	var a int       // a’yı tanımladık
	var b float32
	var c string

	fmt.Printf("%v bytes\n", unsafe.Sizeof(a)) // 8 bytes
	fmt.Printf("%v bytes\n", unsafe.Sizeof(b)) // 4 bytes
	fmt.Printf("%v bytes\n", unsafe.Sizeof(c)) // 16 bytes

	a = 1000000000000000000 // tanımladığımız a’yı kullandık, değerini değiştirdik.
	fmt.Printf("%v bytes\n", unsafe.Sizeof(a)) // 8 bytes
}
```

Kısa değişken tanımlamanın `:=` olduğunu söylemiştik;

https://go.dev/play/p/eBGbxe1ZYNl

```go
package main

import "fmt"

func main() {
	a := 5

	fmt.Printf("%T\n", a)  // int
	fmt.Printf("%v\n", a)	// 5
}
```

Go bizim yerimize `a`’nın `int` olacağını anladı (type inference) ve gerekli
işlemi yaptı. Peki `a`’nın değerini değiştirmek istiyoruz;

```go
package main

import "fmt"

func main() {
	a := 5

	fmt.Printf("%T\n", a)
	fmt.Printf("%v\n", a)

	a := 8 // hata!! no new variables on left side of :=
    // Tekli atamalarda 2 kere tekrar edilemiyor
    
	fmt.Printf("%v\n", a)
}
```

Eğer kısa şekilde değişkeni tanımlamışsak artık değeri değiştirmek
istediğimizde `:=` yerine `=` kullanmamız gerekiyor. Çünkü artık `a` tipi
belli olan bir değişken:

https://go.dev/play/p/e5iRQdQTDCC

```go
package main

import "fmt"

func main() {
	a := 5

	fmt.Printf("%T\n", a) // 
	fmt.Printf("%v\n", a) // 

	a = 8
	fmt.Printf("%v\n", a)
}
```

Kısa değişken tanımlamanın bazı kısıtları var;

- Sadece **fonksiyon** içinde çalışıyor
- Tekli atamalarda **2 kere** tekrar edilemiyor
- Çoklu atamalarda tekrar oluyor ama her seferinde değeri değişiyor
- Kapsama göre tekrar olabiliyor

```go
package main

import "fmt"

a := 5 // error
       // non-declaration statement outside function body

func main() {
	fmt.Printf("a: %v\n", a)
}
```

Çoklu atamalarda tekrar oluyor ama her seferinde değeri değişiyor:

https://go.dev/play/p/gjnj3Dp-UcI

```go
package main

import "fmt"

func main() {
	number1, number2 := example1()
	fmt.Printf("number1: %d , number2: %d\n", number1, number2) // number1: 1 , number2: 2

	number1, number3 := example2()
	fmt.Printf("number1: %d , number3: %d\n", number1, number3) // number1: 100 , number3: 200

	number4, number1 := example1()
	fmt.Printf("number4: %d , number1: %d\n", number4, number1) // number4: 1 , number1: 2
}

func example1() (int, int) {
	return 1, 2
}

func example2() (int, int) {
	return 100, 200
}
```

Kapsama göre kısıtlar (scope):

https://go.dev/play/p/cLGZlfoSyqU

```go
package main

import "fmt"

var number int = 999

func main() {
	fmt.Printf("main - number: %d\n", number) // main - number: 999
	example1()

	fmt.Printf("main - example1 sonrası: %d\n", number) // main - example1 sonrası: 999
}

func example1() {
	number := 1                                   // inner-scope
	fmt.Printf("example1 - number: %d\n", number) // example1 - number: 1

	number = 666                                  // inner-scope
	fmt.Printf("example1 - number: %d\n", number) // example1 - number: 666
}
```

İleriki konularda `if`, `for`, `switch` ifadelerinde bu kısa tanımlamanın özel
kullanımlarını da göreceğiz ama hızlı bir örnek;

https://go.dev/play/p/1aw4ofBRm5d

```go
package main

import "fmt"

func main() {
	number := 100
	fmt.Printf("number: %d\n", number)     // number: 100
	if number := example1(); number == 1 { // iç kapsamda number'ı 1 yaptık
		// inner-scope
		fmt.Printf("(if) number: %d\n", number) // (if) number: 1
	}
	fmt.Printf("number halen: %d\n", number) // dış kapsamda: number halen: 100
}

func example1() int {
	return 1
}
```

Diğer çoklu tanımlama/atama şekilleri de yapmak mümkün;

https://go.dev/play/p/2wHBRk4RbZr

```go
package main

import "fmt"

var a, b int // a ve b int tipinde
var (
	x       = 1  // x dinamik tip 1
	y       = 2  // y dinamik tip 2
	abc int = 99 // abc statik tip, int 99
)

func main() {
	a = 5
	b = 10

	fmt.Printf("a: %v\n", a) // a: 5
	fmt.Printf("b: %v\n", b) // b: 10

	fmt.Printf("x: %v\n", x) // x: 1
	fmt.Printf("y: %v\n", y) // y: 2

	fmt.Printf("abc: %v\n", abc) // abc: 99

	num1, num2 := 101, 201         // num1’e 101, num2’ye 201
	fmt.Printf("num1: %v\n", num1) // num1: 101
	fmt.Printf("num2: %v\n", num2) // num2: 201
}
```

Değişken isimlendirmesinde dikkat edeceğimiz kurallar;

1. Mutlaka harf ile başlamalı
1. İçinde harf, sayı ve `_` (*underscore*) olabilir ama olmasa iyi olur
1. `camelCase`, `BumpyCaps`, `mixedCase` şeklinde tanımlama yapılabilir
1. Anlaşılır olmalıdır

Örneğin veritabanından gelen kayıtların sayısı için bir değişken tanımlamak
gerekese; "NUMBER OF RECORDS" ya da "LENGTH OF RECORDS" ya da "RECORDS LENGTH"
kafamızda olsa;

```go
var lengthOfRecords int // ya da
var recordsLength       // ya da
var numRecs             // çok tercih edilmememli
var recordsAmount       //
```

gibi varyasyonlar olabilir. Eğer imkan varsa tek bir kelime ile ifade etmek en
iyi yöntemdir. Tüm bu kurallar tüm **identifier**’lar için geçerlidir. 

Nerede bir değişken kullanımı görürseniz mutlaka o değişkenin değerini yani
**Value of**’unu kullandığınızı **unutmayın**!

---

## Kapsama Durumu

```go
package main

import "fmt"

func main() {
	n, err := fmt.Println("merhaba") // err deklara edildi ama kullanılmadı!
	if _, err := fmt.Println(n); err != nil { // bu err ile yukarıda farklı, burada inner-scope durumu var
		fmt.Println(err)
	}
}
```

`if` bloğu içindeki `err` kapsam (scope) olarak işlendi bitti. Baştaki `err`
ise deklare edildi ama kullanılmadı... Ancak aşağıdaki gibi olsa
derlenebilirdi:

```go
package main

import "fmt"

func main() {
	n, err := fmt.Println("merhaba") // merhaba
	if _, err := fmt.Println(n); err != nil { // 8 - inner-scope
		fmt.Println(err)
	} // '_, err' err artık tükendi bitti

	fmt.Println(err) // <nil> bu ise ilk satırdaki 'err'
}
```

---

## Değişkenleri Gölgeleme

**Variable Shadowing** yani bir değişkenin başka bir değişkeni gölgelemesidir:

https://go.dev/play/p/typjDuCQPN8

```go
package main

import (
	"fmt"
	"os"
)

func otherFunc(n int, buf []byte) {
	fmt.Println(n, buf)
}

// BadRead variable-shadowing örneği için bir fonksiyon.
func BadRead(f *os.File, buf []byte) error {
	var err error // zero-value nil

	for {
		n, err := f.Read(buf) // bu 'err' yukarıdaki 'err' değişkenini gölgeledi
		if err != nil {
			break
			// f aslında hatalı olduğu için 'err' nil olmayacak ve buraya girecek
			// for'dan çıkacak, otherFunc çağırılmadan loop'dan çıkılacak
			// bu 'err' -> 'var err error' kısmını gölgeleyecek
		}
		otherFunc(n, buf)
	}
	return err // her zaman nil dönecek çünkü ilk tanımlandığı gibi zero-value durumunda kaldı
}

func main() {
	f, _ := os.Open("/tmp/fake") // error'ü yutuyoruz
	var b []byte                 // BadRead için buffer

	if err := BadRead(f, b); err != nil { // bu kısım hata dönmeli çünkü /tmp/fake dosyası yok!
		panic(err)
	}

	fmt.Println("merhaba") // merhaba
}
```

[01]: https://pkg.go.dev/fmt