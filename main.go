package main 
import (
	"github.com/gin-gonic/gin"
	"dummy_api/controllers"
	"dummy_api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"fmt"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	dsn := "root:@tcp(127.0.0.1:3306)/people?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	
	fmt.Println("Connected to database")
	db.AutoMigrate(&models.User{})

	v1.GET("/", controllers.RootMain)
	v1.GET("/user", controllers.UserRoute)
	v1.POST("/user", controllers.UserPostRoute)
	router.Run()
}

