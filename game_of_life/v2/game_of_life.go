package v2

type Grid struct {
	Width  int
	Height int
	Cells  map[Position]bool
}

type Position struct {
	X int
	Y int
}

func NewGrid(width int, height int) *Grid {
	return &Grid{
		Width:  width,
		Height: height,
		Cells:  make(map[Position]bool),
	}
}

func (g *Grid) Toggle(x int, y int) {
	pos := Position{x, y}
	g.Cells[pos] = !g.Cells[pos]
}

func (g *Grid) IsAlive(x int, y int) bool {
	return g.Cells[Position{x, y}]
}

func (g *Grid) CountLiveNeighbours(x int, y int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}

			nx := x + i
			ny := y + j

			isCellWithinGridLimits := nx >= 0 && nx < g.Width && ny >= 0 && ny < g.Height

			if isCellWithinGridLimits && g.IsAlive(nx, ny) {
				count++
			}
		}
	}
	return count
}

func WillBeAlive(isAlive bool, liveNeighbours int) bool {
	if isAlive {
		return liveNeighbours == 2 || liveNeighbours == 3
	}

	return liveNeighbours == 3
}

func (g *Grid) Next() *Grid {
	newGrid := NewGrid(g.Width, g.Height)

	for x := range g.Width {
		for y := range g.Height {
			neighbours := g.CountLiveNeighbours(x, y)
			isAlive := g.IsAlive(x, y)
			if WillBeAlive(isAlive, neighbours) {
				newGrid.Toggle(x, y)
			}
		}
	}

	return newGrid
}
