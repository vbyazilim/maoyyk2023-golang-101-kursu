# Bölüm 14/01: Test

## Test Nedir? Neden Yazılır?

Yazdığımız kodun, fonksiyonun, metotların çalıştığını garanti etmek için
uygulanan yöntemdir. Basit mantık şu, fonksiyon girdi (input) alır, bir iş
yapar ve çıktı (output) üretir. Toplama işlemi yapan bir fonksiyon, iki tane
sayı alır, işlemi yapar ve geriye sonuç döner:

    girdi: 1,2  ->  1+2 => çıktı: 3
    girdi: 0,0  ->  0+0 => çıktı: 0
    girdi: -1,1 -> -1+1 => çıktı: 0

İşte tüm bu varyasyonları (use case’leri) denediğimiz kod parçasına test kodu,
birim test kodu ya da **unit testing** denir.

Go, kendi içinde built-in test araçlarıyla birlikte gelir. `testing` paketi bu
iş için kullanılır. Sonu `_test.go` ile biten dosyalar testlerin yazıldığı
dosyalardır. Go, test’leri çalıştırırken otomatik olarak bu dosyaları işler.

Bizden istenen bir fonksiyon olsa; biz o fonksiyona kimi zaten bir tane, kimi
zaman `n` tane isim versek (string), fonksiyon bizi selamlasa;

```go
SayHi()                 // hi everybody!
SayHi("uğur")           // hi uğur 
SayHi("uğur", "erhan")  // hi uğur
                        // hi erhan
```

Haydi şimdi koda bakalım:

[Örnek](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/14/test-greet)

`greet/greet.go`:

```go
package greet

import "strings"

// SayHi greets given names.
func SayHi(names ...string) string {
	if len(names) == 0 {
		return "hi everybody!"
	}
	out := make([]string, len(names))
	for i, name := range names {
		out[i] = "hi " + name + "!"
	}

	return strings.Join(out, "\n")
}
```

`main.go`:

```go
package main

import (
	"fmt"

	"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet"
)

func main() {
	fmt.Println(greet.SayHi())       // hi everybody!
	fmt.Println(greet.SayHi("uğur")) // hi uğur!
	fmt.Println(greet.SayHi("uğur", "erhan"))
	// hi uğur!
	// hi erhan!
}
```

Şimdi bu fonksiyonun gerçekten bizim verdiğimiz girdileri doğru işleyip
işlemediğini görelim. `greet_test.go` isminde bir dosya oluşturuyoruz ve;

```go
package greet_test

import (
	"fmt"
	"testing"

	"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet"
)

func TestSayHi(t *testing.T) {
	want := "hi vigo!"
	got := greet.SayHi("vigo")

	if got != want {
		t.Errorf("want: %v; got: %v", want, got)
	}
}

func TestSayHiWithNoArgs(t *testing.T) {
	want := "hi everybody!"
	got := greet.SayHi()

	if got != want {
		t.Errorf("want: %v; got: %v", want, got)
	}
}

func TestSayHiWithArgs(t *testing.T) {
	want := "hi vigo!\nhi turbo!\nhi max!"
	got := greet.SayHi("vigo", "turbo", "max")

	if got != want {
		t.Errorf("want: %v; got: %v", want, got)
	}
}

// This is example of single argument usage.
func ExampleSayHi() {
	fmt.Println(greet.SayHi("vigo"))
	// Output: hi vigo!
}

// This is example of no argument usage.
func ExampleSayHi_withNoArg() {
	fmt.Println(greet.SayHi())
	// Output: hi everybody!
}

// This is example of with many arguments usage.
func ExampleSayHi_withArgs() {
	fmt.Println(greet.SayHi("vigo", "turbo", "max"))
	// Output: hi vigo!
	// hi turbo!
	// hi max!
}
```

---

## Test Nasıl Çalıştırılır

Test etmek istediğimiz senaryoları `Test` kelimesiyle başlatan test
fonksiyonları şeklinde yazıyoruz. `Test<NEYİ_TEST_EDIYORUZ?>`. Bizim için 3 konu vardı;

1. Hiç argüman verilmezse `TestSayHiWithNoArgs`
1. Tek argüman verilirse `TestSayHi`
1. N tane argüman verilirse `TestSayHiWithArgs`

Aynı test fonksiyonuları gibi `Example` kelimesiyle başlayan ve aynı test gibi
çalışan örnekler de ekledik. Bu örnekler dokümantasyon için çok faydalıdır.

Şimdi testleri çalıştıralım:

```bash
$ # go test <PAKET-ADI>
$ # go test <TÜM-PAKETLER>

$ go test ./...        # go projesi altındaki tüm paketlerin testlerini çalıştır

$ go list ./...        # go projesi altındaki tüm paketleri listele
$ go test github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet
ok  	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet	(cached)

# verbose mode
$ go test -v github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet
=== RUN   TestSayHi
--- PASS: TestSayHi (0.00s)
=== RUN   TestSayHiWithNoArgs
--- PASS: TestSayHiWithNoArgs (0.00s)
=== RUN   TestSayHiWithArgs
--- PASS: TestSayHiWithArgs (0.00s)
=== RUN   ExampleSayHi
--- PASS: ExampleSayHi (0.00s)
=== RUN   ExampleSayHi_withNoArg
--- PASS: ExampleSayHi_withNoArg (0.00s)
=== RUN   ExampleSayHi_withArgs
--- PASS: ExampleSayHi_withArgs (0.00s)
PASS
ok  	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet	0.406s
```

`(cached)` ifadesi, bu testin daha önce çalıştığı ve kodda bir değişiklik
olmadığını, bu bakımdan da testin aslında tekrar çalışmadığını ifade eder. Eğer
cache’i silmek istersek;

```bash
$ go clean -testcache
```

---

## Examples ve godoc

yeterlidir. Hemen `Example`’lar ne işe yarıyor onu görelim. Öncelikle
`godoc`’u kuralım:

```bash
$ go install golang.org/x/tools/cmd/godoc@latest
$ godoc -http=:6060  # doc sunucusu 6060 portundan çalışacak
```

sonra;

http://127.0.0.1:6060/pkg/github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet/

Example’ları yazarken;

```go
func Example() { ... }
func ExampleF() { ... }
func ExampleT() { ... }
func ExampleT_M() { ... }
```

Yani;

- `Example()`
- `ExampleSayHi()` -> `F()` => `SayHi()`’a denk geliyor
- `ExampleT()` -> Eğer custom bir type yazıp test etmek istesek, `ExampleUser()` gibi
- `ExampleT_M()` -> Custom type’ın metotu yani `ExampleUser_List()` gibi..

durumlarda kullanıyoruz. Keza;

```go
func Example_suffix() { ... }
func ExampleF_suffix() { ... }
func ExampleT_suffix() { ... }
func ExampleT_M_suffix() { ... }
```

şeklinde de kullanabiliyoruz. Aynı şekilde biz de;

- `ExampleSayHi_withNoArg()`
- `ExampleSayHi_withArgs()`

kullandık.

```go
// Output: .....
```

ile çıktısını yazıyoruz. Bu sayede aynı test gibi çalışmış oluyorlar. Eğer bir
örneğin `Output` **comment**’i yoksa kod derleniyor **ama çalıştırılmıyor**.

`Example` sayesinde yazdığımız kodu kullanacak kişiye copy/paste yaparak örnek
kod parçasını hemen alıp kendi koduna entegre edecek bir kolaylık sağlıyoruz.

Test fonksiyonlarını teker teker de çalıştırabiliriz:

```bash
$ go test -v -run ^TestSayHi$ ./...
# go projesi altındaki tüm paketleri tara, "TestSayHi" olanı çalıştır (regex match)

$ go test -v -run ^TestSayHi ./...
# go projesi altındaki tüm paketleri tara, "TestSayHi" ile başlayanları çalıştır (regex match)

?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/02/01-init	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/04/05-struct-custom-tag	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/04/05-struct-field-access	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/04/05-struct-field-access/person	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/04/05-struct-field-access-getter	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/04/05-struct-field-access-getter/person	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/04/05-struct-field-alignment	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/04/05-struct-validate	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/11/generics-constraints	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/11/generics-custom-types	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/11/generics-function-calls	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/11/generics-function-calls-and-types	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/11/generics-function-calls-and-types2	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/11/generics-functions	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/11/generics-functions-interface	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/11/generics-in-maps	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/11/generics-in-structs	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/12/reflect-clearvalue	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/12/reflect-typecheck	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/13/json-generic-interface	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/13/json-custom-decode	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/13/json-marshal-custom-time	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/13/json-streaming	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet	[no test files]
=== RUN   TestSayHi
--- PASS: TestSayHi (0.00s)
=== RUN   TestSayHiWithNoArgs
--- PASS: TestSayHiWithNoArgs (0.00s)
=== RUN   TestSayHiWithArgs
--- PASS: TestSayHiWithArgs (0.00s)
PASS
ok  	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet	(cached)

$ go test -v -run ^TestSayHi github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet
# github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet
# paketi içinde "TestSayHi" ile başlayanları çalıştır (regex match)

=== RUN   TestSayHi
--- PASS: TestSayHi (0.00s)
=== RUN   TestSayHiWithNoArgs
--- PASS: TestSayHiWithNoArgs (0.00s)
=== RUN   TestSayHiWithArgs
--- PASS: TestSayHiWithArgs (0.00s)
PASS
ok  	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet	(cached)
```

Test fonksiyonları `testing` paketinden duruma göre argüman alır. Örneğimizde
`t *testing.T` alıyoruz. `T`’nin yardımcı fonksiyonlarını kullanarak testimizi
geliştireceğiz.

Kabaca bu yaptığımız iş **assertion** yani bir şey iddia ediyoruz! Bize şu
fonksiyonu şu argümanlarla çağırdığımız zaman sonucun da bu olması lazım!

Yine adetlere uyarak `want` ve `got` identifer’larını kullanıyoruz. Bu zorunlu
değil ama hep söylediğimiz şey **conventions over configurations** :)

Beklentimizi ve fonksiyon çıktısını tanımladıktan sonra yapacak tek iş bu iki
değişkenin değerlerinin birbirine eşit olup olmadığı? yönünde.

Testi yazarken belli durumlarda;

- `t.Fatal`, `t.Fatalf` çalışma durur, buna **fail fast** denir, test ilerlemez!
- `t.Error`, `t.Errorf` hata olsa bile diğer testler çalışmaya devam eder...
- `b.Fatal`, `b.Fatalf` çalışma durur, buna **fail fast** denir, test ilerlemez!
- `b.Error`, `b.Errorf` hata olsa bile diğer testler çalışmaya devam eder...
- `t.Fail`, `b.Fail` testi fail olarak işaretler ama diğerlerini run etmeye devam eder...
- `t.FailNow`, `b.FailNow`: çalışma durur,  `runtime.Goexit()` çağrılır, defer edenler devam eder, test durmaz

Test çıktısını `json` formatında da alabiliriz:

```bash
$ go test -v -json github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet
{"Time":"2023-08-15T18:10:23.539298+03:00","Action":"start","Package":"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet"}
{"Time":"2023-08-15T18:10:23.539368+03:00","Action":"run","Package":"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet","Test":"TestSayHi"}
{"Time":"2023-08-15T18:10:23.539371+03:00","Action":"output","Package":"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet","Test":"TestSayHi","Output":"=== RUN   TestSayHi\n"}
{"Time":"2023-08-15T18:10:23.539379+03:00","Action":"output","Package":"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet","Test":"TestSayHi","Output":"--- PASS: TestSayHi (0.00s)\n"}
{"Time":"2023-08-15T18:10:23.539381+03:00","Action":"pass","Package":"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet","Test":"TestSayHi","Elapsed":0}
{"Time":"2023-08-15T18:10:23.539384+03:00","Action":"run","Package":"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet","Test":"TestSayHiWithNoArgs"}
{"Time":"2023-08-15T18:10:23.539385+03:00","Action":"output","Package":"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet","Test":"TestSayHiWithNoArgs","Output":"=== RUN   TestSayHiWithNoArgs\n"}
{"Time":"2023-08-15T18:10:23.539395+03:00","Action":"output","Package":"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet","Test":"TestSayHiWithNoArgs","Output":"--- PASS: TestSayHiWithNoArgs (0.00s)\n"}
{"Time":"2023-08-15T18:10:23.539398+03:00","Action":"pass","Package":"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet","Test":"TestSayHiWithNoArgs","Elapsed":0}
{"Time":"2023-08-15T18:10:23.539399+03:00","Action":"run","Package":"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet","Test":"TestSayHiWithArgs"}
{"Time":"2023-08-15T18:10:23.539401+03:00","Action":"output","Package":"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet","Test":"TestSayHiWithArgs","Output":"=== RUN   TestSayHiWithArgs\n"}
{"Time":"2023-08-15T18:10:23.539403+03:00","Action":"output","Package":"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet","Test":"TestSayHiWithArgs","Output":"--- PASS: TestSayHiWithArgs (0.00s)\n"}
{"Time":"2023-08-15T18:10:23.539417+03:00","Action":"pass","Package":"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet","Test":"TestSayHiWithArgs","Elapsed":0}
{"Time":"2023-08-15T18:10:23.539419+03:00","Action":"run","Package":"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet","Test":"ExampleSayHi"}
{"Time":"2023-08-15T18:10:23.539421+03:00","Action":"output","Package":"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet","Test":"ExampleSayHi","Output":"=== RUN   ExampleSayHi\n"}
{"Time":"2023-08-15T18:10:23.539422+03:00","Action":"output","Package":"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet","Test":"ExampleSayHi","Output":"--- PASS: ExampleSayHi (0.00s)\n"}
{"Time":"2023-08-15T18:10:23.539428+03:00","Action":"pass","Package":"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet","Test":"ExampleSayHi","Elapsed":0}
{"Time":"2023-08-15T18:10:23.539429+03:00","Action":"run","Package":"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet","Test":"ExampleSayHi_withNoArg"}
{"Time":"2023-08-15T18:10:23.53943+03:00","Action":"output","Package":"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet","Test":"ExampleSayHi_withNoArg","Output":"=== RUN   ExampleSayHi_withNoArg\n"}
{"Time":"2023-08-15T18:10:23.539537+03:00","Action":"output","Package":"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet","Test":"ExampleSayHi_withNoArg","Output":"--- PASS: ExampleSayHi_withNoArg (0.00s)\n"}
{"Time":"2023-08-15T18:10:23.539538+03:00","Action":"pass","Package":"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet","Test":"ExampleSayHi_withNoArg","Elapsed":0}
{"Time":"2023-08-15T18:10:23.53954+03:00","Action":"run","Package":"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet","Test":"ExampleSayHi_withArgs"}
{"Time":"2023-08-15T18:10:23.539541+03:00","Action":"output","Package":"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet","Test":"ExampleSayHi_withArgs","Output":"=== RUN   ExampleSayHi_withArgs\n"}
{"Time":"2023-08-15T18:10:23.539543+03:00","Action":"output","Package":"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet","Test":"ExampleSayHi_withArgs","Output":"--- PASS: ExampleSayHi_withArgs (0.00s)\n"}
{"Time":"2023-08-15T18:10:23.539545+03:00","Action":"pass","Package":"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet","Test":"ExampleSayHi_withArgs","Elapsed":0}
{"Time":"2023-08-15T18:10:23.539547+03:00","Action":"output","Package":"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet","Output":"PASS\n"}
{"Time":"2023-08-15T18:10:23.539548+03:00","Action":"output","Package":"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet","Output":"ok  \tgithub.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet\t(cached)\n"}
{"Time":"2023-08-15T18:10:23.539555+03:00","Action":"pass","Package":"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet","Elapsed":0}
```

Keza diğer bir komut satırı argümanı da `-list`. `-list <regex>` ile test
fonksiyonlarını çalıştırır ve listeler;

```bash
$ go test -v -list Say ./...
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/02/01-init	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/04/05-struct-custom-tag	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/04/05-struct-field-access	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/04/05-struct-field-access/person	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/04/05-struct-field-access-getter	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/04/05-struct-field-access-getter/person	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/04/05-struct-field-alignment	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/04/05-struct-validate	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/11/generics-constraints	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/11/generics-custom-types	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/11/generics-function-calls	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/11/generics-function-calls-and-types	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/11/generics-function-calls-and-types2	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/11/generics-functions	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/11/generics-functions-interface	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/11/generics-in-maps	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/11/generics-in-structs	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/12/reflect-clearvalue	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/12/reflect-typecheck	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/13/json-custom-decode	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/13/json-generic-interface	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/13/json-marshal-custom-time	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/13/json-streaming	[no test files]
?   	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet	[no test files]
ok  	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-datarace/kvstore	0.680s
TestSayHi
TestSayHiWithNoArgs
TestSayHiWithArgs
ExampleSayHi
ExampleSayHi_withNoArg
ExampleSayHi_withArgs
ok  	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-greet/greet	0.550s
```

`go test`’in aldığı tüm parametrelerin detaylarını:

```bash
$ go help testflag
```

ile görebilirsiniz.

---

## Data Race Detection

**Race Condition**, concurrent (eş zamanlı) programla ile ilgili sık yaşanan
bir durum. Basitçe şöyle açıklayalım; bir web uygulaması var, veritabanına
bağlı. Veritabanı tablosunda `count` adında bir alan ver, sayısal değer
tutuyor. Kullanıcı `example.com` sayfasını çağırdığı zaman uygulama tablodaki
`count` alanının değerini `+1` arttırıyor.

- İlk kullanıcı geldi, değer 0 -> 1 oldu
- İkinci kullanıcı geldi, değer 1 -> 2 oldu

peki 3. ve 4. kullanıcı aynı anda gelince ne olacak?

    3. Kullanıcı    | 4. Kullanıcı
          |               |
          ↓               ↓
        2 -> 3          2 -> 3

Aynı anda geldiler ve ikisi için de önceki değer `2` idi, 3. kullanıcı değeri
`3` yaptı ama 4. için değer halen `2` idi. Bu durumda istekler birbirleriyle 
yarış yapmış oluyorlar ve bu durum yarış durumu yani **race condition** oluyor.

Doğru olan durum ise şöyle olmalıydı;

- 3. geldi, bir tür kitleme (lock) olmalıydı.
- 4. geldi, lock olduğu için beklemeliydi.
- 3. işini bitirdi kiliti açtı (unlock)
- 4. lock edip doğru değeri okuyup, artırıp unlock etmeliydi

Benzer senaryo bizim için **go routine** kullandığımız zaman yaşanacak. İşte
bu tür kaçakları test esnasında yakalamak için `-race` parametresini
kullanıyoruz. Şimdi aynısı go ile yapalım:

[Örnek](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/14/test-datarace)

```go
package kvstore

import "errors"

var errKeyNotFound = errors.New("key not found")

// Store is key-value store!
type Store struct {
	db map[string]string
}

// Set new key to store.
func (s *Store) Set(k, v string) error {
	s.db[k] = v
	return nil
}

// Get accepts key, returns value and error.
func (s *Store) Get(k string) (string, error) {
	v, ok := s.db[k]
	if !ok {
		return "", errKeyNotFound
	}
	return v, nil
}

// New returns new Store instance.
func New(db map[string]string) Store {
	return Store{db: db}
}
```

Test ise;

```go
package kvstore_test

import (
	"testing"

	"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-datarace/kvstore"
)

func TestDataRace(t *testing.T) {
	st := make(map[string]string)
	done := make(chan struct{})

	s := kvstore.New(st)
	_ = s.Set("foo", "bar")

	go func() {
		_ = s.Set("foo", "data race...")
		done <- struct{}{}
	}()

	want := "bar"
	got, _ := s.Get("foo") // always returns "bar"
	<-done                 // after line 19, blocking ends... map changes but doesn't affect got variable!

	// fmt.Println(s.Get("foo")) data race... <nil>

	if got != want {
		t.Errorf("want: %v; got: %v", want, got)
	}
}
```

Önce testi düz çalıştıralım:

```bash
$ go test -v github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-datarace/kvstore
=== RUN   TestDataRace
--- PASS: TestDataRace (0.00s)
PASS
ok  	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-datarace/kvstore	0.608s
```

Test `pass` etti, yani testi geçtik, halbuki:

```bash
$ go test -v -race github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-datarace/kvstore
=== RUN   TestDataRace
==================
WARNING: DATA RACE
Write at 0x00c00011e450 by goroutine 7:
  runtime.mapaccess2_faststr()
      /opt/homebrew/Cellar/go/1.21.0/libexec/src/runtime/map_faststr.go:108 +0x42c
  github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-datarace/kvstore.(*Store).Set()
      /Users/vigo/Development/VBYazilim/maoyyk2023-golang-101-kursu/src/14/test-datarace/kvstore/kvstore.go:14 +0x5c
  github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-datarace/kvstore_test.TestDataRace.func1()
      /Users/vigo/Development/VBYazilim/maoyyk2023-golang-101-kursu/src/14/test-datarace/kvstore/kvstore_test.go:17 +0x38

Previous read at 0x00c00011e450 by goroutine 6:
  runtime.mapaccess1_faststr()
      /opt/homebrew/Cellar/go/1.21.0/libexec/src/runtime/map_faststr.go:13 +0x40c
  github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-datarace/kvstore.(*Store).Get()
      /Users/vigo/Development/VBYazilim/maoyyk2023-golang-101-kursu/src/14/test-datarace/kvstore/kvstore.go:20 +0x1b8
  github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-datarace/kvstore_test.TestDataRace()
      /Users/vigo/Development/VBYazilim/maoyyk2023-golang-101-kursu/src/14/test-datarace/kvstore/kvstore_test.go:22 +0x1dc
  testing.tRunner()
      /opt/homebrew/Cellar/go/1.21.0/libexec/src/testing/testing.go:1595 +0x194
  testing.(*T).Run.func1()
      /opt/homebrew/Cellar/go/1.21.0/libexec/src/testing/testing.go:1648 +0x40

Goroutine 7 (running) created at:
  github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-datarace/kvstore_test.TestDataRace()
      /Users/vigo/Development/VBYazilim/maoyyk2023-golang-101-kursu/src/14/test-datarace/kvstore/kvstore_test.go:16 +0x190
  testing.tRunner()
      /opt/homebrew/Cellar/go/1.21.0/libexec/src/testing/testing.go:1595 +0x194
  testing.(*T).Run.func1()
      /opt/homebrew/Cellar/go/1.21.0/libexec/src/testing/testing.go:1648 +0x40

Goroutine 6 (running) created at:
  testing.(*T).Run()
      /opt/homebrew/Cellar/go/1.21.0/libexec/src/testing/testing.go:1648 +0x5d8
  testing.runTests.func1()
      /opt/homebrew/Cellar/go/1.21.0/libexec/src/testing/testing.go:2054 +0x80
  testing.tRunner()
      /opt/homebrew/Cellar/go/1.21.0/libexec/src/testing/testing.go:1595 +0x194
  testing.runTests()
      /opt/homebrew/Cellar/go/1.21.0/libexec/src/testing/testing.go:2052 +0x6d8
  testing.(*M).Run()
      /opt/homebrew/Cellar/go/1.21.0/libexec/src/testing/testing.go:1925 +0x904
  main.main()
      _testmain.go:49 +0x294
==================
==================
WARNING: DATA RACE
Write at 0x00c000166088 by goroutine 7:
  github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-datarace/kvstore.(*Store).Set()
      /Users/vigo/Development/VBYazilim/maoyyk2023-golang-101-kursu/src/14/test-datarace/kvstore/kvstore.go:14 +0x68
  github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-datarace/kvstore_test.TestDataRace.func1()
      /Users/vigo/Development/VBYazilim/maoyyk2023-golang-101-kursu/src/14/test-datarace/kvstore/kvstore_test.go:17 +0x38

Previous read at 0x00c000166088 by goroutine 6:
  github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-datarace/kvstore.(*Store).Get()
      /Users/vigo/Development/VBYazilim/maoyyk2023-golang-101-kursu/src/14/test-datarace/kvstore/kvstore.go:20 +0x1c4
  github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-datarace/kvstore_test.TestDataRace()
      /Users/vigo/Development/VBYazilim/maoyyk2023-golang-101-kursu/src/14/test-datarace/kvstore/kvstore_test.go:22 +0x1dc
  testing.tRunner()
      /opt/homebrew/Cellar/go/1.21.0/libexec/src/testing/testing.go:1595 +0x194
  testing.(*T).Run.func1()
      /opt/homebrew/Cellar/go/1.21.0/libexec/src/testing/testing.go:1648 +0x40

Goroutine 7 (running) created at:
  github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-datarace/kvstore_test.TestDataRace()
      /Users/vigo/Development/VBYazilim/maoyyk2023-golang-101-kursu/src/14/test-datarace/kvstore/kvstore_test.go:16 +0x190
  testing.tRunner()
      /opt/homebrew/Cellar/go/1.21.0/libexec/src/testing/testing.go:1595 +0x194
  testing.(*T).Run.func1()
      /opt/homebrew/Cellar/go/1.21.0/libexec/src/testing/testing.go:1648 +0x40

Goroutine 6 (running) created at:
  testing.(*T).Run()
      /opt/homebrew/Cellar/go/1.21.0/libexec/src/testing/testing.go:1648 +0x5d8
  testing.runTests.func1()
      /opt/homebrew/Cellar/go/1.21.0/libexec/src/testing/testing.go:2054 +0x80
  testing.tRunner()
      /opt/homebrew/Cellar/go/1.21.0/libexec/src/testing/testing.go:1595 +0x194
  testing.runTests()
      /opt/homebrew/Cellar/go/1.21.0/libexec/src/testing/testing.go:2052 +0x6d8
  testing.(*M).Run()
      /opt/homebrew/Cellar/go/1.21.0/libexec/src/testing/testing.go:1925 +0x904
  main.main()
      _testmain.go:49 +0x294
==================
    testing.go:1465: race detected during execution of test
--- FAIL: TestDataRace (0.00s)
=== NAME  
    testing.go:1465: race detected during execution of test
FAIL
FAIL	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-datarace/kvstore	0.526s
FAIL
```

---

## Kaynaklar

- https://go.dev/blog/examples
- https://elliotchance.medium.com/godoc-tips-tricks-cda6571549b
- https://go.dev/blog/race-detector
- https://en.wikipedia.org/wiki/Race_condition
