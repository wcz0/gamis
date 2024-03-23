package renderers


/**

*/
type RowSelectionOptions struct {
	*BaseRenderer
}

func NewRowSelectionOptions() *RowSelectionOptions {
    a := &RowSelectionOptions{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}
/**
 * 选择类型 选择全部
 */
func (a *RowSelectionOptions) Key(value string) *RowSelectionOptions {
    a.Set("key", value)
    return a
}

/**
 * 选项显示文本
 */
func (a *RowSelectionOptions) Text(value string) *RowSelectionOptions {
    a.Set("text", value)
    return a
}
