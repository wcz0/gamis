package renderers


/**
 * Link 链接展示控件。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/link

 * @author wcz0
 * @version 6.2.2
 */
type Link struct {
	*BaseRenderer
}

func NewLink() *Link {
    a := &Link{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "link")
    return a
}


func (a *Link) Set(name string, value interface{}) *Link {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * href
 */
func (a *Link) Href(value interface{}) *Link {
    a.Set("href", value)
    return a
}

/**
 * hidden
 */
func (a *Link) Hidden(value interface{}) *Link {
    a.Set("hidden", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Link) HiddenOn(value interface{}) *Link {
    a.Set("hiddenOn", value)
    return a
}

/**
 * body
 */
func (a *Link) Body(value interface{}) *Link {
    a.Set("body", value)
    return a
}

/**
 * staticOn
 */
func (a *Link) StaticOn(value interface{}) *Link {
    a.Set("staticOn", value)
    return a
}

/**
 * editorSetting
 */
func (a *Link) EditorSetting(value interface{}) *Link {
    a.Set("editorSetting", value)
    return a
}

/**
 * disabledOn
 */
func (a *Link) DisabledOn(value interface{}) *Link {
    a.Set("disabledOn", value)
    return a
}

/**
 * id
 */
func (a *Link) Id(value interface{}) *Link {
    a.Set("id", value)
    return a
}

/**
 * staticClassName
 */
func (a *Link) StaticClassName(value interface{}) *Link {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Link) StaticInputClassName(value interface{}) *Link {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *Link) StaticSchema(value interface{}) *Link {
    a.Set("staticSchema", value)
    return a
}

/**
 * badge
 */
func (a *Link) Badge(value interface{}) *Link {
    a.Set("badge", value)
    return a
}

/**
 * visibleOn
 */
func (a *Link) VisibleOn(value interface{}) *Link {
    a.Set("visibleOn", value)
    return a
}

/**
 * static
 */
func (a *Link) Static(value interface{}) *Link {
    a.Set("static", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Link) StaticPlaceholder(value interface{}) *Link {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Link) UseMobileUI(value interface{}) *Link {
    a.Set("useMobileUI", value)
    return a
}

/**
 * testid
 */
func (a *Link) Testid(value interface{}) *Link {
    a.Set("testid", value)
    return a
}

/**
 * icon
 */
func (a *Link) Icon(value interface{}) *Link {
    a.Set("icon", value)
    return a
}

/**
 * htmlTarget
 */
func (a *Link) HtmlTarget(value interface{}) *Link {
    a.Set("htmlTarget", value)
    return a
}

/**
 * onEvent
 */
func (a *Link) OnEvent(value interface{}) *Link {
    a.Set("onEvent", value)
    return a
}

/**
 * className
 */
func (a *Link) ClassName(value interface{}) *Link {
    a.Set("className", value)
    return a
}

/**
 * disabled
 */
func (a *Link) Disabled(value interface{}) *Link {
    a.Set("disabled", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Link) StaticLabelClassName(value interface{}) *Link {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * rightIcon
 */
func (a *Link) RightIcon(value interface{}) *Link {
    a.Set("rightIcon", value)
    return a
}

/**
 * visible
 */
func (a *Link) Visible(value interface{}) *Link {
    a.Set("visible", value)
    return a
}

/**
 * style
 */
func (a *Link) Style(value interface{}) *Link {
    a.Set("style", value)
    return a
}

/**
 */
func (a *Link) Type(value interface{}) *Link {
    a.Set("type", value)
    return a
}

/**
 * blank
 */
func (a *Link) Blank(value interface{}) *Link {
    a.Set("blank", value)
    return a
}
