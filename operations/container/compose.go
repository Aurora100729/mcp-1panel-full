package container

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/Aurora100729/mcp-1panel-full/utils"
)

var ListComposeTool = mcp.NewServerTool[ListComposeInput, any](
	"list_compose",
	"List Docker Compose projects managed by 1Panel",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ListComposeInput]) (*mcp.CallToolResultFor[any], error) {
		payload := map[string]interface{}{
			"page":     1,
			"pageSize": 500,
			"name":     params.Arguments.Name,
			"orderBy":  "name",
			"order":    "ascending",
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/containers/compose/search", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type ListComposeInput struct {
	Name string `json:"name,omitempty" jsonschema:"filter by compose project name"`
}

var ComposeUpTool = mcp.NewServerTool[ComposeUpInput, any](
	"compose_up",
	"[DANGEROUS] Create and start a Docker Compose project from YAML content or template",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ComposeUpInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"name":    input.Name,
			"from":    input.From,
			"file":    input.File,
			"path":    input.Path,
			"template": input.Template,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/containers/compose", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type ComposeUpInput struct {
	Name     string `json:"name" jsonschema:"compose project name"`
	From     string `json:"from" jsonschema:"source: edit, path, or template"`
	File     string `json:"file,omitempty" jsonschema:"docker-compose.yml content when from=edit"`
	Path     string `json:"path,omitempty" jsonschema:"path to docker-compose.yml when from=path"`
	Template int    `json:"template,omitempty" jsonschema:"template ID when from=template"`
}

var ComposeOperateTool = mcp.NewServerTool[ComposeOperateInput, any](
	"compose_operate",
	"[DANGEROUS] Operate on a Docker Compose project: up, down, start, stop, restart",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ComposeOperateInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"name":      input.Name,
			"path":      input.Path,
			"operation": input.Operation,
			"withFile":  input.WithFile,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/containers/compose/operate", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type ComposeOperateInput struct {
	Name      string `json:"name" jsonschema:"compose project name"`
	Path      string `json:"path,omitempty" jsonschema:"compose file path"`
	Operation string `json:"operation" jsonschema:"operation: up, down, start, stop, restart"`
	WithFile  bool   `json:"withFile,omitempty" jsonschema:"use file path"`
}

var ListComposeTemplatesTool = mcp.NewServerTool[ListComposeTemplatesInput, any](
	"list_compose_templates",
	"List Docker Compose templates",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ListComposeTemplatesInput]) (*mcp.CallToolResultFor[any], error) {
		payload := map[string]interface{}{
			"page":     1,
			"pageSize": 500,
			"name":     params.Arguments.Name,
			"orderBy":  "name",
			"order":    "ascending",
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/containers/template/search", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type ListComposeTemplatesInput struct {
	Name string `json:"name,omitempty" jsonschema:"filter by template name"`
}

var DockerInfoTool = mcp.NewServerTool[DockerInfoInput, any](
	"docker_info",
	"Get Docker daemon info: version, storage driver, containers count, images count, etc.",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[DockerInfoInput]) (*mcp.CallToolResultFor[any], error) {
		var result interface{}
		client := utils.NewPanelClient("GET", "/containers/daemonjson")
		return client.Request(&result)
	},
)

type DockerInfoInput struct{}

var DockerSystemPruneTool = mcp.NewServerTool[DockerSystemPruneInput, any](
	"docker_system_prune",
	"[DANGEROUS] Prune unused Docker resources: containers, images, networks, volumes, build cache",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[DockerSystemPruneInput]) (*mcp.CallToolResultFor[any], error) {
		payload := map[string]interface{}{
			"pruneType":  params.Arguments.PruneType,
			"withTagAll": params.Arguments.WithTagAll,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/containers/prune", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type DockerSystemPruneInput struct {
	PruneType  string `json:"pruneType" jsonschema:"type to prune: container, image, volume, network, buildcache"`
	WithTagAll bool   `json:"withTagAll,omitempty" jsonschema:"remove all, not just dangling"`
}

var ContainerExecTool = mcp.NewServerTool[ContainerExecInput, any](
	"container_exec",
	"[DANGEROUS] Execute a command inside a running Docker container",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ContainerExecInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		path := fmt.Sprintf("/containers/exec?containerID=%s&command=%s&user=%s",
			input.ContainerID, input.Command, input.User)
		var result interface{}
		client := utils.NewPanelClient("GET", path)
		return client.Request(&result)
	},
)

type ContainerExecInput struct {
	ContainerID string `json:"containerID" jsonschema:"container ID or name"`
	Command     string `json:"command" jsonschema:"command to execute, e.g. ls -la /"`
	User        string `json:"user,omitempty" jsonschema:"user to run as, default root"`
}
