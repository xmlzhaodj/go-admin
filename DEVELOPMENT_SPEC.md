# GoAdmin fork 接入与兼容约定

## 模块身份与发布

- [go.mod](go.mod) 的模块路径为 `github.com/GoAdminGroup/go-admin`，与现有调用方及主题的 import 一致。Git fork 的地址为 `xmlzhaodj/go-admin`。
- `small_admin/go.mod` 当前依赖上游 `v1.2.26`。本地 `go.work` 优先使用 fork 工作树，框架内部版本由 [modules/system/version.go](modules/system/version.go) 给出。这个源码基线的内部版本为 `v1.2.27`。
- 本地接入见 [README_FORK.md](README_FORK.md)。部署或 CI 若需要 fork 行为，应明确固定该版本的 fork 提交或标签，并验证实际构建所选依赖。仅存在于开发机的 `go.work` 不会随部署自动生效。

## 宿主使用的框架接口

| 接入点 | 当前契约 |
| --- | --- |
| `Engine.AddConfigFromYAML`、`Use` | 初始化数据库和站点设置，注册管理插件及路由。宿主先导入对应适配器、驱动和主题 |
| `Engine.HTML` | 宿主回调返回 `(types.Panel, error)`，由主题渲染。回调错误走警告面板 |
| `Engine.Data` | 宿主回调写入响应。导出应核对最终 Content-Type、下载响应头和 XLSX 内容 |
| `ResolveMysqlConnection` | 将 `db.Connection` 交给宿主。`Query(query, args...)` 和 `Exec(query, args...)` 已支持向底层传递参数 |
| `template.Get`、`types.Panel` 和组件 | 宿主用当前主题构建表单、表格和图表，修改通用组件应同时核对宿主正在用的 Sword 主题 |

[引擎接口](engine/engine.go) · [MySQL 实现](modules/db/mysql.go) · [SQL 执行与结果转换](modules/db/performer.go)

## 改动归属

- 通用会话、权限、路由、组件或表单筛选行为在 GoAdmin fork 维护。若修改 `FilterOperatorLike`，对照 [operators.go](template/types/operators.go) 和 [operators_test.go](template/types/operators_test.go) 核对默认筛选控件。
- 游戏报表筛选、业务 SQL、缓存、指标和 Excel 由 `small_admin` 维护。例如业务查询参数绑定可以直接使用现有 `db.Connection.Query`。
- 框架管理数据使用 `goadmin_*`。本项目部署由 `small_admin/admin.sql` 提供表及种子数据，游戏日志表由 `small_tool/log_to_db/small_server.sql` 初始化。[tests/data/admin.sql](tests/data/admin.sql) 是 GoAdmin 集成测试夹具。
