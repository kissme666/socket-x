# CLAUDE.md

本文件为 Claude Code（claude.ai/code）在此仓库中工作时提供指引。

## 模块信息

Go 模块：`github.com/kissme666/socketx`（Go 1.26.2）

项目处于早期阶段，目前除 `go.mod` 外尚无源文件。

## 常用命令

```bash
go build ./...       # 构建所有包
go test ./...        # 运行所有测试
go test -run TestFoo # 运行单个测试
go vet ./...         # 静态分析
```
