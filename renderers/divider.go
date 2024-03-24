package renderers


/**
 * Divider 分割线渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/divider

* @author wcz0
* @version 6.2.2
*/
type Divider struct {
	*BaseRenderer
}

func NewDivider() *Divider {
    a := &Divider{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "divider")
    return a
}
/**
 * 是否禁用
 */
func (a *Divider) Disabled(value string) *Divider {
    a.Set("disabled", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Divider) Id(value string) *Divider {
    a.Set("id", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Divider) EditorSetting(value string) *Divider {
    a.Set("editorSetting", value)
    return a
}

/**
 * 可选值: left | center | right
 */
func (a *Divider) TitlePosition(value string) *Divider {
    a.Set("titlePosition", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Divider) HiddenOn(value string) *Divider {
    a.Set("hiddenOn", value)
    return a
}

/**
 */
func (a *Divider) Rotate(value string) *Divider {
    a.Set("rotate", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Divider) Hidden(value string) *Divider {
    a.Set("hidden", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Divider) OnEvent(value string) *Divider {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Divider) StaticLabelClassName(value string) *Divider {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 可选值: dashed | solid
 */
func (a *Divider) LineStyle(value string) *Divider {
    a.Set("lineStyle", value)
    return a
}

/**
 * 可选值: horizontal | vertical
 */
func (a *Divider) Direction(value string) *Divider {
    a.Set("direction", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Divider) ClassName(value string) *Divider {
    a.Set("className", value)
    return a
}

/**
 */
func (a *Divider) Type(value string) *Divider {
    a.Set("type", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Divider) Static(value string) *Divider {
    a.Set("static", value)
    return a
}

/**
 */
func (a *Divider) Title(value string) *Divider {
    a.Set("title", value)
    return a
}

/**
 * 是否显示
 */
func (a *Divider) Visible(value string) *Divider {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Divider) StaticPlaceholder(value string) *Divider {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Divider) StaticClassName(value string) *Divider {
    a.Set("staticClassName", value)
    return a
}

/**
 */
func (a *Divider) StaticSchema(value string) *Divider {
    a.Set("staticSchema", value)
    return a
}

/**
 */
func (a *Divider) TitleClassName(value string) *Divider {
    a.Set("titleClassName", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Divider) DisabledOn(value string) *Divider {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Divider) VisibleOn(value string) *Divider {
    a.Set("visibleOn", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Divider) StaticOn(value string) *Divider {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Divider) StaticInputClassName(value string) *Divider {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Divider) UseMobileUI(value string) *Divider {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *Divider) Color(value string) *Divider {
    a.Set("color", value)
    return a
}

/**
 * 组件样式
 */
func (a *Divider) Style(value string) *Divider {
    a.Set("style", value)
    return a
}
