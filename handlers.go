package main

import "github.com/gin-gonic/gin"

func get(httpContext *gin.Context) {
	httpContext.JSON(200, "OK")
}

func realizarTransacao(httpContext *gin.Context) {
	resultado := ResultadoTransacao{
		Limite: "10000",
		Saldo:  "-9999",
	}
	httpContext.JSON(200, resultado)
}
