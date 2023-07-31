# Bölüm 01/03: Kurulum

Dilin son sürümünü sitesinden [indirip][01] kurabilirsiniz. Kullandığınız işletim
sisteminin paket yöneticileri de kolayca kurmanıza yardımcı olur.

Kurulum yapmadan web üzerinden denemeler yapmak isterseniz;

https://go.dev/play/

`brew` paket yöneticisi kullanıyorsanız;

```bash
$ brew install go
$ go version
$ go env
$ cd "$(go env GOROOT)/src"    # source code
```

[01]: https://go.dev/dl/