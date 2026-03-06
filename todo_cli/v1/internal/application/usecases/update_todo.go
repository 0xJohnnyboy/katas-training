package usecases

import (
	ports "todo_cli/internal/application/ports"
	domain "todo_cli/internal/domain"
)

type UpdateTodoUseCase struct {
	repo ports.TodoRepository
}

type UpdateTodoInput struct {
	ID          int
	Title       string
	Description string
	Done        bool
}

func NewUpdateTodoUseCase(repo ports.TodoRepository) *UpdateTodoUseCase {
	return &UpdateTodoUseCase{
		repo,
	}
}

func (uc UpdateTodoUseCase) Execute(input *UpdateTodoInput) error {

	todo, err := uc.repo.Get(input.ID)
	if err != nil {
		return err
	}

	todo, err = domain.NewTodo(todo.ID, input.Title, input.Description, todo.Date, input.Done)

	if err != nil {
		return err
	}

	return uc.repo.Update(todo)
}
