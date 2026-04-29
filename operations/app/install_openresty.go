package app

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/Aurora100729/mcp-1panel-full/operations/types"
	"github.com/Aurora100729/mcp-1panel-full/utils"
)

const (
	InstallOpenResty = "install_openresty"
)

var InstallOpenRestyTool = mcp.NewServerTool[InstallOpenRestyInput, any](
	InstallOpenResty,
	"install openresty, if not set name, default is openresty, if not set http_port, default is 80, if not set https_port, default is 443",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[InstallOpenRestyInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		name := input.Name
		if name == "" {
			name = "openresty"
		}

		httpPort := input.HttpPort
		if httpPort == 0 {
			httpPort = 80
		}

		httpsPort := input.HttpsPort
		if httpsPort == 0 {
			httpsPort = 443
		}

		appRes := &types.AppRes{}
		result, err := utils.NewPanelClient("GET", "/apps/openresty").Request(appRes)
		if err != nil {
			return result, err
		}
		version := appRes.Data.Versions[0]
		appID := appRes.Data.ID
		appDetailURL := fmt.Sprintf("/apps/detail/%d/%s/app", appID, version)
		appDetailRes := &types.AppDetailRes{}
		result, err = utils.NewPanelClient("GET", appDetailURL).Request(appDetailRes)
		if err != nil {
			return result, err
		}

		appDetailID := appDetailRes.Data.ID

		req := &types.AppInstallCreate{
			AppDetailID: appDetailID,
			Name:        name,
			Params: map[string]interface{}{
				"PANEL_APP_PORT_HTTP":  httpPort,
				"PANEL_APP_PORT_HTTPS": httpsPort,
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

type InstallOpenRestyInput struct {
	Name      string  `json:"name,omitempty" jsonschema:"openresty name"`
	HttpPort  float64 `json:"http_port,omitempty" jsonschema:"openresty http port"`
	HttpsPort float64 `json:"https_port,omitempty" jsonschema:"openresty https port"`
}
