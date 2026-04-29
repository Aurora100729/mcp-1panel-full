package shell

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

var ShellExecTool = mcp.NewServerTool[ShellExecInput, any](
	"shell_exec",
	"[DANGEROUS] Execute a shell command on the local machine where MCP server runs. Returns stdout, stderr, and exit code. Supports timeout.",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[ShellExecInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments

		timeoutSec := input.Timeout
		if timeoutSec <= 0 {
			timeoutSec = 30
		}
		execCtx, cancel := context.WithTimeout(ctx, time.Duration(timeoutSec)*time.Second)
		defer cancel()

		shell, flag := getShell(input.Shell)
		cmd := exec.CommandContext(execCtx, shell, flag, input.Command)

		if input.Cwd != "" {
			cmd.Dir = input.Cwd
		}
		for _, e := range input.Env {
			cmd.Env = append(cmd.Env, e)
		}

		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		err := cmd.Run()
		exitCode := 0
		if err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				exitCode = exitErr.ExitCode()
			} else {
				return &mcp.CallToolResult{
					Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("exec error: %v", err)}},
					IsError: true,
				}, err
			}
		}

		output := map[string]interface{}{
			"exitCode": exitCode,
			"stdout":   truncate(stdout.String(), 100000),
			"stderr":   truncate(stderr.String(), 50000),
		}

		text := fmt.Sprintf("Exit: %d\n--- STDOUT ---\n%s", exitCode, truncate(stdout.String(), 100000))
		if stderr.Len() > 0 {
			text += fmt.Sprintf("\n--- STDERR ---\n%s", truncate(stderr.String(), 50000))
		}

		return &mcp.CallToolResult{
			Content:           []mcp.Content{&mcp.TextContent{Text: text}},
			StructuredContent: output,
		}, nil
	},
)

type ShellExecInput struct {
	Command string   `json:"command" jsonschema:"shell command to execute"`
	Cwd     string   `json:"cwd,omitempty" jsonschema:"working directory"`
	Env     []string `json:"env,omitempty" jsonschema:"environment variables KEY=VALUE"`
	Shell   string   `json:"shell,omitempty" jsonschema:"shell to use: bash, sh, powershell, cmd. Auto-detected if empty."`
	Timeout int      `json:"timeout,omitempty" jsonschema:"timeout in seconds, default 30"`
}

func getShell(preferred string) (string, string) {
	if preferred != "" {
		switch strings.ToLower(preferred) {
		case "bash":
			return "bash", "-c"
		case "sh":
			return "sh", "-c"
		case "powershell", "pwsh":
			return "powershell", "-Command"
		case "cmd":
			return "cmd", "/C"
		case "zsh":
			return "zsh", "-c"
		}
	}
	if runtime.GOOS == "windows" {
		return "cmd", "/C"
	}
	return "sh", "-c"
}

func truncate(s string, maxLen int) string {
	if len(s) > maxLen {
		return s[:maxLen] + "\n... [truncated]"
	}
	return s
}
