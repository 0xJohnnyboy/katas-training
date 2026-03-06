package domain

import (
	"errors"
	"strings"
	"time"
)

type Todo struct {
	ID          int
	Date        string
	Title       string
	Description string
	Done        bool
}

func NewTodo(ID int, title string, description string, date string, done bool) (*Todo, error) {
	if strings.TrimSpace(title) == "" {
		return nil, errors.New("Missing title")
	}
	if strings.TrimSpace(date) == "" {
		return nil, errors.New("Missing date")
	}
	if !isValidDate(date) {
		return nil, errors.New("Invalid date format")
	}
	return &Todo{ID, date, title, description, done}, nil
}

func isValidDate(s string) bool {
	if strings.TrimSpace(s) == "" {
		return false
	}
	_, err := time.Parse("2006-01-02", s)
	return err == nil
}
