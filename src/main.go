package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ziadmansourm/gin/config"
	"github.com/ziadmansourm/gin/routes"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRouter(router)
	router.Run(":3000")
}
