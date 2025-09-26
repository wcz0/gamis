package renderers


/**
 * 二维布局渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/grid-2d

 * @author wcz0
 * @version 6.2.2
 */
type Grid2D struct {
	*BaseRenderer
}

func NewGrid2D() *Grid2D {
    a := &Grid2D{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "grid-2d")
    return a
}


func (a *Grid2D) Set(name string, value interface{}) *Grid2D {
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
func (a *Grid2D) VisibleOn(value interface{}) *Grid2D {
    a.Set("visibleOn", value)
    return a
}

/**
 * id
 */
func (a *Grid2D) Id(value interface{}) *Grid2D {
    a.Set("id", value)
    return a
}

/**
 * static
 */
func (a *Grid2D) Static(value interface{}) *Grid2D {
    a.Set("static", value)
    return a
}

/**
 * staticClassName
 */
func (a *Grid2D) StaticClassName(value interface{}) *Grid2D {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Grid2D) StaticInputClassName(value interface{}) *Grid2D {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * style
 */
func (a *Grid2D) Style(value interface{}) *Grid2D {
    a.Set("style", value)
    return a
}

/**
 * grids
 */
func (a *Grid2D) Grids(value interface{}) *Grid2D {
    a.Set("grids", value)
    return a
}

/**
 * className
 */
func (a *Grid2D) ClassName(value interface{}) *Grid2D {
    a.Set("className", value)
    return a
}

/**
 * disabled
 */
func (a *Grid2D) Disabled(value interface{}) *Grid2D {
    a.Set("disabled", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Grid2D) StaticPlaceholder(value interface{}) *Grid2D {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * disabledOn
 */
func (a *Grid2D) DisabledOn(value interface{}) *Grid2D {
    a.Set("disabledOn", value)
    return a
}

/**
 * onEvent
 */
func (a *Grid2D) OnEvent(value interface{}) *Grid2D {
    a.Set("onEvent", value)
    return a
}

/**
 * gapRow
 */
func (a *Grid2D) GapRow(value interface{}) *Grid2D {
    a.Set("gapRow", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Grid2D) StaticLabelClassName(value interface{}) *Grid2D {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *Grid2D) StaticSchema(value interface{}) *Grid2D {
    a.Set("staticSchema", value)
    return a
}

/**
 * cols
 */
func (a *Grid2D) Cols(value interface{}) *Grid2D {
    a.Set("cols", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Grid2D) UseMobileUI(value interface{}) *Grid2D {
    a.Set("useMobileUI", value)
    return a
}

/**
 * rowHeight
 */
func (a *Grid2D) RowHeight(value interface{}) *Grid2D {
    a.Set("rowHeight", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Grid2D) HiddenOn(value interface{}) *Grid2D {
    a.Set("hiddenOn", value)
    return a
}

/**
 * visible
 */
func (a *Grid2D) Visible(value interface{}) *Grid2D {
    a.Set("visible", value)
    return a
}

/**
 */
func (a *Grid2D) Type(value interface{}) *Grid2D {
    a.Set("type", value)
    return a
}

/**
 * testid
 */
func (a *Grid2D) Testid(value interface{}) *Grid2D {
    a.Set("testid", value)
    return a
}

/**
 * staticOn
 */
func (a *Grid2D) StaticOn(value interface{}) *Grid2D {
    a.Set("staticOn", value)
    return a
}

/**
 * editorSetting
 */
func (a *Grid2D) EditorSetting(value interface{}) *Grid2D {
    a.Set("editorSetting", value)
    return a
}

/**
 * width
 */
func (a *Grid2D) Width(value interface{}) *Grid2D {
    a.Set("width", value)
    return a
}

/**
 * gap
 */
func (a *Grid2D) Gap(value interface{}) *Grid2D {
    a.Set("gap", value)
    return a
}

/**
 * hidden
 */
func (a *Grid2D) Hidden(value interface{}) *Grid2D {
    a.Set("hidden", value)
    return a
}
