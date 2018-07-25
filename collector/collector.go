package collector

import (
	"io"
)

// Collector interface for getting data
// from an APIs. Returns the data from
// hitting that API with the specified
// query parameters, if supported
type Collector interface {
	GetData(query Query) io.Reader
	SupportedParameters() map[parameterType]bool
}
