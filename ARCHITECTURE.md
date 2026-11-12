# GoAdmin 框架架构

GoAdmin 是由宿主程序嵌入的框架。应用负责启动 Web 服务，并在启动时选择 GoAdmin 适配器、数据库驱动、主题和业务路由。仓库根目录是 Go 模块，[examples/gin/main.go](examples/gin/main.go) 是本项目接入所用框架的示例入口。

## 模块与职责

| 路径 | 职责 |
| --- | --- |
| [engine/engine.go](engine/engine.go) | 装载配置、注册服务、初始化插件、向适配器注册 HTML 与数据路由 |
| [adapter/adapter.go](adapter/adapter.go)、[adapter/gin/gin.go](adapter/gin/gin.go) | Web 框架适配接口。Gin 实现把 HTTP 请求交给 GoAdmin 上下文的处理链，再写回响应 |
| [context/context.go](context/context.go) | 请求与响应上下文、路由和处理链 |
| [plugins/admin/admin.go](plugins/admin/admin.go)、[plugins/admin/router.go](plugins/admin/router.go) | 管理插件、系统表生成器和登录、菜单、CRUD 等内置路由 |
| [plugins/admin/modules/table/table.go](plugins/admin/modules/table/table.go) | 表格生成接口、字段、数据加载和表单操作等配置 |
| [template/template.go](template/template.go)、[template/types](template/types) | 主题与组件注册、面板渲染、字段及筛选配置 |
| [modules/db/connection.go](modules/db/connection.go)、[modules/db/performer.go](modules/db/performer.go) | 连接接口、SQL 执行和结果类型转换 |
| [modules/auth](modules/auth)、[modules/config](modules/config)、[modules/menu](modules/menu) | 认证与权限、应用和站点配置、菜单加载 |

## small_admin 的启动与请求路径

1. 宿主注册 Gin 适配器、MySQL 驱动、外部 Sword 主题和 Chart.js 组件。引擎读取宿主 YAML，初始化连接并核对主题与框架版本要求。
2. `Engine.Use` 依次初始化 CSRF 服务、站点设置和插件，再把插件路由交给适配器。站点设置会访问 `goadmin_site`。
3. 宿主通过 `ResolveMysqlConnection` 获得业务查询连接，并用 `HTML` 注册报表页面、用 `Data` 注册 XLSX 输出。两种 API 默认使用认证中间件。
4. `HTML` 调用业务回调并把 `types.Panel` 交给主题渲染，加入菜单、资源和页面操作。通过 `Data` 注册的回调直接写 GoAdmin 响应，Gin 适配器负责写回 HTTP 客户端。

GoAdmin 在后台提供两条不同的报表路径。管理插件生成的通用表格走表生成器、参数解析和 CRUD 控制器。`small_admin` 的游戏报表由宿主自己查询和构造组件，通过 `Engine.HTML` 和 `Engine.Data` 接入。
