# GoAdmin fork 协作说明

先读 [README_FORK.md](README_FORK.md)。改引擎、适配器、数据库、表单或模板前，按涉及范围查看 [ARCHITECTURE.md](ARCHITECTURE.md) 和 [DEVELOPMENT_SPEC.md](DEVELOPMENT_SPEC.md)。

- 保留 `github.com/GoAdminGroup/go-admin` 模块身份及相关公开接口的兼容性。现有 Go 文件沿用各自的格式，使用 `gofmt` 检查所改 Go 文件。
- 通用框架行为在本仓库维护。游戏查询、报表、导出和指标留在 `small_admin`。日志接收格式及导入逻辑由相应上游维护。
- 改动先跑对应包的窄测试，再用 `small_admin` 本地 Go Workspace 验证实际调用。数据库或浏览器集成测试使用隔离库和合成数据，先检查测试配置及初始化命令的目标。
- 现行接口与行为记入开发规格，验证证据记入 [DEVELOPMENT_HISTORY.md](DEVELOPMENT_HISTORY.md)，未验证事项记入 [DEVELOPMENT_HANDOFF.md](DEVELOPMENT_HANDOFF.md)。
