package main

import (
	"fmt"

	"github.com/AndrewSalko/salkodev.edms.go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("SalkoDev.EDMS Go")

	router := gin.New()
	routes.UserRoutes(router)
	router.Run(":8080")

	fmt.Println("Finished")
}
