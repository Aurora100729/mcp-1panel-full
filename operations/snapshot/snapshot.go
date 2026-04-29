package snapshot

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/Aurora100729/mcp-1panel-full/utils"
)

var ListSnapshotsTool = mcp.NewServerTool[ListSnapshotsInput, any](
	"list_snapshots",
	"List 1Panel system snapshots (full server state backups)",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ListSnapshotsInput]) (*mcp.CallToolResultFor[any], error) {
		payload := map[string]interface{}{
			"page":     1,
			"pageSize": 100,
			"info":     params.Arguments.Info,
			"orderBy":  "created_at",
			"order":    "descending",
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/settings/snapshot/search", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type ListSnapshotsInput struct {
	Info string `json:"info,omitempty" jsonschema:"filter snapshots by name"`
}

var CreateSnapshotTool = mcp.NewServerTool[CreateSnapshotInput, any](
	"create_snapshot",
	"[DANGEROUS] Create a full system snapshot of the 1Panel server",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[CreateSnapshotInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"from":        input.From,
			"description": input.Description,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/settings/snapshot", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type CreateSnapshotInput struct {
	From        string `json:"from" jsonschema:"backup account type to store snapshot"`
	Description string `json:"description,omitempty" jsonschema:"snapshot description"`
}

var RecoverSnapshotTool = mcp.NewServerTool[RecoverSnapshotInput, any](
	"recover_snapshot",
	"[DANGEROUS] Recover/restore the 1Panel server from a snapshot",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[RecoverSnapshotInput]) (*mcp.CallToolResultFor[any], error) {
		payload := map[string]interface{}{
			"id":    params.Arguments.ID,
			"isNew": params.Arguments.IsNew,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/settings/snapshot/recover", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type RecoverSnapshotInput struct {
	ID    int  `json:"id" jsonschema:"snapshot ID to recover from"`
	IsNew bool `json:"isNew,omitempty" jsonschema:"recover as new installation"`
}

var DeleteSnapshotTool = mcp.NewServerTool[DeleteSnapshotInput, any](
	"delete_snapshot",
	"[DANGEROUS] Delete system snapshots",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[DeleteSnapshotInput]) (*mcp.CallToolResultFor[any], error) {
		payload := map[string]interface{}{
			"ids": params.Arguments.IDs,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/settings/snapshot/del", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type DeleteSnapshotInput struct {
	IDs []int `json:"ids" jsonschema:"snapshot IDs to delete"`
}
