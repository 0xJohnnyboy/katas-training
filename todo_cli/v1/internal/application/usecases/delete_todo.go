package usecases

import (
	ports "todo_cli/internal/application/ports"
	domain "todo_cli/internal/domain"
)

type DeleteTodoUseCase struct {
	repo ports.TodoRepository
}

func NewDeleteTodoUseCase(repo ports.TodoRepository) *DeleteTodoUseCase {
	return &DeleteTodoUseCase{
		repo,
	}
}

func (uc DeleteTodoUseCase) Execute(ID int) (*domain.Todo, error) {

	return uc.repo.Delete(ID)
}
