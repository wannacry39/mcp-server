package handler

import (
	"context"
	"fmt"
	"mcp_server/mcp-server/internal/game"

	"github.com/mark3labs/mcp-go/mcp"
)

type service struct {
	gameService game.Game
}

func NewHandler(game game.Game) Handler {
	return &service{
		gameService: game,
	}
}

func (s *service) HandleNewGame(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	s.gameService.NewGame()
	state := s.gameService.GetState()
	return mcp.NewToolResultText(fmt.Sprintf("New game started!\n%s", state.String())), nil
}

func (s *service) HandleMove(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args, ok := req.Params.Arguments.(map[string]any)
	if !ok {
		return mcp.NewToolResultError("invalid arguments"), nil
	}

	tile, ok := args["tile"].(string)
	if !ok {
		return mcp.NewToolResultError("tile must be a number"), nil
	}

	if err := s.gameService.Move(tile); err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(s.gameService.GetState().String()), nil
}

func (s *service) HandleGetState(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	state := s.gameService.GetState()
	return mcp.NewToolResultText(state.String()), nil
}

func (s *service) HandleIsSolved(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	IsSolved := s.gameService.IsSolved()
	if IsSolved {
		return mcp.NewToolResultText(fmt.Sprintf("congrats!!! your game is solved\n%s", s.gameService.GetState().String())), nil
	}
	return mcp.NewToolResultText(fmt.Sprintf("your game is not solved yet\n%s", s.gameService.GetState().String())), nil
}
