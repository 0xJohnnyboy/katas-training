package usecases

import (
	"errors"
	"testing"
	domain "todo_cli/internal/domain"
)

type TestAddUsecase struct {
	desc           string
	inputs         []CreateTodoInput
	expectedLength int
	expectedIds    []int
	expectedError  error
}

func TestAdd(t *testing.T) {
	t.Parallel()
	repo := NewMockTodoRepository()
	uc := NewCreateTodoUseCase(repo)

	testCases := []TestAddUsecase{
		{"should add 1", []CreateTodoInput{{"Test", "Description", "2026-03-09", false}}, 1, []int{1}, nil},
		{"should add 2", []CreateTodoInput{{"Test", "Description", "2026-03-09", false}, {"Test2", "Description2", "2026-03-09", true}}, 2, []int{1, 2}, nil},
		{"should error because empty title", []CreateTodoInput{{"", "Description", "2026-03-09", false}}, 0, []int{}, errors.New("Missing title")},
		{"should error because empty title filled with spaces", []CreateTodoInput{{"   ", "Description", "2026-03-09", false}}, 0, []int{}, errors.New("Missing title")},
		{"should error because empty date", []CreateTodoInput{{"Test", "Description", "", false}}, 0, []int{}, errors.New("Missing date")},
		{"should error because empty date filled with spaces", []CreateTodoInput{{"Test", "Description", "     ", false}}, 0, []int{}, errors.New("Missing date")},
		{"should error because invalid date format", []CreateTodoInput{{"Test", "Description", "12/03/2024", false}}, 0, []int{}, errors.New("Invalid date format")},
	}

	for _, tc := range testCases {
		repo.todos = []domain.Todo{}
		t.Run(tc.desc, func(t *testing.T) {
			for idx, input := range tc.inputs {
				todo, err := uc.Execute(&input)

				isErrorButNotExpected := err != nil && tc.expectedError == nil
				expectedErrorButGotNone := tc.expectedError != nil && err == nil
				expecteErrorButDoesntMatch := err != nil && tc.expectedError != nil && err.Error() != tc.expectedError.Error()

				if isErrorButNotExpected {
					t.Fatalf("Expected no error, got %s", err.Error())
					return
				}
				if expectedErrorButGotNone {
					t.Fatalf("Expected error %s, got none", tc.expectedError.Error())
					return
				}
				if expecteErrorButDoesntMatch {
					t.Fatalf("Expected error %s, got %s", tc.expectedError.Error(), err.Error())
					return
				}
				if err == nil && todo.ID != tc.expectedIds[idx] {
					t.Fatalf("Expected ID: %d, got %d", tc.expectedIds[idx], todo.ID)
				}
			}

			if tc.expectedLength != len(repo.todos) {
				t.Fatalf("Expected %d todos, got %d", tc.expectedLength, len(repo.todos))

			}
		})
	}
}
