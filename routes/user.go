package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ziadmansourm/gin/controller"
)

func UserRouter(router *gin.Engine) {
	list_Create_PATH := "/user"
	router.GET(list_Create_PATH, controller.ListUsers)
	router.POST(list_Create_PATH, controller.CreateUser)
	get_Update_Delete_PATH := "/user/:id"
	router.GET(get_Update_Delete_PATH, controller.GetUser)
	router.PUT(get_Update_Delete_PATH, controller.UpdateUser)
	router.DELETE(get_Update_Delete_PATH, controller.DeleteUser)
}
