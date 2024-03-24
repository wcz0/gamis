package renderers


/**
 * Flex 布局 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/flex

 * @author wcz0
 * @version 6.2.2
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
 * 事件动作配置
 */
func (a *Flex) OnEvent(value interface{}) *Flex {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Flex) Static(value interface{}) *Flex {
    a.Set("static", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Flex) StaticOn(value interface{}) *Flex {
    a.Set("staticOn", value)
    return a
}

/**
 */
func (a *Flex) StaticSchema(value interface{}) *Flex {
    a.Set("staticSchema", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Flex) UseMobileUI(value interface{}) *Flex {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 多行情况下的垂直分布
 * 可选值: normal | flex-start | flex-end | center | space-between | space-around | space-evenly | stretch
 */
func (a *Flex) AlignContent(value interface{}) *Flex {
    a.Set("alignContent", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Flex) HiddenOn(value interface{}) *Flex {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Flex) Hidden(value interface{}) *Flex {
    a.Set("hidden", value)
    return a
}

/**
 * 是否显示
 */
func (a *Flex) Visible(value interface{}) *Flex {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Flex) StaticPlaceholder(value interface{}) *Flex {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Flex) StaticLabelClassName(value interface{}) *Flex {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 自定义样式
 */
func (a *Flex) Style(value interface{}) *Flex {
    a.Set("style", value)
    return a
}

/**
 * 水平分布
 * 可选值: start | flex-start | center | end | flex-end | space-around | space-between | space-evenly
 */
func (a *Flex) Justify(value interface{}) *Flex {
    a.Set("justify", value)
    return a
}

/**
 * 方向
 * 可选值: row | column | row-reverse | column-reverse
 */
func (a *Flex) Direction(value interface{}) *Flex {
    a.Set("direction", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Flex) ClassName(value interface{}) *Flex {
    a.Set("className", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Flex) Id(value interface{}) *Flex {
    a.Set("id", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Flex) StaticClassName(value interface{}) *Flex {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Flex) StaticInputClassName(value interface{}) *Flex {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Flex) EditorSetting(value interface{}) *Flex {
    a.Set("editorSetting", value)
    return a
}

/**
 * 指定为 flex 展示类型
 */
func (a *Flex) Type(value interface{}) *Flex {
    a.Set("type", value)
    return a
}

/**
 * 每个 flex 的设置
 */
func (a *Flex) Items(value interface{}) *Flex {
    a.Set("items", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Flex) VisibleOn(value interface{}) *Flex {
    a.Set("visibleOn", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Flex) Disabled(value interface{}) *Flex {
    a.Set("disabled", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Flex) DisabledOn(value interface{}) *Flex {
    a.Set("disabledOn", value)
    return a
}

/**
 * 垂直布局
 * 可选值: stretch | start | flex-start | flex-end | end | center | baseline
 */
func (a *Flex) AlignItems(value interface{}) *Flex {
    a.Set("alignItems", value)
    return a
}
