package file

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/Aurora100729/mcp-1panel-full/utils"
)

var ListFilesTool = mcp.NewServerTool[ListFilesInput, any](
	"panel_file_list",
	"List files and directories at a given path on the 1Panel server",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ListFilesInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"path":       input.Path,
			"expand":     true,
			"showHidden": input.ShowHidden,
			"page":       1,
			"pageSize":   500,
			"search":     input.Search,
			"containSub": input.ContainSub,
			"sortBy":     "name",
			"sortOrder":  "ascending",
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/files/search", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type ListFilesInput struct {
	Path       string `json:"path" jsonschema:"directory path to list, e.g. /root or /opt"`
	ShowHidden bool   `json:"showHidden,omitempty" jsonschema:"show hidden files (dotfiles)"`
	Search     string `json:"search,omitempty" jsonschema:"filter files by name"`
	ContainSub bool   `json:"containSub,omitempty" jsonschema:"include subdirectories recursively"`
}

var ReadFileTool = mcp.NewServerTool[ReadFileInput, any](
	"panel_file_read",
	"Read the content of a file on the 1Panel server (text files only)",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ReadFileInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"path": input.Path,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/files/content", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type ReadFileInput struct {
	Path string `json:"path" jsonschema:"full path to the file to read"`
}

var WriteFileTool = mcp.NewServerTool[WriteFileInput, any](
	"panel_file_write",
	"[DANGEROUS] Write/overwrite content to a file on the 1Panel server",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[WriteFileInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"path":    input.Path,
			"content": input.Content,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/files/save", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type WriteFileInput struct {
	Path    string `json:"path" jsonschema:"full path to the file"`
	Content string `json:"content" jsonschema:"file content to write"`
}

var CreateFileTool = mcp.NewServerTool[CreateFileInput, any](
	"panel_file_create",
	"Create a new file or directory on the 1Panel server",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[CreateFileInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"path":  input.Path,
			"isDir": input.IsDir,
			"mode":  input.Mode,
			"isLink": input.IsLink,
			"linkPath": input.LinkPath,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/files", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type CreateFileInput struct {
	Path     string `json:"path" jsonschema:"full path for the new file or directory"`
	IsDir    bool   `json:"isDir,omitempty" jsonschema:"true to create directory, false for file"`
	Mode     int    `json:"mode,omitempty" jsonschema:"file permission mode, e.g. 755"`
	IsLink   bool   `json:"isLink,omitempty" jsonschema:"create symbolic link"`
	LinkPath string `json:"linkPath,omitempty" jsonschema:"target path for symbolic link"`
}

var DeleteFileTool = mcp.NewServerTool[DeleteFileInput, any](
	"panel_file_delete",
	"[DANGEROUS] Delete files or directories on the 1Panel server",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[DeleteFileInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"path":    input.Path,
			"isDir":   input.IsDir,
			"forceDelete": input.ForceDelete,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/files/del", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type DeleteFileInput struct {
	Path        string `json:"path" jsonschema:"full path to delete"`
	IsDir       bool   `json:"isDir,omitempty" jsonschema:"true if deleting a directory"`
	ForceDelete bool   `json:"forceDelete,omitempty" jsonschema:"force delete without confirmation"`
}

var MoveFileTool = mcp.NewServerTool[MoveFileInput, any](
	"panel_file_move",
	"Move or rename files/directories on the 1Panel server",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[MoveFileInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		mtype := input.Type
		if mtype == "" {
			mtype = "move"
		}
		payload := map[string]interface{}{
			"oldPaths": input.OldPaths,
			"newPath":  input.NewPath,
			"type":     mtype,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/files/move", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type MoveFileInput struct {
	OldPaths []string `json:"oldPaths" jsonschema:"source paths to move"`
	NewPath  string   `json:"newPath" jsonschema:"destination directory"`
	Type     string   `json:"type,omitempty" jsonschema:"operation: move or copy"`
}

var CompressFileTool = mcp.NewServerTool[CompressFileInput, any](
	"panel_file_compress",
	"Compress files into an archive (zip/tar.gz/gz/bz2/xz/tar) on the 1Panel server",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[CompressFileInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		ctype := input.Type
		if ctype == "" {
			ctype = "zip"
		}
		payload := map[string]interface{}{
			"type":    ctype,
			"name":    input.Name,
			"dst":     input.Dst,
			"files":   input.Files,
			"replace": input.Replace,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/files/compress", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type CompressFileInput struct {
	Type    string   `json:"type,omitempty" jsonschema:"archive format: zip, tar.gz, gz, bz2, xz, tar"`
	Name    string   `json:"name" jsonschema:"archive filename"`
	Dst     string   `json:"dst" jsonschema:"destination directory for archive"`
	Files   []string `json:"files" jsonschema:"paths of files to compress"`
	Replace bool     `json:"replace,omitempty" jsonschema:"replace if archive exists"`
}

var DecompressFileTool = mcp.NewServerTool[DecompressFileInput, any](
	"panel_file_decompress",
	"Decompress/extract an archive file on the 1Panel server",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[DecompressFileInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"type": input.Type,
			"dst":  input.Dst,
			"path": input.Path,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/files/decompress", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type DecompressFileInput struct {
	Type string `json:"type" jsonschema:"archive type: zip, tar.gz, gz, bz2, xz, tar"`
	Path string `json:"path" jsonschema:"path to archive file"`
	Dst  string `json:"dst" jsonschema:"destination directory for extraction"`
}

var ChmodFileTool = mcp.NewServerTool[ChmodFileInput, any](
	"panel_file_chmod",
	"Change file/directory permissions on the 1Panel server",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ChmodFileInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"paths": input.Paths,
			"mode":  input.Mode,
			"sub":   input.Sub,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/files/mode", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type ChmodFileInput struct {
	Paths []string `json:"paths" jsonschema:"file/directory paths"`
	Mode  string   `json:"mode" jsonschema:"permission mode, e.g. 0755"`
	Sub   bool     `json:"sub,omitempty" jsonschema:"apply recursively to subdirectories"`
}

var WgetFileTool = mcp.NewServerTool[WgetFileInput, any](
	"panel_file_wget",
	"Download a file from a URL to the 1Panel server",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[WgetFileInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"url":  input.URL,
			"path": input.Path,
			"name": input.Name,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/files/wget", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type WgetFileInput struct {
	URL  string `json:"url" jsonschema:"URL to download from"`
	Path string `json:"path" jsonschema:"directory to save the file"`
	Name string `json:"name" jsonschema:"filename to save as"`
}

var FileSearchContentTool = mcp.NewServerTool[FileSearchContentInput, any](
	"panel_file_search_content",
	"Search for text content within files on the 1Panel server (grep-like)",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[FileSearchContentInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		payload := map[string]interface{}{
			"path":   input.Path,
			"search": input.Search,
		}
		var result interface{}
		client := utils.NewPanelClient("POST", "/files/content", utils.WithPayload(payload))
		return client.Request(&result)
	},
)

type FileSearchContentInput struct {
	Path   string `json:"path" jsonschema:"directory path to search in"`
	Search string `json:"search" jsonschema:"text pattern to search for"`
}
