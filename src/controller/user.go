package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ziadmansourm/gin/config"
	"github.com/ziadmansourm/gin/models"
)

func ListUsers(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(http.StatusOK, &users)
}

func GetUser(c *gin.Context) {
	var user models.User
	config.DB.First(&user, c.Param("id"))
	c.JSON(http.StatusOK, &user)
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(http.StatusOK, &user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(http.StatusOK, &user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(http.StatusOK, &user)
}
