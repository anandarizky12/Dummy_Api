package main 
import (
	"github.com/gin-gonic/gin"
	"dummy_api/controllers"
		
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")


	v1.GET("/", controllers.RootMain)
	v1.GET("/user", controllers.UserRoute)
	v1.POST("/user", controllers.UserPostRoute)
	router.Run()
}

