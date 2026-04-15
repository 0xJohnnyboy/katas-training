package movie_rental

import (
	"testing"
)

func TestPlainTextFormatter(t *testing.T) {

	t.Run("should return plain statement", func(t *testing.T) {
		customer := Customer{"Bob", []Rental{}}
		customer.AddRental(Rental{Movie{"Jaws", REGULAR}, 2})
		customer.AddRental(Rental{Movie{"Golden Eye", REGULAR}, 3})
		customer.AddRental(Rental{Movie{"Short New", NEW_RELEASE}, 1})
		customer.AddRental(Rental{Movie{"Long New", NEW_RELEASE}, 2})
		customer.AddRental(Rental{Movie{"Bambi", CHILDREN}, 3})
		customer.AddRental(Rental{Movie{"Toy Story", CHILDREN}, 4})
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
		actual := formatter.FormatStatement(customer)

		if expected != actual {
			t.Fatalf("Expected:\n%s\nActual:\n%s\n", expected, actual)
		}

	})
}
