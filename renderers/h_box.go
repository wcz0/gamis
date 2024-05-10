package renderers


/**
 * Hbox 水平布局渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/hbox

 * @author wcz0
 * @version 6.2.2
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
 * 静态展示表单项类名
 */
func (a *HBox) StaticClassName(value interface{}) *HBox {
    a.Set("staticClassName", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *HBox) UseMobileUI(value interface{}) *HBox {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 如果是水平排版，这个属性可以细化水平排版的左右宽度占比。
 */
func (a *HBox) SubFormHorizontal(value interface{}) *HBox {
    a.Set("subFormHorizontal", value)
    return a
}

/**
 * 是否显示
 */
func (a *HBox) Visible(value interface{}) *HBox {
    a.Set("visible", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *HBox) VisibleOn(value interface{}) *HBox {
    a.Set("visibleOn", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *HBox) HiddenOn(value interface{}) *HBox {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 指定为each展示类型
 */
func (a *HBox) Type(value interface{}) *HBox {
    a.Set("type", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *HBox) StaticOn(value interface{}) *HBox {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *HBox) StaticInputClassName(value interface{}) *HBox {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *HBox) Columns(value interface{}) *HBox {
    a.Set("columns", value)
    return a
}

/**
 * 水平间距
 * 可选值: xs | sm | base | none | md | lg
 */
func (a *HBox) Gap(value interface{}) *HBox {
    a.Set("gap", value)
    return a
}

/**
 * 是否禁用
 */
func (a *HBox) Disabled(value interface{}) *HBox {
    a.Set("disabled", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *HBox) Hidden(value interface{}) *HBox {
    a.Set("hidden", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *HBox) Static(value interface{}) *HBox {
    a.Set("static", value)
    return a
}

/**
 * 组件样式
 */
func (a *HBox) Style(value interface{}) *HBox {
    a.Set("style", value)
    return a
}

/**
 * 配置子表单项默认的展示方式。
 * 可选值: normal | inline | horizontal
 */
func (a *HBox) SubFormMode(value interface{}) *HBox {
    a.Set("subFormMode", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *HBox) Id(value interface{}) *HBox {
    a.Set("id", value)
    return a
}

/**
 * 垂直对齐方式
 * 可选值: top | middle | bottom | between
 */
func (a *HBox) Valign(value interface{}) *HBox {
    a.Set("valign", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *HBox) DisabledOn(value interface{}) *HBox {
    a.Set("disabledOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *HBox) StaticPlaceholder(value interface{}) *HBox {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *HBox) StaticLabelClassName(value interface{}) *HBox {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *HBox) ClassName(value interface{}) *HBox {
    a.Set("className", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *HBox) OnEvent(value interface{}) *HBox {
    a.Set("onEvent", value)
    return a
}

/**
 */
func (a *HBox) StaticSchema(value interface{}) *HBox {
    a.Set("staticSchema", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *HBox) EditorSetting(value interface{}) *HBox {
    a.Set("editorSetting", value)
    return a
}

/**
 * 水平对齐方式
 * 可选值: left | right | between | center
 */
func (a *HBox) Align(value interface{}) *HBox {
    a.Set("align", value)
    return a
}
