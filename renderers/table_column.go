package renderers


/**
 * 表格列，不指定类型时默认为文本类型。

 * @author wcz0
 * @version 6.2.2
 */
type TableColumn struct {
	*BaseRenderer
}

func NewTableColumn() *TableColumn {
    a := &TableColumn{
        BaseRenderer: NewBaseRenderer(),
    }
    return a
}

/**
 * 配置点击复制功能
 */
func (a *TableColumn) Copyable(value interface{}) *TableColumn {
    a.Set("copyable", value)
    return a
}

/**
 * 是否唯一, 只有在 inputTable 里面才有用
 */
func (a *TableColumn) Unique(value interface{}) *TableColumn {
    a.Set("unique", value)
    return a
}

/**
 * 表格列单元格是否可以获取父级数据域值，默认为true，该配置对当前列内单元格生效
 */
func (a *TableColumn) CanAccessSuperData(value interface{}) *TableColumn {
    a.Set("canAccessSuperData", value)
    return a
}

/**
 * 配置查看详情功能
 */
func (a *TableColumn) PopOver(value interface{}) *TableColumn {
    a.Set("popOver", value)
    return a
}

/**
 * 作为表单项时，可以单独配置编辑时的快速编辑面板。
 */
func (a *TableColumn) QuickEditOnUpdate(value interface{}) *TableColumn {
    a.Set("quickEditOnUpdate", value)
    return a
}

/**
 * 配置是否默认展示
 */
func (a *TableColumn) Toggled(value interface{}) *TableColumn {
    a.Set("toggled", value)
    return a
}

/**
 * 列宽度
 */
func (a *TableColumn) Width(value interface{}) *TableColumn {
    a.Set("width", value)
    return a
}

/**
 * 提示信息
 */
func (a *TableColumn) Remark(value interface{}) *TableColumn {
    a.Set("remark", value)
    return a
}

/**
 * 默认值, 只有在 inputTable 里面才有用
 */
func (a *TableColumn) Value(value interface{}) *TableColumn {
    a.Set("value", value)
    return a
}

/**
 * 绑定字段名
 */
func (a *TableColumn) Name(value interface{}) *TableColumn {
    a.Set("name", value)
    return a
}

/**
 * 配置是否可以排序
 */
func (a *TableColumn) Sortable(value interface{}) *TableColumn {
    a.Set("sortable", value)
    return a
}

/**
 * 列头样式表
 */
func (a *TableColumn) LabelClassName(value interface{}) *TableColumn {
    a.Set("labelClassName", value)
    return a
}

/**
 * 结合表格的 footable 一起使用。 填写 *、xs、sm、md、lg指定 footable 的触发条件，可以填写多个用空格隔开
 * 可选值: * | xs | sm | md | lg
 */
func (a *TableColumn) Breakpoint(value interface{}) *TableColumn {
    a.Set("breakpoint", value)
    return a
}

/**
 * 当一次性渲染太多列上有用，默认为 100，可以用来提升表格渲染性能
 */
func (a *TableColumn) LazyRenderAfter(value interface{}) *TableColumn {
    a.Set("lazyRenderAfter", value)
    return a
}

/**
 * 单元格内部组件自定义样式 style作为单元格自定义样式的配置
 */
func (a *TableColumn) InnerStyle(value interface{}) *TableColumn {
    a.Set("innerStyle", value)
    return a
}

/**
 * 单元格样式表达式
 */
func (a *TableColumn) ClassNameExpr(value interface{}) *TableColumn {
    a.Set("classNameExpr", value)
    return a
}

/**
 * todo
 */
func (a *TableColumn) Filterable(value interface{}) *TableColumn {
    a.Set("filterable", value)
    return a
}

/**
 * 列标题
 */
func (a *TableColumn) Label(value interface{}) *TableColumn {
    a.Set("label", value)
    return a
}

/**
 * 配置是否固定当前列
 * 可选值: left | right | none
 */
func (a *TableColumn) Fixed(value interface{}) *TableColumn {
    a.Set("fixed", value)
    return a
}

/**
 * 配置快速编辑功能
 */
func (a *TableColumn) QuickEdit(value interface{}) *TableColumn {
    a.Set("quickEdit", value)
    return a
}

/**
 * 是否可快速搜索
 */
func (a *TableColumn) Searchable(value interface{}) *TableColumn {
    a.Set("searchable", value)
    return a
}

/**
 * 列对齐方式
 * 可选值: left | right | center | justify
 */
func (a *TableColumn) Align(value interface{}) *TableColumn {
    a.Set("align", value)
    return a
}

/**
 * 列样式表
 */
func (a *TableColumn) ClassName(value interface{}) *TableColumn {
    a.Set("className", value)
    return a
}
