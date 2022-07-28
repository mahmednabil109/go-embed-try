//go:generate esc -prefix ./ui/build/ -pkg main -o static.go  ./ui/build
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.StaticFS("/", FS(false))
	r.Run(":8000")
}
