package infrastructure

import (
	. "movie-rental-golang/domain"
	"testing"
)

func TestHTMLFormatter(t *testing.T) {
	t.Run("Format stmt as html", func(t *testing.T) {
		customer := NewCustomer("Bob")
		customer.AddRental(NewRental(NewMovie("Jaws", REGULAR), 2))
		customer.AddRental(NewRental(NewMovie("Golden Eye", REGULAR), 3))
		customer.AddRental(NewRental(NewMovie("Short New", NEW_RELEASE), 1))
		customer.AddRental(NewRental(NewMovie("Long New", NEW_RELEASE), 2))
		customer.AddRental(NewRental(NewMovie("Bambi", CHILDREN), 3))
		customer.AddRental(NewRental(NewMovie("Toy Story", CHILDREN), 4))

		expected := "" +
			"<h1>Rental Record for <em>Bob</em></h1>" +
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
		formatter := FormattersRegistry[HTML]()
		actual, err := formatter.Format(customer)
		if err != nil {
			t.Fatal(err)
		}
		if expected != actual {
			t.Errorf("Expected %s rentals got %s", expected, actual)
		}
	})
}
