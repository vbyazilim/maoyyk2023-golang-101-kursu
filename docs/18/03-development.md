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
$ git commit -m 'add go.mod file'
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

Şimdi projeyi kod editöründe açalım ve
`src/internal/storage/memory/kvstorage/base.go` dosyasına şunu yazalım ve
kaydedelim:

```go
package kvstorage
```

Go koduna başladık, hemen [linter konfigürasyon](https://github.com/vbyazilim/maoyyk2023-golang-101-kursu/tree/main/.golangci.yml)
dosyamızı root dizine atalım, sonra `base.go` dosyasını aşağıdaki gibi
düzenleyelim;

```go
package kvstorage

import (
	"sync"
)

var _ Storer = (*memoryStorage)(nil) // compile time proof

// MemoryDB is a custom type definition uses map[string]any for in memory-db type.
type MemoryDB map[string]any

// Storer defines storage behaviours.
type Storer interface {
	Set(key string, value any) (any, error)
	Get(key string) (any, error)
	Update(key string, value any) (any, error)
	Delete(key string) error
	List() MemoryDB
}

type memoryStorage struct {
	mu sync.RWMutex // guarding db only
	db MemoryDB
}

// StorageOption represents storage option type.
type StorageOption func(*memoryStorage)

// WithMemoryDB sets db option.
func WithMemoryDB(db MemoryDB) StorageOption {
	return func(s *memoryStorage) {
		s.db = db
	}
}

// New instantiates new storage instance.
func New(options ...StorageOption) Storer {
	ms := &memoryStorage{}

	for _, o := range options {
		o(ms)
	}

	return ms
}
```

sonra;

```bash
$ git add .
$ git commit -m 'start storage implementation'
```

Şimdi tüm metotları implemente edelim:

```bash
$ touch src/internal/storage/memory/kvstorage/{delete,get,list,set,update}.go
$ tree .
.
├── go.mod
└── src
    └── internal
        └── storage
            └── memory
                └── kvstorage
                    ├── base.go
                    ├── delete.go
                    ├── get.go
                    ├── list.go
                    ├── set.go
                    └── update.go
```

Şimdi bize özel error tipimizi oluşturalım;

```bash
$ mkdir -p src/internal/kverror
$ touch src/internal/kverror/kverror.go
```

`src/internal/kverror/kverror.go`

```go
package kverror

var (
	_ error   = (*Error)(nil) // compile time proof
	_ KVError = (*Error)(nil) // compile time proof
)

// sentinel errors.
var (
	ErrKeyExists   = New("key exist", true)
	ErrKeyNotFound = New("key not found", false)
	ErrUnknown     = New("unknown error", true)
)

// KVError defines custom error behaviours.
type KVError interface {
	Wrap(err error) KVError
	Unwrap() error
	AddData(any) KVError
	DestoryData() KVError
	Error() string
}

// Error is a custom type definition uses struct, custom error.
type Error struct {
	Err      error
	Message  string
	Data     any `json:"-"`
	Loggable bool
}

// AddData adds extra data to error.
func (e *Error) AddData(data any) KVError {
	e.Data = data
	return e
}

// Unwrap unwraps error.
func (e *Error) Unwrap() error {
	return e.Err
}

// DestoryData removes added data from error.
func (e *Error) DestoryData() KVError {
	e.Data = nil
	return e
}

// Wrap wraps given error.
func (e *Error) Wrap(err error) KVError {
	e.Err = err
	return e
}

func (e *Error) Error() string {
	if e.Err != nil {
		return e.Err.Error() + ", " + e.Message
	}
	return e.Message
}

// New instantiates new Error instance.
func New(m string, l bool) KVError {
	return &Error{
		Message:  m,
		Loggable: l,
	}
}
```

sonra;

```bash
$ git add src/internal/kverror/kverror.go
$ git commit -m 'implement custom error type'
```

---

`src/internal/storage/memory/kvstorage/delete.go`

```go
package kvstorage

func (ms *memoryStorage) Delete(key string) error {
	if _, err := ms.Get(key); err != nil { // can not delete! key doesn't exist
		return err
	}

	ms.mu.Lock()
	defer ms.mu.Unlock()

	delete(ms.db, key)
	return nil
}
```

---

`src/internal/storage/memory/kvstorage/get.go`

```go
package kvstorage

import (
	"fmt"

	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/kverror"
)

func (ms *memoryStorage) Get(key string) (any, error) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	value, ok := ms.db[key]
	if !ok {
		return nil, fmt.Errorf("%w", kverror.ErrKeyNotFound.AddData("'"+key+"' does not exist"))
	}
	return value, nil
}
```

---

`src/internal/storage/memory/kvstorage/list.go`

```go
package kvstorage

func (ms *memoryStorage) List() MemoryDB {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	return ms.db
}
```

---

`src/internal/storage/memory/kvstorage/set.go`

```go
package kvstorage

import (
	"fmt"

	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/kverror"
)

func (ms *memoryStorage) Set(key string, value any) (any, error) {
	if _, err := ms.Get(key); err == nil {
		return nil, fmt.Errorf("%w", kverror.ErrKeyExists.AddData("'"+key+"' already exist"))
	}

	ms.mu.Lock()
	defer ms.mu.Unlock()

	ms.db[key] = value
	return value, nil
}
```

---

`src/internal/storage/memory/kvstorage/update.go`

```go
package kvstorage

func (ms *memoryStorage) Update(key string, value any) (any, error) {
	if _, err := ms.Get(key); err != nil { // can not update! key doesn't exist
		return nil, err
	}

	ms.mu.Lock()
	defer ms.mu.Unlock()

	ms.db[key] = value
	return value, nil
}
```

sonra;

```bash
$ git add .
$ git commit -m 'implement memory storage'
```

Nedir son durum ?

```bash
$ tree .
.
├── go.mod
└── src
    └── internal
        ├── kverror
        │   └── kverror.go
        └── storage
            └── memory
                └── kvstorage
                    ├── base.go
                    ├── delete.go
                    ├── get.go
                    ├── list.go
                    ├── set.go
                    └── update.go
```

---

## Service Layer

```bash
$ mkdir -p src/internal/service/kvstoreservice
$ touch src/internal/service/kvstoreservice/base.go
```

`src/internal/service/kvstoreservice/base.go`

```go
package kvstoreservice

import (
	"context"

	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/storage/memory/kvstorage"
)

var _ KVStoreService = (*kvStoreService)(nil) // compile time proof

// KVStoreService defines service behaviours.
type KVStoreService interface {
	Set(context.Context, *SetRequest) (*ItemResponse, error)
	Get(context.Context, string) (*ItemResponse, error)
	Update(context.Context, *UpdateRequest) (*ItemResponse, error)
	Delete(context.Context, string) error
	List(context.Context) (*ListResponse, error)
}

type kvStoreService struct {
	storage kvstorage.Storer
}

// ServiceOption represents service option type.
type ServiceOption func(*kvStoreService)

// WithStorage sets storage option.
func WithStorage(strg kvstorage.Storer) ServiceOption {
	return func(s *kvStoreService) {
		s.storage = strg
	}
}

// New instantiates new service instance.
func New(options ...ServiceOption) KVStoreService {
	kvs := &kvStoreService{}

	for _, o := range options {
		o(kvs)
	}

	return kvs
}
```

sonra;

```bash
$ touch src/internal/service/kvstoreservice/{delete,get,list,requests,responses,set,update}.go
$ tree .
.
├── go.mod
└── src
    └── internal
        ├── kverror
        │   └── kverror.go
        ├── service
        │   └── kvstoreservice
        │       ├── base.go
        │       ├── delete.go
        │       ├── get.go
        │       ├── list.go
        │       ├── requests.go
        │       ├── responses.go
        │       ├── set.go
        │       └── update.go
        └── storage
            └── memory
                └── kvstorage
                    ├── base.go
                    ├── delete.go
                    ├── get.go
                    ├── list.go
                    ├── set.go
                    └── update.go
```

---

`src/internal/service/kvstoreservice/delete.go`

```go
package kvstoreservice

import (
	"context"
	"fmt"
)

func (s *kvStoreService) Delete(ctx context.Context, key string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		if err := s.storage.Delete(key); err != nil {
			return fmt.Errorf("kvstoreservice.Set storage.Delete err: %w", err)
		}
		return nil
	}
}
```

---

`src/internal/service/kvstoreservice/get.go`

```go
package kvstoreservice

import (
	"context"
	"fmt"
)

func (s *kvStoreService) Get(ctx context.Context, key string) (*ItemResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		value, err := s.storage.Get(key)
		if err != nil {
			return nil, fmt.Errorf("kvstoreservice.Set storage.Get err: %w", err)
		}
		return &ItemResponse{
			Key:   key,
			Value: value,
		}, nil
	}
}
```

---

`src/internal/service/kvstoreservice/list.go`

```go
package kvstoreservice

import (
	"context"
)

func (s *kvStoreService) List(ctx context.Context) (*ListResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		items := s.storage.List()
		response := make(ListResponse, len(items))

		var i int
		for k, v := range items {
			response[i] = ItemResponse{
				Key:   k,
				Value: v,
			}
			i++
		}
		return &response, nil
	}
}
```

---

`src/internal/service/kvstoreservice/requests.go`

```go
package kvstoreservice

// SetRequest is an input payload for Set behaviour.
type SetRequest struct {
	Key   string
	Value any
}

// UpdateRequest is an input payload for Update behaviour.
type UpdateRequest struct {
	Key   string
	Value any
}
```

---

`src/internal/service/kvstoreservice/responses.go`

```go
package kvstoreservice

// ItemResponse represents common k/v response element.
type ItemResponse struct {
	Key   string
	Value any
}

// ListResponse is a collection on ItemResponse.
type ListResponse []ItemResponse
```

---

`src/internal/service/kvstoreservice/set.go`

```go
package kvstoreservice

import (
	"context"
	"fmt"
)

func (s *kvStoreService) Set(ctx context.Context, sr *SetRequest) (*ItemResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		if _, err := s.storage.Set(sr.Key, sr.Value); err != nil {
			return nil, fmt.Errorf("kvstoreservice.Set storage.Set err: %w", err)
		}

		return &ItemResponse{
			Key:   sr.Key,
			Value: sr.Value,
		}, nil
	}
}
```

---

`src/internal/service/kvstoreservice/update.go`

```go
package kvstoreservice

import (
	"context"
	"fmt"
)

func (s *kvStoreService) Update(ctx context.Context, sr *UpdateRequest) (*ItemResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		value, err := s.storage.Update(sr.Key, sr.Value)
		if err != nil {
			return nil, fmt.Errorf("kvstoreservice.Set storage.Update err: %w", err)
		}
		return &ItemResponse{
			Key:   sr.Key,
			Value: value,
		}, nil
	}
}
```

sonra;

```bash
$ git add .
$ git commit -m 'implement service layer'
```

---

## HTTP Handler Layer

```bash
$ mkdir -p src/internal/transport/http/{basehttp,kvstore}handler
$ touch src/internal/transport/http/basehttphandler/basehttphandler.go
$ touch src/internal/transport/http/kvstorehandler/base.go

$ tree .
.
├── go.mod
└── src
    └── internal
        ├── kverror
        │   └── kverror.go
        ├── service
        │   └── kvstoreservice
        │       ├── base.go
        │       ├── delete.go
        │       ├── get.go
        │       ├── list.go
        │       ├── requests.go
        │       ├── responses.go
        │       ├── set.go
        │       └── update.go
        ├── storage
        │   └── memory
        │       └── kvstorage
        │           ├── base.go
        │           ├── delete.go
        │           ├── get.go
        │           ├── list.go
        │           ├── set.go
        │           └── update.go
        └── transport
            └── http
                ├── basehttphandler
                │   └── basehttphandler.go
                └── kvstorehandler
                    ├── base.go
                    ├── delete.go
                    ├── get.go
                    ├── list.go
                    ├── set.go
                    └── update.go
```

`src/internal/transport/http/basehttphandler/basehttphandler.go`

```go
package basehttphandler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

// Handler respresents common http handler functionality.
type Handler struct {
	ServerEnv     string
	Logger        *slog.Logger
	CancelTimeout time.Duration
}

// JSON generates json response.
func (h *Handler) JSON(w http.ResponseWriter, status int, d any) {
	j, err := json.Marshal(d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, _ = w.Write(j)
}
```

---

`src/internal/transport/http/kvstorehandler/base.go`

```go
package kvstorehandler

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/service/kvstoreservice"
	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/transport/http/basehttphandler"
)

var _ KVStoreHTTPHandler = (*kvstoreHandler)(nil) // compile time proof

// KVStoreHTTPHandler defines /store/ http handler behaviours.
type KVStoreHTTPHandler interface {
	Set(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	List(http.ResponseWriter, *http.Request)
}

type kvstoreHandler struct {
	basehttphandler.Handler

	service kvstoreservice.KVStoreService
}

// StoreHandlerOption represents store handler option type.
type StoreHandlerOption func(*kvstoreHandler)

// WithService sets service option.
func WithService(srvc kvstoreservice.KVStoreService) StoreHandlerOption {
	return func(s *kvstoreHandler) {
		s.service = srvc
	}
}

// WithContextTimeout sets handler context cancel timeout.
func WithContextTimeout(d time.Duration) StoreHandlerOption {
	return func(s *kvstoreHandler) {
		s.Handler.CancelTimeout = d
	}
}

// WithServerEnv sets handler server env.
func WithServerEnv(env string) StoreHandlerOption {
	return func(s *kvstoreHandler) {
		s.Handler.ServerEnv = env
	}
}

// WithLogger sets handler logger.
func WithLogger(l *slog.Logger) StoreHandlerOption {
	return func(s *kvstoreHandler) {
		s.Handler.Logger = l
	}
}

// New instantiates new kvstoreHandler instance.
func New(options ...StoreHandlerOption) KVStoreHTTPHandler {
	kvsh := &kvstoreHandler{
		Handler: basehttphandler.Handler{},
	}

	for _, o := range options {
		o(kvsh)
	}

	return kvsh
}
```

---

`src/internal/transport/http/kvstorehandler/delete.go`

```go
package kvstorehandler

import (
	"context"
	"errors"
	"net/http"

	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/kverror"
)

func (h *kvstoreHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		h.JSON(
			w,
			http.StatusMethodNotAllowed,
			map[string]string{"error": "method " + r.Method + " not allowed"},
		)
		return
	}

	if len(r.URL.Query()) == 0 {
		h.JSON(
			w,
			http.StatusNotFound,
			map[string]string{"error": "key query param required"},
		)
		return
	}

	keys, ok := r.URL.Query()["key"]
	if !ok {
		h.JSON(
			w,
			http.StatusNotFound,
			map[string]string{"error": "key not present"},
		)
		return
	}

	key := keys[0]

	ctx, cancel := context.WithTimeout(r.Context(), h.CancelTimeout)
	defer cancel()

	if err := h.service.Delete(ctx, key); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			h.JSON(
				w,
				http.StatusGatewayTimeout,
				map[string]string{"error": err.Error()},
			)
			return
		}

		var kvErr *kverror.Error

		if errors.As(err, &kvErr) {
			clientMessage := kvErr.Message
			if kvErr.Data != nil {
				data, ok := kvErr.Data.(string)
				if ok {
					clientMessage = clientMessage + ", " + data
				}
			}

			if kvErr.Loggable {
				h.Logger.Error("kvstorehandler Delete service.Delete", "err", clientMessage)
			}

			if kvErr == kverror.ErrKeyNotFound {
				h.JSON(w, http.StatusNotFound, map[string]string{"error": clientMessage})
				return
			}
		}
		h.JSON(
			w,
			http.StatusInternalServerError,
			map[string]string{"error": err.Error()},
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
```

---

`src/internal/transport/http/kvstorehandler/get.go`

```bash
package kvstorehandler

import (
	"context"
	"errors"
	"net/http"

	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/kverror"
)

func (h *kvstoreHandler) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.JSON(
			w,
			http.StatusMethodNotAllowed,
			map[string]string{"error": "method " + r.Method + " not allowed"},
		)
		return
	}

	if len(r.URL.Query()) == 0 {
		h.JSON(
			w,
			http.StatusNotFound,
			map[string]string{"error": "key query param required"},
		)
		return
	}

	keys, ok := r.URL.Query()["key"]
	if !ok {
		h.JSON(
			w,
			http.StatusNotFound,
			map[string]string{"error": "key not present"},
		)
		return
	}

	key := keys[0]

	ctx, cancel := context.WithTimeout(r.Context(), h.CancelTimeout)
	defer cancel()

	serviceResponse, err := h.service.Get(ctx, key)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			h.JSON(
				w,
				http.StatusGatewayTimeout,
				map[string]string{"error": err.Error()},
			)
			return
		}

		var kvErr *kverror.Error

		if errors.As(err, &kvErr) {
			clientMessage := kvErr.Message
			if kvErr.Data != nil {
				data, ok := kvErr.Data.(string)
				if ok {
					clientMessage = clientMessage + ", " + data
				}
			}

			if kvErr.Loggable {
				h.Logger.Error("kvstorehandler Get service.Get", "err", clientMessage)
			}

			if kvErr == kverror.ErrKeyNotFound {
				h.JSON(w, http.StatusNotFound, map[string]string{"error": clientMessage})
				return
			}
		}
		h.JSON(
			w,
			http.StatusInternalServerError,
			map[string]string{"error": err.Error()},
		)
		return
	}

	handlerResponse := ItemResponse{
		Key:   serviceResponse.Key,
		Value: serviceResponse.Value,
	}

	h.JSON(
		w,
		http.StatusOK,
		handlerResponse,
	)
}
```

---

`src/internal/transport/http/kvstorehandler/list.go`

```go
package kvstorehandler

import (
	"context"
	"errors"
	"net/http"

	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/kverror"
)

func (h *kvstoreHandler) List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.JSON(
			w,
			http.StatusMethodNotAllowed,
			map[string]string{"error": "method " + r.Method + " not allowed"},
		)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), h.CancelTimeout)
	defer cancel()

	serviceResponse, err := h.service.List(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			h.JSON(
				w,
				http.StatusGatewayTimeout,
				map[string]string{"error": err.Error()},
			)
			return
		}

		var kvErr *kverror.Error
		if errors.As(err, &kvErr) {
			clientMessage := kvErr.Message
			if kvErr.Data != nil {
				data, ok := kvErr.Data.(string)
				if ok {
					clientMessage = clientMessage + ", " + data
				}
			}

			if kvErr.Loggable {
				h.Logger.Error("kvstorehandler List service.List", "err", clientMessage)
			}
		}

		h.JSON(
			w,
			http.StatusInternalServerError,
			map[string]string{"error": err.Error()},
		)
		return
	}

	var handlerResponse ListResponse
	for _, item := range *serviceResponse {
		handlerResponse = append(handlerResponse, ItemResponse{
			Key:   item.Key,
			Value: item.Value,
		})
	}

	if len(handlerResponse) == 0 {
		h.JSON(
			w,
			http.StatusNotFound,
			map[string]string{"error": "nothing found"},
		)
		return
	}

	h.JSON(
		w,
		http.StatusOK,
		handlerResponse,
	)
}
```

---

`src/internal/transport/http/kvstorehandler/requests.go`

```go
package kvstorehandler

// SetRequest is an input payload for creating new k/v item.
type SetRequest struct {
	Key   string `json:"key"`
	Value any    `json:"value"`
}

// UpdateRequest is an input payload for updating existing k/v item.
type UpdateRequest struct {
	Key   string `json:"key"`
	Value any    `json:"value"`
}
```

---

`src/internal/transport/http/kvstorehandler/responses.go`

```go
package kvstorehandler

// ItemResponse represents k/v item.
type ItemResponse struct {
	Key   string `json:"key"`
	Value any    `json:"value"`
}

// ListResponse represents collection of ItemResponse.
type ListResponse []ItemResponse
```

---

`src/internal/transport/http/kvstorehandler/set.go`

```go
package kvstorehandler

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/kverror"
	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/service/kvstoreservice"
)

func (h *kvstoreHandler) Set(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.JSON(
			w,
			http.StatusMethodNotAllowed,
			map[string]string{"error": "method " + r.Method + " not allowed"},
		)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.JSON(
			w,
			http.StatusBadRequest,
			map[string]string{"error": err.Error()},
		)
		return
	}

	if len(body) == 0 {
		h.JSON(
			w,
			http.StatusBadRequest,
			map[string]string{"error": "empty body/payload"},
		)
		return
	}

	var handlerRequest SetRequest
	if err = json.Unmarshal(body, &handlerRequest); err != nil {
		h.JSON(
			w,
			http.StatusInternalServerError,
			map[string]string{"error": err.Error()},
		)
		return
	}

	if handlerRequest.Key == "" {
		h.JSON(
			w,
			http.StatusBadRequest,
			map[string]string{"error": "key is empty"},
		)
		return
	}

	if handlerRequest.Value == nil {
		h.JSON(
			w,
			http.StatusBadRequest,
			map[string]string{"error": "value is empty"},
		)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), h.CancelTimeout)
	defer cancel()

	existingItem, err := h.service.Get(ctx, handlerRequest.Key)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			h.JSON(
				w,
				http.StatusGatewayTimeout,
				map[string]string{"error": err.Error()},
			)
			return
		}

		var kvErr *kverror.Error
		if errors.As(err, &kvErr) {
			clientMessage := kvErr.Message
			if kvErr.Data != nil {
				data, ok := kvErr.Data.(string)
				if ok {
					clientMessage = clientMessage + ", " + data
				}
			}

			if kvErr.Loggable {
				h.Logger.Error("kvstorehandler Set service.Get", "err", clientMessage)
			}

			if kvErr != kverror.ErrKeyNotFound {
				h.JSON(
					w,
					http.StatusBadRequest,
					map[string]string{"error": clientMessage},
				)
				return
			}
		}
	}

	// this should be nil. means, key does not exist
	if existingItem != nil {
		h.JSON(
			w,
			http.StatusConflict,
			map[string]string{"error": "can not set, '" + handlerRequest.Key + "' already exists"},
		)
		return
	}

	serviceRequest := kvstoreservice.SetRequest{
		Key:   handlerRequest.Key,
		Value: handlerRequest.Value,
	}

	serviceResponse, err := h.service.Set(ctx, &serviceRequest)
	if err != nil {
		var kvErr *kverror.Error

		if errors.As(err, &kvErr) {
			clientMessage := kvErr.Message
			if kvErr.Data != nil {
				data, ok := kvErr.Data.(string)
				if ok {
					clientMessage = clientMessage + ", " + data
				}
			}

			if kvErr.Loggable {
				h.Logger.Error("kvstorehandler Set service.Set", "err", clientMessage)
			}

			if kvErr == kverror.ErrKeyExists {
				h.JSON(w, http.StatusConflict, map[string]string{"error": clientMessage})
				return
			}
		}

		h.JSON(
			w,
			http.StatusInternalServerError,
			map[string]string{"error": err.Error()},
		)
		return
	}

	handlerResponse := ItemResponse{
		Key:   serviceResponse.Key,
		Value: serviceResponse.Value,
	}

	h.JSON(
		w,
		http.StatusCreated,
		handlerResponse,
	)
}
```

---

`src/internal/transport/http/kvstorehandler/update.go`

```go
package kvstorehandler

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/kverror"
	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/service/kvstoreservice"
)

func (h *kvstoreHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		h.JSON(
			w,
			http.StatusMethodNotAllowed,
			map[string]string{"error": "method " + r.Method + " not allowed"},
		)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.JSON(
			w,
			http.StatusBadRequest,
			map[string]string{"error": err.Error()},
		)
		return
	}

	if len(body) == 0 {
		h.JSON(
			w,
			http.StatusBadRequest,
			map[string]string{"error": "empty body/payload"},
		)
		return
	}

	var handlerRequest UpdateRequest
	if err = json.Unmarshal(body, &handlerRequest); err != nil {
		h.JSON(
			w,
			http.StatusInternalServerError,
			map[string]string{"error": err.Error()},
		)
		return
	}

	if handlerRequest.Key == "" {
		h.JSON(
			w,
			http.StatusBadRequest,
			map[string]string{"error": "key is empty"},
		)
		return
	}

	if handlerRequest.Value == nil {
		h.JSON(
			w,
			http.StatusBadRequest,
			map[string]string{"error": "value is empty"},
		)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), h.CancelTimeout)
	defer cancel()

	serviceRequest := kvstoreservice.UpdateRequest{
		Key:   handlerRequest.Key,
		Value: handlerRequest.Value,
	}

	serviceResponse, err := h.service.Update(ctx, &serviceRequest)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			h.JSON(
				w,
				http.StatusGatewayTimeout,
				map[string]string{"error": err.Error()},
			)
			return
		}

		var kvErr *kverror.Error

		if errors.As(err, &kvErr) {
			clientMessage := kvErr.Message
			if kvErr.Data != nil {
				data, ok := kvErr.Data.(string)
				if ok {
					clientMessage = clientMessage + ", " + data
				}
			}

			if kvErr.Loggable {
				h.Logger.Error("kvstorehandler Update service.Update", "err", clientMessage)
			}

			if kvErr == kverror.ErrKeyNotFound {
				h.JSON(w, http.StatusNotFound, map[string]string{"error": clientMessage})
				return
			}
		}

		h.JSON(
			w,
			http.StatusInternalServerError,
			map[string]string{"error": err.Error()},
		)
		return
	}

	handlerResponse := ItemResponse{
		Key:   serviceResponse.Key,
		Value: serviceResponse.Value,
	}

	h.JSON(
		w,
		http.StatusOK,
		handlerResponse,
	)
}
```

sonra;

```bash
$ git add .
$ git commit -m 'implement http handlers'
```

---

## releaseinfo paketi

```bash
$ mkdir -p src/releaseinfo
$ touch src/releaseinfo/releaseinfo.go
```

`src/releaseinfo/releaseinfo.go`

```go
package releaseinfo

// Version is the current version of service.
const Version string = "0.0.0"

// BuildInformation holds current build information.
var BuildInformation string
```

sonra;

```bash
$ git add src/releaseinfo/releaseinfo.go
$ git commit -m 'add release information package'
```

---

## apiserver paketi

```bash
$ mkdir -p src/apiserver
$ touch src/apiserver/{apiserver,middlewares}.go

$ tree .
.
├── go.mod
└── src
    ├── apiserver
    │   ├── apiserver.go
    │   └── middlewares.go
    ├── internal
    │   ├── kverror
    │   │   └── kverror.go
    │   ├── service
    │   │   └── kvstoreservice
    │   │       ├── base.go
    │   │       ├── delete.go
    │   │       ├── get.go
    │   │       ├── list.go
    │   │       ├── requests.go
    │   │       ├── responses.go
    │   │       ├── set.go
    │   │       └── update.go
    │   ├── storage
    │   │   └── memory
    │   │       └── kvstorage
    │   │           ├── base.go
    │   │           ├── delete.go
    │   │           ├── get.go
    │   │           ├── list.go
    │   │           ├── set.go
    │   │           └── update.go
    │   └── transport
    │       └── http
    │           ├── basehttphandler
    │           │   └── basehttphandler.go
    │           └── kvstorehandler
    │               ├── base.go
    │               ├── delete.go
    │               ├── get.go
    │               ├── list.go
    │               ├── set.go
    │               └── update.go
    └── releaseinfo
        └── releaseinfo.go
```

---

`src/apiserver/apiserver.go`

```go
package apiserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/service/kvstoreservice"
	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/storage/memory/kvstorage"
	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/transport/http/kvstorehandler"
	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/releaseinfo"
)

// constants.
const (
	ContextCancelTimeout = 5 * time.Second
	ShutdownTimeout      = 10 * time.Second
	ServerReadTimeout    = 10 * time.Second
	ServerWriteTimeout   = 10 * time.Second
	ServerIdleTimeout    = 60 * time.Second

	apiV1Prefix = "/api/v1"
)

type apiServer struct {
	db        kvstorage.MemoryDB
	logLevel  slog.Level
	logger    *slog.Logger
	serverEnv string
}

// Option represents api server option type.
type Option func(*apiServer)

// WithLogger sets logger option.
func WithLogger(l *slog.Logger) Option {
	return func(s *apiServer) {
		s.logger = l
	}
}

// WithServerEnv sets serverEnv option.
func WithServerEnv(env string) Option {
	return func(s *apiServer) {
		s.serverEnv = env
	}
}

// WithLogLevel sets logLevel option.
func WithLogLevel(level string) Option {
	return func(s *apiServer) {
		var logLevel slog.Level

		switch level {
		case "DEBUG":
			logLevel = slog.LevelDebug
		case "WARN":
			logLevel = slog.LevelWarn
		case "ERROR":
			logLevel = slog.LevelError
		default:
			logLevel = slog.LevelInfo
		}

		s.logLevel = logLevel
	}
}

// New instantiates new server instance.
func New(options ...Option) error {
	apisrvr := &apiServer{
		db:       kvstorage.MemoryDB(make(map[string]any)), // default db
		logLevel: slog.LevelInfo,
	}

	for _, o := range options {
		o(apisrvr)
	}

	// default logging options if logger not present.
	if apisrvr.logger == nil {
		logHandlerOpts := &slog.HandlerOptions{Level: apisrvr.logLevel}
		logHandler := slog.NewJSONHandler(os.Stdout, logHandlerOpts)
		apisrvr.logger = slog.New(logHandler)
	}
	slog.SetDefault(apisrvr.logger)

	if apisrvr.serverEnv == "" {
		apisrvr.serverEnv = "production" // default server environment
	}

	logger := apisrvr.logger

	storage := kvstorage.New(
		kvstorage.WithMemoryDB(apisrvr.db),
	)
	service := kvstoreservice.New(
		kvstoreservice.WithStorage(storage),
	)
	kvStoreHandler := kvstorehandler.New(
		kvstorehandler.WithService(service),
		kvstorehandler.WithContextTimeout(ContextCancelTimeout),
		kvstorehandler.WithServerEnv(apisrvr.serverEnv),
		kvstorehandler.WithLogger(logger),
	)

	mux := http.NewServeMux()

	mux.HandleFunc("/healthz/live/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		j, _ := json.Marshal(map[string]any{
			"server":            apisrvr.serverEnv,
			"version":           releaseinfo.Version,
			"build_information": releaseinfo.BuildInformation,
			"message":           "liveness is OK!, server is ready to accept connections",
		})
		_, _ = w.Write(j)
	})
	mux.HandleFunc("/healthz/ready/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		j, _ := json.Marshal(map[string]any{
			"server":            apisrvr.serverEnv,
			"version":           releaseinfo.Version,
			"build_information": releaseinfo.BuildInformation,
			"message":           "readiness is OK!, server is ready to accept connections",
		})
		_, _ = w.Write(j)
	})

	mux.HandleFunc(apiV1Prefix+"/set/", kvStoreHandler.Set)
	mux.HandleFunc(apiV1Prefix+"/get/", kvStoreHandler.Get)
	mux.HandleFunc(apiV1Prefix+"/update/", kvStoreHandler.Update)
	mux.HandleFunc(apiV1Prefix+"/delete/", kvStoreHandler.Delete)
	mux.HandleFunc(apiV1Prefix+"/list/", kvStoreHandler.List)

	api := &http.Server{
		Addr:         ":8000",
		Handler:      appendSlashMiddleware(httpLoggingMiddleware(logger, mux)),
		ReadTimeout:  ServerReadTimeout,
		WriteTimeout: ServerWriteTimeout,
		IdleTimeout:  ServerIdleTimeout,
	}

	shutdown := make(chan os.Signal, 1)
	apiError := make(chan error, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.Info("starting api server", "listening", api.Addr, "env", apisrvr.serverEnv)
		apiError <- api.ListenAndServe()
	}()

	select {
	case err := <-apiError:
		return fmt.Errorf("listen and server err: %w", err)
	case sig := <-shutdown:
		logger.Info("starting shutdown", "pid", sig)
		defer logger.Info("shutdown completed", "pid", sig)

		ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
		defer cancel()

		if err := api.Shutdown(ctx); err != nil {
			if errr := api.Close(); errr != nil {
				logger.Error("api close", "err", errr)
			}
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}

	return nil
}
```

---

`src/apiserver/middlewares.go`

```go
package apiserver

import (
	"log/slog"
	"net/http"
	"strings"
)

func httpLoggingMiddleware(l *slog.Logger, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)

		uri := r.URL.String()
		method := r.Method

		l.Info("http request", "method", method, "uri", uri)
	}

	return http.HandlerFunc(fn)
}

func appendSlashMiddleware(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" && !strings.HasSuffix(r.URL.Path, "/") {
			redirectURL := r.URL.Path + "/"
			if r.URL.RawQuery != "" {
				redirectURL += "?" + r.URL.RawQuery
			}
			http.Redirect(w, r, redirectURL, http.StatusPermanentRedirect)
			return
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
```

sonra;

```bash
$ git add .
$ git commit -m 'add apiserver'
```

---

Artık esas sunucuyu çalıştıracak kısma geldik;

```bash
$ mkdir -p cmd/server
$ touch cmd/server/main.go
```

`cmd/server/main.go`

```go
package main

import (
	"log"
	"os"

	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/apiserver"
)

func main() {
	if err := apiserver.New(
		apiserver.WithServerEnv(os.Getenv("SERVER_ENV")),
		apiserver.WithLogLevel(os.Getenv("LOG_LEVEL")),
	); err != nil {
		log.Fatal(err)
	}
}
```

sonra;

```bash
$ git add .
$ git commit -m 'add server'
```

---

Evet, şimdi kodumuzu linter’dan geçirelim;

```bash
$ golangci-lint version    # v1.54.1
$ golangci-lint run
```

eğer her şey OK ise;

```bash
$ go run -race cmd/server/main.go
```

---

## İstekleri Yapalım

Evet, şu an sunucumuz çalışıyor. İster `curl` ister `httpie` ile denemelere
başlayalım:

`curl` örnekleri:

```bash
# add new key/value
$ curl -L -s -X POST -H "Content-Type: application/json" -d '{"key": "success", "value": true}' "http://localhost:8000/api/v1/set" | jq
{
  "key": "success",
  "value": true
}

$ curl -L -s -X POST -H "Content-Type: application/json" -d '{"key": "server_env", "value": "production"}' "http://localhost:8000/api/v1/set" | jq
{
  "key": "server_env",
  "value": "production"
}

$ curl -L -s -H "Content-Type: application/json" "http://localhost:8000/api/v1/list" | jq
[
  {
    "key": "success",
    "value": true
  },
  {
    "key": "server_env",
    "value": "production"
  }
]

$ curl -L -s -X PUT -H "Content-Type: application/json" -d '{"key": "success", "value": false}' "http://localhost:8000/api/v1/update" | jq
{
  "key": "success",
  "value": false
}

$ curl -L -s -H "Content-Type: application/json" "http://localhost:8000/api/v1/list" | jq
[
  {
    "key": "success",
    "value": false
  },
  {
    "key": "server_env",
    "value": "production"
  }
]

$ curl -L -s -H "Content-Type: application/json" "http://localhost:8000/api/v1/get?key=success" | jq
{
  "key": "success",
  "value": false
}

$ curl -L -s -X DELETE -H "Content-Type: application/json" -o /dev/null -w '%{http_code}\n' "http://localhost:8000/api/v1/delete?key=success"
204

$ curl -L -s -H "Content-Type: application/json" "http://localhost:8000/api/v1/list" | jq
[
  {
    "key": "server_env",
    "value": "production"
  }
]
```

---

`httpie` örnekleri:

```bash
$ http POST "http://localhost:8000/api/v1/set" key="success" value:=true
$ http POST "http://localhost:8000/api/v1/set" key="server_env" value="production"
$ http "http://localhost:8000/api/v1/list"
$ http PUT "http://localhost:8000/api/v1/update" key="success" value:=false
$ http "http://localhost:8000/api/v1/get?key=success"
$ http DELETE "http://localhost:8000/api/v1/delete?key=success"
```

---

## Kaynaklar

- https://github.com/avelino/awesome-go#project-layout
- https://go.dev/blog/slog

---
