# Bölüm 04/05: Veri Tipleri

## Structs

Array ve Slice gibi Struct’da **composite types** ailesinden bir tiptir.

**Structure** yani yapı kelimesinin kısaltılmış halidir `struct`. Yapısal veri
saklamanın en kısa ve basit yöntemidir. İçinde alanlardan oluşan koleksiyonlar
tutar. Bence go’nun en önemli iki konusundan biridir.

Konuyu anlamak için `struct` tipini, veritabanındaki tablo gibi
düşünebilirsiniz. Örneğin kullanıcıları sakladığımız bir tablo olsa.
Kullanıcının adı, soyadı, e-posta adresi, şifresi ve yaşı olsa;

```go
type user struct {
	firstName string
	lastName  string
	email     string
	password  string
	age       int
}
```

Bu durumda `user`’ın alanları;

- `firstName`: Tipi `string`
- `lastName`: Tipi `string`
- `email`: Tipi `string`
- `password`: Tipi `string`
- `age`: Tipi `int`

Go’da tip tanımı yapmadığımız neredeyse hiç bir yer yok. Alanların adı olduğu
gibi tipi de olmak zorunda. Eğer tip tanımı varsa, tiplerin
**zero-value**’ları da var. Yani `firstName` alanının başlangıç (initial)
değeri boş string yani `""`.

Alan adlarını gruplamak da mümkün;

```go
type user struct {
	firstName, lastName, email, password string
	age                                  int
}
```

Gerçek dünyada genelde her şeyi açık açık yazmak ve görmek istiyoruz, bu
bakımdan gruplama stilini pek de kullanmıyoruz.

`type user struct` go açısından **Named Structure** yani ismi olan bir yapı.
Aynı mantıkla ismi olmayan yapılar yani **Anonymous Structure** da mümkün.
Şimdi her ikisini de kullanan örneğe bakalım:

https://go.dev/play/p/mQbUP-GG40Q

```go
package main

import "fmt"

var user struct {
	firstName string
	lastName  string
	email     string
	password  string
	age       int
}

func main() {
	user1 := user
	user1.firstName = "Uğur"
	user1.lastName = "Özyılmazel"
	user1.email = "vigo@xxx.com"
	user1.password = "1234"
	user1.age = 51

	user2 := user
	user2.firstName = "Erhan"
	user2.lastName = "Akpınar"
	user2.email = "erhan@xxx.com"
	user2.password = "1234"
	user2.age = 38

	// anonymous struct
	user3 := struct {
		firstName string
		lastName  string
		email     string
		password  string
		age       int
	}{
		firstName: "Ezel",
		lastName:  "Özyılmazel",
		email:     "ezel@yyy.com",
		password:  "1234",
		age:       12,
	}

	// anonymous struct
	user4 := struct {
		firstName string
		lastName  string
		email     string
		password  string
		age       int
	}{
		"Ali", // kod okunaklığı açısından iyi değil
		"Desidero",
		"alide@me.com",
		"1234",
		77,
	}

	fmt.Printf("user1.firstName: %s\n", user1.firstName) // Uğur
	fmt.Printf("user2.firstName: %s\n", user2.firstName) // Erhan
	fmt.Printf("user3.firstName: %s\n", user3.firstName) // Ezel
	fmt.Printf("user4.firstName: %s\n", user4.firstName) // Ali
}
```

Tekrar zero-value olayını hatırlayalım:

https://go.dev/play/p/o0NIIV0ss1P

```go
package main

import "fmt"

type user struct {
	firstName string // string’lerin zero-value’su yani ""
	lastName  string // string’lerin zero-value’su yani ""
	email     string // string’lerin zero-value’su yani ""
	password  string // string’lerin zero-value’su yani ""
	age       int    // int’lerin zero-value’su yani 0
}

func main() {

	user1 := user{} // boş yapı

	fmt.Printf("%v\n", user1)  // {    0}
	fmt.Printf("%+v\n", user1) // {firstName: lastName: email: password: age:0}
}
```

Atama esnasında bazı alanlara değer atayıp bazı alanları pas geçebiliriz, bu durumda
pas geçilenler yine zero-value’larını alır:

https://go.dev/play/p/qNkYB9_A5cr

```go
package main

import "fmt"

type user struct {
	firstName string
	lastName  string
	email     string
	password  string
	age       int
}

func main() {

	user1 := user{
		firstName: "Uğur",
		lastName:  "Özyılmazel",
	}
	user2 := user{firstName: "Ezel"}
	user3 := user{age: 11}

	fmt.Printf("%+v\n", user1) // {firstName:Uğur lastName:Özyılmazel email: password: age:0}
	fmt.Printf("%+v\n", user2) // {firstName:Ezel lastName: email: password: age:0}
	fmt.Printf("%+v\n", user3) // {firstName: lastName: email: password: age:11}
}
```

Boş bir struct tanımı yapıp içini sonradan da doldurabiliriz:

https://go.dev/play/p/KOP80WwJaYZ

```go
package main

import "fmt"

type user struct {
	firstName string
	lastName  string
	email     string
	password  string
	age       int
}

func main() {

	var user1 user // user1, user tipinde bir değişken

	user1.firstName = "Uğur"
	user1.lastName = "Özyılmazel"

	fmt.Printf("%+v\n", user1) // {firstName:Uğur lastName:Özyılmazel email: password: age:0}
}
```

Keza struct’ı `new` anahtar kelimesiyle **initialize** edip, hafızada yer
rezervasyonu yapabiliriz. `new` ile tanımladığımız zaman bize **pointer**
döner (hafıza adresi) ve initialize olduğu için hafızada yer kaplamış
(allocation yapmış) oluruz ve zero-value’u atamış oluruz:

https://go.dev/play/p/sU6WL_9oK2Y

```go
package main

import "fmt"

type user struct {
	firstName string
	lastName  string
	email     string
	password  string
	age       int
}

func main() {
	user1 := new(user) // hafızayı user tipi için gereken yer kadar rezerve et
	user2 := user{}

	fmt.Printf("user1: %T\n", user1) // user1: *main.user (pointer geldi)
	fmt.Printf("user2: %T\n", user2) // user2: main.user

	fmt.Printf("%v\n", user1) // &{    0}
	fmt.Printf("%v\n", user2) // {    0}

	user1.firstName = "Uğur"
	user2.firstName = "Ezel"

	fmt.Printf("%s\n", user1.firstName) // Uğur
	fmt.Printf("%s\n", user2.firstName) // Ezel

	fmt.Printf("%v\n", *user1) // {Uğur    0} * ile "value of", dereferencing
}
```

Dikkat ettiyseniz `fmt.Printf("user1: %T\n", user1)` ile `user1`’in tipini
yazdırdığımızda bize `*main.user` geldi. Hafızada ayrılan adresi işaret eden,
yani **pointer**’ı döndü. **Pointer** konusunu ileride işleyeceğiz ama hızlıca
geçmek gerekirse;

| Sembol | Açıklaması |
|:------:|:------------------------------------------------------------|
| `*`    | **value of** : yani değeri, `C`’deki **dereference** işlemi |
| `&`    | **address of**: yani hafızadaki hexadecimal adresi          | 

Go bize **explicit dereference** yani `(*user1).firstName` yaparak erişmek
yerine direkt olarak `user1.firstName` şeklinde erişmeye imkan sağlar;

`new` sadece `struct` için değil tüm **concrete type**’lar için geçerlidir:

https://go.dev/play/p/0K26qhK2VjD

```go
package main

import "fmt"

type myInt int

type intPool []myInt

type runner interface {
	Run() error
}

func main() {
	a := new(int)
	b := new(string)
	c := new(bool)
	d := new(float64)
	e := new(myInt)
	f := new(intPool)
	g := new(map[string]string)
	h := new([]byte)

	i := new(runner)

	fmt.Printf("a, type: %T value: %[1]v *value: %v\n", a, *a)
	// a, type: *int value: 0x1400001a100 *value: 0

	fmt.Printf("b, type: %T value: %[1]v *value: %v\n", b, *b)
	// b, type: *string value: 0x14000010250 *value:

	fmt.Printf("c, type: %T value: %[1]v *value: %v\n", c, *c)
	// c, type: *bool value: 0x1400001a108 *value: false

	fmt.Printf("d, type: %T value: %[1]v *value: %v\n", d, *d)
	// d, type: *float64 value: 0x1400001a110 *value: 0

	fmt.Printf("e, type: %T value: %[1]v *value: %v\n", e, *e)
	// e, type: *main.myInt value: 0x1400001a118 *value: 0

	fmt.Printf("f, type: %T value: %[1]v *value: %v\n", f, *f)
	// f, type: *main.intPool value: &[] *value: []

	fmt.Printf("g, type: %T value: %[1]v *value: %v\n", g, *g)
	// g, type: *map[string]string value: &map[] *value: map[]

	fmt.Printf("h ([]byte), type: %T value: %[1]v *value: %v\n", h, *h)
	// h ([]byte), type: *[]uint8 value: &[] *value: []

	fmt.Printf("i (interface), type: %T value: %[1]v *value: %v\n", i, *i)
	// i (interface), type: *main.runner value: 0x14000010260 *value: <nil>
}
```

**Explicit dereference** (açık) örneği;

https://go.dev/play/p/DlUMpQN2b40

```go
package main

import "fmt"

type user struct {
	firstName string
	lastName  string
	email     string
	password  string
	age       int
}

func main() {

	user1 := new(user)
	user1.firstName = "Uğur"

	fmt.Printf("%s\n", (*user1).firstName) // Uğur
	fmt.Printf("%s\n", user1.firstName)    // Uğur

	fmt.Println(user1.firstName == (*user1).firstName) // true
}
```

Hem `new(user)` hem de `&user{}` aynı işi yaparlar, hafızada "zero user"
allocation yaparlar ve rezerve edilen hafızanın adresini (pointer’ını)
dönerler. 

`new` tüm tipler için kullanılabilir; `new(int)` gibi ama `&TYPE` sadece
struct için geçerlidir.

Struct içinde anonim alanlar yapmak da mümkün;

```go
package main

import "fmt"

type user struct {
	string
	int
}

func main() {
	user1 := user{"Uğur Özyılmazel", 46}
	fmt.Printf("%+v\n", user1) // {string:Uğur Özyılmazel int:46}
}
```

Peki bu anonim yapının alanlarına (field’larına) nasıl erişeceğiz ? Doğal
olarak alan adları belirtilen tip adı oluyor:

https://go.dev/play/p/vjpd0v0UY9o

```go
package main

import "fmt"

type user struct {
	string
	int
}

func main() {
	var user1 user

	user1.string = "Uğur Özyılmazel"
	user1.int = 46

	fmt.Printf("%+v\n", user1) // {string:Uğur Özyılmazel int:46}

	fmt.Printf("%s\n", user1.string) // Uğur Özyılmazel
	fmt.Printf("%d\n", user1.int)    // 46
}
```

Bu örnek sadece **proof-of-concept** yani çalıştığını göstermek için, gündelik
hayatta hiç de iyi bir pratik değil. Unutmayın ki iki tane aynı anonim alan olamaz:

```go
type user struct {
	string
	string
	int
	int
}

// derlemez! duplicate field!
```

İç-içe geçmiş, yani **Nested Structures** yapmak da mümkün:

https://go.dev/play/p/Mhlg79fGGbH

```go
package main

import "fmt"

type person struct {
	name    string
	age     int
	address address
}

type address struct {
	city, country string
}

func main() {
	p1 := person{}
	p1.name = "Uğur Özyılmazel"
	p1.age = 46
	p1.address = address{
		city:    "İstanbul",
		country: "Türkiye",
	}

	fmt.Printf("%+v\n", p1) // {name:Uğur Özyılmazel age:46 address:{city:İstanbul country:Türkiye}}

	fmt.Printf("city: %s\n", p1.address.city)       // city: İstanbul
	fmt.Printf("country: %s\n", p1.address.country) // country: Türkiye
}
```

İç-içe struct’ların güzel bir özelliği de **Promoted Fields** yani
`p1.address.city` yerine, `p1.city` şeklinde erişmek mümkün, sadece küçük bir
değişiklik yaparak;

https://go.dev/play/p/OxoWXMYMzgQ

```go
package main

import "fmt"

type person struct {
	name    string
	age     int
	address // address address eski haliydi
}

type address struct {
	city, country string
}

func main() {
	p1 := person{}
	p1.name = "Uğur Özyılmazel"
	p1.age = 46
	p1.address = address{
		city:    "İstanbul",
		country: "Türkiye",
	}

	fmt.Printf("%+v\n", p1) // {name:Uğur Özyılmazel age:46 address:{city:İstanbul country:Türkiye}}

	fmt.Printf("city: %s\n", p1.city)       // city: İstanbul
	fmt.Printf("country: %s\n", p1.country) // country: Türkiye
}
```

Anonim struct’a ait olan alanlar **Promoted Fields** oluyor! Eğer promoted
field adı, içine gömüldüğü struct’ın içindeki bir field ile çakışırsa,
promotion suya düşer :)

https://go.dev/play/p/ch8zh16UpcP

```go
package main

import "fmt"

type person struct {
	name string
	age  int
	city string
	address
}

type address struct {
	city, country string
}

func main() {
	p1 := person{}
	p1.name = "Uğur Özyılmazel"
	p1.age = 46
	p1.city = "New York"
	p1.address = address{
		city:    "İstanbul",
		country: "Türkiye",
	}

	fmt.Printf("%+v\n", p1) // {name:Uğur Özyılmazel age:46 address:{city:İstanbul country:Türkiye}}

	fmt.Printf("city (promoted): %s\n", p1.city) // city: New York
	fmt.Printf("city: %s\n", p1.address.city)    // city: İstanbul
	fmt.Printf("country: %s\n", p1.country)      // country: Türkiye
}
```

Promote edilen field’lara kolay erişim olmasına rağmen, yeni bir kopya (instance)
çıkarılacağı zaman, açık açık gömülü struct ve alanlarını yazmak gerekir. Yani
`p1.city` ile ulaşırız ama `p1.city = ...` şeklinde bir ifade yazamayız.

Peki bir şekilde bu alanların bazılarını erişime açmak kapamak gerekse? 

Nesne yönelimli dillerin sınıf konusunda bahsi çokça geçen **public/private
access control** yani sınıfın dışından ya da içinde erişilenler... Unutmayalım
ki go’da sınıf kavramı yok, **composition** yani birleşme/kompozisyon
mantığı var.

Dersin başında `fmt.Println` fonksiyonundan bahsederken **Exportable**
kavramına hafifçe dokunmuştuk. Go, değişken/sabit/fonksiyon/alan gibi her ne
tanımlıyorsanız, eğer **Büyük** harfle başlamışsa bu dışarıdan erişilebilir
anlamına geliyordu.

Örneğin `import "fmt"` diyoruz ve `fmt.Println("Hello")` dediğimizde, adı
`fmt` olan bir paketi yani ilk satırında `package fmt` yazan paketi içeri
alıyoruz ve `Println`’ın `P`’si büyük olduğu için bu fonksiyonu
çağırabiliyoruz.

Eğer biz `person` diye bir paket yapıyor olsaydık;

```go
package person

// Person represents the Person model
type Person struct {
	FirstName string // Exportable
	LastName  string // Exportable
	secret    string // Unexportable (private)
}
```

ve başka bir paketten `person` paketini `import` edip kullansak, [örnek kod](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/04/05-struct-field-access);

```go
package main

import (
	"fmt"

	"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/04/05-struct-field-access/person"
)

func main() {
	p := person.Person{} // boş bir kopya (instance)

	p.FirstName = "Uğur"
	p.LastName = "Özyılmazel"

	fmt.Printf("p: %#v\n", p) // p: person.Person{FirstName:"Uğur", LastName:"Özyılmazel", secret:""}

	fmt.Println(p.secret) // p.secret undefined (type person.Person has no field or method secret)
}
```

`p.secret` dış dünyadan erişime kapalı. `secret` sadece içeriden erişilen bir
şey. Bu bakımdan `person` paketi içinde hem bu `secret` field’ına atama yapan
hem de `secret`’a erişmeyi sağlacak bir **Getter** ve **Setter** metotlarına
ihtiyacımız olacak; [örnek kod](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/04/05-struct-field-access-getter);

`person.go`

```go
package person

// Person represents the Person model.
type Person struct {
	FirstName string // Exportable
	LastName  string // Exportable
	secret    string // Unexportable (private)
}

// Secret returns private secret field.
func (u Person) Secret() string {
	return u.secret
}

// SetSecret sets private secret value.
func (u *Person) SetSecret(s string) {
	u.secret = s
}
```


`main.go`

```go
package main

import (
	"fmt"

	"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/04/05-struct-field-access-getter/person"
)

func main() {
	p := person.Person{} // boş bir kopya (instance)

	p.FirstName = "Uğur"
	p.LastName = "Özyılmazel"

	fmt.Printf("%+v\n", p) // {FirstName:Uğur LastName:Özyılmazel secret:}

	p.SetSecret("<secret>")

	fmt.Printf("%+v\n", p)  // {FirstName:Uğur LastName:Özyılmazel secret:<secret>}
	fmt.Println(p.Secret()) // <secret>
}
```

Struct’lar **value type** oldukları için karşılaştırılabilirler (comparable):

https://go.dev/play/p/v8KG2KFp6Xl

```go
package main

import "fmt"

type person struct {
	name string
}

func main() {
	p1 := person{"Uğur"}
	p2 := person{"Uğur"}

	fmt.Printf("%v\n", p1)       // Uğur
	fmt.Printf("%v\n", p2)       // Uğur
	fmt.Printf("%v\n", p1 == p2) // true
}
```

Bu karşılaştırma için alanların tipine de bağlıdır, eğer alan tipleri
**comparable** değilse karşılaştırma yapılamaz:

https://go.dev/play/p/vNOLwgPPhnJ

```go
package main

import (
	"fmt"
)

type image struct {
	data map[int]int
}

func main() {
	image1 := image{data: map[int]int{0: 155}}
	image2 := image{data: map[int]int{0: 155}}

	if image1 == image2 {
		fmt.Println("image1 and image2 are equal")
	}
}
// invalid operation: image1 == image2 
// (struct containing map[int]int cannot be compared)
```

Son olarak, struct tasarlarken hafızada kaplayacağı yeri de düşünmemiz
gerekebilir. Alanların tiplerinin kapladığı yere göre küçükten büyüğe göre
sıralama yapmak iyi bir pratiktir:

https://go.dev/play/p/Ab2qYHxklau

```go
package main

import (
	"fmt"
	"unsafe"
)

type bad struct {
	field1 bool    // bool -> 1 byte, padding yüzünden 8 byte yedi
	field2 int64   // int64 -> 8 byte
	field3 bool    // bool -> 1 byte, padding yüzünden 8 byte yedi
	field4 float64 // float64 -> 8 byte

	// aslında 18 byte'lık yer kaplaması lazımken;
	// 7 + 7 = 14 byte daha geldi
	// 32 byte oldu
}

type good struct {
	field2 int64   // int64 -> 8 byte
	field4 float64 // int64 -> 8 byte
	field1 bool    // bool -> 1 byte
	field3 bool    // bool -> 1 byte

	// aslında 18 byte'lık yer kaplaması lazımken;
	// bool'ları 8'in içine sığdırdı (1+1=2), padding'i sağlamak için 6 byte ekledi
	// 24 byte oldu
}

func main() {
	fmt.Println(unsafe.Sizeof(bad{}), "bytes")  // 32 bytes
	fmt.Println(unsafe.Sizeof(good{}), "bytes") // 24 bytes
}
```

Her alan için minimum `8 byte`’lık blok (chunk) rezerve ediyor. Yetmezse bir 8
daha ekliyor (slice capacity gibi düşünün) eğer 8’den az gelirse 8’e
tamamlıyor, buna da **padding** deniyor.

Struct alanlarının ya da bir tipin hafızada kaç byte harcadığını `unsafe`
paketini kullanarak bulabilirsiniz:

https://go.dev/play/p/1LtOv0__law

```go
package main

import (
	"fmt"
	"unsafe"
)

type user struct {
	email    string
	isActive bool
}

func main() {
	var a []int
	var b string

	u := user{} // yeni bir user instance

	fmt.Println(unsafe.Sizeof(a))          // 24 byte
	fmt.Println(unsafe.Sizeof(b))          // 16 byte
	fmt.Println(unsafe.Sizeof(u.isActive)) // 1 byte
}
```

Go bu işi kolay çözmek için bir tool yayınladı: `fieldalignment`

```go
go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@latest
```

Kurulumu yaptıktan sonra, eğitim projesi altından;

```bash
$ cd /path/to/maoyyk2023-golang-101-kursu/
$ fieldalignment src/04/05-struct-field-alignment/main.go
main.go:6:10: struct of size 56 could be 48
```

Dosyanın üzerine yazarak otomatik düzeltme yapmak için;

```bash
$ fieldalignment -fix src/04/05-struct-field-alignment/main.go  # main.go dosyasını değiştirir
```

Şu struct:

```go
type Bad struct {
	Field1 bool    // 1 (+7) = 8
	Field2 int64   // 8
	Field3 bool    // 1 (+7) = 8
	Field4 float64 // 8
	Field5 []bool  // 24
	// 8 + 8 + 8 + 24 = 56
}
```

Düzenleme sonrası;

```go
type Bad struct {
	Field5 []bool  // 24
	Field2 int64   // 8
	Field4 float64 // 8
	Field1 bool    // 1 + 1 = 2 (+6) = 8
	Field3 bool    // ----^
	// 24 + 8 + 8 + 8 = 48
}
```

şeklini aldı. Konu ile ilgili [şirket blogumuzda bir makale][01] de yayınlamıştık.

### Empty Struct

`0` byte yer tutan boş bir struct:

https://go.dev/play/p/2b4hMxnuXRM

```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := struct{}{}

	fmt.Println(a)                // {}
	fmt.Println(unsafe.Sizeof(a)) // 0
}
```

Nerelerde kullanırız?

- [Concurrency][02] konusunda `channel` kullanımında
- [Map][03]’de value olarak


[01]: https://vbyazilim.com/blog/2022/11/28/struct-field-alignment-in-golang/
[02]: https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/docs/15/01-concurrency.md
[03]: https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/docs/04/08-map.md
