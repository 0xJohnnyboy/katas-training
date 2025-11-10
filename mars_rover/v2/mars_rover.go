package v2

import (
	"fmt"
)

type Rover struct {
	X         int
	Y         int
	Direction Direction
	Grid      Grid
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

const (
	Forward  = 1
	Backward = -1
)

const (
	Left  Direction = 3
	Right Direction = 1
)

func (d *Direction) ToString() string {
	return directionsMap[*d]
}

func (r *Rover) Turn(side Direction) {
	r.Direction = (r.Direction + side) % 4
}

var directionsMap = map[Direction]string{
	North: "N",
	East:  "E",
	South: "S",
	West:  "W",
}

type Grid struct {
	Width     int
	Height    int
	Obstacles []Obstacle
}

func NewGrid(width int, height int, obstacles []Obstacle) Grid {
	return Grid{width, height, obstacles}
}

type Obstacle struct {
	X int
	Y int
}

func NewRover(x int, y int, dir Direction, grid Grid) *Rover {
	return &Rover{
		x,
		y,
		dir,
		grid,
	}
}

func (r *Rover) ToString() string {
	return fmt.Sprintf("%d:%d:%s", r.X, r.Y, r.Direction.ToString())
}

var cmdsMap map[rune]func(r *Rover) = map[rune]func(r *Rover){
	'F': func(r *Rover) { r.move(Forward) },
	'B': func(r *Rover) { r.move(Backward) },
	'L': func(r *Rover) { r.Turn(Left) },
	'R': func(r *Rover) { r.Turn(Right) },
}

func (r *Rover) Execute(cmds string) string {
	cmdSlice := []rune(cmds)

	for _, cmd := range cmdSlice {
		x, y := r.X, r.Y

		cmdsMap[cmd](r)

		if r.IsObstacle(r.X, r.Y) {
			r.X = x
			r.Y = y
			break
		}

	}

	r.wrap()
	return r.ToString()
}

func (r *Rover) wrap() {
	r.X = (r.X + r.Grid.Width) % r.Grid.Width
	r.Y = (r.Y + r.Grid.Height) % r.Grid.Height
}

var moveMap map[Direction]func(r *Rover, amount int) = map[Direction]func(r *Rover, amount int){
	North: func(r *Rover, amount int) { r.Y -= amount },
	East:  func(r *Rover, amount int) { r.X += amount },
	South: func(r *Rover, amount int) { r.Y += amount },
	West:  func(r *Rover, amount int) { r.X -= amount },
}

func (r *Rover) move(amount int) {
	moveMap[r.Direction](r, amount)
}

func (r *Rover) IsObstacle(x int, y int) bool {
	for _, o := range r.Grid.Obstacles {
		if o.X == x && o.Y == y {
			return true
		}
	}

	return false
}
