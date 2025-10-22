package website

import (
	"context"
	"errors"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/1Panel-dev/mcp-1panel/operations/types"
	"github.com/1Panel-dev/mcp-1panel/utils"
)

const (
	CreateWebsite = "create_website"
)

var CreateWebsiteTool = mcp.NewServerTool[CreateWebsiteInput, any](
	CreateWebsite,
	"create website",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[CreateWebsiteInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		if input.Domain == "" {
			err := errors.New("domain is required")
			return &mcp.CallToolResult{
				Content: []mcp.Content{&mcp.TextContent{Text: err.Error()}},
				IsError: true,
			}, err
		}

		domain := input.Domain
		alias := domain
		var proxyAddress string
		if input.WebsiteType == "proxy" {
			if input.ProxyAddress == "" {
				err := errors.New("proxy_address is required")
				return &mcp.CallToolResult{
					Content: []mcp.Content{&mcp.TextContent{Text: err.Error()}},
					IsError: true,
				}, err
			}
			proxyAddress = input.ProxyAddress
		}

		groupReq := &types.GroupRequest{
			Type: "website",
		}
		groupRes := &types.GroupRes{}
		result, err := utils.NewPanelClient("POST", "/groups/search", utils.WithPayload(groupReq)).Request(groupRes)
		if err != nil {
			return result, err
		}
		var groupID uint
		for _, group := range groupRes.Data {
			if group.IsDefault {
				groupID = group.ID
				break
			}
		}

		req := &types.CreateWebsiteRequest{
			Domains:  []types.WebsiteDomain{
				{
					Domain: domain,
					Port:   80,
					SSL:    false,
				},
			},
			Alias:          alias,
			Type:           input.WebsiteType,
			WebsiteGroupID: groupID,
			Proxy:          proxyAddress,
			AppType:        "new",
		}
		res := &types.Response{}
		result, err = utils.NewPanelClient("POST", "/websites", utils.WithPayload(req)).Request(res)
		if result != nil {
			result.StructuredContent = res
		}
		return result, err
	},
)

type CreateWebsiteInput struct {
	Domain       string `json:"domain" jsonschema:"domain,required"`
	WebsiteType  string `json:"website_type" jsonschema:"website type,only support static and proxy,required"`
	ProxyAddress string `json:"proxy_address,omitempty" jsonschema:"proxy address,only support for proxy website"`
}
