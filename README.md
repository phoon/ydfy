# 有道翻译库

有道翻译网页版爬虫，仅支持中译英及英译中，且不处理输出结果

## 使用

```bash
go get github.com/iPeven/ydfy
```

## 示例

```go
package main

import (
	"fmt"
	"github.com/iPeven/ydfy"
)

func main() {
	/*英译中*/
	fmt.Println(ydfy.En2zhCHS("hello"))
	/*中译英*/
	fmt.Println(ydfy.ZhCHS2en("你好"))
}
```

输出：

```bash
# 英译中
{"translateResult":[[{"tgt":"你好","src":"hello"}]],"errorCode":0,"type":"en2zh-CHS","smartResult":{"entries":["","n. 表示问候， 惊奇或唤起注意时的用语\r\n","int. 喂；哈罗\r\n","n. (Hello)人名；(法)埃洛\r\n"],"type":1}}
# 中译英
{"translateResult":[[{"tgt":"hello","src":"你好"}]],"errorCode":0,"type":"zh-CHS2en"}
```