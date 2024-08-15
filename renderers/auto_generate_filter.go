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
 * 是否默认收起
 */
func (a *AutoGenerateFilter) Defaultcollapsed(value interface{}) *AutoGenerateFilter {
    a.Set("defaultCollapsed", value)
    return a
}

/**
 * 过滤条件单行列数
 */
func (a *AutoGenerateFilter) Columnsnum(value interface{}) *AutoGenerateFilter {
    a.Set("columnsNum", value)
    return a
}

/**
 * 是否显示设置查询字段
 */
func (a *AutoGenerateFilter) Showbtntoolbar(value interface{}) *AutoGenerateFilter {
    a.Set("showBtnToolbar", value)
    return a
}
