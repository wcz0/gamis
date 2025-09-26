package renderers


/**
 * 栏目容器渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/portlet

 * @author wcz0
 * @version 6.2.2
 */
type PortletTab struct {
	*BaseRenderer
}

func NewPortletTab() *PortletTab {
    a := &PortletTab{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}


func (a *PortletTab) Set(name string, value interface{}) *PortletTab {
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
func (a *PortletTab) DisabledOn(value interface{}) *PortletTab {
    a.Set("disabledOn", value)
    return a
}

/**
 * static
 */
func (a *PortletTab) Static(value interface{}) *PortletTab {
    a.Set("static", value)
    return a
}

/**
 * body
 */
func (a *PortletTab) Body(value interface{}) *PortletTab {
    a.Set("body", value)
    return a
}

/**
 * unmountOnExit
 */
func (a *PortletTab) UnmountOnExit(value interface{}) *PortletTab {
    a.Set("unmountOnExit", value)
    return a
}

/**
 * id
 */
func (a *PortletTab) Id(value interface{}) *PortletTab {
    a.Set("id", value)
    return a
}

/**
 * onEvent
 */
func (a *PortletTab) OnEvent(value interface{}) *PortletTab {
    a.Set("onEvent", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *PortletTab) StaticLabelClassName(value interface{}) *PortletTab {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * useMobileUI
 */
func (a *PortletTab) UseMobileUI(value interface{}) *PortletTab {
    a.Set("useMobileUI", value)
    return a
}

/**
 * tab
 */
func (a *PortletTab) Tab(value interface{}) *PortletTab {
    a.Set("tab", value)
    return a
}

/**
 * iconPosition
 */
func (a *PortletTab) IconPosition(value interface{}) *PortletTab {
    a.Set("iconPosition", value)
    return a
}

/**
 * hidden
 */
func (a *PortletTab) Hidden(value interface{}) *PortletTab {
    a.Set("hidden", value)
    return a
}

/**
 * hiddenOn
 */
func (a *PortletTab) HiddenOn(value interface{}) *PortletTab {
    a.Set("hiddenOn", value)
    return a
}

/**
 * visibleOn
 */
func (a *PortletTab) VisibleOn(value interface{}) *PortletTab {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticSchema
 */
func (a *PortletTab) StaticSchema(value interface{}) *PortletTab {
    a.Set("staticSchema", value)
    return a
}

/**
 * toolbar
 */
func (a *PortletTab) Toolbar(value interface{}) *PortletTab {
    a.Set("toolbar", value)
    return a
}

/**
 * reload
 */
func (a *PortletTab) Reload(value interface{}) *PortletTab {
    a.Set("reload", value)
    return a
}

/**
 * staticClassName
 */
func (a *PortletTab) StaticClassName(value interface{}) *PortletTab {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *PortletTab) StaticInputClassName(value interface{}) *PortletTab {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * editorSetting
 */
func (a *PortletTab) EditorSetting(value interface{}) *PortletTab {
    a.Set("editorSetting", value)
    return a
}

/**
 * title
 */
func (a *PortletTab) Title(value interface{}) *PortletTab {
    a.Set("title", value)
    return a
}

/**
 * mountOnEnter
 */
func (a *PortletTab) MountOnEnter(value interface{}) *PortletTab {
    a.Set("mountOnEnter", value)
    return a
}

/**
 * disabled
 */
func (a *PortletTab) Disabled(value interface{}) *PortletTab {
    a.Set("disabled", value)
    return a
}

/**
 * visible
 */
func (a *PortletTab) Visible(value interface{}) *PortletTab {
    a.Set("visible", value)
    return a
}

/**
 * icon
 */
func (a *PortletTab) Icon(value interface{}) *PortletTab {
    a.Set("icon", value)
    return a
}

/**
 * className
 */
func (a *PortletTab) ClassName(value interface{}) *PortletTab {
    a.Set("className", value)
    return a
}

/**
 * staticOn
 */
func (a *PortletTab) StaticOn(value interface{}) *PortletTab {
    a.Set("staticOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *PortletTab) StaticPlaceholder(value interface{}) *PortletTab {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * style
 */
func (a *PortletTab) Style(value interface{}) *PortletTab {
    a.Set("style", value)
    return a
}
