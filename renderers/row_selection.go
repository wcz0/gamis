package renderers


/**

*/
type RowSelection struct {
	*BaseRenderer
}

func NewRowSelection() *RowSelection {
    a := &RowSelection{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}
/**
 * 行是否禁用表达式
 */
func (a *RowSelection) DisableOn(value string) *RowSelection {
    a.Set("disableOn", value)
    return a
}

/**
 * 自定义选择菜单
 */
func (a *RowSelection) Selections(value string) *RowSelection {
    a.Set("selections", value)
    return a
}

/**
 * 已选择的key值
 */
func (a *RowSelection) SelectedRowKeys(value string) *RowSelection {
    a.Set("selectedRowKeys", value)
    return a
}

/**
 * 已选择的key值表达式
 */
func (a *RowSelection) SelectedRowKeysExpr(value string) *RowSelection {
    a.Set("selectedRowKeysExpr", value)
    return a
}

/**
 * 已选择的key值表达式
 */
func (a *RowSelection) ColumnWidth(value string) *RowSelection {
    a.Set("columnWidth", value)
    return a
}

/**
 * 是否点击行触发选中或取消选中
 */
func (a *RowSelection) RowClick(value string) *RowSelection {
    a.Set("rowClick", value)
    return a
}

/**
 * 选择类型 单选/多选
 */
func (a *RowSelection) Type(value string) *RowSelection {
    a.Set("type", value)
    return a
}

/**
 * 对应数据源的key值
 */
func (a *RowSelection) KeyField(value string) *RowSelection {
    a.Set("keyField", value)
    return a
}
