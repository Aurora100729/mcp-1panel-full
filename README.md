# mcp-1panel-full

**深度优化版 1Panel MCP Server** — 支持 1Panel 全量 API + 本地 Docker / SSH / Shell / 文件系统操作。

> 基于 [1Panel-dev/mcp-1panel](https://github.com/1Panel-dev/mcp-1panel) 深度重构，工具数量从 11 个扩展到 **90+**。

---

## 功能总览

| 分类 | 工具数 | 说明 |
|------|--------|------|
| **通用透传** | 1 | `panel_request` — 任意 1Panel API 端点 |
| **系统/仪表盘** | 2 | 系统信息、仪表盘监控 |
| **网站管理** | 10 | 列表/创建/删除/配置/HTTPS/域名绑定 |
| **SSL 证书** | 2 | 列表/申请 |
| **应用商店** | 9 | 商店搜索/详情/安装/卸载/启停/参数修改 |
| **数据库** | 7 | MySQL/PostgreSQL/Redis 列表/创建/删除/状态/备份 |
| **Docker 容器** | 8 | 列表/创建/启停/删除/日志/统计/exec/prune |
| **Docker 镜像** | 5 | 列表/拉取/删除/构建/清理 |
| **Docker 网络/卷** | 7 | 列表/创建/删除/清理 |
| **Docker Compose** | 6 | 列表/创建/操作/模板/Docker info/清理 |
| **文件管理** | 11 | 列表/读写/创建/删除/移动/压缩/解压/权限/下载/搜索 |
| **防火墙** | 8 | 状态/启停/端口规则/IP规则/转发规则 |
| **定时任务** | 5 | 列表/创建/删除/手动触发/启停 |
| **进程管理** | 2 | 列表/停止 |
| **SSH 管理** | 5 | SSH 状态/启停/配置/日志/密钥生成 |
| **日志** | 3 | 操作日志/登录日志/系统日志 |
| **监控** | 2 | 历史监控数据/清理 |
| **备份** | 4 | 备份账号/记录/创建备份/恢复 |
| **快照** | 4 | 列表/创建/恢复/删除 |
| **系统设置** | 5 | 获取/更新设置/密码/端口/升级 |
| **工具箱** | 10 | DNS/Hosts/Swap/时区/Fail2Ban/ClamAV |
| **运行环境** | 3 | 列表/创建/操作 |
| **本地 Shell** | 1 | 本地命令执行 (bash/sh/cmd/powershell) |
| **本地文件系统** | 6 | 读/写/删/列/stat/搜索 |
| **远程 SSH** | 2 | SSH 远程执行/端口检测 |

## 快速开始

### 编译

```bash
git clone https://github.com/Aurora100729/mcp-1panel-full.git
cd mcp-1panel-full
go build -o mcp-1panel-full .
```

### 运行

```bash
# stdio 模式 (MCP 标准)
./mcp-1panel-full --token YOUR_API_KEY --host http://127.0.0.1:9999

# SSE HTTP 模式
./mcp-1panel-full --transport sse --addr 0.0.0.0:8000 --token YOUR_API_KEY --host http://127.0.0.1:9999

# Streamable HTTP 模式
./mcp-1panel-full --transport streamable-http --addr 0.0.0.0:8000 --token YOUR_API_KEY --host http://127.0.0.1:9999
```

### Docker

```bash
docker build -t mcp-1panel-full .
docker run -d \
  -p 8000:8000 \
  mcp-1panel-full \
  --transport sse --addr 0.0.0.0:8000 --token YOUR_API_KEY --host http://host.docker.internal:9999
```

### MCP 客户端配置 (Windsurf / Claude Desktop / Cursor)

```json
{
  "mcpServers": {
    "mcp-1panel-full": {
      "command": "path/to/mcp-1panel-full",
      "args": ["--token", "YOUR_API_KEY", "--host", "http://127.0.0.1:9999"]
    }
  }
}
```

## CLI 参数

| 参数 | 环境变量 | 默认值 | 说明 |
|------|----------|--------|------|
| `--transport` | - | `stdio` | 传输模式: `stdio`, `sse`, `streamable-http` |
| `--addr` | - | `http://localhost:8000` | HTTP 监听地址 |
| `--token` | `PANEL_ACCESS_TOKEN` | - | 1Panel API Key |
| `--host` | `PANEL_HOST` | - | 1Panel 地址，如 `http://127.0.0.1:9999` |

## 项目结构

```
├── main.go                    # 入口 + 工具注册
├── utils/                     # HTTP 客户端 + 工具函数
├── operations/                # 1Panel API 工具 (按领域分组)
│   ├── generic/               #   panel_request 通用透传
│   ├── system/                #   系统信息
│   ├── website/               #   网站管理
│   ├── ssl/                   #   SSL 证书
│   ├── app/                   #   应用商店
│   ├── database/              #   数据库管理
│   ├── container/             #   Docker 容器/镜像/网络/卷/Compose
│   ├── file/                  #   文件管理
│   ├── firewall/              #   防火墙
│   ├── cron/                  #   定时任务
│   ├── process/               #   进程管理
│   ├── sshmanage/             #   SSH 服务管理
│   ├── panellog/              #   日志
│   ├── monitor/               #   监控
│   ├── backup/                #   备份
│   ├── snapshot/              #   快照
│   ├── setting/               #   系统设置
│   ├── toolbox/               #   工具箱
│   ├── runtime/               #   运行环境
│   └── types/                 #   公共类型定义
├── tools/                     # 本地能力工具
│   ├── shell/                 #   本地命令执行
│   ├── localfs/               #   本地文件系统
│   └── localssh/              #   SSH 远程执行
├── Dockerfile
└── README.md
```

## 安全说明

所有工具默认全部开放，标记为 `[DANGEROUS]` 的工具涉及破坏性操作。请确保 MCP Server 运行在可信环境中。

## 致谢

- [1Panel](https://github.com/1Panel-dev/1Panel) - 现代化 Linux 服务器管理面板
- [1Panel-dev/mcp-1panel](https://github.com/1Panel-dev/mcp-1panel) - 原始 MCP Server 项目
- [Model Context Protocol](https://modelcontextprotocol.io/) - MCP 规范

## License

Apache License 2.0
