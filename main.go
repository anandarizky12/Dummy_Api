package main

import (
	"dummy_api/config"
	"dummy_api/controllers"

	"github.com/gin-gonic/gin"
	// "dummy_api/models"
	// "fmt"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")

	config.ConnectDb()

	// v1.GET("/", controllers.RootMain)
	// v1.GET("/user", controllers.UserRoute)\
	v1.GET("/user/:id", controllers.GetUserById)
	v1.GET("/user", controllers.GetAllUserstRoute)
	v1.POST("/create", controllers.CreateUserController)
	v1.DELETE("/delete/:id", controllers.DeleteUserController)
	v1.PATCH("/update/:id", controllers.UpdateUserController)
	router.Run()
}
