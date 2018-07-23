package storer

import "encoding/json"

// Storer writes data collected from Collector to some form
// of persistent storage
type Storer interface {
	initialize() error
	saveEntries(json json.Encoder, tag string) error
}
