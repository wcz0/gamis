package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type NavItem struct {
	*BaseRenderer
}

func NewNavItem() *NavItem {
    a := &NavItem{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}


func (a *NavItem) Set(name string, value interface{}) *NavItem {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * hidden
 */
func (a *NavItem) Hidden(value interface{}) *NavItem {
    a.Set("hidden", value)
    return a
}

/**
 * staticClassName
 */
func (a *NavItem) StaticClassName(value interface{}) *NavItem {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *NavItem) StaticSchema(value interface{}) *NavItem {
    a.Set("staticSchema", value)
    return a
}

/**
 * className
 */
func (a *NavItem) ClassName(value interface{}) *NavItem {
    a.Set("className", value)
    return a
}

/**
 * onEvent
 */
func (a *NavItem) OnEvent(value interface{}) *NavItem {
    a.Set("onEvent", value)
    return a
}

/**
 * defer
 */
func (a *NavItem) Defer(value interface{}) *NavItem {
    a.Set("defer", value)
    return a
}

/**
 * visible
 */
func (a *NavItem) Visible(value interface{}) *NavItem {
    a.Set("visible", value)
    return a
}

/**
 * static
 */
func (a *NavItem) Static(value interface{}) *NavItem {
    a.Set("static", value)
    return a
}

/**
 * staticOn
 */
func (a *NavItem) StaticOn(value interface{}) *NavItem {
    a.Set("staticOn", value)
    return a
}

/**
 * active
 */
func (a *NavItem) Active(value interface{}) *NavItem {
    a.Set("active", value)
    return a
}

/**
 * deferApi
 */
func (a *NavItem) DeferApi(value interface{}) *NavItem {
    a.Set("deferApi", value)
    return a
}

/**
 * children
 */
func (a *NavItem) Children(value interface{}) *NavItem {
    a.Set("children", value)
    return a
}

/**
 * disabledOn
 */
func (a *NavItem) DisabledOn(value interface{}) *NavItem {
    a.Set("disabledOn", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *NavItem) StaticLabelClassName(value interface{}) *NavItem {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * disabled
 */
func (a *NavItem) Disabled(value interface{}) *NavItem {
    a.Set("disabled", value)
    return a
}

/**
 * visibleOn
 */
func (a *NavItem) VisibleOn(value interface{}) *NavItem {
    a.Set("visibleOn", value)
    return a
}

/**
 * id
 */
func (a *NavItem) Id(value interface{}) *NavItem {
    a.Set("id", value)
    return a
}

/**
 * editorSetting
 */
func (a *NavItem) EditorSetting(value interface{}) *NavItem {
    a.Set("editorSetting", value)
    return a
}

/**
 * target
 */
func (a *NavItem) Target(value interface{}) *NavItem {
    a.Set("target", value)
    return a
}

/**
 * unfolded
 */
func (a *NavItem) Unfolded(value interface{}) *NavItem {
    a.Set("unfolded", value)
    return a
}

/**
 * key
 */
func (a *NavItem) Key(value interface{}) *NavItem {
    a.Set("key", value)
    return a
}

/**
 * mode
 */
func (a *NavItem) Mode(value interface{}) *NavItem {
    a.Set("mode", value)
    return a
}

/**
 * hiddenOn
 */
func (a *NavItem) HiddenOn(value interface{}) *NavItem {
    a.Set("hiddenOn", value)
    return a
}

/**
 * useMobileUI
 */
func (a *NavItem) UseMobileUI(value interface{}) *NavItem {
    a.Set("useMobileUI", value)
    return a
}

/**
 * label
 */
func (a *NavItem) Label(value interface{}) *NavItem {
    a.Set("label", value)
    return a
}

/**
 * disabledTip
 */
func (a *NavItem) DisabledTip(value interface{}) *NavItem {
    a.Set("disabledTip", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *NavItem) StaticPlaceholder(value interface{}) *NavItem {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *NavItem) StaticInputClassName(value interface{}) *NavItem {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * style
 */
func (a *NavItem) Style(value interface{}) *NavItem {
    a.Set("style", value)
    return a
}

/**
 * icon
 */
func (a *NavItem) Icon(value interface{}) *NavItem {
    a.Set("icon", value)
    return a
}

/**
 * to
 */
func (a *NavItem) To(value interface{}) *NavItem {
    a.Set("to", value)
    return a
}
