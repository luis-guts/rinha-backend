package main

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
)

func get(httpContext *gin.Context) {
	httpContext.JSON(200, "OK")
}

func realizarTransacao(httpContext *gin.Context) {

	idCliente := httpContext.Param("id")
	transacao := Transacao{}
	if err := httpContext.BindJSON(&transacao); err != nil {
		httpContext.Error(err)
		httpContext.AbortWithStatus(400)
		return
	}
	id, err := strconv.ParseInt(idCliente, 10, 64)
	if err != nil {
		httpContext.AbortWithStatus(400)
		return
	}
	cliente, err := salvarTransacao(id, &transacao)
	if err != nil {
		if err == sql.ErrNoRows {
			httpContext.Error(err)
			httpContext.AbortWithStatus(404)
			return
		}
		if err.Error() == ("saldo_insuficiente") {
			httpContext.Error(err)
			httpContext.AbortWithStatus(422)
			return
		}
		httpContext.Error(err)
		httpContext.AbortWithStatus(400)
		return
	}

	httpContext.JSON(200, cliente)
}

func getExtrato(httpContext *gin.Context) {
	resultado := ResultadoTransacao{
		Limite: "10000",
		Saldo:  "-9999",
	}
	httpContext.JSON(200, resultado)
}
