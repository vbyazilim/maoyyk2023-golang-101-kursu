# Bölüm 04/05: Veri Tipleri

## Struct Annotations

`struct` alanlarına takılan meta bilgi alanlarına `tag` deniyor. Bu bilginin
eklenmesi işine de **annotation** yani not alma işlemi deniyor. Bu ek bilgiyi
struct’larla çalışırken kullanıyoruz. Struct field’ları ile ilgili ek işlemler
yapacağımız zaman bu tag’leri kullanıyoruz.

Tag; backtick **\`** karakterleri arasında tanımlanıyor ve 
`key1:"value1" key2:"value2"` şeklinde **N** tane tag alabiliyor:

```go
type S struct {
  Field fieldtype `key1:"value1" key2:"value2"`
}
```

Nerelerde işimize yarar ?

- Encoding/Decoding işlemlerinde; `json.Marshal` / `json.Unmarshal` gibi...
- Field validation (doğrulama) işlemlerinde; **alan boş olamaz**’ı kontrol
  ederken...
- Veritabanı işlerinde; database field type definition/validation...

---

## Custom Tag

Şimdi kendimize ait bir tag yapalım. Bu tag’e sahip olan string alanları
otomatik olarak içindeki yazan değere göre **upper**/**lower** case haline
dönüşsün.

https://go.dev/play/p/J3QLYdSjXtb

[Örnek kod](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/04/05-struct-custom-tag)

```go
package main

import (
	"fmt"
	"reflect"
	"strings"
)

// User holds user model data.
type User struct {
	FirstName string `case:"upper"`
	LastName  string `case:"lower"`
	Age       int    `case:"lower"` // won't affect
}

// Set sets case tag declarations.
// "case" tag is only operational for strings!
func (u *User) Set() {
	v := reflect.Indirect(reflect.ValueOf(u))
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if !field.CanSet() {
			continue
		}

		k := field.Type().Kind()
		if k == reflect.String {
			tagCase, ok := t.Field(i).Tag.Lookup("case")
			if !ok {
				continue
			}

			switch tagCase {
			case "upper":
				if field.String() != "" {
					field.SetString(strings.ToUpper(field.String()))
				}
			case "lower":
				if field.String() != "" {
					field.SetString(strings.ToLower(field.String()))
				}
			}
		}
	}
}

func main() {
	u1 := User{
		FirstName: "Uğur",
		LastName:  "Özyılmazel",
		Age:       51,
	}

	fmt.Printf("%+v\n", u1) // {FirstName:Uğur LastName:Özyılmazel Age:49}

	u1.Set()
	fmt.Println(u1.FirstName) // UĞUR
	fmt.Println(u1.LastName)  // özyılmazel
	fmt.Println(u1.Age)       // 51
}
```

`struct`’a eklediğimiz `Set` method’u ile tanımlanan alanların field’larında
modifikasyon yapıyoruz. Önceki derslerden hatırlayacağınız gibi, method
aslında bir **Pointer Receiver** : `func (u *User) Set()`.

Önce değer analizi yapıyoruz, sonra tipi buluyoruz. Tüm field’ların içinde
dolaşıp tipi `string`’e uygun olan alanı bulup tag’ine bakıyoruz. Tag değeri
**upper** ya da **lower** ise gerekli işlemi yapıp field’ın value’sunu
değiştiriyoruz.

---

## Validation (Doğrulama)

https://github.com/isacikgoz/defaults

https://go.dev/play/p/-b7GxOVnxjX

[Örnek kod](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/04/05-struct-validate)

```go
package main

import (
	"fmt"
	"log"

	"github.com/isacikgoz/defaults"
)

// User holds basic user information.
type User struct {
	Name     string `validate:"notempty"`
	Email    string `validate:"email"`
	Homepage string `validate:"url"      default:"https://vbyazilim.com"`
}

func main() {
	var u User

	// set defaults
	if err := defaults.Set(&u); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", u) // {Name: Email: Homepage:https://vbyazilim.com}

	u.Name = "Uğur Özyılmazel"
	u.Email = "ugur@fake.com"
	// u.Email = ""
	// u.Homepage = ""
	// u.Homepage = "foooo"

	if err := defaults.Validate(&u); err != nil {
		log.Fatal(err)
	}
}
```

Daha kapsamlı doğrulamalar yapmak için genelde;

https://github.com/go-playground/validator

kütüphanesini kullanıyoruz.
