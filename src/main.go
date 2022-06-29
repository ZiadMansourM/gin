package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ziadmansourm/gin/config"
	"github.com/ziadmansourm/gin/routes"
)

func main() {
	f, err := os.OpenFile(config.LOG_FILE_PATH, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		panic(err.Error())
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())
	config.Connect()
	routes.UserRouter(engine)
	routes.BookRouter(engine)
	engine.Run(":3000")
}
