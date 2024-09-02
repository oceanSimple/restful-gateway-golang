package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"server1/balance"
	"server1/output"
	"server1/settings"
	"strconv"
	"strings"
)

func main() {
	engine := gin.Default()

	routeForwarding(engine)

	fmt.Println(output.Info() + "Server is running on port 8080")

	err := engine.Run(":" + strconv.Itoa(int(settings.Config.Port)))
	if err != nil {
		fmt.Println(output.Error() + "Failed to start server")
		return
	}
}

func routeForwarding(engine *gin.Engine) {
	engine.Any("/*path", func(context *gin.Context) {
		// Get request url path
		url := context.Param("path")
		route, done := parseUrlToGetRoute(context, url)
		if !done {
			context.JSON(404, gin.H{"message": "Route not found."})
			return
		}
		// according to balance algorithm, select a path
		prefix := getRoutePrefix(route)

		// Check if the route requires jwt
		if route.Jwt {
			done, errMsg := checkJwt(context)
			if !done {
				context.JSON(401, gin.H{"message": errMsg})
				return
			}
		}

		// Create new request
		request, err := http.NewRequest(context.Request.Method, prefix+url, context.Request.Body)
		if err != nil {
			context.JSON(500, gin.H{"message": "Failed to create request." + err.Error()})
			return
		}
		request.Header = context.Request.Header // copy headers

		// Create a client to send the request
		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			context.JSON(500, gin.H{"message": "Failed to send request." + err.Error()})
			return
		} else {
			// fmt.Println(output.Info() + "Request sent to " + prefix + url)
		}
		defer response.Body.Close()

		// Copy the response to the context
		// 1. response header
		for key, value := range response.Header {
			context.Writer.Header()[key] = value
		}
		// 2. response status code
		context.Writer.WriteHeader(response.StatusCode)
		// 3. response body
		io.Copy(context.Writer, response.Body)
	})
}

// According to the url, find the route that matches the url
// example: /api/v1/user, we will gradually match /api, /api/v1, /api/v1/user
func parseUrlToGetRoute(context *gin.Context, url string) (*settings.Route, bool) {
	splits := strings.Split(url, "/")
	for i := 0; i < len(splits); i++ {
		tempUrl := strings.Join(splits[0:i], "/")
		if route, exist := settings.Config.Routes[tempUrl]; exist {
			return route, true
		}
	}
	return nil, false
}

// According to the route and the balance algorithm, get the prefix of the route
func getRoutePrefix(route *settings.Route) string {
	return balance.Round(route)
}

// Check jwt
func checkJwt(context *gin.Context) (bool, string) {
	// TODO check jwt
	// fmt.Println(output.Info() + "Checking jwt and jwt key is " + settings.Config.JwtKey)
	return true, ""
}
