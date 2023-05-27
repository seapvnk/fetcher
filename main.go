package main

import (
	f "github.com/seapvnk/fetcher/lib"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/", f.FetchUrlHandler)
	r.Run()
}