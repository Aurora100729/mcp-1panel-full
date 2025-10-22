package app

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/1Panel-dev/mcp-1panel/operations/types"
	"github.com/1Panel-dev/mcp-1panel/utils"
)

const (
	ListInstalledApps = "list_installed_apps"
)

var ListInstalledAppsTool = mcp.NewServerTool[ListInstalledAppsInput, any](
	ListInstalledApps,
	"list installed apps",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ListInstalledAppsInput]) (*mcp.CallToolResultFor[any], error) {
		req := &types.PageRequest{
			Page:     1,
			PageSize: 500,
		}
		appListRes := &types.AppInstalledListResponse{}
		result, err := utils.NewPanelClient("POST", "/apps/installed/search", utils.WithPayload(req)).Request(appListRes)
		if result != nil {
			result.StructuredContent = appListRes
		}
		return result, err
	},
)

type ListInstalledAppsInput struct {
}
