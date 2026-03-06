package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	usecases "todo_cli/internal/application/usecases"
)

type CliCommand string

const (
	ADD    CliCommand = "add"
	UPDATE CliCommand = "update"
	DELETE CliCommand = "delete"
	LIST   CliCommand = "list"
	HELP   CliCommand = "help"
)

type CLI struct {
	addTodo    *usecases.CreateTodoUseCase
	updateTodo *usecases.UpdateTodoUseCase
	listTodo   *usecases.ListTodoUseCase
	deleteTodo *usecases.DeleteTodoUseCase
}

func NewCLI(
	addTodo *usecases.CreateTodoUseCase,
	updateTodo *usecases.UpdateTodoUseCase,
	listTodo *usecases.ListTodoUseCase,
	deleteTodo *usecases.DeleteTodoUseCase,
) *CLI {
	return &CLI{
		addTodo,
		updateTodo,
		listTodo,
		deleteTodo,
	}

}

func (c CLI) Run() error {
	cmd, options, err := c.getCliCommandAndOptions()

	if err != nil {
		printHelp()
		return err
	}

	switch *cmd {
	case ADD:
		err := c.addTodoHandler(options)
		if err != nil {
			return err
		}
	case UPDATE:
		err := c.updateTodoHandler(options)
		if err != nil {
			return err
		}
	case DELETE:
		err := c.deleteTodoHandler(options)
		if err != nil {
			return err
		}
	case LIST:
		err := c.listTodosHandler()
		if err != nil {
			return err
		}
	default:
		printHelp()
	}

	return nil
}

func getCreateTodoInputFromOptions(options []string) (*usecases.CreateTodoInput, error) {
	var done bool

	if options[3] == "true" {
		done = true
	}

	todoInput := &usecases.CreateTodoInput{
		Title:       options[0],
		Description: options[1],
		Date:        options[2],
		Done:        done,
	}
	return todoInput, nil
}
func getUpdateTodoInputFromOptions(options []string) (*usecases.UpdateTodoInput, error) {
	ID, err := strconv.Atoi(options[0])
	if err != nil {
		return nil, err
	}
	var done bool
	if options[3] == "true" {
		done = true
	}
	todoInput := &usecases.UpdateTodoInput{
		ID:          ID,
		Title:       options[1],
		Description: options[2],
		Done:        done,
	}
	return todoInput, nil
}

func (c CLI) getCliCommandAndOptions() (*CliCommand, []string, error) {
	args := os.Args

	if len(args) == 1 {
		return nil, []string{}, errors.New("Not enough parameters")
	}

	var cmd CliCommand

	switch args[1] {
	case "add":
		cmd = ADD
	case "update":
		cmd = UPDATE
	case "delete":
		cmd = DELETE
	case "list":
		cmd = LIST
	}

	return &cmd, args[2:], nil
}

func printHelp() {
	fmt.Println("Todo CLI")
	fmt.Println("Available commands:")
	fmt.Println("	- add <title> <description> <date: YYYY-MM-DD> <done: true|false>")
	fmt.Println("	- update <id> <title> <description> <done: true|false>")
	fmt.Println("	- delete <id>")
	fmt.Println("	- list")
}

func (c CLI) addTodoHandler(options []string) error {
	todoInput, err := getCreateTodoInputFromOptions(options)
	if err != nil {
		printHelp()
		return err
	}
	todo, err := c.addTodo.Execute(todoInput)
	if err != nil {
		return err
	}

	formattedDone := " "
	if todo.Done {
		formattedDone = "X"
	}
	msg := fmt.Sprintf("ID: %d | %s | [%s] -- %s | %s", todo.ID, todo.Date, formattedDone, todo.Title, todo.Description)
	fmt.Println(msg)
	return nil
}

func (c CLI) updateTodoHandler(options []string) error {
	todoInput, err := getUpdateTodoInputFromOptions(options)
	if err != nil {
		printHelp()
		return err
	}
	err = c.updateTodo.Execute(todoInput)
	if err != nil {
		return err
	}
	fmt.Printf("Updated todo <%d>\n", todoInput.ID)
	return nil
}
func (c CLI) deleteTodoHandler(options []string) error {
	ID, err := strconv.Atoi(options[0])
	if err != nil {
		return err
	}
	todo, err := c.deleteTodo.Execute(ID)
	if err != nil {
		return err
	}
	fmt.Printf("Deleted todo <%d>\n", todo.ID)
	return nil
}

func (c CLI) listTodosHandler() error {
	todos, err := c.listTodo.Execute()
	if err != nil {
		return err
	}
	for _, t := range todos {
		formattedDone := " "
		if t.Done {
			formattedDone = "X"
		}
		msg := fmt.Sprintf("ID: %d | %s | [%s] -- %s | %s", t.ID, t.Date, formattedDone, t.Title, t.Description)
		fmt.Println(msg)
	}
	return nil
}
