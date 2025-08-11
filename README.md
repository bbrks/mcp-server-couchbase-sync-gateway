# mcp-server-couchbase-sync-gateway

Experimental and unofficial MCP Server for interacting with Couchbase Sync Gateway.

## Disclaimer

⚠️ This is a personal experiment in building an MCP Server. It is not production ready nor officially supported by Couchbase. For official Couchbase MCP support, see  [mcp-server-couchbase](https://github.com/Couchbase-Ecosystem/mcp-server-couchbase)

## Demo

The below was run *completely offline* in LM Studio using OpenAI's `gpt-oss-20b` model running a on a **16GB RTX 4080 Super**

https://github.com/user-attachments/assets/943cd42e-c613-4e51-b4ae-611f97156e4f

## Installation

Download from releases, or install binary via Go

```bash
go get github.com/bbrks/mcp-server-couchbase-sync-gateway@latest
```

## Configuration

Configure via flags and env vars (usually set in your `mcp.json` configuration):

- `--sg-admin-api` or `COUCHBASE_SYNC_GATEWAY_ADMIN_API` (required): Admin API base URL, e.g. `https://localhost:4985`
- `COUCHBASE_SYNC_GATEWAY_ADMIN_USERNAME` (optional)
- `COUCHBASE_SYNC_GATEWAY_ADMIN_PASSWORD` (optional)

### LM Studio `mcp.json` config example:

```json
{
  "mcpServers": {
    "couchbase-sync-gateway": {
      "command": "couchbase-sync-gateway.exe",
      "args": [
        "--sg-admin-api-url",
        "http://192.168.68.254:4985"
      ],
      "env": {
        "COUCHBASE_SYNC_GATEWAY_ADMIN_USERNAME": "sgw-mcp-admin-user",
        "COUCHBASE_SYNC_GATEWAY_ADMIN_PASSWORD": "hunter2"
      }
    }
  }
}
```
