package renderers


/**
 * Table 表格渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/table
 *

*/
type Table struct {
	*BaseRenderer
}

func NewTable() *Table {
    a := &Table{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "table")
    return a
}
/**
 * 是否固底
 */
func (a *Table) AffixFooter(value string) *Table {
    a.Set("affixFooter", value)
    return a
}

/**
 * 表格的列信息
 */
func (a *Table) Columns(value string) *Table {
    a.Set("columns", value)
    return a
}

/**
 * 底部外层 CSS 类名
 */
func (a *Table) FooterClassName(value string) *Table {
    a.Set("footerClassName", value)
    return a
}

/**
 * 表格 CSS 类名
 */
func (a *Table) TableClassName(value string) *Table {
    a.Set("tableClassName", value)
    return a
}

/**
 * 底部总结行
 */
func (a *Table) AffixRow(value string) *Table {
    a.Set("affixRow", value)
    return a
}

/**
 * 行角标
 */
func (a *Table) ItemBadge(value string) *Table {
    a.Set("itemBadge", value)
    return a
}

/**
 * 懒加载 API，当行数据中用 defer: true 标记了，则其孩子节点将会用这个 API 来拉取数据。
 */
func (a *Table) DeferApi(value string) *Table {
    a.Set("deferApi", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Table) EditorSetting(value string) *Table {
    a.Set("editorSetting", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Table) Static(value string) *Table {
    a.Set("static", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Table) StaticOn(value string) *Table {
    a.Set("staticOn", value)
    return a
}

/**
 * 是否开启底部展示功能，适合移动端展示
 */
func (a *Table) Footable(value string) *Table {
    a.Set("footable", value)
    return a
}

/**
 * 是否显示底部
 */
func (a *Table) ShowFooter(value string) *Table {
    a.Set("showFooter", value)
    return a
}

/**
 * 数据源：绑定当前环境变量
 */
func (a *Table) Source(value string) *Table {
    a.Set("source", value)
    return a
}

/**
 * 是否显示
 */
func (a *Table) Visible(value string) *Table {
    a.Set("visible", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Table) Hidden(value string) *Table {
    a.Set("hidden", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Table) Id(value string) *Table {
    a.Set("id", value)
    return a
}

/**
 */
func (a *Table) StaticSchema(value string) *Table {
    a.Set("staticSchema", value)
    return a
}

/**
 * 是否固定表头
 */
func (a *Table) AffixHeader(value string) *Table {
    a.Set("affixHeader", value)
    return a
}

/**
 * 展示列显示开关，自动即：列数量大于或等于5个时自动开启
 */
func (a *Table) ColumnsTogglable(value string) *Table {
    a.Set("columnsTogglable", value)
    return a
}

/**
 * 标题
 */
func (a *Table) Title(value string) *Table {
    a.Set("title", value)
    return a
}

/**
 * 合并单元格配置，配置数字表示从左到右的多少列自动合并单元格。
 */
func (a *Table) CombineNum(value string) *Table {
    a.Set("combineNum", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Table) DisabledOn(value string) *Table {
    a.Set("disabledOn", value)
    return a
}

/**
 * 表格自动计算高度
 */
func (a *Table) AutoFillHeight(value string) *Table {
    a.Set("autoFillHeight", value)
    return a
}

/**
 * 合并单元格配置，配置从第几列开始合并。
 */
func (a *Table) CombineFromIndex(value string) *Table {
    a.Set("combineFromIndex", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Table) HiddenOn(value string) *Table {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Table) OnEvent(value string) *Table {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Table) StaticInputClassName(value string) *Table {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Table) UseMobileUI(value string) *Table {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 表格是否可以获取父级数据域值，默认为false
 */
func (a *Table) CanAccessSuperData(value string) *Table {
    a.Set("canAccessSuperData", value)
    return a
}

/**
 * table layout
 * 可选值: fixed | auto
 */
func (a *Table) TableLayout(value string) *Table {
    a.Set("tableLayout", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Table) Disabled(value string) *Table {
    a.Set("disabled", value)
    return a
}

/**
 * 组件样式
 */
func (a *Table) Style(value string) *Table {
    a.Set("style", value)
    return a
}

/**
 * 指定为表格渲染器。
 * 可选值: table | static-table
 */
func (a *Table) Type(value string) *Table {
    a.Set("type", value)
    return a
}

/**
 * 是否可调整列宽
 */
func (a *Table) Resizable(value string) *Table {
    a.Set("resizable", value)
    return a
}

/**
 * 开启查询区域，会根据列元素的searchable属性值，自动生成查询条件表单
 */
func (a *Table) AutoGenerateFilter(value string) *Table {
    a.Set("autoGenerateFilter", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Table) ClassName(value string) *Table {
    a.Set("className", value)
    return a
}

/**
 * 占位符
 */
func (a *Table) Placeholder(value string) *Table {
    a.Set("placeholder", value)
    return a
}

/**
 * 是否显示头部
 */
func (a *Table) ShowHeader(value string) *Table {
    a.Set("showHeader", value)
    return a
}

/**
 * 行样式表表达式
 */
func (a *Table) RowClassNameExpr(value string) *Table {
    a.Set("rowClassNameExpr", value)
    return a
}

/**
 * 设置数据
 */
func (a *Table) Data(value string) *Table {
    a.Set("data", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Table) StaticPlaceholder(value string) *Table {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Table) StaticClassName(value string) *Table {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Table) StaticLabelClassName(value string) *Table {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 顶部外层 CSS 类名
 */
func (a *Table) HeaderClassName(value string) *Table {
    a.Set("headerClassName", value)
    return a
}

/**
 * 工具栏 CSS 类名
 */
func (a *Table) ToolbarClassName(value string) *Table {
    a.Set("toolbarClassName", value)
    return a
}

/**
 * 顶部总结行
 */
func (a *Table) PrefixRow(value string) *Table {
    a.Set("prefixRow", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Table) VisibleOn(value string) *Table {
    a.Set("visibleOn", value)
    return a
}