package movie_rental

import "testing"

func TestCustomer(t *testing.T) {
	t.Run("adding a rental should add frequent renter points", func(t *testing.T) {
		t.Parallel()
		customer := &Customer{"Bob", []Rental{}, 0}
		customer.addRental(Rental{Movie{"Golden Eye", REGULAR}, 1})
		var expected uint
		expected = 1
		actual := customer.FrequentRenterPoints()

		if actual != expected {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	})

	t.Run("adding a new release rental for over 1 day should add frequent renter points with a bonus", func(t *testing.T) {
		t.Parallel()
		customer := &Customer{"Bob", []Rental{}, 0}
		customer.addRental(Rental{Movie{"Golden Eye", NEW_RELEASE}, 3})
		var expected uint
		expected = 2
		actual := customer.FrequentRenterPoints()

		if actual != expected {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	})

}
