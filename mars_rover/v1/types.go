package v1

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
	Size      Coordinates
	Obstacles []Obstacle
}

type Obstacle struct {
	Position Coordinates
}

func (g Grid) HasObstacle(p Coordinates) bool {
	for _, o := range g.Obstacles {
		if o.Position.Equals(p) {
			return true
		}
	}
	return false
}
