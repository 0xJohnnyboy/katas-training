package movie_rental

type Movie struct {
	title     string
	priceCode PriceCode
}

type PriceCode uint

const (
	REGULAR PriceCode = iota
	NEW_RELEASE
	CHILDREN
)

func (m Movie) PriceCode() PriceCode {
	return m.priceCode
}

func (m Movie) Title() string {
	return m.title
}
