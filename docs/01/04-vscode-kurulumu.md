# Bölüm 01/04: VSCode Kurulumu

Go için geliştirilmiş birçok IDE (Integrated Development Environment) var.
Bunlardan biri de **Visual Studio Code**. Kurulumu için;

https://code.visualstudio.com/download adresinden işletim sisteminize uygun
olanı indirmeniz gerekiyor.

[Debian ve Ubuntu][01] için;

```bash
$ sudo apt install ./<file>.deb
```

Daha eski Linux dağıtımları için;

```bash
$ sudo dpkg -i <file>.deb
$ sudo apt-get install -f # Install dependencies
```

Kurulum tamamlandıktan sonra vscode [go ekletisini][02] kurmanız gerekiyor. Go
eklentisi kurulumu tamamlandıktan sonra vscode’da `go` ile gerekli ayarları
yapmak için user `settings`’de değişiklikler yapmamız gerekiyor. Bunun için;

    View -> Command Palette -> Open User Settings

**Settings** sayfasında aşağıdaki ayarları yapmanız gerekiyor.

```json
 "[go]": {
    "editor.insertSpaces": false,
    "editor.formatOnSave": true,
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    }
  }, 
  "go.lintOnSave": "workspace",
  "go.lintTool": "golangci-lint",
}
```

[01]: https://code.visualstudio.com/docs/setup/linux
[02]: https://marketplace.visualstudio.com/items?itemName=golang.go
