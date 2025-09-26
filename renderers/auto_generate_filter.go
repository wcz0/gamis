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


func (a *AutoGenerateFilter) Set(name string, value interface{}) *AutoGenerateFilter {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 是否显示设置查询字段
 */
func (a *AutoGenerateFilter) ShowBtnToolbar(value interface{}) *AutoGenerateFilter {
    a.Set("showBtnToolbar", value)
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
 * 是否启用多选框
 */
func (a *AutoGenerateFilter) EnableBulkActions(value interface{}) *AutoGenerateFilter {
    a.Set("enableBulkActions", value)
    return a
}

/**
 * 启用批量操作的表达式
 */
func (a *AutoGenerateFilter) EnableBulkActionsOn(value interface{}) *AutoGenerateFilter {
    a.Set("enableBulkActionsOn", value)
    return a
}

/**
 * 过滤条件单行列数
 */
func (a *AutoGenerateFilter) ColumnsNum(value interface{}) *AutoGenerateFilter {
    a.Set("columnsNum", value)
    return a
}
