package renderers


/**
 * Audio 音频渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/audio

 * @author wcz0
 * @version 6.2.2
 */
type Audio struct {
	*BaseRenderer
}

func NewAudio() *Audio {
    a := &Audio{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "audio")
    return a
}


func (a *Audio) Set(name string, value interface{}) *Audio {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * visibleOn
 */
func (a *Audio) VisibleOn(value interface{}) *Audio {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Audio) StaticPlaceholder(value interface{}) *Audio {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * inline
 */
func (a *Audio) Inline(value interface{}) *Audio {
    a.Set("inline", value)
    return a
}

/**
 * visible
 */
func (a *Audio) Visible(value interface{}) *Audio {
    a.Set("visible", value)
    return a
}

/**
 * id
 */
func (a *Audio) Id(value interface{}) *Audio {
    a.Set("id", value)
    return a
}

/**
 * style
 */
func (a *Audio) Style(value interface{}) *Audio {
    a.Set("style", value)
    return a
}

/**
 */
func (a *Audio) Type(value interface{}) *Audio {
    a.Set("type", value)
    return a
}

/**
 * controls
 */
func (a *Audio) Controls(value interface{}) *Audio {
    a.Set("controls", value)
    return a
}

/**
 * className
 */
func (a *Audio) ClassName(value interface{}) *Audio {
    a.Set("className", value)
    return a
}

/**
 * src
 */
func (a *Audio) Src(value interface{}) *Audio {
    a.Set("src", value)
    return a
}

/**
 * disabledOn
 */
func (a *Audio) DisabledOn(value interface{}) *Audio {
    a.Set("disabledOn", value)
    return a
}

/**
 * staticOn
 */
func (a *Audio) StaticOn(value interface{}) *Audio {
    a.Set("staticOn", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Audio) StaticInputClassName(value interface{}) *Audio {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * autoPlay
 */
func (a *Audio) AutoPlay(value interface{}) *Audio {
    a.Set("autoPlay", value)
    return a
}

/**
 * rates
 */
func (a *Audio) Rates(value interface{}) *Audio {
    a.Set("rates", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Audio) UseMobileUI(value interface{}) *Audio {
    a.Set("useMobileUI", value)
    return a
}

/**
 * static
 */
func (a *Audio) Static(value interface{}) *Audio {
    a.Set("static", value)
    return a
}

/**
 * staticClassName
 */
func (a *Audio) StaticClassName(value interface{}) *Audio {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *Audio) StaticSchema(value interface{}) *Audio {
    a.Set("staticSchema", value)
    return a
}

/**
 * testid
 */
func (a *Audio) Testid(value interface{}) *Audio {
    a.Set("testid", value)
    return a
}

/**
 * onEvent
 */
func (a *Audio) OnEvent(value interface{}) *Audio {
    a.Set("onEvent", value)
    return a
}

/**
 * loop
 */
func (a *Audio) Loop(value interface{}) *Audio {
    a.Set("loop", value)
    return a
}

/**
 * disabled
 */
func (a *Audio) Disabled(value interface{}) *Audio {
    a.Set("disabled", value)
    return a
}

/**
 * hidden
 */
func (a *Audio) Hidden(value interface{}) *Audio {
    a.Set("hidden", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Audio) HiddenOn(value interface{}) *Audio {
    a.Set("hiddenOn", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Audio) StaticLabelClassName(value interface{}) *Audio {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * editorSetting
 */
func (a *Audio) EditorSetting(value interface{}) *Audio {
    a.Set("editorSetting", value)
    return a
}
