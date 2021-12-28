# 简单的中文昵称生成


## 如何使用
```sh
# 准备
go get github.com/cuipeiyu/go-makename
```

```go
package main

import (
	"log"

	"github.com/cuipeiyu/go-makename"
)

func main() {
	nickname := makename.Generate()
	log.Println(nickname)
}

```
