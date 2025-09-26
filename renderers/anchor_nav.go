package renderers


/**
 * AnchorNav 锚点导航渲染器 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/anchor-nav

 * @author wcz0
 * @version 6.2.2
 */
type AnchorNav struct {
	*BaseRenderer
}

func NewAnchorNav() *AnchorNav {
    a := &AnchorNav{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "anchor-nav")
    return a
}


func (a *AnchorNav) Set(name string, value interface{}) *AnchorNav {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * onEvent
 */
func (a *AnchorNav) OnEvent(value interface{}) *AnchorNav {
    a.Set("onEvent", value)
    return a
}

/**
 * static
 */
func (a *AnchorNav) Static(value interface{}) *AnchorNav {
    a.Set("static", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *AnchorNav) StaticPlaceholder(value interface{}) *AnchorNav {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 */
func (a *AnchorNav) Type(value interface{}) *AnchorNav {
    a.Set("type", value)
    return a
}

/**
 * linkClassName
 */
func (a *AnchorNav) LinkClassName(value interface{}) *AnchorNav {
    a.Set("linkClassName", value)
    return a
}

/**
 * disabled
 */
func (a *AnchorNav) Disabled(value interface{}) *AnchorNav {
    a.Set("disabled", value)
    return a
}

/**
 * hidden
 */
func (a *AnchorNav) Hidden(value interface{}) *AnchorNav {
    a.Set("hidden", value)
    return a
}

/**
 * hiddenOn
 */
func (a *AnchorNav) HiddenOn(value interface{}) *AnchorNav {
    a.Set("hiddenOn", value)
    return a
}

/**
 * visible
 */
func (a *AnchorNav) Visible(value interface{}) *AnchorNav {
    a.Set("visible", value)
    return a
}

/**
 * staticClassName
 */
func (a *AnchorNav) StaticClassName(value interface{}) *AnchorNav {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *AnchorNav) StaticLabelClassName(value interface{}) *AnchorNav {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * className
 */
func (a *AnchorNav) ClassName(value interface{}) *AnchorNav {
    a.Set("className", value)
    return a
}

/**
 * staticOn
 */
func (a *AnchorNav) StaticOn(value interface{}) *AnchorNav {
    a.Set("staticOn", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *AnchorNav) StaticInputClassName(value interface{}) *AnchorNav {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *AnchorNav) StaticSchema(value interface{}) *AnchorNav {
    a.Set("staticSchema", value)
    return a
}

/**
 * editorSetting
 */
func (a *AnchorNav) EditorSetting(value interface{}) *AnchorNav {
    a.Set("editorSetting", value)
    return a
}

/**
 * useMobileUI
 */
func (a *AnchorNav) UseMobileUI(value interface{}) *AnchorNav {
    a.Set("useMobileUI", value)
    return a
}

/**
 * links
 */
func (a *AnchorNav) Links(value interface{}) *AnchorNav {
    a.Set("links", value)
    return a
}

/**
 * active
 */
func (a *AnchorNav) Active(value interface{}) *AnchorNav {
    a.Set("active", value)
    return a
}

/**
 * style
 */
func (a *AnchorNav) Style(value interface{}) *AnchorNav {
    a.Set("style", value)
    return a
}

/**
 * testid
 */
func (a *AnchorNav) Testid(value interface{}) *AnchorNav {
    a.Set("testid", value)
    return a
}

/**
 * sectionClassName
 */
func (a *AnchorNav) SectionClassName(value interface{}) *AnchorNav {
    a.Set("sectionClassName", value)
    return a
}

/**
 * direction
 */
func (a *AnchorNav) Direction(value interface{}) *AnchorNav {
    a.Set("direction", value)
    return a
}

/**
 * disabledOn
 */
func (a *AnchorNav) DisabledOn(value interface{}) *AnchorNav {
    a.Set("disabledOn", value)
    return a
}

/**
 * visibleOn
 */
func (a *AnchorNav) VisibleOn(value interface{}) *AnchorNav {
    a.Set("visibleOn", value)
    return a
}

/**
 * id
 */
func (a *AnchorNav) Id(value interface{}) *AnchorNav {
    a.Set("id", value)
    return a
}
