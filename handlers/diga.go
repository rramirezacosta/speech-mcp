package handlers

import (
	"context"
	"fmt"
	"speech-mcp/speech"

	"github.com/mark3labs/mcp-go/mcp"
)

func DigaHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	text, err := request.RequireString("texto")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	speech.Say(text, "es")

	return mcp.NewToolResultText(fmt.Sprintf("Saying out loud: %s", text)), nil
}
