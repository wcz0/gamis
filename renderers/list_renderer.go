package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type ListRenderer struct {
	*BaseRenderer
}

func NewListRenderer() *ListRenderer {
    a := &ListRenderer{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "list")
    return a
}


func (a *ListRenderer) Set(name string, value interface{}) *ListRenderer {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * visibleOn
 */
func (a *ListRenderer) VisibleOn(value interface{}) *ListRenderer {
    a.Set("visibleOn", value)
    return a
}

/**
 * useMobileUI
 */
func (a *ListRenderer) UseMobileUI(value interface{}) *ListRenderer {
    a.Set("useMobileUI", value)
    return a
}

/**
 * source
 */
func (a *ListRenderer) Source(value interface{}) *ListRenderer {
    a.Set("source", value)
    return a
}

/**
 * affixHeader
 */
func (a *ListRenderer) AffixHeader(value interface{}) *ListRenderer {
    a.Set("affixHeader", value)
    return a
}

/**
 * affixFooter
 */
func (a *ListRenderer) AffixFooter(value interface{}) *ListRenderer {
    a.Set("affixFooter", value)
    return a
}

/**
 * 可选值: list | static-list
 */
func (a *ListRenderer) Type(value interface{}) *ListRenderer {
    a.Set("type", value)
    return a
}

/**
 * testid
 */
func (a *ListRenderer) Testid(value interface{}) *ListRenderer {
    a.Set("testid", value)
    return a
}

/**
 * footer
 */
func (a *ListRenderer) Footer(value interface{}) *ListRenderer {
    a.Set("footer", value)
    return a
}

/**
 * hiddenOn
 */
func (a *ListRenderer) HiddenOn(value interface{}) *ListRenderer {
    a.Set("hiddenOn", value)
    return a
}

/**
 * visible
 */
func (a *ListRenderer) Visible(value interface{}) *ListRenderer {
    a.Set("visible", value)
    return a
}

/**
 * id
 */
func (a *ListRenderer) Id(value interface{}) *ListRenderer {
    a.Set("id", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *ListRenderer) StaticLabelClassName(value interface{}) *ListRenderer {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *ListRenderer) StaticInputClassName(value interface{}) *ListRenderer {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * disabledOn
 */
func (a *ListRenderer) DisabledOn(value interface{}) *ListRenderer {
    a.Set("disabledOn", value)
    return a
}

/**
 * onEvent
 */
func (a *ListRenderer) OnEvent(value interface{}) *ListRenderer {
    a.Set("onEvent", value)
    return a
}

/**
 * editorSetting
 */
func (a *ListRenderer) EditorSetting(value interface{}) *ListRenderer {
    a.Set("editorSetting", value)
    return a
}

/**
 * header
 */
func (a *ListRenderer) Header(value interface{}) *ListRenderer {
    a.Set("header", value)
    return a
}

/**
 * hideCheckToggler
 */
func (a *ListRenderer) HideCheckToggler(value interface{}) *ListRenderer {
    a.Set("hideCheckToggler", value)
    return a
}

/**
 * itemCheckableOn
 */
func (a *ListRenderer) ItemCheckableOn(value interface{}) *ListRenderer {
    a.Set("itemCheckableOn", value)
    return a
}

/**
 * size
 */
func (a *ListRenderer) Size(value interface{}) *ListRenderer {
    a.Set("size", value)
    return a
}

/**
 * showIndexBar
 */
func (a *ListRenderer) ShowIndexBar(value interface{}) *ListRenderer {
    a.Set("showIndexBar", value)
    return a
}

/**
 * staticClassName
 */
func (a *ListRenderer) StaticClassName(value interface{}) *ListRenderer {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *ListRenderer) StaticSchema(value interface{}) *ListRenderer {
    a.Set("staticSchema", value)
    return a
}

/**
 * title
 */
func (a *ListRenderer) Title(value interface{}) *ListRenderer {
    a.Set("title", value)
    return a
}

/**
 * footerClassName
 */
func (a *ListRenderer) FooterClassName(value interface{}) *ListRenderer {
    a.Set("footerClassName", value)
    return a
}

/**
 * valueField
 */
func (a *ListRenderer) ValueField(value interface{}) *ListRenderer {
    a.Set("valueField", value)
    return a
}

/**
 * indexBarOffset
 */
func (a *ListRenderer) IndexBarOffset(value interface{}) *ListRenderer {
    a.Set("indexBarOffset", value)
    return a
}

/**
 * headerClassName
 */
func (a *ListRenderer) HeaderClassName(value interface{}) *ListRenderer {
    a.Set("headerClassName", value)
    return a
}

/**
 * indexField
 */
func (a *ListRenderer) IndexField(value interface{}) *ListRenderer {
    a.Set("indexField", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *ListRenderer) StaticPlaceholder(value interface{}) *ListRenderer {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * className
 */
func (a *ListRenderer) ClassName(value interface{}) *ListRenderer {
    a.Set("className", value)
    return a
}

/**
 * disabled
 */
func (a *ListRenderer) Disabled(value interface{}) *ListRenderer {
    a.Set("disabled", value)
    return a
}

/**
 * staticOn
 */
func (a *ListRenderer) StaticOn(value interface{}) *ListRenderer {
    a.Set("staticOn", value)
    return a
}

/**
 * itemDraggableOn
 */
func (a *ListRenderer) ItemDraggableOn(value interface{}) *ListRenderer {
    a.Set("itemDraggableOn", value)
    return a
}

/**
 * hidden
 */
func (a *ListRenderer) Hidden(value interface{}) *ListRenderer {
    a.Set("hidden", value)
    return a
}

/**
 * static
 */
func (a *ListRenderer) Static(value interface{}) *ListRenderer {
    a.Set("static", value)
    return a
}

/**
 * style
 */
func (a *ListRenderer) Style(value interface{}) *ListRenderer {
    a.Set("style", value)
    return a
}

/**
 * showFooter
 */
func (a *ListRenderer) ShowFooter(value interface{}) *ListRenderer {
    a.Set("showFooter", value)
    return a
}

/**
 * showHeader
 */
func (a *ListRenderer) ShowHeader(value interface{}) *ListRenderer {
    a.Set("showHeader", value)
    return a
}

/**
 * placeholder
 */
func (a *ListRenderer) Placeholder(value interface{}) *ListRenderer {
    a.Set("placeholder", value)
    return a
}

/**
 * checkOnItemClick
 */
func (a *ListRenderer) CheckOnItemClick(value interface{}) *ListRenderer {
    a.Set("checkOnItemClick", value)
    return a
}

/**
 * itemAction
 */
func (a *ListRenderer) ItemAction(value interface{}) *ListRenderer {
    a.Set("itemAction", value)
    return a
}

/**
 * listItem
 */
func (a *ListRenderer) ListItem(value interface{}) *ListRenderer {
    a.Set("listItem", value)
    return a
}
