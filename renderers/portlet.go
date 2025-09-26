package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type Portlet struct {
	*BaseRenderer
}

func NewPortlet() *Portlet {
    a := &Portlet{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "portlet")
    return a
}


func (a *Portlet) Set(name string, value interface{}) *Portlet {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * id
 */
func (a *Portlet) Id(value interface{}) *Portlet {
    a.Set("id", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Portlet) StaticInputClassName(value interface{}) *Portlet {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *Portlet) Type(value interface{}) *Portlet {
    a.Set("type", value)
    return a
}

/**
 * linksClassName
 */
func (a *Portlet) LinksClassName(value interface{}) *Portlet {
    a.Set("linksClassName", value)
    return a
}

/**
 * unmountOnExit
 */
func (a *Portlet) UnmountOnExit(value interface{}) *Portlet {
    a.Set("unmountOnExit", value)
    return a
}

/**
 * scrollable
 */
func (a *Portlet) Scrollable(value interface{}) *Portlet {
    a.Set("scrollable", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Portlet) HiddenOn(value interface{}) *Portlet {
    a.Set("hiddenOn", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Portlet) UseMobileUI(value interface{}) *Portlet {
    a.Set("useMobileUI", value)
    return a
}

/**
 * visible
 */
func (a *Portlet) Visible(value interface{}) *Portlet {
    a.Set("visible", value)
    return a
}

/**
 * visibleOn
 */
func (a *Portlet) VisibleOn(value interface{}) *Portlet {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Portlet) StaticPlaceholder(value interface{}) *Portlet {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticClassName
 */
func (a *Portlet) StaticClassName(value interface{}) *Portlet {
    a.Set("staticClassName", value)
    return a
}

/**
 * contentClassName
 */
func (a *Portlet) ContentClassName(value interface{}) *Portlet {
    a.Set("contentClassName", value)
    return a
}

/**
 * disabled
 */
func (a *Portlet) Disabled(value interface{}) *Portlet {
    a.Set("disabled", value)
    return a
}

/**
 * hidden
 */
func (a *Portlet) Hidden(value interface{}) *Portlet {
    a.Set("hidden", value)
    return a
}

/**
 * source
 */
func (a *Portlet) Source(value interface{}) *Portlet {
    a.Set("source", value)
    return a
}

/**
 * tabsClassName
 */
func (a *Portlet) TabsClassName(value interface{}) *Portlet {
    a.Set("tabsClassName", value)
    return a
}

/**
 * disabledOn
 */
func (a *Portlet) DisabledOn(value interface{}) *Portlet {
    a.Set("disabledOn", value)
    return a
}

/**
 * editorSetting
 */
func (a *Portlet) EditorSetting(value interface{}) *Portlet {
    a.Set("editorSetting", value)
    return a
}

/**
 * onEvent
 */
func (a *Portlet) OnEvent(value interface{}) *Portlet {
    a.Set("onEvent", value)
    return a
}

/**
 * mountOnEnter
 */
func (a *Portlet) MountOnEnter(value interface{}) *Portlet {
    a.Set("mountOnEnter", value)
    return a
}

/**
 * description
 */
func (a *Portlet) Description(value interface{}) *Portlet {
    a.Set("description", value)
    return a
}

/**
 * tabsMode
 */
func (a *Portlet) TabsMode(value interface{}) *Portlet {
    a.Set("tabsMode", value)
    return a
}

/**
 * staticOn
 */
func (a *Portlet) StaticOn(value interface{}) *Portlet {
    a.Set("staticOn", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Portlet) StaticLabelClassName(value interface{}) *Portlet {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *Portlet) StaticSchema(value interface{}) *Portlet {
    a.Set("staticSchema", value)
    return a
}

/**
 * style
 */
func (a *Portlet) Style(value interface{}) *Portlet {
    a.Set("style", value)
    return a
}

/**
 * tabs
 */
func (a *Portlet) Tabs(value interface{}) *Portlet {
    a.Set("tabs", value)
    return a
}

/**
 * className
 */
func (a *Portlet) ClassName(value interface{}) *Portlet {
    a.Set("className", value)
    return a
}

/**
 * static
 */
func (a *Portlet) Static(value interface{}) *Portlet {
    a.Set("static", value)
    return a
}

/**
 * toolbar
 */
func (a *Portlet) Toolbar(value interface{}) *Portlet {
    a.Set("toolbar", value)
    return a
}

/**
 * divider
 */
func (a *Portlet) Divider(value interface{}) *Portlet {
    a.Set("divider", value)
    return a
}

/**
 * hideHeader
 */
func (a *Portlet) HideHeader(value interface{}) *Portlet {
    a.Set("hideHeader", value)
    return a
}
