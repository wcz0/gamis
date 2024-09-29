package renderers


/**
 * 不指定类型默认就是文本

 * @author wcz0
 * @version 6.2.2
 */
type ListBodyField struct {
	*BaseRenderer
}

func NewListBodyField() *ListBodyField {
    a := &ListBodyField{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}


func (a *ListBodyField) Set(name string, value interface{}) *ListBodyField {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 配置点击复制功能
 */
func (a *ListBodyField) Copyable(value interface{}) *ListBodyField {
    a.Set("copyable", value)
    return a
}

/**
 * 列标题
 */
func (a *ListBodyField) Label(value interface{}) *ListBodyField {
    a.Set("label", value)
    return a
}

/**
 * label 类名
 */
func (a *ListBodyField) LabelClassName(value interface{}) *ListBodyField {
    a.Set("labelClassName", value)
    return a
}

/**
 * 内层组件的CSS类名
 */
func (a *ListBodyField) InnerClassName(value interface{}) *ListBodyField {
    a.Set("innerClassName", value)
    return a
}

/**
 * 绑定字段名
 */
func (a *ListBodyField) Name(value interface{}) *ListBodyField {
    a.Set("name", value)
    return a
}

/**
 * 配置查看详情功能
 */
func (a *ListBodyField) PopOver(value interface{}) *ListBodyField {
    a.Set("popOver", value)
    return a
}

/**
 * 配置快速编辑功能
 */
func (a *ListBodyField) QuickEdit(value interface{}) *ListBodyField {
    a.Set("quickEdit", value)
    return a
}
