package mars_rover

import (
	"fmt"
)

type Coordinates struct {
	X int
	Y int
}

func (c Coordinates) Equals(other Coordinates) bool {
	return c.X == other.X && c.Y == other.Y
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"N", "E", "S", "W"}[d]
}

func (d Direction) TurnLeft() Direction {
	return Direction((int(d) + 3) % 4)
}

func (d Direction) TurnRight() Direction {
	return Direction((int(d) + 1) % 4)
}

const (
	Backward int = -1
	Forward  int = 1
)

type Grid struct {
	Size Coordinates
	Obstacles []Obstacle
}

type Rover struct {
	Position Coordinates
	Direction Direction
	Grid      Grid
}

type Obstacle struct {
	Position Coordinates
}


func NewRover(x int, y int, d Direction, grid Grid) *Rover {
	return &Rover{Position: Coordinates{X: x, Y: y}, Direction: d, Grid: grid}
}


func (r *Rover) String() string {
	return fmt.Sprintf("%d:%d:%s", r.Position.X, r.Position.Y, r.Direction.String())
}

func (r *Rover) Execute(command string) string {
	cmdSlice := []rune(command)

	for _, cmd := range cmdSlice {
		switch cmd {
		case 'L':
			r.Direction = r.Direction.TurnLeft()
		case 'R':
			r.Direction = r.Direction.TurnRight()
		case 'F':
			r.Move(Forward)
		case 'B':
			r.Move(Backward)
		}
	}

	r.wrap()

	return r.String()
}

func (r *Rover) wrap() {
	r.Position.Y = ((r.Position.Y % r.Grid.Size.Y) + r.Grid.Size.Y) % r.Grid.Size.Y
	r.Position.X = ((r.Position.X % r.Grid.Size.X) + r.Grid.Size.X) % r.Grid.Size.X
}

func (r *Rover) Move(amount int) {
	switch r.Direction {
	case North:
		r.Position.Y -= amount
	case East:
		r.Position.X += amount
	case South:
		r.Position.Y += amount
	case West:
		r.Position.X -= amount
	}
}
