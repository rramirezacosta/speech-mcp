package handlers

import (
	"context"
	"fmt"
	"mcp_speech/speech"

	"github.com/mark3labs/mcp-go/mcp"
)

func SayHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	text, err := request.RequireString("text")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	speech.Say(text, "en")

	return mcp.NewToolResultText(fmt.Sprintf("Saying out loud: %s", text)), nil
}
