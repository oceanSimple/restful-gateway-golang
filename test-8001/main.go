package main

import "github.com/gin-gonic/gin"

const prefix = "/test-p1"

func main() {
	g := gin.Default()

	g.GET(prefix+"/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "8001-pong",
		})
	})

	g.Run(":8001")
}
