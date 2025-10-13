package string_calculator

import (
	"testing"
	"errors"
)

type TestCase struct {
	desc     string
	input    string
	expected int
	error error
}

var testCases = []TestCase{
	{"empty string", "", 0, nil},
	{"single number", "1", 1, nil},
	{"two numbers", "2,3", 5, nil},
	{"four numbers", "1,2,5,10", 18, nil},
	{"with newlines", "1\n2,3", 6, nil},
	{"error if negative", "-1", 0, ErrNegativeNumber},
	{"error if negative with multiple numbers", "1,2,3,-4", 0, ErrNegativeNumber},
}

func TestAdd(t *testing.T) {
	for _, testCase := range testCases {

		t.Run(testCase.desc, func(t *testing.T) {
			actual, err := Add(testCase.input)

			if !errors.Is(err, testCase.error) {
				t.Errorf("Add(%s) expected error %v, got %v", testCase.input, testCase.error, err)
			}

			if actual != testCase.expected {
				t.Errorf("Add(%s) expected %d, got %d", testCase.input, testCase.expected, actual)
			}
		})
	}
}
