package collector

type parameterType int

const (
	keywords parameterType = iota
	dateRange
	user
	searchPhrase
)

func (pType parameterType) String() string {
	switch pType {
	case keywords:
		return "keywords"
	case dateRange:
		return "date range"
	case user:
		return "user"
	case searchPhrase:
		return "search phrase"
	default:
		return "invalid type"
	}
}
