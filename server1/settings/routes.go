package settings

import (
	"fmt"
	"github.com/spf13/viper"
	"server1/output"
)

func initRoutes() {
	var routes []map[string]interface{}
	err := viper.UnmarshalKey("routes", &routes)
	if err != nil {
		fmt.Println(output.Error() + "Failed to read routes")
	}

	for _, route := range routes {
		var name = route["name"].(string)

		var tempRoute = &Route{
			Name:  name,
			Url:   checkPrefix(name, route["prefix"].(string)),
			Jwt:   checkJwt(name, route["jwt"]),
			Paths: checkPaths(name, route["route"]),
		}

		Config.Routes[tempRoute.Url] = tempRoute
	}
}

// 1. if prefix is empty, panic
// 2. if prefix does not start with a '/', add it
func checkPrefix(name, prefix string) string {
	if prefix == "" {
		panic(output.Error() + "Route: " + name + " does not have a prefix")
	} else if prefix[0] != '/' {
		return "/" + prefix
	} else {
		return prefix
	}
}

// if jwt is not a boolean or don't have a jwt attribute, default to false
func checkJwt(name string, jwt interface{}) bool {
	b, ok := jwt.(bool)
	if !ok {
		fmt.Println(output.Default() + "route: " + name + " does not have a jwt field or it is not a boolean. Defaulting to false")
		return false
	}
	return b
}

// step by step to transform path
func checkPaths(name string, origin any) []*RouteInfo {
	var routeInfos []*RouteInfo

	// transform origin into []interface{}
	var paths, ok = origin.([]interface{})
	if !ok {
		panic(output.Error() + "Route: " + name + " failed to transform config into []interface{}")
	}
	for i := 0; i < len(paths); i++ {
		m, ok := paths[i].(map[string]interface{})
		if !ok {
			panic(output.Error() + "Route: " + name + " failed to transform config into map[string]interface{}")
		}
		var tempRouteInfo = &RouteInfo{
			Path:   m["path"].(string),
			Weight: checkWeight(name, m["weight"]),
		}
		routeInfos = append(routeInfos, tempRouteInfo)
	}

	return routeInfos
}

func checkWeight(name string, weight interface{}) int {
	w, ok := weight.(int)
	if !ok {
		fmt.Println(output.Default() + "route: " + name + " does not have a weight field or it is not an integer. Defaulting to 1")
		return 1
	}
	return w
}
