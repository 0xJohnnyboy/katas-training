package infrastructure

import (
	domain "movie-rental-golang/domain"

	"fmt"
)

type PlainTextFormatter struct{}

func NewPlainTextFormatter() *PlainTextFormatter {
	return &PlainTextFormatter{}
}

func (f PlainTextFormatter) Format(c *domain.Customer) string {
	totalCost := 0.0
	frequentRenterPoints := 0
	result := "Rental Record for " + c.Name() + "\n"

	for _, rental := range c.Rentals() {
		currentRentalCost, _ := rental.Cost()
		frequentRenterPoints += rental.RenterPoints()

		// show figures for this rental
		result += fmt.Sprintf("\t"+rental.Movie().Title()+"\t%.1f\n", currentRentalCost)
		totalCost += currentRentalCost
	}

	// add footer lines
	result += fmt.Sprintf("Amount owed is %.1f\n", totalCost)
	result += fmt.Sprintf("You earned %d frequent renter points", frequentRenterPoints)

	return result

}
