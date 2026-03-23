package infrastructure

import (
	"testing"

	. "movie-rental-golang/domain"
)

func TestPlainTextFormatter(t *testing.T) {
	customer := NewCustomer("Bob", []Rental{})
	customer.AddRental(*NewRental(*NewMovie("Jaws", REGULAR), 2))
	customer.AddRental(*NewRental(*NewMovie("Golden Eye", REGULAR), 3))
	customer.AddRental(*NewRental(*NewMovie("Short New", NEW_RELEASE), 1))
	customer.AddRental(*NewRental(*NewMovie("Long New", NEW_RELEASE), 2))
	customer.AddRental(*NewRental(*NewMovie("Bambi", CHILDREN), 3))
	customer.AddRental(*NewRental(*NewMovie("Toy Story", CHILDREN), 4))

	expected := "" +
		"Rental Record for Bob\n" +
		"\tJaws\t2.0\n" +
		"\tGolden Eye\t3.5\n" +
		"\tShort New\t3.0\n" +
		"\tLong New\t6.0\n" +
		"\tBambi\t1.5\n" +
		"\tToy Story\t3.0\n" +
		"Amount owed is 19.0\n" +
		"You earned 7 frequent renter points"

	formatter := NewPlainTextFormatter()
	stmt := formatter.Format(customer)

	if expected != stmt {
		t.Fatalf("Expected:\n%s\nActual:\n%s\n", expected, stmt)
	}
}
