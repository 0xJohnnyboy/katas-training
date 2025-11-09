package main

import (
	"fmt"
	"time"

	gol "katas/game_of_life"
)

func main() {
	fmt.Print("\033[H\033[2J") // clear term

	grid := gol.NewGrid(10, 10)
	positions := []gol.Position{
		{X: 1, Y: 0},
		{X: 2, Y: 1},
		{X: 0, Y: 2},
		{X: 1, Y: 2},
		{X: 2, Y: 2},
	}

	for _, p := range positions {
		grid.Toggle(p.X, p.Y)
	}

	render(grid)

	for range 10 {
		grid = grid.NextGeneration()
		render(grid)
	}
}

func render(grid *gol.Grid) {
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("\033[%dA", grid.Height)
	fmt.Println(grid.ToString())
}
