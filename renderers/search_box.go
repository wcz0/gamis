package renderers


/**
 * 搜索框渲染器
 *

*/
type SearchBox struct {
	*BaseRenderer
}

func NewSearchBox() *SearchBox {
    a := &SearchBox{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "search-box")
    return a
}
/**
 * 是否处于加载状态
 */
func (a *SearchBox) Loading(value string) *SearchBox {
    a.Set("loading", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *SearchBox) StaticOn(value string) *SearchBox {
    a.Set("staticOn", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *SearchBox) EditorSetting(value string) *SearchBox {
    a.Set("editorSetting", value)
    return a
}

/**
 * 是否为加强样式
 */
func (a *SearchBox) Enhance(value string) *SearchBox {
    a.Set("enhance", value)
    return a
}

/**
 */
func (a *SearchBox) StaticSchema(value string) *SearchBox {
    a.Set("staticSchema", value)
    return a
}

/**
 * 是否显示
 */
func (a *SearchBox) Visible(value string) *SearchBox {
    a.Set("visible", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *SearchBox) OnEvent(value string) *SearchBox {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *SearchBox) StaticLabelClassName(value string) *SearchBox {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *SearchBox) Style(value string) *SearchBox {
    a.Set("style", value)
    return a
}

/**
 * 是否立马搜索。
 */
func (a *SearchBox) SearchImediately(value string) *SearchBox {
    a.Set("searchImediately", value)
    return a
}

/**
 * 外层 css 类名
 */
func (a *SearchBox) ClassName(value string) *SearchBox {
    a.Set("className", value)
    return a
}

/**
 * 是否禁用
 */
func (a *SearchBox) Disabled(value string) *SearchBox {
    a.Set("disabled", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *SearchBox) StaticClassName(value string) *SearchBox {
    a.Set("staticClassName", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *SearchBox) VisibleOn(value string) *SearchBox {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *SearchBox) StaticPlaceholder(value string) *SearchBox {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *SearchBox) Id(value string) *SearchBox {
    a.Set("id", value)
    return a
}

/**
 * 关键字名字。
 */
func (a *SearchBox) Name(value string) *SearchBox {
    a.Set("name", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *SearchBox) StaticInputClassName(value string) *SearchBox {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *SearchBox) UseMobileUI(value string) *SearchBox {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 指定为搜索框。文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/search-box
 */
func (a *SearchBox) Type(value string) *SearchBox {
    a.Set("type", value)
    return a
}

/**
 * 占位符
 */
func (a *SearchBox) Placeholder(value string) *SearchBox {
    a.Set("placeholder", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *SearchBox) Hidden(value string) *SearchBox {
    a.Set("hidden", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *SearchBox) HiddenOn(value string) *SearchBox {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *SearchBox) Static(value string) *SearchBox {
    a.Set("static", value)
    return a
}

/**
 * 是否开启清空内容后立即重新搜索
 */
func (a *SearchBox) ClearAndSubmit(value string) *SearchBox {
    a.Set("clearAndSubmit", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *SearchBox) DisabledOn(value string) *SearchBox {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否为 Mini 样式。
 */
func (a *SearchBox) Mini(value string) *SearchBox {
    a.Set("mini", value)
    return a
}

/**
 * 是否可清除
 */
func (a *SearchBox) Clearable(value string) *SearchBox {
    a.Set("clearable", value)
    return a
}
