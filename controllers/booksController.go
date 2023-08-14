package controllers

import (
	"micros/models"
	"micros/repositories"
	"micros/utils"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func CreateBook(c *gin.Context) {
	var request models.Book

	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(400, utils.ResError(err.Error()))
		return
	}
	_, err = repositories.CreateBook(request)

	if err != nil {
		c.JSON(400, utils.ResError(err.Error()))
		return
	}

	c.JSON(200, utils.ResSuccess("OK"))
}

func GetAllBooks(c *gin.Context) {
	results, err := repositories.GetBooks()
	if err != nil {
		c.JSON(400, utils.ResError(err.Error()))
		return
	}
	c.JSON(200, utils.ResSuccess(results))
}

func FindBook(c *gin.Context) {
	var request models.Book
	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(400, utils.ResError("Invalid JSON Request"))
		return
	}

	result, err := repositories.FindBook(int(request.ID))

	if err != nil {
		c.JSON(400, utils.ResError(err.Error()))
		return
	}

	c.JSON(200, utils.ResSuccess(result))
}

func UpdateBook(c *gin.Context) {
	var request models.Book

	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(400, utils.ResError("Invalid JSON Request"))
		return
	}

	_, err = repositories.UpdateBook(request)

	if err != nil {
		c.JSON(400, utils.ResError(err.Error()))
		return
	}

	c.JSON(200, utils.ResSuccess("OK"))
}

func DeleteBook(c *gin.Context) {
	var request models.Book

	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(400, utils.ResError("Invalid JSON Request"))
		return
	}

	_, err = repositories.DeleteBook(request)

	if err != nil {
		c.JSON(400, utils.ResError(err.Error()))
		return
	}

	c.JSON(200, utils.ResSuccess("OK"))

}
