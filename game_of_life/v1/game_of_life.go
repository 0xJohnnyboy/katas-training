package v1

type Position struct {
	X int
	Y int
}

type Grid struct {
	Width  int
	Height int
	Cells  map[Position]bool
}

func NewGrid(width int, height int) *Grid {
	return &Grid{
		Width:  width,
		Height: height,
		Cells:  make(map[Position]bool),
	}
}

func (g *Grid) Toggle(x int, y int) {
	g.Cells[Position{x, y}] = !g.Cells[Position{x, y}]
}

func (g *Grid) IsAlive(x int, y int) bool {
	return g.Cells[Position{x, y}]
}

func (g *Grid) CountLiveNeighbours(x int, y int) int {
	count := 0

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue // skip self
			}

			// cell position to check
			nx := x + dx
			ny := y + dy

			isCellWithinGridLimits := nx >= 0 && nx < g.Width && ny >= 0 && ny < g.Height

			if isCellWithinGridLimits {
				if g.IsAlive(nx, ny) {
					count++
				}
			}
		}
	}

	return count
}

func (g *Grid) NextGeneration() *Grid {
	next := NewGrid(g.Width, g.Height)

	for x := 0; x <= g.Width; x++ {
		for y := 0; y < g.Height; y++ {
			liveNeighbours := g.CountLiveNeighbours(x, y)
			isAlive := g.IsAlive(x, y)

			if WillBeAlive(isAlive, liveNeighbours) {
				next.Toggle(x, y)
			}
		}
	}

	return next
}

func (g *Grid) ToString() string {
	dead := "⬛"
	alive := "🔴"

	var str string

	for y := range g.Height {
		for x := range g.Width {
			if g.IsAlive(x, y) {
				str += alive
				continue
			}
			str += dead
		}

		if y < g.Height-1 {
			str += "\n"
		}
	}

	return str
}

func WillBeAlive(isAlive bool, liveNeighboursCount int) bool {
	if isAlive {
		return liveNeighboursCount == 2 || liveNeighboursCount == 3
	}

	return liveNeighboursCount == 3
}
