package routes

import (
	"github.com/AndrewSalko/salkodev.edms.go/controller"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.POST("users/register", controller.Register)
	incomingRoutes.POST("users/login", controller.Login)
	incomingRoutes.POST("users/changepassword", controller.ChangePassword)

}
