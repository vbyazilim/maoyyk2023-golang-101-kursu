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

Linux için;

https://go.dev/doc/install

Önce kaynak kodu [indirin][01], sonra;

```bash
$ rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.6.linux-amd64.tar.gz

# $HOME/.profile
$ export PATH=$PATH:/usr/local/go/bin
$ go version
$ go env
```

[01]: https://go.dev/dl/
