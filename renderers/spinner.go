package renderers

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

func (a *Spinner) Set(name string, value interface{}) *Spinner {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Spinner) Body(value interface{}) *Spinner {
	a.Set("body", value)
	return a
}

func (a *Spinner) ClassName(value interface{}) *Spinner {
	a.Set("className", value)
	return a
}

func (a *Spinner) Delay(value interface{}) *Spinner {
	a.Set("delay", value)
	return a
}

func (a *Spinner) Disabled(value interface{}) *Spinner {
	a.Set("disabled", value)
	return a
}

func (a *Spinner) DisabledOn(value interface{}) *Spinner {
	a.Set("disabledOn", value)
	return a
}

func (a *Spinner) EditorSetting(value interface{}) *Spinner {
	a.Set("editorSetting", value)
	return a
}

func (a *Spinner) Hidden(value interface{}) *Spinner {
	a.Set("hidden", value)
	return a
}

func (a *Spinner) HiddenOn(value interface{}) *Spinner {
	a.Set("hiddenOn", value)
	return a
}

func (a *Spinner) Icon(value interface{}) *Spinner {
	a.Set("icon", value)
	return a
}

func (a *Spinner) Id(value interface{}) *Spinner {
	a.Set("id", value)
	return a
}

func (a *Spinner) LoadingConfig(value interface{}) *Spinner {
	a.Set("loadingConfig", value)
	return a
}

func (a *Spinner) Mode(value interface{}) *Spinner {
	a.Set("mode", value)
	return a
}

func (a *Spinner) OnEvent(value interface{}) *Spinner {
	a.Set("onEvent", value)
	return a
}

func (a *Spinner) Overlay(value interface{}) *Spinner {
	a.Set("overlay", value)
	return a
}

func (a *Spinner) Show(value interface{}) *Spinner {
	a.Set("show", value)
	return a
}

func (a *Spinner) Size(value interface{}) *Spinner {
	a.Set("size", value)
	return a
}

func (a *Spinner) SpinnerClassName(value interface{}) *Spinner {
	a.Set("spinnerClassName", value)
	return a
}

func (a *Spinner) SpinnerWrapClassName(value interface{}) *Spinner {
	a.Set("spinnerWrapClassName", value)
	return a
}

func (a *Spinner) Static(value interface{}) *Spinner {
	a.Set("static", value)
	return a
}

func (a *Spinner) StaticClassName(value interface{}) *Spinner {
	a.Set("staticClassName", value)
	return a
}

func (a *Spinner) StaticInputClassName(value interface{}) *Spinner {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Spinner) StaticLabelClassName(value interface{}) *Spinner {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Spinner) StaticOn(value interface{}) *Spinner {
	a.Set("staticOn", value)
	return a
}

func (a *Spinner) StaticPlaceholder(value interface{}) *Spinner {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Spinner) StaticSchema(value interface{}) *Spinner {
	a.Set("staticSchema", value)
	return a
}

func (a *Spinner) Style(value interface{}) *Spinner {
	a.Set("style", value)
	return a
}

func (a *Spinner) TestIdBuilder(value interface{}) *Spinner {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Spinner) Testid(value interface{}) *Spinner {
	a.Set("testid", value)
	return a
}

func (a *Spinner) Tip(value interface{}) *Spinner {
	a.Set("tip", value)
	return a
}

func (a *Spinner) TipPlacement(value interface{}) *Spinner {
	a.Set("tipPlacement", value)
	return a
}

func (a *Spinner) Type(value interface{}) *Spinner {
	a.Set("type", value)
	return a
}

func (a *Spinner) UseMobileUI(value interface{}) *Spinner {
	a.Set("useMobileUI", value)
	return a
}

func (a *Spinner) Visible(value interface{}) *Spinner {
	a.Set("visible", value)
	return a
}

func (a *Spinner) VisibleOn(value interface{}) *Spinner {
	a.Set("visibleOn", value)
	return a
}
