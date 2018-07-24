package collector

import "time"

// Query specifies the parameters
// to filter data by from an API
// endpoint. For use with Collector
type Query struct {
	Keywords     []string
	StartDate    time.Time
	EndDate      time.Time
	User         string
	SearchPhrase string
}

// UnsupportedWarning provides a warning to log to users
// about which of their specified query parameters are
// unsupported on a particular Collector
func (query *Query) UnsupportedWarning(collector *Collector) string {
	return ""
}
