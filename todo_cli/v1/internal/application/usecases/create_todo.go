package usecases

import (
	ports "todo_cli/internal/application/ports"
	domain "todo_cli/internal/domain"
)

type CreateTodoUseCase struct {
	repo ports.TodoRepository
}

type CreateTodoInput struct {
	Title       string
	Description string
	Date        string
	Done        bool
}

func NewCreateTodoUseCase(repo ports.TodoRepository) *CreateTodoUseCase {
	return &CreateTodoUseCase{
		repo,
	}
}

func (uc CreateTodoUseCase) Execute(input *CreateTodoInput) (*domain.Todo, error) {

	todo, err := domain.NewTodo(0, input.Title, input.Description, input.Date, input.Done)

	if err != nil {
		return nil, err
	}

	return uc.repo.Add(todo)
}
