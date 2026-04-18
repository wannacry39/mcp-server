package game

import (
	"math/rand"
	"mcp_server/mcp-server/internal/model"
)

type service struct {
	state model.State
	moves int
}

func NewGame() Game {
	g := &service{}
	g.NewGame()
	return g
}

func (g *service) NewGame() {
	g.state = model.NewSolvedState()
	g.moves = 0

	dirs := []string{"up", "down", "left", "right"}

	for i := 0; i < 100; i++ {
		_ = g.state.MoveByDirection(dirs[rand.Intn(len(dirs))])
	}
}

func (g *service) Move(tile string) error {
	err := g.state.Move(tile)
	if err != nil {
		return err
	}

	g.moves++
	return nil
}

func (g *service) GetState() model.State {
	return g.state.Clone()
}

func (g *service) IsSolved() bool {
	expected := 1

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {

			if i == 3 && j == 3 {
				return g.state.Board[i][j] == 0
			}

			if g.state.Board[i][j] != expected {
				return false
			}

			expected++
		}
	}

	return true
}
