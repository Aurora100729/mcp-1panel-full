package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/1Panel-dev/mcp-1panel/operations/app"
	"github.com/1Panel-dev/mcp-1panel/operations/database"
	"github.com/1Panel-dev/mcp-1panel/operations/ssl"
	"github.com/1Panel-dev/mcp-1panel/operations/system"
	"github.com/1Panel-dev/mcp-1panel/operations/website"
	"github.com/1Panel-dev/mcp-1panel/utils"
)

var (
	Version = utils.Version
)

func setupLogger() (*os.File, error) {
	logDir := "logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		fmt.Printf("create log dir error: %v\n", err)
		return nil, err
	}

	logFilePath := filepath.Join(logDir, "mcp-1panel.log")
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open log file error: %v\n", err)
		return nil, err
	}

	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	return logFile, nil
}

func newMCPServer() *mcp.Server {
	return mcp.NewServer(
		"github.com/1Panel-dev/mcp-1panel",
		Version,
		nil,
	)
}

func addTools(s *mcp.Server) {
	s.AddTools(
		system.GetSystemInfoTool,
		system.GetDashboardInfoTool,
		website.ListWebsitesTool,
		website.CreateWebsiteTool,
		ssl.ListSSLsTool,
		ssl.CreateSSLTool,
		app.InstallMySQLTool,
		app.InstallOpenRestyTool,
		app.ListInstalledAppsTool,
		database.ListDatabasesTool,
		database.CreateDatabaseTool,
	)
}

func runServer(transport string, addr string) error {
	mcpServer := newMCPServer()
	addTools(mcpServer)

	log.Printf("Starting MCP server with transport=%s addr=%s", transport, addr)

	switch strings.ToLower(transport) {
	case "stdio":
		ctx := context.Background()
		log.Printf("Run Stdio server")
		stdioTransport := mcp.NewStdioTransport()
		if err := mcpServer.Run(ctx, stdioTransport); err != nil {
			return fmt.Errorf("server error: %w", err)
		}
		return nil
	case "sse":
		return serveSSE(addr, mcpServer)
	case "streamable", "streamable-http":
		return serveStreamableHTTP(addr, mcpServer)
	default:
		return fmt.Errorf("unsupported transport %q", transport)
	}
}

func serveSSE(addr string, server *mcp.Server) error {
	handler := mcp.NewSSEHandler(func(*http.Request) *mcp.Server { return server })
	return serveHTTPTransport("SSE", addr, handler)
}

func serveStreamableHTTP(addr string, server *mcp.Server) error {
	handler := mcp.NewStreamableHTTPHandler(func(*http.Request) *mcp.Server { return server }, nil)
	return serveHTTPTransport("Streamable HTTP", addr, handler)
}

func serveHTTPTransport(label, addr string, handler http.Handler) error {
	listenAddr, basePath, displayAddr, err := parseHTTPAddr(addr)
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle(basePath, handler)
	if basePath != "/" && !strings.HasSuffix(basePath, "/") {
		mux.Handle(basePath+"/", handler)
	}

	log.Printf("%s transport listening on %s", label, displayAddr)
	return http.ListenAndServe(listenAddr, mux)
}

func parseHTTPAddr(raw string) (listenAddr, basePath, displayAddr string, err error) {
	if raw == "" {
		return "", "", "", fmt.Errorf("addr must not be empty")
	}

	parsedInput := raw
	if !strings.Contains(parsedInput, "://") {
		parsedInput = "http://" + parsedInput
	}

	u, err := url.Parse(parsedInput)
	if err != nil {
		return "", "", "", fmt.Errorf("invalid addr %q: %w", raw, err)
	}

	host := u.Host
	if host == "" {
		return "", "", "", fmt.Errorf("addr %q must include host and port (e.g. http://localhost:8000)", raw)
	}

	path := u.Path
	if path == "" {
		path = "/"
	}
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	display := fmt.Sprintf("%s://%s%s", defaultScheme(u.Scheme), host, path)
	return host, path, display, nil
}

func defaultScheme(s string) string {
	if s == "" {
		return "http"
	}
	return s
}

func main() {
	var (
		transport   string
		accessToken string
		host        string
		addr        string
	)
	flag.StringVar(&transport, "transport", "stdio", "Transport type (stdio, sse, streamable-http)")
	flag.StringVar(&addr, "addr", "http://localhost:8000", "Base URL (host, port, optional path) for HTTP transports")
	flag.StringVar(&accessToken, "token", "", "1Panel api key")
	flag.StringVar(&host, "host", "", "1Panel host (example:http://127.0.0.1:9999)")
	flag.Parse()

	if accessToken != "" {
		utils.SetAccessToken(accessToken)
	}
	if host != "" {
		utils.SetHost(host)
	}

	if err := runServer(transport, addr); err != nil {
		fmt.Printf("server run error: %v\n", err)
		panic(err)
	}
}
