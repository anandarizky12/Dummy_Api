package main 
import (
	"github.com/gin-gonic/gin"
	"dummy_api/controllers"
	// "dummy_api/models"
	"dummy_api/config"
	// "fmt"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	
	config.ConnectDb();

	v1.GET("/", controllers.RootMain)
	v1.GET("/user", controllers.UserRoute)
	v1.POST("/user", controllers.UserPostRoute)
	v1.POST("/create", controllers.CreateUserController)
	router.Run()
}

