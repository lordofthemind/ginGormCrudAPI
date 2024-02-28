package main

import "github.com/gin-gonic/gin"

func main() {

	routes := gin.Default()

	routes.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routes.Run(":9090")

}
