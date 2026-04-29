package toolbox

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/Aurora100729/mcp-1panel-full/utils"
)

var GetDNSTool = mcp.NewServerTool[GetDNSInput, any](
	"toolbox_dns",
	"Get or update DNS nameserver configuration",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[GetDNSInput]) (*mcp.CallToolResultFor[any], error) {
		var result interface{}
		client := utils.NewPanelClient("GET", "/toolbox/dns/search")
		return client.Request(&result)
	},
)

type GetDNSInput struct{}

var UpdateDNSTool = mcp.NewServerTool[UpdateDNSInput, any](
	"toolbox_dns_update",
	"[DANGEROUS] Update DNS nameserver configuration",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[UpdateDNSInput]) (*mcp.CallToolResultFor[any], error) {
		payload := map[string]interface{}{
			"dns": params.Arguments.DNS,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/toolbox/dns/update", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type UpdateDNSInput struct {
	DNS []string `json:"dns" jsonschema:"list of DNS servers, e.g. [\"8.8.8.8\", \"1.1.1.1\"]"`
}

var GetHostsTool = mcp.NewServerTool[GetHostsInput, any](
	"toolbox_hosts",
	"Get /etc/hosts file content",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[GetHostsInput]) (*mcp.CallToolResultFor[any], error) {
		var result interface{}
		client := utils.NewPanelClient("GET", "/toolbox/hosts")
		return client.Request(&result)
	},
)

type GetHostsInput struct{}

var UpdateHostsTool = mcp.NewServerTool[UpdateHostsInput, any](
	"toolbox_hosts_update",
	"[DANGEROUS] Update /etc/hosts file entries",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[UpdateHostsInput]) (*mcp.CallToolResultFor[any], error) {
		payload := map[string]interface{}{
			"content": params.Arguments.Content,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/toolbox/hosts/update", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type UpdateHostsInput struct {
	Content string `json:"content" jsonschema:"full /etc/hosts content"`
}

var GetSwapTool = mcp.NewServerTool[GetSwapInput, any](
	"toolbox_swap",
	"Get swap configuration status",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[GetSwapInput]) (*mcp.CallToolResultFor[any], error) {
		var result interface{}
		client := utils.NewPanelClient("GET", "/toolbox/swap")
		return client.Request(&result)
	},
)

type GetSwapInput struct{}

var SwapOperateTool = mcp.NewServerTool[SwapOperateInput, any](
	"toolbox_swap_operate",
	"[DANGEROUS] Enable, disable, or resize swap space",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[SwapOperateInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"operate": input.Operate,
			"path":    input.Path,
			"size":    input.Size,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/toolbox/swap/operate", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type SwapOperateInput struct {
	Operate string `json:"operate" jsonschema:"operation: create, delete"`
	Path    string `json:"path,omitempty" jsonschema:"swap file path, e.g. /swapfile"`
	Size    int    `json:"size,omitempty" jsonschema:"swap size in MB"`
}

var GetTimezoneTool = mcp.NewServerTool[GetTimezoneInput, any](
	"toolbox_timezone",
	"Get current system timezone and NTP configuration",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[GetTimezoneInput]) (*mcp.CallToolResultFor[any], error) {
		var result interface{}
		client := utils.NewPanelClient("GET", "/toolbox/time/option")
		return client.Request(&result)
	},
)

type GetTimezoneInput struct{}

var Fail2BanStatusTool = mcp.NewServerTool[Fail2BanStatusInput, any](
	"toolbox_fail2ban_status",
	"Get Fail2Ban status and configuration",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[Fail2BanStatusInput]) (*mcp.CallToolResultFor[any], error) {
		var result interface{}
		client := utils.NewPanelClient("GET", "/toolbox/fail2ban/base")
		return client.Request(&result)
	},
)

type Fail2BanStatusInput struct{}

var Fail2BanOperateTool = mcp.NewServerTool[Fail2BanOperateInput, any](
	"toolbox_fail2ban_operate",
	"[DANGEROUS] Start, stop or configure Fail2Ban",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[Fail2BanOperateInput]) (*mcp.CallToolResultFor[any], error) {
		payload := map[string]interface{}{
			"operate": params.Arguments.Operate,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/toolbox/fail2ban/operate", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type Fail2BanOperateInput struct {
	Operate string `json:"operate" jsonschema:"operation: start, stop, restart, enable, disable"`
}

var ClamAVScanTool = mcp.NewServerTool[ClamAVScanInput, any](
	"toolbox_clamav_scan",
	"[DANGEROUS] Trigger ClamAV antivirus scan",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ClamAVScanInput]) (*mcp.CallToolResultFor[any], error) {
		payload := map[string]interface{}{
			"path": params.Arguments.Path,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/toolbox/clam/operate", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type ClamAVScanInput struct {
	Path string `json:"path,omitempty" jsonschema:"path to scan, default /"`
}
