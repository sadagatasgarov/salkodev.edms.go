package routes

import (
	"github.com/AndrewSalko/salkodev.edms.go/controller_departments"
	"github.com/AndrewSalko/salkodev.edms.go/controller_orgs"
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
	routes.POST("users/modify", AuthMiddleware(), controller_users.ModifyUser)
	routes.DELETE("users/delete", AuthMiddleware(), controller_users.DeleteUser)

	routes.POST("users/groups/add", AuthMiddleware(), controller_users.AddToGroup)
	routes.POST("users/groups/remove", AuthMiddleware(), controller_users.RemoveFromGroup)

	routes.POST("orgs/create", AuthMiddleware(), controller_orgs.CreateOrganization)
	routes.POST("orgs/modify", AuthMiddleware(), controller_orgs.ModifyOrganization)
	routes.DELETE("orgs/delete", AuthMiddleware(), controller_orgs.DeleteOrganization)

	routes.POST("departments/create", AuthMiddleware(), controller_departments.CreateDepartment)
	routes.POST("departments/modify", AuthMiddleware(), controller_departments.ModifyDepartment)
	routes.DELETE("departments/delete", AuthMiddleware(), controller_departments.DeleteDepartment)

}
