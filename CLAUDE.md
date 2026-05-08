# CLAUDE.md

本文件为 Claude Code（claude.ai/code）在此仓库中工作时提供指引。

## 模块信息

- Go 模块：`github.com/kissme666/socket-x`
- Go 版本：1.26.2
- GitHub：https://github.com/kissme666/socket-x

## 常用命令

```bash
make build       # 静态编译当前平台，产物输出到 build/
make build-all   # 交叉编译全部平台（linux/darwin/windows × amd64/arm64）
make clean       # 删除 build/ 目录
go test -race ./...        # 运行所有测试
go test -run TestFoo ./... # 运行单个测试
go vet ./...               # 静态分析
```

## 项目结构

```
socket-x/
├── cmd/socketx/
│   ├── root.go      # cobra 根命令、banner、logger 初始化
│   ├── server.go    # server 子命令
│   └── client.go    # client 子命令
├── internal/
│   ├── config/      # viper 配置加载，支持 yaml/json
│   ├── logger/      # slog 封装，支持 json/text 格式 + lumberjack 日志轮转
│   └── banner/      # ASCII 启动 banner
├── config.yaml      # 示例配置文件
└── Dockerfile       # 多阶段构建，最终基于 scratch
```

## 配置文件（config.yaml）

```yaml
server:
  addr: ":8888"
  max_conn: 1000

log:
  filename: "logs/socketx.log"
  level: "info"       # debug / info / warn / error
  format: "json"      # json / text
  max_size: 100       # MB，单文件最大体积
  max_backups: 7      # 最多保留文件数
  max_age: 30         # 天，最多保留天数
  compress: true
```

## CI / GitHub Actions

工作流文件：`.github/workflows/ci.yml`

| 触发条件 | 执行 Job |
|---------|---------|
| push / PR → main | test → lint → build |
| push → main | 额外上传产物 Artifact + 构建并推送 Docker 镜像 |
| 打 tag（v*） | 额外发布产物到 GitHub Release |

- Docker 镜像推送到 `ghcr.io/kissme666/socket-x:latest`
- `lint` 使用 `install-mode: goinstall` 确保与 go.mod Go 版本一致

发布新版本：
```bash
git tag v1.0.0
git push origin v1.0.0
```

## 注意事项

- 包名使用 `socketx`（无连字符），目录名为 `socket-x`
- 错误信息遵循 Go 惯例：全小写、无标点、用冒号分隔上下文，如 `load config: %w`
- 日志使用标准库 `slog`，格式和级别通过配置文件控制；性能瓶颈时可替换为 zap 底层实现而无需改调用处
- `.gitignore` 已忽略：`build/`、`*.env`、`.vscode/`、`.idea/`
