package storer

import "io"

// Storer writes data collected from Collector to some form
// of persistent storage
type Storer interface {
	Initialize() error
	SaveEntries(json io.Reader, tag string) error
}
