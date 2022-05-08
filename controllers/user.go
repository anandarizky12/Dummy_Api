package controllers

import (
	"dummy_api/config"
	"dummy_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootMain(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
		"status":  http.StatusOK,
	})
}

// func UserRoute(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"name":      "Gohan Gin Goku",
// 		"age":       "20",
// 		"power":     "4000K",
// 		"tier":      "B",
// 		"abilities": []string{"Super Saiyan", "Super Saiyan 2", "Super Saiyan 3", "Super Saiyan 4"},
// 		"image":     "https://images5.alphacoders.com/610/610758.png",
// 	})
// }

type User struct {
	Name  string `json:"name" binding:"required"`
	Age   int    `json:"age"`
	Power int    `json:"power"`
}

func GetAllUserstRoute(c *gin.Context) {
	var users []models.User
	err := config.DB.Find(&users).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func GetUserById(c *gin.Context) {
	var user models.User

	id := c.Param("id")
	err := config.DB.Where("id = ?", id).First(&user).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Data Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func CreateUserController(c *gin.Context) {
	var input models.User
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// db := config InitDb()

	err = config.DB.Create(&input).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": input,
	})

}

func DeleteUserController(c *gin.Context) {

	var user models.User

	id := c.Param("id")

	err := config.DB.Where("id = ?", id).First(&user).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Data Not Found",
		})
		return
	}

	config.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"data":    user,
		"message": "Succesfully Deleted The data",
	})
}
