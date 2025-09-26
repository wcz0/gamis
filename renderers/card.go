package renderers

type Card struct {
	*BaseRenderer
}

func NewCard() *Card {
	a := &Card{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "card")
	return a
}

func (a *Card) Set(name string, value interface{}) *Card {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Card) Actions(value interface{}) *Card {
	a.Set("actions", value)
	return a
}

func (a *Card) Body(value interface{}) *Card {
	a.Set("body", value)
	return a
}

func (a *Card) ClassName(value interface{}) *Card {
	a.Set("className", value)
	return a
}

func (a *Card) Disabled(value interface{}) *Card {
	a.Set("disabled", value)
	return a
}

func (a *Card) DisabledOn(value interface{}) *Card {
	a.Set("disabledOn", value)
	return a
}

func (a *Card) EditorSetting(value interface{}) *Card {
	a.Set("editorSetting", value)
	return a
}

func (a *Card) Header(value interface{}) *Card {
	a.Set("header", value)
	return a
}

func (a *Card) Hidden(value interface{}) *Card {
	a.Set("hidden", value)
	return a
}

func (a *Card) HiddenOn(value interface{}) *Card {
	a.Set("hiddenOn", value)
	return a
}

func (a *Card) Id(value interface{}) *Card {
	a.Set("id", value)
	return a
}

func (a *Card) ItemAction(value interface{}) *Card {
	a.Set("itemAction", value)
	return a
}

func (a *Card) Media(value interface{}) *Card {
	a.Set("media", value)
	return a
}

func (a *Card) OnEvent(value interface{}) *Card {
	a.Set("onEvent", value)
	return a
}

func (a *Card) Secondary(value interface{}) *Card {
	a.Set("secondary", value)
	return a
}

func (a *Card) Static(value interface{}) *Card {
	a.Set("static", value)
	return a
}

func (a *Card) StaticClassName(value interface{}) *Card {
	a.Set("staticClassName", value)
	return a
}

func (a *Card) StaticInputClassName(value interface{}) *Card {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Card) StaticLabelClassName(value interface{}) *Card {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Card) StaticOn(value interface{}) *Card {
	a.Set("staticOn", value)
	return a
}

func (a *Card) StaticPlaceholder(value interface{}) *Card {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Card) StaticSchema(value interface{}) *Card {
	a.Set("staticSchema", value)
	return a
}

func (a *Card) Style(value interface{}) *Card {
	a.Set("style", value)
	return a
}

func (a *Card) TestIdBuilder(value interface{}) *Card {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Card) Testid(value interface{}) *Card {
	a.Set("testid", value)
	return a
}

func (a *Card) Toolbar(value interface{}) *Card {
	a.Set("toolbar", value)
	return a
}

func (a *Card) Type(value interface{}) *Card {
	a.Set("type", value)
	return a
}

func (a *Card) UseCardLabel(value interface{}) *Card {
	a.Set("useCardLabel", value)
	return a
}

func (a *Card) UseMobileUI(value interface{}) *Card {
	a.Set("useMobileUI", value)
	return a
}

func (a *Card) Visible(value interface{}) *Card {
	a.Set("visible", value)
	return a
}

func (a *Card) VisibleOn(value interface{}) *Card {
	a.Set("visibleOn", value)
	return a
}
