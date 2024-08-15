package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type HBoxColumn struct {
	*BaseRenderer
}

func NewHBoxColumn() *HBoxColumn {
    a := &HBoxColumn{
        BaseRenderer: NewBaseRenderer(),
    }

func (a *HBoxColumn) Set(name string, value interface{}) *HBoxColumn {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}

    return a
}

/**
 * 配置子表单项默认的展示方式。
 * 可选值: normal | inline | horizontal
 */
func (a *HBoxColumn) Mode(value interface{}) *HBoxColumn {
    a.Set("mode", value)
    return a
}

/**
 * 如果是水平排版，这个属性可以细化水平排版的左右宽度占比。
 */
func (a *HBoxColumn) Horizontal(value interface{}) *HBoxColumn {
    a.Set("horizontal", value)
    return a
}

/**
 * 是否显示
 */
func (a *HBoxColumn) Visible(value interface{}) *HBoxColumn {
    a.Set("visible", value)
    return a
}

/**
 * 列上 CSS 类名
 */
func (a *HBoxColumn) Columnclassname(value interface{}) *HBoxColumn {
    a.Set("columnClassName", value)
    return a
}

/**
 * 宽度
 */
func (a *HBoxColumn) Width(value interface{}) *HBoxColumn {
    a.Set("width", value)
    return a
}

/**
 * 高度
 */
func (a *HBoxColumn) Height(value interface{}) *HBoxColumn {
    a.Set("height", value)
    return a
}

/**
 * 其他样式
 */
func (a *HBoxColumn) Style(value interface{}) *HBoxColumn {
    a.Set("style", value)
    return a
}

/**
 * 垂直对齐方式
 * 可选值: top | middle | bottom | between
 */
func (a *HBoxColumn) Valign(value interface{}) *HBoxColumn {
    a.Set("valign", value)
    return a
}

/**
 * 内容区
 */
func (a *HBoxColumn) Body(value interface{}) *HBoxColumn {
    a.Set("body", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *HBoxColumn) Visibleon(value interface{}) *HBoxColumn {
    a.Set("visibleOn", value)
    return a
}
