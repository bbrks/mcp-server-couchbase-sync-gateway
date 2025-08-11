package main

import (
	"errors"
	"flag"
	"os"
	"strings"
)

type serverConfig struct {
	adminURL      string
	adminUsername string
	adminPassword string
}

func loadConfig() (*serverConfig, error) {
	var (
		urlFlag = flag.String("sg-admin-api-url", "", "Sync Gateway Admin API base URL, e.g. https://sgw.example.com:4985")
	)
	flag.Parse()

	if *urlFlag == "" {
		return nil, errors.New("missing required Sync Gateway Admin API URL (set --sg-admin-api-url)")
	}

	username := os.Getenv("COUCHBASE_SYNC_GATEWAY_ADMIN_USERNAME")
	password := os.Getenv("COUCHBASE_SYNC_GATEWAY_ADMIN_PASSWORD")

	return &serverConfig{
		adminURL:      strings.TrimRight(*urlFlag, "/"),
		adminUsername: username,
		adminPassword: password,
	}, nil
}
