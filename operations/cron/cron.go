package cron

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/Aurora100729/mcp-1panel-full/utils"
)

var ListCronsTool = mcp.NewServerTool[ListCronsInput, any](
	"list_crons",
	"List scheduled cron jobs on the 1Panel server",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ListCronsInput]) (*mcp.CallToolResultFor[any], error) {
		payload := map[string]interface{}{
			"page":     1,
			"pageSize": 500,
			"name":     params.Arguments.Name,
			"type":     params.Arguments.Type,
			"orderBy":  "created_at",
			"order":    "null",
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/cronjobs/search", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type ListCronsInput struct {
	Name string `json:"name,omitempty" jsonschema:"filter by cron name"`
	Type string `json:"type,omitempty" jsonschema:"filter by type: shell, website, database, directory, curl, log, cutWebsiteLog, clean, app"`
}

var CreateCronTool = mcp.NewServerTool[CreateCronInput, any](
	"create_cron",
	"Create a new cron job (scheduled task)",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[CreateCronInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"name":        input.Name,
			"type":        input.Type,
			"specType":    input.SpecType,
			"week":        input.Week,
			"day":         input.Day,
			"hour":        input.Hour,
			"minute":      input.Minute,
			"second":      input.Second,
			"script":      input.Script,
			"url":         input.URL,
			"website":     input.Website,
			"dbName":      input.DBName,
			"dbType":      input.DBType,
			"sourceDir":   input.SourceDir,
			"exclusionRules": input.ExclusionRules,
			"retainCopies": input.RetainCopies,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/cronjobs", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type CreateCronInput struct {
	Name           string `json:"name" jsonschema:"cron job name"`
	Type           string `json:"type" jsonschema:"type: shell, website, database, directory, curl, log, cutWebsiteLog, clean, app"`
	SpecType       string `json:"specType" jsonschema:"schedule type: perWeek, perNDay, perDay, perHour, perNMinute, perNSecond"`
	Week           int    `json:"week,omitempty" jsonschema:"day of week 0-6 for perWeek"`
	Day            int    `json:"day,omitempty" jsonschema:"day interval for perNDay"`
	Hour           int    `json:"hour,omitempty" jsonschema:"hour 0-23"`
	Minute         int    `json:"minute,omitempty" jsonschema:"minute 0-59"`
	Second         int    `json:"second,omitempty" jsonschema:"second 0-59"`
	Script         string `json:"script,omitempty" jsonschema:"shell script content for type=shell"`
	URL            string `json:"url,omitempty" jsonschema:"URL for type=curl"`
	Website        string `json:"website,omitempty" jsonschema:"website name for type=website"`
	DBName         string `json:"dbName,omitempty" jsonschema:"database name for type=database"`
	DBType         string `json:"dbType,omitempty" jsonschema:"database type for type=database"`
	SourceDir      string `json:"sourceDir,omitempty" jsonschema:"source directory for type=directory"`
	ExclusionRules string `json:"exclusionRules,omitempty" jsonschema:"exclusion rules"`
	RetainCopies   int    `json:"retainCopies,omitempty" jsonschema:"number of backup copies to keep"`
}

var DeleteCronTool = mcp.NewServerTool[DeleteCronInput, any](
	"delete_cron",
	"[DANGEROUS] Delete a cron job",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[DeleteCronInput]) (*mcp.CallToolResultFor[any], error) {
		payload := map[string]interface{}{
			"ids":       params.Arguments.IDs,
			"cleanData": params.Arguments.CleanData,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/cronjobs/del", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type DeleteCronInput struct {
	IDs       []int `json:"ids" jsonschema:"cron job IDs to delete"`
	CleanData bool  `json:"cleanData,omitempty" jsonschema:"also clean backup data"`
}

var HandleCronTool = mcp.NewServerTool[HandleCronInput, any](
	"handle_cron",
	"Manually trigger/run a cron job, or enable/disable it",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[HandleCronInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		path := fmt.Sprintf("/cronjobs/handle?id=%d", input.ID)
		var result interface{}
		client := utils.NewPanelClient("POST", path)
		return client.Request(&result)
	},
)

type HandleCronInput struct {
	ID int `json:"id" jsonschema:"cron job ID to handle"`
}

var UpdateCronStatusTool = mcp.NewServerTool[UpdateCronStatusInput, any](
	"update_cron_status",
	"Enable or disable a cron job",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[UpdateCronStatusInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"id":     input.ID,
			"status": input.Status,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/cronjobs/status", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type UpdateCronStatusInput struct {
	ID     int    `json:"id" jsonschema:"cron job ID"`
	Status string `json:"status" jsonschema:"status: enable or disable"`
}
