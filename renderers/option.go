package renderers


/**

*/
type Option struct {
	*BaseRenderer
}

func NewOption() *Option {
    a := &Option{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}
/**
 * 最好不要用！因为有 visible 就够了。
 */
func (a *Option) Hidden(value string) *Option {
    a.Set("hidden", value)
    return a
}

/**
 * 描述，部分控件支持
 */
func (a *Option) Description(value string) *Option {
    a.Set("description", value)
    return a
}

/**
 * 如果设置了，优先级更高，不设置走 source 接口加载。
 */
func (a *Option) DeferApi(value string) *Option {
    a.Set("deferApi", value)
    return a
}

/**
 * 用来显示的文字
 */
func (a *Option) Label(value string) *Option {
    a.Set("label", value)
    return a
}

/**
 * 请保证数值唯一，多个选项值一致会认为是同一个选项。
 */
func (a *Option) Value(value string) *Option {
    a.Set("value", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Option) Disabled(value string) *Option {
    a.Set("disabled", value)
    return a
}

/**
 * 支持嵌套
 */
func (a *Option) Children(value string) *Option {
    a.Set("children", value)
    return a
}

/**
 * 是否可见
 */
func (a *Option) Visible(value string) *Option {
    a.Set("visible", value)
    return a
}

/**
 * 只有设置了 defer 才有意义，内部字段不可以外部设置
 */
func (a *Option) Loaded(value string) *Option {
    a.Set("loaded", value)
    return a
}

/**
 * 可以用来给 Option 标记个范围，让数据展示更清晰。这个只有在数值展示的时候显示。
 */
func (a *Option) ScopeLabel(value string) *Option {
    a.Set("scopeLabel", value)
    return a
}

/**
 * 标记后数据延时加载
 */
func (a *Option) Defer(value string) *Option {
    a.Set("defer", value)
    return a
}

/**
 * 标记正在加载。只有 defer 为 true 时有意义。内部字段不可以外部设置
 */
func (a *Option) Loading(value string) *Option {
    a.Set("loading", value)
    return a
}
