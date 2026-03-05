package domain

import (
	"testing"
)

func TestRental(t *testing.T) {

	regularMovie := NewMovie("regular", REGULAR)
	newReleaseMovie := NewMovie("new release", NEW_RELEASE)
	childrenMovie := NewMovie("children", CHILDREN)
	badMovie := NewMovie("bad", 12)
	unknownPcError := "unknown price code: 12"

	type RentalCostTestCase struct {
		desc                 string
		rental               *Rental
		expectedCost         float64
		expectedRenterPoints uint
		expectedError        *string
	}

	testCases := []RentalCostTestCase{
		{"regular movie 2 days", NewRental(regularMovie, 2), 2.0, 1, nil},
		{"regular movie 3 days", NewRental(regularMovie, 3), 3.5, 1, nil},
		{"new release movie 1 day", NewRental(newReleaseMovie, 1), 3.0, 1, nil},
		{"new release movie 2 days", NewRental(newReleaseMovie, 2), 6.0, 2, nil},
		{"children movie 1 day", NewRental(childrenMovie, 1), 1.5, 1, nil},
		{"children movie 4 days", NewRental(childrenMovie, 4), 3.0, 1, nil},
		{"non existent price code", NewRental(badMovie, 1), 1.0, 1, &unknownPcError},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			actualCost, err := tc.rental.Cost()

			if err == nil && tc.expectedError != nil {
				t.Fatalf("Expected error %s got %v", *tc.expectedError, err)
			}

			if tc.expectedError != nil && *tc.expectedError != err.Error() {
				t.Fatalf("Expected error %s got %v", *tc.expectedError, err)
			}

			if tc.expectedError == nil && err != nil {

				t.Fatalf("Expected cost %f got %v", tc.expectedCost, err)
			}

			if tc.expectedError != nil {
				return
			}

			if tc.expectedCost != actualCost {
				t.Errorf("Expected cost %f got %f", tc.expectedCost, actualCost)
			}
			actualRenterPoints := tc.rental.RenterPoints()
			if tc.expectedRenterPoints != actualRenterPoints {
				t.Errorf("Expected %d renter points got %d", tc.expectedRenterPoints, actualRenterPoints)
			}
		})
	}

}
