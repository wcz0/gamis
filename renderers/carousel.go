package renderers


/**
 * Carousel 轮播图渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/carousel

 * @author wcz0
 * @version 6.2.2
 */
type Carousel struct {
	*BaseRenderer
}

func NewCarousel() *Carousel {
    a := &Carousel{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "carousel")
    return a
}


func (a *Carousel) Set(name string, value interface{}) *Carousel {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * visible
 */
func (a *Carousel) Visible(value interface{}) *Carousel {
    a.Set("visible", value)
    return a
}

/**
 * loop
 */
func (a *Carousel) Loop(value interface{}) *Carousel {
    a.Set("loop", value)
    return a
}

/**
 * thumbMode
 */
func (a *Carousel) ThumbMode(value interface{}) *Carousel {
    a.Set("thumbMode", value)
    return a
}

/**
 * options
 */
func (a *Carousel) Options(value interface{}) *Carousel {
    a.Set("options", value)
    return a
}

/**
 * icons
 */
func (a *Carousel) Icons(value interface{}) *Carousel {
    a.Set("icons", value)
    return a
}

/**
 * hidden
 */
func (a *Carousel) Hidden(value interface{}) *Carousel {
    a.Set("hidden", value)
    return a
}

/**
 * id
 */
func (a *Carousel) Id(value interface{}) *Carousel {
    a.Set("id", value)
    return a
}

/**
 * mouseEvent
 */
func (a *Carousel) MouseEvent(value interface{}) *Carousel {
    a.Set("mouseEvent", value)
    return a
}

/**
 * interval
 */
func (a *Carousel) Interval(value interface{}) *Carousel {
    a.Set("interval", value)
    return a
}

/**
 * height
 */
func (a *Carousel) Height(value interface{}) *Carousel {
    a.Set("height", value)
    return a
}

/**
 * controlsTheme
 */
func (a *Carousel) ControlsTheme(value interface{}) *Carousel {
    a.Set("controlsTheme", value)
    return a
}

/**
 * multiple
 */
func (a *Carousel) Multiple(value interface{}) *Carousel {
    a.Set("multiple", value)
    return a
}

/**
 * itemSchema
 */
func (a *Carousel) ItemSchema(value interface{}) *Carousel {
    a.Set("itemSchema", value)
    return a
}

/**
 * name
 */
func (a *Carousel) Name(value interface{}) *Carousel {
    a.Set("name", value)
    return a
}

/**
 * staticClassName
 */
func (a *Carousel) StaticClassName(value interface{}) *Carousel {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Carousel) StaticInputClassName(value interface{}) *Carousel {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * style
 */
func (a *Carousel) Style(value interface{}) *Carousel {
    a.Set("style", value)
    return a
}

/**
 * editorSetting
 */
func (a *Carousel) EditorSetting(value interface{}) *Carousel {
    a.Set("editorSetting", value)
    return a
}

/**
 * animation
 */
func (a *Carousel) Animation(value interface{}) *Carousel {
    a.Set("animation", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Carousel) HiddenOn(value interface{}) *Carousel {
    a.Set("hiddenOn", value)
    return a
}

/**
 * static
 */
func (a *Carousel) Static(value interface{}) *Carousel {
    a.Set("static", value)
    return a
}

/**
 */
func (a *Carousel) Type(value interface{}) *Carousel {
    a.Set("type", value)
    return a
}

/**
 * testid
 */
func (a *Carousel) Testid(value interface{}) *Carousel {
    a.Set("testid", value)
    return a
}

/**
 * direction
 */
func (a *Carousel) Direction(value interface{}) *Carousel {
    a.Set("direction", value)
    return a
}

/**
 * disabledOn
 */
func (a *Carousel) DisabledOn(value interface{}) *Carousel {
    a.Set("disabledOn", value)
    return a
}

/**
 * onEvent
 */
func (a *Carousel) OnEvent(value interface{}) *Carousel {
    a.Set("onEvent", value)
    return a
}

/**
 * staticOn
 */
func (a *Carousel) StaticOn(value interface{}) *Carousel {
    a.Set("staticOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Carousel) StaticPlaceholder(value interface{}) *Carousel {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * auto
 */
func (a *Carousel) Auto(value interface{}) *Carousel {
    a.Set("auto", value)
    return a
}

/**
 * alwaysShowArrow
 */
func (a *Carousel) AlwaysShowArrow(value interface{}) *Carousel {
    a.Set("alwaysShowArrow", value)
    return a
}

/**
 * visibleOn
 */
func (a *Carousel) VisibleOn(value interface{}) *Carousel {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Carousel) StaticLabelClassName(value interface{}) *Carousel {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * className
 */
func (a *Carousel) ClassName(value interface{}) *Carousel {
    a.Set("className", value)
    return a
}

/**
 * disabled
 */
func (a *Carousel) Disabled(value interface{}) *Carousel {
    a.Set("disabled", value)
    return a
}

/**
 * staticSchema
 */
func (a *Carousel) StaticSchema(value interface{}) *Carousel {
    a.Set("staticSchema", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Carousel) UseMobileUI(value interface{}) *Carousel {
    a.Set("useMobileUI", value)
    return a
}

/**
 * duration
 */
func (a *Carousel) Duration(value interface{}) *Carousel {
    a.Set("duration", value)
    return a
}

/**
 * width
 */
func (a *Carousel) Width(value interface{}) *Carousel {
    a.Set("width", value)
    return a
}

/**
 * placeholder
 */
func (a *Carousel) Placeholder(value interface{}) *Carousel {
    a.Set("placeholder", value)
    return a
}

/**
 * controls
 */
func (a *Carousel) Controls(value interface{}) *Carousel {
    a.Set("controls", value)
    return a
}
