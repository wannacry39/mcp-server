package main

import (
	"mcp_server/mcp-server/internal/game"
	"mcp_server/mcp-server/internal/handler"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	game := game.NewGame()
	handler := handler.NewHandler(game)

	s := server.NewMCPServer(
		"15-puzzle",
		"v0.0.1",
	)

	s.AddTool(mcp.NewTool(
		"new_game",
		mcp.WithDescription("Create a new 15 puzzle game with randomly placed puzzles"),
	), handler.HandleNewGame)

	s.AddTool(mcp.NewTool(
		"move",
		mcp.WithDescription("Move a tile to the empty space"),
		mcp.WithString("tile",
			mcp.Required(),
			mcp.Description("Number of the tile to move, e.g. 15"),
		),
	), handler.HandleMove)

	s.AddTool(mcp.NewTool(
		"get_state",
		mcp.WithDescription("Returns current state of the game"),
	), handler.HandleGetState)

	s.AddTool(mcp.NewTool(
		"is_solved",
		mcp.WithDescription("Returns if the game solved or not"),
	), handler.HandleIsSolved)

	// if err := server.ServeStdio(s); err != nil {
	// 	panic(err)
	// }

	httpServer := server.NewSSEServer(s)
	if err := httpServer.Start(":9999"); err != nil {
		panic(err)
	}
}
