package v2

import (
	"testing"
)

func TestGameOfLife(t *testing.T) {
	t.Run("should create a 5x5 grid", func(t *testing.T) {
		grid := NewGrid(5, 5)

		if grid.Width != 5 || grid.Height != 5 {
			t.Errorf("Expected 5x5 grid, got %dx%d", grid.Width, grid.Height)
		}
	})

	t.Run("should toggle a cell alive", func(t *testing.T) {
		grid := NewGrid(5, 5)

		grid.Toggle(2, 2)
		if !grid.IsAlive(2, 2) {
			t.Errorf("Expected cell (2,2) to be alive, was dead")
		}
	})

	t.Run("should return the count of live neighbours", func(t *testing.T) {
		grid := NewGrid(5, 5)

		grid.Toggle(2, 2)
		grid.Toggle(2, 3)
		grid.Toggle(1, 3)

		expected := 2

		actual := grid.CountLiveNeighbours(2, 2)

		if actual != expected {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	})

	t.Run("should tell wether cell will be alive in the next generation", func(t *testing.T) {

		cases := []struct {
			Desc       string
			Expected   bool
			IsAlive    bool
			Neighbours int
		}{
			{Desc: "Is alive with 2 neighbours should become alive", Expected: true, IsAlive: true, Neighbours: 2},
			{Desc: "Is alive with 3 neighbours should become alive", Expected: true, IsAlive: true, Neighbours: 3},
			{Desc: "Is alive with 4 neighbours should become dead", Expected: false, IsAlive: true, Neighbours: 4},
			{Desc: "Is alive with 1 neighbours should become dead", Expected: false, IsAlive: true, Neighbours: 1},
			{Desc: "Is not alive with 3 neighbours should become alive", Expected: true, IsAlive: false, Neighbours: 3},
			{Desc: "Is not alive with 5 neighbours should become dead", Expected: false, IsAlive: false, Neighbours: 5},
			{Desc: "Is not alive with 2 neighbours should become dead", Expected: false, IsAlive: false, Neighbours: 2},
		}

		for _, c := range cases {
			t.Run(c.Desc, func(t *testing.T) {
				actual := WillBeAlive(c.IsAlive, c.Neighbours)
				if actual != c.Expected {
					t.Errorf("Expected result to be %v, got %v", c.Expected, actual)
				}

			})
		}
	})
	t.Run("should give the next generation", func(t *testing.T) {

		cases := []struct {
			Desc              string
			Generations       int
			AliveCellsAtStart []Position
			AliveCellsAtEnd   []Position
		}{
			{"Blinker Pattern: 1x3 should become 3x1", 1, []Position{{2, 3}, {3, 3}, {4, 3}}, []Position{{3, 2}, {3, 3}, {3, 4}}},
			{"Block Pattern: 2x2 should remain 2x2", 1, []Position{{3, 3}, {3, 4}, {4, 3}, {4, 4}}, []Position{{3, 3}, {3, 4}, {4, 3}, {4, 4}}},
			{"Glider Pattern: should translate on 4 gens", 4, []Position{{1, 2}, {2, 3}, {3, 1}, {3, 2}, {3, 3}}, []Position{{2, 3}, {3, 4}, {4, 2}, {4, 3}, {4, 4}}},
		}

		for _, c := range cases {
			t.Run(c.Desc, func(t *testing.T) {
				grid := NewGrid(10, 10)
				for _, a := range c.AliveCellsAtStart {
					grid.Toggle(a.X, a.Y)
				}

				for range c.Generations {
					grid = grid.Next()
				}

				for _, a := range c.AliveCellsAtEnd {
					if !grid.IsAlive(a.X, a.Y) {
						t.Errorf("Expected cell (%d,%d) to be alive, was dead", a.X, a.Y)
					}
				}
			})
		}

	})
}
