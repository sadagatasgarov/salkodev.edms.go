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
	routes.POST("users/create", AuthMiddleware(), controller_users.CreateUser)

	routes.POST("users/groups/add", AuthMiddleware(), controller_users.AddToGroup)
	routes.POST("users/groups/remove", AuthMiddleware(), controller_users.RemoveFromGroup)
}
