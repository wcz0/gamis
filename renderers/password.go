package renderers


/**
 * Password

* @author wcz0
* @version 6.2.2
*/
type Password struct {
	*BaseRenderer
}

func NewPassword() *Password {
    a := &Password{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "password")
    return a
}
/**
 * 静态展示表单项类名
 */
func (a *Password) StaticClassName(value string) *Password {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Password) StaticLabelClassName(value string) *Password {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *Password) Style(value string) *Password {
    a.Set("style", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Password) Hidden(value string) *Password {
    a.Set("hidden", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Password) HiddenOn(value string) *Password {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *Password) Visible(value string) *Password {
    a.Set("visible", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Password) OnEvent(value string) *Password {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Password) StaticPlaceholder(value string) *Password {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Password) VisibleOn(value string) *Password {
    a.Set("visibleOn", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Password) StaticOn(value string) *Password {
    a.Set("staticOn", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Password) UseMobileUI(value string) *Password {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 打码模式的文本
 */
func (a *Password) MosaicText(value string) *Password {
    a.Set("mosaicText", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Password) EditorSetting(value string) *Password {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *Password) Type(value string) *Password {
    a.Set("type", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Password) ClassName(value string) *Password {
    a.Set("className", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Password) Disabled(value string) *Password {
    a.Set("disabled", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Password) DisabledOn(value string) *Password {
    a.Set("disabledOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Password) Id(value string) *Password {
    a.Set("id", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Password) Static(value string) *Password {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Password) StaticInputClassName(value string) *Password {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *Password) StaticSchema(value string) *Password {
    a.Set("staticSchema", value)
    return a
}
