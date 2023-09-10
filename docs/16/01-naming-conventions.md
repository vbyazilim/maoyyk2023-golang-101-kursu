# Bölüm 16/01: İsimlendirme Kuralları

Kuralımız hep şu:

> Conventions over configuration / Coding by convention

Yani, şartlar ne olursa olsun, **kafamıza göre iş yapmayacağız** ve
geleneklere bağlı kalacağız. Her zaman belirlenmiş standartlara sadık kalmak
olarak da izah edebiliriz.

https://en.wikipedia.org/wiki/Convention_over_configuration

Bir yazılım ürününün ya da kütüphanesinin, kullanıcılara bol miktarda
yapılandırma seçeneği sunmak yerine, belirli varsayılan davranışlara veya
'standartlara' sahip olmasını tercih eden bir yaklaşımdır.

Bu bakımdan, isimlendirme yaparken bu kurallara dikkat etmemiz gerekir.

---

## Değişken Adları

Değişken adı, değişkenin **tuttuğu değeri tarif etmeli**, değişkenin tipini
değil! Yanlış (kötü) örnekler;

```go
var usersMap     map[string]*User
var companiesMap map[string]*Company
var productsMap  map[string]*Product
var usersList    []User
```

Doğrusu:

```go
// ...Maps son ekine ihtiyaç yok

var users     map[string]*User      // içinde User’lar olan koleksiyon (map ya da slice)
var users     []User

var companies map[string]*Company   // içinde Company’ler olan koleksiyon (map ya da slice)
var companies []Company             // içinde Company’ler olan koleksiyon (map ya da slice)
var products  []Product             // içinde Product’lar olan koleksiyon (map ya da slice)
```

Tahmin edilebilir, anlaşılır adlar kullanın:

- `i`, `j`, `k` gibi kısa değişkenleri `for` loop’larında
- `n` sayaç, toplam ya da miktar’ı temsil ederken
- `map`’lerde, `v` -> `value`, `k` -> `key` gibi...
- `a`, `b` aynı tipteki nesneleri ifade ederken (karşılaştırma vs), yerleri de değişebilir
- `x`, `y` karşılaştırma yaparken oluşturulan yerel (local) değişkenlerin genel adı
- `s` genelde `string` tipindekilerin kısaltması olarak
- Koleksiyonlar (map, slice, array) mutlaka çoğul olmalı

---

## Fonksiyon Adları

Fonksiyonlar **döndürdükleri sonuca göre** adlandırılmalıdır.

- Mutlaka karakter ile başlamalıdır, sayı ile başlamaz, içinde `<space>`
  karakteri olamaz
- Exportable’lar büyük harfle başlar ve mutlaka `comment` olarak dokümanı yazılır
- Büyük/küçük harfe duyarlıdır (case-sensitive)

```go
func Add(a, b) int {}
// describes only the operation
```

ama daha da iyisi;

```go
func Sum(a, b) int {}  // sonuç ne? iki sayının toplamından çıkan yeni değer
// returned thing is a sum of a and b...
// this describes the result, not the operation...
```

Kötü örnek:

```go
package grpc

func NewClient() *Client
func NewClientWithTimeout(timeout time.Duration) *Client
```

Bu şekilde daha iyi yapılabilir (functional options pattern);

```go
type Option func(*Client) *Client

func NewClient(opts ...Option)

func WithTimeout(timeout time.Duration) func(c *Client) *Client

client := grpc.NewClient(grpc.WithTimeout(10 * time.Seconds))

// same constructor with different options
```

---

## Metot Adları

Yaptıkları **eylemi anlatacak şekilde** adlandırılmalıdır. Fonksiyon
adlandırmasının **tam tersidir**:

https://go.dev/play/p/lYSUe8VC-qG

```go
package main

import "fmt"

type user struct {
	email    string
	password string
	fullName string
}

// Email is a getter for user.email
func (u user) Email() string {
	return u.email
}

// SetEmail is a setter for user.email
func (u *user) SetEmail(email string) {
	u.email = email
}

// resetPassword resets user's password
func (u *user) resetPassword() error {
	fmt.Println("example reset password")
	u.password = "reset"
	return nil
}

func main() {
	u := &user{}
	u.SetEmail("vigo@me.com")
	u.resetPassword()

	fmt.Println("email", u.Email())
	fmt.Printf("%+v\n", u)
}
// example reset password
// email vigo@me.com
// &{email:vigo@me.com password:reset fullName:}
```

---

## Interface Adları

`interface` için davranışları belirler dedik, bu bakımdan da sonuna `er`
takısı alır; içinde `Read()` tanımı olan paketin `XxxReader` olma ihtimali
yüksek. Sadece `Read()` ve `Write()` (yani sadece 2 fonksiyon) varsa
`XxxReadWriter`, eğer `Read()`, `Write()`, `Count()` (3 fonksiyon) varsa;
`XxxReadWriteCounter` olabilir.

Kodun testini yazmak için illaki `interface` tanımı yapmak ve bu
`interface`’leri test içinde kandırmak (mock’lamak) gerekir. Bu bakımdan da;
gerçek dünyada, go’nun standart kütüphanesindeki gibi bir, iki ya da maksimum
3 tane fonksiyon tanımı olan interface yapmak neredeyse imkansıza yakındır.

Veritabanı katmanı için (storage) tanım yapıyorsunuz:

```go
type Storer interface {  // şimdi içeride Store diye bir fonksiyon olmasını bekliyoruz
    Get()
    Create()
    Update()
    Delete()
    List()
}
```

Belkide;

```go
type Getter interface {
    Get()
}

type Creater interface {
    Create()
}

type Updater interface {
    Update()
}

type Deleter interface {
    Delete()
}

type Lister interface {
    List()
}

type Storage interface {
    Getter
    Creater
    Updater
    Deleter
    Lister
}

// ya da
type GetCreateUpdateDeleteLister interface {
    Getter
    Creater
    Updater
    Deleter
    Lister
}
```

Ben olsam; `type Storer interface` ya da `type Storage interface` ile ilerlerim.

---

## Paket Adları

Belkide go’daki en zor, en kritik isimlendirme paket isimlendirmesidir. İsmi,
paketin amacını anlatmalıdır. 

- İyi bir paket adında sadece harfler olur; `strings`, `strconv`, `fmt`, `io`, `os`...
- `stringUtils`, `foo_tools`, `x11Package` gibi isimler **olmaz**!
- `base`, `common`, `util`, `helpers` gibi genel-geçer paket adı olmamalı ***
- Paket adı, olası güzel değişken adı kullanımına engel olmamalı *
- Paketinizi içerdiklerine göre değil, sağladıklarına göre adlandırın
- Sınıfa ya da türe göre adlandırmayın
- Paket düzeyindeki değişkenler, tüm programı kapsadığı için daha uzun
  tanımlayıcıları (method adı, değişken adı vs...) olmalıdır.
- Başka paketlerin ya da fonksiyon / metotların da kullanabileceği isimlerden kaçının
- İçeriği bağlamında doğru miktarda bilgi taşıyan **en kısa adı** kullanın

```go
import "github.com/pkg/term/v2" // kötü
import "github.com/pkg/v2/term" // daha iyi

func WriteLog(context context.Context, message string) // Don’t, context is stolen
func WriteLog(ctx context.Context, message string)     // Good
```

Peki kötü paket adları nasıl olur? Anlamsız paket adlarından kaçınmak lazım.
`util`, `common` ya da `misc` adlı paketler, kullanıcıya paketin ne içerdiği
konusunda hiçbir fikir vermez. Bu, kullanıcının paketi kullanmasını
zorlaştırır ve paketin bakımını (maintenance) zorlaştırır. Örneğin;

```go
package util
func NewStringSet(...string) map[string]bool {...}
func SortStringSet(map[string]bool) []string {...}
```

olsa aşağıdaki gibi kullanılır:

```go
set := util.NewStringSet("c", "a", "b")
fmt.Println(util.SortStringSet(set))
```

Halbuki;

```go
package stringset
func New(...string) map[string]bool {...}
func Sort(map[string]bool) []string {...}
```

şeklinde olsa;

```go
set := stringset.New("c", "a", "b")
fmt.Println(stringset.Sort(set))
```

olur ve daha **idiomatic** (dilin özelliklerini taşıyan) bir hal alır! Yeni
özellikler geliştirmeler geldikçe;

```go
package stringset

type Set map[string]bool

func (s Set) Sort() []string {...}  // artık bu method’a dönüştü!


func New(...string) Set {...}
```

ilerler..

## Kaynaklar

- https://go.dev/blog/package-names
- https://rakyll.org/style-packages/
- https://go.dev/blog/package-names#bad-package-names-h2
- https://github.com/vigo/stringutils-demo

---

Paket isminin düşünürken hep kafamda şu anı canlandırırım: `paketAdı.New()`
ile çağıracağım; örneğin 3. parti bir servis için client geliştirmesi yapmam
gerekiyor; servis sağlayıcı adı: `acme`, ben de bu servisten kullanıcının
upload ettiği dosyaların listesini çeken bir client yazacağım;

- Aklıma `http` paketi geliyor, `http.Client` var;
- `xxxxx.New` dediğim zaman bana acme http client vermeli
- `acme` diye firmanın çıkarttığı bir paket var mı?
- `acmeclient.New`

**Paket adı, olası güzel değişken adı kullanımına engel olmamalı**

https://dave.cheney.net/2019/01/29/you-shouldnt-name-your-variables-after-their-types-for-the-same-reason-you-wouldnt-name-your-pets-dog-or-cat

Örneğin `context` paket adı yüzünden, kullanıldığı yerlede `var context = ...`
şeklinde bir kullanım yapamıyoruz çünkü paketi içeri aldığımı için `context`
anahtar kelimesi artık o kapsam içinde kullanılır durumda.

Bu bakımdan da `func WriteLog(ctx context.Context, message string)` olduğu
gibi `ctx` şeklinde kullanmak zorunda kalıyoruz.
