# go-sqlite

[![Go Reference](https://pkg.go.dev/badge/github.com/next-bin/go-sqlite/v2.svg)](https://pkg.go.dev/github.com/next-bin/go-sqlite/v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/next-bin/go-sqlite/v2)](https://goreportcard.com/report/github.com/next-bin/go-sqlite/v2)
[![License: BSD](https://img.shields.io/badge/license-BSD-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)](https://go.dev/)

[English](README.md) | 简体中文

一个无 CGo 依赖的 SQLite 驱动，适配 Go 的 `database/sql` 接口，基于 SQLite 3.53.0 合并版本。

本项目源自 CZ.NIC z.s.p.o. 及贡献者开发的 [modernc.org/sqlite](https://gitlab.com/cznic/sqlite)，在此基础上以独立方式持续演进，并重新组织了模块结构。

## 背景

原项目是一个包含 50 多个紧密耦合 Go 模块的整体式代码库，涵盖 C 语言解析器、编译器、汇编器和代码生成工具链。虽然工程实现颇具规模，但这种架构给下游用户带来了实际挑战：

- **依赖面过大** — 导入 SQLite 驱动会传递性地引入词法分析器、中间表示框架和编译工具链，而这些与数据库操作无关。
- **审计复杂度高** — 数十万行生成代码和手写代码横跨完整工具链，难以进行聚焦的安全审查。
- **维护耦合** — 底层模块的变更会沿依赖链不可预测地传播，增加了意外回归的风险。
- **贡献门槛高** — 整体式结构和相互依赖的构建流程，为希望针对驱动本身进行贡献的开发者设置了较高的入门门槛。

本项目通过将代码重构为边界清晰、自包含的模块来应对这些问题——更易审计、维护和扩展。

## 特性

- 无 CGo：纯 Go 实现，轻松交叉编译
- 完全兼容 `database/sql` 接口
- 纯 Go 实现的虚拟表（vtab）
- 自定义 SQL 函数（标量和聚合）
- 自定义排序规则
- 在线备份与恢复
- 预更新、提交和回滚钩子
- 数据库序列化/反序列化
- WAL 模式支持
- Context 取消支持
- 多种时间格式支持
- 支持 17+ 操作系统/架构组合

## 快速开始

    go get github.com/next-bin/go-sqlite/v2

```go
import (
    "database/sql"
    _ "github.com/next-bin/go-sqlite/v2"
)

func main() {
    db, err := sql.Open("sqlite", ":memory:")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    // ... 使用 database/sql 接口操作数据库
}
```

## 文档

完整 API 文档请参阅 [Go 包文档](https://pkg.go.dev/github.com/next-bin/go-sqlite/v2)。

## 示例

`examples/` 目录包含演示所有驱动功能的可运行程序：

- `file_basic` — 基本的 database/sql 文件数据库操作
- `basics` — CRUD、连接模式、数据类型
- `transactions` — 事务、保存点
- `functions` — 自定义标量和聚合函数
- `collation` — 自定义排序规则
- `backup` — 在线备份与恢复
- `hooks` — 预更新、提交、回滚钩子
- `advanced_types` — 时间格式、BLOB、NULL、布尔
- `concurrency` — 并发协程访问
- `serialization` — 数据库序列化/反序列化
- `context_cancel` — 通过 context 取消查询
- `wal` — WAL 模式与检查点
- `columninfo` — 列元数据查询
- `vtab_basic`、`vtab_csv`、`vtab_match`、`vtab_regexp` — 虚拟表示例

## 虚拟表（vtab）

驱动通过 `github.com/next-bin/go-sqlite/v2/vtab` 包暴露纯 Go API，用于实现 SQLite 虚拟表模块。可以将任意数据源（如向量索引、CSV 文件、远程 API）作为 SQL 表，并与 SQLite 的查询优化器集成。

- 注册：`vtab.RegisterModule(db, name, module)`。注册仅对新连接生效。
- 模式声明：在 `Create` 或 `Connect` 中调用 `ctx.Declare("CREATE TABLE <name>(<cols...>)")`。
- 查询规划（BestIndex）：检查 `info.Constraints`、`info.OrderBy` 和 `info.ColUsed`。设置 `ArgIndex` 和 `Omit` 控制约束处理。
- 执行：`Cursor.Filter(idxNum, idxStr, vals)` 按 `ArgIndex` 指定的顺序接收参数。
- 运算符：常见 SQLite 运算符映射到 `ConstraintOp`（EQ/NE/GT/GE/LT/LE/MATCH/LIKE/GLOB/REGEXP 等）。

## 致谢

基于 CZ.NIC z.s.p.o. 开发的 [modernc.org/sqlite](https://gitlab.com/cznic/sqlite)。
原始 SQLite 由 [D. Richard Hipp](https://www.sqlite.org/) 开发（公共领域）。
