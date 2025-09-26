package renderers


/**
 * AnchorNavSection 锚点区域渲染器 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/anchor-nav

 * @author wcz0
 * @version 6.2.2
 */
type AnchorNavSection struct {
	*BaseRenderer
}

func NewAnchorNavSection() *AnchorNavSection {
    a := &AnchorNavSection{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}


func (a *AnchorNavSection) Set(name string, value interface{}) *AnchorNavSection {
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
func (a *AnchorNavSection) Style(value interface{}) *AnchorNavSection {
    a.Set("style", value)
    return a
}

/**
 * title
 */
func (a *AnchorNavSection) Title(value interface{}) *AnchorNavSection {
    a.Set("title", value)
    return a
}

/**
 * children
 */
func (a *AnchorNavSection) Children(value interface{}) *AnchorNavSection {
    a.Set("children", value)
    return a
}

/**
 * className
 */
func (a *AnchorNavSection) ClassName(value interface{}) *AnchorNavSection {
    a.Set("className", value)
    return a
}

/**
 * disabled
 */
func (a *AnchorNavSection) Disabled(value interface{}) *AnchorNavSection {
    a.Set("disabled", value)
    return a
}

/**
 * hidden
 */
func (a *AnchorNavSection) Hidden(value interface{}) *AnchorNavSection {
    a.Set("hidden", value)
    return a
}

/**
 * hiddenOn
 */
func (a *AnchorNavSection) HiddenOn(value interface{}) *AnchorNavSection {
    a.Set("hiddenOn", value)
    return a
}

/**
 * visibleOn
 */
func (a *AnchorNavSection) VisibleOn(value interface{}) *AnchorNavSection {
    a.Set("visibleOn", value)
    return a
}

/**
 * id
 */
func (a *AnchorNavSection) Id(value interface{}) *AnchorNavSection {
    a.Set("id", value)
    return a
}

/**
 * editorSetting
 */
func (a *AnchorNavSection) EditorSetting(value interface{}) *AnchorNavSection {
    a.Set("editorSetting", value)
    return a
}

/**
 * disabledOn
 */
func (a *AnchorNavSection) DisabledOn(value interface{}) *AnchorNavSection {
    a.Set("disabledOn", value)
    return a
}

/**
 * visible
 */
func (a *AnchorNavSection) Visible(value interface{}) *AnchorNavSection {
    a.Set("visible", value)
    return a
}

/**
 * static
 */
func (a *AnchorNavSection) Static(value interface{}) *AnchorNavSection {
    a.Set("static", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *AnchorNavSection) StaticPlaceholder(value interface{}) *AnchorNavSection {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *AnchorNavSection) StaticInputClassName(value interface{}) *AnchorNavSection {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *AnchorNavSection) StaticSchema(value interface{}) *AnchorNavSection {
    a.Set("staticSchema", value)
    return a
}

/**
 * onEvent
 */
func (a *AnchorNavSection) OnEvent(value interface{}) *AnchorNavSection {
    a.Set("onEvent", value)
    return a
}

/**
 * staticOn
 */
func (a *AnchorNavSection) StaticOn(value interface{}) *AnchorNavSection {
    a.Set("staticOn", value)
    return a
}

/**
 * staticClassName
 */
func (a *AnchorNavSection) StaticClassName(value interface{}) *AnchorNavSection {
    a.Set("staticClassName", value)
    return a
}

/**
 * useMobileUI
 */
func (a *AnchorNavSection) UseMobileUI(value interface{}) *AnchorNavSection {
    a.Set("useMobileUI", value)
    return a
}

/**
 * href
 */
func (a *AnchorNavSection) Href(value interface{}) *AnchorNavSection {
    a.Set("href", value)
    return a
}

/**
 * body
 */
func (a *AnchorNavSection) Body(value interface{}) *AnchorNavSection {
    a.Set("body", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *AnchorNavSection) StaticLabelClassName(value interface{}) *AnchorNavSection {
    a.Set("staticLabelClassName", value)
    return a
}
