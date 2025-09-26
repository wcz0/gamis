package renderers


/**
 * Plain 纯文本渲染器 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/plain

 * @author wcz0
 * @version 6.2.2
 */
type Plain struct {
	*BaseRenderer
}

func NewPlain() *Plain {
    a := &Plain{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "plain")
    return a
}


func (a *Plain) Set(name string, value interface{}) *Plain {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * inline
 */
func (a *Plain) Inline(value interface{}) *Plain {
    a.Set("inline", value)
    return a
}

/**
 * className
 */
func (a *Plain) ClassName(value interface{}) *Plain {
    a.Set("className", value)
    return a
}

/**
 * hidden
 */
func (a *Plain) Hidden(value interface{}) *Plain {
    a.Set("hidden", value)
    return a
}

/**
 * testid
 */
func (a *Plain) Testid(value interface{}) *Plain {
    a.Set("testid", value)
    return a
}

/**
 * text
 */
func (a *Plain) Text(value interface{}) *Plain {
    a.Set("text", value)
    return a
}

/**
 * placeholder
 */
func (a *Plain) Placeholder(value interface{}) *Plain {
    a.Set("placeholder", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Plain) StaticPlaceholder(value interface{}) *Plain {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * disabled
 */
func (a *Plain) Disabled(value interface{}) *Plain {
    a.Set("disabled", value)
    return a
}

/**
 * id
 */
func (a *Plain) Id(value interface{}) *Plain {
    a.Set("id", value)
    return a
}

/**
 * onEvent
 */
func (a *Plain) OnEvent(value interface{}) *Plain {
    a.Set("onEvent", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Plain) StaticInputClassName(value interface{}) *Plain {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * editorSetting
 */
func (a *Plain) EditorSetting(value interface{}) *Plain {
    a.Set("editorSetting", value)
    return a
}

/**
 * tpl
 */
func (a *Plain) Tpl(value interface{}) *Plain {
    a.Set("tpl", value)
    return a
}

/**
 * disabledOn
 */
func (a *Plain) DisabledOn(value interface{}) *Plain {
    a.Set("disabledOn", value)
    return a
}

/**
 * static
 */
func (a *Plain) Static(value interface{}) *Plain {
    a.Set("static", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Plain) StaticLabelClassName(value interface{}) *Plain {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *Plain) StaticSchema(value interface{}) *Plain {
    a.Set("staticSchema", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Plain) UseMobileUI(value interface{}) *Plain {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 可选值: plain | text
 */
func (a *Plain) Type(value interface{}) *Plain {
    a.Set("type", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Plain) HiddenOn(value interface{}) *Plain {
    a.Set("hiddenOn", value)
    return a
}

/**
 * visible
 */
func (a *Plain) Visible(value interface{}) *Plain {
    a.Set("visible", value)
    return a
}

/**
 * visibleOn
 */
func (a *Plain) VisibleOn(value interface{}) *Plain {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticOn
 */
func (a *Plain) StaticOn(value interface{}) *Plain {
    a.Set("staticOn", value)
    return a
}

/**
 * staticClassName
 */
func (a *Plain) StaticClassName(value interface{}) *Plain {
    a.Set("staticClassName", value)
    return a
}

/**
 * style
 */
func (a *Plain) Style(value interface{}) *Plain {
    a.Set("style", value)
    return a
}
