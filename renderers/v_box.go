package renderers


/**
 * 垂直布局控件 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/vbox

 * @author wcz0
 * @version 6.2.2
 */
type VBox struct {
	*BaseRenderer
}

func NewVBox() *VBox {
    a := &VBox{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "vbox")
    return a
}


func (a *VBox) Set(name string, value interface{}) *VBox {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * className
 */
func (a *VBox) ClassName(value interface{}) *VBox {
    a.Set("className", value)
    return a
}

/**
 * disabled
 */
func (a *VBox) Disabled(value interface{}) *VBox {
    a.Set("disabled", value)
    return a
}

/**
 * hiddenOn
 */
func (a *VBox) HiddenOn(value interface{}) *VBox {
    a.Set("hiddenOn", value)
    return a
}

/**
 * onEvent
 */
func (a *VBox) OnEvent(value interface{}) *VBox {
    a.Set("onEvent", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *VBox) StaticPlaceholder(value interface{}) *VBox {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticClassName
 */
func (a *VBox) StaticClassName(value interface{}) *VBox {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *VBox) StaticLabelClassName(value interface{}) *VBox {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * disabledOn
 */
func (a *VBox) DisabledOn(value interface{}) *VBox {
    a.Set("disabledOn", value)
    return a
}

/**
 * visible
 */
func (a *VBox) Visible(value interface{}) *VBox {
    a.Set("visible", value)
    return a
}

/**
 * static
 */
func (a *VBox) Static(value interface{}) *VBox {
    a.Set("static", value)
    return a
}

/**
 */
func (a *VBox) Type(value interface{}) *VBox {
    a.Set("type", value)
    return a
}

/**
 * hidden
 */
func (a *VBox) Hidden(value interface{}) *VBox {
    a.Set("hidden", value)
    return a
}

/**
 * visibleOn
 */
func (a *VBox) VisibleOn(value interface{}) *VBox {
    a.Set("visibleOn", value)
    return a
}

/**
 * id
 */
func (a *VBox) Id(value interface{}) *VBox {
    a.Set("id", value)
    return a
}

/**
 * staticOn
 */
func (a *VBox) StaticOn(value interface{}) *VBox {
    a.Set("staticOn", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *VBox) StaticInputClassName(value interface{}) *VBox {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * style
 */
func (a *VBox) Style(value interface{}) *VBox {
    a.Set("style", value)
    return a
}

/**
 * editorSetting
 */
func (a *VBox) EditorSetting(value interface{}) *VBox {
    a.Set("editorSetting", value)
    return a
}

/**
 * staticSchema
 */
func (a *VBox) StaticSchema(value interface{}) *VBox {
    a.Set("staticSchema", value)
    return a
}

/**
 * useMobileUI
 */
func (a *VBox) UseMobileUI(value interface{}) *VBox {
    a.Set("useMobileUI", value)
    return a
}

/**
 * testid
 */
func (a *VBox) Testid(value interface{}) *VBox {
    a.Set("testid", value)
    return a
}

/**
 * rows
 */
func (a *VBox) Rows(value interface{}) *VBox {
    a.Set("rows", value)
    return a
}
