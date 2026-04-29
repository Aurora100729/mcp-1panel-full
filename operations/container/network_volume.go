package container

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/Aurora100729/mcp-1panel-full/utils"
)

var ListNetworksTool = mcp.NewServerTool[ListNetworksInput, any](
	"list_networks",
	"List all Docker networks",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ListNetworksInput]) (*mcp.CallToolResultFor[any], error) {
		payload := map[string]interface{}{
			"page":     1,
			"pageSize": 500,
			"name":     params.Arguments.Name,
			"orderBy":  "created_at",
			"order":    "null",
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/containers/network/search", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type ListNetworksInput struct {
	Name string `json:"name,omitempty" jsonschema:"filter by network name"`
}

var CreateNetworkTool = mcp.NewServerTool[CreateNetworkInput, any](
	"create_network",
	"Create a new Docker network",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[CreateNetworkInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"name":    input.Name,
			"driver":  input.Driver,
			"subnet":  input.Subnet,
			"gateway": input.Gateway,
			"ipRange": input.IPRange,
			"options": input.Options,
			"labels":  input.Labels,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/containers/network", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type CreateNetworkInput struct {
	Name    string   `json:"name" jsonschema:"network name"`
	Driver  string   `json:"driver,omitempty" jsonschema:"driver: bridge, overlay, macvlan, host, none. Default: bridge"`
	Subnet  string   `json:"subnet,omitempty" jsonschema:"subnet CIDR, e.g. 172.20.0.0/16"`
	Gateway string   `json:"gateway,omitempty" jsonschema:"gateway IP"`
	IPRange string   `json:"ipRange,omitempty" jsonschema:"IP range"`
	Options []string `json:"options,omitempty" jsonschema:"driver options"`
	Labels  []string `json:"labels,omitempty" jsonschema:"network labels"`
}

var DeleteNetworkTool = mcp.NewServerTool[DeleteNetworkInput, any](
	"delete_network",
	"[DANGEROUS] Delete a Docker network by ID",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[DeleteNetworkInput]) (*mcp.CallToolResultFor[any], error) {
		payload := map[string]interface{}{
			"names": params.Arguments.Names,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/containers/network/del", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type DeleteNetworkInput struct {
	Names []string `json:"names" jsonschema:"network names or IDs to delete"`
}

var ListVolumesTool = mcp.NewServerTool[ListVolumesInput, any](
	"list_volumes",
	"List all Docker volumes",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ListVolumesInput]) (*mcp.CallToolResultFor[any], error) {
		payload := map[string]interface{}{
			"page":     1,
			"pageSize": 500,
			"name":     params.Arguments.Name,
			"orderBy":  "created_at",
			"order":    "null",
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/containers/volume/search", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type ListVolumesInput struct {
	Name string `json:"name,omitempty" jsonschema:"filter by volume name"`
}

var CreateVolumeTool = mcp.NewServerTool[CreateVolumeInput, any](
	"create_volume",
	"Create a new Docker volume",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[CreateVolumeInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"name":   input.Name,
			"driver": input.Driver,
			"labels": input.Labels,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/containers/volume", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type CreateVolumeInput struct {
	Name   string   `json:"name" jsonschema:"volume name"`
	Driver string   `json:"driver,omitempty" jsonschema:"volume driver, default: local"`
	Labels []string `json:"labels,omitempty" jsonschema:"volume labels"`
}

var DeleteVolumeTool = mcp.NewServerTool[DeleteVolumeInput, any](
	"delete_volume",
	"[DANGEROUS] Delete Docker volumes",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[DeleteVolumeInput]) (*mcp.CallToolResultFor[any], error) {
		payload := map[string]interface{}{
			"names": params.Arguments.Names,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/containers/volume/del", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type DeleteVolumeInput struct {
	Names []string `json:"names" jsonschema:"volume names to delete"`
}

var VolumePruneTool = mcp.NewServerTool[VolumePruneInput, any](
	"volume_prune",
	"[DANGEROUS] Remove all unused Docker volumes",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[VolumePruneInput]) (*mcp.CallToolResultFor[any], error) {
		payload := map[string]interface{}{
			"pruneType":  "volume",
			"withTagAll": false,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/containers/prune", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type VolumePruneInput struct{}
