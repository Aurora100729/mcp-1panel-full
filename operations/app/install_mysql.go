package app

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/Aurora100729/mcp-1panel-full/operations/types"
	"github.com/Aurora100729/mcp-1panel-full/utils"
)

const (
	InstallMySQL = "install_mysql"
)

var InstallMySQLTool = mcp.NewServerTool[InstallMySQLInput, any](
	InstallMySQL,
	"install mysql, if not set name, default is mysql, if not set version, default is '', if not set root_password, default is '')",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[InstallMySQLInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		name := input.Name
		if name == "" {
			name = "mysql"
		}

		version := input.Version
		if version == "latest" {
			version = ""
		}

		appRes := &types.AppRes{}
		result, err := utils.NewPanelClient("GET", "/apps/mysql").Request(appRes)
		if err != nil {
			return result, err
		}
		exist := false
		for _, v := range appRes.Data.Versions {
			if v == version || strings.Contains(v, version) {
				version = v
				exist = true
				break
			}
		}
		if !exist {
			err := errors.New("version not found")
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{Text: err.Error()},
				},
				IsError: true,
			}, err
		}
		if version == "" {
			version = appRes.Data.Versions[0]
		}
		appID := appRes.Data.ID
		appDetailURL := fmt.Sprintf("/apps/detail/%d/%s/app", appID, version)
		appDetailRes := &types.AppDetailRes{}
		result, err = utils.NewPanelClient("GET", appDetailURL).Request(appDetailRes)
		if err != nil {
			return result, err
		}
		appDetailID := appDetailRes.Data.ID

		port := input.Port
		if port == 0 {
			port = 3306
		}

		rootPassword := input.RootPassword
		if rootPassword == "" {
			rootPassword = fmt.Sprintf("mysql_%s", utils.GetRandomStr(6))
		}

		req := &types.AppInstallCreate{
			AppDetailID: appDetailID,
			Name:        name,
			Params: map[string]interface{}{
				"PANEL_APP_PORT_HTTP":    port,
				"PANEL_DB_ROOT_PASSWORD": rootPassword,
			},
		}
		res := &types.Response{}
		result, err = utils.NewPanelClient("POST", "/apps/install", utils.WithPayload(req)).Request(res)
		if result != nil {
			result.StructuredContent = res
		}
		return result, err
	},
)

type InstallMySQLInput struct {
	Name         string  `json:"name" jsonschema:"mysql name"`
	Version      string  `json:"version,omitempty" jsonschema:"mysql version, not support latest version"`
	RootPassword string  `json:"root_password,omitempty" jsonschema:"mysql root password"`
	Port         float64 `json:"port,omitempty" jsonschema:"mysql port"`
}
