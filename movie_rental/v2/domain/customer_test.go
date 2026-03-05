package domain

import (
	"testing"
)

func TestCustomer(t *testing.T) {
	t.Run("add rental to customer", func(t *testing.T) {
		movie := NewMovie("Jaws", REGULAR)
		rental := NewRental(movie, 2)
		customer := NewCustomer("Bob")

		customer.AddRental(rental)

		expectedRentalsCount := 1
		expectedMovieTitle := "Jaws"
		expectedDaysRented := uint(2)

		actualRentals := customer.Rentals()

		if expectedRentalsCount != len(actualRentals) {
			t.Fatalf("Expected %d rentals got %d", expectedRentalsCount, len(actualRentals))
		}

		actualMovieTitle := actualRentals[0].Movie().Title()
		if expectedMovieTitle != actualMovieTitle {
			t.Errorf("Expected %s got %s", expectedMovieTitle, actualMovieTitle)
		}

		actualDaysRented := actualRentals[0].DaysRented()
		if expectedDaysRented != actualDaysRented {
			t.Errorf("Expected %d got %d", expectedDaysRented, actualDaysRented)
		}
	})

}
