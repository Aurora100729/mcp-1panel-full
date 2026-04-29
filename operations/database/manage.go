package database

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/Aurora100729/mcp-1panel-full/utils"
)

var DeleteDatabaseTool = mcp.NewServerTool[DeleteDatabaseInput, any](
	"delete_database",
	"[DANGEROUS] Delete a database",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[DeleteDatabaseInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		var path string
		if input.Type == "postgresql" {
			path = "/databases/pg/del"
		} else {
			path = "/databases/del"
		}
		payload := map[string]interface{}{
			"ids":         input.IDs,
			"deleteBackup": input.DeleteBackup,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", path, utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type DeleteDatabaseInput struct {
	IDs          []int  `json:"ids" jsonschema:"database IDs to delete"`
	Type         string `json:"type,omitempty" jsonschema:"database type: mysql or postgresql"`
	DeleteBackup bool   `json:"deleteBackup,omitempty" jsonschema:"also delete backups"`
}

var DatabaseBackupTool = mcp.NewServerTool[DatabaseBackupInput, any](
	"database_backup",
	"Create a backup of a database",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[DatabaseBackupInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"type":       input.Type,
			"name":       input.Name,
			"detailName": input.DetailName,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/settings/backup/backup", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type DatabaseBackupInput struct {
	Type       string `json:"type" jsonschema:"backup type: mysql or postgresql"`
	Name       string `json:"name" jsonschema:"installed database app name"`
	DetailName string `json:"detailName" jsonschema:"database name to backup"`
}

var DatabaseStatusTool = mcp.NewServerTool[DatabaseStatusInput, any](
	"database_status",
	"Get MySQL/PostgreSQL server status and variables",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[DatabaseStatusInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		var path string
		if input.Type == "postgresql" {
			path = fmt.Sprintf("/databases/pg/status/%s", input.Database)
		} else {
			path = fmt.Sprintf("/databases/status/%s", input.Database)
		}
		var result interface{}
		client := utils.NewPanelClient("GET", path)
		return client.Request(&result)
	},
)

type DatabaseStatusInput struct {
	Type     string `json:"type" jsonschema:"database type: mysql or postgresql"`
	Database string `json:"database" jsonschema:"installed database app name"`
}

var ListRedisTool = mcp.NewServerTool[ListRedisInput, any](
	"list_redis",
	"List Redis databases/keys",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ListRedisInput]) (*mcp.CallToolResultFor[any], error) {
		payload := map[string]interface{}{
			"page":     1,
			"pageSize": 500,
			"database": params.Arguments.Database,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/databases/redis/search", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type ListRedisInput struct {
	Database string `json:"database" jsonschema:"installed Redis app name"`
}

var RedisStatusTool = mcp.NewServerTool[RedisStatusInput, any](
	"redis_status",
	"Get Redis server status information",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[RedisStatusInput]) (*mcp.CallToolResultFor[any], error) {
		path := fmt.Sprintf("/databases/redis/status/%s", params.Arguments.Database)
		var result interface{}
		client := utils.NewPanelClient("GET", path)
		return client.Request(&result)
	},
)

type RedisStatusInput struct {
	Database string `json:"database" jsonschema:"installed Redis app name"`
}
