package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ziadmansourm/gin/books"
	"github.com/ziadmansourm/gin/config"
	"github.com/ziadmansourm/gin/users"
)

func main() {
	// [1]: Write to custome logger file
	f, err := os.OpenFile(config.LOG_FILE_PATH, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		panic(err.Error())
	}
	// [2]: Create DB Connection && Migrate changes
	config.Connect()
	config.DB.AutoMigrate(
		&users.UserModel{},
		&books.BookModel{},
	)
	// [3]: Craete gin Engine && middlewares
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())
	// [4]: Register routers
	users.Router(engine)
	books.Router(engine)
	// [5]: Engine.Run:3000
	engine.Run(config.PORT_NUM)
}
