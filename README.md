# go-hutool

go版的工具集，java的hutool，适合java体质的go开发者、go开发新手。

## 网站

官网：https://github.com/xcocean/go-hutool

## 安装

要求go版本>=1.18

```shell
go get -u github.com/xcocean/go-hutool
```

## 使用

```go
package main

// 导入依赖
import (
	"github.com/xcocean/go-hutool/gg"
)

// 使用
func main() {
	println(gg.IsWindows())          // true
	println(gg.File_exist("d://gg")) // true
	gg.Assert_isTure2(gg.IsWindows(), "当前系统不是Windows")

	// 字符串去空格
	println(gg.Str_Trim("  github.com/xcocean/go-hutool/gg "))
}
```
