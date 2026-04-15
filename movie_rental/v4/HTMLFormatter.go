package movie_rental

import (
	"fmt"
)

type HTMLFormatter struct{}

func NewHTMLFormatter() *HTMLFormatter {
	return &HTMLFormatter{}
}

func (f HTMLFormatter) FormatStatement(c Customer) string {
	totalAmount := 0.0
	frequentRenterPoints := 0
	result := "<h1>Rental Record for <em>" + c.Name() + "</em></h1>\n"
	result += "<table>\n"

	for _, rental := range c.rentals {
		thisAmount := rental.Cost()
		frequentRenterPoints += rental.RenterPoints()

		// show figures for this rental
		result += fmt.Sprintf("<tr><td>"+rental.Movie().Title()+"</td><td>%.1f</td></tr>\n", thisAmount)
		totalAmount += thisAmount
	}

	result += "</table>\n"
	// add footer lines
	result += fmt.Sprintf("<p>Amount owed is <em>%.1f</em></p>\n", totalAmount)
	result += fmt.Sprintf("<p>You earned <em>%d</em> frequent renter points</p>", frequentRenterPoints)

	return result
}
