package settings

type ConfigClass struct {
	Port   int64
	Routes []Route
}

type Route struct {
	Name   string
	Prefix string
	Jwt    bool
	Paths  []RouteInfo
}

type RouteInfo struct {
	Path   string
	Weight int
}
