package renderers

type GridNav struct {
    *BaseRenderer
}

func NewGridNav() *GridNav {
    g := &GridNav{
        BaseRenderer: NewBaseRenderer(),
    }
    g.Set("type", "grid-nav")
    return g
}

/**
 * 是否显示列表项边框
 */
func (g *GridNav) Border(value bool) *GridNav {
    g.Set("border", value)
    return g
}

/**
 * 是否将列表项内容居中显示
 */
func (g *GridNav) Center(value bool) *GridNav {
    g.Set("center", value)
    return g
}

/**
 * 外层 CSS 类名
 */
func (g *GridNav) ClassName(value string) *GridNav {
    g.Set("className", value)
    return g
}

/**
 * 列数
 */
func (g *GridNav) ColumnNum(value string) *GridNav {
    g.Set("columnNum", value)
    return g
}

/**
 * 列表项内容排列的方向，可选值为 horizontal 、vertical
 */
func (g *GridNav) Direction(value string) *GridNav {
    g.Set("direction", value)
    return g
}

/**
 * 列表项之间的间距，默认单位为px
 */
func (g *GridNav) Gutter(value string) *GridNav {
    g.Set("gutter", value)
    return g
}

/**
 * 图标宽度占比，单位%
 */
func (g *GridNav) IconRatio(value string) *GridNav {
    g.Set("iconRatio", value)
    return g
}

/**
 * 列表项 css 类名
 */
func (g *GridNav) ItemClassName(value string) *GridNav {
    g.Set("itemClassName", value)
    return g
}

/**
 * 列表项图标
 */
func (g *GridNav) Options(value string) *GridNav {
    g.Set("options", value)
    return g
}

/**
 * 是否调换图标和文本的位置
 */
func (g *GridNav) Reverse(value bool) *GridNav {
    g.Set("reverse", value)
    return g
}

/**
 * 数据源
 */
func (g *GridNav) Source(value string) *GridNav {
    g.Set("source", value)
    return g
}

/**
 * 是否将列表项固定为正方形
 */
func (g *GridNav) Square(value bool) *GridNav {
    g.Set("square", value)
    return g
}

/**
 * 指定为 grid-nav 渲染器。
 */
func (g *GridNav) Type(value string) *GridNav {
    g.Set("type", value)
    return g
}

/**
 * 图片数组
 */
func (g *GridNav) Value(value string) *GridNav {
    g.Set("value", value)
    return g
}
