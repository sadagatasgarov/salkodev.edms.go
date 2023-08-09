package main

import (
	"context"
	"fmt"

	"github.com/AndrewSalko/salkodev.edms.go/database_groups"
	"github.com/AndrewSalko/salkodev.edms.go/database_users"
	"github.com/AndrewSalko/salkodev.edms.go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("SalkoDev.EDMS Go")

	database_groups.ValidateSchema()
	database_users.ValidateSchema()

	database_users.ValidateAdminAccount(context.TODO())

	router := gin.New()
	routes.InitRoutes(router)
	router.Run(":8080")

	fmt.Println("Finished")
}
