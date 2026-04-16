package game

import "mcp_server/mcp-server/internal/model"

type service struct {
	state model.State
}

func New() Game {
	s := &service{}
	s.NewGame()
	return s
}

func (s *service) NewGame() {
	startState := model.NewState()
	s.state = startState

}

func (s *service) GetState() model.State {
	return s.state

}

func (s *service) MoveUp() {
	s.state.MoveUp()

}

func (s *service) MoveDown() {
	s.state.MoveDown()

}

func (s *service) MoveLeft() {
	s.state.MoveLeft()

}

func (s *service) MoveRight() {
	s.state.MoveRight()

}
