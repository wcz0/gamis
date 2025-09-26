package renderers


/**
 * List 列表展示控件。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/card

 * @author wcz0
 * @version 6.2.2
 */
type BaseList struct {
	*BaseRenderer
}

func NewBaseList() *BaseList {
    a := &BaseList{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}


func (a *BaseList) Set(name string, value interface{}) *BaseList {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * listItem
 */
func (a *BaseList) ListItem(value interface{}) *BaseList {
    a.Set("listItem", value)
    return a
}

/**
 * hideCheckToggler
 */
func (a *BaseList) HideCheckToggler(value interface{}) *BaseList {
    a.Set("hideCheckToggler", value)
    return a
}

/**
 * itemCheckableOn
 */
func (a *BaseList) ItemCheckableOn(value interface{}) *BaseList {
    a.Set("itemCheckableOn", value)
    return a
}

/**
 * checkOnItemClick
 */
func (a *BaseList) CheckOnItemClick(value interface{}) *BaseList {
    a.Set("checkOnItemClick", value)
    return a
}

/**
 * size
 */
func (a *BaseList) Size(value interface{}) *BaseList {
    a.Set("size", value)
    return a
}

/**
 * disabledOn
 */
func (a *BaseList) DisabledOn(value interface{}) *BaseList {
    a.Set("disabledOn", value)
    return a
}

/**
 * staticSchema
 */
func (a *BaseList) StaticSchema(value interface{}) *BaseList {
    a.Set("staticSchema", value)
    return a
}

/**
 * style
 */
func (a *BaseList) Style(value interface{}) *BaseList {
    a.Set("style", value)
    return a
}

/**
 * headerClassName
 */
func (a *BaseList) HeaderClassName(value interface{}) *BaseList {
    a.Set("headerClassName", value)
    return a
}

/**
 * affixFooter
 */
func (a *BaseList) AffixFooter(value interface{}) *BaseList {
    a.Set("affixFooter", value)
    return a
}

/**
 * indexField
 */
func (a *BaseList) IndexField(value interface{}) *BaseList {
    a.Set("indexField", value)
    return a
}

/**
 * staticOn
 */
func (a *BaseList) StaticOn(value interface{}) *BaseList {
    a.Set("staticOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *BaseList) StaticPlaceholder(value interface{}) *BaseList {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * useMobileUI
 */
func (a *BaseList) UseMobileUI(value interface{}) *BaseList {
    a.Set("useMobileUI", value)
    return a
}

/**
 * itemAction
 */
func (a *BaseList) ItemAction(value interface{}) *BaseList {
    a.Set("itemAction", value)
    return a
}

/**
 * showIndexBar
 */
func (a *BaseList) ShowIndexBar(value interface{}) *BaseList {
    a.Set("showIndexBar", value)
    return a
}

/**
 * disabled
 */
func (a *BaseList) Disabled(value interface{}) *BaseList {
    a.Set("disabled", value)
    return a
}

/**
 * title
 */
func (a *BaseList) Title(value interface{}) *BaseList {
    a.Set("title", value)
    return a
}

/**
 * placeholder
 */
func (a *BaseList) Placeholder(value interface{}) *BaseList {
    a.Set("placeholder", value)
    return a
}

/**
 * affixHeader
 */
func (a *BaseList) AffixHeader(value interface{}) *BaseList {
    a.Set("affixHeader", value)
    return a
}

/**
 * itemDraggableOn
 */
func (a *BaseList) ItemDraggableOn(value interface{}) *BaseList {
    a.Set("itemDraggableOn", value)
    return a
}

/**
 * className
 */
func (a *BaseList) ClassName(value interface{}) *BaseList {
    a.Set("className", value)
    return a
}

/**
 * id
 */
func (a *BaseList) Id(value interface{}) *BaseList {
    a.Set("id", value)
    return a
}

/**
 * editorSetting
 */
func (a *BaseList) EditorSetting(value interface{}) *BaseList {
    a.Set("editorSetting", value)
    return a
}

/**
 * testid
 */
func (a *BaseList) Testid(value interface{}) *BaseList {
    a.Set("testid", value)
    return a
}

/**
 * showHeader
 */
func (a *BaseList) ShowHeader(value interface{}) *BaseList {
    a.Set("showHeader", value)
    return a
}

/**
 * indexBarOffset
 */
func (a *BaseList) IndexBarOffset(value interface{}) *BaseList {
    a.Set("indexBarOffset", value)
    return a
}

/**
 * visibleOn
 */
func (a *BaseList) VisibleOn(value interface{}) *BaseList {
    a.Set("visibleOn", value)
    return a
}

/**
 * footerClassName
 */
func (a *BaseList) FooterClassName(value interface{}) *BaseList {
    a.Set("footerClassName", value)
    return a
}

/**
 * source
 */
func (a *BaseList) Source(value interface{}) *BaseList {
    a.Set("source", value)
    return a
}

/**
 * showFooter
 */
func (a *BaseList) ShowFooter(value interface{}) *BaseList {
    a.Set("showFooter", value)
    return a
}

/**
 * hidden
 */
func (a *BaseList) Hidden(value interface{}) *BaseList {
    a.Set("hidden", value)
    return a
}

/**
 * hiddenOn
 */
func (a *BaseList) HiddenOn(value interface{}) *BaseList {
    a.Set("hiddenOn", value)
    return a
}

/**
 * visible
 */
func (a *BaseList) Visible(value interface{}) *BaseList {
    a.Set("visible", value)
    return a
}

/**
 * onEvent
 */
func (a *BaseList) OnEvent(value interface{}) *BaseList {
    a.Set("onEvent", value)
    return a
}

/**
 * footer
 */
func (a *BaseList) Footer(value interface{}) *BaseList {
    a.Set("footer", value)
    return a
}

/**
 * valueField
 */
func (a *BaseList) ValueField(value interface{}) *BaseList {
    a.Set("valueField", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *BaseList) StaticLabelClassName(value interface{}) *BaseList {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *BaseList) StaticInputClassName(value interface{}) *BaseList {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * type
 */
func (a *BaseList) Type(value interface{}) *BaseList {
    a.Set("type", value)
    return a
}

/**
 * static
 */
func (a *BaseList) Static(value interface{}) *BaseList {
    a.Set("static", value)
    return a
}

/**
 * staticClassName
 */
func (a *BaseList) StaticClassName(value interface{}) *BaseList {
    a.Set("staticClassName", value)
    return a
}

/**
 * header
 */
func (a *BaseList) Header(value interface{}) *BaseList {
    a.Set("header", value)
    return a
}
