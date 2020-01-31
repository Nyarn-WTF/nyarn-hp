package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("pages/*.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})

	router.GET("/test", func(ctx *gin.Context) {
		ctx.HTML(200, "test.html", gin.H{})
	})

	router.Run()
}
