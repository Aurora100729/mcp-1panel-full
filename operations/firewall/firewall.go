package firewall

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/Aurora100729/mcp-1panel-full/utils"
)

var FirewallStatusTool = mcp.NewServerTool[FirewallStatusInput, any](
	"firewall_status",
	"Get firewall status (active/inactive) and basic info",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[FirewallStatusInput]) (*mcp.CallToolResultFor[any], error) {
		var result interface{}
		client := utils.NewPanelClient("POST", "/hosts/firewall/base")
		return client.Request(&result)
	},
)

type FirewallStatusInput struct{}

var FirewallOperateTool = mcp.NewServerTool[FirewallOperateInput, any](
	"firewall_operate",
	"[DANGEROUS] Enable or disable the system firewall",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[FirewallOperateInput]) (*mcp.CallToolResultFor[any], error) {
		payload := map[string]interface{}{
			"operation": params.Arguments.Operation,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/hosts/firewall/operate", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type FirewallOperateInput struct {
	Operation string `json:"operation" jsonschema:"operation: start, stop, restart, reload"`
}

var ListFirewallRulesTool = mcp.NewServerTool[ListFirewallRulesInput, any](
	"list_firewall_rules",
	"List firewall port rules",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ListFirewallRulesInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		ruleType := input.Type
		if ruleType == "" {
			ruleType = "port"
		}
		payload := map[string]interface{}{
			"page":     1,
			"pageSize": 500,
			"info":     input.Info,
			"type":     ruleType,
			"status":   input.Status,
			"strategy": input.Strategy,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/hosts/firewall/search", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type ListFirewallRulesInput struct {
	Info     string `json:"info,omitempty" jsonschema:"filter by port or description"`
	Type     string `json:"type,omitempty" jsonschema:"port type filter"`
	Status   string `json:"status,omitempty" jsonschema:"status filter"`
	Strategy string `json:"strategy,omitempty" jsonschema:"strategy: accept, drop"`
}

var CreateFirewallRuleTool = mcp.NewServerTool[CreateFirewallRuleInput, any](
	"create_firewall_rule",
	"[DANGEROUS] Create a new firewall port rule",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[CreateFirewallRuleInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"port":     input.Port,
			"protocol": input.Protocol,
			"strategy": input.Strategy,
			"address":  input.Address,
			"description": input.Description,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/hosts/firewall/port", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type CreateFirewallRuleInput struct {
	Port        string `json:"port" jsonschema:"port number or range (e.g. 80 or 8000-9000)"`
	Protocol    string `json:"protocol" jsonschema:"protocol: tcp, udp, tcp/udp"`
	Strategy    string `json:"strategy" jsonschema:"strategy: accept or drop"`
	Address     string `json:"address,omitempty" jsonschema:"source IP address or CIDR"`
	Description string `json:"description,omitempty" jsonschema:"rule description"`
}

var DeleteFirewallRuleTool = mcp.NewServerTool[DeleteFirewallRuleInput, any](
	"delete_firewall_rule",
	"[DANGEROUS] Delete a firewall port rule",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[DeleteFirewallRuleInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"port":     input.Port,
			"protocol": input.Protocol,
			"strategy": input.Strategy,
			"address":  input.Address,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/hosts/firewall/batch", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type DeleteFirewallRuleInput struct {
	Port     string `json:"port" jsonschema:"port number or range"`
	Protocol string `json:"protocol" jsonschema:"protocol: tcp, udp, tcp/udp"`
	Strategy string `json:"strategy" jsonschema:"strategy: accept or drop"`
	Address  string `json:"address,omitempty" jsonschema:"source IP address"`
}

var ListFirewallIPRulesTool = mcp.NewServerTool[ListFirewallIPRulesInput, any](
	"list_firewall_ip_rules",
	"List firewall IP address rules",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ListFirewallIPRulesInput]) (*mcp.CallToolResultFor[any], error) {
		payload := map[string]interface{}{
			"page":     1,
			"pageSize": 500,
			"type":     "address",
			"info":     params.Arguments.Info,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/hosts/firewall/search", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type ListFirewallIPRulesInput struct {
	Info string `json:"info,omitempty" jsonschema:"filter by IP or description"`
}

var CreateFirewallIPRuleTool = mcp.NewServerTool[CreateFirewallIPRuleInput, any](
	"create_firewall_ip_rule",
	"[DANGEROUS] Create a firewall IP rule (allow/deny specific IP)",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[CreateFirewallIPRuleInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"address":  input.Address,
			"strategy": input.Strategy,
			"description": input.Description,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/hosts/firewall/ip", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type CreateFirewallIPRuleInput struct {
	Address     string `json:"address" jsonschema:"IP address or CIDR block"`
	Strategy    string `json:"strategy" jsonschema:"strategy: accept or drop"`
	Description string `json:"description,omitempty" jsonschema:"rule description"`
}

var ListFirewallForwardsTool = mcp.NewServerTool[ListFirewallForwardsInput, any](
	"list_firewall_forwards",
	"List firewall port forwarding rules",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ListFirewallForwardsInput]) (*mcp.CallToolResultFor[any], error) {
		payload := map[string]interface{}{
			"page":     1,
			"pageSize": 500,
			"info":     params.Arguments.Info,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/hosts/firewall/forward", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type ListFirewallForwardsInput struct {
	Info string `json:"info,omitempty" jsonschema:"filter"`
}
