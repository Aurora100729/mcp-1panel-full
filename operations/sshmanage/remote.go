package sshmanage

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"golang.org/x/crypto/ssh"
)

// Default connection settings populated from CLI flags.
var (
	DefaultHost     string
	DefaultUser     string
	DefaultPort     int
	DefaultKeyPath  string
	DefaultPassword string
)

// SetDefaults configures default SSH connection parameters used when tool arguments are not provided.
func SetDefaults(host, user, keyPath, password string, port int) {
	DefaultHost = host
	DefaultUser = user
	DefaultKeyPath = keyPath
	DefaultPassword = password
	DefaultPort = port
}

// SSHRemoteExecTool executes a command on a remote host via SSH using password or key auth.
var SSHRemoteExecTool = mcp.NewServerTool[SSHRemoteExecInput, any](
	"ssh_remote_exec",
	"[DANGEROUS] Execute a command on a remote host via SSH. Supports password and key authentication.",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[SSHRemoteExecInput]) (*mcp.CallToolResultFor[any], error) {
		in := params.Arguments
		if in.Host == "" {
			in.Host = DefaultHost
		}
		if in.User == "" {
			in.User = DefaultUser
		}
		if in.Password == "" && in.KeyPath == "" && in.KeyContent == "" {
			in.KeyPath = DefaultKeyPath
			in.Password = DefaultPassword
		}
		if in.Port == 0 {
			in.Port = DefaultPort
		}
		if in.Host == "" || in.User == "" || in.Command == "" {
			return errResult("host, user, and command are required")
		}
		if in.Password == "" && in.KeyPath == "" && in.KeyContent == "" {
			return errResult("one of password, keyPath, or keyContent is required")
		}
		port := in.Port
		if port == 0 {
			port = 22
		}
		timeout := time.Duration(in.Timeout) * time.Second
		if timeout == 0 {
			timeout = 15 * time.Second
		}

		auths, err := buildAuthMethods(in)
		if err != nil {
			return errResult(err.Error())
		}

		cfg := &ssh.ClientConfig{
			User:            in.User,
			Auth:            auths,
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Timeout:         timeout,
		}
		addr := fmt.Sprintf("%s:%d", in.Host, port)
		client, err := ssh.Dial("tcp", addr, cfg)
		if err != nil {
			return errResult(fmt.Sprintf("ssh dial failed: %v", err))
		}
		defer client.Close()

		session, err := client.NewSession()
		if err != nil {
			return errResult(fmt.Sprintf("ssh session failed: %v", err))
		}
		defer session.Close()

		var stdout, stderr strings.Builder
		session.Stdout = &stdout
		session.Stderr = &stderr

		runErr := session.Run(in.Command)
		exitCode := 0
		if runErr != nil {
			if ee, ok := runErr.(*ssh.ExitError); ok {
				exitCode = ee.ExitStatus()
			} else {
				exitCode = -1
			}
		}

		out := map[string]interface{}{
			"host":     in.Host,
			"user":     in.User,
			"port":     port,
			"command":  in.Command,
			"stdout":   stdout.String(),
			"stderr":   stderr.String(),
			"exitCode": exitCode,
		}
		text := fmt.Sprintf("exit=%d\n--- stdout ---\n%s\n--- stderr ---\n%s", exitCode, stdout.String(), stderr.String())
		return &mcp.CallToolResultFor[any]{
			Content:           []mcp.Content{&mcp.TextContent{Text: text}},
			StructuredContent: out,
		}, nil
	},
)

type SSHRemoteExecInput struct {
	Host          string `json:"host" jsonschema:"remote host address (IP or domain)"`
	Port          int    `json:"port,omitempty" jsonschema:"SSH port (default 22)"`
	User          string `json:"user" jsonschema:"SSH username"`
	Password      string `json:"password,omitempty" jsonschema:"SSH password (use either password or key)"`
	KeyPath       string `json:"keyPath,omitempty" jsonschema:"path to private key file on local machine"`
	KeyContent    string `json:"keyContent,omitempty" jsonschema:"raw private key content (PEM)"`
	KeyPassphrase string `json:"keyPassphrase,omitempty" jsonschema:"passphrase for encrypted private key"`
	Command       string `json:"command" jsonschema:"command to execute on the remote host"`
	Timeout       int    `json:"timeout,omitempty" jsonschema:"connection timeout in seconds (default 15)"`
}

// SSHPortCheckTool tests TCP connectivity to a remote SSH port.
var SSHPortCheckTool = mcp.NewServerTool[SSHPortCheckInput, any](
	"ssh_port_check",
	"Check if an SSH port is open/reachable on a remote host",
	func(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[SSHPortCheckInput]) (*mcp.CallToolResultFor[any], error) {
		in := params.Arguments
		if in.Host == "" {
			in.Host = DefaultHost
		}
		if in.Port == 0 {
			in.Port = DefaultPort
		}
		if in.Host == "" {
			return errResult("host is required")
		}
		port := in.Port
		if port == 0 {
			port = 22
		}
		addr := fmt.Sprintf("%s:%d", in.Host, port)
		conn, err := net.DialTimeout("tcp", addr, 5*time.Second)
		reachable := err == nil
		var msg string
		if reachable {
			conn.Close()
			msg = fmt.Sprintf("%s reachable", addr)
		} else {
			msg = fmt.Sprintf("%s unreachable: %v", addr, err)
		}
		return &mcp.CallToolResultFor[any]{
			Content: []mcp.Content{&mcp.TextContent{Text: msg}},
			StructuredContent: map[string]interface{}{
				"host":      in.Host,
				"port":      port,
				"reachable": reachable,
			},
		}, nil
	},
)

type SSHPortCheckInput struct {
	Host string `json:"host" jsonschema:"remote host address"`
	Port int    `json:"port,omitempty" jsonschema:"SSH port (default 22)"`
}

func buildAuthMethods(in SSHRemoteExecInput) ([]ssh.AuthMethod, error) {
	var auths []ssh.AuthMethod
	if in.Password != "" {
		auths = append(auths, ssh.Password(in.Password))
		auths = append(auths, ssh.KeyboardInteractive(func(name, instruction string, questions []string, echos []bool) ([]string, error) {
			answers := make([]string, len(questions))
			for i := range answers {
				answers[i] = in.Password
			}
			return answers, nil
		}))
	}
	if in.KeyContent != "" {
		signer, err := parseKey([]byte(in.KeyContent), in.KeyPassphrase)
		if err != nil {
			return nil, fmt.Errorf("parse keyContent: %w", err)
		}
		auths = append(auths, ssh.PublicKeys(signer))
	}
	if in.KeyPath != "" {
		data, err := os.ReadFile(in.KeyPath)
		if err != nil {
			return nil, fmt.Errorf("read keyPath: %w", err)
		}
		signer, err := parseKey(data, in.KeyPassphrase)
		if err != nil {
			return nil, fmt.Errorf("parse keyPath: %w", err)
		}
		auths = append(auths, ssh.PublicKeys(signer))
	}
	if len(auths) == 0 {
		return nil, errors.New("no auth methods provided")
	}
	return auths, nil
}

func parseKey(data []byte, passphrase string) (ssh.Signer, error) {
	if passphrase != "" {
		return ssh.ParsePrivateKeyWithPassphrase(data, []byte(passphrase))
	}
	return ssh.ParsePrivateKey(data)
}

func errResult(msg string) (*mcp.CallToolResultFor[any], error) {
	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: msg}},
		IsError: true,
	}, errors.New(msg)
}
