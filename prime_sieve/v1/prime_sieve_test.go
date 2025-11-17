package v1

import (
	"fmt"
	"slices"
	"testing"
)

func TestPrimeSieve(t *testing.T) {

	testCases := []struct {
		Name     string
		Max      int
		Expected []int
	}{
		{
			Name:     "Expects 4 prime number with max= 10",
			Max:      10,
			Expected: []int{2, 3, 5, 7},
		},
		{
			Name:     "Expects 8 prime number with max= 10",
			Max:      20,
			Expected: []int{2, 3, 5, 7, 11, 13, 17, 19}},
		{
			Name:     "Expects 10 prime number with max= 10",
			Max:      30,
			Expected: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29},
		},
		{
			Name:     "Expects 15 prime number with max= 10",
			Max:      50,
			Expected: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47},
		},
	}

	for _, tc := range testCases {
		desc := fmt.Sprintf(tc.Name, tc.Max, len(tc.Expected))

		t.Run(desc, func(t *testing.T) {
			actual := GetPrimeNumbers(tc.Max)
			if !slices.Equal(tc.Expected, actual) {
				t.Errorf("Expected %v, got %v", tc.Expected, actual)

			}
		})
	}

}
