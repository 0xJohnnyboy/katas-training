package mars_rover

import (
	"fmt"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type Rover struct {
	X         int
	Y         int
	Direction Direction
}

func NewRover(x int, y int, d Direction) *Rover {
	return &Rover{X: x, Y: y, Direction: d}
}

func (d Direction) String() string {
	return [...]string{"N", "E", "S", "W"}[d]
}

func (d Direction) TurnLeft() Direction {
	return Direction((int(d) + 3) % 4)
}

func (d Direction) TurnRight() Direction {
	return Direction((int(d) + 1) % 4)
}

func (r *Rover) String() string {
	return fmt.Sprintf("%d:%d:%s", r.X, r.Y, r.Direction.String())
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
			r.Move(1)
		case 'B':
			r.Move(-1)
		}
	}

	return r.String()
}

func (r *Rover) Move(amount int) {
	switch r.Direction {
	case North:
		r.Y -= amount
	case East:
		r.X += amount
	case South:
		r.Y += amount
	case West:
		r.X -= amount
	}
}
