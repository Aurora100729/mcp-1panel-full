package process

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/Aurora100729/mcp-1panel-full/utils"
)

var ListProcessesTool = mcp.NewServerTool[ListProcessesInput, any](
	"list_processes",
	"List running processes on the 1Panel server with CPU/memory usage",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ListProcessesInput]) (*mcp.CallToolResultFor[any], error) {
		payload := map[string]interface{}{
			"type": params.Arguments.Type,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/process", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type ListProcessesInput struct {
	Type string `json:"type,omitempty" jsonschema:"process type: all, system, user. Default: all"`
}

var StopProcessTool = mcp.NewServerTool[StopProcessInput, any](
	"stop_process",
	"[DANGEROUS] Kill/stop a running process by PID",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[StopProcessInput]) (*mcp.CallToolResultFor[any], error) {
		payload := map[string]interface{}{
			"PID": params.Arguments.PID,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/process/stop", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type StopProcessInput struct {
	PID int `json:"PID" jsonschema:"process ID to stop/kill"`
}
