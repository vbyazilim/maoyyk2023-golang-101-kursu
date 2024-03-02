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

https://go.dev/play/p/7yav9dmBRb7

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// UserLevel is a custom type definition uses string for defining user level.
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

https://go.dev/play/p/E6L7GhKpJmu

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// UserLevel is a custom type definition uses string.
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

Şimdi [örneğe](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/13/json-marshal-custom-time) bakalım, `struct`
alanları içinde gezen, tag’leri bulan ve özelleştirilmiş `time` kullanan bir
yapı bulunuyor. `go`, `Marshal` işlemi yaparken, eğer custom type varsa ve bu
type’ın `MarshalJSON()` metotu varsa onu kullanıyor. `describeStruct` ise
`reflect` paketini kullanarak struct içinde geziyor.

https://go.dev/play/p/7ny1yu09idp

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

// CustomTime is a custom type definition uses time.Time, uses custom marshal format.
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

## Custom Encoding

Aslında `pretty print` yaptığımız kısımda `marshalStruct(v any)` fonksiyonu
içinde kullanmıştık, ek olarak html karakterlerini de escape etmeyi görelim;

https://go.dev/play/p/PaLSl8ya-ql

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	data := map[string]any{
		"message": `this is <strong>bold</strong> text`,
	}

	buffer := new(bytes.Buffer)        // create buffer, returns pointer.
	encoder := json.NewEncoder(buffer) // point to buffer
	encoder.SetEscapeHTML(true)        // enable html escape
	encoder.SetIndent("", "    ")      // indent 4 spaces each key

	// serialize data to buffer
	if err := encoder.Encode(data); err != nil {
		log.Fatal(err)
	}

	fmt.Println(buffer)
	// fmt.Println(buffer.String())
	// fmt.Printf("%s\n", buffer)
	// {
	//     "message": "this is \u003cstrong\u003ebold\u003c/strong\u003e text"
	// }

	buffer = new(bytes.Buffer)        // create buffer, returns pointer.
	encoder = json.NewEncoder(buffer) // point to buffer
	encoder.SetEscapeHTML(false)      // disable html escape
	encoder.SetIndent("", "    ")     // indent 4 spaces each key

	// serialize data to buffer
	if err := encoder.Encode(data); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buffer)
	// {
	//     "message": "this is <strong>bold</strong> text"
	// }
}
```


---

## Custom Decoding

Dış dünyadan gelen veriyi `go` tipine `Unmarshal` ile çeviriyorduk. Bazen bu
işi biraz daha kontrollü yapmak gerekebilir. Bu durumda `json.NewDecoder`
kullanırız. Beklenmeyen bir alan geldiğinde hata yakalama ile bunu yakalayıp
istersek işlemi ilerletmeyebiliriz.

Şimdi [örneğe](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/13/json-custom-decode) bakalım;

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

---

## Generic Interface

Elimizde raw string olduğunda bunu `map[string]any`’e cast ederek, içindeki
diğer verilere de ulaşabiliriz.

Eğer bazı alanların opsiyonel olmasını istiyorsak, özellikle `struct`’a
`Unmarshal` ederken, struct içine bir tane joker alan koyup bunu
`map[string]any` yapıyoruz ve eksik alanları bu map’ten siliyoruz:

https://go.dev/play/p/u1mm5nr-jJx

[Örnek](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/13/json-generic-interface)

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// User holds user data
type User struct {
	Name      string `json:"name"`
	Email     string
	Age       int
	Optionals map[string]any `json:"-"`
}

func main() {
	incoming := `{
		"name": "Uğur",
		"email": "vigo@example.com",
		"age": 51,
		"foo": 1,
		"bar": "2"
	}`

	u := User{}
	if err := json.Unmarshal([]byte(incoming), &u.Optionals); err != nil {
		log.Fatal(err)
	}

	fmt.Println("u.Optionals", u.Optionals)
	// u.Optionals map[age:51 bar:2 email:vigo@example.com foo:1 name:Uğur]

	if v, ok := u.Optionals["name"].(string); ok {
		u.Name = string(v)
		delete(u.Optionals, "name")
	}
	if v, ok := u.Optionals["email"].(string); ok {
		u.Email = string(v)
		delete(u.Optionals, "email")
	}
	if v, ok := u.Optionals["age"].(float64); ok {
		u.Age = int(v)
		delete(u.Optionals, "age")
	}

	fmt.Printf("u: %+v\n", u)
	// u: {Name:Uğur Email:vigo@example.com Age:51 Optionals:map[bar:2 foo:1]}

	if u.Email != "" {
		fmt.Println("u.Email", u.Email)
		// u.Email vigo@example.com
	}
	if u.Age != 0 {
		fmt.Println("u.Age", u.Age)
		// u.Age 51
	}

	if len(u.Optionals) > 0 {
		fmt.Printf("you have %d invalid field(s)\n", len(u.Optionals))
		// you have 2 invalid field(s)

		for v := range u.Optionals {
			fmt.Printf("%q\n", v)
		}
		// "foo"
		// "bar"
	}
}
```

`Unmarshal` işlemine başlamadan önce, gelen verinin json’a uygun olup
olmadığını doğrulamak için `json.Valid` kullanıyoruz;

https://go.dev/play/p/Z0a7mVQpXoU

```go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	goodJSON := `{"example": 1}`
	badJSON := `{"example":2:]}}`

	fmt.Println("goodJSON is valid?", json.Valid([]byte(goodJSON)))
	// goodJSON is valid? true

	fmt.Println("badJSON is valid?", json.Valid([]byte(badJSON)))
	// badJSON is valid? false
}
// goodJSON is valid? true
// badJSON is valid? false
```

---

## Streaming

Uygulananız başka bir servise rest-api üzerinden `json` ile veri çekiyor fakat
gelen veri json array şeklinde değil;

https://go.dev/play/p/H95BTncScTb

[Örnek](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/13/json-streaming)

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// Person represents person model.
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// People represents collection on Person slice.
type People []Person

func main() {
	// bulk data, looks like json array?
	incoming := `
	{"name": "Fred", "age": 40}
	{"name": "Mary", "age": 21}
	{"name": "Pat", "age": 30}
	`

	decoder := json.NewDecoder(strings.NewReader(incoming))

	var b bytes.Buffer
	encoder := json.NewEncoder(&b)

	var p Person
	var people People

	for decoder.More() {
		if err := decoder.Decode(&p); err != nil {
			log.Print("decode err", err)
			continue
		}

		fmt.Printf("p: %+v\n", p)
		// p: {Name:Fred Age:40}
		// p: {Name:Mary Age:21}
		// p: {Name:Pat Age:30}

		people = append(people, p)

		if err := encoder.Encode(p); err != nil {
			log.Panic("encode err", err) // do not panic!
		}
	}

	fmt.Println(b.Bytes())
	// [123 34 110 97 109 101 34 58 34 70 114 101 100 34 44 34 97 103 101 34 58 52 48 125 10 123 34 110 97 109 101 34 58 34 77 97 114 121 34 44 34 97 103 101 34 58 50 49 125 10 123 34 110 97 109 101 34 58 34 80 97 116 34 44 34 97 103 101 34 58 51 48 125 10]

	// fmt.Println(string(b.Bytes()))
	fmt.Println(b.String())
	// {"name":"Fred","age":40}
	// {"name":"Mary","age":21}
	// {"name":"Pat","age":30}

	fmt.Printf("people: %+v\n", people)
	// people: [{Name:Fred Age:40} {Name:Mary Age:21} {Name:Pat Age:30}]

	j, _ := json.MarshalIndent(people, "", "    ")
	fmt.Printf("%s\n", j)
	// [
	//     {
	//         "name": "Fred",
	//         "age": 40
	//     },
	//     {
	//         "name": "Mary",
	//         "age": 21
	//     },
	//     {
	//         "name": "Pat",
	//         "age": 30
	//     }
	// ]
}
```
