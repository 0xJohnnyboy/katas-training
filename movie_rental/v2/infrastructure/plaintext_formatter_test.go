package infrastructure

import (
	. "movie-rental-golang/domain"
	"testing"
)

func TestPlainTextFormatter(t *testing.T) {
	t.Run("Format stmt as plain text", func(t *testing.T) {
		customer := NewCustomer("Bob")
		customer.AddRental(NewRental(NewMovie("Jaws", REGULAR), 2))
		customer.AddRental(NewRental(NewMovie("Golden Eye", REGULAR), 3))
		customer.AddRental(NewRental(NewMovie("Short New", NEW_RELEASE), 1))
		customer.AddRental(NewRental(NewMovie("Long New", NEW_RELEASE), 2))
		customer.AddRental(NewRental(NewMovie("Bambi", CHILDREN), 3))
		customer.AddRental(NewRental(NewMovie("Toy Story", CHILDREN), 4))

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
		formatter := FormattersRegistry[PlainText]()
		actual, err := formatter.Format(customer)
		if err != nil {
			t.Fatal(err)
		}
		if expected != actual {
			t.Errorf("Expected %s rentals got %s", expected, actual)
		}
	})
}
