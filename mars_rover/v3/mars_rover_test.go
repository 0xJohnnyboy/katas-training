package v3

import (
	"testing"
)

func TestMarsRover(t *testing.T) {
	t.Run("Should create a 5x5 grid", func(t *testing.T) {
		grid := NewGrid(5, 5)
		if grid.Width != 5 || grid.Height != 5 {
			t.Errorf("Expected 5x5 grid, got %dx%d", grid.Width, grid.Height)
		}
	})

	t.Run("Should return wether 2 positions are equal", func(t *testing.T) {
		position1 := Position{4, 4}
		position2 := Position{4, 4}

		expected := true

		actual := position1.Equals(position2)
		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("Should create a rover on a 5x5 grid with position (2,2) and direction N", func(t *testing.T) {
		grid := NewGrid(5, 5)
		rover := NewRover(2, 2, North, grid)

		expectedPos := Position{2, 2}
		expectedDir := North

		if expectedDir != rover.Direction {
			t.Errorf("Expected rover to have direction %v, got %v", expectedDir, rover.Direction)
		}
		if expectedPos != rover.Position {
			t.Errorf("Expected rover to have position %v, got %v", expectedPos, rover.Position)
		}
	})
	t.Run("Rover should return position as a string", func(t *testing.T) {
		grid := NewGrid(5, 5)
		rover := NewRover(2, 2, North, grid)

		expected := "2:2:N"
		actual := rover.ToString()
		if actual != expected {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
	})

	t.Run("Rover should turn left", func(t *testing.T) {
		grid := NewGrid(5, 5)
		rover := NewRover(2, 2, North, grid)

		expected := "2:2:W"
		actual := rover.Execute("L")

		if expected != actual {
			t.Errorf("Expected position %s, got %s", expected, actual)
		}
	})

	t.Run("Rover should handle multiple turns", func(t *testing.T) {
		grid := NewGrid(5, 5)
		rover := NewRover(2, 2, North, grid)

		expected := "2:2:S"
		actual := rover.Execute("LRRR")

		if expected != actual {
			t.Errorf("Expected position %s, got %s", expected, actual)
		}
	})

	t.Run("Rover should move forward", func(t *testing.T) {
		grid := NewGrid(5, 5)
		rover := NewRover(2, 2, North, grid)

		expected := "2:1:N"
		actual := rover.Execute("F")

		if expected != actual {
			t.Errorf("Expected position %s, got %s", expected, actual)
		}
	})

	t.Run("Rover should move backward with wrapping", func(t *testing.T) {
		grid := NewGrid(3, 3)
		rover := NewRover(0, 2, East, grid)

		expected := "2:2:E"
		actual := rover.Execute("B")

		if expected != actual {
			t.Errorf("Expected position %s, got %s", expected, actual)
		}
	})

	t.Run("Rover should handle complex command", func(t *testing.T) {
		grid := NewGrid(5, 5)
		rover := NewRover(0, 0, South, grid)

		expected := "0:3:E"
		actual := rover.Execute("FFBFLFFRFLFFF")

		if expected != actual {
			t.Errorf("Expected position %s, got %s", expected, actual)
		}
	})

	t.Run("Rover should stop at obstacle", func(t *testing.T) {
		grid := NewGrid(5, 5)
		grid.Obstacles = []Position{
			{1, 2},
			{2, 1},
		}

		rover := NewRover(2, 2, North, grid)

		expected := "2:2:N"
		actual := rover.Execute("FFF")

		if expected != actual {
			t.Errorf("Expected position %s, got %s", expected, actual)
		}
	})
}
