package v2

import (
	"testing"
)

func TestMarsRover(t *testing.T) {
	t.Run("should create a rover with initial position", func(t *testing.T) {
		grid := NewGrid(5, 5, []Obstacle{})
		rover := NewRover(2, 3, North, grid)

		if rover.Grid.Height != 5 || rover.Grid.Width != 5 {
			t.Errorf("Grid should be 5x5, got %dx%d", rover.Grid.Height, rover.Grid.Width)
		}

		if rover.X != 2 || rover.Y != 3 || rover.Direction != North {
			t.Errorf("Rover should be init with (2, 3, North), got (%d, %d, %v)", rover.X, rover.Y, rover.Direction)
		}
	})

	t.Run("should return a string with the rover's position", func(t *testing.T) {
		grid := NewGrid(5, 5, []Obstacle{})
		rover := NewRover(3, 4, South, grid)

		expected := "3:4:S"
		actual := rover.ToString()
		if actual != expected {
			t.Errorf("Expected %s, got %s", expected, actual)
		}

	})

	t.Run("rover should move forward 1 step", func(t *testing.T) {
		grid := NewGrid(5, 5, []Obstacle{})
		rover := NewRover(2, 3, North, grid)

		expected := "2:2:N"

		actual := rover.Execute("F")

		if actual != expected {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
	})

	t.Run("rover should have moved forward 2 steps", func(t *testing.T) {
		grid := NewGrid(5, 5, []Obstacle{})
		rover := NewRover(2, 3, North, grid)

		expected := "2:1:N"

		actual := rover.Execute("BFFFBF")

		if actual != expected {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
	})

	t.Run("rover should turn left", func(t *testing.T) {
		grid := NewGrid(5, 5, []Obstacle{})
		rover := NewRover(2, 3, North, grid)

		expected := "2:3:W"

		actual := rover.Execute("L")

		if actual != expected {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
	})

	t.Run("rover should have turned right", func(t *testing.T) {
		grid := NewGrid(5, 5, []Obstacle{})
		rover := NewRover(2, 3, North, grid)

		expected := "2:3:E"

		actual := rover.Execute("LLRLRRLRR")

		if actual != expected {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
	})

	t.Run("rover should be able to move forward, backward and turn both directions", func(t *testing.T) {
		grid := NewGrid(5, 5, []Obstacle{})
		rover := NewRover(2, 3, North, grid)
		expected := "1:4:E"

		actual := rover.Execute("FLBLFFRFFRR")

		if actual != expected {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
	})

	t.Run("grid should be wrapping", func(t *testing.T) {
		grid := NewGrid(3, 3, []Obstacle{})
		rover := NewRover(0, 2, West, grid)
		expected := "2:0:S"

		actual := rover.Execute("FLF")

		if actual != expected {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
	})

	t.Run("should return true if position has an obstacle", func(t *testing.T) {
		obstacles := []Obstacle{
			{1, 0},
			{3, 3},
		}
		grid := NewGrid(4, 4, obstacles)
		rover := NewRover(1, 2, North, grid)

		for _, o := range obstacles {
			if !rover.IsObstacle(o.X, o.Y) {
				t.Errorf("Expected position (%d,%d) to be an obstacle", o.X, o.Y)
			}
		}
	})

	t.Run("obstacles should stop the rover", func(t *testing.T) {
		grid := NewGrid(4, 4, []Obstacle{})
		rover := NewRover(1, 2, North, grid)
		rover.Grid.Obstacles = []Obstacle{
			{1, 0},
		}

		expected := "1:1:N"

		actual := rover.Execute("FF")

		if actual != expected {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
	})
}
