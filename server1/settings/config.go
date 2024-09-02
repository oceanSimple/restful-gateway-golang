package settings

type ConfigClass struct {
	Port   int64
	JwtKey string
	Routes map[string]*Route
}

type Route struct {
	Name    string
	Url     string
	Jwt     bool
	Paths   []*RouteInfo
	Counter int64 // counter for round robin
}

type RouteInfo struct {
	Path   string
	Weight int
}
