# GoAdmin fork 开发记录

## 2026-11-12 文档与本地接入基线

以 `c47763c7` 为源码基线。核对框架启动、Gin 适配、管理插件、主题、MySQL 连接，以及 `small_admin` 的页面和导出入口。

- Go 1.27.4（macOS ARM64）下，`context`、`template`、`template/types` 和后台插件 `parameter` 包测试通过，`table` 包编译通过且该次无可运行测试。
- 从 `small_admin` 的 `go.work` 构建成功。模块解析指向同级 `go-admin`，构建产物中 `Engine.Use` 的源文件路径也指向本地 fork。
- 本轮验证覆盖框架窄测试及宿主编译。登录、权限、业务查询、Excel 和多数据库或多适配器矩阵未在本轮运行。复验步骤见 [DEVELOPMENT_HANDOFF.md](DEVELOPMENT_HANDOFF.md)。
