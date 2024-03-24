package renderers


/**
 * IFrame 渲染器 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/iframe
 *

*/
type IFrame struct {
	*BaseRenderer
}

func NewIFrame() *IFrame {
    a := &IFrame{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "iframe")
    return a
}
/**
 * 是否隐藏
 */
func (a *IFrame) Hidden(value string) *IFrame {
    a.Set("hidden", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *IFrame) HiddenOn(value string) *IFrame {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否禁用
 */
func (a *IFrame) Disabled(value string) *IFrame {
    a.Set("disabled", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *IFrame) DisabledOn(value string) *IFrame {
    a.Set("disabledOn", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *IFrame) UseMobileUI(value string) *IFrame {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *IFrame) Allow(value string) *IFrame {
    a.Set("allow", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *IFrame) StaticLabelClassName(value string) *IFrame {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *IFrame) StaticSchema(value string) *IFrame {
    a.Set("staticSchema", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *IFrame) Id(value string) *IFrame {
    a.Set("id", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *IFrame) Static(value string) *IFrame {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *IFrame) StaticInputClassName(value string) *IFrame {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *IFrame) Style(value string) *IFrame {
    a.Set("style", value)
    return a
}

/**
 */
func (a *IFrame) Width(value string) *IFrame {
    a.Set("width", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *IFrame) ClassName(value string) *IFrame {
    a.Set("className", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *IFrame) OnEvent(value string) *IFrame {
    a.Set("onEvent", value)
    return a
}

/**
 */
func (a *IFrame) Type(value string) *IFrame {
    a.Set("type", value)
    return a
}

/**
 * 是否显示
 */
func (a *IFrame) Visible(value string) *IFrame {
    a.Set("visible", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *IFrame) VisibleOn(value string) *IFrame {
    a.Set("visibleOn", value)
    return a
}

/**
 * 页面地址
 */
func (a *IFrame) Src(value string) *IFrame {
    a.Set("src", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *IFrame) StaticOn(value string) *IFrame {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *IFrame) StaticClassName(value string) *IFrame {
    a.Set("staticClassName", value)
    return a
}

/**
 * 事件相应，配置后当 iframe 通过 postMessage 发送事件时，可以触发 AMIS 内部的动作。
 */
func (a *IFrame) Events(value string) *IFrame {
    a.Set("events", value)
    return a
}

/**
 */
func (a *IFrame) Height(value string) *IFrame {
    a.Set("height", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *IFrame) StaticPlaceholder(value string) *IFrame {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *IFrame) EditorSetting(value string) *IFrame {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *IFrame) Name(value string) *IFrame {
    a.Set("name", value)
    return a
}

/**
 */
func (a *IFrame) Sandbox(value string) *IFrame {
    a.Set("sandbox", value)
    return a
}

/**
 * 可选值: no-referrer | no-referrer-when-downgrade | origin | origin-when-cross-origin | same-origin | strict-origin | strict-origin-when-cross-origin | unsafe-url
 */
func (a *IFrame) Referrerpolicy(value string) *IFrame {
    a.Set("referrerpolicy", value)
    return a
}
