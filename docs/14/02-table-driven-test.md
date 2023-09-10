# Bölüm 14/02: Test

## Table Driven Tests

Test etmek istediğimiz fonksiyon için farklı farklı 10 tane ayrı test yazmak
yerine, girdi ve çıktıyı bir `slice` içinde okuyarak fonksiyona parametre
geçme ve sonucu test etme işlemine denir.

`SayHi()` fonksiyonumuz için;

```go
type test struct {
	input    []string
	want     string
}

tests := []test{
	{input: []string{"vigo"}, want: "hi vigo!"},
	{input: []string{"vigo", "turbo"}, want: "hi vigo!\nhi turbo!"},
	{input: []string{}, want: "hi everybody!"},
}
```

şeklinde tablomuz olsun. Testleri çalıştırmak için;

```go
// tests'de kaç tane örnek varsa...
for _, tc := range tests {
	got := SayHi(tc.input...)

	if got != tc.want {
		t.Errorf("want: %v; got: %v", tc.want, got)
	}
}
```

yapabiliriz. Eğer bir noktada patlarsa, tam olarak nerede ya da ne yaparken
sorun çıktığını anlamak için `test` struct’ımıza ilave bir alan ekliyoruz:

```go
type test struct {
	testName string
	input    []string
	want     string
}

tests := []test{
	{testName: "run with single arg", input: []string{"vigo"}, want: "hi vigo!\n"},
	{testName: "run with multiple args", input: []string{"vigo", "turbo"}, want: "hi vigo!\nhi turbo!\n"},
	{testName: "run with no arg", input: []string{}, want: "hi everybody"},
}

for _, tc := range tests {
	got := SayHi(tc.input...)

	if got != tc.want {
		t.Errorf("[%s]: want: %v; got: %v", tc.testName, tc.want, got)
	}
}
```

Testi çalıştıralım:

[Örnek](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/14/test-table-driven)

```bash
$ go test -v github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-table-driven/greet
```

---

## Sub Tests

Nasıl ki testleri bir şekilde isimlendirerek kategorize ettiysek, go bize bunu
daha rahat yapma ve her parçacığa erişebilecek isim verme imkanı sağlıyor.

Bu sayede istediğimiz an istediğimiz alt parçacığı çalıştırabiliyoruz. Buna
da **sub test** deniyor:

[Örnek](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/14/test-table-driven-sub-tests)

```go
package greet_test

import (
	"testing"

	"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-table-driven-sub-tests/greet"
)

func TestSayHi(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  string
	}{
		"run with single arg":    {input: []string{"vigo"}, want: "hi vigo!"},
		"run with multiple args": {input: []string{"vigo", "turbo"}, want: "hi vigo!\nhi turbo!"},
		"run with no arg":        {input: []string{}, want: "hi everybody!"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := greet.SayHi(tc.input...)

			if got != tc.want {
				t.Errorf("want: %v; got: %v", tc.want, got)
			}
		})
	}
}
```

Testi çalıştıralım:

```bash
$ go test -v github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-table-driven-sub-tests/greet
=== RUN   TestSayHi
=== RUN   TestSayHi/run_with_single_arg
=== RUN   TestSayHi/run_with_multiple_args
=== RUN   TestSayHi/run_with_no_arg
--- PASS: TestSayHi (0.00s)
    --- PASS: TestSayHi/run_with_single_arg (0.00s)
    --- PASS: TestSayHi/run_with_multiple_args (0.00s)
    --- PASS: TestSayHi/run_with_no_arg (0.00s)
PASS
ok  	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-table-driven-sub-tests/greet	0.371s
```

Yine `tests` içinde yineleme (iteration) yapıyoruz fakat bu kez `t.Run()`
ile testi çalıştırıyoruz.

---

## SetUp ve TearDown

Testleri çalıştırmadan önce ve çalıştırdıktan sonra bazı işlemler yapmak
isteyebiliriz. Test başlamadan önce test için veritabanı oluşturup test
bitiminde bunu silebiliriz. Bu gibi durumlar `SetUp` ve `TearDown` anları
olarak ifade edilir.

Go’da iki yöntem yaygındır. `func TestMain(m *testing.M)`ile ya da **sub
test** ile.

`TestMain`:

Teste başlamadan önce bazı environment variable’ları set edelim, test
bitiminde de silelim.

```go
func TestMain(m *testing.M) {
	fmt.Println("do setup operations...")
	os.Setenv("CUSTOM_HOST", "localhost")
	os.Setenv("CUSTOM_PORT", "9000")

	result := m.Run()

	fmt.Println("do teardown operations...")
	os.Unsetenv("CUSTOM_HOST")
	os.Unsetenv("CUSTOM_PORT")

	os.Exit(result)
}
```

Eğer sub test içinden yapmak istersek; `t.Run()` öncesi ve sonrasını
kullanırız:

[Örnek](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/14/test-setup-teardown)

```go
package greet_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-setup-teardown/greet"
)

func TestMain(m *testing.M) {
	fmt.Println("do setup operations...")
	_ = os.Setenv("CUSTOM_HOST", "localhost")
	_ = os.Setenv("CUSTOM_PORT", "9000")

	result := m.Run()

	fmt.Println("do teardown operations...")
	_ = os.Unsetenv("CUSTOM_HOST")
	_ = os.Unsetenv("CUSTOM_PORT")

	os.Exit(result)
}

func TestSayHi(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  string
	}{
		"run with single arg":    {input: []string{"vigo"}, want: "hi vigo!"},
		"run with multiple args": {input: []string{"vigo", "turbo"}, want: "hi vigo!\nhi turbo!"},
		"run with no arg":        {input: []string{}, want: "hi everybody!"},
	}

	for name, tc := range tests {
		// <setup code>
		fmt.Println("setup code from sub test initiated!")
		t.Run(name, func(t *testing.T) {
			if val, ok := os.LookupEnv("CUSTOM_PORT"); ok && val == "9000" {
				fmt.Println("using port 9000")
			}
			got := greet.SayHi(tc.input...)

			if got != tc.want {
				t.Errorf("want: %v; got: %v", tc.want, got)
			}
		})
		// <tear-down code>
		fmt.Println("teardown code from sub test initiated!")
	}
}
```

Çalıştıralım:

```bash
$ go test -v github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-setup-teardown/greet
do setup operations...
=== RUN   TestSayHi
setup code from sub test initiated!
=== RUN   TestSayHi/run_with_single_arg
using port 9000
teardown code from sub test initiated!
setup code from sub test initiated!
=== RUN   TestSayHi/run_with_multiple_args
using port 9000
teardown code from sub test initiated!
setup code from sub test initiated!
=== RUN   TestSayHi/run_with_no_arg
using port 9000
teardown code from sub test initiated!
--- PASS: TestSayHi (0.00s)
    --- PASS: TestSayHi/run_with_single_arg (0.00s)
    --- PASS: TestSayHi/run_with_multiple_args (0.00s)
    --- PASS: TestSayHi/run_with_no_arg (0.00s)
PASS
do teardown operations...
ok  	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-setup-teardown/greet	0.292s
```

---

## Paralel Test

Her **test** dediğimiz şey aslında bir test fonksiyonu ile ilişkilendirilmiş
durumda. Yani `TestSayHi` bir test fonksiyonu. Eğer hangi test fonksiyonu
`t.Paralel()` metotunu çağırırsa, o test artık paralelde çalışan bir test
haline dönüşür.

Aynı sub test örneğimizi paralel çalışan test haline getirelim:

[Örnek](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/14/test-parallel)

```go
package greet_test

import (
	"testing"

	"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-parallel/greet"
)

// TestSayHi will not complete until all parallel tests started by Run have completed.
// As a result, no other parallel tests can run in parallel to these parallel tests.
func TestSayHi(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  string
	}{
		"run with single arg":    {input: []string{"vigo"}, want: "hi vigo!"},
		"run with multiple args": {input: []string{"vigo", "turbo"}, want: "hi vigo!\nhi turbo!"},
		"run with no arg":        {input: []string{}, want: "hi everybody!"},
	}

	for name, tc := range tests {
		tc := tc // capture range variable to ensure that
		// tc gets bound to the correct instance.

		t.Run(name, func(t *testing.T) {
			t.Parallel() // run in parallel

			got := greet.SayHi(tc.input...)
			if got != tc.want {
				t.Errorf("want: %v; got: %v", tc.want, got)
			}
		})
	}
}
```

Çalıştıralım:

```bash
$ go test -v github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-parallel/greet
=== RUN   TestSayHi
=== RUN   TestSayHi/run_with_no_arg
=== PAUSE TestSayHi/run_with_no_arg
=== RUN   TestSayHi/run_with_single_arg
=== PAUSE TestSayHi/run_with_single_arg
=== RUN   TestSayHi/run_with_multiple_args
=== PAUSE TestSayHi/run_with_multiple_args
=== CONT  TestSayHi/run_with_no_arg
=== CONT  TestSayHi/run_with_multiple_args
=== CONT  TestSayHi/run_with_single_arg
--- PASS: TestSayHi (0.00s)
    --- PASS: TestSayHi/run_with_no_arg (0.00s)
    --- PASS: TestSayHi/run_with_multiple_args (0.00s)
    --- PASS: TestSayHi/run_with_single_arg (0.00s)
PASS
ok  	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-parallel/greet	0.282s
```

Paralel bir test hiçbir zaman sıralı (sequential) bir testle aynı anda
(concurrent) çalışmaz ve testin çalışması, üst testin çağıran test
fonksiyonu `return` edene kadar kadar askıya alınır.

`CONT`: Continue anlamında. `-parallel` komut satırı argümanı ile çalışabilecek
maksimum paralel test sayısını da belirtebiliyoruz:

Default paralel test sayısı `runtime.GOMAXPROCS`’la set edilmiş sayıdır. Eğer
aksi belirtilmemişse, `GOMAXPROCS`’un default değeri `runtime.NumCPU()`
değeridir. Kullandığım bilgisayarda `runtime.NumCPU()` sonucu `10`. Eğer
`-parallel` ile hiçbir şey belirlenmediyse benim kullandığım bilgisayarda `10`
paralel test çalışacak.

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU())
}
```

Çalıştıralım:

```bash
$ go test -v -parallel 4 github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-parallel/greet
```

---

## Kaynaklar

- https://dave.cheney.net/2019/05/07/prefer-table-driven-tests
- https://go.dev/blog/subtests
- https://www.gopherguides.com/articles/table-driven-testing-in-parallel
- https://splice.com/blog/lesser-known-features-go-test/
- https://eleni.blog/2019/05/11/parallel-test-execution-in-go/
