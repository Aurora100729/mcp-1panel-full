package database

import (
	"context"
	"errors"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/Aurora100729/mcp-1panel-full/operations/types"
	"github.com/Aurora100729/mcp-1panel-full/utils"
)

const (
	ListDatabases = "list_databases"
)

var ListDatabasesTool = mcp.NewServerTool[ListDatabasesInput, any](
	ListDatabases,
	"list databases by name",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ListDatabasesInput]) (*mcp.CallToolResultFor[any], error) {
		database := params.Arguments.Name
		if database == "" {
			err := errors.New("database name is required")
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{Text: err.Error()},
				},
				IsError: true,
			}, err
		}
		pageReq := &types.ListDatabaseRequest{
			PageRequest: types.PageRequest{
				Page:     1,
				PageSize: 500,
			},
			Order:    "null",
			OrderBy:  "created_at",
			Database: database,
		}
		databaseListRes := &types.DatabaseListResponse{}
		result, err := utils.NewPanelClient("POST", "/databases/search", utils.WithPayload(pageReq)).Request(databaseListRes)
		if result != nil {
			result.StructuredContent = databaseListRes
		}
		return result, err
	},
)

type ListDatabasesInput struct {
	Name string `json:"name" jsonschema:"database name"`
}
