package panellog

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/Aurora100729/mcp-1panel-full/utils"
)

var OperationLogsTool = mcp.NewServerTool[OperationLogsInput, any](
	"operation_logs",
	"Get 1Panel operation logs (audit trail of all actions performed)",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[OperationLogsInput]) (*mcp.CallToolResultFor[any], error) {
		payload := map[string]interface{}{
			"page":     1,
			"pageSize": 100,
			"group":    params.Arguments.Group,
			"orderBy":  "name",
			"order":    "ascending",
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/logs/tasks/search", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type OperationLogsInput struct {
	Group string `json:"group,omitempty" jsonschema:"filter by group"`
}

var LoginLogsTool = mcp.NewServerTool[LoginLogsInput, any](
	"login_logs",
	"Get 1Panel login history logs",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[LoginLogsInput]) (*mcp.CallToolResultFor[any], error) {
		payload := map[string]interface{}{
			"page":     1,
			"pageSize": 100,
			"ip":       params.Arguments.IP,
			"status":   params.Arguments.Status,
			"orderBy":  "name",
			"order":    "ascending",
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/logs/tasks/search", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type LoginLogsInput struct {
	IP     string `json:"ip,omitempty" jsonschema:"filter by IP address"`
	Status string `json:"status,omitempty" jsonschema:"filter: Success, Failed"`
}

var SystemLogsTool = mcp.NewServerTool[SystemLogsInput, any](
	"system_logs",
	"Get 1Panel system logs",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[SystemLogsInput]) (*mcp.CallToolResultFor[any], error) {
		var result interface{}
		client := utils.NewPanelClient("GET", "/logs/system/files")
		return client.Request(&result)
	},
)

type SystemLogsInput struct{}
