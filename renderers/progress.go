package renderers


/**
 * 进度展示控件。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/progress

 * @author wcz0
 * @version 6.2.2
 */
type Progress struct {
	*BaseRenderer
}

func NewProgress() *Progress {
    a := &Progress{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "progress")
    a.Set("mode", "line")
    return a
}
/**
 * 配置不同的值段，用不同的样式提示用户
 */
func (a *Progress) Map(value interface{}) *Progress {
    a.Set("map", value)
    return a
}

/**
 * 内容的模板函数
 */
func (a *Progress) ValueTpl(value interface{}) *Progress {
    a.Set("valueTpl", value)
    return a
}

/**
 * 阈值
 */
func (a *Progress) Threshold(value interface{}) *Progress {
    a.Set("threshold", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Progress) StaticOn(value interface{}) *Progress {
    a.Set("staticOn", value)
    return a
}

/**
 */
func (a *Progress) StaticSchema(value interface{}) *Progress {
    a.Set("staticSchema", value)
    return a
}

/**
 * 占位符
 */
func (a *Progress) Placeholder(value interface{}) *Progress {
    a.Set("placeholder", value)
    return a
}

/**
 * 仪表盘进度条缺口角度，可取值 0 ~ 295
 */
func (a *Progress) GapDegree(value interface{}) *Progress {
    a.Set("gapDegree", value)
    return a
}

/**
 * 仪表盘进度条缺口位置
 * 可选值: top | bottom | left | right
 */
func (a *Progress) GapPosition(value interface{}) *Progress {
    a.Set("gapPosition", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Progress) HiddenOn(value interface{}) *Progress {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Progress) StaticClassName(value interface{}) *Progress {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Progress) StaticLabelClassName(value interface{}) *Progress {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 进度值
 */
func (a *Progress) Value(value interface{}) *Progress {
    a.Set("value", value)
    return a
}

/**
 * 进度条 CSS 类名
 */
func (a *Progress) ProgressClassName(value interface{}) *Progress {
    a.Set("progressClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *Progress) Style(value interface{}) *Progress {
    a.Set("style", value)
    return a
}

/**
 * 关联字段名
 */
func (a *Progress) Name(value interface{}) *Progress {
    a.Set("name", value)
    return a
}

/**
 * 是否显示动画（只有在开启的时候才能看出来）
 */
func (a *Progress) Animate(value interface{}) *Progress {
    a.Set("animate", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Progress) OnEvent(value interface{}) *Progress {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Progress) Static(value interface{}) *Progress {
    a.Set("static", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Progress) ClassName(value interface{}) *Progress {
    a.Set("className", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Progress) Hidden(value interface{}) *Progress {
    a.Set("hidden", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Progress) VisibleOn(value interface{}) *Progress {
    a.Set("visibleOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Progress) Id(value interface{}) *Progress {
    a.Set("id", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Progress) Disabled(value interface{}) *Progress {
    a.Set("disabled", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Progress) StaticInputClassName(value interface{}) *Progress {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Progress) DisabledOn(value interface{}) *Progress {
    a.Set("disabledOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Progress) StaticPlaceholder(value interface{}) *Progress {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Progress) EditorSetting(value interface{}) *Progress {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *Progress) Type(value interface{}) *Progress {
    a.Set("type", value)
    return a
}

/**
 * 进度条线的宽度
 */
func (a *Progress) StrokeWidth(value interface{}) *Progress {
    a.Set("strokeWidth", value)
    return a
}

/**
 * 是否显示阈值数值
 */
func (a *Progress) ShowThresholdText(value interface{}) *Progress {
    a.Set("showThresholdText", value)
    return a
}

/**
 * 是否显示
 */
func (a *Progress) Visible(value interface{}) *Progress {
    a.Set("visible", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Progress) UseMobileUI(value interface{}) *Progress {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 进度条类型
 * 可选值: line | circle | dashboard
 */
func (a *Progress) Mode(value interface{}) *Progress {
    a.Set("mode", value)
    return a
}

/**
 * 是否显示值
 */
func (a *Progress) ShowLabel(value interface{}) *Progress {
    a.Set("showLabel", value)
    return a
}

/**
 * 是否显示背景间隔
 */
func (a *Progress) Stripe(value interface{}) *Progress {
    a.Set("stripe", value)
    return a
}
