package main

import (
	"fmt"
	"server1/output"
	"server1/settings"
	"testing"
)

func TestSettingInit(t *testing.T) {
	config := settings.Config

	fmt.Println(output.Info()+"Port: ", config.Port)
	fmt.Println("----------------------------------------")
	for _, route := range config.Routes {
		fmt.Println(output.Info()+"Route Name: ", route.Name)
		fmt.Println(output.Info()+"Route Url: ", route.Url)
		fmt.Println(output.Info()+"Route Jwt: ", route.Jwt)
		for _, path := range route.Paths {
			fmt.Println(output.Info()+"Route Path: ", path.Path, " Weight: ", path.Weight)
		}
		fmt.Println("----------------------------------------")
	}
}
