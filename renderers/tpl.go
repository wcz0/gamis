package renderers


/**
 * tpl 渲染器

 * @author wcz0
 * @version 6.2.2
 */
type Tpl struct {
	*BaseRenderer
}

func NewTpl() *Tpl {
    a := &Tpl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "tpl")
    return a
}


func (a *Tpl) Set(name string, value interface{}) *Tpl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * useMobileUI
 */
func (a *Tpl) UseMobileUI(value interface{}) *Tpl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 可选值: tpl | html
 */
func (a *Tpl) Type(value interface{}) *Tpl {
    a.Set("type", value)
    return a
}

/**
 * tpl
 */
func (a *Tpl) Tpl(value interface{}) *Tpl {
    a.Set("tpl", value)
    return a
}

/**
 * className
 */
func (a *Tpl) ClassName(value interface{}) *Tpl {
    a.Set("className", value)
    return a
}

/**
 * editorSetting
 */
func (a *Tpl) EditorSetting(value interface{}) *Tpl {
    a.Set("editorSetting", value)
    return a
}

/**
 * html
 */
func (a *Tpl) Html(value interface{}) *Tpl {
    a.Set("html", value)
    return a
}

/**
 * inline
 */
func (a *Tpl) Inline(value interface{}) *Tpl {
    a.Set("inline", value)
    return a
}

/**
 * disabled
 */
func (a *Tpl) Disabled(value interface{}) *Tpl {
    a.Set("disabled", value)
    return a
}

/**
 * id
 */
func (a *Tpl) Id(value interface{}) *Tpl {
    a.Set("id", value)
    return a
}

/**
 * testid
 */
func (a *Tpl) Testid(value interface{}) *Tpl {
    a.Set("testid", value)
    return a
}

/**
 * raw
 */
func (a *Tpl) Raw(value interface{}) *Tpl {
    a.Set("raw", value)
    return a
}

/**
 * wrapperComponent
 */
func (a *Tpl) WrapperComponent(value interface{}) *Tpl {
    a.Set("wrapperComponent", value)
    return a
}

/**
 * hidden
 */
func (a *Tpl) Hidden(value interface{}) *Tpl {
    a.Set("hidden", value)
    return a
}

/**
 * visibleOn
 */
func (a *Tpl) VisibleOn(value interface{}) *Tpl {
    a.Set("visibleOn", value)
    return a
}

/**
 * text
 */
func (a *Tpl) Text(value interface{}) *Tpl {
    a.Set("text", value)
    return a
}

/**
 * disabledOn
 */
func (a *Tpl) DisabledOn(value interface{}) *Tpl {
    a.Set("disabledOn", value)
    return a
}

/**
 * onEvent
 */
func (a *Tpl) OnEvent(value interface{}) *Tpl {
    a.Set("onEvent", value)
    return a
}

/**
 * staticOn
 */
func (a *Tpl) StaticOn(value interface{}) *Tpl {
    a.Set("staticOn", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Tpl) StaticInputClassName(value interface{}) *Tpl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * visible
 */
func (a *Tpl) Visible(value interface{}) *Tpl {
    a.Set("visible", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Tpl) StaticPlaceholder(value interface{}) *Tpl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticSchema
 */
func (a *Tpl) StaticSchema(value interface{}) *Tpl {
    a.Set("staticSchema", value)
    return a
}

/**
 * badge
 */
func (a *Tpl) Badge(value interface{}) *Tpl {
    a.Set("badge", value)
    return a
}

/**
 * staticClassName
 */
func (a *Tpl) StaticClassName(value interface{}) *Tpl {
    a.Set("staticClassName", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Tpl) HiddenOn(value interface{}) *Tpl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * static
 */
func (a *Tpl) Static(value interface{}) *Tpl {
    a.Set("static", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Tpl) StaticLabelClassName(value interface{}) *Tpl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * style
 */
func (a *Tpl) Style(value interface{}) *Tpl {
    a.Set("style", value)
    return a
}

/**
 * testidBuilder
 */
func (a *Tpl) TestidBuilder(value interface{}) *Tpl {
    a.Set("testidBuilder", value)
    return a
}
