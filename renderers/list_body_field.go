package renderers


/**
 * 不指定类型默认就是文本
 *

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
/**
 * label 类名
 */
func (a *ListBodyField) LabelClassName(value string) *ListBodyField {
    a.Set("labelClassName", value)
    return a
}

/**
 * 内层组件的CSS类名
 */
func (a *ListBodyField) InnerClassName(value string) *ListBodyField {
    a.Set("innerClassName", value)
    return a
}

/**
 * 绑定字段名
 */
func (a *ListBodyField) Name(value string) *ListBodyField {
    a.Set("name", value)
    return a
}

/**
 * 配置查看详情功能
 */
func (a *ListBodyField) PopOver(value string) *ListBodyField {
    a.Set("popOver", value)
    return a
}

/**
 * 配置快速编辑功能
 */
func (a *ListBodyField) QuickEdit(value string) *ListBodyField {
    a.Set("quickEdit", value)
    return a
}

/**
 * 配置点击复制功能
 */
func (a *ListBodyField) Copyable(value string) *ListBodyField {
    a.Set("copyable", value)
    return a
}

/**
 * 列标题
 */
func (a *ListBodyField) Label(value string) *ListBodyField {
    a.Set("label", value)
    return a
}
