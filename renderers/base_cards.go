package renderers


/**
 * Cards 卡片集合渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/cards

 * @author wcz0
 * @version 6.2.2
 */
type BaseCards struct {
	*BaseRenderer
}

func NewBaseCards() *BaseCards {
    a := &BaseCards{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}


func (a *BaseCards) Set(name string, value interface{}) *BaseCards {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * style
 */
func (a *BaseCards) Style(value interface{}) *BaseCards {
    a.Set("style", value)
    return a
}

/**
 * useMobileUI
 */
func (a *BaseCards) UseMobileUI(value interface{}) *BaseCards {
    a.Set("useMobileUI", value)
    return a
}

/**
 * type
 */
func (a *BaseCards) Type(value interface{}) *BaseCards {
    a.Set("type", value)
    return a
}

/**
 * showHeader
 */
func (a *BaseCards) ShowHeader(value interface{}) *BaseCards {
    a.Set("showHeader", value)
    return a
}

/**
 * className
 */
func (a *BaseCards) ClassName(value interface{}) *BaseCards {
    a.Set("className", value)
    return a
}

/**
 * visible
 */
func (a *BaseCards) Visible(value interface{}) *BaseCards {
    a.Set("visible", value)
    return a
}

/**
 * id
 */
func (a *BaseCards) Id(value interface{}) *BaseCards {
    a.Set("id", value)
    return a
}

/**
 * footerClassName
 */
func (a *BaseCards) FooterClassName(value interface{}) *BaseCards {
    a.Set("footerClassName", value)
    return a
}

/**
 * title
 */
func (a *BaseCards) Title(value interface{}) *BaseCards {
    a.Set("title", value)
    return a
}

/**
 * itemCheckableOn
 */
func (a *BaseCards) ItemCheckableOn(value interface{}) *BaseCards {
    a.Set("itemCheckableOn", value)
    return a
}

/**
 * onEvent
 */
func (a *BaseCards) OnEvent(value interface{}) *BaseCards {
    a.Set("onEvent", value)
    return a
}

/**
 * staticClassName
 */
func (a *BaseCards) StaticClassName(value interface{}) *BaseCards {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *BaseCards) StaticInputClassName(value interface{}) *BaseCards {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * card
 */
func (a *BaseCards) Card(value interface{}) *BaseCards {
    a.Set("card", value)
    return a
}

/**
 * hideCheckToggler
 */
func (a *BaseCards) HideCheckToggler(value interface{}) *BaseCards {
    a.Set("hideCheckToggler", value)
    return a
}

/**
 * affixFooter
 */
func (a *BaseCards) AffixFooter(value interface{}) *BaseCards {
    a.Set("affixFooter", value)
    return a
}

/**
 * itemDraggableOn
 */
func (a *BaseCards) ItemDraggableOn(value interface{}) *BaseCards {
    a.Set("itemDraggableOn", value)
    return a
}

/**
 * masonryLayout
 */
func (a *BaseCards) MasonryLayout(value interface{}) *BaseCards {
    a.Set("masonryLayout", value)
    return a
}

/**
 * visibleOn
 */
func (a *BaseCards) VisibleOn(value interface{}) *BaseCards {
    a.Set("visibleOn", value)
    return a
}

/**
 * header
 */
func (a *BaseCards) Header(value interface{}) *BaseCards {
    a.Set("header", value)
    return a
}

/**
 * headerClassName
 */
func (a *BaseCards) HeaderClassName(value interface{}) *BaseCards {
    a.Set("headerClassName", value)
    return a
}

/**
 * itemClassName
 */
func (a *BaseCards) ItemClassName(value interface{}) *BaseCards {
    a.Set("itemClassName", value)
    return a
}

/**
 * hidden
 */
func (a *BaseCards) Hidden(value interface{}) *BaseCards {
    a.Set("hidden", value)
    return a
}

/**
 * hiddenOn
 */
func (a *BaseCards) HiddenOn(value interface{}) *BaseCards {
    a.Set("hiddenOn", value)
    return a
}

/**
 * valueField
 */
func (a *BaseCards) ValueField(value interface{}) *BaseCards {
    a.Set("valueField", value)
    return a
}

/**
 * disabledOn
 */
func (a *BaseCards) DisabledOn(value interface{}) *BaseCards {
    a.Set("disabledOn", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *BaseCards) StaticLabelClassName(value interface{}) *BaseCards {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *BaseCards) StaticSchema(value interface{}) *BaseCards {
    a.Set("staticSchema", value)
    return a
}

/**
 * showFooter
 */
func (a *BaseCards) ShowFooter(value interface{}) *BaseCards {
    a.Set("showFooter", value)
    return a
}

/**
 * static
 */
func (a *BaseCards) Static(value interface{}) *BaseCards {
    a.Set("static", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *BaseCards) StaticPlaceholder(value interface{}) *BaseCards {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * placeholder
 */
func (a *BaseCards) Placeholder(value interface{}) *BaseCards {
    a.Set("placeholder", value)
    return a
}

/**
 * checkOnItemClick
 */
func (a *BaseCards) CheckOnItemClick(value interface{}) *BaseCards {
    a.Set("checkOnItemClick", value)
    return a
}

/**
 * staticOn
 */
func (a *BaseCards) StaticOn(value interface{}) *BaseCards {
    a.Set("staticOn", value)
    return a
}

/**
 * editorSetting
 */
func (a *BaseCards) EditorSetting(value interface{}) *BaseCards {
    a.Set("editorSetting", value)
    return a
}

/**
 * source
 */
func (a *BaseCards) Source(value interface{}) *BaseCards {
    a.Set("source", value)
    return a
}

/**
 * affixHeader
 */
func (a *BaseCards) AffixHeader(value interface{}) *BaseCards {
    a.Set("affixHeader", value)
    return a
}

/**
 * footer
 */
func (a *BaseCards) Footer(value interface{}) *BaseCards {
    a.Set("footer", value)
    return a
}

/**
 * testid
 */
func (a *BaseCards) Testid(value interface{}) *BaseCards {
    a.Set("testid", value)
    return a
}

/**
 * loadingConfig
 */
func (a *BaseCards) LoadingConfig(value interface{}) *BaseCards {
    a.Set("loadingConfig", value)
    return a
}

/**
 * disabled
 */
func (a *BaseCards) Disabled(value interface{}) *BaseCards {
    a.Set("disabled", value)
    return a
}
