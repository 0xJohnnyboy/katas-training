package infrastructure

import (
	"fmt"
	domain "movie-rental-golang/domain"
)

type HTMLFormatter struct{}

func NewHTMLFormatter() *HTMLFormatter {
	return &HTMLFormatter{}
}

func (f HTMLFormatter) Format(c *domain.Customer) (string, error) {
	var totalAmount float64
	var frequentRenterPoints uint

	result := wrapWith(H1, "Rental Record for "+wrapWith(EM, c.Name()))
	var htmlTable string

	for _, rental := range c.Rentals() {
		currentCost, err := rental.Cost()
		if err != nil {
			return "", err
		}
		frequentRenterPoints += rental.RenterPoints()

		title := wrapWith(TD, rental.Movie().Title())
		cost := wrapWith(TD, fmt.Sprintf("%.1f", currentCost))
		htmlTable += wrapWith(TR, title+cost)
		totalAmount += currentCost
	}

	result += wrapWith(TABLE, htmlTable)

	formattedAmount := wrapWith(EM, fmt.Sprintf("%.1f", totalAmount))
	formattedPoints := wrapWith(EM, fmt.Sprintf("%d", frequentRenterPoints))

	result += wrapWith(P, "Amount owed is "+formattedAmount)
	result += wrapWith(P, "You earned "+formattedPoints+" frequent renter points")

	return result, nil
}

type Markup string

const (
	H1    Markup = "h1"
	EM    Markup = "em"
	TABLE Markup = "table"
	TR    Markup = "tr"
	TD    Markup = "td"
	P     Markup = "p"
)

func wrapWith(m Markup, s string) string {
	return fmt.Sprintf("<%s>%s</%s>", m, s, m)
}
