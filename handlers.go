package main

import "github.com/gin-gonic/gin"

func get(httpContext *gin.Context) {
	httpContext.JSON(200, "OK")
}
