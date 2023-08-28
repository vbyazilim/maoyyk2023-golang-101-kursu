# Bölüm 02/02: Golang Uygulamasına Genel Bakış

## Executable Nedir?

İngilizce **execute** edilebilen yani yürütülebilen, çalıştırılabilen
uygulamaya **executable** denir. Gündelik hayatta komut satırından
çalıştırdığınız komutların büyük çoğunluğu (örneğin `ls`, `cat` gibi...) bu
kategoriden uygulamalardır.

Go ile bu şekide, üzerinde koştuğu işletim sistemine göre çalışan uygulamalar
geliştirebiliriz. Yazdığımız go uygulamasını işletim sisteminden bağımsız
(Linux, macOS ya da Windows) derleyebilir ve çalıştırabiliriz.

Kodu çalıştırmak için bir kaç farklı yöntem var;

### `go run /path/to/main.go`

Önce dosyayı derler, ürettiği (build ettiği) binary dosyayı işletim sisteminde
gizli bir yere atar, genelde `/tmp/` altında, sonra oradan binary’i çağırarak
çalıştırır.

### `go run .`

`.` shell ortamında **current working directory** yani o an içinde buluduğumuz
dizin anlamındadır, bulunduğumuz yerdeki modül yapısına göre uygun `main.go`
dosyasını bulur ve çalıştırır. Genelde içinde `go.mod` olan bir proje dizininde
olmamız gerekir aksi halde;

```bash
$ go run .
go: go.mod file not found in current directory or any parent directory; see 'go help modules'
```

hata alırız.

### `go build`

Bulunduğumuz dizin içindeki `go.mod` yapısına göre ilgili `main.go` dosyasını
bulur ve derler. Ürettiği binary’i yine `.` yani **current working directory**
altına atar; sonra elle biz çalıştırırız:

```bash
$ cd /tmp/
$ mkdir demo
$ cd demo
$ go mod init demo
$ cat << EOF > main.go
package main

import "fmt"

func main() {
    fmt.Println("merhaba dünya!")
}
EOF

$ go build
$ ./demo
merhaba dünya!
```

Adı `demo` olan binary / executable dosya üretildi. Peki bu dosyanın adı neden
`demo` oldu? Çünkü modülümüzün adı `demo`. Başka bir isim vermek için;

- Ya modülün adını değişeceğiz
- Ya da derlerken `-o NAME` ile `go build -o projem`

şeklinde değiştirebiliriz.

### `go install`

Bu komut sayesinde, verilen paketi uzaktan ya da yerelden önce indirip, derleyip,
go kurulumundaki çalıştırılabilir dosyaların olduğu yere otomatik olarak atıp,
işletim sistemi seviyesinde (`$PATH`) çalıştırabiliriz;

```bash
$ go env GOPATH
/Users/vigo/.local/go

$ ls -al "$(go env GOPATH)/bin"
$ echo $PATH

$ # go install PACKAGE@TAG
$ go install github.com/vigo/statoo/v2@latest
#            ^                         ^
#            +--- paket adı            +--- hangi revizyon? en son tag’i indir

# tüm release’ler: https://github.com/vigo/statoo/releases

# eski versiyon için:
$ go list -m -versions github.com/vigo/statoo@latest
github.com/vigo/statoo v0.1.0 v0.1.1 v0.1.2 v0.1.3 v0.2.0 v0.2.1 v0.2.2 v0.2.3 v1.0.0 v1.0.1 v1.1.0 v1.1.1 v1.1.2 v1.1.3 v1.2.0 v1.2.1 v1.2.2 v1.2.3 v1.3.0 v1.3.1 v1.4.0

$ go list -m -versions github.com/vigo/statoo/v2@latest
github.com/vigo/statoo/v2 v2.0.3

# go install github.com/vigo/statoo@latest  # v1 familyasındaki son sürüm
# go install github.com/vigo/statoo@v1.0.0  # v1.0.0
```

---

## Library Nedir?

İçinde `main.go` olma zorunluğu bulunmayan, yardımcı fonksiyon
koleksiyonlarının bulunduğu üçüncü parti kodları tarif ederken kullanırız.
Örneğin, "acaba bu servisin bir go-sdk’i ya da client’ı var mı?" dediğimizde
aslında ilgili servisin api’ını kullanmak için geliştirilmiş bir
paketin/kütüphanenin olup olmadığını sormuş oluruz.

https://github.com/vigo/stringutils-demo

Eğer projemiz içinde `stringutils-demo` paketini kullanmak istersek, bağımlılık
olarak projemiz içine almak için; projemizin ana dizininde (yani `go.mod`)
dosyasının olduğu yerde;

```bash
$ go get github.com/vigo/stringutils-demo
```

ile eklememiz yeterlidir. Paket / kütüphane konularını ve kullanımlarını
daha ileriki bölümlerde detaylı şekilde inceleyeceğiz.
