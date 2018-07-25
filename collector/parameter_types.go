package collector

type parameterType int

const (
	keywords parameterType = iota
	dateRange
	user
	searchPhrase
	limit
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
	case limit:
		return "limit"
	default:
		return "invalid type"
	}
}
