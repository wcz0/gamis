package renderers


/**
 * 提示渲染器，默认会显示个小图标，鼠标放上来的时候显示配置的内容。

 * @author wcz0
 * @version 6.2.2
 */
type BaseRemark struct {
	*BaseRenderer
}

func NewBaseRemark() *BaseRemark {
    a := &BaseRemark{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}


func (a *BaseRemark) Set(name string, value interface{}) *BaseRemark {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * onEvent
 */
func (a *BaseRemark) OnEvent(value interface{}) *BaseRemark {
    a.Set("onEvent", value)
    return a
}

/**
 * testid
 */
func (a *BaseRemark) Testid(value interface{}) *BaseRemark {
    a.Set("testid", value)
    return a
}

/**
 * tooltipClassName
 */
func (a *BaseRemark) TooltipClassName(value interface{}) *BaseRemark {
    a.Set("tooltipClassName", value)
    return a
}

/**
 * disabledOn
 */
func (a *BaseRemark) DisabledOn(value interface{}) *BaseRemark {
    a.Set("disabledOn", value)
    return a
}

/**
 * hidden
 */
func (a *BaseRemark) Hidden(value interface{}) *BaseRemark {
    a.Set("hidden", value)
    return a
}

/**
 * visibleOn
 */
func (a *BaseRemark) VisibleOn(value interface{}) *BaseRemark {
    a.Set("visibleOn", value)
    return a
}

/**
 * editorSetting
 */
func (a *BaseRemark) EditorSetting(value interface{}) *BaseRemark {
    a.Set("editorSetting", value)
    return a
}

/**
 * disabled
 */
func (a *BaseRemark) Disabled(value interface{}) *BaseRemark {
    a.Set("disabled", value)
    return a
}

/**
 * visible
 */
func (a *BaseRemark) Visible(value interface{}) *BaseRemark {
    a.Set("visible", value)
    return a
}

/**
 * id
 */
func (a *BaseRemark) Id(value interface{}) *BaseRemark {
    a.Set("id", value)
    return a
}

/**
 * className
 */
func (a *BaseRemark) ClassName(value interface{}) *BaseRemark {
    a.Set("className", value)
    return a
}

/**
 * staticClassName
 */
func (a *BaseRemark) StaticClassName(value interface{}) *BaseRemark {
    a.Set("staticClassName", value)
    return a
}

/**
 * style
 */
func (a *BaseRemark) Style(value interface{}) *BaseRemark {
    a.Set("style", value)
    return a
}

/**
 * type
 */
func (a *BaseRemark) Type(value interface{}) *BaseRemark {
    a.Set("type", value)
    return a
}

/**
 * shape
 */
func (a *BaseRemark) Shape(value interface{}) *BaseRemark {
    a.Set("shape", value)
    return a
}

/**
 * useMobileUI
 */
func (a *BaseRemark) UseMobileUI(value interface{}) *BaseRemark {
    a.Set("useMobileUI", value)
    return a
}

/**
 * icon
 */
func (a *BaseRemark) Icon(value interface{}) *BaseRemark {
    a.Set("icon", value)
    return a
}

/**
 * title
 */
func (a *BaseRemark) Title(value interface{}) *BaseRemark {
    a.Set("title", value)
    return a
}

/**
 * placement
 */
func (a *BaseRemark) Placement(value interface{}) *BaseRemark {
    a.Set("placement", value)
    return a
}

/**
 * hiddenOn
 */
func (a *BaseRemark) HiddenOn(value interface{}) *BaseRemark {
    a.Set("hiddenOn", value)
    return a
}

/**
 * static
 */
func (a *BaseRemark) Static(value interface{}) *BaseRemark {
    a.Set("static", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *BaseRemark) StaticPlaceholder(value interface{}) *BaseRemark {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * content
 */
func (a *BaseRemark) Content(value interface{}) *BaseRemark {
    a.Set("content", value)
    return a
}

/**
 * rootClose
 */
func (a *BaseRemark) RootClose(value interface{}) *BaseRemark {
    a.Set("rootClose", value)
    return a
}

/**
 * staticOn
 */
func (a *BaseRemark) StaticOn(value interface{}) *BaseRemark {
    a.Set("staticOn", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *BaseRemark) StaticInputClassName(value interface{}) *BaseRemark {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *BaseRemark) StaticSchema(value interface{}) *BaseRemark {
    a.Set("staticSchema", value)
    return a
}

/**
 * label
 */
func (a *BaseRemark) Label(value interface{}) *BaseRemark {
    a.Set("label", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *BaseRemark) StaticLabelClassName(value interface{}) *BaseRemark {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * trigger
 */
func (a *BaseRemark) Trigger(value interface{}) *BaseRemark {
    a.Set("trigger", value)
    return a
}
