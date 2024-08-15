package renderers


/**
 * MultilineText

 * @author wcz0
 * @version 6.2.2
 */
type MultilineText struct {
	*BaseRenderer
}

func NewMultilineText() *MultilineText {
    a := &MultilineText{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "multiline-text")
    return a
}


func (a *MultilineText) Set(name string, value interface{}) *MultilineText {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 静态展示表单项类名
 */
func (a *MultilineText) Staticclassname(value interface{}) *MultilineText {
    a.Set("staticClassName", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *MultilineText) Staticon(value interface{}) *MultilineText {
    a.Set("staticOn", value)
    return a
}

/**
 */
func (a *MultilineText) Testid(value interface{}) *MultilineText {
    a.Set("testid", value)
    return a
}

/**
 * 收起按钮文本
 */
func (a *MultilineText) Collapsebuttontext(value interface{}) *MultilineText {
    a.Set("collapseButtonText", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *MultilineText) Classname(value interface{}) *MultilineText {
    a.Set("className", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *MultilineText) Visibleon(value interface{}) *MultilineText {
    a.Set("visibleOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *MultilineText) Id(value interface{}) *MultilineText {
    a.Set("id", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *MultilineText) Onevent(value interface{}) *MultilineText {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *MultilineText) Staticlabelclassname(value interface{}) *MultilineText {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 最大行数
 */
func (a *MultilineText) Maxrows(value interface{}) *MultilineText {
    a.Set("maxRows", value)
    return a
}

/**
 * 展开按钮文本
 */
func (a *MultilineText) Expendbuttontext(value interface{}) *MultilineText {
    a.Set("expendButtonText", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *MultilineText) Usemobileui(value interface{}) *MultilineText {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 文本
 */
func (a *MultilineText) Text(value interface{}) *MultilineText {
    a.Set("text", value)
    return a
}

/**
 * 是否禁用
 */
func (a *MultilineText) Disabled(value interface{}) *MultilineText {
    a.Set("disabled", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *MultilineText) Disabledon(value interface{}) *MultilineText {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *MultilineText) Static(value interface{}) *MultilineText {
    a.Set("static", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *MultilineText) Staticplaceholder(value interface{}) *MultilineText {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 */
func (a *MultilineText) Staticschema(value interface{}) *MultilineText {
    a.Set("staticSchema", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *MultilineText) Hidden(value interface{}) *MultilineText {
    a.Set("hidden", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *MultilineText) Hiddenon(value interface{}) *MultilineText {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *MultilineText) Visible(value interface{}) *MultilineText {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *MultilineText) Staticinputclassname(value interface{}) *MultilineText {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *MultilineText) Style(value interface{}) *MultilineText {
    a.Set("style", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *MultilineText) Editorsetting(value interface{}) *MultilineText {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *MultilineText) Testidbuilder(value interface{}) *MultilineText {
    a.Set("testIdBuilder", value)
    return a
}

/**
 */
func (a *MultilineText) Type(value interface{}) *MultilineText {
    a.Set("type", value)
    return a
}
