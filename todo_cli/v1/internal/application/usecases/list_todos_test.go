package usecases

import (
	"testing"
	domain "todo_cli/internal/domain"
)

type TestListUsecase struct {
	desc           string
	input          int
	expectedLength int
}

func TestList(t *testing.T) {
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

	repoTodos := []domain.Todo{*todo1, *todo2, *todo3}

	uc := NewListTodoUseCase(repo)

	t.Run("Should return empty list ", func(t *testing.T) {
		todos, err := uc.Execute()

		if err != nil {
			t.Fatalf("Expected no error, got %s", err.Error())
		}

		if len(todos) != 0 {
			t.Fatalf("Expected 0 todos, got %d", len(todos))
		}
	})

	t.Run("Should list with 3 elements ", func(t *testing.T) {
		repo.todos = repoTodos

		todos, err := uc.Execute()

		if err != nil {
			t.Fatalf("Expected no error, got %s", err.Error())
		}

		if len(todos) != len(repoTodos) {
			t.Fatalf("Expected %d todos, got %d", len(repoTodos), len(todos))
		}
	})

	t.Run("Should only list 20 last elements ", func(t *testing.T) {
		repo.todos = repoTodos
		for range 5 {
			repo.todos = append(repo.todos, repo.todos...)
		}

		todos, err := uc.Execute()

		if err != nil {
			t.Fatalf("Expected no error, got %s", err.Error())
		}

		if len(todos) != 20 {
			t.Fatalf("Expected 20 todos, got %d", len(todos))
		}
	})
}
