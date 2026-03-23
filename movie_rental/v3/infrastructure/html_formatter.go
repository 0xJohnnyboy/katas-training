package infrastructure

import (
	domain "movie-rental-golang/domain"

	"fmt"
)

type HTMLFormatter struct{}

func NewHTMLFormatter() *HTMLFormatter {
	return &HTMLFormatter{}
}

func (f HTMLFormatter) Format(c *domain.Customer) string {
	totalCost := 0.0
	frequentRenterPoints := 0
	result := "<h1>Rental Record for <em>" + c.Name() + "</em></h1>\n"
	result += "<table>\n"

	for _, rental := range c.Rentals() {
		currentRentalCost, _ := rental.Cost()
		frequentRenterPoints += rental.RenterPoints()
		result += fmt.Sprintf("<tr><td>"+rental.Movie().Title()+"</td><td>%.1f</td></tr>\n", currentRentalCost)
		totalCost += currentRentalCost
	}

	result += "</table>\n"

	// add footer lines
	result += fmt.Sprintf("<p>Amount owed is <em>%.1f</em></p>\n", totalCost)
	result += fmt.Sprintf("<p>You earned <em>%d</em> frequent renter points</p>", frequentRenterPoints)

	return result

}
