package process

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/Aurora100729/mcp-1panel-full/utils"
)

var ListProcessesTool = mcp.NewServerTool[ListProcessesInput, any](
	"list_processes",
	"Get process info by PID on the 1Panel server. Provide a PID to inspect (default: 1)",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ListProcessesInput]) (*mcp.CallToolResultFor[any], error) {
		pid := params.Arguments.PID
		if pid == "" {
			pid = "1"
		}
		var result interface{}
		client := utils.NewPanelClient("GET", "/process/"+pid)
		return client.Request(&result)
	},
)

type ListProcessesInput struct {
	PID  string `json:"pid,omitempty" jsonschema:"process ID to inspect (default: 1)"`
	Type string `json:"type,omitempty" jsonschema:"deprecated, ignored"`
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
