package system

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/Aurora100729/mcp-1panel-full/operations/types"
	"github.com/Aurora100729/mcp-1panel-full/utils"
)

const (
	GetSystemInfo = "get_system_info"
)

var GetSystemInfoTool = mcp.NewServerTool[GetSystemInfoInput, any](
	GetSystemInfo,
	"show host system information, The unit of diskSize is bytes",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[GetSystemInfoInput]) (*mcp.CallToolResultFor[any], error) {
		client := utils.NewPanelClient("GET", "/dashboard/base/os")
		osInfo := &types.OsInfoRes{}
		result, err := client.Request(osInfo)
		if result != nil {
			result.StructuredContent = osInfo
		}
		return result, err
	},
)

type GetSystemInfoInput struct{}
