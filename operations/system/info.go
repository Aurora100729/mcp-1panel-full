package system

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/1Panel-dev/mcp-1panel/operations/types"
	"github.com/1Panel-dev/mcp-1panel/utils"
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
