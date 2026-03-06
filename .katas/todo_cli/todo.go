package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
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

func main() {
	csvPath := os.Getenv("TODO_CSV")
	if strings.TrimSpace(csvPath) == "" {
		csvPath = "todos.csv"
	}

	out, err := run(os.Args[1:], csvPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	if strings.TrimSpace(out) != "" {
		fmt.Println(out)
	}
}

// NOTE: update only changes title/description, keep date as is.
func run(args []string, csvPath string) (string, error) {
	if len(args) == 0 {
		return usage(), nil
	}

	cmd := strings.ToLower(strings.TrimSpace(args[0]))

	if cmd == "add" {
		if len(args) < 4 {
			return "", errors.New("add needs: add <YYYY-MM-DD> <title> <description>")
		}
		date := strings.TrimSpace(args[1])
		title := strings.TrimSpace(args[2])
		desc := strings.TrimSpace(strings.Join(args[3:], " "))
		if !isValidDate(date) {
			return "", errors.New("invalid date format, expected YYYY-MM-DD")
		}
		if title == "" {
			return "", errors.New("title cannot be empty")
		}
		todos, err := readTodos(csvPath)
		if err != nil {
			return "", err
		}
		nextID := 1
		for i := 0; i < len(todos); i++ {
			if todos[i].ID >= nextID {
				nextID = todos[i].ID + 1
			}
		}
		todos = append(todos, Todo{ID: nextID, Date: date, Title: title, Description: desc, Done: false})
		if err := writeTodos(csvPath, todos); err != nil {
			return "", err
		}
		return "todo added", nil
	}

	if cmd == "update" {
		if len(args) < 6 {
			return "", errors.New("update needs: update <id> <YYYY-MM-DD> <title> <description> <true|false>")
		}
		idRaw := strings.TrimSpace(args[1])
		date := strings.TrimSpace(args[2])
		title := strings.TrimSpace(args[3])
		doneRaw := strings.TrimSpace(args[len(args)-1])
		desc := strings.TrimSpace(strings.Join(args[4:len(args)-1], " "))

		id, err := strconv.Atoi(idRaw)
		if err != nil || id < 1 {
			return "", errors.New("id must be a positive integer")
		}
		if !isValidDate(date) {
			return "", errors.New("invalid date format, expected YYYY-MM-DD")
		}
		doneVal, err := parseDone(doneRaw)
		if err != nil {
			return "", err
		}

		todos, err := readTodos(csvPath)
		if err != nil {
			return "", err
		}
		pos := -1
		for i := 0; i < len(todos); i++ {
			if todos[i].ID == id {
				pos = i
				break
			}
		}
		if pos == -1 {
			return "", errors.New("todo id out of range")
		}

		item := todos[pos]
		item.Date = date
		item.Title = title
		item.Description = desc
		item.Done = doneVal
		todos[pos] = item

		if err := writeTodos(csvPath, todos); err != nil {
			return "", err
		}
		return "todo updated", nil
	}

	if cmd == "delete" {
		if len(args) != 2 {
			return "", errors.New("delete needs: delete <id>")
		}
		id, err := strconv.Atoi(strings.TrimSpace(args[1]))
		if err != nil || id < 1 {
			return "", errors.New("id must be a positive integer")
		}
		todos, err := readTodos(csvPath)
		if err != nil {
			return "", err
		}
		pos := -1
		for i := 0; i < len(todos); i++ {
			if todos[i].ID == id {
				pos = i
				break
			}
		}
		if pos == -1 {
			return "", errors.New("todo id out of range")
		}
		todos = append(todos[:pos], todos[pos+1:]...)
		if err := writeTodos(csvPath, todos); err != nil {
			return "", err
		}
		return "todo deleted", nil
	}

	if cmd == "list" {
		todos, err := readTodos(csvPath)
		if err != nil {
			return "", err
		}
		if len(todos) == 0 {
			return "No todos.", nil
		}
		var out []string
		start := 0
		if len(todos) > 20 {
			start = len(todos) - 20
		}
		for i := start; i < len(todos); i++ {
			status := "[ ]"
			if todos[i].Done {
				status = "[x]"
			}
			out = append(out, fmt.Sprintf("%d | %s | %s | %s | %s", todos[i].ID, status, todos[i].Date, todos[i].Title, todos[i].Description))
		}
		return strings.Join(out, "\n"), nil
	}

	if cmd == "help" {
		return usage(), nil
	}

	return "", errors.New("unknown command: " + cmd)
}

func usage() string {
	return strings.TrimSpace(`todo_cli commands:
  add <YYYY-MM-DD> <title> <description>
  update <id> <YYYY-MM-DD> <title> <description> <true|false>
  delete <id>
  list
  help`)
}

func isValidDate(s string) bool {
	if strings.TrimSpace(s) == "" {
		return false
	}
	_, err := time.Parse("2006-01-02", s)
	return err == nil
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

func readTodos(csvPath string) ([]Todo, error) {
	f, err := os.Open(csvPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Todo{}, nil
		}
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	recs, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	out := make([]Todo, 0, len(recs))
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
			out = append(out, Todo{
				ID:          id,
				Date:        rec[1],
				Title:       rec[2],
				Description: rec[3],
				Done:        d,
			})
			continue
		}
		d, _ := parseDone(rec[3])
		out = append(out, Todo{
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

func writeTodos(csvPath string, todos []Todo) error {
	f, err := os.Create(csvPath)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	for i := 0; i < len(todos); i++ {
		t := todos[i]
		if err := w.Write([]string{strconv.Itoa(t.ID), t.Date, t.Title, t.Description, strconv.FormatBool(t.Done)}); err != nil {
			return err
		}
	}
	w.Flush()
	return w.Error()
}
