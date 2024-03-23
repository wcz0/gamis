package renderers


/**

*/
type AutoGenerateFilter struct {
	*BaseRenderer
}

func NewAutoGenerateFilter() *AutoGenerateFilter {
    a := &AutoGenerateFilter{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}
/**
 * 过滤条件单行列数
 */
func (a *AutoGenerateFilter) ColumnsNum(value string) *AutoGenerateFilter {
    a.Set("columnsNum", value)
    return a
}

/**
 * 是否显示设置查询字段
 */
func (a *AutoGenerateFilter) ShowBtnToolbar(value string) *AutoGenerateFilter {
    a.Set("showBtnToolbar", value)
    return a
}

/**
 * 是否默认收起
 */
func (a *AutoGenerateFilter) DefaultCollapsed(value string) *AutoGenerateFilter {
    a.Set("defaultCollapsed", value)
    return a
}