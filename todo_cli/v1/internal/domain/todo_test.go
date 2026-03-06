package domain

import (
	"errors"
	"testing"
)

type TestCaseInput struct {
	ID    int
	Title string
	Desc  string
	Date  string
	Done  bool
}
type NewTodoTestCase struct {
	desc          string
	input         TestCaseInput
	expectedError error
}

func TestTodo(t *testing.T) {
	testCases := []NewTodoTestCase{
		{"should create todo", TestCaseInput{1, "test title", "test desc", "2026-03-10", false}, nil},
		{"should create todo with empty desc", TestCaseInput{1, "test title", "", "2026-03-10", false}, nil},
		{"should create todo with empty desc filled with spaces", TestCaseInput{1, "test title", "     ", "2026-03-10", false}, nil},
		{"should create todo with done true", TestCaseInput{1, "test title", "test desc", "2026-03-10", false}, nil},
		{"should error creating todo with empty title", TestCaseInput{1, "", "test desc", "2026-03-10", false}, errors.New("Missing title")},
		{"should error creating todo with empty title filled with spaces", TestCaseInput{1, "    ", "test desc", "2026-03-10", false}, errors.New("Missing title")},
		{"should error creating todo with empty date", TestCaseInput{1, "test title", "test desc", "", false}, errors.New("Missing date")},
		{"should error creating todo with empty date filled with spaces", TestCaseInput{1, "test title", "test desc", "    ", false}, errors.New("Missing date")},
		{"should error creating todo with wrong date format", TestCaseInput{1, "test title", "test desc", "12/03/2026", false}, errors.New("Invalid date format")},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()

			todo, err := NewTodo(tc.input.ID, tc.input.Title, tc.input.Desc, tc.input.Date, tc.input.Done)

			isErrorButNotExpected := err != nil && tc.expectedError == nil
			expectedErrorButGotNone := tc.expectedError != nil && err == nil
			expectedErrorButDoesntMatch := err != nil && tc.expectedError != nil && err.Error() != tc.expectedError.Error()

			if isErrorButNotExpected {
				t.Fatalf("Expected no error, got %s", err.Error())
			}
			if expectedErrorButGotNone {
				t.Fatalf("Expected error %s, got none", tc.expectedError.Error())
			}
			if expectedErrorButDoesntMatch {
				t.Fatalf("Expected error %s, got %s", tc.expectedError.Error(), err.Error())
			}

			if err != nil {
				return
			}

			todoMatchesInput := tc.input.ID == todo.ID && tc.input.Title == todo.Title && tc.input.Desc == todo.Description && tc.input.Date == todo.Date && tc.input.Done == todo.Done
			if !todoMatchesInput {
				t.Fatalf("Expected todo to match %v, got %v", tc.input, todo)
			}
		})
	}

}
