package website

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/1Panel-dev/mcp-1panel/operations/types"
	"github.com/1Panel-dev/mcp-1panel/utils"
)

const (
	ListWebsites = "list_websites"
)

var ListWebsitesTool = mcp.NewServerTool[ListWebsitesInput, any](
	ListWebsites,
	"list websites",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ListWebsitesInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		req := &types.ListWebsiteRequest{
			Order:   "null",
			OrderBy: "created_at",
			PageRequest: types.PageRequest{
				Page:     1,
				PageSize: 500,
				Name:     input.Name,
			},
		}
		client := utils.NewPanelClient("POST", "/websites/search", utils.WithPayload(req))
		listWebsiteRes := &types.ListWebsiteRes{}
		result, err := client.Request(listWebsiteRes)
		if result != nil {
			result.StructuredContent = listWebsiteRes
		}
		return result, err
	},
)

type ListWebsitesInput struct {
	Name string `json:"name,omitempty" jsonschema:"search by website name"`
}
