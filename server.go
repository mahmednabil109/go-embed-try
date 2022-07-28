//go:generate go run /Users/mahmednabil/go/pkg/mod/github.com/mjibson/esc@v0.2.0/main.go -o static.go -pkg=main -prefix "ui/build" ./ui/build
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.StaticFS("/", FS(false))

	r.Run(":8000")
}
