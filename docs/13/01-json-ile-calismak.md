# Bölüm 13/01: JSON İle Çalışmak

Go ile çok yüksek ihtimal ağ servisleri geliştireceksiniz. Dış dünyadan veri
almak ya da dış dünyadan istek atan bir istemciye cevap vermek için bir kaç
farklı yöntem mevcut. Bunlardan en sık kullanacaklarınız arasında `json` yani
hafif siklet veri değişim formatını kullancaksınız. Bu bakımdan;

1. Ya, `go` daki bir tipi, dış dünyanın anlayacağı `json` hale (`Marshal`)
1. Ya da dış dünyadan `json` formatında gelen veriyi `go`’nun anlayacağı hale (`Unmarshal`)

getirmek çok sık yaptığımız bir iş olacak. Bu dönüşüm esnasında 

---

## `encoding/json` Marshal

Elimizdeki `go` tipini `[]byte` haline dönüştürme işlemidir:

https://go.dev/play/p/nOzvhUsD40p

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// UserLevel is a type alias for defining user level.
type UserLevel string

// UserLevels holds collection of UserLevel types.
type UserLevels []UserLevel

func main() {
	UserLevelAdmin := UserLevel("admin")
	UserLevelModerator := UserLevel("moderator")
	UserLevelAnonymous := UserLevel("anonymous")

	userLevels := UserLevels{
		UserLevelAdmin,
		UserLevelModerator,
		UserLevelAnonymous,
	}

	j, err := json.Marshal(userLevels)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(j)
	// [91 34 97 100 109 105 110 34 44 34 109 111 100 101 114 97 116 111 114 34 44 34 97 110 111 110 121 109 111 117 115 34 93]

	fmt.Println(string(j))
	// ["admin","moderator","anonymous"]
}
```

`j`’nin string görüntüsü: `["admin","moderator","anonymous"]` şeklindedir. Bu
aslında JavaScript’teki `array` tipidir.

---

## `encoding/json` Unmarshal

Şimdi, tam ters işlemi yapalım; dış dünyadan bize `["admin","moderator","anonymous"]`
gelsin:

https://go.dev/play/p/7l7sIFzUh1S

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// UserLevel is a type alias for defining user level.
type UserLevel string

// UserLevels holds collection of UserLevel types.
type UserLevels []UserLevel

func main() {
	// byte slice from raw string.
	input := []byte(`["admin","moderator","anonymous"]`)

	// target data type for serialization.
	var userLevels UserLevels

	if err := json.Unmarshal(input, &userLevels); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", userLevels)
	// main.UserLevels{"admin", "moderator", "anonymous"}
}
```

---

## `json:"FIELD"` Tag’i

Genelde bu çevirme işlerini `struct` tipi üzerinde yaparız. Dışarıdan gelen ya
da dışarıya gidecek veride daha yapısal (structural) veri kullanmak isteriz:

https://go.dev/play/p/BB5D_VrLpLa

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// User represents user model.
type User struct {
	Name  string
	Email string
	Age   int
}

func main() {
	u := User{
		Name:  "Uğur Özyılmazel",
		Email: "vigo@example.com",
		Age:   51,
	}

	j, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(j)
	// [123 34 78 97 109 101 34 58 34 85 196 159 117 114 32 195 150 122 121 196 177 108 109 97 122 101 108 34 44 34 69 109 97 105 108 34 58 34 118 105 103 111 64 101 120 97 109 112 108 101 46 99 111 109 34 44 34 65 103 101 34 58 53 49 125]
	
	fmt.Println(string(j))
	// {"Name":"Uğur Özyılmazel","Email":"vigo@example.com","Age":51}
}
```

Dikkat ettiyseniz, `[]byte`’a baktığımızda; `{"Name":"Uğur
Özyılmazel","Email":"vigo@example.com","Age":51}` alan adlarının `struct`
field name’leri ile aynı olduğunu görürüz. `json` convention’a baktığımızda
alan adlarının bir kuralı var. Ana kural `key`’lerin (Name, Email ...) mutlaka
**Unicode** karakterlerden oluşmasıdır. Kimileri `camelCase` kimileri
`snake_case` de kullanabilir.

Kuralına uygun yazılmış bir rest-api servisi, alan adı kuralı olarak
`snake_case` kullanmalıdır.

Peki, biz `Name` yerine `name` nasıl döneceğiz? Bu durumda `json:"FIELD"` tag
devreye girer:

https://go.dev/play/p/EY_4u3pniOf

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// User represents user model.
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	u := User{
		Name:  "Uğur Özyılmazel",
		Email: "vigo@example.com",
		Age:   51,
	}

	j, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(j)
	// [123 34 110 97 109 101 34 58 34 85 196 159 117 114 32 195 150 122 121 196 177 108 109 97 122 101 108 34 44 34 101 109 97 105 108 34 58 34 118 105 103 111 64 101 120 97 109 112 108 101 46 99 111 109 34 44 34 97 103 101 34 58 53 49 125]

	fmt.Println(string(j))
	// {"name":"Uğur Özyılmazel","email":"vigo@example.com","age":51}
}
```

`json` tag’i (aslında bu bir struct annotation) sadece bu işe yaramaz;

```go
// JSON çıktıda key "myName" olarak görünür.
Field int `json:"myName"`

// JSON çıktıda key "myName" olarak görünür fakat değer "empty" (zero-value) ise
// omit edilir (yani dışarıda bırakılır) ve bu field çıktıda yer almaz!
Field int `json:"myName,omitempty"`

// JSON çıktıda key "Field" olarak görünür fakat değer "empty" (zero-value) ise
// omit edilir (yani dışarıda bırakılır) ve bu field çıktıda yer almaz!
Field int `json:",omitempty"`

// Field komple görmezden gelinir (ignore) ve çıktıda yer almaz!
Field int `json:"-"`

// JSON çıktıda key "-" olarak görünür.
Field int `json:"-,"`
```

Ek olarak, dış dünyadan gelen veri `string` görünümünde ama değer olarak
`integer`;

```json
{"age": "51"}
```

`age`’in tip olarak değeri `51` yani `string` ama biz bunu içeride `int`
olarak kullanmak istiyoruz (bazı kötü tasarlanmış api’larda bu tür durumlar oluyor):

```json
Age int `json:"age,string"`
```

şeklinde de kullanabiliyoruz:

https://go.dev/play/p/Qp0lnPOPBRQ

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// User represents user model.
type User struct {
	Name     string `json:"name"`
	Email    string `json:"email,omitempty"`
	Age      int    `json:"age,string,omitempty"`
	Password string `json:"-"`
}

func main() {
	u := User{
		Name:  "Uğur Özyılmazel",
		Email: "vigo@example.com",
		Age:   51,
	}

	j, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(j)
	// [123 34 110 97 109 101 34 58 34 85 196 159 117 114 32 195 150 122 121 196 177 108 109 97 122 101 108 34 44 34 101 109 97 105 108 34 58 34 118 105 103 111 64 101 120 97 109 112 108 101 46 99 111 109 34 44 34 97 103 101 34 58 34 53 49 34 125]

	fmt.Println(string(j))
	// {"name":"Uğur Özyılmazel","email":"vigo@example.com","age":"51"}

	u2 := User{
		Name: "Uğur Özyılmazel",
	}

	j, err = json.Marshal(u2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(j)
	// [123 34 110 97 109 101 34 58 34 85 196 159 117 114 32 195 150 122 121 196 177 108 109 97 122 101 108 34 125]

	fmt.Println(string(j))
	// {"name":"Uğur Özyılmazel"}
}
```

Şimdi [örneğe](../../src/13/json-marshal-custom-time) bakalım, `struct`
alanları içinde gezen, tag’leri bulan ve özelleştirilmiş `time` kullanan bir
yapı bulunuyor. `go`, `Marshal` işlemi yaparken, eğer custom type varsa ve bu
type’ın `MarshalJSON()` metodu varsa onu kullanıyor. `describeStruct` ise
`reflect` paketini kullanarak struct içinde geziyor.

https://go.dev/play/p/fsbNbTOMVav

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"
)

// User holds user model data. User struct has json tags!
type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Age       int       `json:"age"`
	BirthDate time.Time `json:"birth_date"`
	Admin     bool      `json:"admin"`
	LastVisit time.Time `json:"-"` // omitted, ignore
}

const customTimeLayout = "2006-01-02T15:04:05-07:00"

// CustomTime is a type alias for time.Time, uses custom marshal format.
type CustomTime struct {
	time.Time
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + ct.Time.Format(customTimeLayout) + `"`), nil
}

// UserWithCustomTime holds user model data with custom time type!
type UserWithCustomTime struct {
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Age       int        `json:"age"`
	BirthDate time.Time  `json:"birth_date"`
	Admin     bool       `json:"admin"`
	LastVisit CustomTime `json:",omitempty"`
}

// OtherUser holds user model data, only one field has tag!
type OtherUser struct {
	FirstName string
	LastName  string
	Age       int
	BirthDate time.Time
	Admin     bool
	LastVisit *time.Time `json:",omitempty"` // this field uses pointer, can be nil, can be omitted if unset!
}

func describeStruct(v any) {
	t := reflect.TypeOf(v)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf(
			"Name: %-12s Type: %-18s Tag: %-20s json: %s\n",
			field.Name,
			field.Type,
			field.Tag,
			field.Tag.Get("json"),
		)
	}
	fmt.Println()
}

func marshalStruct(v any) error {
	j, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return err
	}
	fmt.Printf("%s\n\n", j)
	return nil
}

func main() {
	trTZ := time.FixedZone("UTC+3", +3*60*60)
	now := time.Now()

	u1 := User{
		FirstName: "Uğur",
		LastName:  "Özy",
		Age:       49,
		BirthDate: time.Date(1972, time.August, 13, 10, 0, 0, 0, trTZ),
		Admin:     true,
		LastVisit: now,
	}
	describeStruct(u1)
	// Name: FirstName    Type: string             Tag: json:"first_name"    json: first_name
	// Name: LastName     Type: string             Tag: json:"last_name"     json: last_name
	// Name: Age          Type: int                Tag: json:"age"           json: age
	// Name: BirthDate    Type: time.Time          Tag: json:"birth_date"    json: birth_date
	// Name: Admin        Type: bool               Tag: json:"admin"         json: admin
	// Name: LastVisit    Type: time.Time          Tag: json:"-"             json: -

	u2 := OtherUser{
		FirstName: "Ezel",
		LastName:  "Özy",
		Age:       10,
		BirthDate: time.Date(2011, time.August, 13, 7, 0, 0, 0, trTZ),
		// we don't have LastVisit field!
	}
	describeStruct(u2)
	// Name: FirstName    Type: string             Tag:                      json:
	// Name: LastName     Type: string             Tag:                      json:
	// Name: Age          Type: int                Tag:                      json:
	// Name: BirthDate    Type: time.Time          Tag:                      json:
	// Name: Admin        Type: bool               Tag:                      json:
	// Name: LastVisit    Type: *time.Time         Tag: json:",omitempty"    json: ,omitempty

	u3 := OtherUser{
		FirstName: "Ezel",
		LastName:  "Özy",
		Age:       12,
		BirthDate: time.Date(2011, time.August, 13, 7, 0, 0, 0, trTZ),
		LastVisit: &now,
	}
	describeStruct(u3)
	// Name: FirstName    Type: string             Tag:                      json:
	// Name: LastName     Type: string             Tag:                      json:
	// Name: Age          Type: int                Tag:                      json:
	// Name: BirthDate    Type: time.Time          Tag:                      json:
	// Name: Admin        Type: bool               Tag:                      json:
	// Name: LastVisit    Type: *time.Time         Tag: json:",omitempty"    json: ,omitempty

	u4 := UserWithCustomTime{
		FirstName: "Ezel",
		LastName:  "Özyılmazel",
		Age:       12,
		BirthDate: time.Date(2011, time.August, 13, 7, 0, 0, 0, trTZ),
		LastVisit: CustomTime{now},
	}
	describeStruct(u4)
	// Name: FirstName    Type: string             Tag: json:"first_name"    json: first_name
	// Name: LastName     Type: string             Tag: json:"last_name"     json: last_name
	// Name: Age          Type: int                Tag: json:"age"           json: age
	// Name: BirthDate    Type: time.Time          Tag: json:"birth_date"    json: birth_date
	// Name: Admin        Type: bool               Tag: json:"admin"         json: admin
	// Name: LastVisit    Type: main.CustomTime    Tag: json:",omitempty"    json: ,omitempty

	if err := marshalStruct(u1); err != nil {
		log.Fatal(err)
	}
	// {
	//     "first_name": "Uğur",
	//     "last_name": "Özy",
	//     "age": 49,
	//     "birth_date": "1972-08-13T10:00:00+03:00",
	//     "admin": true
	// }

	if err := marshalStruct(u2); err != nil {
		log.Fatal(err)
	}
	// {
	//     "FirstName": "Ezel",
	//     "LastName": "Özy",
	//     "Age": 10,
	//     "BirthDate": "2011-08-13T07:00:00+03:00",
	//     "Admin": false
	// }

	if err := marshalStruct(u3); err != nil {
		log.Fatal(err)
	}
	// {
	//     "FirstName": "Ezel",
	//     "LastName": "Özy",
	//     "Age": 12,
	//     "BirthDate": "2011-08-13T07:00:00+03:00",
	//     "Admin": false,
	//     "LastVisit": "2023-08-13T14:07:09.537756+03:00"
	// }

	if err := marshalStruct(u4); err != nil {
		log.Fatal(err)
	}
	// {
	//     "first_name": "Ezel",
	//     "last_name": "Özyılmazel",
	//     "age": 12,
	//     "birth_date": "2011-08-13T07:00:00+03:00",
	//     "admin": false,
	//     "LastVisit": "2023-08-13T14:07:09+03:00"
	// }
}
```

---

## Custom Decoding

Dış dünyadan gelen veriyi `go` tipine `Unmarshal` ile çeviriyorduk. Bazen bu
işi biraz daha kontrollü yapmak gerekebilir. Bu durumda `json.NewDecoder`
kullanırız. Beklenmeyen bir alan geldiğinde hata yakalama ile bunu yakalayıp
istersek işlemi ilerletmeyebiliriz.

Şimdi [örneğe](../../src/13/json-custom-decode) bakalım;

https://go.dev/play/p/hCGfrUGRn5-

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// User holds user model data. User struct has json tags!
type User struct {
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Age       int        `json:"age"`
	BirthDate time.Time  `json:"birth_date"`
	Admin     bool       `json:"admin"`
	LastVisit *time.Time `json:"last_visit"`
}

// Users holds user slice of Users.
type Users []User

func main() {
	input := `[
		{
			"first_name": "Uğur",
			"last_name": "Özyılmazel",
			"age": 51,
			"birth_date": "1972-08-13T09:00:00.0+03:00"
		},
		{
			"first_name": "Ömer",
			"last_name": "Özyılmazel",
			"age": 41,
			"birth_date": "1982-08-13T14:00:00.0+03:00",
			"admin": true
		},
		{
			"this": "This",
			"is": "is",
			"fake": "fake"
		}
	]`

	b := []byte(input)
	d := json.NewDecoder(bytes.NewReader(b))
	// accepts io.Reader, convert byte slice -> io.Reader satisfier

	d.DisallowUnknownFields()
	// raise error for unknown fields

	var users Users

	// if err := json.Unmarshal(b, &users); err != nil {
	// 	log.Fatal(err)
	// }

	if err := d.Decode(&users); err != nil {
		fmt.Println(err) // json: unknown field "this"
	}

	fmt.Printf("%+v\n", users)
	// [{FirstName:Uğur LastName:Özyılmazel Age:51 BirthDate:1972-08-13 09:00:00 +0300 +0300 Admin:false LastVisit:<nil>} {FirstName:Ömer LastName:Özyılmazel Age:41 BirthDate:1982-08-13 14:00:00 +0300 +03 Admin:true LastVisit:<nil>} {FirstName: LastName: Age:0 BirthDate:0001-01-01 00:00:00 +0000 UTC Admin:false LastVisit:<nil>}]

	for _, user := range users {
		fmt.Printf("%+v\n", user)
		fmt.Println(user.LastVisit, user.LastVisit == nil)
	}
	// {FirstName:Uğur LastName:Özyılmazel Age:51 BirthDate:1972-08-13 09:00:00 +0300 +0300 Admin:false LastVisit:<nil>}
	// <nil> true
	// {FirstName:Ömer LastName:Özyılmazel Age:41 BirthDate:1982-08-13 14:00:00 +0300 +03 Admin:true LastVisit:<nil>}
	// <nil> true
	// {FirstName: LastName: Age:0 BirthDate:0001-01-01 00:00:00 +0000 UTC Admin:false LastVisit:<nil>}
	// <nil> true
	

	j, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n\n", j)
	// [
	//     {
	//         "first_name": "Uğur",
	//         "last_name": "Özyılmazel",
	//         "age": 51,
	//         "birth_date": "1972-08-13T09:00:00+03:00",
	//         "admin": false,
	//         "last_visit": null
	//     },
	//     {
	//         "first_name": "Ömer",
	//         "last_name": "Özyılmazel",
	//         "age": 41,
	//         "birth_date": "1982-08-13T14:00:00+03:00",
	//         "admin": true,
	//         "last_visit": null
	//     },
	//     {
	//         "first_name": "",
	//         "last_name": "",
	//         "age": 0,
	//         "birth_date": "0001-01-01T00:00:00Z",
	//         "admin": false,
	//         "last_visit": null
	//     }
	// ]
	
}
```

Son alan; `"this": "This",` ile başlayan kısım işlenmedi ama `user`’ın son
değeri `User` struct’ının zero-value (empty)’larıyla doldu. Yani bu decode
işleminde **3 tane elemanı** olan bir slice çıktı.