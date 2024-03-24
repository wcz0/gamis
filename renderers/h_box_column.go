package renderers


/**

*/
type HBoxColumn struct {
	*BaseRenderer
}

func NewHBoxColumn() *HBoxColumn {
    a := &HBoxColumn{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}
/**
 * 垂直对齐方式
 * 可选值: top | middle | bottom | between
 */
func (a *HBoxColumn) Valign(value string) *HBoxColumn {
    a.Set("valign", value)
    return a
}

/**
 * 高度
 */
func (a *HBoxColumn) Height(value string) *HBoxColumn {
    a.Set("height", value)
    return a
}

/**
 * 配置子表单项默认的展示方式。
 * 可选值: normal | inline | horizontal
 */
func (a *HBoxColumn) Mode(value string) *HBoxColumn {
    a.Set("mode", value)
    return a
}

/**
 * 如果是水平排版，这个属性可以细化水平排版的左右宽度占比。
 */
func (a *HBoxColumn) Horizontal(value string) *HBoxColumn {
    a.Set("horizontal", value)
    return a
}

/**
 * 是否显示
 */
func (a *HBoxColumn) Visible(value string) *HBoxColumn {
    a.Set("visible", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *HBoxColumn) VisibleOn(value string) *HBoxColumn {
    a.Set("visibleOn", value)
    return a
}

/**
 * 列上 CSS 类名
 */
func (a *HBoxColumn) ColumnClassName(value string) *HBoxColumn {
    a.Set("columnClassName", value)
    return a
}

/**
 * 宽度
 */
func (a *HBoxColumn) Width(value string) *HBoxColumn {
    a.Set("width", value)
    return a
}

/**
 * 其他样式
 */
func (a *HBoxColumn) Style(value string) *HBoxColumn {
    a.Set("style", value)
    return a
}

/**
 * 内容区
 */
func (a *HBoxColumn) Body(value string) *HBoxColumn {
    a.Set("body", value)
    return a
}
