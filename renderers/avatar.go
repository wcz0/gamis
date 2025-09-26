package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type Avatar struct {
	*BaseRenderer
}

func NewAvatar() *Avatar {
    a := &Avatar{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "avatar")
    return a
}


func (a *Avatar) Set(name string, value interface{}) *Avatar {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * disabledOn
 */
func (a *Avatar) DisabledOn(value interface{}) *Avatar {
    a.Set("disabledOn", value)
    return a
}

/**
 * shape
 */
func (a *Avatar) Shape(value interface{}) *Avatar {
    a.Set("shape", value)
    return a
}

/**
 * id
 */
func (a *Avatar) Id(value interface{}) *Avatar {
    a.Set("id", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Avatar) StaticLabelClassName(value interface{}) *Avatar {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Avatar) StaticInputClassName(value interface{}) *Avatar {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * size
 */
func (a *Avatar) Size(value interface{}) *Avatar {
    a.Set("size", value)
    return a
}

/**
 * gap
 */
func (a *Avatar) Gap(value interface{}) *Avatar {
    a.Set("gap", value)
    return a
}

/**
 * className
 */
func (a *Avatar) ClassName(value interface{}) *Avatar {
    a.Set("className", value)
    return a
}

/**
 * hidden
 */
func (a *Avatar) Hidden(value interface{}) *Avatar {
    a.Set("hidden", value)
    return a
}

/**
 * visibleOn
 */
func (a *Avatar) VisibleOn(value interface{}) *Avatar {
    a.Set("visibleOn", value)
    return a
}

/**
 * src
 */
func (a *Avatar) Src(value interface{}) *Avatar {
    a.Set("src", value)
    return a
}

/**
 * fit
 */
func (a *Avatar) Fit(value interface{}) *Avatar {
    a.Set("fit", value)
    return a
}

/**
 * visible
 */
func (a *Avatar) Visible(value interface{}) *Avatar {
    a.Set("visible", value)
    return a
}

/**
 * onEvent
 */
func (a *Avatar) OnEvent(value interface{}) *Avatar {
    a.Set("onEvent", value)
    return a
}

/**
 * editorSetting
 */
func (a *Avatar) EditorSetting(value interface{}) *Avatar {
    a.Set("editorSetting", value)
    return a
}

/**
 * badge
 */
func (a *Avatar) Badge(value interface{}) *Avatar {
    a.Set("badge", value)
    return a
}

/**
 * draggable
 */
func (a *Avatar) Draggable(value interface{}) *Avatar {
    a.Set("draggable", value)
    return a
}

/**
 * onError
 */
func (a *Avatar) OnError(value interface{}) *Avatar {
    a.Set("onError", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Avatar) HiddenOn(value interface{}) *Avatar {
    a.Set("hiddenOn", value)
    return a
}

/**
 * staticClassName
 */
func (a *Avatar) StaticClassName(value interface{}) *Avatar {
    a.Set("staticClassName", value)
    return a
}

/**
 * style
 */
func (a *Avatar) Style(value interface{}) *Avatar {
    a.Set("style", value)
    return a
}

/**
 * testid
 */
func (a *Avatar) Testid(value interface{}) *Avatar {
    a.Set("testid", value)
    return a
}

/**
 * alt
 */
func (a *Avatar) Alt(value interface{}) *Avatar {
    a.Set("alt", value)
    return a
}

/**
 * crossOrigin
 */
func (a *Avatar) CrossOrigin(value interface{}) *Avatar {
    a.Set("crossOrigin", value)
    return a
}

/**
 * disabled
 */
func (a *Avatar) Disabled(value interface{}) *Avatar {
    a.Set("disabled", value)
    return a
}

/**
 * staticOn
 */
func (a *Avatar) StaticOn(value interface{}) *Avatar {
    a.Set("staticOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Avatar) StaticPlaceholder(value interface{}) *Avatar {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * static
 */
func (a *Avatar) Static(value interface{}) *Avatar {
    a.Set("static", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Avatar) UseMobileUI(value interface{}) *Avatar {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *Avatar) Type(value interface{}) *Avatar {
    a.Set("type", value)
    return a
}

/**
 * icon
 */
func (a *Avatar) Icon(value interface{}) *Avatar {
    a.Set("icon", value)
    return a
}

/**
 * text
 */
func (a *Avatar) Text(value interface{}) *Avatar {
    a.Set("text", value)
    return a
}

/**
 * staticSchema
 */
func (a *Avatar) StaticSchema(value interface{}) *Avatar {
    a.Set("staticSchema", value)
    return a
}

/**
 * defaultAvatar
 */
func (a *Avatar) DefaultAvatar(value interface{}) *Avatar {
    a.Set("defaultAvatar", value)
    return a
}
