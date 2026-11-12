# GoAdmin fork 开发交接

## 当前接入

- fork 基线为 `c47763c7`。本地 `main`、`origin/main` 与已存在的 `upstream/main` 引用在本轮检查时一致。
- 同级 `small_admin/go.work` 已生效。构建产物中的 `Engine.Use` 指向本地 fork。部署环境是否采用 fork 仍需在实际发布配置中核对。
- 本轮建立文档和验证基线。首次扩展时，先确定它属于 GoAdmin 的通用行为还是 `small_admin` 的业务行为，再选择对应入口。

## 验证顺序

在 GoAdmin 仓库根目录从对应包开始。下列命令覆盖请求上下文、模板、筛选参数以及表格包的编译。

```sh
GOWORK=off go test -mod=readonly ./context ./template ./template/types ./plugins/admin/modules/parameter ./plugins/admin/modules/table
```

再用 [README_FORK.md](README_FORK.md) 的命令，从 `small_admin` 工作区构建并核对 fork 路径。改动登录、权限、模板或导出时，用隔离的 MySQL schema 和测试角色复验登录、报表、权限和下载。

完整 [Makefile](Makefile) 测试目标包含多种适配器、数据库和浏览器验收。它会导入或重建测试数据，部分目标会修改依赖文件。运行前核对目标和测试配置，将测试库与现有 `small_server` 隔离。当前未执行完整测试矩阵或浏览器验收。

[已验证基线](DEVELOPMENT_HISTORY.md)
