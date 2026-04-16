package game

import "mcp_server/mcp-server/internal/model"

type Game interface {
	NewGame()
	MoveUp()
	MoveDown()
	MoveLeft()
	MoveRight()
	GetState() model.State
}
