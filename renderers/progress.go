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

func (a *Progress) Set(name string, value interface{}) *Progress {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}

    a.Set("mode", "line")
    a.Set("type", "progress")
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Progress) Editorsetting(value interface{}) *Progress {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *Progress) Testidbuilder(value interface{}) *Progress {
    a.Set("testIdBuilder", value)
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
 * 进度条线的宽度
 */
func (a *Progress) Strokewidth(value interface{}) *Progress {
    a.Set("strokeWidth", value)
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
 */
func (a *Progress) Staticschema(value interface{}) *Progress {
    a.Set("staticSchema", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Progress) Usemobileui(value interface{}) *Progress {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 是否显示阈值数值
 */
func (a *Progress) Showthresholdtext(value interface{}) *Progress {
    a.Set("showThresholdText", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Progress) Staticplaceholder(value interface{}) *Progress {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Progress) Staticinputclassname(value interface{}) *Progress {
    a.Set("staticInputClassName", value)
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
 * 仪表盘进度条缺口角度，可取值 0 ~ 295
 */
func (a *Progress) Gapdegree(value interface{}) *Progress {
    a.Set("gapDegree", value)
    return a
}

/**
 * 仪表盘进度条缺口位置
 * 可选值: top | bottom | left | right
 */
func (a *Progress) Gapposition(value interface{}) *Progress {
    a.Set("gapPosition", value)
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
 * 配置不同的值段，用不同的样式提示用户
 */
func (a *Progress) Map(value interface{}) *Progress {
    a.Set("map", value)
    return a
}

/**
 * 是否显示背景间隔
 */
func (a *Progress) Stripe(value interface{}) *Progress {
    a.Set("stripe", value)
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
 * 占位符
 */
func (a *Progress) Placeholder(value interface{}) *Progress {
    a.Set("placeholder", value)
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
 * 是否静态展示表达式
 */
func (a *Progress) Staticon(value interface{}) *Progress {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Progress) Staticlabelclassname(value interface{}) *Progress {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *Progress) Testid(value interface{}) *Progress {
    a.Set("testid", value)
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
 * 内容的模板函数
 */
func (a *Progress) Valuetpl(value interface{}) *Progress {
    a.Set("valueTpl", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Progress) Hiddenon(value interface{}) *Progress {
    a.Set("hiddenOn", value)
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
 * 静态展示表单项类名
 */
func (a *Progress) Staticclassname(value interface{}) *Progress {
    a.Set("staticClassName", value)
    return a
}

/**
 * 是否显示值
 */
func (a *Progress) Showlabel(value interface{}) *Progress {
    a.Set("showLabel", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Progress) Disabledon(value interface{}) *Progress {
    a.Set("disabledOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Progress) Onevent(value interface{}) *Progress {
    a.Set("onEvent", value)
    return a
}

/**
 */
func (a *Progress) Type(value interface{}) *Progress {
    a.Set("type", value)
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
 * 组件唯一 id，主要用于日志采集
 */
func (a *Progress) Id(value interface{}) *Progress {
    a.Set("id", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Progress) Visibleon(value interface{}) *Progress {
    a.Set("visibleOn", value)
    return a
}

/**
 * 进度条 CSS 类名
 */
func (a *Progress) Progressclassname(value interface{}) *Progress {
    a.Set("progressClassName", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Progress) Classname(value interface{}) *Progress {
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
