# SocketX

基于 Go 实现的轻量级 TCP 代理框架，用于突破网络限制。

English | [中文](README.zh_CN.md)

![CI](https://github.com/kissme666/socket-x/actions/workflows/ci.yml/badge.svg)

## 特性

- TCP 服务端 / 客户端模式
- 支持 YAML / JSON 配置文件
- 结构化日志 + 日志轮转（slog + lumberjack）
- 跨平台静态编译二进制
- Docker 支持

## 安装

**二进制**

从 [GitHub Releases](https://github.com/kissme666/socket-x/releases) 下载对应平台的最新版本。

**Docker**

```bash
docker pull ghcr.io/kissme666/socket-x:latest
```

**从源码构建**

```bash
git clone https://github.com/kissme666/socket-x.git
cd socket-x
make build
```

## 使用

```bash
# 启动服务端
./socketx server -c config.yaml

# 启动客户端
./socketx client -c config.yaml
```

**Docker**

```bash
docker run -p 8888:8888 ghcr.io/kissme666/socket-x:latest

# 挂载自定义配置
docker run -p 8888:8888 -v ./config.yaml:/config.yaml ghcr.io/kissme666/socket-x:latest
```

## 配置

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

## License

MIT
