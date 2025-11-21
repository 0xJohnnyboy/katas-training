package movie_rental

import "testing"

func TestRental(t *testing.T) {
	t.Run("NEW_RELEASE: 3 days", func(t *testing.T) {
		t.Parallel()
		movie := Movie{"Matrix", NEW_RELEASE}
		rental := Rental{movie, 3}
		expected := 9.0 // 3 days * 3

		actual := rental.CalculateCost()

		if actual != expected {
			t.Errorf("Expected %.2f, got %.2f", expected, actual)
		}
	})

	t.Run("REGULAR: 2 days or less", func(t *testing.T) {
		t.Parallel()
		movie := Movie{"Titanic", REGULAR}
		rental := Rental{movie, 2}
		expected := 2.0 // flat rate

		actual := rental.CalculateCost()

		if actual != expected {
			t.Errorf("Expected %.2f, got %.2f", expected, actual)
		}
	})

	t.Run("REGULAR: more than 2 days", func(t *testing.T) {
		t.Parallel()
		movie := Movie{"Titanic", REGULAR}
		rental := Rental{movie, 5}
		expected := 6.5 // 2 + (5-2)*1.5

		actual := rental.CalculateCost()

		if actual != expected {
			t.Errorf("Expected %.2f, got %.2f", expected, actual)
		}
	})

	t.Run("CHILDREN: 3 days or less", func(t *testing.T) {
		t.Parallel()
		movie := Movie{"Toy Story", CHILDREN}
		rental := Rental{movie, 3}
		expected := 1.5 // flat rate

		actual := rental.CalculateCost()

		if actual != expected {
			t.Errorf("Expected %.2f, got %.2f", expected, actual)
		}
	})

	t.Run("CHILDREN: more than 3 days", func(t *testing.T) {
		t.Parallel()
		movie := Movie{"Toy Story", CHILDREN}
		rental := Rental{movie, 5}
		expected := 4.5 // 1.5 + (5-3)*1.5

		actual := rental.CalculateCost()

		if actual != expected {
			t.Errorf("Expected %.2f, got %.2f", expected, actual)
		}
	})
}
