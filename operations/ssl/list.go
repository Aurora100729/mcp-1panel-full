package ssl

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/Aurora100729/mcp-1panel-full/operations/types"
	"github.com/Aurora100729/mcp-1panel-full/utils"
)

const (
	ListSSLs = "list_ssls"
)

var ListSSLsTool = mcp.NewServerTool[ListSSLsInput, any](
	ListSSLs,
	"list ssls",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ListSSLsInput]) (*mcp.CallToolResultFor[any], error) {
		req := &types.PageRequest{
			Page:     1,
			PageSize: 500,
		}
		listWebsiteSSLRes := &types.ListWebsiteSSLRes{}
		result, err := utils.NewPanelClient("POST", "/websites/ssl/search", utils.WithPayload(req)).Request(listWebsiteSSLRes)
		if result != nil {
			result.StructuredContent = listWebsiteSSLRes
		}
		return result, err
	},
)

type ListSSLsInput struct {
}
