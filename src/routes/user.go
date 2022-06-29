package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ziadmansourm/gin/controller"
)

func UserRouter(engine *gin.Engine) {
	router := engine.Group("/users")

	router.GET("/", controller.ListUsers)
	router.POST("/", controller.CreateUser)

	router.GET("/:id", controller.GetUser)
	router.PUT("/:id", controller.UpdateUser)
	router.DELETE("/:id", controller.DeleteUser)
}
