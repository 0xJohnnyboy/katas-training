package domain

import (
	"testing"
)

func TestCustomer(t *testing.T) {
	customer := NewCustomer("Bob", []Rental{})
	customer.AddRental(*NewRental(*NewMovie("Jaws", REGULAR), 2))
	customer.AddRental(*NewRental(*NewMovie("Golden Eye", REGULAR), 3))
	customer.AddRental(*NewRental(*NewMovie("Short New", NEW_RELEASE), 1))
	customer.AddRental(*NewRental(*NewMovie("Long New", NEW_RELEASE), 2))
	customer.AddRental(*NewRental(*NewMovie("Bambi", CHILDREN), 3))
	customer.AddRental(*NewRental(*NewMovie("Toy Story", CHILDREN), 4))
}
