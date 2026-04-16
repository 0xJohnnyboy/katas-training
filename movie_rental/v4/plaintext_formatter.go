package movie_rental

import (
	"fmt"
)

type PlainTextFormatter struct{}

func NewPlainTextFormatter() *PlainTextFormatter {
	return &PlainTextFormatter{}
}

func (f PlainTextFormatter) FormatStatement(c Customer) string {
	totalAmount := 0.0
	frequentRenterPoints := 0
	result := "Rental Record for " + c.Name() + "\n"

	for _, rental := range c.rentals {
		thisAmount := rental.Cost()

		frequentRenterPoints += rental.RenterPoints()

		// show figures for this rental
		result += fmt.Sprintf("\t"+rental.Movie().Title()+"\t%.1f\n", thisAmount)
		totalAmount += thisAmount
	}

	// add footer lines
	result += fmt.Sprintf("Amount owed is %.1f\n", totalAmount)
	result += fmt.Sprintf("You earned %d frequent renter points", frequentRenterPoints)

	return result
}
