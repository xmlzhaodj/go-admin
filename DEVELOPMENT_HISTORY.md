# GoAdmin fork 开发记录

## 2026-11-13 网站设置表单扩展

宿主可用 `Admin.SetSiteForm` 扩展内置网站设置。在 `small_admin` 的隔离 MySQL 环境验证下拉框渲染、保存、校验和重启恢复。通过 `go test ./plugins/admin ./plugins/admin/modules/table`。

## 2026-11-12 减少查询结果的重复类型解析

- 在本机 MySQL 隔离库核对普通查询与事务查询的整数、定点数、文本和空值。用合成查询返回 8727 行、49 列，改动前后的结果摘要一致。五次运行的中位耗时由 80.983 毫秒降至 41.810 毫秒。通过 `go vet ./modules/db`，`small_admin` 工作区构建使用了本地 fork。
- 原有 `TestMssqlSQL_WhereIn` 在改动前后的源码均因测试驱动为空而失败，原有 MySQL 类型测试使用的账号配置无法通过本机认证。本轮使用单独创建的 MySQL 库做结果对照。

## 2026-11-12 文档与本地接入基线

以 `c47763c7` 为源码基线。核对框架启动、Gin 适配、管理插件、主题、MySQL 连接，以及 `small_admin` 的页面和导出入口。

- Go 1.27.4（macOS ARM64）下，`context`、`template`、`template/types` 和后台插件 `parameter` 包测试通过，`table` 包编译通过且该次无可运行测试。
- 从 `small_admin` 的 `go.work` 构建成功。模块解析指向同级 `go-admin`，构建产物中 `Engine.Use` 的源文件路径也指向本地 fork。
- 本轮验证覆盖框架窄测试及宿主编译。登录、权限、业务查询、Excel 和多数据库或多适配器矩阵未在本轮运行。复验步骤见 [DEVELOPMENT_HANDOFF.md](DEVELOPMENT_HANDOFF.md)。
