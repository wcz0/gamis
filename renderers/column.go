package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type Column struct {
	*BaseRenderer
}

func NewColumn() *Column {
    a := &Column{
        BaseRenderer: NewBaseRenderer(),
    }

func (a *Column) Set(name string, value interface{}) *Column {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}

    return a
}

/**
 * 指定列合并表达式
 */
func (a *Column) Colspanexpr(value interface{}) *Column {
    a.Set("colSpanExpr", value)
    return a
}

/**
 * 列表头提示
 */
func (a *Column) Remark(value interface{}) *Column {
    a.Set("remark", value)
    return a
}

/**
 */
func (a *Column) Width(value interface{}) *Column {
    a.Set("width", value)
    return a
}

/**
 * 表头单元格样式
 */
func (a *Column) Titleclassname(value interface{}) *Column {
    a.Set("titleClassName", value)
    return a
}

/**
 * 指定列唯一标识
 */
func (a *Column) Name(value interface{}) *Column {
    a.Set("name", value)
    return a
}

/**
 * 指定列内容渲染器
 */
func (a *Column) Type(value interface{}) *Column {
    a.Set("type", value)
    return a
}

/**
 * 表头分组
 */
func (a *Column) Children(value interface{}) *Column {
    a.Set("children", value)
    return a
}

/**
 * 可复制
 */
func (a *Column) Copyable(value interface{}) *Column {
    a.Set("copyable", value)
    return a
}

/**
 * 快速搜索
 */
func (a *Column) Searchable(value interface{}) *Column {
    a.Set("searchable", value)
    return a
}

/**
 * 是否固定在左侧/右侧
 */
func (a *Column) Fixed(value interface{}) *Column {
    a.Set("fixed", value)
    return a
}

/**
 * 当前列是否展示
 */
func (a *Column) Toggled(value interface{}) *Column {
    a.Set("toggled", value)
    return a
}

/**
 * 指定行合并表达式
 */
func (a *Column) Rowspanexpr(value interface{}) *Column {
    a.Set("rowSpanExpr", value)
    return a
}

/**
 * 兼容table列筛选
 */
func (a *Column) Filterable(value interface{}) *Column {
    a.Set("filterable", value)
    return a
}

/**
 * 单元格样式
 */
func (a *Column) Classnameexpr(value interface{}) *Column {
    a.Set("classNameExpr", value)
    return a
}

/**
 * 表格列单元格是否可以获取父级数据域值，默认为true，该配置对当前列内单元格生效
 */
func (a *Column) Canaccesssuperdata(value interface{}) *Column {
    a.Set("canAccessSuperData", value)
    return a
}

/**
 * 指定列标题
 */
func (a *Column) Title(value interface{}) *Column {
    a.Set("title", value)
    return a
}

/**
 * 快速排序
 */
func (a *Column) Sorter(value interface{}) *Column {
    a.Set("sorter", value)
    return a
}

/**
 * 兼容table快速排序
 */
func (a *Column) Sortable(value interface{}) *Column {
    a.Set("sortable", value)
    return a
}

/**
 * 内容居左、居中、居右
 */
func (a *Column) Align(value interface{}) *Column {
    a.Set("align", value)
    return a
}

/**
 * 列样式
 */
func (a *Column) Classname(value interface{}) *Column {
    a.Set("className", value)
    return a
}

/**
 * 配置快速编辑功能
 */
func (a *Column) Quickedit(value interface{}) *Column {
    a.Set("quickEdit", value)
    return a
}
