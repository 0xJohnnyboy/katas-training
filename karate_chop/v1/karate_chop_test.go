package v1

import (
	"slices"
	"testing"
)

type TestCase struct {
	Target   int
	Array    []int
	Expected int
}

type TestCaseGroup struct {
	Desc  string
	Cases []TestCase
}

var testCaseGroups []TestCaseGroup = []TestCaseGroup{
	{
		Desc: "Empty slice should return -1",
		Cases: []TestCase{
			{3, []int{}, -1},
			{0, []int{}, -1},
		},
	},
	{
		Desc: "Array with 1 element",
		Cases: []TestCase{
			{1, []int{1}, 0},
			{3, []int{1}, -1},
		},
	},
	{
		Desc: "Array with multiple elements, should find",
		Cases: []TestCase{
			{1, []int{1, 3, 5}, 0},
			{3, []int{1, 3, 5}, 1},
			{5, []int{1, 3, 5}, 2},
			{1, []int{1, 3, 5, 7}, 0},
			{3, []int{1, 3, 5, 7}, 1},
			{5, []int{1, 3, 5, 7}, 2},
			{7, []int{1, 3, 5, 7}, 3},
		},
	},
	{
		Desc: "Array with multiple elements, should not find",
		Cases: []TestCase{
			{0, []int{1, 3, 5}, -1},
			{2, []int{1, 3, 5}, -1},
			{4, []int{1, 3, 5}, -1},
			{6, []int{1, 3, 5}, -1},
			{0, []int{1, 3, 5, 7}, -1},
			{2, []int{1, 3, 5, 7}, -1},
			{4, []int{1, 3, 5, 7}, -1},
			{6, []int{1, 3, 5, 7}, -1},
			{8, []int{1, 3, 5, 7}, -1},
		},
	},
	{
		Desc: "Bigger slices",
		Cases: []TestCase{
			{1, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 0},
			{5, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 4},
			{9, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 8},
			{0, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, -1},
			{10, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, -1},
			{3, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2},
			{7, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 6},
		},
	},
	{
		Desc: "Negative numbers",
		Cases: []TestCase{
			{-5, []int{-5, -3, -1, 0, 2}, 0},
			{-3, []int{-5, -3, -1, 0, 2}, 1},
			{0, []int{-5, -3, -1, 0, 2}, 3},
			{2, []int{-5, -3, -1, 0, 2}, 4},
			{-4, []int{-5, -3, -1, 0, 2}, -1},
			{1, []int{-5, -3, -1, 0, 2}, -1},
		},
	},
}

func runCases(t *testing.T, testCases []TestCase) {
	for _, tc := range testCases {
		result := Chop(tc.Target, tc.Array)
		if result != tc.Expected {
			t.Errorf("Array: %v with target %d, Expected %d, got %d", tc.Array, tc.Target, tc.Expected, result)
		}
	}
}

func TestKarateChop(t *testing.T) {
	for _, tcg := range testCaseGroups {
		t.Run(tcg.Desc, func(t *testing.T) {
			runCases(t, tcg.Cases)
		})
	}

	t.Run("Edge Cases", func(t *testing.T) {
		t.Run("Duplicates: should return index 1, 2 or 3", func(t *testing.T) {
			target := 3
			array := []int{1, 3, 3, 3, 5}
			expected := []int{1, 2, 3}

			result := Chop(target, array)
			if !slices.Contains(expected, result) {

				t.Errorf("Array %v with target %d, expected one of %v. Got %d", array, target, expected, result)
			}

		})
		t.Run("Duplicates: should return index 2 or 3", func(t *testing.T) {
			target := 5
			array := []int{1, 3, 5, 5, 7}
			expected := []int{2, 3}

			result := Chop(target, array)
			if !slices.Contains(expected, result) {

				t.Errorf("Array %v with target %d, expected one of %v. Got %d", array, target, expected, result)
			}

		})
		t.Run("Duplicates: should return index 0 or 1", func(t *testing.T) {
			target := 1
			array := []int{1, 1}
			expected := []int{0, 1}

			result := Chop(target, array)
			if !slices.Contains(expected, result) {

				t.Errorf("Array %v with target %d, expected one of %v. Got %d", array, target, expected, result)
			}

		})
		t.Run("Multiples of 10: should return 4", func(t *testing.T) {
			target := 50
			array := []int{10, 20, 30, 40, 50}
			expected := 4

			actual := Chop(target, array)
			if expected != actual {

				t.Errorf("Array %v with target %d, expected one of %v. Got %d", array, target, expected, actual)
			}

		})
		t.Run("Multiples of 10: should not find", func(t *testing.T) {
			target := 100
			array := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
			expected := -1

			actual := Chop(target, array)
			if expected != actual {

				t.Errorf("Array %v with target %d, expected one of %v. Got %d", array, target, expected, actual)
			}
		})

	})
}
