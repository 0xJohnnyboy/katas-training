package main

import (
	"todo_cli/cmd"
	"todo_cli/internal/application/usecases"
	"todo_cli/internal/infrastructure"
)

func main() {
	repo := infrastructure.NewCsv()

	addTodoUseCase := usecases.NewCreateTodoUseCase(repo)
	updateTodoUseCase := usecases.NewUpdateTodoUseCase(repo)
	listTodoUseCase := usecases.NewListTodoUseCase(repo)
	deleteTodoUseCase := usecases.NewDeleteTodoUseCase(repo)

	CLI := cmd.NewCLI(
		addTodoUseCase,
		updateTodoUseCase,
		listTodoUseCase,
		deleteTodoUseCase,
	)

	err := CLI.Run()

	if err != nil {
		panic(err.Error())
	}
}
