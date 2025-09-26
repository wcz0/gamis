package renderers


/**
 * CollapseGroup 折叠渲染器，格式说明。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/collapse

 * @author wcz0
 * @version 6.2.2
 */
type CollapseGroup struct {
	*BaseRenderer
}

func NewCollapseGroup() *CollapseGroup {
    a := &CollapseGroup{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "collapse-group")
    return a
}


func (a *CollapseGroup) Set(name string, value interface{}) *CollapseGroup {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * activeKey
 */
func (a *CollapseGroup) ActiveKey(value interface{}) *CollapseGroup {
    a.Set("activeKey", value)
    return a
}

/**
 * accordion
 */
func (a *CollapseGroup) Accordion(value interface{}) *CollapseGroup {
    a.Set("accordion", value)
    return a
}

/**
 */
func (a *CollapseGroup) Type(value interface{}) *CollapseGroup {
    a.Set("type", value)
    return a
}

/**
 * static
 */
func (a *CollapseGroup) Static(value interface{}) *CollapseGroup {
    a.Set("static", value)
    return a
}

/**
 * className
 */
func (a *CollapseGroup) ClassName(value interface{}) *CollapseGroup {
    a.Set("className", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *CollapseGroup) StaticInputClassName(value interface{}) *CollapseGroup {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * enableFieldSetStyle
 */
func (a *CollapseGroup) EnableFieldSetStyle(value interface{}) *CollapseGroup {
    a.Set("enableFieldSetStyle", value)
    return a
}

/**
 * staticOn
 */
func (a *CollapseGroup) StaticOn(value interface{}) *CollapseGroup {
    a.Set("staticOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *CollapseGroup) StaticPlaceholder(value interface{}) *CollapseGroup {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * useMobileUI
 */
func (a *CollapseGroup) UseMobileUI(value interface{}) *CollapseGroup {
    a.Set("useMobileUI", value)
    return a
}

/**
 * hidden
 */
func (a *CollapseGroup) Hidden(value interface{}) *CollapseGroup {
    a.Set("hidden", value)
    return a
}

/**
 * visible
 */
func (a *CollapseGroup) Visible(value interface{}) *CollapseGroup {
    a.Set("visible", value)
    return a
}

/**
 * staticSchema
 */
func (a *CollapseGroup) StaticSchema(value interface{}) *CollapseGroup {
    a.Set("staticSchema", value)
    return a
}

/**
 * expandIcon
 */
func (a *CollapseGroup) ExpandIcon(value interface{}) *CollapseGroup {
    a.Set("expandIcon", value)
    return a
}

/**
 * expandIconPosition
 */
func (a *CollapseGroup) ExpandIconPosition(value interface{}) *CollapseGroup {
    a.Set("expandIconPosition", value)
    return a
}

/**
 * id
 */
func (a *CollapseGroup) Id(value interface{}) *CollapseGroup {
    a.Set("id", value)
    return a
}

/**
 * disabledOn
 */
func (a *CollapseGroup) DisabledOn(value interface{}) *CollapseGroup {
    a.Set("disabledOn", value)
    return a
}

/**
 * hiddenOn
 */
func (a *CollapseGroup) HiddenOn(value interface{}) *CollapseGroup {
    a.Set("hiddenOn", value)
    return a
}

/**
 * staticClassName
 */
func (a *CollapseGroup) StaticClassName(value interface{}) *CollapseGroup {
    a.Set("staticClassName", value)
    return a
}

/**
 * disabled
 */
func (a *CollapseGroup) Disabled(value interface{}) *CollapseGroup {
    a.Set("disabled", value)
    return a
}

/**
 * onEvent
 */
func (a *CollapseGroup) OnEvent(value interface{}) *CollapseGroup {
    a.Set("onEvent", value)
    return a
}

/**
 * editorSetting
 */
func (a *CollapseGroup) EditorSetting(value interface{}) *CollapseGroup {
    a.Set("editorSetting", value)
    return a
}

/**
 * testid
 */
func (a *CollapseGroup) Testid(value interface{}) *CollapseGroup {
    a.Set("testid", value)
    return a
}

/**
 * visibleOn
 */
func (a *CollapseGroup) VisibleOn(value interface{}) *CollapseGroup {
    a.Set("visibleOn", value)
    return a
}

/**
 * style
 */
func (a *CollapseGroup) Style(value interface{}) *CollapseGroup {
    a.Set("style", value)
    return a
}

/**
 * body
 */
func (a *CollapseGroup) Body(value interface{}) *CollapseGroup {
    a.Set("body", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *CollapseGroup) StaticLabelClassName(value interface{}) *CollapseGroup {
    a.Set("staticLabelClassName", value)
    return a
}
