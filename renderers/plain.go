package renderers


/**
 * Plain 纯文本渲染器 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/plain
 *

*/
type Plain struct {
	*BaseRenderer
}

func NewPlain() *Plain {
    a := &Plain{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "plain")
    return a
}
/**
 * 是否隐藏
 */
func (a *Plain) Hidden(value string) *Plain {
    a.Set("hidden", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Plain) HiddenOn(value string) *Plain {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Plain) Id(value string) *Plain {
    a.Set("id", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Plain) StaticOn(value string) *Plain {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Plain) StaticPlaceholder(value string) *Plain {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Plain) StaticLabelClassName(value string) *Plain {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *Plain) Tpl(value string) *Plain {
    a.Set("tpl", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Plain) DisabledOn(value string) *Plain {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否内联显示？
 */
func (a *Plain) Inline(value string) *Plain {
    a.Set("inline", value)
    return a
}

/**
 * 是否显示
 */
func (a *Plain) Visible(value string) *Plain {
    a.Set("visible", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Plain) VisibleOn(value string) *Plain {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Plain) StaticClassName(value string) *Plain {
    a.Set("staticClassName", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Plain) UseMobileUI(value string) *Plain {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *Plain) Text(value string) *Plain {
    a.Set("text", value)
    return a
}

/**
 * 占位符
 */
func (a *Plain) Placeholder(value string) *Plain {
    a.Set("placeholder", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Plain) ClassName(value string) *Plain {
    a.Set("className", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Plain) Disabled(value string) *Plain {
    a.Set("disabled", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Plain) OnEvent(value string) *Plain {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Plain) Static(value string) *Plain {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Plain) StaticInputClassName(value string) *Plain {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *Plain) StaticSchema(value string) *Plain {
    a.Set("staticSchema", value)
    return a
}

/**
 * 组件样式
 */
func (a *Plain) Style(value string) *Plain {
    a.Set("style", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Plain) EditorSetting(value string) *Plain {
    a.Set("editorSetting", value)
    return a
}

/**
 * 指定为模板渲染器。文档：https://aisuda.bce.baidu.com/amis/zh-CN/docs/concepts/template
 * 可选值: plain | text
 */
func (a *Plain) Type(value string) *Plain {
    a.Set("type", value)
    return a
}
