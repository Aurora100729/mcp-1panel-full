package system

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/1Panel-dev/mcp-1panel/operations/types"
	"github.com/1Panel-dev/mcp-1panel/utils"
)

const (
	GetDashboardInfo = "get_dashboard_info"
)

var GetDashboardInfoTool = mcp.NewServerTool[GetDashboardInfoInput, any](
	GetDashboardInfo,
	"show dashboard info",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[GetDashboardInfoInput]) (*mcp.CallToolResultFor[any], error) {
		client := utils.NewPanelClient("GET", "/dashboard/base/all/all")
		info := &types.DashboardRes{}
		result, err := client.Request(info)
		if result != nil {
			result.StructuredContent = info
		}
		return result, err
	},
)

type GetDashboardInfoInput struct{}
