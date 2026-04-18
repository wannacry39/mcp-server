package handler

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
)

type Handler interface {
	HandleNewGame(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
	HandleMove(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
	HandleGetState(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
	HandleIsSolved(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}
