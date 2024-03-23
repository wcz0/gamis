package renderers


/**
 * Badge 角标。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/badge
 *

*/
type Badge struct {
	*BaseRenderer
}

func NewBadge() *Badge {
    a := &Badge{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}
/**
 * 大小
 */
func (a *Badge) Size(value string) *Badge {
    a.Set("size", value)
    return a
}

/**
 * 角标位置，相对于position的位置进行偏移
 */
func (a *Badge) Offset(value string) *Badge {
    a.Set("offset", value)
    return a
}

/**
 * 封顶的数字值
 */
func (a *Badge) OverflowCount(value string) *Badge {
    a.Set("overflowCount", value)
    return a
}

/**
 * 动态控制是否显示
 */
func (a *Badge) VisibleOn(value string) *Badge {
    a.Set("visibleOn", value)
    return a
}

/**
 * 是否显示动画
 */
func (a *Badge) Animation(value string) *Badge {
    a.Set("animation", value)
    return a
}

/**
 * 角标的自定义样式
 */
func (a *Badge) Style(value string) *Badge {
    a.Set("style", value)
    return a
}

/**
 */
func (a *Badge) ClassName(value string) *Badge {
    a.Set("className", value)
    return a
}

/**
 * 角标类型
 * 可选值: text | dot | ribbon
 */
func (a *Badge) Mode(value string) *Badge {
    a.Set("mode", value)
    return a
}

/**
 * 角标位置
 * 可选值: top-right | top-left | bottom-right | bottom-left
 */
func (a *Badge) Position(value string) *Badge {
    a.Set("position", value)
    return a
}

/**
 * 提示类型
 */
func (a *Badge) Level(value string) *Badge {
    a.Set("level", value)
    return a
}

/**
 * 文本内容
 */
func (a *Badge) Text(value string) *Badge {
    a.Set("text", value)
    return a
}
