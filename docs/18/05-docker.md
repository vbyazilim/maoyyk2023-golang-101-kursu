# Bölüm 18/05: In-Memory Key-Value Store

## Docker

Şimdi bu go servisimizi `docker` ile bir container içinden çalıştıralım.
Öncelikle `Dockerfile` oluşturalım:

```bash
$ touch Dockerfile
```

sonra;

```docker
# build application
FROM golang:1.21.0-alpine AS builder

ENV GOPRIVATE=github.com/vbyazilim

ARG GITHUB_ACCESS_TOKEN
ARG BUILD_INFORMATION

# hadolint ignore=DL3018
RUN apk add --update --no-cache git \
    && git config --global url.https://${GITHUB_ACCESS_TOKEN}@github.com/.insteadOf https://github.com/

WORKDIR /build
COPY ./go.mod /build/

# COPY ./go.mod ./go.sum /build/
# RUN go mod download

COPY . /build
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-X 'github.com/vbyazilim/kvstore/src/releaseinfo.BuildInformation=${BUILD_INFORMATION}'" -o app ./cmd/server

# get certificates
FROM alpine:3.18.3 AS certs

# hadolint ignore=DL3018
RUN apk add --update --no-cache ca-certificates

FROM busybox:1.36
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /build/app /kvstoreapp

EXPOSE 8000
CMD ["/kvstoreapp"]
```

Hemen deneyelim oldu mu?

```bash
$ export BUILD_INFO="$(git rev-parse HEAD)-$(go env GOOS)-$(go env GOARCH)"
$ echo $BUILD_INFO
$ docker build --build-arg="BUILD_INFORMATION=${BUILD_INFO}" -t kvstore:latest .
```

şimdi çalıştıralım;

```bash
$ SERVER_ENV="production" LOG_LEVEL="ERROR" docker run --cpus="2" --env SERVER_ENV --env LOG_LEVEL -p 9000:8000 kvstore:latest
```

Şimdi service `:9000`’den erişelim?

```bash
$ http POST "http://localhost:9000/api/v1/set" key="success" value:=true
$ http POST "http://localhost:9000/api/v1/set" key="server_env" value="production"
$ http "http://localhost:9000/api/v1/list"
$ http PUT "http://localhost:9000/api/v1/update" key="success" value:=false
$ http "http://localhost:9000/api/v1/get?key=success"
$ http DELETE "http://localhost:9000/api/v1/delete?key=success"
```

Tamamsa;

```bash
$ git add Dockerfile
$ git commit -m 'add Dockerfile'
```

---
