package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type Cards struct {
	*BaseRenderer
}

func NewCards() *Cards {
    a := &Cards{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "cards")
    return a
}


func (a *Cards) Set(name string, value interface{}) *Cards {
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
func (a *Cards) DisabledOn(value interface{}) *Cards {
    a.Set("disabledOn", value)
    return a
}

/**
 * staticClassName
 */
func (a *Cards) StaticClassName(value interface{}) *Cards {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Cards) StaticInputClassName(value interface{}) *Cards {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *Cards) Type(value interface{}) *Cards {
    a.Set("type", value)
    return a
}

/**
 * headerClassName
 */
func (a *Cards) HeaderClassName(value interface{}) *Cards {
    a.Set("headerClassName", value)
    return a
}

/**
 * itemClassName
 */
func (a *Cards) ItemClassName(value interface{}) *Cards {
    a.Set("itemClassName", value)
    return a
}

/**
 * showHeader
 */
func (a *Cards) ShowHeader(value interface{}) *Cards {
    a.Set("showHeader", value)
    return a
}

/**
 * itemDraggableOn
 */
func (a *Cards) ItemDraggableOn(value interface{}) *Cards {
    a.Set("itemDraggableOn", value)
    return a
}

/**
 * disabled
 */
func (a *Cards) Disabled(value interface{}) *Cards {
    a.Set("disabled", value)
    return a
}

/**
 * hidden
 */
func (a *Cards) Hidden(value interface{}) *Cards {
    a.Set("hidden", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Cards) HiddenOn(value interface{}) *Cards {
    a.Set("hiddenOn", value)
    return a
}

/**
 * id
 */
func (a *Cards) Id(value interface{}) *Cards {
    a.Set("id", value)
    return a
}

/**
 * onEvent
 */
func (a *Cards) OnEvent(value interface{}) *Cards {
    a.Set("onEvent", value)
    return a
}

/**
 * static
 */
func (a *Cards) Static(value interface{}) *Cards {
    a.Set("static", value)
    return a
}

/**
 * loadingConfig
 */
func (a *Cards) LoadingConfig(value interface{}) *Cards {
    a.Set("loadingConfig", value)
    return a
}

/**
 * style
 */
func (a *Cards) Style(value interface{}) *Cards {
    a.Set("style", value)
    return a
}

/**
 * card
 */
func (a *Cards) Card(value interface{}) *Cards {
    a.Set("card", value)
    return a
}

/**
 * placeholder
 */
func (a *Cards) Placeholder(value interface{}) *Cards {
    a.Set("placeholder", value)
    return a
}

/**
 * source
 */
func (a *Cards) Source(value interface{}) *Cards {
    a.Set("source", value)
    return a
}

/**
 * title
 */
func (a *Cards) Title(value interface{}) *Cards {
    a.Set("title", value)
    return a
}

/**
 * affixFooter
 */
func (a *Cards) AffixFooter(value interface{}) *Cards {
    a.Set("affixFooter", value)
    return a
}

/**
 * header
 */
func (a *Cards) Header(value interface{}) *Cards {
    a.Set("header", value)
    return a
}

/**
 * testid
 */
func (a *Cards) Testid(value interface{}) *Cards {
    a.Set("testid", value)
    return a
}

/**
 * showFooter
 */
func (a *Cards) ShowFooter(value interface{}) *Cards {
    a.Set("showFooter", value)
    return a
}

/**
 * hideCheckToggler
 */
func (a *Cards) HideCheckToggler(value interface{}) *Cards {
    a.Set("hideCheckToggler", value)
    return a
}

/**
 * footer
 */
func (a *Cards) Footer(value interface{}) *Cards {
    a.Set("footer", value)
    return a
}

/**
 * itemCheckableOn
 */
func (a *Cards) ItemCheckableOn(value interface{}) *Cards {
    a.Set("itemCheckableOn", value)
    return a
}

/**
 * masonryLayout
 */
func (a *Cards) MasonryLayout(value interface{}) *Cards {
    a.Set("masonryLayout", value)
    return a
}

/**
 * valueField
 */
func (a *Cards) ValueField(value interface{}) *Cards {
    a.Set("valueField", value)
    return a
}

/**
 * className
 */
func (a *Cards) ClassName(value interface{}) *Cards {
    a.Set("className", value)
    return a
}

/**
 * visible
 */
func (a *Cards) Visible(value interface{}) *Cards {
    a.Set("visible", value)
    return a
}

/**
 * staticOn
 */
func (a *Cards) StaticOn(value interface{}) *Cards {
    a.Set("staticOn", value)
    return a
}

/**
 * staticSchema
 */
func (a *Cards) StaticSchema(value interface{}) *Cards {
    a.Set("staticSchema", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Cards) UseMobileUI(value interface{}) *Cards {
    a.Set("useMobileUI", value)
    return a
}

/**
 * visibleOn
 */
func (a *Cards) VisibleOn(value interface{}) *Cards {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Cards) StaticPlaceholder(value interface{}) *Cards {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Cards) StaticLabelClassName(value interface{}) *Cards {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * editorSetting
 */
func (a *Cards) EditorSetting(value interface{}) *Cards {
    a.Set("editorSetting", value)
    return a
}

/**
 * checkOnItemClick
 */
func (a *Cards) CheckOnItemClick(value interface{}) *Cards {
    a.Set("checkOnItemClick", value)
    return a
}

/**
 * footerClassName
 */
func (a *Cards) FooterClassName(value interface{}) *Cards {
    a.Set("footerClassName", value)
    return a
}

/**
 * affixHeader
 */
func (a *Cards) AffixHeader(value interface{}) *Cards {
    a.Set("affixHeader", value)
    return a
}
