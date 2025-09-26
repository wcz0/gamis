package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type Tab struct {
	*BaseRenderer
}

func NewTab() *Tab {
    a := &Tab{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}


func (a *Tab) Set(name string, value interface{}) *Tab {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * badge
 */
func (a *Tab) Badge(value interface{}) *Tab {
    a.Set("badge", value)
    return a
}

/**
 * hash
 */
func (a *Tab) Hash(value interface{}) *Tab {
    a.Set("hash", value)
    return a
}

/**
 * unmountOnExit
 */
func (a *Tab) UnmountOnExit(value interface{}) *Tab {
    a.Set("unmountOnExit", value)
    return a
}

/**
 * title
 */
func (a *Tab) Title(value interface{}) *Tab {
    a.Set("title", value)
    return a
}

/**
 * horizontal
 */
func (a *Tab) Horizontal(value interface{}) *Tab {
    a.Set("horizontal", value)
    return a
}

/**
 * disabled
 */
func (a *Tab) Disabled(value interface{}) *Tab {
    a.Set("disabled", value)
    return a
}

/**
 * id
 */
func (a *Tab) Id(value interface{}) *Tab {
    a.Set("id", value)
    return a
}

/**
 * staticClassName
 */
func (a *Tab) StaticClassName(value interface{}) *Tab {
    a.Set("staticClassName", value)
    return a
}

/**
 * style
 */
func (a *Tab) Style(value interface{}) *Tab {
    a.Set("style", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Tab) UseMobileUI(value interface{}) *Tab {
    a.Set("useMobileUI", value)
    return a
}

/**
 * tab
 */
func (a *Tab) Tab(value interface{}) *Tab {
    a.Set("tab", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Tab) HiddenOn(value interface{}) *Tab {
    a.Set("hiddenOn", value)
    return a
}

/**
 * staticSchema
 */
func (a *Tab) StaticSchema(value interface{}) *Tab {
    a.Set("staticSchema", value)
    return a
}

/**
 * disabledOn
 */
func (a *Tab) DisabledOn(value interface{}) *Tab {
    a.Set("disabledOn", value)
    return a
}

/**
 * visibleOn
 */
func (a *Tab) VisibleOn(value interface{}) *Tab {
    a.Set("visibleOn", value)
    return a
}

/**
 * body
 */
func (a *Tab) Body(value interface{}) *Tab {
    a.Set("body", value)
    return a
}

/**
 * reload
 */
func (a *Tab) Reload(value interface{}) *Tab {
    a.Set("reload", value)
    return a
}

/**
 * mode
 */
func (a *Tab) Mode(value interface{}) *Tab {
    a.Set("mode", value)
    return a
}

/**
 * iconPosition
 */
func (a *Tab) IconPosition(value interface{}) *Tab {
    a.Set("iconPosition", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Tab) StaticPlaceholder(value interface{}) *Tab {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Tab) StaticLabelClassName(value interface{}) *Tab {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * icon
 */
func (a *Tab) Icon(value interface{}) *Tab {
    a.Set("icon", value)
    return a
}

/**
 * className
 */
func (a *Tab) ClassName(value interface{}) *Tab {
    a.Set("className", value)
    return a
}

/**
 * onEvent
 */
func (a *Tab) OnEvent(value interface{}) *Tab {
    a.Set("onEvent", value)
    return a
}

/**
 * static
 */
func (a *Tab) Static(value interface{}) *Tab {
    a.Set("static", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Tab) StaticInputClassName(value interface{}) *Tab {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * mountOnEnter
 */
func (a *Tab) MountOnEnter(value interface{}) *Tab {
    a.Set("mountOnEnter", value)
    return a
}

/**
 * closable
 */
func (a *Tab) Closable(value interface{}) *Tab {
    a.Set("closable", value)
    return a
}

/**
 * hidden
 */
func (a *Tab) Hidden(value interface{}) *Tab {
    a.Set("hidden", value)
    return a
}

/**
 * visible
 */
func (a *Tab) Visible(value interface{}) *Tab {
    a.Set("visible", value)
    return a
}

/**
 * staticOn
 */
func (a *Tab) StaticOn(value interface{}) *Tab {
    a.Set("staticOn", value)
    return a
}

/**
 * editorSetting
 */
func (a *Tab) EditorSetting(value interface{}) *Tab {
    a.Set("editorSetting", value)
    return a
}
