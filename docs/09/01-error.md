# Bölüm 09/01: `error`

`error` standart kütüphaneyle gelen bir `interface` aslında:

```go
type error interface {
    Error() string
}
```

Herhangi bir tip, eğer `Error() string` metotuna sahipse o artık go için
gerçek (valid) bir error değeri olur. Go’da hata mesajları mutlaka 
**küçük harfle başlar!**

---

## Custom Error Types

Go paketleri içinde kendine özel bir kısım `error` değerleri ile gelir. Bizler
de geliştirme yaparken, paketlerimize özel `error` tipleri implemente ederiz.

https://go.dev/play/p/nd_kcYA9FAc

```go
package main

import (
	"errors"
	"fmt"
)

// MyError is a custom type.
type MyError struct {
	Message string
}

// Error implements error interface.
func (m MyError) Error() string {
	return m.Message
}

// generateError return an error, a type which implements error interface.
func generateError(s string) error {
	return MyError{
		Message: s,
	}
}

func fakeFunc1() error {
	return errors.New("this is a fake error")
}

func fakeFunc2() error {
	return MyError{"this is a fake error"}
}

func main() {
	if err := generateError("this is a test error"); err != nil {
		fmt.Printf("generateError -> %[1]T\n", err)
	}

	if err := fakeFunc1(); err != nil {
		fmt.Printf("fakeFunc1 -> %[1]T\n", err)
	}

	if err := fakeFunc2(); err != nil {
		fmt.Printf("fakeFunc2 -> %[1]T\n", err)
	}
}
// generateError -> main.MyError
// fakeFunc1 -> *errors.errorString
// fakeFunc2 -> main.MyError
```

İşte milyon dolarlık mülakat sorusu:

> Neden go’da error kontrolü `nil` olup olmadığına bakılarak yapılır?

Neden;

- `true` ya da `false` ?
- `== 1` ya da `== 0` ?

Cevabı sizden bekliyorum :)

https://go.dev/play/p/VOE28-nwG88

```go
package main

import "fmt"

// errKind is a custom type definition uses int.
type errKind int

// create automatic constants with iota.
const (
	_              errKind = iota // skip first value, which is 0
	invalidUser                   // 1
	invalidRequest                // 2
)

// customError is a custom for representing error.
type customError struct {
	kind errKind
}

// Error implements error interface.
func (e customError) Error() string {
	switch e.kind {
	case invalidUser:
		return "invalid user"
	case invalidRequest:
		return "invalid request"
	}

	return "unknown error"
}

var (
	errInvalidUser    = customError{kind: invalidUser}
	errInvalidRequest = customError{kind: invalidRequest}
	errUnknown        = customError{kind: 9999}
)

func checkUser(name string) error {
	switch name {
	case "admin":
		return errInvalidUser
	case "hack-attack":
		return errInvalidRequest
	case "":
		return errUnknown
	}

	return nil
}

func main() {
	// err: invalid user , type: main.customError
	if err := checkUser("admin"); err != nil {
		fmt.Printf("err: %[1]v , type: %[1]T\n", err)
	}

	// err: invalid request , type: main.customError
	if err := checkUser("hack-attack"); err != nil {
		fmt.Printf("err: %[1]v , type: %[1]T\n", err)
	}

	// err: unknown error , type: main.customError
	if err := checkUser(""); err != nil {
		fmt.Printf("err: %[1]v , type: %[1]T\n", err)
	}

	// all ok!
	if err := checkUser("vigo"); err != nil {
		fmt.Printf("err: %[1]v , type: %[1]T\n", err)
	}
}
```

---

## Wrapping

Kullandığımız paket bize `error` döndü, ama biz de mesaja ilave bir kısım
bilgiler eklemek istiyoruz, bu durumda gelen `error`’ü kendi mesajımızla
**wrap** ederiz, yani sarmalarız:

https://go.dev/play/p/W53nyqebqDB

```go
package main

import (
	"errors"
	"fmt"
	"log"
)

var errUnknown = errors.New("unknown error")

// getUser is a phony function, returns an error.
func getUser() error {
	// do real operations here.
	return errUnknown
}

// getStats is also a fake function, calls getUser and wraps incoming error.
func getStats() error {
	if err := getUser(); err != nil {
		return fmt.Errorf("getStats has an error: %w", err)
	}
	return nil // all goes ok!
}

func main() {
	if err := getStats(); err != nil {
		log.Fatal(err) // prints error and calls sys.exit(1)
	}
}

// 2023/02/10 22:04:14 getStats has an error: unknown error
// exit status 1
```

Çalışma sırası;

1. `main` çalışır, `getStats()` çağırılır
1. `getStats()` -> `getUser()` çağırılır
1. `getUser()` gerite `errUnknown` döner: `"unknown error"`
1. `getStats()` gelen `errUnknown` hatasını şu mesajla sarmalar: `"getStats has an error"`

---

## Unwrapping

Sarılı mesajı geri sarmak için yani **wrap** edilmiş `error`’ü **unwrap**
etmek için kullanırız. Az önceki örneğe ufak bir modifikasyon yapalım:

```go
func main() {
	if err := getStats(); err != nil {
		fmt.Printf("err: %q\n", err)
		fmt.Printf("unwrapped err: %q\n", errors.Unwrap(err))
	}
	// err: "getStats has an error: unknown error"
	// unwrapped err: "unknown error"
}
```

`errors.Unwrap` fonksiyonu bu işlemi yapmamızı sağladı.

---

## Tip Kontrolleri

Uygulama çalışırken bir kısım hata kontrolleri, geri dönen `error` değerli
kodun doğal akışı içinde normal olaylardır. Bazı durumlarda dönen `error`
tipine göre başka şeyler yapmak gerekir.

Örneğin veritabanında kayıt bulunamadıysa `ErrRecordNotFound`, kayıt eklerken
**unique index** hatası aldıysak `ErrUniqueConstraint` gibi hata şekline göre
farklı `error` değerleri dönmek gerekir.

Uygulama bu dönen tiplere göre isteyene (api consumer ya da client) makul bir
bilgi dönmelidir. Bazen hatasına göre log’a ekleme ya da pas geçme, ya da
slack kanalına mesaj atmak gibi başka aksiyonlar da alınabilir.

Bu durunlarda `errors.Is` ve `errors.As` yardımımıza koşar:

## `errors.Is`

Hatayı `value` (değeri) anlamında kontrol etmek için kullanılır.

https://go.dev/play/p/omoy_TntsMQ

```go
package main

import (
	"errors"
	"fmt"
	"os"
)

var (
	errCustom = errors.New("custom error")

	// this is an error from standard library.
	errPath = &os.PathError{
		Err: errors.New("path error"),
	}
)

// raiseErrors will generate errors.
func raiseErrors(n int) error {
	switch n {
	case 1:
		return errCustom
	case 2:
		return errPath
	default:
		return nil
	}
}

func main() {
	if err := raiseErrors(1); err != nil {
		if errors.Is(err, errCustom) {
			fmt.Println("err, value is errCustom")
		} else {
			fmt.Println("err, value is %v", err)
		}
	}

	if err := raiseErrors(2); err != nil {
		if errors.Is(err, errPath) {
			fmt.Println("err, value is errPath")
		} else {
			fmt.Println("err, value is %v", err)
		}
	}

	if err := raiseErrors(0); err != nil {
		fmt.Println("err, value is %v", err)
	}
}

// err, value is errCustom
// err, value is errPath
```

## `errors.As`

Hatayı `type` olarak kontrol etmeye yarar.

https://go.dev/play/p/JhM3JXdy-15

```go
package main

import (
	"errors"
	"fmt"
)

// customError is a custom type definition uses string, will
// implement error interface.
type customError string

func (c customError) Error() string {
	return string(c)
}

var (
	errSpecial = customError("special error")
	errOther   = customError("other error")
)

// errorizer generates fake errors.
func errorizer(n int) error {
	switch n {
	case 1:
		return errSpecial
	case 2:
		return errOther
	default:
		return nil
	}
}

func main() {
	if err := errorizer(1); err != nil {
		var cErr customError

		if errors.As(err, &cErr) {
			fmt.Println("errSpecial", cErr == errSpecial)
			fmt.Println("errOther", cErr == errOther)
			fmt.Println()
		}
	}

	if err := errorizer(2); err != nil {
		var cErr customError

		if errors.As(err, &cErr) {
			fmt.Println("errSpecial", cErr == errSpecial)
			fmt.Println("errOther", cErr == errOther)
			fmt.Println()
		}
	}
}
```

İlginç olan kısım burası:

```go
var cErr customError
if errors.As(err, &cErr) {
	// ...    
}
```

`errors.As`, `err` hata zincirinde `customError`’ü bulmaya çalışır, bulursa `cErr`’ün içini
set eder ve `true` döner, bulamazsa `false` döner. Eğer `true` dönmüşse artık `cErr`’de
ilgili hatanın tüm metotları vs mevcut olur.

Dikkat ettiyseniz pointer’ın pointer’ını verdik `errors.As`’e:

    +------+
    | 0001 | *0010
    +------+
    :
    :
    +------+
    | 0010 | *customError
    +------+

İlk pointer’dan `customError`’ün nereye yerleştiğini buluyor, 2. pointer’ı
kullanarak içini dolduruyor.

Go’nun doc’undan bir örnek:

https://go.dev/play/p/gpeCpbUvib_f

```go
package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

// source from /opt/homebrew/Cellar/go/1.19.5/libexec/src/io/fs/fs.go
// PathError records an error and the operation and file path that caused it.
// type PathError struct {
// 	Op   string
// 	Path string
// 	Err  error
// }

func main() {
	if _, err := os.Open("non-existing"); err != nil {
		// fs.PathError is a struct type
		// need to dereference with *
		var pathError *fs.PathError

		// passing pointer of pointer
		// pathError holds address,
		// &pathError is the address of address, errors.As need it!
		if errors.As(err, &pathError) {
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}
}

// Failed at path: non-existing
```

---

## `panic` ve `recover`

`panic` alsında **sadece geliştirme yaparken** kullanılması gereken bir
built-in fonksiyon. Asıl amacı bize **stack trace**’i yani hata oluştu,
oluşurken sırsıyla neler çağırıldı, patlak nerede oluştu bunu görmemizi sağlar.

https://go.dev/play/p/PKHGlDOjL0H

```go
package main

import (
	"errors"
)

func main() {
	if err := errors.New("err is here"); err != nil {
		panic(err)
	}
}

// panic: err is here
//
// goroutine 1 [running]:
// main.main()
// 	untitled:9 +0x54
// exit status 2
```

Hatta Rob Pike derki:

> Don’t panic

Bazen henüz yazılmamış, daha sonra yazılacak kod için **placeholder** (yer
tutucu) olarak kullanılır, metotu yazarken henüz kodu planlamadık ama kodu
**compile** etmek istiyoruz, bu durumlarda da kullanırız:

```go
package main

import "fmt"

func greet(name string) error {
	panic("not implemented")
}

func main() {
	fmt.Println("hello")
}
```

yukarıdaki kod sorunsuz derlenir.

`panic` ile ilgili esas sıkıntı, `defer` edilen fonksiyonlar çalışmaya devam
eder, bu sayede `recover` ile sanki hiçbir şey olmamış gibi hayat devam eder:

https://go.dev/play/p/fISb2H5TNL_8

```go
package main

import "fmt"

func makePanic() {
	panic("oh my god!")
}

func main() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("recover", p)
		}
	}()

	makePanic()
}

// recover oh my god!
```

Bazı web framework’leri bu taktiği kullanır, hatta bazı kötü tasarlanmış
rest-api servisleri hatayı tek noktadan yönetmek için, hataları `panic` ile
yapar ve `recover` ile yola devam eder. Böyle yazılmış bir proje içindeyseniz
mutlaka arkadaşlarınızı uyarın.

`panic` ve `recover` sadece geliştirme esnasında kullanılmalıdır!

---

## Yaygın Pratikler

- Go’da nesne yönelimli dillerde olduğu gibi **exception handling** diye bir yaklaşım yoktur
- `error` bir tip’tir, diğer tipler gibi işlenmelidir
- Asla `error`’leri yutmamalı, **ignore** etmemelidir, mümkünse hataları `_` **blank identifier**’a (linux /dev/null gibi) göndermemek gerekir
- Her zaman **fail fast** yaklaşımı olmalıdır, `error` yakalandığı an (nil değilse) fonksiyon **return** etmelidir.
- Hata mesajları açık ve anlaşılır olmalıdır, hangi paketten hangi fonksiyondan geldiği belirtilmelidir
- Hata yakalandığı zaman **bir kere** işlenmelidir (handle only once)

https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully

```go
func AuthenticateRequest(r *Request) error {
	err := authenticate(r.User)
	if err != nil {
		return err // ??? bu hata kimden geldi? 
	}
	return nil
}
```

şöyle olsa;

```go
func AuthenticateRequest(r *Request) error {
	err := authenticate(r.User)
	if err != nil {
		return fmt.Errorf("authenticate failed: %v", err) // wrap errors with an extra message
	}
	return nil
}
```

Kodu şu şekilde;

```go
func Write(w io.Writer, buf []byte) error {
	_, err := w.Write(buf)
	if err != nil {
		// annotated error goes to log file
		log.Println("unable to write:", err)

		// unannotated error returned to caller
		return err
	}
	return nil
}
```

şöyle yazmak daha iyidir:

```go
func Write(w io.Write, buf []byte) error {
	_, err := w.Write(buf)
	return errors.Wrap(err, "write failed") // <-- Wrap method from github.com/pkg/errors package
}
```
