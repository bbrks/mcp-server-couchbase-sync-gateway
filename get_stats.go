package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func getStatsHandler(cfg *serverConfig) func(ctx context.Context, sess *mcp.ServerSession, params *mcp.CallToolParamsFor[struct{}]) (*mcp.CallToolResultFor[any], error) {
	return func(ctx context.Context, sess *mcp.ServerSession, params *mcp.CallToolParamsFor[struct{}]) (*mcp.CallToolResultFor[any], error) {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, cfg.adminURL+"/_expvar", nil)
		if err != nil {
			return nil, err
		}
		if cfg.adminUsername != "" || cfg.adminPassword != "" {
			req.SetBasicAuth(cfg.adminUsername, cfg.adminPassword)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, err := io.ReadAll(io.LimitReader(resp.Body, 4096))
			if err != nil {
				return nil, fmt.Errorf("failed to read error response body: %w", err)
			}
			return nil, fmt.Errorf("failed to get stats - %d: %s", resp.StatusCode, string(body))
		}

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		// parse expvar JSON just to check we got a sane looking result back before returning it
		var parsed map[string]any
		if err := json.Unmarshal(data, &parsed); err != nil {
			return nil, fmt.Errorf("failed to parse expvar JSON: %w", err)
		}

		result := &mcp.CallToolResultFor[any]{
			Content: []mcp.Content{&mcp.TextContent{Text: string(data)}},
		}
		return result, nil
	}
}
