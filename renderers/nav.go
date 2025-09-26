package renderers


/**
 * Nav 导航渲染器 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/nav

 * @author wcz0
 * @version 6.2.2
 */
type Nav struct {
	*BaseRenderer
}

func NewNav() *Nav {
    a := &Nav{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "nav")
    return a
}


func (a *Nav) Set(name string, value interface{}) *Nav {
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
func (a *Nav) VisibleOn(value interface{}) *Nav {
    a.Set("visibleOn", value)
    return a
}

/**
 * onEvent
 */
func (a *Nav) OnEvent(value interface{}) *Nav {
    a.Set("onEvent", value)
    return a
}

/**
 * staticClassName
 */
func (a *Nav) StaticClassName(value interface{}) *Nav {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *Nav) StaticSchema(value interface{}) *Nav {
    a.Set("staticSchema", value)
    return a
}

/**
 * saveOrderApi
 */
func (a *Nav) SaveOrderApi(value interface{}) *Nav {
    a.Set("saveOrderApi", value)
    return a
}

/**
 * overflow
 */
func (a *Nav) Overflow(value interface{}) *Nav {
    a.Set("overflow", value)
    return a
}

/**
 * mode
 */
func (a *Nav) Mode(value interface{}) *Nav {
    a.Set("mode", value)
    return a
}

/**
 * expandIcon
 */
func (a *Nav) ExpandIcon(value interface{}) *Nav {
    a.Set("expandIcon", value)
    return a
}

/**
 * staticOn
 */
func (a *Nav) StaticOn(value interface{}) *Nav {
    a.Set("staticOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Nav) StaticPlaceholder(value interface{}) *Nav {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * style
 */
func (a *Nav) Style(value interface{}) *Nav {
    a.Set("style", value)
    return a
}

/**
 */
func (a *Nav) Type(value interface{}) *Nav {
    a.Set("type", value)
    return a
}

/**
 * links
 */
func (a *Nav) Links(value interface{}) *Nav {
    a.Set("links", value)
    return a
}

/**
 * itemBadge
 */
func (a *Nav) ItemBadge(value interface{}) *Nav {
    a.Set("itemBadge", value)
    return a
}

/**
 * showKey
 */
func (a *Nav) ShowKey(value interface{}) *Nav {
    a.Set("showKey", value)
    return a
}

/**
 * expandPosition
 */
func (a *Nav) ExpandPosition(value interface{}) *Nav {
    a.Set("expandPosition", value)
    return a
}

/**
 * indentSize
 */
func (a *Nav) IndentSize(value interface{}) *Nav {
    a.Set("indentSize", value)
    return a
}

/**
 * popupClassName
 */
func (a *Nav) PopupClassName(value interface{}) *Nav {
    a.Set("popupClassName", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Nav) HiddenOn(value interface{}) *Nav {
    a.Set("hiddenOn", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Nav) StaticInputClassName(value interface{}) *Nav {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * level
 */
func (a *Nav) Level(value interface{}) *Nav {
    a.Set("level", value)
    return a
}

/**
 * searchConfig
 */
func (a *Nav) SearchConfig(value interface{}) *Nav {
    a.Set("searchConfig", value)
    return a
}

/**
 * disabledOn
 */
func (a *Nav) DisabledOn(value interface{}) *Nav {
    a.Set("disabledOn", value)
    return a
}

/**
 * id
 */
func (a *Nav) Id(value interface{}) *Nav {
    a.Set("id", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Nav) UseMobileUI(value interface{}) *Nav {
    a.Set("useMobileUI", value)
    return a
}

/**
 * themeColor
 */
func (a *Nav) ThemeColor(value interface{}) *Nav {
    a.Set("themeColor", value)
    return a
}

/**
 * className
 */
func (a *Nav) ClassName(value interface{}) *Nav {
    a.Set("className", value)
    return a
}

/**
 * disabled
 */
func (a *Nav) Disabled(value interface{}) *Nav {
    a.Set("disabled", value)
    return a
}

/**
 * deferApi
 */
func (a *Nav) DeferApi(value interface{}) *Nav {
    a.Set("deferApi", value)
    return a
}

/**
 * stacked
 */
func (a *Nav) Stacked(value interface{}) *Nav {
    a.Set("stacked", value)
    return a
}

/**
 * badge
 */
func (a *Nav) Badge(value interface{}) *Nav {
    a.Set("badge", value)
    return a
}

/**
 * collapsed
 */
func (a *Nav) Collapsed(value interface{}) *Nav {
    a.Set("collapsed", value)
    return a
}

/**
 * hidden
 */
func (a *Nav) Hidden(value interface{}) *Nav {
    a.Set("hidden", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Nav) StaticLabelClassName(value interface{}) *Nav {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * editorSetting
 */
func (a *Nav) EditorSetting(value interface{}) *Nav {
    a.Set("editorSetting", value)
    return a
}

/**
 * source
 */
func (a *Nav) Source(value interface{}) *Nav {
    a.Set("source", value)
    return a
}

/**
 * itemActions
 */
func (a *Nav) ItemActions(value interface{}) *Nav {
    a.Set("itemActions", value)
    return a
}

/**
 * dragOnSameLevel
 */
func (a *Nav) DragOnSameLevel(value interface{}) *Nav {
    a.Set("dragOnSameLevel", value)
    return a
}

/**
 * visible
 */
func (a *Nav) Visible(value interface{}) *Nav {
    a.Set("visible", value)
    return a
}

/**
 * static
 */
func (a *Nav) Static(value interface{}) *Nav {
    a.Set("static", value)
    return a
}

/**
 * testid
 */
func (a *Nav) Testid(value interface{}) *Nav {
    a.Set("testid", value)
    return a
}

/**
 * draggable
 */
func (a *Nav) Draggable(value interface{}) *Nav {
    a.Set("draggable", value)
    return a
}

/**
 * defaultOpenLevel
 */
func (a *Nav) DefaultOpenLevel(value interface{}) *Nav {
    a.Set("defaultOpenLevel", value)
    return a
}

/**
 * accordion
 */
func (a *Nav) Accordion(value interface{}) *Nav {
    a.Set("accordion", value)
    return a
}

/**
 * searchable
 */
func (a *Nav) Searchable(value interface{}) *Nav {
    a.Set("searchable", value)
    return a
}
