package tennis

import (
	"testing"
	// "errors"
)

func TestTennis(t *testing.T) {

	t.Run("Initial score is love-all", func(t *testing.T) {
		game := NewGame()
		expected := "Love-All"
		actual := game.GetScore()

		if actual != expected {
			t.Errorf("GetScore() expected %s, got %s", expected, actual)
		}
	})

	t.Run("Player scores once", func(t *testing.T) {
		game := NewGame()
		game.PointWonBy(Player1)
		expected := "15-Love"
		actual := game.GetScore()

		if actual != expected {
			t.Errorf("GetScore() expected %s, got %s", expected, actual)
		}
	})

	t.Run("Player 1 scores once, player 2 scores once, player 1 scores twice", func(t *testing.T) {
		game := NewGame()
		game.PointWonBy(Player1)
		game.PointWonBy(Player2)
		game.PointWonBy(Player1)
		game.PointWonBy(Player1)
		expected := "40-15"
		actual := game.GetScore()

		if actual != expected {
			t.Errorf("GetScore() expected %s, got %s", expected, actual)
		}
	})

	t.Run("Player 1 scores followed by player 2, 3 times", func(t *testing.T) {
		game := NewGame()
		for range 3 {
			game.PointWonBy(Player1)
			game.PointWonBy(Player2)
		}
		expected := "Deuce"
		actual := game.GetScore()

		if actual != expected {
			t.Errorf("GetScore() expected %s, got %s", expected, actual)
		}
	})

	t.Run("At Deuce, player 1 scores once", func(t *testing.T) {
		game := NewGame()
		for range 3 {
			game.PointWonBy(Player1)
			game.PointWonBy(Player2)
		}

		game.PointWonBy(Player1)
		expected := "Advantage player 1"
		actual := game.GetScore()

		if actual != expected {
			t.Errorf("GetScore() expected %s, got %s", expected, actual)
		}
	})

	t.Run("At Advantage player 1, player 2 scores once", func(t *testing.T) {
		game := NewGame()
		for range 3 {
			game.PointWonBy(Player1)
			game.PointWonBy(Player2)
		}

		game.PointWonBy(Player1)
		game.PointWonBy(Player2)
		expected := "Deuce"
		actual := game.GetScore()

		if actual != expected {
			t.Errorf("GetScore() expected %s, got %s", expected, actual)
		}
	})

	t.Run("At Advantage player 2, player 2 scores once", func(t *testing.T) {
		game := NewGame()
		for range 3 {
			game.PointWonBy(Player1)
			game.PointWonBy(Player2)
		}

		game.PointWonBy(Player2)
		game.PointWonBy(Player2)
		expected := "player 2 wins"

		actual := game.GetScore()

		if actual != expected {
			t.Errorf("GetScore() expected %s, got %s", expected, actual)
		}
	})
}
