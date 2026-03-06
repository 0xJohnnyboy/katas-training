package usecases

import (
	"errors"
	"testing"
	domain "todo_cli/internal/domain"
)

type TestDeleteUsecase struct {
	desc           string
	input          int
	expectedTodo   *domain.Todo
	expectedLength int
	expectedError  error
}

func TestDelete(t *testing.T) {
	t.Parallel()
	repo := NewMockTodoRepository()
	todo1 := &domain.Todo{
		ID:          1,
		Date:        "2026-03-12",
		Title:       "test title",
		Description: "test description",
		Done:        false,
	}
	todo2 := &domain.Todo{
		ID:          2,
		Date:        "2026-04-12",
		Title:       "test title 2",
		Description: "test description 2",
		Done:        true,
	}
	todo3 := &domain.Todo{
		ID:          3,
		Date:        "2026-05-12",
		Title:       "test title 3",
		Description: "test description 3",
		Done:        false,
	}
	todos := []domain.Todo{*todo1, *todo2, *todo3}

	uc := NewDeleteTodoUseCase(repo)

	testCases := []TestDeleteUsecase{
		{"should delete todo for provided id", 2, todo2, len(todos) - 1, nil},
		{"should error trying to delete todo for out of range id", 11, nil, len(todos), errors.New("No todo for provided ID")},
	}

	for _, tc := range testCases {
		repo.todos = todos

		t.Run(tc.desc, func(t *testing.T) {
			actual, err := uc.Execute(tc.input)

			isErrorButNotExpected := err != nil && tc.expectedError == nil
			expectedErrorButGotNone := tc.expectedError != nil && err == nil
			expectedErrorButDoesntMatch := err != nil && tc.expectedError != nil && err.Error() != tc.expectedError.Error()

			if isErrorButNotExpected {
				t.Fatalf("Expected no error, got %s", err.Error())
				return
			}
			if expectedErrorButGotNone {
				t.Fatalf("Expected error %s, got none", tc.expectedError.Error())
				return
			}
			if expectedErrorButDoesntMatch {
				t.Fatalf("Expected error %s, got %s", tc.expectedError.Error(), err.Error())
				return
			}

			if tc.expectedLength != len(repo.todos) {
				t.Fatalf("Expected %d items, got %d", tc.expectedLength, len(repo.todos))
			}
			if err == nil {
				if actual.Date != tc.expectedTodo.Date {
					t.Fatalf("Expected Date %s, got %s", tc.expectedTodo.Date, actual.Date)
				}
				if actual.Title != tc.expectedTodo.Title {
					t.Fatalf("Expected Title %s, got %s", tc.expectedTodo.Title, actual.Title)
				}
				if actual.Description != tc.expectedTodo.Description {
					t.Fatalf("Expected Description %s, got %s", tc.expectedTodo.Description, actual.Description)
				}
				if actual.Done != tc.expectedTodo.Done {
					t.Fatalf("Expected Done %v, got %v", tc.expectedTodo.Done, actual.Done)
				}
			}
		})
	}
}
