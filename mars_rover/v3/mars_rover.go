package v3

import (
	"fmt"
)

type Grid struct {
	Width     int
	Height    int
	Obstacles []Position
}

func NewGrid(width int, height int) Grid {
	return Grid{
		Width:  width,
		Height: height,
	}
}

type Direction int

var directionsMap map[Direction]string = map[Direction]string{
	North: "N",
	East:  "E",
	South: "S",
	West:  "W",
}

func (d *Direction) ToString() string {
	return directionsMap[*d]
}

const (
	North Direction = iota
	East
	South
	West
)

type Move int

const (
	Backward Move = -1
	Forward  Move = 1
)

type Position struct {
	X int
	Y int
}

func (p *Position) Equals(pos Position) bool {
	return p.X == pos.X && p.Y == pos.Y
}

type Rover struct {
	Position  Position
	Direction Direction
	Grid      Grid
}

func (r *Rover) ToString() string {
	return fmt.Sprintf("%d:%d:%s", r.Position.X, r.Position.Y, r.Direction.ToString())
}

func (r *Rover) IsObstacle(pos Position) bool {
	for _, o := range r.Grid.Obstacles {
		if o.X == pos.X && o.Y == pos.Y {
			return true
		}
	}

	return false
}

var actionsMap map[rune]func(r *Rover) = map[rune]func(r *Rover){
	'L': func(r *Rover) {
		r.Direction = (r.Direction + 3) % 4
	},
	'R': func(r *Rover) {
		r.Direction = (r.Direction + 1) % 4
	},
	'F': func(r *Rover) {
		r.Move(Forward)
	},
	'B': func(r *Rover) {
		r.Move(Backward)
	},
}

var moveMap map[Direction]func(r *Rover, m Move) = map[Direction]func(r *Rover, m Move){
	North: func(r *Rover, m Move) {
		r.Position.Y -= int(m)
	},
	East: func(r *Rover, m Move) {
		r.Position.X += int(m)
	},
	South: func(r *Rover, m Move) {
		r.Position.Y += int(m)
	},
	West: func(r *Rover, m Move) {
		r.Position.X -= int(m)
	},
}

func (r *Rover) Move(m Move) {
	moveMap[r.Direction](r, m)
}

func (r *Rover) Execute(commands string) string {
	cmdsSlice := []rune(commands)
	for _, cmd := range cmdsSlice {
		pos := r.Position
		actionsMap[cmd](r)

		if r.IsObstacle(r.Position) {
			r.Position = pos
			break
		}
	}

	r.wrap()

	return r.ToString()
}

func (r *Rover) wrap() {
	r.Position.X = (r.Position.X + r.Grid.Width) % r.Grid.Width
	r.Position.Y = (r.Position.Y + r.Grid.Height) % r.Grid.Height
}

func NewRover(width int, height int, dir Direction, grid Grid) *Rover {
	return &Rover{
		Position:  Position{width, height},
		Direction: dir,
		Grid:      grid,
	}
}
