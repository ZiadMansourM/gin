package books

import (
	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) {
	router := engine.Group("/books")

	router.GET("/", List)
	router.POST("/", Create)

	router.GET("/:id", Get)
	router.PUT("/:id", Update)
	router.DELETE("/:id", Delete)
}
