package game

import "mcp_server/mcp-server/internal/model"

type Game interface {
	NewGame()
	Move(string) error
	GetState() model.State
	IsSolved() bool
}
