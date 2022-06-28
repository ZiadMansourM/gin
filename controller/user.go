package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ziadmansourm/gin/config"
	"github.com/ziadmansourm/gin/models"
)

func ListUsers(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)
}

func GetUser(c *gin.Context) {
	var user models.User
	config.DB.First(&user, c.Param("id"))
	c.JSON(200, &user)
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(200, &user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(200, &user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(200, &user)
}
