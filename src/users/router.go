package users

import (
	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) {
	router := engine.Group("/users")

	router.GET("/", List)
	router.POST("/", Create)

	router.GET("/:id", Get)
	router.PUT("/:id", Update)
	router.DELETE("/:id", Delete)
}
