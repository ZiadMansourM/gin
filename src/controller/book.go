package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ziadmansourm/gin/config"
	"github.com/ziadmansourm/gin/models"
)

type Book struct {
	Id     uint64      `json:"id"`
	Title  string      `json:"title"`
	Author models.User `json:"author" gorm:"foreignkey:PersonID"`
}

func CreateResponseBook(book models.Book, user models.User) Book {
	return Book{Id: book.Id, Title: book.Title, Author: user}
}

func ListBooks(c *gin.Context) {
	books := []models.Book{}
	config.DB.Find(&books)
	responseBooks := []Book{}
	for _, book := range books {
		var user models.User
		config.DB.First(&user, book.PersonID)
		responseBook := CreateResponseBook(book, user)
		responseBooks = append(responseBooks, responseBook)
	}
	c.JSON(http.StatusOK, &responseBooks)
}

func GetBook(c *gin.Context) {
	var book models.Book
	if err := config.DB.First(&book, c.Param("id")).Error; err != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{"errors": err.Error()},
		)
		return
	}
	var user models.User
	config.DB.First(&user, book.PersonID)
	responseBook := CreateResponseBook(book, user)
	c.JSON(http.StatusOK, &responseBook)
}

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	var user models.User
	if err := config.DB.First(&user, book.PersonID).Error; err != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{"errors": err.Error()},
		)
		return
	}
	config.DB.Create(&book)

	responseBook := CreateResponseBook(book, user)
	c.JSON(http.StatusOK, &responseBook)
}

func UpdateBook(c *gin.Context) {
	var book models.Book
	if err := config.DB.First(&book, c.Param("id")).Error; err != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{"errors": err.Error()},
		)
		return
	}
	type UpdateBook struct {
		Title string `json:"title"`
	}
	var updatedData UpdateBook
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}
	book.Title = updatedData.Title
	config.DB.Save(&book)
	c.JSON(http.StatusOK, &book)
}

func DeleteBook(c *gin.Context) {
	var book models.Book
	config.DB.Where("id = ?", c.Param("id")).Delete(&book)
	c.JSON(http.StatusOK, &book)
}
