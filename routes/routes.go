package routes

import (
	controller "github.com/AndrewSalko/salkodev.edms.go/controller/users"
	"github.com/gin-gonic/gin"
)

func InitRoutes(routes *gin.Engine) {

	routes.POST("users/register", controller.Register)
	routes.POST("users/login", controller.Login)
	routes.GET("users/confirmregistration", controller.ConfirmRegistration)
	routes.POST("users/changepassword", AuthMiddleware(), controller.ChangePassword)

	routes.POST("users/refreshtoken", AuthMiddleware(), controller.RefreshToken)
}
