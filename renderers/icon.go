package renderers


/**
 * Icon 图标渲染器 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/icon

 * @author wcz0
 * @version 6.2.2
 */
type Icon struct {
	*BaseRenderer
}

func NewIcon() *Icon {
    a := &Icon{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "icon")
    return a
}


func (a *Icon) Set(name string, value interface{}) *Icon {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * staticSchema
 */
func (a *Icon) StaticSchema(value interface{}) *Icon {
    a.Set("staticSchema", value)
    return a
}

/**
 * editorSetting
 */
func (a *Icon) EditorSetting(value interface{}) *Icon {
    a.Set("editorSetting", value)
    return a
}

/**
 * badge
 */
func (a *Icon) Badge(value interface{}) *Icon {
    a.Set("badge", value)
    return a
}

/**
 * visibleOn
 */
func (a *Icon) VisibleOn(value interface{}) *Icon {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticOn
 */
func (a *Icon) StaticOn(value interface{}) *Icon {
    a.Set("staticOn", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Icon) StaticInputClassName(value interface{}) *Icon {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Icon) UseMobileUI(value interface{}) *Icon {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *Icon) Type(value interface{}) *Icon {
    a.Set("type", value)
    return a
}

/**
 * testid
 */
func (a *Icon) Testid(value interface{}) *Icon {
    a.Set("testid", value)
    return a
}

/**
 * icon
 */
func (a *Icon) Icon(value interface{}) *Icon {
    a.Set("icon", value)
    return a
}

/**
 * id
 */
func (a *Icon) Id(value interface{}) *Icon {
    a.Set("id", value)
    return a
}

/**
 * onEvent
 */
func (a *Icon) OnEvent(value interface{}) *Icon {
    a.Set("onEvent", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Icon) StaticLabelClassName(value interface{}) *Icon {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * style
 */
func (a *Icon) Style(value interface{}) *Icon {
    a.Set("style", value)
    return a
}

/**
 * vendor
 */
func (a *Icon) Vendor(value interface{}) *Icon {
    a.Set("vendor", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Icon) StaticPlaceholder(value interface{}) *Icon {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * disabled
 */
func (a *Icon) Disabled(value interface{}) *Icon {
    a.Set("disabled", value)
    return a
}

/**
 * disabledOn
 */
func (a *Icon) DisabledOn(value interface{}) *Icon {
    a.Set("disabledOn", value)
    return a
}

/**
 * hidden
 */
func (a *Icon) Hidden(value interface{}) *Icon {
    a.Set("hidden", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Icon) HiddenOn(value interface{}) *Icon {
    a.Set("hiddenOn", value)
    return a
}

/**
 * staticClassName
 */
func (a *Icon) StaticClassName(value interface{}) *Icon {
    a.Set("staticClassName", value)
    return a
}

/**
 * className
 */
func (a *Icon) ClassName(value interface{}) *Icon {
    a.Set("className", value)
    return a
}

/**
 * visible
 */
func (a *Icon) Visible(value interface{}) *Icon {
    a.Set("visible", value)
    return a
}

/**
 * static
 */
func (a *Icon) Static(value interface{}) *Icon {
    a.Set("static", value)
    return a
}
