package localfs

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

var LocalFileReadTool = mcp.NewServerTool[LocalFileReadInput, any](
	"local_file_read",
	"Read a file from the local filesystem where MCP server runs",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[LocalFileReadInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		data, err := os.ReadFile(input.Path)
		if err != nil {
			return &mcp.CallToolResult{
				Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("read error: %v", err)}},
				IsError: true,
			}, err
		}
		content := string(data)
		if len(content) > 200000 {
			content = content[:200000] + "\n... [truncated]"
		}
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: content}},
		}, nil
	},
)

type LocalFileReadInput struct {
	Path string `json:"path" jsonschema:"absolute file path to read"`
}

var LocalFileWriteTool = mcp.NewServerTool[LocalFileWriteInput, any](
	"local_file_write",
	"[DANGEROUS] Write content to a file on the local filesystem",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[LocalFileWriteInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		mode := os.FileMode(0644)
		if input.Mode != 0 {
			mode = os.FileMode(input.Mode)
		}
		dir := filepath.Dir(input.Path)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return &mcp.CallToolResult{
				Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("mkdir error: %v", err)}},
				IsError: true,
			}, err
		}
		if input.Append {
			f, err := os.OpenFile(input.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, mode)
			if err != nil {
				return &mcp.CallToolResult{
					Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("open error: %v", err)}},
					IsError: true,
				}, err
			}
			defer f.Close()
			if _, err := f.WriteString(input.Content); err != nil {
				return &mcp.CallToolResult{
					Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("write error: %v", err)}},
					IsError: true,
				}, err
			}
		} else {
			if err := os.WriteFile(input.Path, []byte(input.Content), mode); err != nil {
				return &mcp.CallToolResult{
					Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("write error: %v", err)}},
					IsError: true,
				}, err
			}
		}
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("Successfully wrote to %s", input.Path)}},
		}, nil
	},
)

type LocalFileWriteInput struct {
	Path    string `json:"path" jsonschema:"absolute file path to write"`
	Content string `json:"content" jsonschema:"file content"`
	Append  bool   `json:"append,omitempty" jsonschema:"append to file instead of overwrite"`
	Mode    int    `json:"mode,omitempty" jsonschema:"file permission mode, e.g. 0644"`
}

var LocalFileDeleteTool = mcp.NewServerTool[LocalFileDeleteInput, any](
	"local_file_delete",
	"[DANGEROUS] Delete a file or directory from the local filesystem",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[LocalFileDeleteInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		var err error
		if input.Recursive {
			err = os.RemoveAll(input.Path)
		} else {
			err = os.Remove(input.Path)
		}
		if err != nil {
			return &mcp.CallToolResult{
				Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("delete error: %v", err)}},
				IsError: true,
			}, err
		}
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("Deleted: %s", input.Path)}},
		}, nil
	},
)

type LocalFileDeleteInput struct {
	Path      string `json:"path" jsonschema:"path to delete"`
	Recursive bool   `json:"recursive,omitempty" jsonschema:"recursively delete directory"`
}

var LocalFileListTool = mcp.NewServerTool[LocalFileListInput, any](
	"local_file_list",
	"List files and directories on the local filesystem",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[LocalFileListInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		entries, err := os.ReadDir(input.Path)
		if err != nil {
			return &mcp.CallToolResult{
				Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("readdir error: %v", err)}},
				IsError: true,
			}, err
		}
		var lines []string
		for _, e := range entries {
			if !input.ShowHidden && strings.HasPrefix(e.Name(), ".") {
				continue
			}
			info, _ := e.Info()
			size := int64(0)
			perm := fs.FileMode(0)
			if info != nil {
				size = info.Size()
				perm = info.Mode()
			}
			typeStr := "f"
			if e.IsDir() {
				typeStr = "d"
			}
			lines = append(lines, fmt.Sprintf("%s %s %10d %s", typeStr, perm, size, e.Name()))
		}
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: strings.Join(lines, "\n")}},
		}, nil
	},
)

type LocalFileListInput struct {
	Path       string `json:"path" jsonschema:"directory path to list"`
	ShowHidden bool   `json:"showHidden,omitempty" jsonschema:"show hidden files (dotfiles)"`
}

var LocalFileStatTool = mcp.NewServerTool[LocalFileStatInput, any](
	"local_file_stat",
	"Get file/directory info (size, permissions, modification time)",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[LocalFileStatInput]) (*mcp.CallToolResultFor[any], error) {
		info, err := os.Stat(params.Arguments.Path)
		if err != nil {
			return &mcp.CallToolResult{
				Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("stat error: %v", err)}},
				IsError: true,
			}, err
		}
		text := fmt.Sprintf("Name: %s\nSize: %d\nMode: %s\nIsDir: %v\nModTime: %s",
			info.Name(), info.Size(), info.Mode(), info.IsDir(), info.ModTime().Format("2006-01-02 15:04:05"))
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: text}},
		}, nil
	},
)

type LocalFileStatInput struct {
	Path string `json:"path" jsonschema:"file or directory path"`
}

var LocalFileSearchTool = mcp.NewServerTool[LocalFileSearchInput, any](
	"local_file_search",
	"Search for files by name pattern in a directory tree",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[LocalFileSearchInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		maxDepth := input.MaxDepth
		if maxDepth <= 0 {
			maxDepth = 5
		}
		var matches []string
		count := 0
		_ = filepath.WalkDir(input.Path, func(path string, d fs.DirEntry, err error) error {
			if err != nil || count >= 500 {
				return filepath.SkipDir
			}
			rel, _ := filepath.Rel(input.Path, path)
			depth := strings.Count(rel, string(os.PathSeparator))
			if depth > maxDepth {
				return filepath.SkipDir
			}
			matched, _ := filepath.Match(input.Pattern, d.Name())
			if matched {
				matches = append(matches, path)
				count++
			}
			return nil
		})
		if len(matches) == 0 {
			return &mcp.CallToolResult{
				Content: []mcp.Content{&mcp.TextContent{Text: "No files matched"}},
			}, nil
		}
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: strings.Join(matches, "\n")}},
		}, nil
	},
)

type LocalFileSearchInput struct {
	Path     string `json:"path" jsonschema:"root directory to search"`
	Pattern  string `json:"pattern" jsonschema:"filename glob pattern, e.g. *.go or config.*"`
	MaxDepth int    `json:"maxDepth,omitempty" jsonschema:"max directory depth, default 5"`
}
