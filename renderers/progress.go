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
    return a
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
/**
 * stripe
 */
func (a *Progress) Stripe(value interface{}) *Progress {
    a.Set("stripe", value)
    return a
}

/**
 * visible
 */
func (a *Progress) Visible(value interface{}) *Progress {
    a.Set("visible", value)
    return a
}

/**
 * static
 */
func (a *Progress) Static(value interface{}) *Progress {
    a.Set("static", value)
    return a
}

/**
 * value
 */
func (a *Progress) Value(value interface{}) *Progress {
    a.Set("value", value)
    return a
}

/**
 * threshold
 */
func (a *Progress) Threshold(value interface{}) *Progress {
    a.Set("threshold", value)
    return a
}

/**
 * name
 */
func (a *Progress) Name(value interface{}) *Progress {
    a.Set("name", value)
    return a
}

/**
 * disabled
 */
func (a *Progress) Disabled(value interface{}) *Progress {
    a.Set("disabled", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Progress) StaticLabelClassName(value interface{}) *Progress {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * showLabel
 */
func (a *Progress) ShowLabel(value interface{}) *Progress {
    a.Set("showLabel", value)
    return a
}

/**
 * placeholder
 */
func (a *Progress) Placeholder(value interface{}) *Progress {
    a.Set("placeholder", value)
    return a
}

/**
 * showThresholdText
 */
func (a *Progress) ShowThresholdText(value interface{}) *Progress {
    a.Set("showThresholdText", value)
    return a
}

/**
 * className
 */
func (a *Progress) ClassName(value interface{}) *Progress {
    a.Set("className", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Progress) UseMobileUI(value interface{}) *Progress {
    a.Set("useMobileUI", value)
    return a
}

/**
 * mode
 */
func (a *Progress) Mode(value interface{}) *Progress {
    a.Set("mode", value)
    return a
}

/**
 * map
 */
func (a *Progress) Map(value interface{}) *Progress {
    a.Set("map", value)
    return a
}

/**
 * valueTpl
 */
func (a *Progress) ValueTpl(value interface{}) *Progress {
    a.Set("valueTpl", value)
    return a
}

/**
 * disabledOn
 */
func (a *Progress) DisabledOn(value interface{}) *Progress {
    a.Set("disabledOn", value)
    return a
}

/**
 * hidden
 */
func (a *Progress) Hidden(value interface{}) *Progress {
    a.Set("hidden", value)
    return a
}

/**
 * onEvent
 */
func (a *Progress) OnEvent(value interface{}) *Progress {
    a.Set("onEvent", value)
    return a
}

/**
 * gapDegree
 */
func (a *Progress) GapDegree(value interface{}) *Progress {
    a.Set("gapDegree", value)
    return a
}

/**
 * id
 */
func (a *Progress) Id(value interface{}) *Progress {
    a.Set("id", value)
    return a
}

/**
 * staticClassName
 */
func (a *Progress) StaticClassName(value interface{}) *Progress {
    a.Set("staticClassName", value)
    return a
}

/**
 * style
 */
func (a *Progress) Style(value interface{}) *Progress {
    a.Set("style", value)
    return a
}

/**
 * gapPosition
 */
func (a *Progress) GapPosition(value interface{}) *Progress {
    a.Set("gapPosition", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Progress) HiddenOn(value interface{}) *Progress {
    a.Set("hiddenOn", value)
    return a
}

/**
 * visibleOn
 */
func (a *Progress) VisibleOn(value interface{}) *Progress {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticOn
 */
func (a *Progress) StaticOn(value interface{}) *Progress {
    a.Set("staticOn", value)
    return a
}

/**
 * staticSchema
 */
func (a *Progress) StaticSchema(value interface{}) *Progress {
    a.Set("staticSchema", value)
    return a
}

/**
 * testid
 */
func (a *Progress) Testid(value interface{}) *Progress {
    a.Set("testid", value)
    return a
}

/**
 * animate
 */
func (a *Progress) Animate(value interface{}) *Progress {
    a.Set("animate", value)
    return a
}

/**
 * strokeWidth
 */
func (a *Progress) StrokeWidth(value interface{}) *Progress {
    a.Set("strokeWidth", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Progress) StaticPlaceholder(value interface{}) *Progress {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Progress) StaticInputClassName(value interface{}) *Progress {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * editorSetting
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
 * progressClassName
 */
func (a *Progress) ProgressClassName(value interface{}) *Progress {
    a.Set("progressClassName", value)
    return a
}
