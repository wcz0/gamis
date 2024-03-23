package renderers


/**
 * Panel渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/panel
 *

*/
type Panel struct {
	*BaseRenderer
}

func NewPanel() *Panel {
    a := &Panel{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "panel")
    return a
}
/**
 * 容器 css 类名
 */
func (a *Panel) ClassName(value string) *Panel {
    a.Set("className", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Panel) EditorSetting(value string) *Panel {
    a.Set("editorSetting", value)
    return a
}

/**
 * 按钮集合外层类名
 */
func (a *Panel) ActionsClassName(value string) *Panel {
    a.Set("actionsClassName", value)
    return a
}

/**
 * 内容区域
 */
func (a *Panel) Body(value string) *Panel {
    a.Set("body", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Panel) DisabledOn(value string) *Panel {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Panel) HiddenOn(value string) *Panel {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 按钮集合
 */
func (a *Panel) Actions(value string) *Panel {
    a.Set("actions", value)
    return a
}

/**
 * 配置 Body 容器 className
 */
func (a *Panel) BodyClassName(value string) *Panel {
    a.Set("bodyClassName", value)
    return a
}

/**
 * 配置 header 容器 className
 */
func (a *Panel) HeaderClassName(value string) *Panel {
    a.Set("headerClassName", value)
    return a
}

/**
 * 如果是水平排版，这个属性可以细化水平排版的左右宽度占比。
 */
func (a *Panel) SubFormHorizontal(value string) *Panel {
    a.Set("subFormHorizontal", value)
    return a
}

/**
 */
func (a *Panel) StaticSchema(value string) *Panel {
    a.Set("staticSchema", value)
    return a
}

/**
 * 组件样式
 */
func (a *Panel) Style(value string) *Panel {
    a.Set("style", value)
    return a
}

/**
 * Panel 标题
 */
func (a *Panel) Title(value string) *Panel {
    a.Set("title", value)
    return a
}

/**
 * 配置子表单项默认的展示方式。
 * 可选值: normal | inline | horizontal
 */
func (a *Panel) SubFormMode(value string) *Panel {
    a.Set("subFormMode", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Panel) Disabled(value string) *Panel {
    a.Set("disabled", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Panel) VisibleOn(value string) *Panel {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Panel) StaticClassName(value string) *Panel {
    a.Set("staticClassName", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Panel) Hidden(value string) *Panel {
    a.Set("hidden", value)
    return a
}

/**
 * 是否显示
 */
func (a *Panel) Visible(value string) *Panel {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Panel) StaticLabelClassName(value string) *Panel {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 指定为Panel渲染器。
 */
func (a *Panel) Type(value string) *Panel {
    a.Set("type", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Panel) Static(value string) *Panel {
    a.Set("static", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Panel) StaticPlaceholder(value string) *Panel {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Panel) OnEvent(value string) *Panel {
    a.Set("onEvent", value)
    return a
}

/**
 * footer 和 actions 外层 div 类名。
 */
func (a *Panel) FooterWrapClassName(value string) *Panel {
    a.Set("footerWrapClassName", value)
    return a
}

/**
 * 头部内容, 和 title 二选一。
 */
func (a *Panel) Header(value string) *Panel {
    a.Set("header", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Panel) Id(value string) *Panel {
    a.Set("id", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Panel) StaticOn(value string) *Panel {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Panel) StaticInputClassName(value string) *Panel {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Panel) UseMobileUI(value string) *Panel {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 底部内容区域
 */
func (a *Panel) Footer(value string) *Panel {
    a.Set("footer", value)
    return a
}

/**
 * 配置 footer 容器 className
 */
func (a *Panel) FooterClassName(value string) *Panel {
    a.Set("footerClassName", value)
    return a
}

/**
 * 固定底部, 想要把按钮固定在底部的时候配置。
 */
func (a *Panel) AffixFooter(value string) *Panel {
    a.Set("affixFooter", value)
    return a
}
