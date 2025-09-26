package renderers


/**
 * Hbox 水平布局渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/hbox

 * @author wcz0
 * @version 6.2.2
 */
type HBox struct {
	*BaseRenderer
}

func NewHBox() *HBox {
    a := &HBox{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "hbox")
    return a
}


func (a *HBox) Set(name string, value interface{}) *HBox {
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
func (a *HBox) Visible(value interface{}) *HBox {
    a.Set("visible", value)
    return a
}

/**
 * hiddenOn
 */
func (a *HBox) HiddenOn(value interface{}) *HBox {
    a.Set("hiddenOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *HBox) StaticPlaceholder(value interface{}) *HBox {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticSchema
 */
func (a *HBox) StaticSchema(value interface{}) *HBox {
    a.Set("staticSchema", value)
    return a
}

/**
 * useMobileUI
 */
func (a *HBox) UseMobileUI(value interface{}) *HBox {
    a.Set("useMobileUI", value)
    return a
}

/**
 * columns
 */
func (a *HBox) Columns(value interface{}) *HBox {
    a.Set("columns", value)
    return a
}

/**
 * disabledOn
 */
func (a *HBox) DisabledOn(value interface{}) *HBox {
    a.Set("disabledOn", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *HBox) StaticInputClassName(value interface{}) *HBox {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * subFormHorizontal
 */
func (a *HBox) SubFormHorizontal(value interface{}) *HBox {
    a.Set("subFormHorizontal", value)
    return a
}

/**
 * valign
 */
func (a *HBox) Valign(value interface{}) *HBox {
    a.Set("valign", value)
    return a
}

/**
 * hidden
 */
func (a *HBox) Hidden(value interface{}) *HBox {
    a.Set("hidden", value)
    return a
}

/**
 * onEvent
 */
func (a *HBox) OnEvent(value interface{}) *HBox {
    a.Set("onEvent", value)
    return a
}

/**
 * staticClassName
 */
func (a *HBox) StaticClassName(value interface{}) *HBox {
    a.Set("staticClassName", value)
    return a
}

/**
 * style
 */
func (a *HBox) Style(value interface{}) *HBox {
    a.Set("style", value)
    return a
}

/**
 */
func (a *HBox) Type(value interface{}) *HBox {
    a.Set("type", value)
    return a
}

/**
 * staticOn
 */
func (a *HBox) StaticOn(value interface{}) *HBox {
    a.Set("staticOn", value)
    return a
}

/**
 * testid
 */
func (a *HBox) Testid(value interface{}) *HBox {
    a.Set("testid", value)
    return a
}

/**
 * align
 */
func (a *HBox) Align(value interface{}) *HBox {
    a.Set("align", value)
    return a
}

/**
 * gap
 */
func (a *HBox) Gap(value interface{}) *HBox {
    a.Set("gap", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *HBox) StaticLabelClassName(value interface{}) *HBox {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * className
 */
func (a *HBox) ClassName(value interface{}) *HBox {
    a.Set("className", value)
    return a
}

/**
 * disabled
 */
func (a *HBox) Disabled(value interface{}) *HBox {
    a.Set("disabled", value)
    return a
}

/**
 * visibleOn
 */
func (a *HBox) VisibleOn(value interface{}) *HBox {
    a.Set("visibleOn", value)
    return a
}

/**
 * id
 */
func (a *HBox) Id(value interface{}) *HBox {
    a.Set("id", value)
    return a
}

/**
 * static
 */
func (a *HBox) Static(value interface{}) *HBox {
    a.Set("static", value)
    return a
}

/**
 * editorSetting
 */
func (a *HBox) EditorSetting(value interface{}) *HBox {
    a.Set("editorSetting", value)
    return a
}

/**
 * subFormMode
 */
func (a *HBox) SubFormMode(value interface{}) *HBox {
    a.Set("subFormMode", value)
    return a
}
