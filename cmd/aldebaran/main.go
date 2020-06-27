package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/irvanherz/aldebaran/models"
)

func main() {
	print("ALDEBARAN v1")
	print("API Server is starting...")
	route := gin.Default()
	models.Testee()

	route.GET("/go", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "0",
		})
	})
	route.Run(":3001")
	print("API Server is exiting...")
}
