package renderers


/**
 * 操作栏渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/operation

 * @author wcz0
 * @version 6.2.2
 */
type Operation struct {
	*BaseRenderer
}

func NewOperation() *Operation {
    a := &Operation{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "operation")
    return a
}


func (a *Operation) Set(name string, value interface{}) *Operation {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * buttons
 */
func (a *Operation) Buttons(value interface{}) *Operation {
    a.Set("buttons", value)
    return a
}

/**
 * 设置label
 */
func (a *Operation) Label(value interface{}) *Operation {
    a.Set("label", value)
    return a
}

/**
 * disabledOn
 */
func (a *Operation) DisabledOn(value interface{}) *Operation {
    a.Set("disabledOn", value)
    return a
}

/**
 * visibleOn
 */
func (a *Operation) VisibleOn(value interface{}) *Operation {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticOn
 */
func (a *Operation) StaticOn(value interface{}) *Operation {
    a.Set("staticOn", value)
    return a
}

/**
 * placeholder
 */
func (a *Operation) Placeholder(value interface{}) *Operation {
    a.Set("placeholder", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Operation) HiddenOn(value interface{}) *Operation {
    a.Set("hiddenOn", value)
    return a
}

/**
 * visible
 */
func (a *Operation) Visible(value interface{}) *Operation {
    a.Set("visible", value)
    return a
}

/**
 * static
 */
func (a *Operation) Static(value interface{}) *Operation {
    a.Set("static", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Operation) StaticLabelClassName(value interface{}) *Operation {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * editorSetting
 */
func (a *Operation) EditorSetting(value interface{}) *Operation {
    a.Set("editorSetting", value)
    return a
}

/**
 * className
 */
func (a *Operation) ClassName(value interface{}) *Operation {
    a.Set("className", value)
    return a
}

/**
 * disabled
 */
func (a *Operation) Disabled(value interface{}) *Operation {
    a.Set("disabled", value)
    return a
}

/**
 * hidden
 */
func (a *Operation) Hidden(value interface{}) *Operation {
    a.Set("hidden", value)
    return a
}

/**
 * onEvent
 */
func (a *Operation) OnEvent(value interface{}) *Operation {
    a.Set("onEvent", value)
    return a
}

/**
 * staticClassName
 */
func (a *Operation) StaticClassName(value interface{}) *Operation {
    a.Set("staticClassName", value)
    return a
}

/**
 * style
 */
func (a *Operation) Style(value interface{}) *Operation {
    a.Set("style", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Operation) UseMobileUI(value interface{}) *Operation {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *Operation) Type(value interface{}) *Operation {
    a.Set("type", value)
    return a
}

/**
 * id
 */
func (a *Operation) Id(value interface{}) *Operation {
    a.Set("id", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Operation) StaticPlaceholder(value interface{}) *Operation {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Operation) StaticInputClassName(value interface{}) *Operation {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *Operation) StaticSchema(value interface{}) *Operation {
    a.Set("staticSchema", value)
    return a
}

/**
 * testid
 */
func (a *Operation) Testid(value interface{}) *Operation {
    a.Set("testid", value)
    return a
}
