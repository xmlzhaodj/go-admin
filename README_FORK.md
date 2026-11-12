# GoAdmin fork 开发入口

本仓库是 `xmlzhaodj/go-admin`，用于为同级项目 `small_admin` 开发和验证框架层改动。GoAdmin 提供后台路由、登录与权限、菜单、通用表格与表单、主题和数据库连接。游戏报表、指标和 Excel 在 `small_admin`，日志导入和 `ss_*` 表定义在 `small_tool/log_to_db`。

上游文档见 [README_CN.md](README_CN.md)、[README.md](README.md) 和 [CONTRIBUTING_CN.md](CONTRIBUTING_CN.md)。

## 本地接入

fork 的 [go.mod](go.mod) 保留模块名 `github.com/GoAdminGroup/go-admin`，声明 Go 1.22.10。`small_admin` 声明 Go 1.24.4，本地联合开发使用至少 Go 1.24.4。

两仓库为同级目录时，在 `small_admin` 根目录运行下列命令。已有 `go.work` 时跳过初始化，直接核对模块目录。

```sh
go work init . ../go-admin
go list -m -f '{{.Dir}}' github.com/GoAdminGroup/go-admin
go build -o /tmp/small_admin-fork-check .
```

模块目录应指向同级的 `go-admin`。可以用 `go version -m /tmp/small_admin-fork-check` 查看依赖模块标为 `(devel)`，再用 `go tool objdump -s 'github.com/GoAdminGroup/go-admin/engine\.\(\*Engine\)\.Use$' /tmp/small_admin-fork-check` 检查二进制中的源文件路径。业务项目的本地工作区与正式部署的依赖固定方式见[开发规格](DEVELOPMENT_SPEC.md)。

## 文档

| 文档 | 内容 |
| --- | --- |
| [AGENTS.md](AGENTS.md) | 修改与验证约定 |
| [ARCHITECTURE.md](ARCHITECTURE.md) | 框架分层、启动和请求路径 |
| [DEVELOPMENT_SPEC.md](DEVELOPMENT_SPEC.md) | 对 `small_admin` 的接入契约和改动归属 |
| [DEVELOPMENT_HANDOFF.md](DEVELOPMENT_HANDOFF.md) | 当前接入状态和复验步骤 |
| [DEVELOPMENT_HISTORY.md](DEVELOPMENT_HISTORY.md) | 代码基线与实测记录 |
