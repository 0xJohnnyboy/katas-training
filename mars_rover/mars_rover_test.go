package mars_rover

import (
	"testing"
)

func TestRover(t *testing.T) {
	t.Run("creates a rover with initial position", func(t *testing.T) {
		rover := NewRover(2, 3, North, 5, 5)

		if rover.X != 2 || rover.Y != 3 {
			t.Errorf("expected position (2, 3), got (%d, %d)", rover.X, rover.Y)
		}

		if rover.Direction != North {
			t.Errorf("expected direction North, got %v", rover.Direction)
		}
	})

	t.Run("rover should return its position", func(t *testing.T) {
		rover := NewRover(2, 4, West, 5, 5)
		if rover.String() != "2:4:W" {
			t.Errorf("expected position 2:4:W, got %s", rover.String())
		}
	})

	t.Run("rover facing north should face west", func(t *testing.T) {
		rover := NewRover(0, 0, North, 5, 5)

		rover.Execute("L")

		if rover.Direction != West {
			t.Errorf("expected direction West, got %v", rover.Direction)
		}
	})

	t.Run("rover facing north should face north", func(t *testing.T) {
		rover := NewRover(0, 0, North, 5, 5)

		rover.Execute("LRLR")

		if rover.Direction != North {
			t.Errorf("expected direction North, got %v", rover.Direction)
		}
	})

	t.Run("rover should move forward", func(t *testing.T) {
		rover := NewRover(2, 2, North, 5, 5)

		position := rover.Execute("F")

		if position != "2:1:N" {
			t.Errorf("expected position 2:1:N, got %s", rover.Direction)
		}
	})

	t.Run("rover should move", func(t *testing.T) {
		rover := NewRover(2, 2, North, 5, 5)

		position := rover.Execute("FLFRFRFFBRFFLL")

		if position != "2:2:N" {
			t.Errorf("expected position 2:2:N, got %s", position)
		}
	})

	t.Run("rover should move with wrapping", func(t *testing.T) {
		rover := NewRover(2, 2, North, 5, 5)

		position := rover.Execute("LFRFFFRFLFF")

		if position != "2:2:N" {
			t.Errorf("expected position 2:2:N, got %s", position)
		}
	})
}
