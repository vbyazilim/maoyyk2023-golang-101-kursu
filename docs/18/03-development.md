# Bölüm 18/03: In-Memory Key-Value Store

## Development

Diğer pek çok **framework**’ün belli bir yoğurt yeme tarzı olur.
Python’cuların çok sevdiği **Django**, Ruby’cilerin **Ruby on Rails**,
PHP’cilerin **Laravel**, Node’cuların **Express** gibi sizin adınıza pek çok
sorunu çözdükleri harika **framework**’leri var. İçinde bir çok fonksiyon ve
mantık barındıran devasa kod yığınları.

Go’da bu tür uçtan-uca her derde deva bir **framework** ne yazık ki yok. Eğer
http server ihtiyacımız varsa, sadece http katmanını çözen, sadece veritabanı
katmanını çözen, aynı lego parçaları gibi ayrı-ayrı kütüphaneler bulunur.
Bunları birbirine bağlamak da geliştirciye kalır :)

Dolayısıyla, go için en fazla **best practice**’ler (en iyi pratikler) söz
konusudur. GitHub’ta pek çok "go application / project structre" gibi repo’lar
bulmak mümkün.

Genelde ben, go’nun kaynak kodunu referans alıyorum, acaba go’yu icad edenler
ne tür bir yaklaşım içine girmişler, neler uygulamışlar hep bunlara bakıyorum.

Bu bağlamda proje yapımız:

    .
    ├── Dockerfile
    ├── README.md
    ├── cmd
    │   └── server
    ├── go.mod
    └── src
        ├── apiserver
        │   ├── apiserver.go
        │   └── middlewares.go
        ├── internal
        │   ├── kverror
        │   ├── service
        │   │   └── kvstoreservice
        │   ├── storage
        │   │   └── memory
        │   └── transport
        │       └── http
        └── releaseinfo

şeklinde. `src/internal/` dizini altında;

1. `storage/`
1. `service/`
1. `transport/http`

paket tanımlarımızı yapıyoruz. Neden `internal/` kullanıyoruz, eğer bu projeyi
herhangi bir kullanıcı `go get` ile sanki bir kütüphaneymiş gibi projesine
eklerse, `internal/` altındaki hiçbir pakete erişimi olamayacak! Şu an sadece;

- `src/apiserver`
- `src/releaseinfo`

Paketleri **exportable** yani `import` edilebilir durumda. Esas uygulamanın
çalışacağı yer `cmd/server/` altındaki `main.go` dosyası olacak. Server
ile ilgili tanımlamaları `apiserver/` altında yapacağız. Özel bir `error`
tipimiz var: `kverror`.

`storage/memory/` kullanacağımız in-memory storage davranışı ve işi yapan
fonksiyonlar burada olacak. Yarın "artık veritabanı kullananım" dersek;
`storage/postgresql/` altına gereken davranışları ve fonksiyonları
tanımlayabiliriz.

Keza, aynı şekilde, yarın sadece `http` protokolü yerine `rpc` ya da `grpc`
sunmak istersek: `transport/rpc/` ya da `transport/grpc/` gibi ilerleyebiliriz.

## go mod init

Şimdi geliştirme yapacağımız dizine gidip projeyi başlatalım:

```bash
$ cd /path/to/development/
$ mkdir kvstore
$ cd kvstore/
$ git init
$ git commit --allow-empty -m '[root] add initial commit'
```

Şimdi https://gitignore.io sitesinden projemiz için gereken `.gitignore`
dosyasını alıyoruz ve projenin ana dizininde (root) `touch .gitignore` yaparak
içine paste ediyoruz;

```bash
$ git add .
$ git commit -m 'add gitignore file'
```

Şimdi go modülümüzü oluşturalım;

```bash
$ go mod init github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore
$ git add .
$ git ci -m 'add go.mod file'
```

İlk olarak storage katmanından başlıyoruz;

```bash
$ mkdir -p src/internal/storage/memory/kvstorage
$ tree .
.
├── go.mod
└── src
    └── internal
        └── storage
            └── memory
                └── kvstorage   <--- paket adı
```

Şimdi storage ile ilgili tanımları yapmak için;

```bash
$ touch src/internal/storage/memory/kvstorage/base.go
```

Şimdi projeyi kod editöründe açalım ve `base.go` dosyasına şunu yazalım ve kaydedelim:

```go
package kvstorage
```

Go koduna başladık, hemen [linter konfigürasyon](../../.golangci.yml)
dosyamızı root dizine atalım:

---

## Kaynaklar

- https://github.com/avelino/awesome-go#project-layout

---
