package renderers


/**
 * JSON 数据展示控件。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/json

 * @author wcz0
 * @version 6.2.2
 */
type Json struct {
	*BaseRenderer
}

func NewJson() *Json {
    a := &Json{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "json")
    return a
}


func (a *Json) Set(name string, value interface{}) *Json {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * staticPlaceholder
 */
func (a *Json) StaticPlaceholder(value interface{}) *Json {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * editorSetting
 */
func (a *Json) EditorSetting(value interface{}) *Json {
    a.Set("editorSetting", value)
    return a
}

/**
 * displayDataTypes
 */
func (a *Json) DisplayDataTypes(value interface{}) *Json {
    a.Set("displayDataTypes", value)
    return a
}

/**
 * sortKeys
 */
func (a *Json) SortKeys(value interface{}) *Json {
    a.Set("sortKeys", value)
    return a
}

/**
 * mutable
 */
func (a *Json) Mutable(value interface{}) *Json {
    a.Set("mutable", value)
    return a
}

/**
 * disabledOn
 */
func (a *Json) DisabledOn(value interface{}) *Json {
    a.Set("disabledOn", value)
    return a
}

/**
 * visible
 */
func (a *Json) Visible(value interface{}) *Json {
    a.Set("visible", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Json) StaticInputClassName(value interface{}) *Json {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * enableClipboard
 */
func (a *Json) EnableClipboard(value interface{}) *Json {
    a.Set("enableClipboard", value)
    return a
}

/**
 * iconStyle
 */
func (a *Json) IconStyle(value interface{}) *Json {
    a.Set("iconStyle", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Json) HiddenOn(value interface{}) *Json {
    a.Set("hiddenOn", value)
    return a
}

/**
 * static
 */
func (a *Json) Static(value interface{}) *Json {
    a.Set("static", value)
    return a
}

/**
 * staticOn
 */
func (a *Json) StaticOn(value interface{}) *Json {
    a.Set("staticOn", value)
    return a
}

/**
 * style
 */
func (a *Json) Style(value interface{}) *Json {
    a.Set("style", value)
    return a
}

/**
 * quotesOnKeys
 */
func (a *Json) QuotesOnKeys(value interface{}) *Json {
    a.Set("quotesOnKeys", value)
    return a
}

/**
 * className
 */
func (a *Json) ClassName(value interface{}) *Json {
    a.Set("className", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Json) UseMobileUI(value interface{}) *Json {
    a.Set("useMobileUI", value)
    return a
}

/**
 * testid
 */
func (a *Json) Testid(value interface{}) *Json {
    a.Set("testid", value)
    return a
}

/**
 * disabled
 */
func (a *Json) Disabled(value interface{}) *Json {
    a.Set("disabled", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Json) StaticLabelClassName(value interface{}) *Json {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *Json) StaticSchema(value interface{}) *Json {
    a.Set("staticSchema", value)
    return a
}

/**
 * source
 */
func (a *Json) Source(value interface{}) *Json {
    a.Set("source", value)
    return a
}

/**
 * hidden
 */
func (a *Json) Hidden(value interface{}) *Json {
    a.Set("hidden", value)
    return a
}

/**
 * id
 */
func (a *Json) Id(value interface{}) *Json {
    a.Set("id", value)
    return a
}

/**
 * ellipsisThreshold
 */
func (a *Json) EllipsisThreshold(value interface{}) *Json {
    a.Set("ellipsisThreshold", value)
    return a
}

/**
 * visibleOn
 */
func (a *Json) VisibleOn(value interface{}) *Json {
    a.Set("visibleOn", value)
    return a
}

/**
 * onEvent
 */
func (a *Json) OnEvent(value interface{}) *Json {
    a.Set("onEvent", value)
    return a
}

/**
 * staticClassName
 */
func (a *Json) StaticClassName(value interface{}) *Json {
    a.Set("staticClassName", value)
    return a
}

/**
 * 可选值: json | static-json
 */
func (a *Json) Type(value interface{}) *Json {
    a.Set("type", value)
    return a
}

/**
 * value
 */
func (a *Json) Value(value interface{}) *Json {
    a.Set("value", value)
    return a
}

/**
 * levelExpand
 */
func (a *Json) LevelExpand(value interface{}) *Json {
    a.Set("levelExpand", value)
    return a
}
