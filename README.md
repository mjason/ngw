ngw
===

一个简单的web framework 用来开发web api，不适合用来开发web site。

## install:

- go get labix.org/v2/mgo
- go get github.com/gorilla/mux
- go get github.com/mjason/ngw

## 例子:

```go
package main

import (
  "ngw"
)

func main() {
  // 添加一个路由
  ngw.Route("/", func(a ngw.Action) {
    a.Render([]byte(`你好吗？`))
  })
  // 设置监听端口，如果不设置使用的是127.0.0.1:3000
  ngw.Listen = "127.0.0.1:5000"
  // 启动服务器
  ngw.Start()
}

```
使用浏览器打开http://127.0.0.1:5000, 即可看到一端文字

===

### 版本

目前版本为0.1，很多东西都不稳定你懂的。

===

### 协议

采用BSD开源协议



