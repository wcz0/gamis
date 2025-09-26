package renderers


/**
 * Tag

 * @author wcz0
 * @version 6.2.2
 */
type Tag struct {
	*BaseRenderer
}

func NewTag() *Tag {
    a := &Tag{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "tag")
    return a
}


func (a *Tag) Set(name string, value interface{}) *Tag {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * static
 */
func (a *Tag) Static(value interface{}) *Tag {
    a.Set("static", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Tag) StaticPlaceholder(value interface{}) *Tag {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticClassName
 */
func (a *Tag) StaticClassName(value interface{}) *Tag {
    a.Set("staticClassName", value)
    return a
}

/**
 * testid
 */
func (a *Tag) Testid(value interface{}) *Tag {
    a.Set("testid", value)
    return a
}

/**
 * hidden
 */
func (a *Tag) Hidden(value interface{}) *Tag {
    a.Set("hidden", value)
    return a
}

/**
 * editorSetting
 */
func (a *Tag) EditorSetting(value interface{}) *Tag {
    a.Set("editorSetting", value)
    return a
}

/**
 * visibleOn
 */
func (a *Tag) VisibleOn(value interface{}) *Tag {
    a.Set("visibleOn", value)
    return a
}

/**
 * id
 */
func (a *Tag) Id(value interface{}) *Tag {
    a.Set("id", value)
    return a
}

/**
 * disabledOn
 */
func (a *Tag) DisabledOn(value interface{}) *Tag {
    a.Set("disabledOn", value)
    return a
}

/**
 */
func (a *Tag) Type(value interface{}) *Tag {
    a.Set("type", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Tag) StaticLabelClassName(value interface{}) *Tag {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *Tag) StaticSchema(value interface{}) *Tag {
    a.Set("staticSchema", value)
    return a
}

/**
 * style
 */
func (a *Tag) Style(value interface{}) *Tag {
    a.Set("style", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Tag) UseMobileUI(value interface{}) *Tag {
    a.Set("useMobileUI", value)
    return a
}

/**
 * checked
 */
func (a *Tag) Checked(value interface{}) *Tag {
    a.Set("checked", value)
    return a
}

/**
 * color
 */
func (a *Tag) Color(value interface{}) *Tag {
    a.Set("color", value)
    return a
}

/**
 * displayMode
 */
func (a *Tag) DisplayMode(value interface{}) *Tag {
    a.Set("displayMode", value)
    return a
}

/**
 * icon
 */
func (a *Tag) Icon(value interface{}) *Tag {
    a.Set("icon", value)
    return a
}

/**
 * className
 */
func (a *Tag) ClassName(value interface{}) *Tag {
    a.Set("className", value)
    return a
}

/**
 * disabled
 */
func (a *Tag) Disabled(value interface{}) *Tag {
    a.Set("disabled", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Tag) HiddenOn(value interface{}) *Tag {
    a.Set("hiddenOn", value)
    return a
}

/**
 * onEvent
 */
func (a *Tag) OnEvent(value interface{}) *Tag {
    a.Set("onEvent", value)
    return a
}

/**
 * staticOn
 */
func (a *Tag) StaticOn(value interface{}) *Tag {
    a.Set("staticOn", value)
    return a
}

/**
 * closable
 */
func (a *Tag) Closable(value interface{}) *Tag {
    a.Set("closable", value)
    return a
}

/**
 * closeIcon
 */
func (a *Tag) CloseIcon(value interface{}) *Tag {
    a.Set("closeIcon", value)
    return a
}

/**
 * checkable
 */
func (a *Tag) Checkable(value interface{}) *Tag {
    a.Set("checkable", value)
    return a
}

/**
 * visible
 */
func (a *Tag) Visible(value interface{}) *Tag {
    a.Set("visible", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Tag) StaticInputClassName(value interface{}) *Tag {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * label
 */
func (a *Tag) Label(value interface{}) *Tag {
    a.Set("label", value)
    return a
}
