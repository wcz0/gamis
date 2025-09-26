package renderers


/**

 * @author wcz0
 * @version 6.2.2
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


func (a *Option) Set(name string, value interface{}) *Option {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 如果设置了，优先级更高，不设置走 source 接口加载。
 */
func (a *Option) DeferApi(value interface{}) *Option {
    a.Set("deferApi", value)
    return a
}

/**
 * 标记正在加载。只有 defer 为 true 时有意义。内部字段不可以外部设置
 */
func (a *Option) Loading(value interface{}) *Option {
    a.Set("loading", value)
    return a
}

/**
 * 用来显示的文字
 */
func (a *Option) Label(value interface{}) *Option {
    a.Set("label", value)
    return a
}

/**
 * 可以用来给 Option 标记个范围，让数据展示更清晰。这个只有在数值展示的时候显示。
 */
func (a *Option) ScopeLabel(value interface{}) *Option {
    a.Set("scopeLabel", value)
    return a
}

/**
 * 禁用提示
 */
func (a *Option) DisabledTip(value interface{}) *Option {
    a.Set("disabledTip", value)
    return a
}

/**
 * 标记后数据延时加载
 */
func (a *Option) Defer(value interface{}) *Option {
    a.Set("defer", value)
    return a
}

/**
 * 只有设置了 defer 才有意义，内部字段不可以外部设置
 */
func (a *Option) Loaded(value interface{}) *Option {
    a.Set("loaded", value)
    return a
}

/**
 * 请保证数值唯一，多个选项值一致会认为是同一个选项。
 */
func (a *Option) Value(value interface{}) *Option {
    a.Set("value", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Option) Disabled(value interface{}) *Option {
    a.Set("disabled", value)
    return a
}

/**
 * 支持嵌套
 */
func (a *Option) Children(value interface{}) *Option {
    a.Set("children", value)
    return a
}

/**
 * 是否可见
 */
func (a *Option) Visible(value interface{}) *Option {
    a.Set("visible", value)
    return a
}

/**
 * 最好不要用！因为有 visible 就够了。
 */
func (a *Option) Hidden(value interface{}) *Option {
    a.Set("hidden", value)
    return a
}

/**
 * 描述，部分控件支持
 */
func (a *Option) Description(value interface{}) *Option {
    a.Set("description", value)
    return a
}
