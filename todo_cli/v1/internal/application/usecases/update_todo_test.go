package usecases

import (
	"errors"
	"testing"
	domain "todo_cli/internal/domain"
)

type TestUpdateUsecase struct {
	desc          string
	input         *UpdateTodoInput
	expectedError error
}

func TestUpdate(t *testing.T) {
	t.Parallel()
	repo := NewMockTodoRepository()
	todo := &domain.Todo{
		ID:          1,
		Date:        "2026-03-12",
		Title:       "test title",
		Description: "test description",
		Done:        false,
	}
	uc := NewUpdateTodoUseCase(repo)

	testCases := []TestUpdateUsecase{
		{"should update title", &UpdateTodoInput{1, "updated title", "test description", false}, nil},
		{"should error updating title with empty", &UpdateTodoInput{1, "", "test description", true}, errors.New("Missing title")},
		{"should error updating title with empty filled with space", &UpdateTodoInput{1, "   ", "test description", true}, errors.New("Missing title")},
		{"should update description", &UpdateTodoInput{1, "test title", "updated description", false}, nil},
		{"should update done", &UpdateTodoInput{1, "test title", "test description", true}, nil},
		{"should error updating non existent ID", &UpdateTodoInput{4, "test title", "test description", true}, errors.New("No todo found for provided ID")},
	}

	for _, tc := range testCases {
		repo.todos = append([]domain.Todo{}, *todo)

		t.Run(tc.desc, func(t *testing.T) {
			err := uc.Execute(tc.input)

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

			if err == nil {
				actualTodo := repo.todos[0]
				if actualTodo.Date != todo.Date {
					t.Fatalf("Expected Date %s, got %s", todo.Date, actualTodo.Date)
				}
				if actualTodo.Title != tc.input.Title {
					t.Fatalf("Expected Title %s, got %s", tc.input.Title, actualTodo.Title)
				}
				if actualTodo.Description != tc.input.Description {
					t.Fatalf("Expected Description %s, got %s", tc.input.Description, actualTodo.Description)
				}
				if actualTodo.Done != tc.input.Done {
					t.Fatalf("Expected Done %v, got %v", tc.input.Done, actualTodo.Done)
				}
			}
		})
	}
}
