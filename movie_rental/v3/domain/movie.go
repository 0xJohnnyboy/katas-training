package domain

const (
	REGULAR uint = iota
	NEW_RELEASE
	CHILDREN
)

type Movie struct {
	title     string
	priceCode uint
}

func NewMovie(title string, priceCode uint) *Movie {
	return &Movie{
		title,
		priceCode,
	}
}

func (m Movie) PriceCode() uint {
	return m.priceCode
}

func (m Movie) Title() string {
	return m.title
}
