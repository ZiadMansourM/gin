package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ziadmansourm/gin/config"
	"github.com/ziadmansourm/gin/users"
)

func CreateResponseBook(book BookModel, user users.UserSerializer) BookSerializer {
	return BookSerializer{Id: book.Id, Title: book.Title, Author: user}
}

func List(c *gin.Context) {
	books := []BookModel{}
	config.DB.Find(&books)
	responseBooks := []BookSerializer{}
	for _, book := range books {
		var user users.UserModel
		config.DB.First(&user, book.PersonID)
		responseBook := CreateResponseBook(book, users.CreateResponseUser(user))
		responseBooks = append(responseBooks, responseBook)
	}
	c.JSON(http.StatusOK, &responseBooks)
}

func Get(c *gin.Context) {
	var book BookModel
	if err := config.DB.First(&book, c.Param("id")).Error; err != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{"errors": err.Error()},
		)
		return
	}
	var user users.UserModel
	config.DB.First(&user, book.PersonID)
	responseBook := CreateResponseBook(book, users.CreateResponseUser(user))
	c.JSON(http.StatusOK, &responseBook)
}

func Create(c *gin.Context) {
	var book BookModel
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	var user users.UserModel
	if err := config.DB.First(&user, book.PersonID).Error; err != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{"errors": err.Error()},
		)
		return
	}
	config.DB.Create(&book)

	responseBook := CreateResponseBook(book, users.CreateResponseUser(user))
	c.JSON(http.StatusOK, &responseBook)
}

func Update(c *gin.Context) {
	var book BookModel
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

func Delete(c *gin.Context) {
	var book BookModel
	config.DB.Where("id = ?", c.Param("id")).Delete(&book)
	c.JSON(http.StatusOK, &book)
}
