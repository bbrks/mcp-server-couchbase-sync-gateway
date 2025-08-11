package main

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	cfg, err := loadConfig()
	if err != nil {
		fatal(err)
	}

	server := mcp.NewServer(&mcp.Implementation{
		Name:    "mcp-server-couchbase-sync-gateway",
		Version: "0.1.0",
	}, nil)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_stats",
		Annotations: &mcp.ToolAnnotations{Title: "Get Sync Gateway Metrics", ReadOnlyHint: true},
		Description: "Retrieve point-in-time snapshot of Couchbase Sync Gateway statistics (or metrics) as JSON. This does not include historical data, but will give an updated view of the live system each time this is called. Prefer to re-run this tool whenever there is a request about Sync Gateway and its stats or metrics.",
	}, getStatsHandler(cfg))

	if err := server.Run(context.Background(), mcp.NewStdioTransport()); err != nil {
		fatal(err)
	}
}
