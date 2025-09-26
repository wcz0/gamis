package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type NavOverflow struct {
	*BaseRenderer
}

func NewNavOverflow() *NavOverflow {
    a := &NavOverflow{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}


func (a *NavOverflow) Set(name string, value interface{}) *NavOverflow {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 菜单触发按钮CSS类名
 */
func (a *NavOverflow) OverflowClassName(value interface{}) *NavOverflow {
    a.Set("overflowClassName", value)
    return a
}

/**
 * Popover浮层CSS类名
 */
func (a *NavOverflow) OverflowPopoverClassName(value interface{}) *NavOverflow {
    a.Set("overflowPopoverClassName", value)
    return a
}

/**
 * 菜单外层CSS类名
 */
func (a *NavOverflow) OverflowListClassName(value interface{}) *NavOverflow {
    a.Set("overflowListClassName", value)
    return a
}

/**
 * 导航横向布局时，开启开启响应式收纳后最大可显示数量，超出此数量的导航将被收纳到下拉菜单中
 */
func (a *NavOverflow) MaxVisibleCount(value interface{}) *NavOverflow {
    a.Set("maxVisibleCount", value)
    return a
}

/**
 * 包裹导航的外层标签名，可以使用其他标签渲染
 */
func (a *NavOverflow) WrapperComponent(value interface{}) *NavOverflow {
    a.Set("wrapperComponent", value)
    return a
}

/**
 * 导航列表后缀节点
 */
func (a *NavOverflow) OverflowSuffix(value interface{}) *NavOverflow {
    a.Set("overflowSuffix", value)
    return a
}

/**
 * 自定义样式
 */
func (a *NavOverflow) Style(value interface{}) *NavOverflow {
    a.Set("style", value)
    return a
}

/**
 * 导航超出后响应式收纳方案。
 * 可选值: popup | swipe
 */
func (a *NavOverflow) Mode(value interface{}) *NavOverflow {
    a.Set("mode", value)
    return a
}

/**
 * 是否开启响应式收纳
 */
func (a *NavOverflow) Enable(value interface{}) *NavOverflow {
    a.Set("enable", value)
    return a
}

/**
 * 菜单触发按钮的文字
 */
func (a *NavOverflow) OverflowLabel(value interface{}) *NavOverflow {
    a.Set("overflowLabel", value)
    return a
}

/**
 * 菜单触发按钮的图标
 */
func (a *NavOverflow) OverflowIndicator(value interface{}) *NavOverflow {
    a.Set("overflowIndicator", value)
    return a
}

/**
 * 导航项目宽度
 */
func (a *NavOverflow) ItemWidth(value interface{}) *NavOverflow {
    a.Set("itemWidth", value)
    return a
}
