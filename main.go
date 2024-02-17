package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", get)
	router.POST("/clientes/:id/transacoes", realizarTransacao)
	router.Run(":3000")
}
