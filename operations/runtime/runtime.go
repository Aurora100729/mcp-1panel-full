package runtime

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/Aurora100729/mcp-1panel-full/utils"
)

var ListRuntimesTool = mcp.NewServerTool[ListRuntimesInput, any](
	"list_runtimes",
	"List installed runtimes (PHP, Node.js, Java, Go, Python, .NET)",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ListRuntimesInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"page":     1,
			"pageSize": 500,
			"name":     input.Name,
			"type":     input.Type,
			"orderBy":  "name",
			"order":    "ascending",
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/runtimes/search", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type ListRuntimesInput struct {
	Name string `json:"name,omitempty" jsonschema:"filter by runtime name"`
	Type string `json:"type,omitempty" jsonschema:"filter by type: php, node, java, go, python, dotnet"`
}

var CreateRuntimeTool = mcp.NewServerTool[CreateRuntimeInput, any](
	"create_runtime",
	"Create a new runtime environment (PHP, Node.js, etc.)",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[CreateRuntimeInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"name":    input.Name,
			"type":    input.Type,
			"appDetailId": input.AppDetailID,
			"image":   input.Image,
			"version": input.Version,
			"params":  input.Params,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/runtimes", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type CreateRuntimeInput struct {
	Name        string      `json:"name" jsonschema:"runtime name"`
	Type        string      `json:"type" jsonschema:"runtime type: php, node, java, go, python, dotnet"`
	AppDetailID int         `json:"appDetailId,omitempty" jsonschema:"app detail ID"`
	Image       string      `json:"image,omitempty" jsonschema:"custom Docker image"`
	Version     string      `json:"version,omitempty" jsonschema:"runtime version"`
	Params      interface{} `json:"params,omitempty" jsonschema:"additional parameters (JSON)"`
}

var RuntimeOperateTool = mcp.NewServerTool[RuntimeOperateInput, any](
	"runtime_operate",
	"[DANGEROUS] Operate on a runtime: start, stop, restart, delete",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[RuntimeOperateInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		path := fmt.Sprintf("/runtimes/operate")
		payload := map[string]interface{}{
			"ID":        input.ID,
			"operate":   input.Operate,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", path, utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type RuntimeOperateInput struct {
	ID      int    `json:"ID" jsonschema:"runtime ID"`
	Operate string `json:"operate" jsonschema:"operation: start, stop, restart, delete"`
}
