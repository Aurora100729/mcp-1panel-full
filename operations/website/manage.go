package website

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/Aurora100729/mcp-1panel-full/utils"
)

var DeleteWebsiteTool = mcp.NewServerTool[DeleteWebsiteInput, any](
	"delete_website",
	"[DANGEROUS] Delete a website and optionally its associated resources",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[DeleteWebsiteInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"id":            input.ID,
			"deleteApp":     input.DeleteApp,
			"deleteBackup":  input.DeleteBackup,
			"forceDelete":   input.ForceDelete,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/websites/del", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type DeleteWebsiteInput struct {
	ID           int  `json:"id" jsonschema:"website ID to delete"`
	DeleteApp    bool `json:"deleteApp,omitempty" jsonschema:"also delete associated app"`
	DeleteBackup bool `json:"deleteBackup,omitempty" jsonschema:"also delete backups"`
	ForceDelete  bool `json:"forceDelete,omitempty" jsonschema:"force delete even if errors"`
}

var GetWebsiteConfigTool = mcp.NewServerTool[GetWebsiteConfigInput, any](
	"get_website_config",
	"Get Nginx configuration for a website",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[GetWebsiteConfigInput]) (*mcp.CallToolResultFor[any], error) {
		path := fmt.Sprintf("/websites/%d/config/%s", params.Arguments.ID, params.Arguments.Type)
		var result interface{}
		client := utils.NewPanelClient("GET", path)
		return client.Request(&result)
	},
)

type GetWebsiteConfigInput struct {
	ID   int    `json:"id" jsonschema:"website ID"`
	Type string `json:"type,omitempty" jsonschema:"config type to get, default is the nginx conf"`
}

var UpdateWebsiteConfigTool = mcp.NewServerTool[UpdateWebsiteConfigInput, any](
	"update_website_config",
	"[DANGEROUS] Update Nginx configuration for a website",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[UpdateWebsiteConfigInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		path := fmt.Sprintf("/websites/%d/config/update", input.ID)
		payload := map[string]interface{}{
			"content": input.Content,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", path, utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type UpdateWebsiteConfigInput struct {
	ID      int    `json:"id" jsonschema:"website ID"`
	Content string `json:"content" jsonschema:"new nginx configuration content"`
}

var WebsiteOperateTool = mcp.NewServerTool[WebsiteOperateInput, any](
	"website_operate",
	"[DANGEROUS] Start, stop or delete a website",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[WebsiteOperateInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"id":      input.ID,
			"operate": input.Operate,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/websites/operate", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type WebsiteOperateInput struct {
	ID      int    `json:"id" jsonschema:"website ID"`
	Operate string `json:"operate" jsonschema:"operation: start, stop"`
}

var GetWebsiteHTTPSTool = mcp.NewServerTool[GetWebsiteHTTPSInput, any](
	"get_website_https",
	"Get HTTPS/SSL configuration for a website",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[GetWebsiteHTTPSInput]) (*mcp.CallToolResultFor[any], error) {
		path := fmt.Sprintf("/websites/%d/https", params.Arguments.ID)
		var result interface{}
		client := utils.NewPanelClient("GET", path)
		return client.Request(&result)
	},
)

type GetWebsiteHTTPSInput struct {
	ID int `json:"id" jsonschema:"website ID"`
}

var UpdateWebsiteHTTPSTool = mcp.NewServerTool[UpdateWebsiteHTTPSInput, any](
	"update_website_https",
	"[DANGEROUS] Enable or configure HTTPS/SSL for a website",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[UpdateWebsiteHTTPSInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		path := fmt.Sprintf("/websites/%d/https", input.ID)
		payload := map[string]interface{}{
			"enable":      input.Enable,
			"websiteSSLId": input.WebsiteSSLID,
			"type":        input.Type,
			"httpConfig":  input.HTTPConfig,
			"SSLProtocol": input.SSLProtocol,
			"algorithm":   input.Algorithm,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", path, utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type UpdateWebsiteHTTPSInput struct {
	ID           int      `json:"id" jsonschema:"website ID"`
	Enable       bool     `json:"enable" jsonschema:"enable HTTPS"`
	WebsiteSSLID int      `json:"websiteSSLId,omitempty" jsonschema:"SSL certificate ID to use"`
	Type         string   `json:"type,omitempty" jsonschema:"SSL type"`
	HTTPConfig   string   `json:"httpConfig,omitempty" jsonschema:"HTTP redirect config: HTTPSOnly, HTTPAlso, HTTPToHTTPS"`
	SSLProtocol  []string `json:"SSLProtocol,omitempty" jsonschema:"SSL protocols to enable"`
	Algorithm    string   `json:"algorithm,omitempty" jsonschema:"SSL algorithm"`
}

var ListWebsiteDomainsTool = mcp.NewServerTool[ListWebsiteDomainsInput, any](
	"list_website_domains",
	"List domains bound to a website",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ListWebsiteDomainsInput]) (*mcp.CallToolResultFor[any], error) {
		path := fmt.Sprintf("/websites/domains/%d", params.Arguments.WebsiteID)
		var result interface{}
		client := utils.NewPanelClient("GET", path)
		return client.Request(&result)
	},
)

type ListWebsiteDomainsInput struct {
	WebsiteID int `json:"websiteID" jsonschema:"website ID"`
}

var CreateWebsiteDomainTool = mcp.NewServerTool[CreateWebsiteDomainInput, any](
	"create_website_domain",
	"Add a new domain/subdomain to a website",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[CreateWebsiteDomainInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"websiteID": input.WebsiteID,
			"domain":    input.Domain,
			"port":      input.Port,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/websites/domains", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type CreateWebsiteDomainInput struct {
	WebsiteID int    `json:"websiteID" jsonschema:"website ID"`
	Domain    string `json:"domain" jsonschema:"domain name to add"`
	Port      int    `json:"port,omitempty" jsonschema:"port, default 80"`
}
