package infrastructure

import (
	"testing"

	. "movie-rental-golang/domain"
)

func TestHTMLFormatter(t *testing.T) {
	customer := NewCustomer("Bob", []Rental{})
	customer.AddRental(*NewRental(*NewMovie("Jaws", REGULAR), 2))
	customer.AddRental(*NewRental(*NewMovie("Golden Eye", REGULAR), 3))
	customer.AddRental(*NewRental(*NewMovie("Short New", NEW_RELEASE), 1))
	customer.AddRental(*NewRental(*NewMovie("Long New", NEW_RELEASE), 2))
	customer.AddRental(*NewRental(*NewMovie("Bambi", CHILDREN), 3))
	customer.AddRental(*NewRental(*NewMovie("Toy Story", CHILDREN), 4))

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
	stmt := formatter.Format(customer)

	if expected != stmt {
		t.Fatalf("Expected:\n%s\nActual:\n%s\n", expected, stmt)
	}
}
