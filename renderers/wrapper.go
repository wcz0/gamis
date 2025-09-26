package renderers


/**
 * Wrapper 容器渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/wrapper

 * @author wcz0
 * @version 6.2.2
 */
type Wrapper struct {
	*BaseRenderer
}

func NewWrapper() *Wrapper {
    a := &Wrapper{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "wrapper")
    return a
}


func (a *Wrapper) Set(name string, value interface{}) *Wrapper {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * staticLabelClassName
 */
func (a *Wrapper) StaticLabelClassName(value interface{}) *Wrapper {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * editorSetting
 */
func (a *Wrapper) EditorSetting(value interface{}) *Wrapper {
    a.Set("editorSetting", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Wrapper) UseMobileUI(value interface{}) *Wrapper {
    a.Set("useMobileUI", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Wrapper) HiddenOn(value interface{}) *Wrapper {
    a.Set("hiddenOn", value)
    return a
}

/**
 * onEvent
 */
func (a *Wrapper) OnEvent(value interface{}) *Wrapper {
    a.Set("onEvent", value)
    return a
}

/**
 * static
 */
func (a *Wrapper) Static(value interface{}) *Wrapper {
    a.Set("static", value)
    return a
}

/**
 */
func (a *Wrapper) Type(value interface{}) *Wrapper {
    a.Set("type", value)
    return a
}

/**
 * body
 */
func (a *Wrapper) Body(value interface{}) *Wrapper {
    a.Set("body", value)
    return a
}

/**
 * size
 */
func (a *Wrapper) Size(value interface{}) *Wrapper {
    a.Set("size", value)
    return a
}

/**
 * className
 */
func (a *Wrapper) ClassName(value interface{}) *Wrapper {
    a.Set("className", value)
    return a
}

/**
 * disabled
 */
func (a *Wrapper) Disabled(value interface{}) *Wrapper {
    a.Set("disabled", value)
    return a
}

/**
 * disabledOn
 */
func (a *Wrapper) DisabledOn(value interface{}) *Wrapper {
    a.Set("disabledOn", value)
    return a
}

/**
 * visible
 */
func (a *Wrapper) Visible(value interface{}) *Wrapper {
    a.Set("visible", value)
    return a
}

/**
 * id
 */
func (a *Wrapper) Id(value interface{}) *Wrapper {
    a.Set("id", value)
    return a
}

/**
 * style
 */
func (a *Wrapper) Style(value interface{}) *Wrapper {
    a.Set("style", value)
    return a
}

/**
 * testid
 */
func (a *Wrapper) Testid(value interface{}) *Wrapper {
    a.Set("testid", value)
    return a
}

/**
 * staticClassName
 */
func (a *Wrapper) StaticClassName(value interface{}) *Wrapper {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Wrapper) StaticInputClassName(value interface{}) *Wrapper {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *Wrapper) StaticSchema(value interface{}) *Wrapper {
    a.Set("staticSchema", value)
    return a
}

/**
 * wrap
 */
func (a *Wrapper) Wrap(value interface{}) *Wrapper {
    a.Set("wrap", value)
    return a
}

/**
 * hidden
 */
func (a *Wrapper) Hidden(value interface{}) *Wrapper {
    a.Set("hidden", value)
    return a
}

/**
 * visibleOn
 */
func (a *Wrapper) VisibleOn(value interface{}) *Wrapper {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticOn
 */
func (a *Wrapper) StaticOn(value interface{}) *Wrapper {
    a.Set("staticOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Wrapper) StaticPlaceholder(value interface{}) *Wrapper {
    a.Set("staticPlaceholder", value)
    return a
}
