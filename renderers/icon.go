package renderers


/**
 * Icon 图标渲染器 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/icon
 *

*/
type Icon struct {
	*BaseRenderer
}

func NewIcon() *Icon {
    a := &Icon{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "icon")
    return a
}
/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Icon) UseMobileUI(value string) *Icon {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *Icon) Type(value string) *Icon {
    a.Set("type", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Icon) Disabled(value string) *Icon {
    a.Set("disabled", value)
    return a
}

/**
 * 是否显示
 */
func (a *Icon) Visible(value string) *Icon {
    a.Set("visible", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Icon) Static(value string) *Icon {
    a.Set("static", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Icon) StaticOn(value string) *Icon {
    a.Set("staticOn", value)
    return a
}

/**
 */
func (a *Icon) StaticSchema(value string) *Icon {
    a.Set("staticSchema", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Icon) DisabledOn(value string) *Icon {
    a.Set("disabledOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Icon) StaticPlaceholder(value string) *Icon {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Icon) StaticClassName(value string) *Icon {
    a.Set("staticClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *Icon) Style(value string) *Icon {
    a.Set("style", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Icon) StaticInputClassName(value string) *Icon {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Icon) ClassName(value string) *Icon {
    a.Set("className", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Icon) Hidden(value string) *Icon {
    a.Set("hidden", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Icon) VisibleOn(value string) *Icon {
    a.Set("visibleOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Icon) Id(value string) *Icon {
    a.Set("id", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Icon) OnEvent(value string) *Icon {
    a.Set("onEvent", value)
    return a
}

/**
 * 可选值: iconfont | fa | 
 */
func (a *Icon) Vendor(value string) *Icon {
    a.Set("vendor", value)
    return a
}

/**
 * 角标
 */
func (a *Icon) Badge(value string) *Icon {
    a.Set("badge", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Icon) HiddenOn(value string) *Icon {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Icon) StaticLabelClassName(value string) *Icon {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Icon) EditorSetting(value string) *Icon {
    a.Set("editorSetting", value)
    return a
}

/**
 * 按钮类型
 */
func (a *Icon) Icon(value string) *Icon {
    a.Set("icon", value)
    return a
}
