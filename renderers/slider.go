package renderers

type Slider struct {
	*BaseRenderer
}

func NewSlider() *Slider {
	a := &Slider{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "slider")
	return a
}

func (a *Slider) Set(name string, value interface{}) *Slider {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Slider) Body(value interface{}) *Slider {
	a.Set("body", value)
	return a
}

func (a *Slider) BodyWidth(value interface{}) *Slider {
	a.Set("bodyWidth", value)
	return a
}

func (a *Slider) ClassName(value interface{}) *Slider {
	a.Set("className", value)
	return a
}

func (a *Slider) Disabled(value interface{}) *Slider {
	a.Set("disabled", value)
	return a
}

func (a *Slider) DisabledOn(value interface{}) *Slider {
	a.Set("disabledOn", value)
	return a
}

func (a *Slider) EditorSetting(value interface{}) *Slider {
	a.Set("editorSetting", value)
	return a
}

func (a *Slider) Hidden(value interface{}) *Slider {
	a.Set("hidden", value)
	return a
}

func (a *Slider) HiddenOn(value interface{}) *Slider {
	a.Set("hiddenOn", value)
	return a
}

func (a *Slider) Id(value interface{}) *Slider {
	a.Set("id", value)
	return a
}

func (a *Slider) Left(value interface{}) *Slider {
	a.Set("left", value)
	return a
}

func (a *Slider) OnEvent(value interface{}) *Slider {
	a.Set("onEvent", value)
	return a
}

func (a *Slider) Right(value interface{}) *Slider {
	a.Set("right", value)
	return a
}

func (a *Slider) Static(value interface{}) *Slider {
	a.Set("static", value)
	return a
}

func (a *Slider) StaticClassName(value interface{}) *Slider {
	a.Set("staticClassName", value)
	return a
}

func (a *Slider) StaticInputClassName(value interface{}) *Slider {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Slider) StaticLabelClassName(value interface{}) *Slider {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Slider) StaticOn(value interface{}) *Slider {
	a.Set("staticOn", value)
	return a
}

func (a *Slider) StaticPlaceholder(value interface{}) *Slider {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Slider) StaticSchema(value interface{}) *Slider {
	a.Set("staticSchema", value)
	return a
}

func (a *Slider) Style(value interface{}) *Slider {
	a.Set("style", value)
	return a
}

func (a *Slider) TestIdBuilder(value interface{}) *Slider {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Slider) Testid(value interface{}) *Slider {
	a.Set("testid", value)
	return a
}

func (a *Slider) Type(value interface{}) *Slider {
	a.Set("type", value)
	return a
}

func (a *Slider) UseMobileUI(value interface{}) *Slider {
	a.Set("useMobileUI", value)
	return a
}

func (a *Slider) Visible(value interface{}) *Slider {
	a.Set("visible", value)
	return a
}

func (a *Slider) VisibleOn(value interface{}) *Slider {
	a.Set("visibleOn", value)
	return a
}
