package backup

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/Aurora100729/mcp-1panel-full/utils"
)

var ListBackupAccountsTool = mcp.NewServerTool[ListBackupAccountsInput, any](
	"list_backup_accounts",
	"List configured backup storage accounts (local, S3, OSS, COS, MinIO, SFTP, WebDAV, OneDrive, etc.)",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ListBackupAccountsInput]) (*mcp.CallToolResultFor[any], error) {
		payload := map[string]interface{}{
			"page":     1,
			"pageSize": 500,
			"orderBy":  "created_at",
			"order":    "null",
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/settings/backup/search", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type ListBackupAccountsInput struct{}

var CreateBackupAccountTool = mcp.NewServerTool[CreateBackupAccountInput, any](
	"create_backup_account",
	"Create a new backup storage account",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[CreateBackupAccountInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"type":   input.Type,
			"vars":   input.Vars,
			"bucket": input.Bucket,
			"credential": input.Credential,
			"backupPath": input.BackupPath,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/settings/backup", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type CreateBackupAccountInput struct {
	Type       string      `json:"type" jsonschema:"backup type: LOCAL, S3, OSS, COS, KODO, MinIO, SFTP, WebDAV, OneDrive"`
	Vars       interface{} `json:"vars,omitempty" jsonschema:"type-specific configuration vars (JSON object)"`
	Bucket     string      `json:"bucket,omitempty" jsonschema:"bucket name for cloud storage"`
	Credential string      `json:"credential,omitempty" jsonschema:"credential string"`
	BackupPath string      `json:"backupPath,omitempty" jsonschema:"backup storage path"`
}

var ListBackupRecordsTool = mcp.NewServerTool[ListBackupRecordsInput, any](
	"list_backup_records",
	"List backup records/history",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ListBackupRecordsInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"page":     1,
			"pageSize": 100,
			"type":     input.Type,
			"name":     input.Name,
			"orderBy":  "created_at",
			"order":    "descending",
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/settings/backup/search/records", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type ListBackupRecordsInput struct {
	Type string `json:"type,omitempty" jsonschema:"backup type filter: website, database, app, directory"`
	Name string `json:"name,omitempty" jsonschema:"backup name filter"`
}

var BackupOperateTool = mcp.NewServerTool[BackupOperateInput, any](
	"backup_operate",
	"[DANGEROUS] Perform backup operations: create a new backup or restore from backup",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[BackupOperateInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"type":      input.Type,
			"name":      input.Name,
			"detailName": input.DetailName,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/settings/backup/backup", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type BackupOperateInput struct {
	Type       string `json:"type" jsonschema:"backup type: website, database, app, directory"`
	Name       string `json:"name" jsonschema:"name of what to backup"`
	DetailName string `json:"detailName,omitempty" jsonschema:"detail name (e.g. database name)"`
}
