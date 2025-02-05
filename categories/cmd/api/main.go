package main


import "github.com/gin-gonic/gin"


func main() {

	r := gin.Default()

	r.GET("/healthy", func (c *gin.Context) {
		c.JSON(200, gin.H{
			"success": false,
		})
	}) 
	r.Run(":8080")
}