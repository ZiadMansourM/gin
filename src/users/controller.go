package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ziadmansourm/gin/config"
)

func CreateResponseUser(user UserModel) UserSerializer {
	return UserSerializer{Id: user.Id, Name: user.Name, Email: user.Email}
}

func List(c *gin.Context) {
	users := []UserModel{}
	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(
			http.StatusConflict,
			gin.H{"errors": err.Error()},
		)
		return
	}
	responseUsers := []UserSerializer{}
	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}
	c.JSON(http.StatusOK, &responseUsers)
}

func Get(c *gin.Context) {
	var user UserModel
	if err := config.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{"errors": err.Error()},
		)
		return
	}
	responseUser := CreateResponseUser(user)
	c.JSON(http.StatusOK, &responseUser)
}

func Create(c *gin.Context) {
	var user UserModel
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
	responseUser := CreateResponseUser(user)
	c.JSON(http.StatusOK, &responseUser)
}

func Update(c *gin.Context) {
	var user UserModel
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
	responseUser := CreateResponseUser(user)
	c.JSON(http.StatusOK, &responseUser)
}

func Delete(c *gin.Context) {
	var user UserModel
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
	responseUser := CreateResponseUser(user)
	c.JSON(http.StatusOK, &responseUser)
}
