package renderers


/**

*/
type Column struct {
	*BaseRenderer
}

func NewColumn() *Column {
    a := &Column{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}
/**
 * 配置快速编辑功能
 */
func (a *Column) QuickEdit(value string) *Column {
    a.Set("quickEdit", value)
    return a
}

/**
 * 指定列标题
 */
func (a *Column) Title(value string) *Column {
    a.Set("title", value)
    return a
}

/**
 * 指定行合并表达式
 */
func (a *Column) RowSpanExpr(value string) *Column {
    a.Set("rowSpanExpr", value)
    return a
}

/**
 * 指定列合并表达式
 */
func (a *Column) ColSpanExpr(value string) *Column {
    a.Set("colSpanExpr", value)
    return a
}

/**
 * 快速搜索
 */
func (a *Column) Searchable(value string) *Column {
    a.Set("searchable", value)
    return a
}

/**
 * 兼容table列筛选
 */
func (a *Column) Filterable(value string) *Column {
    a.Set("filterable", value)
    return a
}

/**
 * 内容居左、居中、居右
 */
func (a *Column) Align(value string) *Column {
    a.Set("align", value)
    return a
}

/**
 * 列样式
 */
func (a *Column) ClassName(value string) *Column {
    a.Set("className", value)
    return a
}

/**
 * 表头分组
 */
func (a *Column) Children(value string) *Column {
    a.Set("children", value)
    return a
}

/**
 * 兼容table快速排序
 */
func (a *Column) Sortable(value string) *Column {
    a.Set("sortable", value)
    return a
}

/**
 * 是否固定在左侧/右侧
 */
func (a *Column) Fixed(value string) *Column {
    a.Set("fixed", value)
    return a
}

/**
 * 单元格样式
 */
func (a *Column) ClassNameExpr(value string) *Column {
    a.Set("classNameExpr", value)
    return a
}

/**
 */
func (a *Column) Width(value string) *Column {
    a.Set("width", value)
    return a
}

/**
 * 列表头提示
 */
func (a *Column) Remark(value string) *Column {
    a.Set("remark", value)
    return a
}

/**
 * 指定列唯一标识
 */
func (a *Column) Name(value string) *Column {
    a.Set("name", value)
    return a
}

/**
 * 指定列内容渲染器
 */
func (a *Column) Type(value string) *Column {
    a.Set("type", value)
    return a
}

/**
 * 可复制
 */
func (a *Column) Copyable(value string) *Column {
    a.Set("copyable", value)
    return a
}

/**
 * 快速排序
 */
func (a *Column) Sorter(value string) *Column {
    a.Set("sorter", value)
    return a
}

/**
 * 当前列是否展示
 */
func (a *Column) Toggled(value string) *Column {
    a.Set("toggled", value)
    return a
}

/**
 * 表头单元格样式
 */
func (a *Column) TitleClassName(value string) *Column {
    a.Set("titleClassName", value)
    return a
}

/**
 * 表格列单元格是否可以获取父级数据域值，默认为true，该配置对当前列内单元格生效
 */
func (a *Column) CanAccessSuperData(value string) *Column {
    a.Set("canAccessSuperData", value)
    return a
}
