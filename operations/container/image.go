package container

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/Aurora100729/mcp-1panel-full/utils"
)

var ListImagesTool = mcp.NewServerTool[ListImagesInput, any](
	"list_images",
	"List all Docker images with size, tags, and creation time",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ListImagesInput]) (*mcp.CallToolResultFor[any], error) {
		payload := map[string]interface{}{
			"page":     1,
			"pageSize": 500,
			"name":     params.Arguments.Name,
			"orderBy":  "created_at",
			"order":    "null",
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/containers/image/search", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type ListImagesInput struct {
	Name string `json:"name,omitempty" jsonschema:"filter images by name"`
}

var ImagePullTool = mcp.NewServerTool[ImagePullInput, any](
	"image_pull",
	"Pull a Docker image from a registry",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ImagePullInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"repoID":  input.RepoID,
			"imageName": input.ImageName,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/containers/image/pull", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type ImagePullInput struct {
	ImageName string `json:"imageName" jsonschema:"image to pull, e.g. nginx:latest"`
	RepoID    int    `json:"repoID,omitempty" jsonschema:"registry repo ID, 0 for Docker Hub"`
}

var ImageRemoveTool = mcp.NewServerTool[ImageRemoveInput, any](
	"image_remove",
	"[DANGEROUS] Remove one or more Docker images",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ImageRemoveInput]) (*mcp.CallToolResultFor[any], error) {
		payload := map[string]interface{}{
			"imageID":  params.Arguments.ImageID,
			"isForce":  params.Arguments.Force,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/containers/image/remove", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type ImageRemoveInput struct {
	ImageID string `json:"imageID" jsonschema:"image ID or name:tag"`
	Force   bool   `json:"isForce,omitempty" jsonschema:"force remove even if in use"`
}

var ImageBuildTool = mcp.NewServerTool[ImageBuildInput, any](
	"image_build",
	"[DANGEROUS] Build a Docker image from a Dockerfile",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ImageBuildInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"name":       input.Name,
			"from":       input.From,
			"dockerfile": input.Dockerfile,
			"tags":       input.Tags,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/containers/image/build", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type ImageBuildInput struct {
	Name       string   `json:"name" jsonschema:"image name"`
	From       string   `json:"from" jsonschema:"build from: path or edit"`
	Dockerfile string   `json:"dockerfile,omitempty" jsonschema:"Dockerfile content when from=edit"`
	Tags       []string `json:"tags,omitempty" jsonschema:"image tags"`
}

var ImagePruneTool = mcp.NewServerTool[ImagePruneInput, any](
	"image_prune",
	"[DANGEROUS] Remove unused Docker images to free disk space",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ImagePruneInput]) (*mcp.CallToolResultFor[any], error) {
		payload := map[string]interface{}{
			"pruneType":  "image",
			"withTagAll": params.Arguments.WithTagAll,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/containers/prune", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type ImagePruneInput struct {
	WithTagAll bool `json:"withTagAll,omitempty" jsonschema:"remove all unused images, not just dangling"`
}
