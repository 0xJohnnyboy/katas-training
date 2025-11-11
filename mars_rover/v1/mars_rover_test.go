package v1

import (
	"testing"
)

func TestRover(t *testing.T) {
	t.Run("creates a rover with initial position", func(t *testing.T) {
		grid := Grid{Size: Coordinates{X: 5, Y: 5}}
		rover := NewRover(2, 3, North, grid)

		expectedCoordinates := Coordinates{X: 2, Y: 3}

		if !rover.Position.Equals(expectedCoordinates) {
			t.Errorf("expected position (2, 3), got (%d, %d)", rover.Position.X, rover.Position.Y)
		}

		if rover.Direction != North {
			t.Errorf("expected direction North, got %v", rover.Direction)
		}
	})

	t.Run("rover should return its position", func(t *testing.T) {
		grid := Grid{Size: Coordinates{X: 5, Y: 5}}
		rover := NewRover(2, 4, West, grid)

		if rover.String() != "2:4:W" {
			t.Errorf("expected position 2:4:W, got %s", rover.String())
		}
	})

	t.Run("rover facing north should face west", func(t *testing.T) {
		grid := Grid{Size: Coordinates{X: 5, Y: 5}}
		rover := NewRover(0, 0, North, grid)

		rover.Execute("L")

		if rover.Direction != West {
			t.Errorf("expected direction West, got %v", rover.Direction)
		}
	})

	t.Run("rover facing north should face north", func(t *testing.T) {
		grid := Grid{Size: Coordinates{X: 5, Y: 5}}
		rover := NewRover(0, 0, North, grid)

		rover.Execute("LRLR")

		if rover.Direction != North {
			t.Errorf("expected direction North, got %v", rover.Direction)
		}
	})

	t.Run("rover should move forward", func(t *testing.T) {
		grid := Grid{Size: Coordinates{X: 5, Y: 5}}
		rover := NewRover(2, 2, North, grid)

		position := rover.Execute("F")

		if position != "2:1:N" {
			t.Errorf("expected position 2:1:N, got %s", rover.Direction)
		}
	})

	t.Run("rover should move", func(t *testing.T) {
		grid := Grid{Size: Coordinates{X: 5, Y: 5}}
		rover := NewRover(2, 2, North, grid)

		position := rover.Execute("FLFRFRFFBRFFLL")

		if position != "2:2:N" {
			t.Errorf("expected position 2:2:N, got %s", position)
		}
	})

	t.Run("rover should move with wrapping", func(t *testing.T) {
		grid := Grid{Size: Coordinates{X: 5, Y: 5}}
		rover := NewRover(2, 2, North, grid)

		position := rover.Execute("LFRFFFRFLFF")

		if position != "2:2:N" {
			t.Errorf("expected position 2:2:N, got %s", position)
		}
	})

	t.Run("grid should have obstacles", func(t *testing.T) {
		grid := Grid{Size: Coordinates{X: 5, Y: 5}}
		rover := NewRover(2, 2, North, grid)

		obstacles := []Obstacle{
			{Position: Coordinates{X: 2, Y: 1}},
			{Position: Coordinates{X: 3, Y: 3}},
		}

		rover.Grid.Obstacles = obstacles

		if len(rover.Grid.Obstacles) != 2 {
			t.Errorf("expected grid to have 2 obstacles, got %d", len(rover.Grid.Obstacles))
		}

		for i, o := range obstacles {
			if !rover.Grid.Obstacles[i].Position.Equals(o.Position) {
				t.Errorf("expected obstacle %d to be %v, got %v", i, o.Position, rover.Grid.Obstacles[i].Position)
			}
		}
	})

	t.Run("rover should stop at obstacle", func(t *testing.T) {
		grid := Grid{Size: Coordinates{X: 5, Y: 5}}
		rover := NewRover(2, 2, North, grid)

		obstacles := []Obstacle{
			{Position: Coordinates{X: 2, Y: 1}},
			{Position: Coordinates{X: 3, Y: 3}},
		}

		rover.Grid.Obstacles = obstacles

		position := rover.Execute("LFRFRFFFF")

		if position != "1:1:E" {
			t.Errorf("expected position 1:1:E, got %s", position)
		}

	})

	t.Run("should print grid", func(t *testing.T) {
		grid := Grid{Size: Coordinates{X: 5, Y: 5}}
		rover := NewRover(2, 2, North, grid)

		obstacles := []Obstacle{
			{Position: Coordinates{X: 2, Y: 1}},
			{Position: Coordinates{X: 3, Y: 3}},
		}

		rover.Grid.Obstacles = obstacles

		rover.Execute("LFFBRRR")

		expected := "-----\n--X--\n-v---\n---X-\n-----"
		actual := rover.ToString()

		if actual != expected {
			t.Errorf("expected grid to be \n%s\n, got \n%s", expected, actual)
		}
	})
}
