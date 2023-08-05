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

ve başka bir paketten `person` paketini `import` edip kullansak, [örnek kod](../src/04/05-struct-field-access);
