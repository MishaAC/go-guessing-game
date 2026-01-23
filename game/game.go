package game

import "math/rand"

type Result int

const (
	TooLow Result = iota
	TooHigh
	Correct
	GameOver
)

type Game struct {
	secretNumber  int
	min           int
	max           int
	attemptsLeft  int
	totalAttempts int
}

func New(min, max, attempts int) *Game {
	return &Game{
		secretNumber:  rand.Intn(max-min+1) + min,
		min:           min,
		max:           max,
		attemptsLeft:  attempts,
		totalAttempts: attempts,
	}
}

func (g *Game) Guess(number int) Result {
	if g.attemptsLeft <= 0 {
		return GameOver
	}
	g.attemptsLeft--
	switch {
	case number < g.secretNumber:
		return TooLow
	case number > g.secretNumber:
		return TooHigh
	default:
		return Correct
	}
}

func (g *Game) AttemptsLeft() int {
	return g.attemptsLeft
}
