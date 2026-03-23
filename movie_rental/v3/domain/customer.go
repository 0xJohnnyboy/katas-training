package domain

import "fmt"

type Customer struct {
	name    string
	rentals []Rental
}

func NewCustomer(name string, rentals []Rental) *Customer {
	return &Customer{
		name,
		rentals,
	}
}

func (c *Customer) AddRental(rental Rental) {
	c.rentals = append(c.rentals, rental)
}

// deprecated
func (c Customer) statement() string {
	totalAmount := 0.0
	frequentRenterPoints := 0
	result := "Rental Record for " + c.Name() + "\n"

	for _, each := range c.rentals {
		thisAmount := 0.0

		//determine amounts for each line
		switch each.Movie().PriceCode() {
		case REGULAR:
			thisAmount += 2
			if each.DaysRented() > 2 {
				thisAmount += (each.DaysRented() - 2) * 1.5
			}
			break
		case NEW_RELEASE:
			thisAmount += each.DaysRented() * 3
			break
		case CHILDREN:
			thisAmount += 1.5
			if each.DaysRented() > 3 {
				thisAmount += (each.DaysRented() - 3) * 1.5
			}
			break
		}

		// add frequent renter points
		frequentRenterPoints++
		// add bonus for a two day new release rental
		if (each.Movie().PriceCode() == NEW_RELEASE) && each.DaysRented() > 1 {
			frequentRenterPoints++
		}

		// show figures for this rental
		result += fmt.Sprintf("\t"+each.Movie().Title()+"\t%.1f\n", thisAmount)
		totalAmount += thisAmount
	}

	// add footer lines
	result += fmt.Sprintf("Amount owed is %.1f\n", totalAmount)
	result += fmt.Sprintf("You earned %d frequent renter points", frequentRenterPoints)

	return result

}

func (c Customer) Name() string {
	return c.name
}

func (c Customer) Rentals() []Rental {
	return c.rentals
}
