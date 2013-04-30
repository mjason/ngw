ngw
===

ngw是一个简单的web开发框架，十分适合用来开发web api项目，支持restful 路由控制，它包括了一个基于github.com/gorilla/mux的一个更加简单强大的路由。 包括一个mgo的两个增强函数。

===
## 如何使用

一个简单的hello word

```
package main

import (
  "github.com/mjason/ngw"
)

func main() {
  ngw.Get("/", SayHello)
  ngw.Start()
}

func SayHello(a ngw.A) {
  a.OK([]byte(`hello word`))
}

```

### todo

完善文档

### 协议

采用BSD开源协议