package renderers


/**
 * 操作栏渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/operation
 *

*/
type Operation struct {
	*BaseRenderer
}

func NewOperation() *Operation {
    a := &Operation{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "operation")
    return a
}
/**
 * 是否禁用表达式
 */
func (a *Operation) DisabledOn(value string) *Operation {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Operation) StaticOn(value string) *Operation {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Operation) StaticClassName(value string) *Operation {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Operation) StaticLabelClassName(value string) *Operation {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Operation) Disabled(value string) *Operation {
    a.Set("disabled", value)
    return a
}

/**
 */
func (a *Operation) StaticSchema(value string) *Operation {
    a.Set("staticSchema", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Operation) UseMobileUI(value string) *Operation {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Operation) HiddenOn(value string) *Operation {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Operation) VisibleOn(value string) *Operation {
    a.Set("visibleOn", value)
    return a
}

/**
 * 设置label
 */
func (a *Operation) Label(value string) *Operation {
    a.Set("label", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Operation) Hidden(value string) *Operation {
    a.Set("hidden", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Operation) Static(value string) *Operation {
    a.Set("static", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Operation) Id(value string) *Operation {
    a.Set("id", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Operation) OnEvent(value string) *Operation {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Operation) StaticPlaceholder(value string) *Operation {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Operation) StaticInputClassName(value string) *Operation {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *Operation) Style(value string) *Operation {
    a.Set("style", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Operation) EditorSetting(value string) *Operation {
    a.Set("editorSetting", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Operation) ClassName(value string) *Operation {
    a.Set("className", value)
    return a
}

/**
 * 是否显示
 */
func (a *Operation) Visible(value string) *Operation {
    a.Set("visible", value)
    return a
}

/**
 */
func (a *Operation) Buttons(value string) *Operation {
    a.Set("buttons", value)
    return a
}

/**
 * 指定为操作栏
 */
func (a *Operation) Type(value string) *Operation {
    a.Set("type", value)
    return a
}

/**
 * 占位符
 */
func (a *Operation) Placeholder(value string) *Operation {
    a.Set("placeholder", value)
    return a
}