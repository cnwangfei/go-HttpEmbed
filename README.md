# go-HttpEmbed
取消http中使用embed嵌入静态资源时必须录入相对路径的问题

静态资源目录
```
D:\WORK\go\src\work\test\ui
D:\WORK\go\src\work\test\ui\page
D:\WORK\go\src\work\test\ui\page\index.html
D:\WORK\go\src\work\test\ui\page\main.wasm
D:\WORK\go\src\work\test\ui\page\wasm.js
```
http://127.0.0.1:8080/ui/index.html 打开index.html
```
package main

import (
	"embed"
	httpEmbed "github.com/cnwangfei/go-HttpEmbed"
	"github.com/gin-gonic/gin"
	"net/http"
)

//go:embed static/ui/page
var uiEmbed embed.FS

func main() {
	g := gin.Default()
	g.GET("")

	var e httpEmbed.Fs
	e.FS = &uiEmbed
	e.Path = "static/ui/page"
	g.StaticFS("/ui", http.FS(e))

	g.Run(":8080")
}
```
