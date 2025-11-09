package game_of_life

import (
	// "fmt"
	"testing"
)

func TestGrid(t *testing.T) {
	t.Run("should create empty grid", func(t *testing.T) {
		grid := NewGrid(3, 3)

		if grid.Width != 3 || grid.Height != 3 {
			t.Errorf("expected 3x3 grid")
		}
	})

	t.Run("should set a cell alive", func(t *testing.T) {
		grid := NewGrid(3, 3)
		grid.Toggle(1, 1)

		if !grid.IsAlive(1, 1) {
			t.Errorf("expected ell at (1,1) to be alive")
		}
	})
}

func TestCountingNeighbours(t *testing.T) {
	t.Run("should count 0 live neighbours", func(t *testing.T) {
		grid := NewGrid(3, 3)
		grid.Toggle(1, 1)

		count := grid.CountLiveNeighbours(1, 1)
		if count != 0 {
			t.Errorf("expected 0 live neighbours, got %d", count)
		}
	})

	t.Run("should count 3 live neighbours", func(t *testing.T) {
		grid := NewGrid(3, 3)
		grid.Toggle(1, 1)
		// neighbours
		grid.Toggle(1, 2)
		grid.Toggle(2, 1)
		grid.Toggle(2, 2)

		count := grid.CountLiveNeighbours(1, 1)
		if count != 3 {
			t.Errorf("expected 0 live neighbours, got %d", count)
		}
	})
}

func TestCellEvolution(t *testing.T) {
	tests := []struct {
		desc            string
		isAlive         bool
		neighboursCount int
		expectedIsAlive bool
	}{
		{"alive with 2 neighbors stays alive", true, 2, true},
		{"alive with 3 neighbors stays alive", true, 3, true},
		{"alive with 1 neighbor dies", true, 1, false},
		{"alive with 4 neighbors dies", true, 4, false},
		{"dead with 3 neighbors becomes alive", false, 3, true},
		{"dead with 2 neighbors stays dead", false, 2, false},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			result := WillBeAlive(tt.isAlive, tt.neighboursCount)
			if result != tt.expectedIsAlive {
				t.Errorf("expected %v, got %v", tt.expectedIsAlive, result)
			}
		})
	}
}

func TestNextGeneration(t *testing.T) {
	t.Run("blinker pattern", func(t *testing.T) {
		// blinker pattern : horizontal line becomes vertical
		grid := NewGrid(5, 5)
		for i := 1; i <= 3; i++ {
			grid.Toggle(2, i)
		}

		next := grid.NextGeneration()

		for j := 1; j <= 3; j++ {
			if !next.IsAlive(j, 2) {
				t.Errorf("expected (%d, 2) cell to be alive, was dead", j)
			}
		}

		if next.IsAlive(2, 1) || next.IsAlive(2, 3) {
			t.Errorf("expexted old cells to be dead")
		}
	})

	t.Run("block pattern", func(t *testing.T) {
		// a block should remain stable
		grid := NewGrid(5, 5)
		positions := []Position{
			{1, 1},
			{1, 2},
			{2, 2},
			{2, 1},
		}

		for _, p := range positions {
			grid.Toggle(p.X, p.Y)
		}

		next := grid.NextGeneration()
		for _, p := range positions {
			if !next.IsAlive(p.X, p.Y) {
				t.Errorf("block is stable: expected (%d, %d) to be alive, was dead", p.X, p.Y)
			}
		}

	})

	t.Run("glider moves", func(t *testing.T) {
		// a block should remain stable
		grid := NewGrid(6, 6)
		positions := []Position{
			{1, 0},
			{2, 1},
			{0, 2},
			{1, 2},
			{2, 2},
		}

		for _, p := range positions {
			grid.Toggle(p.X, p.Y)
		}


		next := grid.NextGeneration()
		for range 3 {
			next = next.NextGeneration()
		}

		for _, p := range positions {
			if !next.IsAlive(p.X + 1, p.Y + 1) {
				t.Errorf("glider should have moved 1 step diagonally to the bottom right, (%d, %d) cell expected to be alive, was dead", p.X + 1, p.Y + 1)
			}
		}
	})
}
