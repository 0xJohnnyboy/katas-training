package infrastructure

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"strings"
	domain "todo_cli/internal/domain"
)

type Csv struct {
	path string
}

func NewCsv() *Csv {
	path := os.Getenv("TODO_CSV")
	if strings.TrimSpace(path) == "" {
		path = "todos.csv"
	}

	return &Csv{path}
}

func (c Csv) Get(ID int) (*domain.Todo, error) {
	todos, err := c.read()
	if err != nil {
		for _, t := range todos {
			if t.ID == ID {
				return &t, nil
			}
		}
	}
	return nil, errors.New("No todo found for provided ID")
}

func (c Csv) Add(todo *domain.Todo) (*domain.Todo, error) {
	todos, err := c.read()
	if err != nil {
		return nil, err
	}

	nextID := 1
	for i := range todos {
		if todos[i].ID >= nextID {
			nextID = todos[i].ID + 1
		}
	}
	todo.ID = nextID
	todos = append(todos, *todo)
	if err := c.save(todos); err != nil {
		return todo, err
	}
	return todo, nil
}

func (c Csv) Update(todo *domain.Todo) error {
	todos, err := c.read()
	if err != nil {
		return err
	}
	pos := -1
	for i := range todos {
		if todos[i].ID == todo.ID {
			pos = i
			break
		}
	}
	if pos == -1 {
		return errors.New("todo id out of range")
	}

	item := todos[pos]
	item.Title = todo.Title
	item.Description = todo.Description
	item.Done = todo.Done
	todos[pos] = item

	if err := c.save(todos); err != nil {
		return err
	}
	return nil
}

func (c Csv) List() ([]domain.Todo, error) {
	todos, err := c.read()
	if err != nil {
		return []domain.Todo{}, err
	}

	return todos, nil
}

func (c Csv) Delete(ID int) (*domain.Todo, error) {
	todos, err := c.read()
	if err != nil {
		return nil, err
	}
	pos := -1
	for i := 0; i < len(todos); i++ {
		if todos[i].ID == ID {
			pos = i
			break
		}
	}
	if pos == -1 {
		return nil, errors.New("No todo for provided ID")
	}
	todo := todos[pos]
	todos = append(todos[:pos], todos[pos+1:]...)
	if err := c.save(todos); err != nil {
		return nil, err
	}
	return &todo, nil
}

func (c Csv) save(todos []domain.Todo) error {
	f, err := os.Create(c.path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	for i := range todos {
		t := todos[i]
		if err := w.Write([]string{strconv.Itoa(t.ID), t.Date, t.Title, t.Description, strconv.FormatBool(t.Done)}); err != nil {
			return err
		}
	}
	w.Flush()
	return w.Error()
}

func (c Csv) read() ([]domain.Todo, error) {
	f, err := os.Open(c.path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []domain.Todo{}, nil
		}
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	recs, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	out := make([]domain.Todo, 0, len(recs))
	fallbackID := 1
	for _, rec := range recs {
		if len(rec) < 4 {
			continue
		}
		if len(rec) >= 5 {
			id, err := strconv.Atoi(rec[0])
			if err != nil || id < 1 {
				continue
			}
			d, _ := parseDone(rec[4])
			out = append(out, domain.Todo{
				ID:          id,
				Date:        rec[1],
				Title:       rec[2],
				Description: rec[3],
				Done:        d,
			})
			continue
		}
		d, _ := parseDone(rec[3])
		out = append(out, domain.Todo{
			ID:          fallbackID,
			Date:        rec[0],
			Title:       rec[1],
			Description: rec[2],
			Done:        d,
		})
		fallbackID++
	}
	return out, nil
}

func parseDone(s string) (bool, error) {
	x := strings.ToLower(strings.TrimSpace(s))
	if x == "true" || x == "1" || x == "yes" || x == "y" {
		return true, nil
	}
	if x == "false" || x == "0" || x == "no" || x == "n" {
		return false, nil
	}
	return false, errors.New("completed must be true or false")
}
