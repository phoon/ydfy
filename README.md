# 有道翻译库

有道翻译网页版爬虫，仅支持中译英及英译中，且不处理输出结果。

## 使用

```bash
go get github.com/phoon/ydfy
```

## 示例

> 假设我们编译出的 二进制文件叫做`fy`

```go
// go build -trimpath -o fy
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"strings"
	"github.com/phoon/ydfy"
)

var (
	c string //中文
	e string //英文
	s bool   //smart
)

func init() {
	flag.StringVar(&c, "c", "", "-c 待翻译的中文")
	flag.StringVar(&e, "e", "", "-e English to translate")
	flag.BoolVar(&s, "s", false, "-s 显示smart result")
}

func main() {
	flag.Parse()
	r := ydfy.Result{}

	switch c {
	case "":
		if e == "" {
			flag.Usage()
		} else {
			// English to Chinese
			json.Unmarshal([]byte(ydfy.En2zhCHS(strings.TrimSpace(e))), &r)
			fmt.Println(r.TranslateResult[0][0].Tgt)
		}
	default:
		// Chinese to English
		json.Unmarshal([]byte(ydfy.ZhCHS2en(strings.TrimSpace(c))), &r)
		fmt.Println(r.TranslateResult[0][0].Tgt)
	}

	// enable smart result?
	if s {
		if len(r.SmartResult.Entries) != 0 {
			fmt.Println("\nMore:")
			for _, v := range r.SmartResult.Entries {
				fmt.Println("\t", v)
			}
		}
	}
}
```

输出：

```bash
# -s 参数开启SmartResult（即有道词典结果
#英译中
$ fy -s -e "development"
发展

More:
         n. 发展；开发；发育；住宅小区（专指由同一开发商开发的）；[摄] 显影

$ fy -e "development" #不开启 -s
发展

# 中译英
$ fy -s -c "发展"
The development of

More:
         develop

         progress

         evolution
```
