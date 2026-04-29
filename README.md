<div align="center">

<img src="https://capsule-render.vercel.app/api?type=waving&color=gradient&customColorList=12,20,24&height=200&section=header&text=mcp-1panel-full&fontSize=60&fontColor=ffffff&animation=fadeIn&fontAlignY=38&desc=AI-Native%20Server%20Management%20via%201Panel%20MCP&descAlignY=58&descSize=18" alt="Header" />

<p>
  <img src="https://img.shields.io/badge/Go-1.25-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go" />
  <img src="https://img.shields.io/badge/MCP-1.0-purple?style=for-the-badge&logo=anthropic&logoColor=white" alt="MCP" />
  <img src="https://img.shields.io/badge/License-GPL%20v3-blue?style=for-the-badge&logo=gnu&logoColor=white" alt="License" />
  <img src="https://img.shields.io/badge/Tools-90%2B-success?style=for-the-badge&logo=tools&logoColor=white" alt="Tools" />
</p>

<p>
  <img src="https://img.shields.io/github/stars/Aurora100729/mcp-1panel-full?style=for-the-badge&color=yellow&labelColor=0D1117" alt="Stars" />
  <img src="https://img.shields.io/github/forks/Aurora100729/mcp-1panel-full?style=for-the-badge&color=blue&labelColor=0D1117" alt="Forks" />
  <img src="https://img.shields.io/github/issues/Aurora100729/mcp-1panel-full?style=for-the-badge&color=red&labelColor=0D1117" alt="Issues" />
  <img src="https://img.shields.io/github/last-commit/Aurora100729/mcp-1panel-full?style=for-the-badge&color=green&labelColor=0D1117" alt="Last commit" />
</p>

<p>
  <strong>让 AI 助手用自然语言全面操控你的 1Panel 服务器</strong><br/>
  <sub>Claude · Windsurf · Cursor · 任何 MCP 兼容客户端</sub>
</p>

</div>

> 🚀 基于 [1Panel-dev/mcp-1panel](https://github.com/1Panel-dev/mcp-1panel) 深度重构。工具数量从 **11** 个扩展到 **90+**，覆盖 1Panel v2.0 全量 API + 本地 Shell / 文件系统 / SSH 远程执行。

---

## ✨ 亮点

- 🚀 **90+ 工具** — 覆盖 1Panel 几乎所有 API
- 🔧 **通用透传** — `panel_request` 直接调用任意未封装端点
- 🔒 **SSH 远程执行** — 支持密码 / 私钥 / 加密私钥 / keyboard-interactive
- ⚙️ **可配置默认值** — CLI 参数预填充 SSH 主机/用户/密钥，调用时只需 `command`
- 📦 **多传输模式** — `stdio` / `sse` / `streamable-http`
- 🐳 **Docker 友好** — 提供官方镜像
- 🛡️ **结构化输出** — 所有工具返回 JSON 结构化内容，便于 AI 二次处理
- 🐛 **针对 v2.0.15 修复** — 修复多个上游 API 参数校验问题

---

## 🛠 技术栈

<div align="center">

<img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go" />
<img src="https://img.shields.io/badge/Model%20Context%20Protocol-512BD4?style=for-the-badge&logo=anthropic&logoColor=white" alt="MCP" />
<img src="https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white" alt="Docker" />
<img src="https://img.shields.io/badge/Linux-FCC624?style=for-the-badge&logo=linux&logoColor=black" alt="Linux" />
<img src="https://img.shields.io/badge/SSH-231F20?style=for-the-badge&logo=openssh&logoColor=white" alt="SSH" />
<img src="https://img.shields.io/badge/1Panel-005CFF?style=for-the-badge&logo=server&logoColor=white" alt="1Panel" />

</div>

---

## 📖 目录

- [功能总览](#功能总览)
- [快速开始](#快速开始)
- [CLI 参数](#cli-参数)
- [MCP 客户端配置](#mcp-客户端配置)
- [SSH 远程执行](#ssh-远程执行)
- [使用示例](#使用示例)
- [项目结构](#项目结构)
- [故障排查](#故障排查)
- [安全说明](#安全说明)
- [更新日志](#更新日志)

---

## 功能总览

| 分类 | 工具数 | 代表工具 |
|------|:---:|------|
| **通用透传** | 1 | `panel_request` — 任意 1Panel API 端点 |
| **系统/仪表盘** | 2 | `get_system_info` `get_dashboard_info` |
| **网站管理** | 10 | `list_websites` `create_website` `update_website_https` |
| **SSL 证书** | 2 | `list_ssls` `create_ssl` |
| **应用商店** | 9 | `app_store_list` `install_mysql` `install_openresty` `app_operate` |
| **数据库** | 7 | `list_databases` `create_database` `database_backup` `redis_status` |
| **Docker 容器** | 8 | `list_containers` `container_create` `container_exec` `container_logs` |
| **Docker 镜像** | 5 | `list_images` `image_pull` `image_build` `image_prune` |
| **Docker 网络/卷** | 7 | `list_networks` `create_volume` `delete_network` |
| **Docker Compose** | 6 | `list_compose` `compose_up` `compose_operate` |
| **文件管理** | 11 | `panel_file_list` `panel_file_read` `panel_file_compress` `panel_file_wget` |
| **防火墙** | 8 | `firewall_status` `list_firewall_rules` `create_firewall_ip_rule` |
| **定时任务** | 5 | `list_crons` `create_cron` `handle_cron` |
| **进程管理** | 2 | `list_processes` `stop_process` |
| **SSH 服务** | 5 | `ssh_info` `ssh_operate` `ssh_logs` `ssh_generate_key` |
| **远程 SSH 执行** ⭐ | 2 | `ssh_remote_exec` `ssh_port_check` |
| **日志** | 3 | `operation_logs` `login_logs` `system_logs` |
| **监控** | 2 | `monitor_search` `monitor_clean` |
| **备份** | 4 | `list_backup_accounts` `backup_operate` `database_backup` |
| **快照** | 4 | `list_snapshots` `create_snapshot` `recover_snapshot` |
| **系统设置** | 5 | `get_settings` `update_password` `update_panel_port` `panel_upgrade` |
| **工具箱** | 10 | `toolbox_dns` `toolbox_hosts` `toolbox_swap` `toolbox_fail2ban_status` |
| **运行环境** | 3 | `list_runtimes` `create_runtime` `runtime_operate` |
| **本地 Shell** | 1 | `shell_exec` 本地命令执行 |
| **本地文件系统** | 6 | `local_file_read` `local_file_write` `local_file_search` |

> 完整工具列表见 `main.go` 中的 `addTools()` 函数。

---

## 快速开始

### 1. 编译

```bash
git clone https://github.com/Aurora100729/mcp-1panel-full.git
cd mcp-1panel-full
go build -o mcp-1panel-full .
```

需要 Go 1.25+。

### 2. 获取 1Panel API Key

1. 登录 1Panel 面板
2. 右上角 **个人信息** → **API 接口**
3. 启用 API 接口，复制 `API Key`
4. **⚠️ 必须将客户端 IP 加入白名单**

### 3. 运行

```bash
# stdio（默认，与 MCP 客户端集成用）
./mcp-1panel-full --token YOUR_API_KEY --host http://127.0.0.1:9999

# SSE（旧 HTTP 模式）
./mcp-1panel-full --transport sse --addr 0.0.0.0:8000 \
  --token YOUR_API_KEY --host http://127.0.0.1:9999

# Streamable HTTP（推荐 HTTP 模式）
./mcp-1panel-full --transport streamable-http --addr 0.0.0.0:8000 \
  --token YOUR_API_KEY --host http://127.0.0.1:9999
```

### 4. Docker 部署

```bash
docker build -t mcp-1panel-full .

docker run -d --name mcp-1panel \
  -p 8000:8000 \
  mcp-1panel-full \
  --transport streamable-http --addr 0.0.0.0:8000 \
  --token YOUR_API_KEY \
  --host http://host.docker.internal:9999
```

---

## CLI 参数

### 1Panel 连接

| 参数 | 默认 | 说明 |
|------|------|------|
| `--token` | _(必填)_ | 1Panel API Key |
| `--host` | _(必填)_ | 1Panel 地址，如 `http://127.0.0.1:9999` |

### 传输

| 参数 | 默认 | 说明 |
|------|------|------|
| `--transport` | `stdio` | `stdio` / `sse` / `streamable-http` |
| `--addr` | `http://localhost:8000` | HTTP 监听地址（仅 HTTP 传输） |

### SSH 远程执行默认值（可选）

预填后调用 `ssh_remote_exec` 只需 `command` 参数。

| 参数 | 默认 | 说明 |
|------|------|------|
| `--ssh-host` | _空_ | 默认 SSH 主机（IP/域名） |
| `--ssh-user` | _空_ | 默认 SSH 用户名 |
| `--ssh-key` | _空_ | 默认私钥文件路径 |
| `--ssh-password` | _空_ | 默认 SSH 密码（不推荐，明文） |
| `--ssh-port` | `22` | 默认 SSH 端口 |

---

## MCP 客户端配置

### Windsurf / Claude Desktop / Cursor

在 `mcp_config.json` 中：

```json
{
  "mcpServers": {
    "mcp-1panel-full": {
      "command": "C:\\path\\to\\mcp-1panel-full.exe",
      "args": [
        "--token", "YOUR_API_KEY",
        "--host", "http://127.0.0.1:9999",
        "--ssh-host", "your.server.ip",
        "--ssh-user", "ubuntu",
        "--ssh-key", "C:\\Users\\You\\.ssh\\id_rsa",
        "--ssh-port", "22"
      ]
    }
  }
}
```

### Linux / macOS

```json
{
  "mcpServers": {
    "mcp-1panel-full": {
      "command": "/usr/local/bin/mcp-1panel-full",
      "args": [
        "--token", "YOUR_API_KEY",
        "--host", "http://127.0.0.1:9999"
      ]
    }
  }
}
```

---

## SSH 远程执行

### 调用方式

**方式 1 — 完整参数（每次提供）：**

```json
{
  "host": "1.2.3.4",
  "user": "ubuntu",
  "keyPath": "/path/to/id_rsa",
  "command": "df -h"
}
```

**方式 2 — 使用 CLI 默认值（启动时已配置）：**

```json
{ "command": "df -h" }
```

### 认证方式

| 字段 | 说明 |
|------|------|
| `password` | 密码认证 |
| `keyPath` | 私钥文件路径 |
| `keyContent` | 私钥内容（PEM 格式直接粘贴） |
| `keyPassphrase` | 加密私钥的口令 |

支持自动 keyboard-interactive 回退（兼容禁用 `PasswordAuthentication` 的服务器）。

### 输出

```
exit=0
--- stdout ---
ubuntu
VM-0-5-ubuntu

--- stderr ---
```

同时返回结构化字段：`{ host, user, port, command, stdout, stderr, exitCode }`。

---

## 使用示例

让 AI 助手用自然语言操作：

### 系统监控

> 帮我看看服务器现在的 CPU、内存、磁盘使用率

→ AI 调用 `get_dashboard_info`

### 应用部署

> 帮我安装一个 Redis，端口 6380，密码 abc123

→ AI 调用 `app_store_list` 找 Redis → `panel_request` POST `/apps/install`

### 容器管理

> 列出所有容器，把停止的都启动

→ `list_containers` → `container_operate(operation: "start", names: [...])`

### 远程命令

> 帮我看看远程服务器的磁盘使用率，并清理 /tmp 下大于 100MB 的文件

→ `ssh_remote_exec(command: "df -h && find /tmp -size +100M -delete")`

### 数据库备份

> 备份所有 MySQL 数据库到本地

→ `list_databases` → 循环 `database_backup`

### 故障排查

> Nginx 启动失败，帮我看看日志

→ `list_installed_apps` 找 nginx → `app_installed_detail` → `container_logs`

---

## 项目结构

```
├── main.go                       # 入口 + CLI flags + 工具注册
├── utils/                        # HTTP 客户端 + 通用工具
├── operations/                   # 1Panel API 工具（按领域分组）
│   ├── generic/                  #   panel_request 通用透传
│   ├── system/                   #   系统/仪表盘
│   ├── website/                  #   网站管理
│   ├── ssl/                      #   SSL 证书
│   ├── app/                      #   应用商店
│   ├── database/                 #   MySQL/PostgreSQL/Redis
│   ├── container/                #   Docker 容器/镜像/网络/卷/Compose
│   ├── file/                     #   文件管理
│   ├── firewall/                 #   防火墙规则
│   ├── cron/                     #   定时任务
│   ├── process/                  #   进程管理
│   ├── sshmanage/                #   SSH 服务管理 + 远程执行 ⭐
│   ├── panellog/                 #   操作/登录/系统日志
│   ├── monitor/                  #   性能监控
│   ├── backup/                   #   备份/恢复
│   ├── snapshot/                 #   系统快照
│   ├── setting/                  #   系统设置
│   ├── toolbox/                  #   DNS/Hosts/Swap/Fail2Ban
│   ├── runtime/                  #   PHP/Node.js/Python 运行环境
│   └── types/                    #   公共类型定义
├── tools/                        # 本地能力工具
│   ├── shell/                    #   本地 Shell 执行
│   └── localfs/                  #   本地文件系统
├── logs/                         # 运行日志（自动创建）
├── Dockerfile
├── go.mod
└── README.md
```

---

## 故障排查

### `dial tcp :22:` SSH 错误

CLI 默认 `--ssh-host` 未生效。检查：

1. `mcp_config.json` 中是否传了 `--ssh-host`
2. 重启 MCP 客户端（不仅是工具刷新）
3. 查看 `logs/mcp-1panel-full.log` 是否有 `[ssh] defaults set: ...`

### `请求参数错误` (1Panel API 400)

某些 1Panel 端点对必填字段校验严格。本项目已为常见端点（如 `list_databases`、`list_backup_records`、`firewall_status`、`monitor_search`）注入合理默认值。如仍失败，使用 `panel_request` 透传并显式提供完整 payload。

### `Not Found (code: 404)` (1Panel API)

不同 1Panel 版本 API 路径有变。本项目兼容 v2.0.15。如使用更新版本，可用 `panel_request` 调用新端点。

### Token 鉴权失败

1. 确认 1Panel 中 API 接口已启用
2. 确认客户端 IP 已加入 API 白名单
3. Token 区分大小写，不要截断

### MCP 客户端连接断开

```
transport error: transport closed
```

通常需重启 MCP 客户端进程：

- **Windsurf**: `Ctrl+Shift+P` → `Reload Window`
- **Claude Desktop**: 完全退出后重启
- **Cursor**: 重启 Cursor

---

## 安全说明

⚠️ 本工具拥有**对 1Panel 完全控制权 + 远程 SSH 执行权 + 本地 Shell 执行权**，等同于服务器 root 权限。

### 必须遵守

1. **不要在公开/不可信环境运行**
2. **不要将 API Token、SSH 私钥提交到仓库**（已在 `.gitignore` 中屏蔽 `*.pem` `*.key`）
3. **不要使用 `--ssh-password` 在配置文件中明文存储密码**，优先使用密钥
4. **`[DANGEROUS]` 标记的工具会修改/删除资源，AI 调用前应确认**

### 推荐做法

- 私钥放在 `~/.ssh/` 目录并设置 `chmod 600`
- 使用专用 1Panel API Key 并加 IP 白名单
- 对生产环境使用只读 API Key（如 1Panel 后续支持权限分级）
- 定期轮换 API Key 与 SSH 密钥

---

## 更新日志

### v1.0.0 (2026-04)

- ✅ 新增 `ssh_remote_exec` / `ssh_port_check` 工具，支持密码/密钥/加密密钥认证
- ✅ 新增 CLI 参数 `--ssh-host` / `--ssh-user` / `--ssh-key` / `--ssh-password` / `--ssh-port` 作为默认值
- ✅ 修复 1Panel v2.0.15 API 兼容性：
  - `firewall_status` 添加 `name` 必填字段
  - `monitor_search` 自动填充时间窗口
  - `list_databases` 修正 `orderBy` 校验
  - `list_backup_records` 添加 `type` 默认值
  - `list_processes` 改用 `/process/:pid` 端点（v2.0.15 移除了 `/process/listening`）
- ✅ 工具数量扩展到 90+
- ✅ 日志路径改为可执行文件同级目录，避免工作目录差异

---

## 致谢

- [1Panel](https://github.com/1Panel-dev/1Panel) — 现代化 Linux 服务器管理面板
- [1Panel-dev/mcp-1panel](https://github.com/1Panel-dev/mcp-1panel) — 原始 MCP Server 项目
- [Model Context Protocol](https://modelcontextprotocol.io/) — MCP 规范
- [golang.org/x/crypto/ssh](https://pkg.go.dev/golang.org/x/crypto/ssh) — Go SSH 客户端

---

## License

GNU General Public License v3.0 — 详见 [LICENSE](LICENSE)

本项目基于 GPL-3.0 协议开源。任何修改、分发、衍生作品必须以相同协议开源并保留版权声明。

---

## ⭐ Star History

<div align="center">

<a href="https://www.star-history.com/#Aurora100729/mcp-1panel-full&Date">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="https://api.star-history.com/svg?repos=Aurora100729/mcp-1panel-full&type=Date&theme=dark" />
    <img src="https://api.star-history.com/svg?repos=Aurora100729/mcp-1panel-full&type=Date" alt="Star History Chart" />
  </picture>
</a>

<br/><br/>

<sub>If you find this project useful, please consider giving it a ⭐!</sub>

<img src="https://capsule-render.vercel.app/api?type=waving&color=gradient&customColorList=12,20,24&height=120&section=footer" alt="Footer" />

</div>
