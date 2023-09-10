# Bölüm 14/03: Test

## Code Coverage Nedir?

Bir proje ya da uygulama ya da paket geliştirdik. Peki yazdığımız kodun ne
kadarını test ettik? Code coverage, yazılan kodun **% kaçının** test
edildiğini ölçme işlemidir.

[Örnek](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/14/test-code-coverage)

```bash
$ go test -v -cover github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-code-coverage/greet
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
coverage: 100.0% of statements
ok  	github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-code-coverage/greet	0.273s	coverage: 100.0% of statements
```

Sonuç: **coverage: 100.0% of statements** mükemmel :)

Şimdi bu coverage’dan profil üretip web arayüzünde görüntüleyelim:

```bash
$ go test -v -cover -coverprofile src/14/test-code-coverage/coverage.out github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-code-coverage/greet
```

Çıkan `coverage.out` dosyası:

```bah
$ cat src/14/test-code-coverage/coverage.out
mode: set
github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-code-coverage/greet/greet.go:6.36,7.21 1 1
github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-code-coverage/greet/greet.go:7.21,9.3 1 1
github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-code-coverage/greet/greet.go:10.2,11.29 2 1
github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-code-coverage/greet/greet.go:11.29,13.3 1 1
github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-code-coverage/greet/greet.go:15.2,15.32 1 1
```

Şimdi web arayüzünden bakalım:

```bash
$ go tool cover -html src/14/test-code-coverage/coverage.out

$ go tool cover -html=src/14/test-code-coverage/coverage.out -o /tmp/test.html
```


---

