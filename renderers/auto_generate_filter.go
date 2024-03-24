package renderers


/**

 * @author wcz0
 * @version 6.2.2
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
 * 是否默认收起
 */
func (a *AutoGenerateFilter) DefaultCollapsed(value interface{}) *AutoGenerateFilter {
    a.Set("defaultCollapsed", value)
    return a
}

/**
 * 过滤条件单行列数
 */
func (a *AutoGenerateFilter) ColumnsNum(value interface{}) *AutoGenerateFilter {
    a.Set("columnsNum", value)
    return a
}

/**
 * 是否显示设置查询字段
 */
func (a *AutoGenerateFilter) ShowBtnToolbar(value interface{}) *AutoGenerateFilter {
    a.Set("showBtnToolbar", value)
    return a
}
