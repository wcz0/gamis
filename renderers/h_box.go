package renderers


/**
 * Hbox 水平布局渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/hbox
 *

*/
type HBox struct {
	*BaseRenderer
}

func NewHBox() *HBox {
    a := &HBox{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "hbox")
    return a
}
/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *HBox) UseMobileUI(value string) *HBox {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 指定为each展示类型
 */
func (a *HBox) Type(value string) *HBox {
    a.Set("type", value)
    return a
}

/**
 * 如果是水平排版，这个属性可以细化水平排版的左右宽度占比。
 */
func (a *HBox) SubFormHorizontal(value string) *HBox {
    a.Set("subFormHorizontal", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *HBox) DisabledOn(value string) *HBox {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *HBox) VisibleOn(value string) *HBox {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *HBox) StaticPlaceholder(value string) *HBox {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *HBox) HiddenOn(value string) *HBox {
    a.Set("hiddenOn", value)
    return a
}

/**
 */
func (a *HBox) StaticSchema(value string) *HBox {
    a.Set("staticSchema", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *HBox) Static(value string) *HBox {
    a.Set("static", value)
    return a
}

/**
 */
func (a *HBox) Columns(value string) *HBox {
    a.Set("columns", value)
    return a
}

/**
 * 垂直对齐方式
 * 可选值: top | middle | bottom | between
 */
func (a *HBox) Valign(value string) *HBox {
    a.Set("valign", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *HBox) Hidden(value string) *HBox {
    a.Set("hidden", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *HBox) Id(value string) *HBox {
    a.Set("id", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *HBox) OnEvent(value string) *HBox {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *HBox) StaticLabelClassName(value string) *HBox {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *HBox) StaticInputClassName(value string) *HBox {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 水平对齐方式
 * 可选值: left | right | between | center
 */
func (a *HBox) Align(value string) *HBox {
    a.Set("align", value)
    return a
}

/**
 * 是否禁用
 */
func (a *HBox) Disabled(value string) *HBox {
    a.Set("disabled", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *HBox) StaticClassName(value string) *HBox {
    a.Set("staticClassName", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *HBox) EditorSetting(value string) *HBox {
    a.Set("editorSetting", value)
    return a
}

/**
 * 配置子表单项默认的展示方式。
 * 可选值: normal | inline | horizontal
 */
func (a *HBox) SubFormMode(value string) *HBox {
    a.Set("subFormMode", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *HBox) ClassName(value string) *HBox {
    a.Set("className", value)
    return a
}

/**
 * 是否显示
 */
func (a *HBox) Visible(value string) *HBox {
    a.Set("visible", value)
    return a
}

/**
 * 组件样式
 */
func (a *HBox) Style(value string) *HBox {
    a.Set("style", value)
    return a
}

/**
 * 水平间距
 * 可选值: xs | sm | base | none | md | lg
 */
func (a *HBox) Gap(value string) *HBox {
    a.Set("gap", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *HBox) StaticOn(value string) *HBox {
    a.Set("staticOn", value)
    return a
}
