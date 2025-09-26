package renderers


/**
 * Collapse 折叠渲染器，格式说明。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/collapse

 * @author wcz0
 * @version 6.2.2
 */
type Collapse struct {
	*BaseRenderer
}

func NewCollapse() *Collapse {
    a := &Collapse{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "collapse")
    return a
}


func (a *Collapse) Set(name string, value interface{}) *Collapse {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * visible
 */
func (a *Collapse) Visible(value interface{}) *Collapse {
    a.Set("visible", value)
    return a
}

/**
 * staticClassName
 */
func (a *Collapse) StaticClassName(value interface{}) *Collapse {
    a.Set("staticClassName", value)
    return a
}

/**
 * testid
 */
func (a *Collapse) Testid(value interface{}) *Collapse {
    a.Set("testid", value)
    return a
}

/**
 * showArrow
 */
func (a *Collapse) ShowArrow(value interface{}) *Collapse {
    a.Set("showArrow", value)
    return a
}

/**
 * hidden
 */
func (a *Collapse) Hidden(value interface{}) *Collapse {
    a.Set("hidden", value)
    return a
}

/**
 * onEvent
 */
func (a *Collapse) OnEvent(value interface{}) *Collapse {
    a.Set("onEvent", value)
    return a
}

/**
 * style
 */
func (a *Collapse) Style(value interface{}) *Collapse {
    a.Set("style", value)
    return a
}

/**
 * editorSetting
 */
func (a *Collapse) EditorSetting(value interface{}) *Collapse {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *Collapse) Type(value interface{}) *Collapse {
    a.Set("type", value)
    return a
}

/**
 * collapsable
 */
func (a *Collapse) Collapsable(value interface{}) *Collapse {
    a.Set("collapsable", value)
    return a
}

/**
 * expandIcon
 */
func (a *Collapse) ExpandIcon(value interface{}) *Collapse {
    a.Set("expandIcon", value)
    return a
}

/**
 * size
 */
func (a *Collapse) Size(value interface{}) *Collapse {
    a.Set("size", value)
    return a
}

/**
 * staticSchema
 */
func (a *Collapse) StaticSchema(value interface{}) *Collapse {
    a.Set("staticSchema", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Collapse) UseMobileUI(value interface{}) *Collapse {
    a.Set("useMobileUI", value)
    return a
}

/**
 * unmountOnExit
 */
func (a *Collapse) UnmountOnExit(value interface{}) *Collapse {
    a.Set("unmountOnExit", value)
    return a
}

/**
 * divideLine
 */
func (a *Collapse) DivideLine(value interface{}) *Collapse {
    a.Set("divideLine", value)
    return a
}

/**
 * static
 */
func (a *Collapse) Static(value interface{}) *Collapse {
    a.Set("static", value)
    return a
}

/**
 * header
 */
func (a *Collapse) Header(value interface{}) *Collapse {
    a.Set("header", value)
    return a
}

/**
 * disabled
 */
func (a *Collapse) Disabled(value interface{}) *Collapse {
    a.Set("disabled", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Collapse) HiddenOn(value interface{}) *Collapse {
    a.Set("hiddenOn", value)
    return a
}

/**
 * id
 */
func (a *Collapse) Id(value interface{}) *Collapse {
    a.Set("id", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Collapse) StaticPlaceholder(value interface{}) *Collapse {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Collapse) StaticInputClassName(value interface{}) *Collapse {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * headerPosition
 */
func (a *Collapse) HeaderPosition(value interface{}) *Collapse {
    a.Set("headerPosition", value)
    return a
}

/**
 * collapsed
 */
func (a *Collapse) Collapsed(value interface{}) *Collapse {
    a.Set("collapsed", value)
    return a
}

/**
 * collapseHeader
 */
func (a *Collapse) CollapseHeader(value interface{}) *Collapse {
    a.Set("collapseHeader", value)
    return a
}

/**
 * disabledOn
 */
func (a *Collapse) DisabledOn(value interface{}) *Collapse {
    a.Set("disabledOn", value)
    return a
}

/**
 * visibleOn
 */
func (a *Collapse) VisibleOn(value interface{}) *Collapse {
    a.Set("visibleOn", value)
    return a
}

/**
 * key
 */
func (a *Collapse) Key(value interface{}) *Collapse {
    a.Set("key", value)
    return a
}

/**
 * mountOnEnter
 */
func (a *Collapse) MountOnEnter(value interface{}) *Collapse {
    a.Set("mountOnEnter", value)
    return a
}

/**
 * body
 */
func (a *Collapse) Body(value interface{}) *Collapse {
    a.Set("body", value)
    return a
}

/**
 * className
 */
func (a *Collapse) ClassName(value interface{}) *Collapse {
    a.Set("className", value)
    return a
}

/**
 * staticOn
 */
func (a *Collapse) StaticOn(value interface{}) *Collapse {
    a.Set("staticOn", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Collapse) StaticLabelClassName(value interface{}) *Collapse {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * bodyClassName
 */
func (a *Collapse) BodyClassName(value interface{}) *Collapse {
    a.Set("bodyClassName", value)
    return a
}

/**
 * headingClassName
 */
func (a *Collapse) HeadingClassName(value interface{}) *Collapse {
    a.Set("headingClassName", value)
    return a
}
