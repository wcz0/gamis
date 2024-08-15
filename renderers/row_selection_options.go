package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type RowSelectionOptions struct {
	*BaseRenderer
}

func NewRowSelectionOptions() *RowSelectionOptions {
    a := &RowSelectionOptions{
        BaseRenderer: NewBaseRenderer(),
    }

func (a *RowSelectionOptions) Set(name string, value interface{}) *RowSelectionOptions {
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
 * 选择类型 选择全部
 */
func (a *RowSelectionOptions) Key(value interface{}) *RowSelectionOptions {
    a.Set("key", value)
    return a
}

/**
 * 选项显示文本
 */
func (a *RowSelectionOptions) Text(value interface{}) *RowSelectionOptions {
    a.Set("text", value)
    return a
}
