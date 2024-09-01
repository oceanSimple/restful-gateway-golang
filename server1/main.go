package main

import (
	"fmt"
	"server1/settings"
)

func main() {
	fmt.Println(settings.Config)

	//engine := gin.Default()
	//
	//engine.GET("/ping", func(c *gin.Context) {
	//})
	//
	//fmt.Println(output.Info() + "Server is running on port 8080")
	//err := engine.Run(":8080")
	//if err != nil {
	//	fmt.Println(output.Error() + "Failed to start server")
	//	return
	//}
}
