package monitor

import (
	"context"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/Aurora100729/mcp-1panel-full/utils"
)

var MonitorSearchTool = mcp.NewServerTool[MonitorSearchInput, any](
	"monitor_search",
	"Get server monitoring data: CPU, memory, disk IO, network history over time",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[MonitorSearchInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		param := input.Param
		if param == "" {
			param = "all"
		}
		startTime := input.StartTime
		endTime := input.EndTime
		if startTime == "" {
			startTime = time.Now().Add(-1 * time.Hour).Format(time.RFC3339)
		}
		if endTime == "" {
			endTime = time.Now().Format(time.RFC3339)
		}
		payload := map[string]interface{}{
			"param":     param,
			"info":      input.Info,
			"startTime": startTime,
			"endTime":   endTime,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/hosts/monitor/search", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type MonitorSearchInput struct {
	Param     string `json:"param,omitempty" jsonschema:"monitor type: all, cpu, memory, load, io, network"`
	Info      string `json:"info,omitempty" jsonschema:"additional filter"`
	StartTime string `json:"startTime,omitempty" jsonschema:"start time ISO 8601"`
	EndTime   string `json:"endTime,omitempty" jsonschema:"end time ISO 8601"`
}

var MonitorCleanTool = mcp.NewServerTool[MonitorCleanInput, any](
	"monitor_clean",
	"[DANGEROUS] Clean up monitoring data",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[MonitorCleanInput]) (*mcp.CallToolResultFor[any], error) {
		var result interface{}
		client := utils.NewPanelClient("POST", "/hosts/monitor/clean")
		return client.Request(&result)
	},
)

type MonitorCleanInput struct{}
