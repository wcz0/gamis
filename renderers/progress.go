package renderers

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

func (a *Progress) Animate(value interface{}) *Progress {
	a.Set("animate", value)
	return a
}

func (a *Progress) ClassName(value interface{}) *Progress {
	a.Set("className", value)
	return a
}

func (a *Progress) Disabled(value interface{}) *Progress {
	a.Set("disabled", value)
	return a
}

func (a *Progress) DisabledOn(value interface{}) *Progress {
	a.Set("disabledOn", value)
	return a
}

func (a *Progress) EditorSetting(value interface{}) *Progress {
	a.Set("editorSetting", value)
	return a
}

func (a *Progress) GapDegree(value interface{}) *Progress {
	a.Set("gapDegree", value)
	return a
}

func (a *Progress) GapPosition(value interface{}) *Progress {
	a.Set("gapPosition", value)
	return a
}

func (a *Progress) Hidden(value interface{}) *Progress {
	a.Set("hidden", value)
	return a
}

func (a *Progress) HiddenOn(value interface{}) *Progress {
	a.Set("hiddenOn", value)
	return a
}

func (a *Progress) Id(value interface{}) *Progress {
	a.Set("id", value)
	return a
}

func (a *Progress) Map(value interface{}) *Progress {
	a.Set("map", value)
	return a
}

func (a *Progress) Mode(value interface{}) *Progress {
	a.Set("mode", value)
	return a
}

func (a *Progress) Name(value interface{}) *Progress {
	a.Set("name", value)
	return a
}

func (a *Progress) OnEvent(value interface{}) *Progress {
	a.Set("onEvent", value)
	return a
}

func (a *Progress) Placeholder(value interface{}) *Progress {
	a.Set("placeholder", value)
	return a
}

func (a *Progress) ProgressClassName(value interface{}) *Progress {
	a.Set("progressClassName", value)
	return a
}

func (a *Progress) ShowLabel(value interface{}) *Progress {
	a.Set("showLabel", value)
	return a
}

func (a *Progress) ShowThresholdText(value interface{}) *Progress {
	a.Set("showThresholdText", value)
	return a
}

func (a *Progress) Static(value interface{}) *Progress {
	a.Set("static", value)
	return a
}

func (a *Progress) StaticClassName(value interface{}) *Progress {
	a.Set("staticClassName", value)
	return a
}

func (a *Progress) StaticInputClassName(value interface{}) *Progress {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Progress) StaticLabelClassName(value interface{}) *Progress {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Progress) StaticOn(value interface{}) *Progress {
	a.Set("staticOn", value)
	return a
}

func (a *Progress) StaticPlaceholder(value interface{}) *Progress {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Progress) StaticSchema(value interface{}) *Progress {
	a.Set("staticSchema", value)
	return a
}

func (a *Progress) Stripe(value interface{}) *Progress {
	a.Set("stripe", value)
	return a
}

func (a *Progress) StrokeWidth(value interface{}) *Progress {
	a.Set("strokeWidth", value)
	return a
}

func (a *Progress) Style(value interface{}) *Progress {
	a.Set("style", value)
	return a
}

func (a *Progress) TestIdBuilder(value interface{}) *Progress {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Progress) Testid(value interface{}) *Progress {
	a.Set("testid", value)
	return a
}

func (a *Progress) Threshold(value interface{}) *Progress {
	a.Set("threshold", value)
	return a
}

func (a *Progress) Type(value interface{}) *Progress {
	a.Set("type", value)
	return a
}

func (a *Progress) UseMobileUI(value interface{}) *Progress {
	a.Set("useMobileUI", value)
	return a
}

func (a *Progress) Value(value interface{}) *Progress {
	a.Set("value", value)
	return a
}

func (a *Progress) ValueTpl(value interface{}) *Progress {
	a.Set("valueTpl", value)
	return a
}

func (a *Progress) Visible(value interface{}) *Progress {
	a.Set("visible", value)
	return a
}

func (a *Progress) VisibleOn(value interface{}) *Progress {
	a.Set("visibleOn", value)
	return a
}
