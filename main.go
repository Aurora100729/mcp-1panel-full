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

	"github.com/Aurora100729/mcp-1panel-full/operations/app"
	"github.com/Aurora100729/mcp-1panel-full/operations/backup"
	"github.com/Aurora100729/mcp-1panel-full/operations/container"
	"github.com/Aurora100729/mcp-1panel-full/operations/cron"
	"github.com/Aurora100729/mcp-1panel-full/operations/database"
	"github.com/Aurora100729/mcp-1panel-full/operations/file"
	"github.com/Aurora100729/mcp-1panel-full/operations/firewall"
	"github.com/Aurora100729/mcp-1panel-full/operations/generic"
	"github.com/Aurora100729/mcp-1panel-full/operations/monitor"
	"github.com/Aurora100729/mcp-1panel-full/operations/panellog"
	"github.com/Aurora100729/mcp-1panel-full/operations/process"
	"github.com/Aurora100729/mcp-1panel-full/operations/runtime"
	"github.com/Aurora100729/mcp-1panel-full/operations/setting"
	"github.com/Aurora100729/mcp-1panel-full/operations/snapshot"
	"github.com/Aurora100729/mcp-1panel-full/operations/sshmanage"
	"github.com/Aurora100729/mcp-1panel-full/operations/ssl"
	"github.com/Aurora100729/mcp-1panel-full/operations/system"
	"github.com/Aurora100729/mcp-1panel-full/operations/toolbox"
	"github.com/Aurora100729/mcp-1panel-full/operations/website"
	"github.com/Aurora100729/mcp-1panel-full/tools/localfs"
	"github.com/Aurora100729/mcp-1panel-full/tools/localssh"
	"github.com/Aurora100729/mcp-1panel-full/tools/shell"
	"github.com/Aurora100729/mcp-1panel-full/utils"
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

	logFilePath := filepath.Join(logDir, "mcp-1panel-full.log")
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
		"github.com/Aurora100729/mcp-1panel-full",
		Version,
		nil,
	)
}

func addTools(s *mcp.Server) {
	s.AddTools(
		// ── Generic passthrough ──
		generic.PanelRequestTool,

		// ── System / Dashboard ──
		system.GetSystemInfoTool,
		system.GetDashboardInfoTool,

		// ── Website ──
		website.ListWebsitesTool,
		website.CreateWebsiteTool,
		website.DeleteWebsiteTool,
		website.GetWebsiteConfigTool,
		website.UpdateWebsiteConfigTool,
		website.WebsiteOperateTool,
		website.GetWebsiteHTTPSTool,
		website.UpdateWebsiteHTTPSTool,
		website.ListWebsiteDomainsTool,
		website.CreateWebsiteDomainTool,

		// ── SSL ──
		ssl.ListSSLsTool,
		ssl.CreateSSLTool,

		// ── App Store ──
		app.InstallMySQLTool,
		app.InstallOpenRestyTool,
		app.ListInstalledAppsTool,
		app.AppStoreTool,
		app.AppDetailTool,
		app.AppInstalledDetailTool,
		app.AppOperateTool,
		app.AppUpdateParamsTool,
		app.AppInstalledParamsTool,

		// ── Database (MySQL / PostgreSQL / Redis) ──
		database.ListDatabasesTool,
		database.CreateDatabaseTool,
		database.DeleteDatabaseTool,
		database.DatabaseBackupTool,
		database.DatabaseStatusTool,
		database.ListRedisTool,
		database.RedisStatusTool,

		// ── Container / Docker ──
		container.ListContainersTool,
		container.ContainerOperateTool,
		container.ContainerInspectTool,
		container.ContainerLogsTool,
		container.ContainerStatsTool,
		container.ContainerCreateTool,
		container.ContainerPruneTool,
		container.ContainerExecTool,

		// ── Image ──
		container.ListImagesTool,
		container.ImagePullTool,
		container.ImageRemoveTool,
		container.ImageBuildTool,
		container.ImagePruneTool,

		// ── Network / Volume ──
		container.ListNetworksTool,
		container.CreateNetworkTool,
		container.DeleteNetworkTool,
		container.ListVolumesTool,
		container.CreateVolumeTool,
		container.DeleteVolumeTool,
		container.VolumePruneTool,

		// ── Compose ──
		container.ListComposeTool,
		container.ComposeUpTool,
		container.ComposeOperateTool,
		container.ListComposeTemplatesTool,
		container.DockerInfoTool,
		container.DockerSystemPruneTool,

		// ── File Management ──
		file.ListFilesTool,
		file.ReadFileTool,
		file.WriteFileTool,
		file.CreateFileTool,
		file.DeleteFileTool,
		file.MoveFileTool,
		file.CompressFileTool,
		file.DecompressFileTool,
		file.ChmodFileTool,
		file.WgetFileTool,
		file.FileSearchContentTool,

		// ── Firewall ──
		firewall.FirewallStatusTool,
		firewall.FirewallOperateTool,
		firewall.ListFirewallRulesTool,
		firewall.CreateFirewallRuleTool,
		firewall.DeleteFirewallRuleTool,
		firewall.ListFirewallIPRulesTool,
		firewall.CreateFirewallIPRuleTool,
		firewall.ListFirewallForwardsTool,

		// ── Cron Jobs ──
		cron.ListCronsTool,
		cron.CreateCronTool,
		cron.DeleteCronTool,
		cron.HandleCronTool,
		cron.UpdateCronStatusTool,

		// ── Process ──
		process.ListProcessesTool,
		process.StopProcessTool,

		// ── SSH Management (1Panel) ──
		sshmanage.SSHInfoTool,
		sshmanage.SSHOperateTool,
		sshmanage.SSHUpdateTool,
		sshmanage.SSHLogsTool,
		sshmanage.SSHGenerateKeyTool,
		sshmanage.SSHRemoteExecTool,
		sshmanage.SSHPortCheckTool,

		// ── Logs ──
		panellog.OperationLogsTool,
		panellog.LoginLogsTool,
		panellog.SystemLogsTool,

		// ── Monitor ──
		monitor.MonitorSearchTool,
		monitor.MonitorCleanTool,

		// ── Backup ──
		backup.ListBackupAccountsTool,
		backup.CreateBackupAccountTool,
		backup.ListBackupRecordsTool,
		backup.BackupOperateTool,

		// ── Snapshot ──
		snapshot.ListSnapshotsTool,
		snapshot.CreateSnapshotTool,
		snapshot.RecoverSnapshotTool,
		snapshot.DeleteSnapshotTool,

		// ── Settings ──
		setting.GetSettingsTool,
		setting.UpdateSettingTool,
		setting.UpdatePasswordTool,
		setting.UpdatePortTool,
		setting.PanelUpgradeTool,

		// ── Toolbox ──
		toolbox.GetDNSTool,
		toolbox.UpdateDNSTool,
		toolbox.GetHostsTool,
		toolbox.UpdateHostsTool,
		toolbox.GetSwapTool,
		toolbox.SwapOperateTool,
		toolbox.GetTimezoneTool,
		toolbox.Fail2BanStatusTool,
		toolbox.Fail2BanOperateTool,
		toolbox.ClamAVScanTool,

		// ── Runtime ──
		runtime.ListRuntimesTool,
		runtime.CreateRuntimeTool,
		runtime.RuntimeOperateTool,

		// ═══════════════════════════════════
		// ── Local Tools (bypass 1Panel) ──
		// ═══════════════════════════════════

		// ── Shell Exec ──
		shell.ShellExecTool,

		// ── Local File System ──
		localfs.LocalFileReadTool,
		localfs.LocalFileWriteTool,
		localfs.LocalFileDeleteTool,
		localfs.LocalFileListTool,
		localfs.LocalFileStatTool,
		localfs.LocalFileSearchTool,

		// ── SSH Remote Exec ──
		localssh.SSHExecTool,
		localssh.SSHPortCheckTool,
	)
}

func runServer(transport string, addr string) error {
	mcpServer := newMCPServer()
	addTools(mcpServer)

	log.Printf("Starting MCP server (v%s) with transport=%s addr=%s", Version, transport, addr)

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
		sshHost     string
		sshUser     string
		sshKey      string
		sshPassword string
		sshPort     int
	)
	flag.StringVar(&transport, "transport", "stdio", "Transport type (stdio, sse, streamable-http)")
	flag.StringVar(&addr, "addr", "http://localhost:8000", "Base URL (host, port, optional path) for HTTP transports")
	flag.StringVar(&accessToken, "token", "", "1Panel api key")
	flag.StringVar(&host, "host", "", "1Panel host (example:http://127.0.0.1:9999)")
	flag.StringVar(&sshHost, "ssh-host", "", "Default SSH host for ssh_remote_exec")
	flag.StringVar(&sshUser, "ssh-user", "", "Default SSH username for ssh_remote_exec")
	flag.StringVar(&sshKey, "ssh-key", "", "Default SSH private key path for ssh_remote_exec")
	flag.StringVar(&sshPassword, "ssh-password", "", "Default SSH password for ssh_remote_exec")
	flag.IntVar(&sshPort, "ssh-port", 22, "Default SSH port for ssh_remote_exec")
	flag.Parse()

	logFile, _ := setupLogger()
	if logFile != nil {
		defer logFile.Close()
	}

	if accessToken != "" {
		utils.SetAccessToken(accessToken)
	}
	if host != "" {
		utils.SetHost(host)
	}
	sshmanage.SetDefaults(sshHost, sshUser, sshKey, sshPassword, sshPort)

	if err := runServer(transport, addr); err != nil {
		fmt.Printf("server run error: %v\n", err)
		panic(err)
	}
}
