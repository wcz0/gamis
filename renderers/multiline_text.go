package renderers

type MultilineText struct {
	*BaseRenderer
}

func NewMultilineText() *MultilineText {
	a := &MultilineText{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "multiline-text")
	return a
}

func (a *MultilineText) Set(name string, value interface{}) *MultilineText {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *MultilineText) ClassName(value interface{}) *MultilineText {
	a.Set("className", value)
	return a
}

func (a *MultilineText) CollapseButtonText(value interface{}) *MultilineText {
	a.Set("collapseButtonText", value)
	return a
}

func (a *MultilineText) Disabled(value interface{}) *MultilineText {
	a.Set("disabled", value)
	return a
}

func (a *MultilineText) DisabledOn(value interface{}) *MultilineText {
	a.Set("disabledOn", value)
	return a
}

func (a *MultilineText) EditorSetting(value interface{}) *MultilineText {
	a.Set("editorSetting", value)
	return a
}

func (a *MultilineText) ExpendButtonText(value interface{}) *MultilineText {
	a.Set("expendButtonText", value)
	return a
}

func (a *MultilineText) Hidden(value interface{}) *MultilineText {
	a.Set("hidden", value)
	return a
}

func (a *MultilineText) HiddenOn(value interface{}) *MultilineText {
	a.Set("hiddenOn", value)
	return a
}

func (a *MultilineText) Id(value interface{}) *MultilineText {
	a.Set("id", value)
	return a
}

func (a *MultilineText) MaxRows(value interface{}) *MultilineText {
	a.Set("maxRows", value)
	return a
}

func (a *MultilineText) OnEvent(value interface{}) *MultilineText {
	a.Set("onEvent", value)
	return a
}

func (a *MultilineText) Static(value interface{}) *MultilineText {
	a.Set("static", value)
	return a
}

func (a *MultilineText) StaticClassName(value interface{}) *MultilineText {
	a.Set("staticClassName", value)
	return a
}

func (a *MultilineText) StaticInputClassName(value interface{}) *MultilineText {
	a.Set("staticInputClassName", value)
	return a
}

func (a *MultilineText) StaticLabelClassName(value interface{}) *MultilineText {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *MultilineText) StaticOn(value interface{}) *MultilineText {
	a.Set("staticOn", value)
	return a
}

func (a *MultilineText) StaticPlaceholder(value interface{}) *MultilineText {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *MultilineText) StaticSchema(value interface{}) *MultilineText {
	a.Set("staticSchema", value)
	return a
}

func (a *MultilineText) Style(value interface{}) *MultilineText {
	a.Set("style", value)
	return a
}

func (a *MultilineText) TestIdBuilder(value interface{}) *MultilineText {
	a.Set("testIdBuilder", value)
	return a
}

func (a *MultilineText) Testid(value interface{}) *MultilineText {
	a.Set("testid", value)
	return a
}

func (a *MultilineText) Text(value interface{}) *MultilineText {
	a.Set("text", value)
	return a
}

func (a *MultilineText) Type(value interface{}) *MultilineText {
	a.Set("type", value)
	return a
}

func (a *MultilineText) UseMobileUI(value interface{}) *MultilineText {
	a.Set("useMobileUI", value)
	return a
}

func (a *MultilineText) Visible(value interface{}) *MultilineText {
	a.Set("visible", value)
	return a
}

func (a *MultilineText) VisibleOn(value interface{}) *MultilineText {
	a.Set("visibleOn", value)
	return a
}
