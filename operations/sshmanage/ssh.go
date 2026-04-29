package sshmanage

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/Aurora100729/mcp-1panel-full/utils"
)

var SSHInfoTool = mcp.NewServerTool[SSHInfoInput, any](
	"ssh_info",
	"Get SSH service configuration and status on the 1Panel server",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[SSHInfoInput]) (*mcp.CallToolResultFor[any], error) {
		var result interface{}
		client := utils.NewPanelClient("POST", "/hosts/ssh/search")
		return client.Request(&result)
	},
)

type SSHInfoInput struct{}

var SSHOperateTool = mcp.NewServerTool[SSHOperateInput, any](
	"ssh_operate",
	"[DANGEROUS] Start, stop, restart or configure SSH service on the 1Panel server",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[SSHOperateInput]) (*mcp.CallToolResultFor[any], error) {
		payload := map[string]interface{}{
			"operate": params.Arguments.Operate,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/hosts/ssh/operate", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type SSHOperateInput struct {
	Operate string `json:"operate" jsonschema:"operation: start, stop, restart"`
}

var SSHUpdateTool = mcp.NewServerTool[SSHUpdateInput, any](
	"ssh_update",
	"[DANGEROUS] Update SSH configuration (port, auth method, etc.)",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[SSHUpdateInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"key":   input.Key,
			"value": input.Value,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/hosts/ssh/update", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type SSHUpdateInput struct {
	Key   string `json:"key" jsonschema:"SSH config key: Port, PasswordAuthentication, PubkeyAuthentication, PermitRootLogin, UseDNS, ListenAddress"`
	Value string `json:"value" jsonschema:"new value for the config key"`
}

var SSHLogsTool = mcp.NewServerTool[SSHLogsInput, any](
	"ssh_logs",
	"Get SSH login logs (successful and failed attempts)",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[SSHLogsInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		status := input.Status
		if status == "" {
			status = "All"
		}
		payload := map[string]interface{}{
			"page":     1,
			"pageSize": 500,
			"info":     input.Info,
			"Status":   status,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/hosts/ssh/log", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type SSHLogsInput struct {
	Info   string `json:"info,omitempty" jsonschema:"filter by IP or username"`
	Status string `json:"status,omitempty" jsonschema:"filter: All, Success, Failed"`
}

var SSHGenerateKeyTool = mcp.NewServerTool[SSHGenerateKeyInput, any](
	"ssh_generate_key",
	"[DANGEROUS] Generate a new SSH key pair on the 1Panel server",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[SSHGenerateKeyInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		encryptionMode := input.EncryptionMode
		if encryptionMode == "" {
			encryptionMode = "ed25519"
		}
		payload := map[string]interface{}{
			"encryptionMode": encryptionMode,
			"password":       input.Password,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/hosts/ssh/generate", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type SSHGenerateKeyInput struct {
	EncryptionMode string `json:"encryptionMode,omitempty" jsonschema:"key type: ed25519, ecdsa, rsa, dsa"`
	Password       string `json:"password,omitempty" jsonschema:"passphrase for the key"`
}
