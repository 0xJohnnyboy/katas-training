package infrastructure

import (
	"fmt"
	domain "movie-rental-golang/domain"
)

type PlainTextFormatter struct {
}

func NewPlainTextFormatter() *PlainTextFormatter {
	return &PlainTextFormatter{}
}

func (f PlainTextFormatter) Format(c *domain.Customer) (string, error) {
	var totalAmount float64
	var frequentRenterPoints uint

	result := fmt.Sprintf("Rental Record for %s\n", c.Name())

	for _, rental := range c.Rentals() {
		currentRentalPrice, err := rental.Cost()
		if err != nil {
			return "", err
		}
		frequentRenterPoints += rental.RenterPoints()

		result += fmt.Sprintf("\t"+rental.Movie().Title()+"\t%.1f\n", currentRentalPrice)
		totalAmount += currentRentalPrice
	}

	result += fmt.Sprintf("Amount owed is %.1f\n", totalAmount)
	result += fmt.Sprintf("You earned %d frequent renter points", frequentRenterPoints)

	return result, nil
}
