# Bölüm 17/02: Golang Paketi Geliştirmek

Şimdi örnek bir `go.mod` dosyasına bakalım;

```go
module github.com/my/library
              |
              └---- bu modülün host edildiği yer, import path
          
go 1.16       <---- modül develop edilirken kullanılan golang’in versiyonu

require (     <---- modülün bağımlılıklarının tanımının bağladığı yer
    github.com/dep/one v1.0.0 
    github.com/dep/two/v2 v2.3.0 
    github.com/vigo/stringutils-demo v0.1.1 // indirect
                                |
                                └---- bu paket "go get github.com/vigo/stringutils-demo" ile kurulmuş ama
                                      henüz kod içinde kullanılmamış, bu bakımdan "indirect"
                                      bağımlılık olarak verilmiş ama hiçbir yerde kullanılmamış.
    github.com/dep/other v0.0.0-20180523231146-b3f5c0f6e5f1
                         |
                         └---- Bu aslında "pseudo-version" yani esas bir tag’e değil, commit’e bakıyor.
                               Tarih bilgisi      - Commit Hash
                               2018-05-23-23:11:46-b3f5c0f6e5f1

    github.com/dep/legacy v2.0.0+incompatible
                         |
                         └---- "incompatible" çünkü bu paket henüz "go mod" yapısına geçmemiş.
)

exclude github.com/dep/legacy v1.9.2
  └---- Belirli bir modül sürümünün kullanılmasını engelle
replace github.com/dep/one => github.com/fork/one
  └---- github.com/dep/one paketini github.com/fork/one ile değiştir

```

Keza `go.mod` üzerinde otomatik değişiklikler yapmak için aşağıdaki komutlara
bakalım:

```bash
$ go get -d github.com/path/to/module       # build ya da install etmeden sadece go.mod’u güncelle
                                            # ve paketi build etmek için gereken kodu çek
                                            # (add/upgrade dependency)

$ go get -d github.com/dep/two/v2@v2.1.0    # v2.1.0 için build ya da install etmeden sadece go.mod’u güncelle
                                            # ve paketi build etmek için gereken kodu çek
                                            # (use specific version)

$ go get -d github.com/dep/commit@branch    # verilen branch için build ya da install etmeden sadece go.mod’u güncelle
                                            # ve paketi build etmek için gereken kodu çek
                                            # (use specific branch)

$ go get -d -u ./...                        # tüm paketler için build ya da install etmeden sadece go.mod’u güncelle
                                            # ve paketi build etmek için gereken kodu çek
                                            # (upgrade all!)

$ go get -d github.com/dep/legacy@none      # @none ile bu pakete bağımlılığı siliyoruz, go.mod’dan da uçuyor
                                            # sonrasında `go mod tidy` lazım
                                            # (remove dependency)

$ go mod tidy                               # go.mod ve go.sum’ı düzenler, temizler, eksikleri fazlalıkları düzenler ayarlar.

$ go mod download                           # bağımlığı module cache’e atar.
                                            # tüm indirilenler, bilgisayarınızda `GOMODCACHE` neresiyse oraya kaydolur...

$ ls -al $(go env GOMODCACHE)/github.com/vigo
total 0
drwxr-xr-x  5 vigo staff  160 Nov 13 22:53 .
drwxr-xr-x 74 vigo staff 2.4K Nov 12 19:35 ..
dr-xr-xr-x 14 vigo staff  448 Nov 10 20:09 lsvirtualenvs@v0.1.0
dr-xr-xr-x 12 vigo staff  384 Nov 13 22:47 stringutils-demo@v0.0.0-20211113192943-449a20582367
dr-xr-xr-x 12 vigo staff  384 Nov 13 22:53 stringutils-demo@v0.1.1

$ go mod init github.com/path/to/module     # yeni modül başlat

$ go mod why -m github.com/path/to/module   # verilen modül neden bir bağımlılık?

$ cd /path/to/github.com/vigo/stringutils-demo/
$ go mod why -m github.com/vigo/stringutils-demo
# github.com/vigo/stringutils-demo
github.com/vigo/golang102-custom-package-demo
github.com/vigo/stringutils-demo

$ go install github.com/path/to/bin@latest  # ilgili paketi hem çek hem de varsa binary’sini build edip
                                            # GOPATH’in altındaki bin/’e at
$ ls -al $(go env GOPATH)/bin

# go mod edit -replace SOURCE=TARGET
$ go mod edit -replace github.com/pselle/bar=/Users/pselle/Projects/bar
```

---

## Kaynaklar

- https://golang.org/doc/tutorial/create-module
- https://golang.org/doc/modules/managing-dependencies
- https://thewebivore.com/using-replace-in-go-mod-to-point-to-your-local-module/
