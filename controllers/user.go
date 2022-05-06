package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"dummy_api/models"	
	"dummy_api/config"
	// "fmt"
	
)

// func main (){
// 	dsn := "root:@tcp(127.0.0.1:3306)/people?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		panic("failed to connect config")
// 	}

// 	db.AutoMigrate(&models.User{})

// }

func RootMain(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
				"message": "Hello World",
				"status" : http.StatusOK,
			})
}

func UserRoute(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
					"name" : "Gohan Gin Goku",
					"age" : "20",
					"power" : "4000K",
					"tier" : "B",
					"abilities" : []string{"Super Saiyan", "Super Saiyan 2", "Super Saiyan 3", "Super Saiyan 4"},
					"image" : "https://images5.alphacoders.com/610/610758.png",
				})
}

func GetUserById(c *gin.Context){
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id" : id,
	})
}

type User struct {
	Name string `json:"name" binding:"required"`
	Age int `json:"age"`
	Power int `json:"power"`

}

func UserPostRoute(c *gin.Context){
	var user User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name" : user.Name,
		"age" : user.Age,
		"power" : user.Power,
	})
}

func CreateUserController(c *gin.Context ){
	var input User
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}

	// db := config.InitDb()	
	user  := models.User{
		Name : input.Name,
		Age : input.Age,
		Power : input.Power,
	}

	config.DB.Create(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
				"errors" : err,
			})
	}	
	c.JSON(http.StatusOK, gin.H{
		"name" : user.Name,
		"age" : user.Age,
		"power" : user.Power,
	})


}