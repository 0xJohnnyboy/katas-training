package domain

type PriceCode uint

const (
	REGULAR PriceCode = iota
	NEW_RELEASE
	CHILDREN
)

type Movie struct {
	title     string
	priceCode PriceCode
}

func NewMovie(title string, priceCode PriceCode) *Movie {
	return &Movie{title, priceCode}
}

func (m Movie) PriceCode() PriceCode {
	return m.priceCode
}

func (m Movie) Title() string {
	return m.title
}
