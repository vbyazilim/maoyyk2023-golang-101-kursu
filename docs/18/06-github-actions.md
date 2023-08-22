# Bölüm 18/06: In-Memory Key-Value Store

## GitHub Actions

GitHub’a `push` yaptığımızda ya da `pull request`’leri `merge` ettiğimizde,
bir kısım kontrollerin yapılmasını istiyoruz. Otomatik olarak testler çalışsın,
linter kodu kontrol etsin:

```bash
$ mkdir -p .github/workflows
$ touch .github/workflows/go-{lint,test}.yml
```

`.github/workflows/go-lint.yml`

```yaml
name: Golang CI Lint

on:
  pull_request:

concurrency:
  group: golangci-lint
  cancel-in-progress: true

jobs:
  golangci:
    name: golangci linter
    runs-on: ubuntu-latest
    env:
      GOPRIVATE: github.com/vbyazilim
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
        with:
          go-version: '1.21'
          cache: false
      - name: golangci-lint
        uses: golangci/golangci-lint-action@v3
        with:
          version: v1.54
          args: --timeout=5m
```

---

`.github/workflows/go-test.yml`

```yaml
name: Golang Tests

on:
  pull_request:
  push:
    branches:
      - main
    tags-ignore:
      - '**'

concurrency:
  group: golang-test
  cancel-in-progress: true

jobs:
  test:
    name: Run tests
    runs-on: ubuntu-latest
    env:
      GOPRIVATE: github.com/vbyazilim
    steps:
      - name: Checkout code
        uses: actions/checkout@v3

      - name: Setup Go
        uses: actions/setup-go@v4
        with:
          go-version-file: "go.mod"
        id: go

      - run: git config --global url.https://${{ github.token }}@github.com/.insteadOf https://github.com/

      - name: Run tests
        run: LOG_LEVEL="error" go test -p 1 -v -race -failfast -coverprofile=coverage.txt -covermode=atomic ./...

      - name: Upload coverage to Codecov
        uses: codecov/codecov-action@v3
        with:
            token: ${{ secrets.CODECOV_TOKEN }}
```

sonra;

```bash
$ git add .
$ git commit -m 'add github action workflows'
```

Action’ları çalıştırmak için kodu GitHub’a push etmemiz lazım.
