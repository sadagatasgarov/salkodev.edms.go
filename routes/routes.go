package routes

import (
	"github.com/AndrewSalko/salkodev.edms.go/controller_users"
	"github.com/gin-gonic/gin"
)

func InitRoutes(routes *gin.Engine) {

	routes.POST("users/register", controller_users.Register)
	routes.POST("users/login", controller_users.Login)
	routes.GET("users/confirmregistration", controller_users.ConfirmRegistration)
	routes.POST("users/changepassword", AuthMiddleware(), controller_users.ChangePassword)

	routes.POST("users/refreshtoken", AuthMiddleware(), controller_users.RefreshToken)
}
