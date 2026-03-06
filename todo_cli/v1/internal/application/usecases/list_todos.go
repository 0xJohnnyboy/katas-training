package usecases

import (
	ports "todo_cli/internal/application/ports"
	domain "todo_cli/internal/domain"
)

type ListTodoUseCase struct {
	repo ports.TodoRepository
}

func NewListTodoUseCase(repo ports.TodoRepository) *ListTodoUseCase {
	return &ListTodoUseCase{
		repo,
	}
}

func (uc ListTodoUseCase) Execute() ([]domain.Todo, error) {
	size := 20

	todos, err := uc.repo.List()
	if err != nil {
		return []domain.Todo{}, err
	}

	if size >= len(todos) {
		return todos, nil
	}

	return todos[len(todos)-size:], nil
}
