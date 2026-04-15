package movie_rental

import (
	"testing"
)

func TestRental(t *testing.T) {
	t.Run("should calculate cost", func(t *testing.T) {
		testcases := []struct {
			name       string
			priceCode  PriceCode
			daysRented float64
			expected   float64
		}{
			{name: "regular under threshold", priceCode: REGULAR, daysRented: 1, expected: 2},
			{name: "regular at threshold", priceCode: REGULAR, daysRented: 2, expected: 2},
			{name: "regular over threshold", priceCode: REGULAR, daysRented: 3, expected: 3.5},
			{name: "new release one day", priceCode: NEW_RELEASE, daysRented: 1, expected: 3},
			{name: "new release multiple days", priceCode: NEW_RELEASE, daysRented: 2, expected: 6},
			{name: "children under threshold", priceCode: CHILDREN, daysRented: 2, expected: 1.5},
			{name: "children at threshold", priceCode: CHILDREN, daysRented: 3, expected: 1.5},
			{name: "children over threshold", priceCode: CHILDREN, daysRented: 4, expected: 3},
		}

		for _, tc := range testcases {
			t.Run(tc.name, func(t *testing.T) {
				rental := Rental{
					movie:      Movie{title: "Movie", priceCode: tc.priceCode},
					daysRented: tc.daysRented,
				}

				actual := rental.Cost()
				if actual != tc.expected {
					t.Fatalf("expected %.1f, got %.1f", tc.expected, actual)
				}
			})
		}
	})
	t.Run("should calculate frequent renter points", func(t *testing.T) {
		testcases := []struct {
			name       string
			priceCode  PriceCode
			daysRented float64
			expected   int
		}{
			{name: "regular rental", priceCode: REGULAR, daysRented: 5, expected: 1},
			{name: "children rental", priceCode: CHILDREN, daysRented: 5, expected: 1},
			{name: "new release one day", priceCode: NEW_RELEASE, daysRented: 1, expected: 1},
			{name: "new release more than one day", priceCode: NEW_RELEASE, daysRented: 2, expected: 2},
		}

		for _, tc := range testcases {
			t.Run(tc.name, func(t *testing.T) {
				rental := Rental{
					movie:      Movie{title: "Movie", priceCode: tc.priceCode},
					daysRented: tc.daysRented,
				}

				actual := rental.RenterPoints()
				if actual != tc.expected {
					t.Fatalf("expected %d, got %d", tc.expected, actual)
				}
			})
		}
	})
}
