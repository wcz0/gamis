package renderers


/**
 * Flex 布局 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/flex
 *

*/
type Flex struct {
	*BaseRenderer
}

func NewFlex() *Flex {
    a := &Flex{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "flex")
    return a
}
/**
 * 是否禁用表达式
 */
func (a *Flex) DisabledOn(value string) *Flex {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Flex) Hidden(value string) *Flex {
    a.Set("hidden", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Flex) VisibleOn(value string) *Flex {
    a.Set("visibleOn", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Flex) Static(value string) *Flex {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Flex) StaticLabelClassName(value string) *Flex {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Flex) UseMobileUI(value string) *Flex {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 每个 flex 的设置
 */
func (a *Flex) Items(value string) *Flex {
    a.Set("items", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Flex) Disabled(value string) *Flex {
    a.Set("disabled", value)
    return a
}

/**
 * 自定义样式
 */
func (a *Flex) Style(value string) *Flex {
    a.Set("style", value)
    return a
}

/**
 * 指定为 flex 展示类型
 */
func (a *Flex) Type(value string) *Flex {
    a.Set("type", value)
    return a
}

/**
 * 多行情况下的垂直分布
 * 可选值: normal | flex-start | flex-end | center | space-between | space-around | space-evenly | stretch
 */
func (a *Flex) AlignContent(value string) *Flex {
    a.Set("alignContent", value)
    return a
}

/**
 * 方向
 * 可选值: row | column | row-reverse | column-reverse
 */
func (a *Flex) Direction(value string) *Flex {
    a.Set("direction", value)
    return a
}

/**
 */
func (a *Flex) StaticSchema(value string) *Flex {
    a.Set("staticSchema", value)
    return a
}

/**
 * 是否显示
 */
func (a *Flex) Visible(value string) *Flex {
    a.Set("visible", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Flex) OnEvent(value string) *Flex {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Flex) StaticPlaceholder(value string) *Flex {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Flex) StaticClassName(value string) *Flex {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Flex) StaticInputClassName(value string) *Flex {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 水平分布
 * 可选值: start | flex-start | center | end | flex-end | space-around | space-between | space-evenly
 */
func (a *Flex) Justify(value string) *Flex {
    a.Set("justify", value)
    return a
}

/**
 * 垂直布局
 * 可选值: stretch | start | flex-start | flex-end | end | center | baseline
 */
func (a *Flex) AlignItems(value string) *Flex {
    a.Set("alignItems", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Flex) HiddenOn(value string) *Flex {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Flex) Id(value string) *Flex {
    a.Set("id", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Flex) StaticOn(value string) *Flex {
    a.Set("staticOn", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Flex) EditorSetting(value string) *Flex {
    a.Set("editorSetting", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Flex) ClassName(value string) *Flex {
    a.Set("className", value)
    return a
}
