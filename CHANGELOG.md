# Changelog

本项目遵循 [Keep a Changelog](https://keepachangelog.com/zh-CN/1.1.0/) 和[语义化版本](https://semver.org/lang/zh-CN/)。

从 v6 起，版本号与所映射的 amis 版本保持一致。当前模块路径为 `github.com/wcz0/gamis/v6`，旧路径 `github.com/wcz0/gamis` 停留在 v0.6.19。

## [6.13.0] - 2026-10-23

### Added

- 基于 amis 6.13.0 映射重新生成渲染器，映射来源为 owl-admin 6.13.0 的 `php-renderers`。
- 新增 19 个渲染器：`BaseButtonGroup`、`BaseCards`、`BaseCollapse`、`BaseComboControl`、`BaseDateRangeControl`、`BaseForm`、`BaseInputFormulaControl`、`BaseList`、`BaseRemark`、`BaseTable`、`BaseTransferControl`、`CRUD2`、`CRUDToolbar`、`FormBaseControl`、`FormOptions`、`InputDateBaseControl`、`InputTextAddOn`、`Schema`、`SchemaQuickEdit`。
- 新增 160+ 个链式配置方法。
- `amis.go` 新增 `CRUD2()`、`CRUDToolbar()`、`Schema()`、`SchemaQuickEdit()`、`InputTextAddOn()`、`FormOptions()`、`BaseForm()` 等便捷函数。

### Changed

- 模块路径升级为 `github.com/wcz0/gamis/v6`。
- `CRUD2` 构造时自动写入 `type: "crud2"`。

### Fixed

- 修复 `CRUD2` 缺少 `type` 字段的问题。
- 保持 `Grid` 的既有映射不变，避免上游 6.13 映射回退影响现有页面。

### Compatibility

- 保留已有链式方法与 `TestIdBuilder`。
- `Root()` 保留为 deprecated 包装，等价于 `Page()`。

## 早期版本

| 版本 | 日期 | 变更 |
| --- | --- | --- |
| 0.6.19 | 2025-09-26 | 删除无用的 `go.sum` |
| 0.6.18 | 2025-09-26 | 移除独立 `renderers` 模块，更新 `Component` 函数签名 |
| 0.6.17 | 2025-09-26 | 添加 `renderers` 模块依赖 |
| 0.6.16 | 2025-09-26 | 移除 `renderers` 模块的本地 `replace` |
| 0.6.15 | 2025-09-26 | 更新多个渲染器组件并添加新功能 |
| 0.6.14 / 0.6.13 | 2025-09-26 | 更新渲染器到 amis 6.13.0 |
| 0.6.12 | 2024-09-29 | 细节修复 |
| 0.6.11 | 2024-09-29 | `html` 返回 `*tpl` |
| 0.6.10 / 0.6.9 | 2024-08-15 | 细节修复 |
| 0.6.8 | 2024-08-15 | 增加 `Set` 方法 |
| 0.6.7 | 2024-08-15 | 更新 amis 版本 |
| 0.6.6 | 2024-07-05 | 发布 0.6.6 |
| 0.6.4 | 2024-05-11 | 文档更新 |

更早版本见 git tag。
