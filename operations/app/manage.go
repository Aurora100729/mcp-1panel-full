package app

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/Aurora100729/mcp-1panel-full/utils"
)

var AppStoreTool = mcp.NewServerTool[AppStoreInput, any](
	"app_store_list",
	"List available apps in the 1Panel app store",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[AppStoreInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"page":     1,
			"pageSize": 500,
			"name":     input.Name,
			"type":     input.Type,
			"tags":     input.Tags,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/apps/search", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type AppStoreInput struct {
	Name string   `json:"name,omitempty" jsonschema:"filter by app name"`
	Type string   `json:"type,omitempty" jsonschema:"filter by app type"`
	Tags []string `json:"tags,omitempty" jsonschema:"filter by tags"`
}

var AppDetailTool = mcp.NewServerTool[AppDetailInput, any](
	"app_detail",
	"Get detailed information about a specific app including available versions",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[AppDetailInput]) (*mcp.CallToolResultFor[any], error) {
		path := fmt.Sprintf("/apps/%s", params.Arguments.Key)
		var result interface{}
		client := utils.NewPanelClient("GET", path)
		return client.Request(&result)
	},
)

type AppDetailInput struct {
	Key string `json:"key" jsonschema:"app key/name, e.g. mysql, redis, nginx, wordpress"`
}

var AppInstalledDetailTool = mcp.NewServerTool[AppInstalledDetailInput, any](
	"app_installed_detail",
	"Get details of an installed app instance including params, env, and status",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[AppInstalledDetailInput]) (*mcp.CallToolResultFor[any], error) {
		path := fmt.Sprintf("/apps/installed/%d", params.Arguments.ID)
		var result interface{}
		client := utils.NewPanelClient("GET", path)
		return client.Request(&result)
	},
)

type AppInstalledDetailInput struct {
	ID int `json:"id" jsonschema:"installed app ID"`
}

var AppOperateTool = mcp.NewServerTool[AppOperateInput, any](
	"app_operate",
	"[DANGEROUS] Operate on an installed app: start, stop, restart, delete, sync, upgrade",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[AppOperateInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"installId":   input.InstallID,
			"operate":     input.Operate,
			"deleteBackup": input.DeleteBackup,
			"forceDelete":  input.ForceDelete,
			"deleteDB":     input.DeleteDB,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/apps/installed/op", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type AppOperateInput struct {
	InstallID    int    `json:"installId" jsonschema:"installed app ID"`
	Operate      string `json:"operate" jsonschema:"operation: start, stop, restart, delete, sync, upgrade, rebuild"`
	DeleteBackup bool   `json:"deleteBackup,omitempty" jsonschema:"delete backups when deleting app"`
	ForceDelete  bool   `json:"forceDelete,omitempty" jsonschema:"force delete"`
	DeleteDB     bool   `json:"deleteDB,omitempty" jsonschema:"delete database when deleting app"`
}

var AppUpdateParamsTool = mcp.NewServerTool[AppUpdateParamsInput, any](
	"app_update_params",
	"[DANGEROUS] Update parameters/environment variables of an installed app",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[AppUpdateParamsInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"installId": input.InstallID,
			"params":    input.Params,
		}
		var result interface{}
		client := utils.NewPanelClient("PUT", "/apps/installed/params/update", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type AppUpdateParamsInput struct {
	InstallID int                    `json:"installId" jsonschema:"installed app ID"`
	Params    map[string]interface{} `json:"params" jsonschema:"key-value parameters to update"`
}

var AppInstalledParamsTool = mcp.NewServerTool[AppInstalledParamsInput, any](
	"app_installed_params",
	"Get current parameters/env of an installed app",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[AppInstalledParamsInput]) (*mcp.CallToolResultFor[any], error) {
		path := fmt.Sprintf("/apps/installed/params/%d", params.Arguments.ID)
		var result interface{}
		client := utils.NewPanelClient("GET", path)
		return client.Request(&result)
	},
)

type AppInstalledParamsInput struct {
	ID int `json:"id" jsonschema:"installed app ID"`
}
