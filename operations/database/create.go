package database

import (
	"context"
	"encoding/base64"
	"errors"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/Aurora100729/mcp-1panel-full/operations/types"
	"github.com/Aurora100729/mcp-1panel-full/utils"
)

const (
	CreateDatabase = "create_database"
)

var CreateDatabaseTool = mcp.NewServerTool[CreateDatabaseInput, any](
	CreateDatabase,
	"create a database by type name and password",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[CreateDatabaseInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		if input.Database == "" {
			err := errors.New("database name is required")
			return &mcp.CallToolResult{
				Content: []mcp.Content{&mcp.TextContent{Text: err.Error()}},
				IsError: true,
			}, err
		}
		if input.DatabaseType == "" {
			err := errors.New("database type is required")
			return &mcp.CallToolResult{
				Content: []mcp.Content{&mcp.TextContent{Text: err.Error()}},
				IsError: true,
			}, err
		}
		if input.DatabaseType != "mysql" && input.DatabaseType != "postgresql" {
			err := errors.New("database type is invalid, support mysql and postgresql")
			return &mcp.CallToolResult{
				Content: []mcp.Content{&mcp.TextContent{Text: err.Error()}},
				IsError: true,
			}, err
		}
		if input.Name == "" {
			err := errors.New("name is required")
			return &mcp.CallToolResult{
				Content: []mcp.Content{&mcp.TextContent{Text: err.Error()}},
				IsError: true,
			}, err
		}

		password := input.Password
		if password == "" {
			password = utils.GetRandomStr(12)
		}
		encodedPassword := base64.StdEncoding.EncodeToString([]byte(password))

		username := input.Username
		if username == "" {
			username = input.Name
		}

		createReq := &types.CreateDatabaseRequest{
			Database: input.Database,
			Password: encodedPassword,
			Type:     input.DatabaseType,
			Name:     input.Name,
			From:     "local",
			Username: username,
		}
		var createURL string
		if input.DatabaseType == "mysql" {
			createURL = "/databases"
			createReq.Format = "utf8mb4"
			createReq.Permission = "%"
		} else {
			createURL = "/databases/pg"
			createReq.Format = "UTF8"
		}
		res := &types.Response{}
		result, err := utils.NewPanelClient("POST", createURL, utils.WithPayload(createReq)).Request(res)
		if result != nil {
			result.StructuredContent = res
		}
		return result, err
	},
)

type CreateDatabaseInput struct {
	DatabaseType string `json:"database_type" jsonschema:"installed database app type, support mysql and postgresql"`
	Database     string `json:"database" jsonschema:"installed database app name"`
	Name         string `json:"name" jsonschema:"database name"`
	Username     string `json:"username,omitempty" jsonschema:"database username"`
	Password     string `json:"password,omitempty" jsonschema:"database password"`
}
