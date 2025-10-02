package main

import (
	"fmt"
	"speech-mcp/handlers"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create a new MCP server
	s := server.NewMCPServer(
		"MCP Speech Server",
		"1.0.0",
		server.WithToolCapabilities(false),
	)

	// Add "Say" tool
	sayTool := mcp.NewTool("say",
		mcp.WithDescription("Speak out loud something in english"),
		mcp.WithString("text",
			mcp.Required(),
			mcp.Description("The text to say out loud"),
		),
	)

	// Add tool handler
	s.AddTool(sayTool, handlers.SayHandler)

	// Add "Diga" tool
	digaTool := mcp.NewTool("diga",
		mcp.WithDescription("Di algo en voz alta en español"),
		mcp.WithString("texto",
			mcp.Required(),
			mcp.Description("El texto a decir en voz alta"),
		),
	)

	// Add tool handler
	s.AddTool(digaTool, handlers.DigaHandler)

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
