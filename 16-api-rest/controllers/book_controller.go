package controllers

import (
	"strconv"

	"github.com/JPaulo-Moura/16-api-rest/database"
	"github.com/JPaulo-Moura/16-api-rest/models"
	"github.com/gin-gonic/gin"
)

//Resposta a rota GET api/v1/books
func ShowBook(ctx *gin.Context) {
	id := ctx.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, gin.H{ //response
			"error:": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	book := models.Book{}
	err = db.First(&book, idInt).Error

	if err != nil {
		ctx.JSON(400, gin.H{
			"error:": "Cannot find book: " + err.Error(),
		})
		return
	}

	ctx.JSON(200, book)
}

func CreateBook(ctx *gin.Context) {
	book := models.Book{}

	err := ctx.ShouldBindJSON(&book) //converte o body da request para a struct Book
	if err != nil {
		ctx.JSON(400, gin.H{
			"error:": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	db := database.GetDatabase()
	err = db.Create(&book).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"error:": "Cannot create book: " + err.Error(),
		})
		return
	}

	ctx.JSON(201, book)
}

func ShowBooks(ctx *gin.Context) {
	db := database.GetDatabase()

	books := []models.Book{}

	err := db.Find(&books).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"error:": "Cannot list books: " + err.Error(),
		})
		return
	}

	ctx.JSON(200, books)
}

func UpdateBook(ctx *gin.Context) {
	book := models.Book{}
	err := ctx.ShouldBindJSON(&book)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error:": "cannot bind JSON" + err.Error(),
		})
		return
	}

	db := database.GetDatabase()

	err = db.Save(&book).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "cannot create book: " + err.Error(),
		})
		return
	}

	ctx.JSON(200, book)
}

func DeleteBook(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "ID has to be integer" + err.Error(),
		})
		return
	}

	db := database.GetDatabase()
	err = db.Delete(&models.Book{}, id).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "cannot delete book: " + err.Error(),
		})
		return
	}

	ctx.Status(204)
}
