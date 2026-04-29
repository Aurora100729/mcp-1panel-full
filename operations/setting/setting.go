package setting

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/Aurora100729/mcp-1panel-full/utils"
)

var GetSettingsTool = mcp.NewServerTool[GetSettingsInput, any](
	"get_settings",
	"Get 1Panel system settings (port, language, theme, security settings, etc.)",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[GetSettingsInput]) (*mcp.CallToolResultFor[any], error) {
		var result interface{}
		client := utils.NewPanelClient("POST", "/settings/search")
		return client.Request(&result)
	},
)

type GetSettingsInput struct{}

var UpdateSettingTool = mcp.NewServerTool[UpdateSettingInput, any](
	"update_setting",
	"[DANGEROUS] Update a 1Panel system setting",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[UpdateSettingInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"key":   input.Key,
			"value": input.Value,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/settings/update", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type UpdateSettingInput struct {
	Key   string `json:"key" jsonschema:"setting key: UserName, Password, Email, SessionTimeout, LocalTime, PanelName, Theme, Language, ServerPort, SecurityEntrance, ExpirationDays, ComplexityVerification, MFAStatus, MonitorStatus, MonitorStoreDays, MessageType, AllowIPs, BindDomain"`
	Value string `json:"value" jsonschema:"new value for the setting"`
}

var UpdatePasswordTool = mcp.NewServerTool[UpdatePasswordInput, any](
	"update_password",
	"[DANGEROUS] Update 1Panel admin password",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[UpdatePasswordInput]) (*mcp.CallToolResultFor[any], error) {
		payload := map[string]interface{}{
			"oldPassword": params.Arguments.OldPassword,
			"newPassword": params.Arguments.NewPassword,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/settings/password/update", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type UpdatePasswordInput struct {
	OldPassword string `json:"oldPassword" jsonschema:"current password"`
	NewPassword string `json:"newPassword" jsonschema:"new password"`
}

var UpdatePortTool = mcp.NewServerTool[UpdatePortInput, any](
	"update_panel_port",
	"[DANGEROUS] Update 1Panel listening port",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[UpdatePortInput]) (*mcp.CallToolResultFor[any], error) {
		payload := map[string]interface{}{
			"serverPort": params.Arguments.Port,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/settings/port/update", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type UpdatePortInput struct {
	Port int `json:"serverPort" jsonschema:"new port number for 1Panel"`
}

var PanelUpgradeTool = mcp.NewServerTool[PanelUpgradeInput, any](
	"panel_upgrade",
	"[DANGEROUS] Upgrade 1Panel to the latest version or a specified version",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[PanelUpgradeInput]) (*mcp.CallToolResultFor[any], error) {
		payload := map[string]interface{}{
			"version": params.Arguments.Version,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/settings/upgrade", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type PanelUpgradeInput struct {
	Version string `json:"version,omitempty" jsonschema:"target version, empty for latest"`
}
