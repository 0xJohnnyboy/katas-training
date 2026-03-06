package usecases

import (
	"errors"
	domain "todo_cli/internal/domain"
)

type MockTodoRepository struct {
	todos []domain.Todo
}

func NewMockTodoRepository() *MockTodoRepository {
	return &MockTodoRepository{[]domain.Todo{}}
}

func (m *MockTodoRepository) Add(todo *domain.Todo) (*domain.Todo, error) {
	ID := 1

	if len(m.todos) == 0 {
		todo.ID = ID
		m.todos = append(m.todos, *todo)
		return todo, nil
	}

	lastID := m.todos[len(m.todos)-1].ID
	todo.ID = lastID + 1
	m.todos = append(m.todos, *todo)
	return todo, nil
}
func (m *MockTodoRepository) Update(todo *domain.Todo) error {
	for i := range m.todos {
		if m.todos[i].ID == todo.ID {
			m.todos[i] = *todo
			return nil
		}
	}
	return nil
}
func (m *MockTodoRepository) Delete(ID int) (*domain.Todo, error) {
	for i, t := range m.todos {
		if t.ID == ID {
			todo := t
			m.todos = append(m.todos[:i], m.todos[i+1:]...)
			return &todo, nil
		}
	}
	return nil, errors.New("No todo for provided ID")
}
func (m *MockTodoRepository) List() ([]domain.Todo, error) {
	return m.todos, nil
}

func (m *MockTodoRepository) Get(ID int) (*domain.Todo, error) {
	for _, t := range m.todos {
		if t.ID == ID {
			return &t, nil
		}
	}
	return nil, errors.New("No todo found for provided ID")
}
