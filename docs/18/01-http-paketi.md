# Bölüm 18/01: In-Memory Key-Value Store

Basit bir rest-api geliştireceğiz. Küçük ve basit bir [REDIS][01] klonu. Klon
ama çok primitif bir klon. Key/Value çiftlerini hafızada tutan, listeleme,
ekleme, silme, okuma ve güncelleme yani `CRUDL` (Create, Read, Update, Delete,
List) operasyonları yapabileceğimiz bir servis. Diğer bir amacımız da hiçbir
ek paket kullanmadan, go ile gelen paketleri kullanarak bu servisi geliştirmek.

Nelere ihtiyacımız var;

- HTTP Server
- Storage (hafızada tutacağımız map)

## HTTP Server

Go, standart kütüphanesine **production grade** yani canlı ortamda gönül
rahatlığıyla kullanabileceğimiz http sunucuyla birlikte geliyor. Hatta sadece
sunucu değil istemcisi de var (http client). Tüm bu özellikler `net/http`
paketi içinde.

`net/http` bize neler sağlar?

- Web uygulamaları yapabiliriz
- Statik dosya sunucusu olarak kullanabiliriz
- Routing
- Cookie yönetimi yapabiliriz

Keza pek çok popüler web frameworkleri de altta bu paketi kullanır:

- https://github.com/go-chi/chi
- https://github.com/valyala/fasthttp
- https://github.com/labstack/echo
- https://github.com/gofiber/fiber
- https://github.com/gin-gonic/gin
- https://github.com/go-kratos/kratos (microservice)
- https://github.com/go-kit/kit (microservice)

`godoc`’ta:

- https://pkg.go.dev/net/http@go1.21.0 (http)
- https://pkg.go.dev/net/http@go1.21.0#hdr-Servers (server)

Bizim ilgilendiğimiz kısım [Server type][02]:

```go
type Server struct {
	// Addr optionally specifies the TCP address for the server to listen on,
	// in the form "host:port". If empty, ":http" (port 80) is used.
	// The service names are defined in RFC 6335 and assigned by IANA.
	// See net.Dial for details of the address format.
	Addr string

	Handler Handler // handler to invoke, http.DefaultServeMux if nil
    
    // others
}
```

`Handler` field’ı, `Handler` tipinde; Peki [nedir][03] bu?

```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

Evet, bu bir `interface`. Herhangi bir tipin `(ResponseWriter, *Request)` alan
bir fonksiyonu olursa o artık bir **HTTP handler** oluyor.

HTTP Handler ne yapar? istemcinin sunucudan yaptığı istekleri yakalayan ve
geriye cevap dönen (byte cinsinden) şeydir.

Soru: `ServeHTTP` neden `*Request` (pointer) alırken `ResponseWriter`’ı
(value) olarak alıyor? [ResponseWriter][04] dokümanı.

Özetle; bizim şöyle bir tipimiz olsa;

```go
type foo struct {}
func (foo)ServeHTTP(http.ResponseWriter, *http.Request){
}
```

artık `foo` bir HTTP Handler olarak kullanılabilir.

[örnek](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/18/basic-http-server)

https://go.dev/play/p/8dMERh1XZvg

```bash
$ go run src/18/basic-http-server/main.go
```

sonra tarayıcıyı açıp;

- http://127.0.0.1:9090/foo
- http://127.0.0.1:9090/bar

tebrikler, artık bir web sunucunuz var!

Biz, servisimizi geliştirirken biraz daha gelişmiş özellikleri olan
`http.Server`’ı kullanacağız. Bizim bir kısım endpoint’leri yakalamamız
gerekiyor. Bunun için **request multiplexer**’a [ihtiyacımız][05] var.

Hemen godoc’tan bir örneğe bakalım:

[örnek](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/src/18/basic-mux)

https://go.dev/play/p/T38GlxCvEpL

```bash
$ go run src/18/basic-mux/main.go
```

sonra tarayıcıyı açıp:

- http://127.0.0.1:9090/
- http://127.0.0.1:9090/api/

`http.Server` kullandığımız zaman;

```go
http.Server{
	Addr:         ":8000",
	Handler:      mux,
	ReadTimeout:  ServerReadTimeout,
	WriteTimeout: ServerWriteTimeout,
	IdleTimeout:  ServerIdleTimeout,
}
```

gibi ek parametrelerde kullanabiliyoruz.

---

[01]: https://redis.io/
[02]: https://pkg.go.dev/net/http@go1.21.0#Server
[03]: https://pkg.go.dev/net/http@go1.21.0#Handler
[04]: https://pkg.go.dev/net/http@go1.21.0#ResponseWriter
[05]: https://pkg.go.dev/net/http@go1.21.0#ServeMux