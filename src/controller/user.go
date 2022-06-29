package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ziadmansourm/gin/config"
	"github.com/ziadmansourm/gin/models"
)

func ListUsers(c *gin.Context) {
	users := []models.User{}
	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(
			http.StatusConflict,
			gin.H{"errors": err.Error()},
		)
		return
	}
	c.JSON(http.StatusOK, &users)
}

func GetUser(c *gin.Context) {
	var user models.User
	if err := config.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{"errors": err.Error()},
		)
		return
	}
	c.JSON(http.StatusOK, &user)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(
			http.StatusConflict,
			gin.H{"errors": err.Error()},
		)
		return
	}
	c.JSON(http.StatusOK, &user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	if err := config.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{"errors": err.Error()},
		)
		return
	}
	type UpdateUser struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var updatedData UpdateUser
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}
	user.Name = updatedData.Name
	user.Email = updatedData.Email
	user.Password = updatedData.Password
	config.DB.Save(&user)
	c.JSON(http.StatusOK, &user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	if err := config.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{"errors": err.Error()},
		)
		return
	}
	if err := config.DB.Delete(&user).Error; err != nil {
		c.JSON(
			http.StatusConflict,
			gin.H{"errors": err.Error()},
		)
		return
	}
	c.JSON(http.StatusOK, &user)
}
