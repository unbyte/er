# ER - re regexp

![version](https://img.shields.io/github/v/tag/unbyte/er?style=flat-square)
![license](https://img.shields.io/github/license/unbyte/er?style=flat-square)
![Go Report Card](https://goreportcard.com/badge/github.com/unbyte/er?style=flat-square)

## 🛠 Installation

### As Executable 

- If you have `go` in your system, what you need to do is just to use `go get`

    ```shell script
    go get -u github.com/unbyte/er/cmd/er
    er --help
    ```

- In other case, you need go to download executable in [Release Page](https://github.com/unbyte/er/releases), 
and then rename the executable to `er` ( or `er.exe` on window ). 

### As Package

If you want to use `er` as a go package, please `go get -u github.com/unbyte/er`.

## 🎨 Usage

### As Executable

```shell script
> er --help
Usage of ER CLI:
  -a, -amount int
        amount of strings to be generated. default to 1. (default 1)
  -p, -pattern string
        pattern string

Syntax: https://golang.org/pkg/regexp/syntax/
Unicode Class: https://en.wikipedia.org/wiki/Unicode_character_property

> er -p "\d{3}-\d{8}|\d{4}-\d{7}" -a 10
454-16390004
0913-1976506
543-75125853
280-27961072
5049-7522609
7833-1752530
822-89737417
7176-8019427
5181-6167904
090-60481568

> er -p "[😂-😍]+"
😇😅😋😊😃😉😊😂😆😍😊😈😍😇😍😄😇😍😂😆😉😌😅😅😊😃😃😇😂😌😉😋😉😆😆😂😄😍😂😍😇😄😃😄😂😈😂😌😉😌😅😇😂

```

You can get the syntax of regexp [here](https://golang.org/pkg/regexp/syntax/)
 
and get all unicode groups [here](https://en.wikipedia.org/wiki/Unicode_character_property)

### As Package

```go
import (
    "github.com/unbyte/er"
    "regexp/syntax"
)
func main(){
    pattern := "^[A-Za-z_]{10,14}$"
    generator, err := er.Parse(pattern, syntax.Perl)
    if err != nil {
        panic(err)
    }

    s, err := generator.Generate()
    if err != nil {
        panic(err)
    }
    fmt.Println(s)

    ss, err := generator.GenerateMultiple(10)
    if err != nil {
        panic(err)
    }
    fmt.Println(strings.Join(ss, "\n"))
}
```

## 📄 License

MIT LICENSE.