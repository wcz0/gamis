package renderers

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

func (a *Carousel) AlwaysShowArrow(value interface{}) *Carousel {
	a.Set("alwaysShowArrow", value)
	return a
}

func (a *Carousel) Animation(value interface{}) *Carousel {
	a.Set("animation", value)
	return a
}

func (a *Carousel) Auto(value interface{}) *Carousel {
	a.Set("auto", value)
	return a
}

func (a *Carousel) ClassName(value interface{}) *Carousel {
	a.Set("className", value)
	return a
}

func (a *Carousel) Controls(value interface{}) *Carousel {
	a.Set("controls", value)
	return a
}

func (a *Carousel) ControlsTheme(value interface{}) *Carousel {
	a.Set("controlsTheme", value)
	return a
}

func (a *Carousel) Direction(value interface{}) *Carousel {
	a.Set("direction", value)
	return a
}

func (a *Carousel) Disabled(value interface{}) *Carousel {
	a.Set("disabled", value)
	return a
}

func (a *Carousel) DisabledOn(value interface{}) *Carousel {
	a.Set("disabledOn", value)
	return a
}

func (a *Carousel) Duration(value interface{}) *Carousel {
	a.Set("duration", value)
	return a
}

func (a *Carousel) EditorSetting(value interface{}) *Carousel {
	a.Set("editorSetting", value)
	return a
}

func (a *Carousel) Height(value interface{}) *Carousel {
	a.Set("height", value)
	return a
}

func (a *Carousel) Hidden(value interface{}) *Carousel {
	a.Set("hidden", value)
	return a
}

func (a *Carousel) HiddenOn(value interface{}) *Carousel {
	a.Set("hiddenOn", value)
	return a
}

func (a *Carousel) Icons(value interface{}) *Carousel {
	a.Set("icons", value)
	return a
}

func (a *Carousel) Id(value interface{}) *Carousel {
	a.Set("id", value)
	return a
}

func (a *Carousel) Interval(value interface{}) *Carousel {
	a.Set("interval", value)
	return a
}

func (a *Carousel) ItemSchema(value interface{}) *Carousel {
	a.Set("itemSchema", value)
	return a
}

func (a *Carousel) Loop(value interface{}) *Carousel {
	a.Set("loop", value)
	return a
}

func (a *Carousel) MouseEvent(value interface{}) *Carousel {
	a.Set("mouseEvent", value)
	return a
}

func (a *Carousel) Multiple(value interface{}) *Carousel {
	a.Set("multiple", value)
	return a
}

func (a *Carousel) Name(value interface{}) *Carousel {
	a.Set("name", value)
	return a
}

func (a *Carousel) OnEvent(value interface{}) *Carousel {
	a.Set("onEvent", value)
	return a
}

func (a *Carousel) Options(value interface{}) *Carousel {
	a.Set("options", value)
	return a
}

func (a *Carousel) Placeholder(value interface{}) *Carousel {
	a.Set("placeholder", value)
	return a
}

func (a *Carousel) Static(value interface{}) *Carousel {
	a.Set("static", value)
	return a
}

func (a *Carousel) StaticClassName(value interface{}) *Carousel {
	a.Set("staticClassName", value)
	return a
}

func (a *Carousel) StaticInputClassName(value interface{}) *Carousel {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Carousel) StaticLabelClassName(value interface{}) *Carousel {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Carousel) StaticOn(value interface{}) *Carousel {
	a.Set("staticOn", value)
	return a
}

func (a *Carousel) StaticPlaceholder(value interface{}) *Carousel {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Carousel) StaticSchema(value interface{}) *Carousel {
	a.Set("staticSchema", value)
	return a
}

func (a *Carousel) Style(value interface{}) *Carousel {
	a.Set("style", value)
	return a
}

func (a *Carousel) TestIdBuilder(value interface{}) *Carousel {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Carousel) Testid(value interface{}) *Carousel {
	a.Set("testid", value)
	return a
}

func (a *Carousel) ThumbMode(value interface{}) *Carousel {
	a.Set("thumbMode", value)
	return a
}

func (a *Carousel) Type(value interface{}) *Carousel {
	a.Set("type", value)
	return a
}

func (a *Carousel) UseMobileUI(value interface{}) *Carousel {
	a.Set("useMobileUI", value)
	return a
}

func (a *Carousel) Visible(value interface{}) *Carousel {
	a.Set("visible", value)
	return a
}

func (a *Carousel) VisibleOn(value interface{}) *Carousel {
	a.Set("visibleOn", value)
	return a
}

func (a *Carousel) Width(value interface{}) *Carousel {
	a.Set("width", value)
	return a
}
