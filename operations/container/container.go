package container

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/Aurora100729/mcp-1panel-full/operations/types"
	"github.com/Aurora100729/mcp-1panel-full/utils"
)

var ListContainersTool = mcp.NewServerTool[ListContainersInput, any](
	"list_containers",
	"List all Docker containers with status, image, ports, and resource usage",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ListContainersInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		req := &types.PageRequest{
			Page:     1,
			PageSize: 500,
			Name:     input.Name,
		}
		payload := map[string]interface{}{
			"page":     req.Page,
			"pageSize": req.PageSize,
			"name":     input.Name,
			"state":    input.State,
			"orderBy":  "created_at",
			"order":    "null",
			"filters":  "",
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/containers/search", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type ListContainersInput struct {
	Name  string `json:"name,omitempty" jsonschema:"filter containers by name"`
	State string `json:"state,omitempty" jsonschema:"filter by state: all, running, stopped, created, exited"`
}

var ContainerOperateTool = mcp.NewServerTool[ContainerOperateInput, any](
	"container_operate",
	"[DANGEROUS] Operate on Docker containers: start, stop, restart, kill, pause, unpause, remove, rename",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ContainerOperateInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"names":     input.Names,
			"operation": input.Operation,
		}
		if input.NewName != "" {
			payload["newName"] = input.NewName
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/containers/operate", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type ContainerOperateInput struct {
	Names     []string `json:"names" jsonschema:"container names or IDs to operate on"`
	Operation string   `json:"operation" jsonschema:"operation: start, stop, restart, kill, pause, unpause, remove, rename"`
	NewName   string   `json:"newName,omitempty" jsonschema:"new name when operation is rename"`
}

var ContainerInspectTool = mcp.NewServerTool[ContainerInspectInput, any](
	"container_inspect",
	"Get detailed inspect information of a Docker container",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ContainerInspectInput]) (*mcp.CallToolResultFor[any], error) {
		var result interface{}
		path := fmt.Sprintf("/containers/inspect?id=%s", params.Arguments.ID)
		client := utils.NewPanelClient("GET", path)
		return client.Request(&result)
	},
)

type ContainerInspectInput struct {
	ID string `json:"id" jsonschema:"container ID or name"`
}

var ContainerLogsTool = mcp.NewServerTool[ContainerLogsInput, any](
	"container_logs",
	"Get logs of a Docker container",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ContainerLogsInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		since := input.Since
		if since == "" {
			since = "all"
		}
		tail := input.Tail
		if tail == "" {
			tail = "200"
		}
		payload := map[string]interface{}{
			"containerID": input.ContainerID,
			"mode":        "all",
			"since":       since,
			"tail":        tail,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/containers/search/log", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type ContainerLogsInput struct {
	ContainerID string `json:"containerID" jsonschema:"container ID or name"`
	Since       string `json:"since,omitempty" jsonschema:"show logs since timestamp or 'all'"`
	Tail        string `json:"tail,omitempty" jsonschema:"number of lines from end, default 200"`
}

var ContainerStatsTool = mcp.NewServerTool[ContainerStatsInput, any](
	"container_stats",
	"Get real-time resource usage stats of a Docker container (CPU, memory, network, IO)",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ContainerStatsInput]) (*mcp.CallToolResultFor[any], error) {
		var result interface{}
		path := fmt.Sprintf("/containers/stats/%s", params.Arguments.ID)
		client := utils.NewPanelClient("GET", path)
		return client.Request(&result)
	},
)

type ContainerStatsInput struct {
	ID string `json:"id" jsonschema:"container ID or name"`
}

var ContainerCreateTool = mcp.NewServerTool[ContainerCreateInput, any](
	"container_create",
	"[DANGEROUS] Create and start a new Docker container with full configuration options",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ContainerCreateInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"name":          input.Name,
			"image":         input.Image,
			"cmdStr":        input.Cmd,
			"publishAllPorts": input.PublishAllPorts,
			"exposedPorts":  input.ExposedPorts,
			"nanoCPUs":      input.NanoCPUs,
			"memory":        input.Memory,
			"autoRemove":    input.AutoRemove,
			"labels":        input.Labels,
			"labelsStr":     input.LabelsStr,
			"env":           input.Env,
			"restartPolicy": input.RestartPolicy,
			"volumes":       input.Volumes,
			"network":       input.Network,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/containers", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type ContainerCreateInput struct {
	Name            string            `json:"name" jsonschema:"container name"`
	Image           string            `json:"image" jsonschema:"docker image, e.g. nginx:latest"`
	Cmd             string            `json:"cmdStr,omitempty" jsonschema:"command to run"`
	PublishAllPorts bool              `json:"publishAllPorts,omitempty" jsonschema:"publish all exposed ports"`
	ExposedPorts    []interface{}     `json:"exposedPorts,omitempty" jsonschema:"port mappings array"`
	NanoCPUs        float64           `json:"nanoCPUs,omitempty" jsonschema:"CPU limit in cores"`
	Memory          float64           `json:"memory,omitempty" jsonschema:"memory limit in bytes"`
	AutoRemove      bool              `json:"autoRemove,omitempty" jsonschema:"auto remove after stop"`
	Labels          []string          `json:"labels,omitempty" jsonschema:"container labels"`
	LabelsStr       []string          `json:"labelsStr,omitempty" jsonschema:"label strings key=value"`
	Env             []string          `json:"env,omitempty" jsonschema:"environment variables KEY=VALUE"`
	RestartPolicy   string            `json:"restartPolicy,omitempty" jsonschema:"restart policy: no, always, on-failure, unless-stopped"`
	Volumes         []interface{}     `json:"volumes,omitempty" jsonschema:"volume mount configurations"`
	Network         string            `json:"network,omitempty" jsonschema:"docker network name"`
}

var ContainerPruneTool = mcp.NewServerTool[ContainerPruneInput, any](
	"container_prune",
	"[DANGEROUS] Remove all stopped containers to free up resources",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ContainerPruneInput]) (*mcp.CallToolResultFor[any], error) {
		payload := map[string]interface{}{
			"pruneType": "container",
			"withTagAll": false,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/containers/prune", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type ContainerPruneInput struct{}
