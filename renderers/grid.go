package renderers


/**
 * Grid 格子布局渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/grid

 * @author wcz0
 * @version 6.2.2
 */
type Grid struct {
	*BaseRenderer
}

func NewGrid() *Grid {
    a := &Grid{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "grid")
    return a
}


func (a *Grid) Set(name string, value interface{}) *Grid {
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
func (a *Grid) StaticSchema(value interface{}) *Grid {
    a.Set("staticSchema", value)
    return a
}

/**
 * className
 */
func (a *Grid) ClassName(value interface{}) *Grid {
    a.Set("className", value)
    return a
}

/**
 * static
 */
func (a *Grid) Static(value interface{}) *Grid {
    a.Set("static", value)
    return a
}

/**
 * staticOn
 */
func (a *Grid) StaticOn(value interface{}) *Grid {
    a.Set("staticOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Grid) StaticPlaceholder(value interface{}) *Grid {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticClassName
 */
func (a *Grid) StaticClassName(value interface{}) *Grid {
    a.Set("staticClassName", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Grid) UseMobileUI(value interface{}) *Grid {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *Grid) Type(value interface{}) *Grid {
    a.Set("type", value)
    return a
}

/**
 * testid
 */
func (a *Grid) Testid(value interface{}) *Grid {
    a.Set("testid", value)
    return a
}

/**
 * valign
 */
func (a *Grid) Valign(value interface{}) *Grid {
    a.Set("valign", value)
    return a
}

/**
 * align
 */
func (a *Grid) Align(value interface{}) *Grid {
    a.Set("align", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Grid) HiddenOn(value interface{}) *Grid {
    a.Set("hiddenOn", value)
    return a
}

/**
 * visible
 */
func (a *Grid) Visible(value interface{}) *Grid {
    a.Set("visible", value)
    return a
}

/**
 * id
 */
func (a *Grid) Id(value interface{}) *Grid {
    a.Set("id", value)
    return a
}

/**
 * onEvent
 */
func (a *Grid) OnEvent(value interface{}) *Grid {
    a.Set("onEvent", value)
    return a
}

/**
 * style
 */
func (a *Grid) Style(value interface{}) *Grid {
    a.Set("style", value)
    return a
}

/**
 * editorSetting
 */
func (a *Grid) EditorSetting(value interface{}) *Grid {
    a.Set("editorSetting", value)
    return a
}

/**
 * columns
 */
func (a *Grid) Columns(value interface{}) *Grid {
    a.Set("columns", value)
    return a
}

/**
 * gap
 */
func (a *Grid) Gap(value interface{}) *Grid {
    a.Set("gap", value)
    return a
}

/**
 * disabled
 */
func (a *Grid) Disabled(value interface{}) *Grid {
    a.Set("disabled", value)
    return a
}

/**
 * disabledOn
 */
func (a *Grid) DisabledOn(value interface{}) *Grid {
    a.Set("disabledOn", value)
    return a
}

/**
 * hidden
 */
func (a *Grid) Hidden(value interface{}) *Grid {
    a.Set("hidden", value)
    return a
}

/**
 * visibleOn
 */
func (a *Grid) VisibleOn(value interface{}) *Grid {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Grid) StaticLabelClassName(value interface{}) *Grid {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Grid) StaticInputClassName(value interface{}) *Grid {
    a.Set("staticInputClassName", value)
    return a
}
