package renderers


/**
 * Color 显示渲染器，格式说明。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/color

 * @author wcz0
 * @version 6.2.2
 */
type Color struct {
	*BaseRenderer
}

func NewColor() *Color {
    a := &Color{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "color")
    return a
}


func (a *Color) Set(name string, value interface{}) *Color {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 */
func (a *Color) StaticSchema(value interface{}) *Color {
    a.Set("staticSchema", value)
    return a
}

/**
 * 组件样式
 */
func (a *Color) Style(value interface{}) *Color {
    a.Set("style", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Color) EditorSetting(value interface{}) *Color {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *Color) TestIdBuilder(value interface{}) *Color {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 默认颜色
 */
func (a *Color) DefaultColor(value interface{}) *Color {
    a.Set("defaultColor", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Color) ClassName(value interface{}) *Color {
    a.Set("className", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Color) Static(value interface{}) *Color {
    a.Set("static", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Color) StaticOn(value interface{}) *Color {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Color) StaticClassName(value interface{}) *Color {
    a.Set("staticClassName", value)
    return a
}

/**
 * 指定为颜色显示控件
 */
func (a *Color) Type(value interface{}) *Color {
    a.Set("type", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Color) Disabled(value interface{}) *Color {
    a.Set("disabled", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Color) HiddenOn(value interface{}) *Color {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Color) OnEvent(value interface{}) *Color {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Color) StaticLabelClassName(value interface{}) *Color {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Color) StaticInputClassName(value interface{}) *Color {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Color) Hidden(value interface{}) *Color {
    a.Set("hidden", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Color) VisibleOn(value interface{}) *Color {
    a.Set("visibleOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Color) Id(value interface{}) *Color {
    a.Set("id", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Color) UseMobileUI(value interface{}) *Color {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *Color) Testid(value interface{}) *Color {
    a.Set("testid", value)
    return a
}

/**
 * 是否用文字显示值。
 */
func (a *Color) ShowValue(value interface{}) *Color {
    a.Set("showValue", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Color) DisabledOn(value interface{}) *Color {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *Color) Visible(value interface{}) *Color {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Color) StaticPlaceholder(value interface{}) *Color {
    a.Set("staticPlaceholder", value)
    return a
}
