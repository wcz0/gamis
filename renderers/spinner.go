package renderers


/**

*/
type Spinner struct {
	*BaseRenderer
}

func NewSpinner() *Spinner {
    a := &Spinner{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "spinner")
    return a
}
/**
 * 是否隐藏
 */
func (a *Spinner) Hidden(value string) *Spinner {
    a.Set("hidden", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Spinner) StaticOn(value string) *Spinner {
    a.Set("staticOn", value)
    return a
}

/**
 * 控制Spinner显示与隐藏
 */
func (a *Spinner) Show(value string) *Spinner {
    a.Set("show", value)
    return a
}

/**
 * 自定义icon
 */
func (a *Spinner) Icon(value string) *Spinner {
    a.Set("icon", value)
    return a
}

/**
 * 作为容器使用时内容
 */
func (a *Spinner) Body(value string) *Spinner {
    a.Set("body", value)
    return a
}

/**
 * 是否显示
 */
func (a *Spinner) Visible(value string) *Spinner {
    a.Set("visible", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Spinner) OnEvent(value string) *Spinner {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Spinner) Static(value string) *Spinner {
    a.Set("static", value)
    return a
}

/**
 * 作为容器使用时最外层元素的class
 */
func (a *Spinner) SpinnerWrapClassName(value string) *Spinner {
    a.Set("spinnerWrapClassName", value)
    return a
}

/**
 * spinner文案位置
 * 可选值: top | right | bottom | left
 */
func (a *Spinner) TipPlacement(value string) *Spinner {
    a.Set("tipPlacement", value)
    return a
}

/**
 * 延迟显示
 */
func (a *Spinner) Delay(value string) *Spinner {
    a.Set("delay", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Spinner) DisabledOn(value string) *Spinner {
    a.Set("disabledOn", value)
    return a
}

/**
 */
func (a *Spinner) LoadingConfig(value string) *Spinner {
    a.Set("loadingConfig", value)
    return a
}

/**
 */
func (a *Spinner) StaticSchema(value string) *Spinner {
    a.Set("staticSchema", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Spinner) EditorSetting(value string) *Spinner {
    a.Set("editorSetting", value)
    return a
}

/**
 * spin图标位置包裹元素的自定义class
 */
func (a *Spinner) SpinnerClassName(value string) *Spinner {
    a.Set("spinnerClassName", value)
    return a
}

/**
 * spinner文案
 */
func (a *Spinner) Tip(value string) *Spinner {
    a.Set("tip", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Spinner) HiddenOn(value string) *Spinner {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Spinner) VisibleOn(value string) *Spinner {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Spinner) StaticInputClassName(value string) *Spinner {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *Spinner) Style(value string) *Spinner {
    a.Set("style", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Spinner) Disabled(value string) *Spinner {
    a.Set("disabled", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Spinner) StaticClassName(value string) *Spinner {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Spinner) StaticLabelClassName(value string) *Spinner {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *Spinner) Mode(value string) *Spinner {
    a.Set("mode", value)
    return a
}

/**
 * 是否显示遮罩层
 */
func (a *Spinner) Overlay(value string) *Spinner {
    a.Set("overlay", value)
    return a
}

/**
 * 自定义spinner的class
 */
func (a *Spinner) ClassName(value string) *Spinner {
    a.Set("className", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Spinner) Id(value string) *Spinner {
    a.Set("id", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Spinner) StaticPlaceholder(value string) *Spinner {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Spinner) UseMobileUI(value string) *Spinner {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 组件类型
 */
func (a *Spinner) Type(value string) *Spinner {
    a.Set("type", value)
    return a
}

/**
 * spinner Icon 大小
 * 可选值: sm | lg | 
 */
func (a *Spinner) Size(value string) *Spinner {
    a.Set("size", value)
    return a
}
