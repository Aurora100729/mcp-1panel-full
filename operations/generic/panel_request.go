package generic

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/Aurora100729/mcp-1panel-full/utils"
)

var PanelRequestTool = mcp.NewServerTool[PanelRequestInput, any](
	"panel_request",
	"[PASSTHROUGH] Send any request to 1Panel API. Use this for any endpoint not covered by dedicated tools. Method: GET/POST/PUT/DELETE. Path example: /containers, /websites/ssl/search. Body is optional JSON object.",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[PanelRequestInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		method := strings.ToUpper(input.Method)
		if method == "" {
			method = "GET"
		}
		path := input.Path
		if path == "" {
			return &mcp.CallToolResult{
				Content: []mcp.Content{&mcp.TextContent{Text: "path is required"}},
				IsError: true,
			}, nil
		}

		opts := []utils.Option{}
		if input.Body != nil {
			opts = append(opts, utils.WithPayload(input.Body))
		}
		if len(input.Query) > 0 {
			q := make(map[string]interface{})
			for k, v := range input.Query {
				q[k] = v
			}
			opts = append(opts, utils.WithQuery(q))
		}
		if len(input.Headers) > 0 {
			opts = append(opts, utils.WithHeaders(input.Headers))
		}

		client := utils.NewPanelClient(method, path, opts...)
		_, err := client.Do()
		if err != nil {
			return &mcp.CallToolResult{
				Content: []mcp.Content{&mcp.TextContent{Text: err.Error()}},
				IsError: true,
			}, err
		}

		body, err := client.GetRespBody()
		if err != nil {
			return &mcp.CallToolResult{
				Content: []mcp.Content{&mcp.TextContent{Text: "Failed to read response: " + err.Error()}},
				IsError: true,
			}, err
		}

		var result interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			return &mcp.CallToolResult{
				Content: []mcp.Content{&mcp.TextContent{Text: string(body)}},
			}, nil
		}

		pretty, _ := json.MarshalIndent(result, "", "  ")
		return &mcp.CallToolResult{
			Content:           []mcp.Content{&mcp.TextContent{Text: string(pretty)}},
			StructuredContent: result,
		}, nil
	},
)

type PanelRequestInput struct {
	Method  string            `json:"method" jsonschema:"HTTP method: GET, POST, PUT, DELETE"`
	Path    string            `json:"path" jsonschema:"API path relative to /api/v2, e.g. /containers or /websites/ssl/search"`
	Body    interface{}       `json:"body,omitempty" jsonschema:"optional JSON request body"`
	Query   map[string]string `json:"query,omitempty" jsonschema:"optional URL query parameters"`
	Headers map[string]string `json:"headers,omitempty" jsonschema:"optional extra HTTP headers"`
}
