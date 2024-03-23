package renderers


/**

*/
type SchemaCopyable struct {
	*BaseRenderer
}

func NewSchemaCopyable() *SchemaCopyable {
    a := &SchemaCopyable{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}
/**
 * 可以配置图标
 */
func (a *SchemaCopyable) Icon(value string) *SchemaCopyable {
    a.Set("icon", value)
    return a
}

/**
 * 配置复制时的内容模板。
 */
func (a *SchemaCopyable) Content(value string) *SchemaCopyable {
    a.Set("content", value)
    return a
}

/**
 * 提示文字内容
 */
func (a *SchemaCopyable) Tooltip(value string) *SchemaCopyable {
    a.Set("tooltip", value)
    return a
}
