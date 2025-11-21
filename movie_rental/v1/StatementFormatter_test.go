package movie_rental

import (
	"testing"
)

func TestStatementFormatter(t *testing.T) {
	customer := Customer{"Bob", []Rental{}, 0}
	customer.addRental(Rental{Movie{"Jaws", REGULAR}, 2})
	customer.addRental(Rental{Movie{"Golden Eye", REGULAR}, 3})
	customer.addRental(Rental{Movie{"Short New", NEW_RELEASE}, 1})
	customer.addRental(Rental{Movie{"Long New", NEW_RELEASE}, 2})
	customer.addRental(Rental{Movie{"Bambi", CHILDREN}, 3})
	customer.addRental(Rental{Movie{"Toy Story", CHILDREN}, 4})

	t.Run("legacy statement method", func(t *testing.T) {
		t.Parallel()
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
		if expected != customer.Statement() {
			t.Fatalf("Expected:\n%s\nActual:\n%s\n", expected, customer.Statement())
		}
	})

	t.Run("plain text printing", func(t *testing.T) {
		t.Parallel()
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

		formatter := GetStatementFormatter(PlainText)
		actual := formatter.Format(customer)
		if expected != actual {
			t.Fatalf("Expected:\n%s\nActual:\n%s\n", expected, actual)
		}
	})

	t.Run("html printing", func(t *testing.T) {
		t.Parallel()
		expected := "<h1>Rental Record for <em>Bob</em></h1>" +
			"<table>" +
			"<tr><td>Jaws</td><td>2.0</td></tr>" +
			"<tr><td>Golden Eye</td><td>3.5</td></tr>" +
			"<tr><td>Short New</td><td>3.0</td></tr>" +
			"<tr><td>Long New</td><td>6.0</td></tr>" +
			"<tr><td>Bambi</td><td>1.5</td></tr>" +
			"<tr><td>Toy Story</td><td>3.0</td></tr>" +
			"</table>" +
			"<p>Amount owed is <em>19.0</em></p>" +
			"<p>You earned <em>7</em> frequent renter points</p>"
		formatter := GetStatementFormatter(HTML)
		actual := formatter.Format(customer)
		if expected != actual {
			t.Fatalf("Expected:\n%s\nActual:\n%s\n", expected, actual)
		}

	})
}
