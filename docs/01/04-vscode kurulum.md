# Bölüm 01/04: Vscode Kurulum

Go için geliştirilmiş birçok IDE var. Bunlardan biri de Visual Studio Code.
Kurulumu için;

https://code.visualstudio.com/download adresinden işletim sisteminize uygun olanı indirmeniz gerekiyor.


Debian ve Ubuntu için;


```bash
$ sudo apt install ./<file>.deb
```

Daha eski linux dağıtımları için;

```bash
$ sudo dpkg -i <file>.deb
$ sudo apt-get install -f # Install dependencies
```

https://code.visualstudio.com/docs/setup/linux

Kurulum tamamlandıktan sonra vscode go ekletisini kurmanız gerekiyor.

https://marketplace.visualstudio.com/items?itemName=golang.go

Go eklentisi kurulumu tamamlandıktan sonra vscode'da go ile gerekli ayarları yapmak için user settings'de değişiklikler yapmamız gerekiyor. Bunun için;

```bash
View -> Command Palette -> Open User Settings
```

Settings sayfasında aşağıdaki ayarları yapmanız gerekiyor.
```
 "[go]": {
    "editor.insertSpaces": false,
    "editor.formatOnSave": true,
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    }
  }, 
  "go.lintOnSave": "workspace",
  "go.lintTool": "golangci-lint",
```

