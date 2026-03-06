package ports

import (
	domain "todo_cli/internal/domain"
)

type TodoRepository interface {
	Add(todo *domain.Todo) (*domain.Todo, error)
	Update(todo *domain.Todo) error
	List() ([]domain.Todo, error)
	Delete(ID int) (*domain.Todo, error)
	Get(ID int) (*domain.Todo, error)
}
