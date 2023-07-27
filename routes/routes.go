package routes

import (
	"github.com/AndrewSalko/salkodev.edms.go/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(routes *gin.Engine) {

	routes.POST("users/register", controller.Register)
	routes.POST("users/login", controller.Login)
	routes.POST("users/changepassword", AuthMiddleware(), controller.ChangePassword)

}
