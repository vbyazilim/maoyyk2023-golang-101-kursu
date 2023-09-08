# Bölüm 18/04: In-Memory Key-Value Store

## Test

Go, **test first** bir dil. Yani testler hayati derecede önemli. Bir projenin
ne kadar iyi test coverage’ı olursa, o proje / kütüphane / servis o kadar
sağlam çalışır anlamına gelir.

Kabaca, yazdığımız her satır kod, fonksiyon, metot, yani her şey test
edilebilir şeylerdir. Eğer yazdığınız kodu test edemiyorsanız, o zaman bir
sıkıntı var demektir. Bir şeyleri hatalı yapmış ya da atlamışsınızdır.

Nelerin testlerini yapmamız iyi olur?

- Ek paket yaptık mı? (`kverror`)
- Storage katmanı (tüm metotları)
- Service katmanı (tüm metotları)
- HTTP Handler katmanı (tüm metotları)

İyi bir test coverage yüzdesi ~ `%80` civarındadır. Yani yazılan kodun en az
`%80`’i cover edilmişse bu iş **OK**’dir. (%80 - %20 yaklaşımı) Coverage ne
kadar yüksek olursa kendimizi o kadar güvende hissederiz.

Şimdi testlere başlayalım; önce `kverror`:

```bash
$ touch src/internal/kverror/kverror_test.go
```

`src/internal/kverror/kverror_test.go`

```go
package kverror_test

import (
	"errors"
	"testing"

	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/kverror"
)

func TestError(t *testing.T) {
	err := kverror.New("some error", true)
	var kvErr *kverror.Error

	if !errors.As(err, &kvErr) {
		t.Errorf("error does not match the target type, want: %T, got: %v", kvErr, err)
	}

	shouldEqual := "some error"
	if kvErr.Message != shouldEqual {
		t.Errorf("error message does not match, want: %s, got: %s", shouldEqual, kvErr.Message)
	}

	shouldLoggable := true
	if kvErr.Loggable != shouldLoggable {
		t.Errorf("error should be loggable, want: %t, got: %t", shouldLoggable, kvErr.Loggable)
	}
}

func TestWrap(t *testing.T) {
	err := kverror.New("some error", false)
	wrappedErr := err.Wrap(errors.New("inner")) // nolint

	var kvErr *kverror.Error

	if !errors.As(wrappedErr, &kvErr) {
		t.Errorf("error does not match the target type, want: %T, got: %v", kvErr, err)
	}

	if kvErr.Err == nil {
		t.Errorf("wrapped error can not be nil, want: %v, got: nil", kvErr.Err)
	}

	shouldEqual := "inner, some error"
	if err.Error() != shouldEqual {
		t.Errorf("wrapped error does not match, want: %s, got: %s", shouldEqual, err.Error())
	}
}

func TestUnwrap(t *testing.T) {
	err := kverror.New("some error", false)
	wrappedErr := err.Wrap(errors.New("inner")) // nolint

	var kvErr *kverror.Error

	if !errors.As(wrappedErr, &kvErr) {
		t.Errorf("error does not match the target type, want: %T, got: %v", kvErr, err)
	}

	shouldEqual := "inner"
	unwrappedErr := kvErr.Unwrap()
	if unwrappedErr.Error() != shouldEqual {
		t.Errorf("unwrapped error does not match, want: %s, got: %s", shouldEqual, unwrappedErr.Error())
	}
}

func TestAddDataDestroyData(t *testing.T) {
	err := kverror.New("some error", false).AddData("hello")

	var kvErr *kverror.Error

	if !errors.As(err, &kvErr) {
		t.Errorf("error does not match the target type, want: %T, got: %v", kvErr, err)
	}

	if kvErr.Data == nil {
		t.Errorf("data should not be nil, want: %v, got: nil", kvErr.Data)
	}

	shouldEqual := "hello"
	data, ok := kvErr.Data.(string)
	if !ok {
		t.Error("data should be assertable to string")
	}

	if data != shouldEqual {
		t.Errorf("data does not match, want: %s, got: %s", shouldEqual, data)
	}

	shouldEqual = "some error"
	if err.Error() != shouldEqual {
		t.Errorf("error does not match, want: %s, got: %s", shouldEqual, err.Error())
	}

	err = err.DestoryData()
	if !errors.As(err, &kvErr) {
		t.Errorf("error does not match the target type, want: %T, got: %v", kvErr, err)
	}

	if kvErr.Data != nil {
		t.Errorf("data should be nil, want: nil, got: %v", kvErr.Data)
	}
}
```

şimdi testi çalıştıralım; önce paketleri bulalım;

```bash
$ go list ./... | grep 'kverror'
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/kverror

$ go test -race -v github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/kverror
$ go test -cover -race -v github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/kverror
:
:
coverage: 100.0% of statements
```

sonra;

```bash
$ git add src/internal/kverror/kverror_test.go
$ git commit -m 'add kverror test'
```

---

## Storage Testleri

```bash
$ touch src/internal/storage/memory/kvstorage/{delete,get,list,set,update}_test.go
```

`src/internal/storage/memory/kvstorage/delete_test.go`

```go
package kvstorage_test

import (
	"testing"

	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/storage/memory/kvstorage"
)

func TestDeleteEmpty(t *testing.T) {
	storage := kvstorage.New()

	if err := storage.Delete("key"); err == nil {
		t.Error("error not occurred")
	}
}

func TestDelete(t *testing.T) {
	key := "key"
	memoryStorage := map[string]any{
		key: "value",
	}
	storage := kvstorage.New(
		kvstorage.WithMemoryDB(memoryStorage),
	)

	if err := storage.Delete(key); err != nil {
		t.Error("error occurred")
	}
}
```

---

`src/internal/storage/memory/kvstorage/get_test.go`

```go
package kvstorage_test

import (
	"testing"

	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/storage/memory/kvstorage"
)

func TestGetEmpty(t *testing.T) {
	storage := kvstorage.New()

	if _, err := storage.Get("key"); err == nil {
		t.Error("error not occurred")
	}
}

func TestGet(t *testing.T) {
	key := "key"
	memoryStorage := map[string]any{
		key: "value",
	}
	storage := kvstorage.New(
		kvstorage.WithMemoryDB(memoryStorage),
	)

	value, err := storage.Get(key)
	if err != nil {
		t.Error("error occurred")
	}

	if value != "value" {
		t.Error("value not equal")
	}
}
```

---

`src/internal/storage/memory/kvstorage/list_test.go`

```go
package kvstorage_test

import (
	"reflect"
	"testing"

	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/storage/memory/kvstorage"
)

func TestList(t *testing.T) {
	key := "key"
	memoryStorage := kvstorage.MemoryDB(map[string]any{
		key: "value",
	})
	storage := kvstorage.New(
		kvstorage.WithMemoryDB(memoryStorage),
	)

	value := storage.List()

	if !reflect.DeepEqual(value, memoryStorage) {
		t.Error("value not equal")
	}
}
```

---

`src/internal/storage/memory/kvstorage/set_test.go`

```go
package kvstorage_test

import (
	"testing"

	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/storage/memory/kvstorage"
)

func TestSet(t *testing.T) {
	key := "key"
	memoryStorage := kvstorage.MemoryDB(map[string]any{})
	storage := kvstorage.New(
		kvstorage.WithMemoryDB(memoryStorage),
	)

	val, err := storage.Set(key, "value")
	if err != nil {
		t.Errorf("want: value, got: %v, err: %v", val, err)
	}

	if _, err := storage.Set(key, "xxx"); err == nil {
		t.Error("error not occurred")
	}
}
```

---

`src/internal/storage/memory/kvstorage/update_test.go`

```go
package kvstorage_test

import (
	"testing"

	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/storage/memory/kvstorage"
)

func TestUpdateEmpty(t *testing.T) {
	storage := kvstorage.New()

	if _, err := storage.Update("key", "value"); err == nil {
		t.Error("error not occurred")
	}
}

func TestUpdate(t *testing.T) {
	key := "key"
	memoryStorage := map[string]any{
		key: "value",
	}
	storage := kvstorage.New(
		kvstorage.WithMemoryDB(memoryStorage),
	)

	value, err := storage.Update(key, "value2")
	if err != nil {
		t.Error("error occurred")
	}

	if value != "value2" {
		t.Error("value not equal")
	}
}
```

Durum ne?

```bash
$ tree .
.
├── cmd
│   └── server
│       └── main.go
├── go.mod
└── src
    ├── apiserver
    │   ├── apiserver.go
    │   └── middlewares.go
    ├── internal
    │   ├── kverror
    │   │   ├── kverror.go
    │   │   └── kverror_test.go
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
    │   │           ├── delete_test.go
    │   │           ├── get.go
    │   │           ├── get_test.go
    │   │           ├── list.go
    │   │           ├── list_test.go
    │   │           ├── set.go
    │   │           ├── set_test.go
    │   │           ├── update.go
    │   │           └── update_test.go
    │   └── transport
    │       └── http
    │           ├── basehttphandler
    │           │   └── basehttphandler.go
    │           └── kvstorehandler
    │               └── base.go
    └── releaseinfo
        └── releaseinfo.go
```

şimdi testi çalıştıralım; önce paketleri bulalım;

```bash
$ go list ./... | grep 'kvstorage'
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/storage/memory/kvstorage

$ go test -race -v github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/storage/memory/kvstorage
$ go test -cover -race -v github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/storage/memory/kvstorage
:
:
coverage: 100.0% of statements
```

sonra;

```bash
$ git add .
$ git commit -m 'add storage tests'
```


---

## Service Testleri

```bash
$ touch src/internal/service/kvstoreservice/{base,delete,get,list,set,update}_test.go
$ tree .
.
├── cmd
│   └── server
│       └── main.go
├── go.mod
└── src
    ├── apiserver
    │   ├── apiserver.go
    │   └── middlewares.go
    ├── internal
    │   ├── kverror
    │   │   ├── kverror.go
    │   │   └── kverror_test.go
    │   ├── service
    │   │   └── kvstoreservice
    │   │       ├── base.go
    │   │       ├── base_test.go
    │   │       ├── delete.go
    │   │       ├── delete_test.go
    │   │       ├── get.go
    │   │       ├── get_test.go
    │   │       ├── list.go
    │   │       ├── list_test.go
    │   │       ├── requests.go
    │   │       ├── responses.go
    │   │       ├── set.go
    │   │       ├── set_test.go
    │   │       ├── update.go
    │   │       └── update_test.go
    │   ├── storage
    │   │   └── memory
    │   │       └── kvstorage
    │   │           ├── base.go
    │   │           ├── delete.go
    │   │           ├── delete_test.go
    │   │           ├── get.go
    │   │           ├── get_test.go
    │   │           ├── list.go
    │   │           ├── list_test.go
    │   │           ├── set.go
    │   │           ├── set_test.go
    │   │           ├── update.go
    │   │           └── update_test.go
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

`src/internal/service/kvstoreservice/base_test.go`

```go
package kvstoreservice_test

import (
	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/storage/memory/kvstorage"
)

var _ kvstorage.Storer = (*mockStorage)(nil) // compile time proof

type mockStorage struct {
	deleteErr error
	getErr    error
	updateErr error
	setErr    error

	memoryDB kvstorage.MemoryDB
}

func (m *mockStorage) Delete(k string) error {
	if m.deleteErr == nil {
		delete(m.memoryDB, k)
		return nil
	}
	return m.deleteErr
}

func (m *mockStorage) Get(k string) (any, error) {
	if m.getErr == nil {
		v, ok := m.memoryDB[k]
		if !ok {
			return nil, m.getErr
		}
		return v, nil
	}
	return nil, m.getErr
}

func (m *mockStorage) List() kvstorage.MemoryDB {
	return m.memoryDB
}

func (m *mockStorage) Set(k string, v any) (any, error) {
	if m.setErr == nil {
		if _, ok := m.memoryDB[k]; ok {
			return nil, m.setErr
		}

		m.memoryDB[k] = v
		return v, nil

	}
	return nil, m.setErr
}

func (m *mockStorage) Update(k string, v any) (any, error) {
	if m.updateErr == nil {
		if _, ok := m.memoryDB[k]; !ok {
			return nil, m.updateErr
		}

		m.memoryDB[k] = v
		return v, nil
	}
	return nil, m.updateErr
}
```

---

`src/internal/service/kvstoreservice/delete_test.go`

```go
package kvstoreservice_test

import (
	"context"
	"errors"
	"testing"

	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/kverror"
	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/service/kvstoreservice"
)

func TestDeleteWithCancel(t *testing.T) {
	mockStorage := &mockStorage{}
	kvsStoreService := kvstoreservice.New(
		kvstoreservice.WithStorage(mockStorage),
	)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if err := kvsStoreService.Delete(ctx, "key"); !errors.Is(err, ctx.Err()) {
		t.Error("error not occurred")
	}
}

func TestDeleteWithStorageError(t *testing.T) {
	mockStorage := &mockStorage{
		deleteErr: kverror.ErrKeyNotFound,
	}
	kvsStoreService := kvstoreservice.New(
		kvstoreservice.WithStorage(mockStorage),
	)

	err := kvsStoreService.Delete(context.Background(), "key")
	if err == nil {
		t.Error("error not occurred")
	}

	var kvErr *kverror.Error

	if !errors.As(err, &kvErr) {
		t.Error("error must be kverror.ErrKeyNotFound")
	}
}

func TestDelete(t *testing.T) {
	mockStorage := &mockStorage{
		memoryDB: map[string]any{
			"key": "value",
		},
	}

	kvsStoreService := kvstoreservice.New(
		kvstoreservice.WithStorage(mockStorage),
	)

	if err := kvsStoreService.Delete(context.Background(), "key"); err != nil {
		t.Error("error occurred")
	}

	_, ok := mockStorage.memoryDB["key"]
	if ok {
		t.Error("delete is not working!")
	}
}
```

---

`src/internal/service/kvstoreservice/get_test.go`

```go
package kvstoreservice_test

import (
	"context"
	"errors"
	"testing"

	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/kverror"
	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/service/kvstoreservice"
)

func TestGetWithCancel(t *testing.T) {
	mockStorage := &mockStorage{}
	kvsStoreService := kvstoreservice.New(kvstoreservice.WithStorage(mockStorage))

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if _, err := kvsStoreService.Get(ctx, "key"); !errors.Is(err, ctx.Err()) {
		t.Error("error not occurred")
	}
}

func TestGetWithStorageError(t *testing.T) {
	mockStorage := &mockStorage{
		getErr: kverror.ErrKeyNotFound, // get raises ErrKeyNotFound
	}
	kvsStoreService := kvstoreservice.New(kvstoreservice.WithStorage(mockStorage))

	res, err := kvsStoreService.Get(context.Background(), "key")
	if err == nil {
		t.Error("error not occurred")
	}

	if res != nil {
		t.Errorf("response must be nil!")
	}

	var kvErr *kverror.Error

	if !errors.As(err, &kvErr) {
		t.Error("error must be kverror.ErrKeyNotFound")
	}
}

func TestGet(t *testing.T) {
	mockStorage := &mockStorage{
		memoryDB: map[string]any{
			"key": "value",
		},
	}
	kvsStoreService := kvstoreservice.New(kvstoreservice.WithStorage(mockStorage))

	res, err := kvsStoreService.Get(context.Background(), "key")
	if err != nil {
		t.Error("error occurred")
	}

	if res == nil {
		t.Error("result should not be nil")
	}

	if res != nil {
		val := *res
		if val.Value != "value" {
			t.Errorf("want: value, got: %s", val.Value)
		}
	}
}
```

---

`src/internal/service/kvstoreservice/list_test.go`

```go
package kvstoreservice_test

import (
	"context"
	"errors"
	"testing"

	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/service/kvstoreservice"
)

func TestListWithCancel(t *testing.T) {
	mockStorage := &mockStorage{}
	kvsStoreService := kvstoreservice.New(kvstoreservice.WithStorage(mockStorage))

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if _, err := kvsStoreService.List(ctx); !errors.Is(err, ctx.Err()) {
		t.Error("error not occurred")
	}
}

func TestList(t *testing.T) {
	mockStorage := &mockStorage{
		memoryDB: map[string]any{
			"key": "value",
		},
	}
	kvsStoreService := kvstoreservice.New(kvstoreservice.WithStorage(mockStorage))

	if _, err := kvsStoreService.List(context.Background()); err != nil {
		t.Error("error occurred")
	}
}
```

---

`src/internal/service/kvstoreservice/set_test.go`

```go
package kvstoreservice_test

import (
	"context"
	"errors"
	"testing"

	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/kverror"
	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/service/kvstoreservice"
)

func TestSetWithCancel(t *testing.T) {
	mockStorage := &mockStorage{}
	kvsStoreService := kvstoreservice.New(
		kvstoreservice.WithStorage(mockStorage),
	)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if _, err := kvsStoreService.Set(ctx, nil); !errors.Is(err, ctx.Err()) {
		t.Error("error not occurred")
	}
}

func TestSetWithStorageError(t *testing.T) {
	mockStorage := &mockStorage{
		setErr: kverror.ErrKeyExists,
	}
	kvsStoreService := kvstoreservice.New(
		kvstoreservice.WithStorage(mockStorage),
	)

	serviceRequest := kvstoreservice.SetRequest{
		Key:   "vigo",
		Value: "lego",
	}

	res, err := kvsStoreService.Set(context.Background(), &serviceRequest)

	if res != nil {
		t.Errorf("response must be nil!")
	}

	var kvErr *kverror.Error

	if !errors.As(err, &kvErr) {
		t.Error("error must be kverror.ErrKeyExists")
	}
}

func TestSet(t *testing.T) {
	mockStorage := &mockStorage{
		memoryDB: map[string]any{},
	}
	kvsStoreService := kvstoreservice.New(
		kvstoreservice.WithStorage(mockStorage),
	)

	setRequest := kvstoreservice.SetRequest{
		Key:   "username",
		Value: "vigo",
	}

	res, err := kvsStoreService.Set(context.Background(), &setRequest)
	if err != nil {
		t.Errorf("error occurred, err: %v", err)
	}

	if res == nil {
		t.Error("result should not be nil")
	}

	if res != nil {
		val := *res

		if val.Value != "vigo" {
			t.Errorf("want: vigo, got: %s", val.Value)
		}
	}
}
```

---

`src/internal/service/kvstoreservice/update_test.go`

```go
package kvstoreservice_test

import (
	"context"
	"errors"
	"testing"

	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/kverror"
	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/service/kvstoreservice"
)

func TestUpdateWithCancel(t *testing.T) {
	mockStorage := &mockStorage{}
	kvsStoreService := kvstoreservice.New(
		kvstoreservice.WithStorage(mockStorage),
	)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if _, err := kvsStoreService.Update(ctx, nil); !errors.Is(err, ctx.Err()) {
		t.Error("error not occurred")
	}
}

func TestUpdateWithStorageError(t *testing.T) {
	mockStorage := &mockStorage{
		updateErr: kverror.ErrKeyNotFound, // raises kverror.ErrKeyNotFound
	}
	kvsStoreService := kvstoreservice.New(
		kvstoreservice.WithStorage(mockStorage),
	)

	updateRequest := kvstoreservice.UpdateRequest{
		Key:   "key",
		Value: "value",
	}

	res, err := kvsStoreService.Update(context.Background(), &updateRequest)
	if res != nil {
		t.Errorf("response must be nil!")
	}

	var kvErr *kverror.Error

	if !errors.As(err, &kvErr) {
		t.Error("error must be kverror.ErrKeyNotFound")
	}
}

func TestUpdate(t *testing.T) {
	mockStorage := &mockStorage{
		memoryDB: map[string]any{
			"key": "value",
		},
	}
	kvsStoreService := kvstoreservice.New(
		kvstoreservice.WithStorage(mockStorage),
	)

	updateRequest := kvstoreservice.UpdateRequest{
		Key:   "key",
		Value: "vigo",
	}

	res, err := kvsStoreService.Update(context.Background(), &updateRequest)
	if err != nil {
		t.Errorf("error occurred, err: %v", err)
	}

	if res == nil {
		t.Error("result should not be nil")
	}

	if res != nil {
		val := *res

		if val.Value != "vigo" {
			t.Errorf("want: vigo, got: %s", val.Value)
		}
	}
}
```

şimdi testi çalıştıralım; önce paketleri bulalım;

```bash
$ go list ./... | grep 'kvstoreservice'
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/service/kvstoreservice

$ go test -race -v github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/service/kvstoreservice
$ go test -cover -race -v github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/service/kvstoreservice
:
:
coverage: 100.0% of statements
```

sonra;

```bash
$ git add .
$ git commit -m 'add service tests'
```

---

## HTTP Handler Testleri

```bash
$ touch src/internal/transport/http/kvstorehandler/{base,delete,get,list,set,update}_test.go
$ tree .
.
├── cmd
│   └── server
│       └── main.go
├── go.mod
└── src
    ├── apiserver
    │   ├── apiserver.go
    │   └── middlewares.go
    ├── internal
    │   ├── kverror
    │   │   ├── kverror.go
    │   │   └── kverror_test.go
    │   ├── service
    │   │   └── kvstoreservice
    │   │       ├── base.go
    │   │       ├── base_test.go
    │   │       ├── delete.go
    │   │       ├── delete_test.go
    │   │       ├── get.go
    │   │       ├── get_test.go
    │   │       ├── list.go
    │   │       ├── list_test.go
    │   │       ├── requests.go
    │   │       ├── responses.go
    │   │       ├── set.go
    │   │       ├── set_test.go
    │   │       ├── update.go
    │   │       └── update_test.go
    │   ├── storage
    │   │   └── memory
    │   │       └── kvstorage
    │   │           ├── base.go
    │   │           ├── delete.go
    │   │           ├── delete_test.go
    │   │           ├── get.go
    │   │           ├── get_test.go
    │   │           ├── list.go
    │   │           ├── list_test.go
    │   │           ├── set.go
    │   │           ├── set_test.go
    │   │           ├── update.go
    │   │           └── update_test.go
    │   └── transport
    │       └── http
    │           ├── basehttphandler
    │           │   └── basehttphandler.go
    │           └── kvstorehandler
    │               ├── base.go
    │               ├── base_test.go
    │               ├── delete.go
    │               ├── delete_test.go
    │               ├── get.go
    │               ├── get_test.go
    │               ├── list.go
    │               ├── list_test.go
    │               ├── set.go
    │               ├── set_test.go
    │               ├── update.go
    │               └── update_test.go
    └── releaseinfo
        └── releaseinfo.go
```

---

`src/internal/transport/http/kvstorehandler/base_test.go`

```go
package kvstorehandler_test

import (
	"context"
	"log/slog"
	"os"

	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/service/kvstoreservice"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

type mockService struct {
	deleteErr      error
	getErr         error
	getResponse    *kvstoreservice.ItemResponse
	listErr        error
	listResponse   *kvstoreservice.ListResponse
	setErr         error
	setResponse    *kvstoreservice.ItemResponse
	updateErr      error
	updateResponse *kvstoreservice.ItemResponse
}

func (m *mockService) Delete(_ context.Context, _ string) error {
	return m.deleteErr
}

func (m *mockService) Get(_ context.Context, _ string) (*kvstoreservice.ItemResponse, error) {
	return m.getResponse, m.getErr
}

func (m *mockService) List(_ context.Context) (*kvstoreservice.ListResponse, error) {
	return m.listResponse, m.listErr
}

func (m *mockService) Set(_ context.Context, _ *kvstoreservice.SetRequest) (*kvstoreservice.ItemResponse, error) {
	return m.setResponse, m.setErr
}

func (m *mockService) Update(_ context.Context, _ *kvstoreservice.UpdateRequest) (*kvstoreservice.ItemResponse, error) {
	return m.updateResponse, m.updateErr
}
```

---

`src/internal/transport/http/kvstorehandler/delete_test.go`

```go
package kvstorehandler_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/kverror"
	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/transport/http/kvstorehandler"
)

func TestDeleteInvalidMethod(t *testing.T) {
	handler := kvstorehandler.New()
	req := httptest.NewRequest(http.MethodGet, "/key", nil)
	w := httptest.NewRecorder()

	handler.Delete(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusMethodNotAllowed, w.Code)
	}

	shouldContain := "method GET not allowed"
	if !strings.Contains(w.Body.String(), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, w.Body.String())
	}
}

func TestDeleteQueryParamRequired(t *testing.T) {
	handler := kvstorehandler.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	w := httptest.NewRecorder()

	handler.Delete(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusNotFound, w.Code)
	}

	shouldContain := "key query param required"
	if !strings.Contains(w.Body.String(), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, w.Body.String())
	}
}

func TestDeleteQueryParamKeyNotFound(t *testing.T) {
	handler := kvstorehandler.New()
	req := httptest.NewRequest(http.MethodDelete, "/?foo=test", nil)
	w := httptest.NewRecorder()

	handler.Delete(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusNotFound, w.Code)
	}

	shouldContain := "key not present"
	if !strings.Contains(w.Body.String(), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, w.Body.String())
	}
}

func TestDeleteTimeout(t *testing.T) {
	handler := kvstorehandler.New(
		kvstorehandler.WithContextTimeout(time.Second*-1),
		kvstorehandler.WithService(&mockService{
			deleteErr: context.DeadlineExceeded,
		}),
	)

	req := httptest.NewRequest(http.MethodDelete, "/?key=test", nil)
	w := httptest.NewRecorder()

	handler.Delete(w, req)

	if w.Code != http.StatusGatewayTimeout {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusGatewayTimeout, w.Code)
	}

	shouldContain := "context deadline exceeded"
	if !strings.Contains(w.Body.String(), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, w.Body.String())
	}
}

func TestDeleteErrUnknown(t *testing.T) {
	handler := kvstorehandler.New(
		kvstorehandler.WithService(&mockService{
			deleteErr: kverror.ErrUnknown,
		}),
		kvstorehandler.WithLogger(logger),
	)

	req := httptest.NewRequest(http.MethodDelete, "/?key=test", nil)
	w := httptest.NewRecorder()

	handler.Delete(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusInternalServerError, w.Code)
	}

	shouldContain := "unknown error"
	if !strings.Contains(w.Body.String(), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, w.Body.String())
	}
}

func TestDeleteErrKeyNotFound(t *testing.T) {
	_ = kverror.ErrKeyNotFound.AddData("key=test") // ignore error.

	handler := kvstorehandler.New(
		kvstorehandler.WithService(&mockService{
			deleteErr: kverror.ErrKeyNotFound,
		}),
		kvstorehandler.WithLogger(logger),
	)

	req := httptest.NewRequest(http.MethodDelete, "/?key=test", nil)
	w := httptest.NewRecorder()

	handler.Delete(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusNotFound, w.Code)
	}

	if !strings.Contains(w.Body.String(), "key not found") {
		t.Error("body not equal")
	}

	shouldContain := "key=test"
	if !strings.Contains(w.Body.String(), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, w.Body.String())
	}

	_ = kverror.ErrKeyNotFound.DestoryData() // ignore error.
}

func TestDeleteSuccess(t *testing.T) {
	handler := kvstorehandler.New(
		kvstorehandler.WithService(&mockService{}),
		kvstorehandler.WithLogger(logger),
	)

	req := httptest.NewRequest(http.MethodDelete, "/?key=test", nil)
	w := httptest.NewRecorder()

	handler.Delete(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusNoContent, w.Code)
	}

	if w.Body.Len() != 0 {
		t.Errorf("wrong body size, want: 0, got: %d", w.Body.Len())
	}
}
```

---

`src/internal/transport/http/kvstorehandler/get_test.go`

```go
package kvstorehandler_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/kverror"
	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/service/kvstoreservice"
	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/transport/http/kvstorehandler"
)

func TestGetInvalidMethod(t *testing.T) {
	handler := kvstorehandler.New()
	req := httptest.NewRequest(http.MethodDelete, "/key", nil)
	w := httptest.NewRecorder()

	handler.Get(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusMethodNotAllowed, w.Code)
	}

	shouldContain := "method DELETE not allowed"
	if !strings.Contains(w.Body.String(), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, w.Body.String())
	}
}

func TestGetQueryParamRequired(t *testing.T) {
	handler := kvstorehandler.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	handler.Get(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusNotFound, w.Code)
	}

	shouldContain := "key query param required"
	if !strings.Contains(w.Body.String(), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, w.Body.String())
	}
}

func TestGetQueryParamKeyNotFound(t *testing.T) {
	handler := kvstorehandler.New()
	req := httptest.NewRequest(http.MethodGet, "/?foo=test", nil)
	w := httptest.NewRecorder()

	handler.Get(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusNotFound, w.Code)
	}

	shouldContain := "key not present"
	if !strings.Contains(w.Body.String(), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, w.Body.String())
	}
}

func TestGetTimeout(t *testing.T) {
	handler := kvstorehandler.New(
		kvstorehandler.WithContextTimeout(time.Second*-1),
		kvstorehandler.WithService(&mockService{
			getErr: context.DeadlineExceeded,
		}),
	)

	req := httptest.NewRequest(http.MethodGet, "/?key=test", nil)
	w := httptest.NewRecorder()

	handler.Get(w, req)

	if w.Code != http.StatusGatewayTimeout {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusGatewayTimeout, w.Code)
	}

	shouldContain := "context deadline exceeded"
	if !strings.Contains(w.Body.String(), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, w.Body.String())
	}
}

func TestGetErrUnknown(t *testing.T) {
	handler := kvstorehandler.New(
		kvstorehandler.WithService(&mockService{
			getErr: kverror.ErrUnknown,
		}),
		kvstorehandler.WithLogger(logger),
	)

	req := httptest.NewRequest(http.MethodGet, "/?key=test", nil)
	w := httptest.NewRecorder()

	handler.Get(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusInternalServerError, w.Code)
	}

	shouldContain := "unknown error"
	if !strings.Contains(w.Body.String(), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, w.Body.String())
	}
}

func TestGetErrKeyNotFound(t *testing.T) {
	_ = kverror.ErrKeyNotFound.AddData("key=test") // ignore error.

	handler := kvstorehandler.New(
		kvstorehandler.WithService(&mockService{
			getErr: kverror.ErrKeyNotFound,
		}),
		kvstorehandler.WithLogger(logger),
	)

	req := httptest.NewRequest(http.MethodGet, "/?key=test", nil)
	w := httptest.NewRecorder()

	handler.Get(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusNotFound, w.Code)
	}

	shouldContain := "key not found"
	if !strings.Contains(w.Body.String(), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, w.Body.String())
	}

	shouldContain = "key=test"
	if !strings.Contains(w.Body.String(), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, w.Body.String())
	}

	_ = kverror.ErrKeyNotFound.DestoryData() // ignore error.
}

func TestGetSuccess(t *testing.T) {
	handler := kvstorehandler.New(
		kvstorehandler.WithService(&mockService{}),
		kvstorehandler.WithLogger(logger),
		kvstorehandler.WithService(&mockService{
			getResponse: &kvstoreservice.ItemResponse{
				Key:   "test",
				Value: "test",
			},
		}),
	)

	req := httptest.NewRequest(http.MethodGet, "/?key=test", nil)
	w := httptest.NewRecorder()

	handler.Get(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusOK, w.Code)
	}

	shouldEqual := `{"key":"test","value":"test"}`
	if w.Body.String() != shouldEqual {
		t.Errorf("wrong body message, want: %s, got: %s", shouldEqual, w.Body.String())
	}
}
```

---

`src/internal/transport/http/kvstorehandler/list_test.go`

```go
package kvstorehandler_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/kverror"
	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/service/kvstoreservice"
	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/transport/http/kvstorehandler"
)

func TestListInvalidMethod(t *testing.T) {
	handler := kvstorehandler.New()
	req := httptest.NewRequest(http.MethodDelete, "/key", nil)
	w := httptest.NewRecorder()

	handler.List(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusMethodNotAllowed, w.Code)
	}

	shouldContain := "method DELETE not allowed"
	if !strings.Contains(w.Body.String(), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, w.Body.String())
	}
}

func TestListTimeout(t *testing.T) {
	handler := kvstorehandler.New(
		kvstorehandler.WithContextTimeout(time.Second*-1),
		kvstorehandler.WithService(&mockService{
			listErr: context.DeadlineExceeded,
		}),
	)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	handler.List(w, req)

	if w.Code != http.StatusGatewayTimeout {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusGatewayTimeout, w.Code)
	}

	shouldContain := "context deadline exceeded"
	if !strings.Contains(w.Body.String(), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, w.Body.String())
	}
}

func TestListErrUnknown(t *testing.T) {
	handler := kvstorehandler.New(
		kvstorehandler.WithService(&mockService{
			listErr: kverror.ErrUnknown.AddData("fake error"),
		}),
		kvstorehandler.WithLogger(logger),
	)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	handler.List(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusInternalServerError, w.Code)
	}

	shouldContain := "unknown error"
	if !strings.Contains(w.Body.String(), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, w.Body.String())
	}
}

func TestEmptyList(t *testing.T) {
	handler := kvstorehandler.New(
		kvstorehandler.WithService(&mockService{}),
		kvstorehandler.WithLogger(logger),
		kvstorehandler.WithService(&mockService{
			listResponse: &kvstoreservice.ListResponse{},
		}),
	)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	handler.List(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusNotFound, w.Code)
	}
}

func TestListSuccess(t *testing.T) {
	handler := kvstorehandler.New(
		kvstorehandler.WithService(&mockService{}),
		kvstorehandler.WithLogger(logger),
		kvstorehandler.WithService(&mockService{
			listResponse: &kvstoreservice.ListResponse{
				{
					Key:   "test",
					Value: "test",
				},
			},
		}),
	)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	handler.List(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusOK, w.Code)
	}

	shouldEqual := `[{"key":"test","value":"test"}]`
	if w.Body.String() != shouldEqual {
		t.Errorf("wrong body message, want: %s, got: %s", shouldEqual, w.Body.String())
	}
}
```

---

`src/internal/transport/http/kvstorehandler/set_test.go`

```go
package kvstorehandler_test

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/kverror"
	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/service/kvstoreservice"
	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/transport/http/kvstorehandler"
)

type errorReader struct{}

func (e *errorReader) Read(_ []byte) (n int, err error) {
	return 0, errors.New("forced error") // nolint
}

func TestSetInvalidMethod(t *testing.T) {
	handler := kvstorehandler.New()
	req := httptest.NewRequest(http.MethodDelete, "/key", nil)
	w := httptest.NewRecorder()

	handler.Set(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusMethodNotAllowed, w.Code)
	}

	shouldContain := "method DELETE not allowed"
	if !strings.Contains(w.Body.String(), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, w.Body.String())
	}
}

func TestSetBodyReadError(t *testing.T) {
	handler := kvstorehandler.New()
	req := httptest.NewRequest(http.MethodPost, "/key", &errorReader{})

	w := httptest.NewRecorder()

	handler.Set(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusBadRequest, w.Code)
	}
}

func TestSetBodyUnmarshal(t *testing.T) {
	handler := kvstorehandler.New()
	handlerRequest := bytes.NewBufferString(`{"key": "key", "value": "123}`)
	req := httptest.NewRequest(http.MethodPost, "/key", handlerRequest)
	w := httptest.NewRecorder()

	handler.Set(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusInternalServerError, w.Code)
	}
}

func TestSetEmptyBody(t *testing.T) {
	handler := kvstorehandler.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	w := httptest.NewRecorder()

	handler.Set(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusBadRequest, w.Code)
	}

	shouldContain := "empty body/payload"
	if !strings.Contains(w.Body.String(), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, w.Body.String())
	}
}

func TestSetKeyIsEmpty(t *testing.T) {
	handler := kvstorehandler.New()

	payload := strings.NewReader("{}")
	req := httptest.NewRequest(http.MethodPost, "/", payload)
	w := httptest.NewRecorder()

	handler.Set(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusBadRequest, w.Code)
	}

	shouldContain := "key is empty"
	if !strings.Contains(w.Body.String(), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, w.Body.String())
	}
}

func TestSetValueIsEmpty(t *testing.T) {
	handler := kvstorehandler.New()

	payload := strings.NewReader(`{"key":"test"}`)
	req := httptest.NewRequest(http.MethodPost, "/", payload)
	w := httptest.NewRecorder()

	handler.Set(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusBadRequest, w.Code)
	}

	shouldContain := "value is empty"
	if !strings.Contains(w.Body.String(), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, w.Body.String())
	}
}

func TestSetTimeout(t *testing.T) {
	handler := kvstorehandler.New(
		kvstorehandler.WithContextTimeout(time.Second*-1),
		kvstorehandler.WithService(&mockService{
			getErr: context.DeadlineExceeded,
		}),
	)

	payload := strings.NewReader(`{"key":"test","value":"test"}`)
	req := httptest.NewRequest(http.MethodPost, "/?key=test", payload)
	w := httptest.NewRecorder()

	handler.Set(w, req)

	if w.Code != http.StatusGatewayTimeout {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusGatewayTimeout, w.Code)
	}

	shouldContain := "context deadline exceeded"
	if !strings.Contains(w.Body.String(), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, w.Body.String())
	}
}

func TestSetErrUnknown(t *testing.T) {
	handler := kvstorehandler.New(
		kvstorehandler.WithService(&mockService{
			setErr: kverror.ErrUnknown,
		}),
		kvstorehandler.WithLogger(logger),
	)

	payload := strings.NewReader(`{"key":"test","value":"test"}`)
	req := httptest.NewRequest(http.MethodPost, "/", payload)
	w := httptest.NewRecorder()

	handler.Set(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusInternalServerError, w.Code)
	}

	shouldContain := "unknown error"
	if !strings.Contains(w.Body.String(), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, w.Body.String())
	}
}

func TestSetServiceUnknownError(t *testing.T) {
	handler := kvstorehandler.New(
		kvstorehandler.WithService(&mockService{
			getErr: kverror.ErrUnknown.AddData("fake error"),
		}),
		kvstorehandler.WithLogger(logger),
	)

	payload := strings.NewReader(`{"key":"test","value":"test"}`)
	req := httptest.NewRequest(http.MethodPost, "/", payload)
	w := httptest.NewRecorder()

	handler.Set(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusBadRequest, w.Code)
	}
}

func TestSetServiceNilExistingItem(t *testing.T) {
	handler := kvstorehandler.New(
		kvstorehandler.WithService(&mockService{}),
		kvstorehandler.WithLogger(logger),
		kvstorehandler.WithService(&mockService{
			getResponse: &kvstoreservice.ItemResponse{
				Key:   "test",
				Value: "test",
			},
		}),
	)

	payload := strings.NewReader(`{"key":"test","value":"test"}`)
	req := httptest.NewRequest(http.MethodPost, "/", payload)
	w := httptest.NewRecorder()

	handler.Set(w, req)

	if w.Code != http.StatusConflict {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusConflict, w.Code)
	}
}

func TestSetErrKeyExists(t *testing.T) {
	_ = kverror.ErrKeyExists.AddData("key=test") // ignore error.

	handler := kvstorehandler.New(
		kvstorehandler.WithService(&mockService{
			setErr: kverror.ErrKeyExists,
		}),
		kvstorehandler.WithLogger(logger),
	)

	payload := strings.NewReader(`{"key":"test","value":"test"}`)
	req := httptest.NewRequest(http.MethodPost, "/", payload)
	w := httptest.NewRecorder()

	handler.Set(w, req)

	if w.Code != http.StatusConflict {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusConflict, w.Code)
	}

	shouldContain := "key exist"
	if !strings.Contains(w.Body.String(), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, w.Body.String())
	}

	shouldContain = "key=test"
	if !strings.Contains(w.Body.String(), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, w.Body.String())
	}

	_ = kverror.ErrKeyExists.DestoryData() // ignore error.
}

func TestSetSuccess(t *testing.T) {
	handler := kvstorehandler.New(
		kvstorehandler.WithService(&mockService{}),
		kvstorehandler.WithLogger(logger),
		kvstorehandler.WithService(&mockService{
			setResponse: &kvstoreservice.ItemResponse{
				Key:   "test",
				Value: "test",
			},
		}),
	)

	payload := strings.NewReader(`{"key":"test","value":"test"}`)
	req := httptest.NewRequest(http.MethodPost, "/", payload)
	w := httptest.NewRecorder()

	handler.Set(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusCreated, w.Code)
	}

	shouldEqual := `{"key":"test","value":"test"}`
	if w.Body.String() != shouldEqual {
		t.Errorf("wrong body message, want: %s, got: %s", shouldEqual, w.Body.String())
	}
}
```

---

`src/internal/transport/http/kvstorehandler/update_test.go`

```go
package kvstorehandler_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/kverror"
	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/service/kvstoreservice"
	"github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/transport/http/kvstorehandler"
)

func TestUpdateInvalidMethod(t *testing.T) {
	handler := kvstorehandler.New()
	req := httptest.NewRequest(http.MethodDelete, "/key", nil)
	w := httptest.NewRecorder()

	handler.Update(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusMethodNotAllowed, w.Code)
	}

	shouldContain := "method DELETE not allowed"
	if !strings.Contains(w.Body.String(), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, w.Body.String())
	}
}

func TestUpdateBodyReadError(t *testing.T) {
	handler := kvstorehandler.New()
	req := httptest.NewRequest(http.MethodPut, "/", &errorReader{})

	w := httptest.NewRecorder()

	handler.Update(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusBadRequest, w.Code)
	}
}

func TestUpdateEmptyBody(t *testing.T) {
	handler := kvstorehandler.New()
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	w := httptest.NewRecorder()

	handler.Update(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusBadRequest, w.Code)
	}

	shouldContain := "empty body/payload"
	if !strings.Contains(w.Body.String(), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, w.Body.String())
	}
}

func TestUpdateBodyUnmarshal(t *testing.T) {
	handler := kvstorehandler.New()
	handlerRequest := bytes.NewBufferString(`{"key": "key", "value": "123}`)
	req := httptest.NewRequest(http.MethodPut, "/", handlerRequest)
	w := httptest.NewRecorder()

	handler.Update(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusInternalServerError, w.Code)
	}
}

func TestUpdateKeyIsEmpty(t *testing.T) {
	handler := kvstorehandler.New()

	payload := strings.NewReader("{}")
	req := httptest.NewRequest(http.MethodPut, "/", payload)
	w := httptest.NewRecorder()

	handler.Update(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusBadRequest, w.Code)
	}

	shouldContain := "key is empty"
	if !strings.Contains(w.Body.String(), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, w.Body.String())
	}
}

func TestUpdateValueIsEmpty(t *testing.T) {
	handler := kvstorehandler.New()

	payload := strings.NewReader(`{"key":"test"}`)
	req := httptest.NewRequest(http.MethodPut, "/", payload)
	w := httptest.NewRecorder()

	handler.Update(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusBadRequest, w.Code)
	}

	shouldContain := "value is empty"
	if !strings.Contains(w.Body.String(), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, w.Body.String())
	}
}

func TestUpdateTimeout(t *testing.T) {
	handler := kvstorehandler.New(
		kvstorehandler.WithContextTimeout(time.Second*-1),
		kvstorehandler.WithService(&mockService{
			updateErr: context.DeadlineExceeded,
		}),
	)

	payload := strings.NewReader(`{"key":"test","value":"test"}`)
	req := httptest.NewRequest(http.MethodPut, "/?key=test", payload)
	w := httptest.NewRecorder()

	handler.Update(w, req)

	if w.Code != http.StatusGatewayTimeout {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusGatewayTimeout, w.Code)
	}

	shouldContain := "context deadline exceeded"
	if !strings.Contains(w.Body.String(), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, w.Body.String())
	}
}

func TestUpdateErrUnknown(t *testing.T) {
	handler := kvstorehandler.New(
		kvstorehandler.WithService(&mockService{
			updateErr: kverror.ErrUnknown,
		}),
		kvstorehandler.WithLogger(logger),
	)

	payload := strings.NewReader(`{"key":"test","value":"test"}`)
	req := httptest.NewRequest(http.MethodPut, "/", payload)
	w := httptest.NewRecorder()

	handler.Update(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusInternalServerError, w.Code)
	}

	shouldContain := "unknown error"
	if !strings.Contains(w.Body.String(), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, w.Body.String())
	}
}

func TestUpdateErrKeyExists(t *testing.T) {
	_ = kverror.ErrKeyNotFound.AddData("key=test") // ignore return no need

	handler := kvstorehandler.New(
		kvstorehandler.WithService(&mockService{
			updateErr: kverror.ErrKeyNotFound,
		}),
		kvstorehandler.WithLogger(logger),
	)

	payload := strings.NewReader(`{"key":"test","value":"test"}`)
	req := httptest.NewRequest(http.MethodPut, "/", payload)
	w := httptest.NewRecorder()

	handler.Update(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusNotFound, w.Code)
	}

	shouldContain := "key not found"
	if !strings.Contains(w.Body.String(), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, w.Body.String())
	}

	shouldContain = "key=test"
	if !strings.Contains(w.Body.String(), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, w.Body.String())
	}

	_ = kverror.ErrKeyNotFound.DestoryData() // ignore error
}

func TestUpdateSuccess(t *testing.T) {
	handler := kvstorehandler.New(
		kvstorehandler.WithService(&mockService{}),
		kvstorehandler.WithLogger(logger),
		kvstorehandler.WithService(&mockService{
			updateResponse: &kvstoreservice.ItemResponse{
				Key:   "test",
				Value: "test",
			},
		}),
	)

	payload := strings.NewReader(`{"key":"test","value":"test"}`)
	req := httptest.NewRequest(http.MethodPut, "/", payload)
	w := httptest.NewRecorder()

	handler.Update(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusOK, w.Code)
	}

	shouldEqual := `{"key":"test","value":"test"}`
	if w.Body.String() != shouldEqual {
		t.Errorf("wrong body message, want: %s, got: %s", shouldEqual, w.Body.String())
	}
}
```

şimdi testi çalıştıralım; önce paketleri bulalım;

```bash
$ go list ./... | grep 'kvstorehandler'
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/transport/http/kvstorehandler

$ go test -race -v github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/transport/http/kvstorehandler
$ go test -cover -race -v github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/transport/http/kvstorehandler
:
:
coverage: 99.1% of statements
```

evet, testler bitti, tüm test coverage ne durumda?

```bash
$ go test -coverpkg=./... -coverprofile=coverage.out ./...
$ go tool cover -func=coverage.out
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/kverror/kverror.go:33:				AddData			100.0%
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/kverror/kverror.go:39:				Unwrap			100.0%
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/kverror/kverror.go:44:				DestoryData		100.0%
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/kverror/kverror.go:50:				Wrap			100.0%
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/kverror/kverror.go:55:				Error			100.0%
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/kverror/kverror.go:63:				New			100.0%
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/service/kvstoreservice/base.go:28:			WithStorage		100.0%
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/service/kvstoreservice/base.go:35:			New			100.0%
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/service/kvstoreservice/delete.go:8:			Delete			100.0%
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/service/kvstoreservice/get.go:8:			Get			100.0%
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/service/kvstoreservice/list.go:7:			List			100.0%
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/service/kvstoreservice/set.go:8:			Set			100.0%
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/service/kvstoreservice/update.go:8:			Update			100.0%
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/storage/memory/kvstorage/base.go:30:			WithMemoryDB		100.0%
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/storage/memory/kvstorage/base.go:37:			New			100.0%
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/storage/memory/kvstorage/delete.go:3:			Delete			100.0%
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/storage/memory/kvstorage/get.go:9:			Get			100.0%
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/storage/memory/kvstorage/list.go:3:			List			100.0%
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/storage/memory/kvstorage/set.go:9:			Set			100.0%
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/storage/memory/kvstorage/update.go:3:			Update			100.0%
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/transport/http/basehttphandler/basehttphandler.go:18:	JSON			71.4%
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/transport/http/kvstorehandler/base.go:33:		WithService		100.0%
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/transport/http/kvstorehandler/base.go:40:		WithContextTimeout	100.0%
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/transport/http/kvstorehandler/base.go:47:		WithServerEnv		0.0%
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/transport/http/kvstorehandler/base.go:54:		WithLogger		100.0%
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/transport/http/kvstorehandler/base.go:61:		New			100.0%
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/transport/http/kvstorehandler/delete.go:11:		Delete			100.0%
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/transport/http/kvstorehandler/get.go:11:		Get			100.0%
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/transport/http/kvstorehandler/list.go:11:		List			100.0%
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/transport/http/kvstorehandler/set.go:14:		Set			100.0%
github.com/<GITHUB-KULLANICI-ADINIZ>/kvstore/src/internal/transport/http/kvstorehandler/update.go:14:		Update			100.0%
total:												(statements)		98.7%
```

yani tüm projenin toplam test coverage’ı **98.7%**. Sonra;

```bash
$ git add .
$ git commit -m 'add http handler tests'
```
