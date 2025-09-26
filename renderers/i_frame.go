package renderers


/**
 * IFrame 渲染器 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/iframe

 * @author wcz0
 * @version 6.2.2
 */
type IFrame struct {
	*BaseRenderer
}

func NewIFrame() *IFrame {
    a := &IFrame{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "iframe")
    return a
}


func (a *IFrame) Set(name string, value interface{}) *IFrame {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * staticInputClassName
 */
func (a *IFrame) StaticInputClassName(value interface{}) *IFrame {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * disabledOn
 */
func (a *IFrame) DisabledOn(value interface{}) *IFrame {
    a.Set("disabledOn", value)
    return a
}

/**
 * hiddenOn
 */
func (a *IFrame) HiddenOn(value interface{}) *IFrame {
    a.Set("hiddenOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *IFrame) StaticPlaceholder(value interface{}) *IFrame {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * height
 */
func (a *IFrame) Height(value interface{}) *IFrame {
    a.Set("height", value)
    return a
}

/**
 * name
 */
func (a *IFrame) Name(value interface{}) *IFrame {
    a.Set("name", value)
    return a
}

/**
 * sandbox
 */
func (a *IFrame) Sandbox(value interface{}) *IFrame {
    a.Set("sandbox", value)
    return a
}

/**
 * useMobileUI
 */
func (a *IFrame) UseMobileUI(value interface{}) *IFrame {
    a.Set("useMobileUI", value)
    return a
}

/**
 * width
 */
func (a *IFrame) Width(value interface{}) *IFrame {
    a.Set("width", value)
    return a
}

/**
 * referrerpolicy
 */
func (a *IFrame) Referrerpolicy(value interface{}) *IFrame {
    a.Set("referrerpolicy", value)
    return a
}

/**
 * hidden
 */
func (a *IFrame) Hidden(value interface{}) *IFrame {
    a.Set("hidden", value)
    return a
}

/**
 * style
 */
func (a *IFrame) Style(value interface{}) *IFrame {
    a.Set("style", value)
    return a
}

/**
 * src
 */
func (a *IFrame) Src(value interface{}) *IFrame {
    a.Set("src", value)
    return a
}

/**
 * visible
 */
func (a *IFrame) Visible(value interface{}) *IFrame {
    a.Set("visible", value)
    return a
}

/**
 * visibleOn
 */
func (a *IFrame) VisibleOn(value interface{}) *IFrame {
    a.Set("visibleOn", value)
    return a
}

/**
 * id
 */
func (a *IFrame) Id(value interface{}) *IFrame {
    a.Set("id", value)
    return a
}

/**
 * staticOn
 */
func (a *IFrame) StaticOn(value interface{}) *IFrame {
    a.Set("staticOn", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *IFrame) StaticLabelClassName(value interface{}) *IFrame {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *IFrame) StaticSchema(value interface{}) *IFrame {
    a.Set("staticSchema", value)
    return a
}

/**
 * editorSetting
 */
func (a *IFrame) EditorSetting(value interface{}) *IFrame {
    a.Set("editorSetting", value)
    return a
}

/**
 * staticClassName
 */
func (a *IFrame) StaticClassName(value interface{}) *IFrame {
    a.Set("staticClassName", value)
    return a
}

/**
 */
func (a *IFrame) Type(value interface{}) *IFrame {
    a.Set("type", value)
    return a
}

/**
 * testid
 */
func (a *IFrame) Testid(value interface{}) *IFrame {
    a.Set("testid", value)
    return a
}

/**
 * allow
 */
func (a *IFrame) Allow(value interface{}) *IFrame {
    a.Set("allow", value)
    return a
}

/**
 * className
 */
func (a *IFrame) ClassName(value interface{}) *IFrame {
    a.Set("className", value)
    return a
}

/**
 * disabled
 */
func (a *IFrame) Disabled(value interface{}) *IFrame {
    a.Set("disabled", value)
    return a
}

/**
 * onEvent
 */
func (a *IFrame) OnEvent(value interface{}) *IFrame {
    a.Set("onEvent", value)
    return a
}

/**
 * static
 */
func (a *IFrame) Static(value interface{}) *IFrame {
    a.Set("static", value)
    return a
}
