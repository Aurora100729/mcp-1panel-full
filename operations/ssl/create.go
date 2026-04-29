package ssl

import (
	"context"
	"errors"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/Aurora100729/mcp-1panel-full/operations/types"
	"github.com/Aurora100729/mcp-1panel-full/utils"
)

const (
	CreateSSL = "create_ssl"
)

var CreateSSLTool = mcp.NewServerTool[CreateSSLInput, any](
	CreateSSL,
	"create ssl",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[CreateSSLInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		if input.Domain == "" {
			err := errors.New("domain is required")
			return &mcp.CallToolResult{
				Content: []mcp.Content{&mcp.TextContent{Text: err.Error()}},
				IsError: true,
			}, err
		}
		if input.Provider == "" {
			err := errors.New("provider is required")
			return &mcp.CallToolResult{
				Content: []mcp.Content{&mcp.TextContent{Text: err.Error()}},
				IsError: true,
			}, err
		}
		if input.Provider != "dnsAccount" && input.Provider != "http" {
			err := errors.New("provider must be dnsAccount or http")
			return &mcp.CallToolResult{
				Content: []mcp.Content{&mcp.TextContent{Text: err.Error()}},
				IsError: true,
			}, err
		}

		acmeRes := &types.ListAcmeRes{}
		pageReq := &types.PageRequest{
			Page:     1,
			PageSize: 500,
		}
		result, err := utils.NewPanelClient("POST", "/websites/acme/search", utils.WithPayload(pageReq)).Request(acmeRes)
		if err != nil {
			return result, err
		}
		if len(acmeRes.Data.Items) == 0 {
			err := errors.New("no acme account found")
			return &mcp.CallToolResult{
				Content: []mcp.Content{&mcp.TextContent{Text: err.Error()}},
				IsError: true,
			}, err
		}
		acme := acmeRes.Data.Items[0]

		var dnsAccountID uint
		if input.Provider == "dnsAccount" {
			dnsAccountRes := &types.ListDNSAccountRes{}
			result, err = utils.NewPanelClient("POST", "/websites/dns/search", utils.WithPayload(pageReq)).Request(dnsAccountRes)
			if err != nil {
				return result, err
			}
			if len(dnsAccountRes.Data.Items) == 0 {
				err := errors.New("no dns account found")
				return &mcp.CallToolResult{
					Content: []mcp.Content{&mcp.TextContent{Text: err.Error()}},
					IsError: true,
				}, err
			}
			dnsName := input.DnsAccount
			if dnsName != "" {
				checkName := strings.ToLower(dnsName)
				for _, dnsAccount := range dnsAccountRes.Data.Items {
					if strings.Contains(strings.ToLower(dnsAccount.Name), checkName) || strings.Contains(strings.ToLower(dnsAccount.Type), checkName) {
						dnsAccountID = dnsAccount.ID
						break
					}
				}
			}
			if dnsAccountID == 0 {
				dnsAccountID = dnsAccountRes.Data.Items[0].ID
			}
		}

		req := &types.CreateSSLRequest{
			PrimaryDomain: input.Domain,
			Provider:      input.Provider,
			AcmeAccountID: acme.ID,
			DnsAccountID:  dnsAccountID,
			KeyType:       "P256",
		}
		res := &types.Response{}
		result, err = utils.NewPanelClient("POST", "/websites/ssl", utils.WithPayload(req)).Request(res)
		if result != nil {
			result.StructuredContent = res
		}
		return result, err
	},
)

type CreateSSLInput struct {
	Domain     string `json:"domain" jsonschema:"domain"`
	Provider   string `json:"provider" jsonschema:"provider support dnsAccount,http"`
	DnsAccount string `json:"dnsAccount,omitempty" jsonschema:"dnsAccount"`
}
