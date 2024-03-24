package renderers


/**
 * Alert 提示渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/alert

* @author wcz0
* @version 6.2.2
*/
type Alert struct {
	*BaseRenderer
}

func NewAlert() *Alert {
    a := &Alert{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "alert")
    return a
}
/**
 * 静态展示表单项Label类名
 */
func (a *Alert) StaticLabelClassName(value string) *Alert {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 关闭按钮CSS类名
 */
func (a *Alert) CloseButtonClassName(value string) *Alert {
    a.Set("closeButtonClassName", value)
    return a
}

/**
 * 图标CSS类名
 */
func (a *Alert) IconClassName(value string) *Alert {
    a.Set("iconClassName", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Alert) HiddenOn(value string) *Alert {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Alert) UseMobileUI(value string) *Alert {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 提示框标题
 */
func (a *Alert) Title(value string) *Alert {
    a.Set("title", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Alert) StaticOn(value string) *Alert {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Alert) StaticInputClassName(value string) *Alert {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 提示类型
 * 可选值: info | warning | success | danger
 */
func (a *Alert) Level(value string) *Alert {
    a.Set("level", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Alert) ClassName(value string) *Alert {
    a.Set("className", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Alert) DisabledOn(value string) *Alert {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Alert) Hidden(value string) *Alert {
    a.Set("hidden", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Alert) Static(value string) *Alert {
    a.Set("static", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Alert) VisibleOn(value string) *Alert {
    a.Set("visibleOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Alert) Id(value string) *Alert {
    a.Set("id", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Alert) EditorSetting(value string) *Alert {
    a.Set("editorSetting", value)
    return a
}

/**
 * 是否显示ICON
 */
func (a *Alert) ShowIcon(value string) *Alert {
    a.Set("showIcon", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Alert) Disabled(value string) *Alert {
    a.Set("disabled", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Alert) StaticClassName(value string) *Alert {
    a.Set("staticClassName", value)
    return a
}

/**
 * 指定为提示框类型
 */
func (a *Alert) Type(value string) *Alert {
    a.Set("type", value)
    return a
}

/**
 * 是否显示关闭按钮
 */
func (a *Alert) ShowCloseButton(value string) *Alert {
    a.Set("showCloseButton", value)
    return a
}

/**
 * 是否显示
 */
func (a *Alert) Visible(value string) *Alert {
    a.Set("visible", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Alert) OnEvent(value string) *Alert {
    a.Set("onEvent", value)
    return a
}

/**
 */
func (a *Alert) StaticSchema(value string) *Alert {
    a.Set("staticSchema", value)
    return a
}

/**
 * 左侧图标
 */
func (a *Alert) Icon(value string) *Alert {
    a.Set("icon", value)
    return a
}

/**
 * 组件样式
 */
func (a *Alert) Style(value string) *Alert {
    a.Set("style", value)
    return a
}

/**
 * 内容区域
 */
func (a *Alert) Body(value string) *Alert {
    a.Set("body", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Alert) StaticPlaceholder(value string) *Alert {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 操作区域
 */
func (a *Alert) Actions(value string) *Alert {
    a.Set("actions", value)
    return a
}
