package localssh

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"golang.org/x/crypto/ssh"
)

var SSHExecTool = mcp.NewServerTool[SSHExecInput, any](
	"ssh_remote_exec",
	"[DANGEROUS] Execute a command on a remote host via SSH. Supports password and key authentication.",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[SSHExecInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		port := input.Port
		if port == 0 {
			port = 22
		}
		timeoutSec := input.Timeout
		if timeoutSec <= 0 {
			timeoutSec = 30
		}

		config := &ssh.ClientConfig{
			User:            input.User,
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Timeout:         time.Duration(timeoutSec) * time.Second,
		}

		if input.Password != "" {
			config.Auth = append(config.Auth, ssh.Password(input.Password))
		}
		if input.KeyPath != "" {
			key, err := os.ReadFile(input.KeyPath)
			if err != nil {
				return &mcp.CallToolResult{
					Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("read key error: %v", err)}},
					IsError: true,
				}, err
			}
			var signer ssh.Signer
			if input.KeyPassphrase != "" {
				signer, err = ssh.ParsePrivateKeyWithPassphrase(key, []byte(input.KeyPassphrase))
			} else {
				signer, err = ssh.ParsePrivateKey(key)
			}
			if err != nil {
				return &mcp.CallToolResult{
					Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("parse key error: %v", err)}},
					IsError: true,
				}, err
			}
			config.Auth = append(config.Auth, ssh.PublicKeys(signer))
		}
		if input.KeyContent != "" {
			var signer ssh.Signer
			var err error
			if input.KeyPassphrase != "" {
				signer, err = ssh.ParsePrivateKeyWithPassphrase([]byte(input.KeyContent), []byte(input.KeyPassphrase))
			} else {
				signer, err = ssh.ParsePrivateKey([]byte(input.KeyContent))
			}
			if err != nil {
				return &mcp.CallToolResult{
					Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("parse key error: %v", err)}},
					IsError: true,
				}, err
			}
			config.Auth = append(config.Auth, ssh.PublicKeys(signer))
		}

		addr := fmt.Sprintf("%s:%d", input.Host, port)
		client, err := ssh.Dial("tcp", addr, config)
		if err != nil {
			return &mcp.CallToolResult{
				Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("SSH connect error: %v", err)}},
				IsError: true,
			}, err
		}
		defer client.Close()

		session, err := client.NewSession()
		if err != nil {
			return &mcp.CallToolResult{
				Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("SSH session error: %v", err)}},
				IsError: true,
			}, err
		}
		defer session.Close()

		var stdout, stderr bytes.Buffer
		session.Stdout = &stdout
		session.Stderr = &stderr

		err = session.Run(input.Command)
		exitCode := 0
		if err != nil {
			if exitErr, ok := err.(*ssh.ExitError); ok {
				exitCode = exitErr.ExitStatus()
			} else {
				return &mcp.CallToolResult{
					Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("exec error: %v", err)}},
					IsError: true,
				}, err
			}
		}

		text := fmt.Sprintf("[%s] Exit: %d\n--- STDOUT ---\n%s", addr, exitCode, stdout.String())
		if stderr.Len() > 0 {
			text += fmt.Sprintf("\n--- STDERR ---\n%s", stderr.String())
		}

		result := map[string]interface{}{
			"host":     input.Host,
			"exitCode": exitCode,
			"stdout":   stdout.String(),
			"stderr":   stderr.String(),
		}

		return &mcp.CallToolResult{
			Content:           []mcp.Content{&mcp.TextContent{Text: text}},
			StructuredContent: result,
		}, nil
	},
)

type SSHExecInput struct {
	Host          string `json:"host" jsonschema:"remote host address"`
	Port          int    `json:"port,omitempty" jsonschema:"SSH port, default 22"`
	User          string `json:"user" jsonschema:"SSH username"`
	Password      string `json:"password,omitempty" jsonschema:"SSH password (if using password auth)"`
	KeyPath       string `json:"keyPath,omitempty" jsonschema:"path to SSH private key file"`
	KeyContent    string `json:"keyContent,omitempty" jsonschema:"SSH private key content (PEM)"`
	KeyPassphrase string `json:"keyPassphrase,omitempty" jsonschema:"passphrase for encrypted key"`
	Command       string `json:"command" jsonschema:"command to execute on remote host"`
	Timeout       int    `json:"timeout,omitempty" jsonschema:"connection timeout in seconds, default 30"`
}

var SSHPortCheckTool = mcp.NewServerTool[SSHPortCheckInput, any](
	"ssh_port_check",
	"Check if an SSH port is open/reachable on a remote host",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[SSHPortCheckInput]) (*mcp.CallToolResultFor[any], error) {
		input := params.Arguments
		port := input.Port
		if port == 0 {
			port = 22
		}
		addr := fmt.Sprintf("%s:%d", input.Host, port)
		conn, err := net.DialTimeout("tcp", addr, 5*time.Second)
		if err != nil {
			return &mcp.CallToolResult{
				Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("Port %d on %s is CLOSED or unreachable: %v", port, input.Host, err)}},
			}, nil
		}
		conn.Close()
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("Port %d on %s is OPEN", port, input.Host)}},
		}, nil
	},
)

type SSHPortCheckInput struct {
	Host string `json:"host" jsonschema:"remote host address"`
	Port int    `json:"port,omitempty" jsonschema:"port to check, default 22"`
}
