package balance

import "server1/settings"

type Tactic interface {
	GetPrefix(route settings.Route) string
}
