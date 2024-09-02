package balance

import "server1/settings"

func Round(route *settings.Route) string {
	route.Counter++
	return route.Paths[route.Counter%int64(len(route.Paths))].Path
}
