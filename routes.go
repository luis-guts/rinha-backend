package main

import (
	"github.com/gin-gonic/gin"
)

func addRoutes(engine *gin.Engine) {
	engine.GET("/", get)
}
