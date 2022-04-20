package main 
import "github.com/gin-gonic/gin"


func main() {
	router := gin.Default()
	var data = [4]string{"1","2","3","4"}

	router.GET("/", func(c *gin.Context){
			c.JSON(200, gin.H{
				"message": "Hello World",
				"data": data,
			})
	})

	router.Run()
}