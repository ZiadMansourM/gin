package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ziadmansourm/gin/controller"
)

func BookRouter(engine *gin.Engine) {
	router := engine.Group("/books")

	router.GET("/", controller.ListBooks)
	router.POST("/", controller.CreateBook)

	router.GET("/:id", controller.GetBook)
	router.PUT("/:id", controller.UpdateBook)
	router.DELETE("/:id", controller.DeleteBook)
}
