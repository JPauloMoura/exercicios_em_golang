package controllers

import "github.com/gin-gonic/gin"

func ShowBooks(ctx *gin.Context) {
	//Resposta a rota GET api/v1/books
	ctx.JSON(200, gin.H{
		"value": "ok",
	})
}
