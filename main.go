package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(200, "OK")
	})
	router.GET("/", get)
	router.Run(":3000")
}
