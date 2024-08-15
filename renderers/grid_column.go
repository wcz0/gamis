package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type GridColumn struct {
	*BaseRenderer
}

func NewGridColumn() *GridColumn {
    a := &GridColumn{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}


func (a *GridColumn) Set(name string, value interface{}) *GridColumn {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 极小屏（<768px）时宽度占比
 */
func (a *GridColumn) Xs(value interface{}) *GridColumn {
    a.Set("xs", value)
    return a
}

/**
 * 中屏时(>=992px)宽度占比
 */
func (a *GridColumn) Md(value interface{}) *GridColumn {
    a.Set("md", value)
    return a
}

/**
 * 垂直对齐方式
 * 可选值: top | middle | bottom | between
 */
func (a *GridColumn) Valign(value interface{}) *GridColumn {
    a.Set("valign", value)
    return a
}

/**
 */
func (a *GridColumn) Body(value interface{}) *GridColumn {
    a.Set("body", value)
    return a
}

/**
 * 样式
 */
func (a *GridColumn) Style(value interface{}) *GridColumn {
    a.Set("style", value)
    return a
}

/**
 */
func (a *GridColumn) WrapperCustomStyle(value interface{}) *GridColumn {
    a.Set("wrapperCustomStyle", value)
    return a
}

/**
 * 组件唯一 id
 */
func (a *GridColumn) Id(value interface{}) *GridColumn {
    a.Set("id", value)
    return a
}

/**
 * 小屏时（>=768px）宽度占比
 */
func (a *GridColumn) Sm(value interface{}) *GridColumn {
    a.Set("sm", value)
    return a
}

/**
 * 大屏时(>=1200px)宽度占比
 */
func (a *GridColumn) Lg(value interface{}) *GridColumn {
    a.Set("lg", value)
    return a
}

/**
 * 配置子表单项默认的展示方式。
 * 可选值: normal | inline | horizontal
 */
func (a *GridColumn) Mode(value interface{}) *GridColumn {
    a.Set("mode", value)
    return a
}

/**
 * 如果是水平排版，这个属性可以细化水平排版的左右宽度占比。
 */
func (a *GridColumn) Horizontal(value interface{}) *GridColumn {
    a.Set("horizontal", value)
    return a
}

/**
 * 列类名
 */
func (a *GridColumn) ColumnClassName(value interface{}) *GridColumn {
    a.Set("columnClassName", value)
    return a
}

/**
 */
func (a *GridColumn) ThemeCss(value interface{}) *GridColumn {
    a.Set("themeCss", value)
    return a
}
