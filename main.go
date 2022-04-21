package main 
import "github.com/gin-gonic/gin"
import "net/http"

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")


	v1.GET("/", rootMain)
	v1.GET("/user", userRoute)
	v1.POST("/user", userPostRoute)
	router.Run()
}


func rootMain(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
				"message": "Hello World",
				"status" : http.StatusOK,
			})
}

func userRoute(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
					"name" : "Gohan Gin Goku",
					"age" : "20",
					"power" : "4000K",
					"tier" : "B",
					"abilities" : []string{"Super Saiyan", "Super Saiyan 2", "Super Saiyan 3", "Super Saiyan 4"},
					"image" : "https://images5.alphacoders.com/610/610758.png",
				})
}

func getUserById(c *gin.Context){
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

func userPostRoute(c *gin.Context){
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