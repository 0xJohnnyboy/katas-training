package tennis

import "fmt"

type Player int

const (
	Player1 Player = iota
	Player2
)

type Game struct {
	score1 int
	score2 int
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) PointWonBy(player Player) {
	if player == Player1 {
		if g.score1 == 3 && g.score2 == 4 {
			g.score2--
			return
		}

		g.score1++
	} else {
		if g.score1 == 4 && g.score2 == 3 {
			g.score1--
			return
		}

		g.score2++
	}
}

var Scores = map[int]string{
	0: "Love",
	1: "15",
	2: "30",
	3: "40",
}

func (g *Game) GetScore() string {
	sum := g.score1 + g.score2
	if sum > 0 && sum < 6 {
		return fmt.Sprintf("%s-%s", Scores[g.score1], Scores[g.score2])
	}

	switch {
	case g.score1 == 3 && g.score2 == 3:
		return "Deuce"
	case g.score1 == 4 && g.score2 == 3:
		return "Advantage player 1"
	case g.score1 == 3 && g.score2 == 4:
		return "Advantage player 2"
	case g.score1 == 5 && g.score2 == 3:
		return "player 1 wins"
	case g.score1 == 3 && g.score2 == 5:
		return "player 2 wins"
	}

	return "Love-All"
}
