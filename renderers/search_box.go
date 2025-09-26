package renderers


/**
 * 搜索框渲染器

 * @author wcz0
 * @version 6.2.2
 */
type SearchBox struct {
	*BaseRenderer
}

func NewSearchBox() *SearchBox {
    a := &SearchBox{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "search-box")
    return a
}


func (a *SearchBox) Set(name string, value interface{}) *SearchBox {
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
func (a *SearchBox) DisabledOn(value interface{}) *SearchBox {
    a.Set("disabledOn", value)
    return a
}

/**
 * id
 */
func (a *SearchBox) Id(value interface{}) *SearchBox {
    a.Set("id", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *SearchBox) StaticLabelClassName(value interface{}) *SearchBox {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *SearchBox) Type(value interface{}) *SearchBox {
    a.Set("type", value)
    return a
}

/**
 * name
 */
func (a *SearchBox) Name(value interface{}) *SearchBox {
    a.Set("name", value)
    return a
}

/**
 * enhance
 */
func (a *SearchBox) Enhance(value interface{}) *SearchBox {
    a.Set("enhance", value)
    return a
}

/**
 * hidden
 */
func (a *SearchBox) Hidden(value interface{}) *SearchBox {
    a.Set("hidden", value)
    return a
}

/**
 * visible
 */
func (a *SearchBox) Visible(value interface{}) *SearchBox {
    a.Set("visible", value)
    return a
}

/**
 * staticClassName
 */
func (a *SearchBox) StaticClassName(value interface{}) *SearchBox {
    a.Set("staticClassName", value)
    return a
}

/**
 * className
 */
func (a *SearchBox) ClassName(value interface{}) *SearchBox {
    a.Set("className", value)
    return a
}

/**
 * clearable
 */
func (a *SearchBox) Clearable(value interface{}) *SearchBox {
    a.Set("clearable", value)
    return a
}

/**
 * clearAndSubmit
 */
func (a *SearchBox) ClearAndSubmit(value interface{}) *SearchBox {
    a.Set("clearAndSubmit", value)
    return a
}

/**
 * onEvent
 */
func (a *SearchBox) OnEvent(value interface{}) *SearchBox {
    a.Set("onEvent", value)
    return a
}

/**
 * testid
 */
func (a *SearchBox) Testid(value interface{}) *SearchBox {
    a.Set("testid", value)
    return a
}

/**
 * searchImediately
 */
func (a *SearchBox) SearchImediately(value interface{}) *SearchBox {
    a.Set("searchImediately", value)
    return a
}

/**
 * hiddenOn
 */
func (a *SearchBox) HiddenOn(value interface{}) *SearchBox {
    a.Set("hiddenOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *SearchBox) StaticPlaceholder(value interface{}) *SearchBox {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * loading
 */
func (a *SearchBox) Loading(value interface{}) *SearchBox {
    a.Set("loading", value)
    return a
}

/**
 * static
 */
func (a *SearchBox) Static(value interface{}) *SearchBox {
    a.Set("static", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *SearchBox) StaticInputClassName(value interface{}) *SearchBox {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * style
 */
func (a *SearchBox) Style(value interface{}) *SearchBox {
    a.Set("style", value)
    return a
}

/**
 * editorSetting
 */
func (a *SearchBox) EditorSetting(value interface{}) *SearchBox {
    a.Set("editorSetting", value)
    return a
}

/**
 * mini
 */
func (a *SearchBox) Mini(value interface{}) *SearchBox {
    a.Set("mini", value)
    return a
}

/**
 * staticOn
 */
func (a *SearchBox) StaticOn(value interface{}) *SearchBox {
    a.Set("staticOn", value)
    return a
}

/**
 * staticSchema
 */
func (a *SearchBox) StaticSchema(value interface{}) *SearchBox {
    a.Set("staticSchema", value)
    return a
}

/**
 * disabled
 */
func (a *SearchBox) Disabled(value interface{}) *SearchBox {
    a.Set("disabled", value)
    return a
}

/**
 * visibleOn
 */
func (a *SearchBox) VisibleOn(value interface{}) *SearchBox {
    a.Set("visibleOn", value)
    return a
}

/**
 * useMobileUI
 */
func (a *SearchBox) UseMobileUI(value interface{}) *SearchBox {
    a.Set("useMobileUI", value)
    return a
}

/**
 * placeholder
 */
func (a *SearchBox) Placeholder(value interface{}) *SearchBox {
    a.Set("placeholder", value)
    return a
}
