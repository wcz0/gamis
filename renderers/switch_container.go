package renderers


/**
 * SwitchContainer 状态容器渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/state-container

 * @author wcz0
 * @version 6.2.2
 */
type SwitchContainer struct {
	*BaseRenderer
}

func NewSwitchContainer() *SwitchContainer {
    a := &SwitchContainer{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "switch-container")
    return a
}


func (a *SwitchContainer) Set(name string, value interface{}) *SwitchContainer {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * staticSchema
 */
func (a *SwitchContainer) StaticSchema(value interface{}) *SwitchContainer {
    a.Set("staticSchema", value)
    return a
}

/**
 * editorSetting
 */
func (a *SwitchContainer) EditorSetting(value interface{}) *SwitchContainer {
    a.Set("editorSetting", value)
    return a
}

/**
 * items
 */
func (a *SwitchContainer) Items(value interface{}) *SwitchContainer {
    a.Set("items", value)
    return a
}

/**
 * className
 */
func (a *SwitchContainer) ClassName(value interface{}) *SwitchContainer {
    a.Set("className", value)
    return a
}

/**
 * visibleOn
 */
func (a *SwitchContainer) VisibleOn(value interface{}) *SwitchContainer {
    a.Set("visibleOn", value)
    return a
}

/**
 * onEvent
 */
func (a *SwitchContainer) OnEvent(value interface{}) *SwitchContainer {
    a.Set("onEvent", value)
    return a
}

/**
 * staticClassName
 */
func (a *SwitchContainer) StaticClassName(value interface{}) *SwitchContainer {
    a.Set("staticClassName", value)
    return a
}

/**
 * useMobileUI
 */
func (a *SwitchContainer) UseMobileUI(value interface{}) *SwitchContainer {
    a.Set("useMobileUI", value)
    return a
}

/**
 * testid
 */
func (a *SwitchContainer) Testid(value interface{}) *SwitchContainer {
    a.Set("testid", value)
    return a
}

/**
 * disabledOn
 */
func (a *SwitchContainer) DisabledOn(value interface{}) *SwitchContainer {
    a.Set("disabledOn", value)
    return a
}

/**
 * hiddenOn
 */
func (a *SwitchContainer) HiddenOn(value interface{}) *SwitchContainer {
    a.Set("hiddenOn", value)
    return a
}

/**
 * id
 */
func (a *SwitchContainer) Id(value interface{}) *SwitchContainer {
    a.Set("id", value)
    return a
}

/**
 * staticOn
 */
func (a *SwitchContainer) StaticOn(value interface{}) *SwitchContainer {
    a.Set("staticOn", value)
    return a
}

/**
 */
func (a *SwitchContainer) Type(value interface{}) *SwitchContainer {
    a.Set("type", value)
    return a
}

/**
 * disabled
 */
func (a *SwitchContainer) Disabled(value interface{}) *SwitchContainer {
    a.Set("disabled", value)
    return a
}

/**
 * hidden
 */
func (a *SwitchContainer) Hidden(value interface{}) *SwitchContainer {
    a.Set("hidden", value)
    return a
}

/**
 * static
 */
func (a *SwitchContainer) Static(value interface{}) *SwitchContainer {
    a.Set("static", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *SwitchContainer) StaticPlaceholder(value interface{}) *SwitchContainer {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *SwitchContainer) StaticLabelClassName(value interface{}) *SwitchContainer {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *SwitchContainer) StaticInputClassName(value interface{}) *SwitchContainer {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * style
 */
func (a *SwitchContainer) Style(value interface{}) *SwitchContainer {
    a.Set("style", value)
    return a
}

/**
 * visible
 */
func (a *SwitchContainer) Visible(value interface{}) *SwitchContainer {
    a.Set("visible", value)
    return a
}
