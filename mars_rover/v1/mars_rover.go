package v1

import (
	"fmt"
)

type Rover struct {
	Position  Coordinates
	Direction Direction
	Grid      Grid
}

func NewRover(x int, y int, d Direction, grid Grid) *Rover {
	return &Rover{Position: Coordinates{X: x, Y: y}, Direction: d, Grid: grid}
}

func (r *Rover) String() string {
	return fmt.Sprintf("%d:%d:%s", r.Position.X, r.Position.Y, r.Direction.String())
}

func (r *Rover) ToString() string {
	var symbols = map[Direction]string{
		North: "^",
		East:  ">",
		South: "v",
		West:  "<",
	}

	str := ""
	obstacle := "X"
	empty := "-"

	for y := range r.Grid.Size.Y {
		for x := range r.Grid.Size.X {
			if r.Grid.HasObstacle(Coordinates{X: x, Y: y}) {
				str += obstacle
				continue
			}

			if x == r.Position.X && y == r.Position.Y {
				str += symbols[r.Direction]
				continue
			}

			str += empty
		}

		if y < r.Grid.Size.Y-1 {
			str += "\n"
		}
	}

	return str
}

func (r *Rover) Execute(command string) string {
	cmdSlice := []rune(command)

	for _, cmd := range cmdSlice {
		state := r.Position
		actionsMap := map[rune]func(){
			'L': func() { r.Direction = r.Direction.TurnLeft() },
			'R': func() { r.Direction = r.Direction.TurnRight() },
			'F': func() { r.move(Forward) },
			'B': func() { r.move(Backward) },
		}

		actionsMap[cmd]()

		if r.Grid.HasObstacle(r.Position) {
			r.Position = state
			break
		}
	}

	r.wrapCoordinates()

	return r.String()
}

func (r *Rover) wrapCoordinates() {
	r.Position.Y = ((r.Position.Y % r.Grid.Size.Y) + r.Grid.Size.Y) % r.Grid.Size.Y
	r.Position.X = ((r.Position.X % r.Grid.Size.X) + r.Grid.Size.X) % r.Grid.Size.X
}

func (r *Rover) move(amount int) {
	moveMap := map[Direction]func(int){
		North: func(amount int) { r.Position.Y -= amount },
		East:  func(amount int) { r.Position.X += amount },
		South: func(amount int) { r.Position.Y += amount },
		West:  func(amount int) { r.Position.X -= amount },
	}

	moveMap[r.Direction](amount)
}
