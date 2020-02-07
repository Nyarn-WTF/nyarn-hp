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

	router.GET("/contact.html", func(ctx *gin.Context) {
		ctx.HTML(200, "contact.html", gin.H{})
	})

	router.GET("/products.html", func(ctx *gin.Context) {
		ctx.HTML(200, "products.html", gin.H{})
	})

	router.Run(":443")
}
