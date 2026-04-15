package movie_rental

import (
	"testing"
)

func TestHTMLFormatter(t *testing.T) {

	t.Run("should return html statement", func(t *testing.T) {
		customer := Customer{"Bob", []Rental{}}
		customer.AddRental(Rental{Movie{"Jaws", REGULAR}, 2})
		customer.AddRental(Rental{Movie{"Golden Eye", REGULAR}, 3})
		customer.AddRental(Rental{Movie{"Short New", NEW_RELEASE}, 1})
		customer.AddRental(Rental{Movie{"Long New", NEW_RELEASE}, 2})
		customer.AddRental(Rental{Movie{"Bambi", CHILDREN}, 3})
		customer.AddRental(Rental{Movie{"Toy Story", CHILDREN}, 4})
		expected := "" +
			"<h1>Rental Record for <em>Bob</em></h1>\n" +
			"<table>\n" +
			"<tr><td>Jaws</td><td>2.0</td></tr>\n" +
			"<tr><td>Golden Eye</td><td>3.5</td></tr>\n" +
			"<tr><td>Short New</td><td>3.0</td></tr>\n" +
			"<tr><td>Long New</td><td>6.0</td></tr>\n" +
			"<tr><td>Bambi</td><td>1.5</td></tr>\n" +
			"<tr><td>Toy Story</td><td>3.0</td></tr>\n" +
			"</table>\n" +
			"<p>Amount owed is <em>19.0</em></p>\n" +
			"<p>You earned <em>7</em> frequent renter points</p>"
		formatter := NewHTMLFormatter()
		actual := formatter.FormatStatement(customer)
		if expected != actual {
			t.Fatalf("Expected:\n%s\nActual:\n%s\n", expected, actual)
		}

	})
}
