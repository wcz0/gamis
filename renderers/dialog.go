package renderers


/**
 * Dialog 弹框渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/dialog
 *

*/
type Dialog struct {
	*BaseRenderer
}

func NewDialog() *Dialog {
    a := &Dialog{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "dialog")
    return a
}
/**
 * 请通过配置 title 设置标题
 */
func (a *Dialog) Title(value string) *Dialog {
    a.Set("title", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Dialog) Disabled(value string) *Dialog {
    a.Set("disabled", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Dialog) UseMobileUI(value string) *Dialog {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *Dialog) Type(value string) *Dialog {
    a.Set("type", value)
    return a
}

/**
 * 是否支持点其它区域关闭 Dialog
 */
func (a *Dialog) CloseOnOutside(value string) *Dialog {
    a.Set("closeOnOutside", value)
    return a
}

/**
 * Dialog 大小
 * 可选值: xs | sm | md | lg | xl | full
 */
func (a *Dialog) Size(value string) *Dialog {
    a.Set("size", value)
    return a
}

/**
 */
func (a *Dialog) Footer(value string) *Dialog {
    a.Set("footer", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Dialog) ClassName(value string) *Dialog {
    a.Set("className", value)
    return a
}

/**
 * 是否显示
 */
func (a *Dialog) Visible(value string) *Dialog {
    a.Set("visible", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Dialog) StaticOn(value string) *Dialog {
    a.Set("staticOn", value)
    return a
}

/**
 * 可拖拽
 */
func (a *Dialog) Draggable(value string) *Dialog {
    a.Set("draggable", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Dialog) DisabledOn(value string) *Dialog {
    a.Set("disabledOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Dialog) StaticInputClassName(value string) *Dialog {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 是否支持按 ESC 关闭 Dialog
 */
func (a *Dialog) CloseOnEsc(value string) *Dialog {
    a.Set("closeOnEsc", value)
    return a
}

/**
 * 是否显示错误信息
 */
func (a *Dialog) ShowErrorMsg(value string) *Dialog {
    a.Set("showErrorMsg", value)
    return a
}

/**
 */
func (a *Dialog) HeaderClassName(value string) *Dialog {
    a.Set("headerClassName", value)
    return a
}

/**
 * 是否显示 spinner
 */
func (a *Dialog) ShowLoading(value string) *Dialog {
    a.Set("showLoading", value)
    return a
}

/**
 * 弹框类型 confirm 确认弹框
 */
func (a *Dialog) DialogType(value string) *Dialog {
    a.Set("dialogType", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Dialog) Hidden(value string) *Dialog {
    a.Set("hidden", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Dialog) VisibleOn(value string) *Dialog {
    a.Set("visibleOn", value)
    return a
}

/**
 */
func (a *Dialog) StaticSchema(value string) *Dialog {
    a.Set("staticSchema", value)
    return a
}

/**
 * 默认不用填写，自动会创建确认和取消按钮。
 */
func (a *Dialog) Actions(value string) *Dialog {
    a.Set("actions", value)
    return a
}

/**
 * Dialog 高度
 */
func (a *Dialog) Height(value string) *Dialog {
    a.Set("height", value)
    return a
}

/**
 * 是否显示关闭按钮
 */
func (a *Dialog) ShowCloseButton(value string) *Dialog {
    a.Set("showCloseButton", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Dialog) OnEvent(value string) *Dialog {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Dialog) StaticClassName(value string) *Dialog {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Dialog) StaticLabelClassName(value string) *Dialog {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *Dialog) Name(value string) *Dialog {
    a.Set("name", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Dialog) HiddenOn(value string) *Dialog {
    a.Set("hiddenOn", value)
    return a
}

/**
 */
func (a *Dialog) Testid(value string) *Dialog {
    a.Set("testid", value)
    return a
}

/**
 * 是否显示蒙层
 */
func (a *Dialog) Overlay(value string) *Dialog {
    a.Set("overlay", value)
    return a
}

/**
 * Dialog 宽度
 */
func (a *Dialog) Width(value string) *Dialog {
    a.Set("width", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Dialog) Id(value string) *Dialog {
    a.Set("id", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Dialog) Static(value string) *Dialog {
    a.Set("static", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Dialog) StaticPlaceholder(value string) *Dialog {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 配置 Body 容器 className
 */
func (a *Dialog) BodyClassName(value string) *Dialog {
    a.Set("bodyClassName", value)
    return a
}

/**
 * 影响自动生成的按钮，如果自己配置了按钮这个配置无效。
 */
func (a *Dialog) Confirm(value string) *Dialog {
    a.Set("confirm", value)
    return a
}

/**
 * 组件样式
 */
func (a *Dialog) Style(value string) *Dialog {
    a.Set("style", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Dialog) EditorSetting(value string) *Dialog {
    a.Set("editorSetting", value)
    return a
}

/**
 * 内容区域
 */
func (a *Dialog) Body(value string) *Dialog {
    a.Set("body", value)
    return a
}

/**
 */
func (a *Dialog) Header(value string) *Dialog {
    a.Set("header", value)
    return a
}
