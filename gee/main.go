package main

import "gee/gee"

func main() {
	engine := gee.New()
	engine.GET("/", func(c *gee.Context) {
		c.HTML(200, "<h1>Hello Gee</h1>")
	})
	engine.POST("/hello", func(c *gee.Context) {
		c.JSON(200, gee.H{
			"message": "hello",
		})
	})
	engine.Run(":9999")
}
