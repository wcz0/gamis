package renderers

type Each struct {
	*BaseRenderer
}

func NewEach() *Each {
	a := &Each{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "each")
	return a
}

func (a *Each) Set(name string, value interface{}) *Each {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Each) ClassName(value interface{}) *Each {
	a.Set("className", value)
	return a
}

func (a *Each) Disabled(value interface{}) *Each {
	a.Set("disabled", value)
	return a
}

func (a *Each) DisabledOn(value interface{}) *Each {
	a.Set("disabledOn", value)
	return a
}

func (a *Each) EditorSetting(value interface{}) *Each {
	a.Set("editorSetting", value)
	return a
}

func (a *Each) Hidden(value interface{}) *Each {
	a.Set("hidden", value)
	return a
}

func (a *Each) HiddenOn(value interface{}) *Each {
	a.Set("hiddenOn", value)
	return a
}

func (a *Each) Id(value interface{}) *Each {
	a.Set("id", value)
	return a
}

func (a *Each) IndexKeyName(value interface{}) *Each {
	a.Set("indexKeyName", value)
	return a
}

func (a *Each) ItemKeyName(value interface{}) *Each {
	a.Set("itemKeyName", value)
	return a
}

func (a *Each) Items(value interface{}) *Each {
	a.Set("items", value)
	return a
}

func (a *Each) Name(value interface{}) *Each {
	a.Set("name", value)
	return a
}

func (a *Each) OnEvent(value interface{}) *Each {
	a.Set("onEvent", value)
	return a
}

func (a *Each) Placeholder(value interface{}) *Each {
	a.Set("placeholder", value)
	return a
}

func (a *Each) Source(value interface{}) *Each {
	a.Set("source", value)
	return a
}

func (a *Each) Static(value interface{}) *Each {
	a.Set("static", value)
	return a
}

func (a *Each) StaticClassName(value interface{}) *Each {
	a.Set("staticClassName", value)
	return a
}

func (a *Each) StaticInputClassName(value interface{}) *Each {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Each) StaticLabelClassName(value interface{}) *Each {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Each) StaticOn(value interface{}) *Each {
	a.Set("staticOn", value)
	return a
}

func (a *Each) StaticPlaceholder(value interface{}) *Each {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Each) StaticSchema(value interface{}) *Each {
	a.Set("staticSchema", value)
	return a
}

func (a *Each) Style(value interface{}) *Each {
	a.Set("style", value)
	return a
}

func (a *Each) TestIdBuilder(value interface{}) *Each {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Each) Testid(value interface{}) *Each {
	a.Set("testid", value)
	return a
}

func (a *Each) Type(value interface{}) *Each {
	a.Set("type", value)
	return a
}

func (a *Each) UseMobileUI(value interface{}) *Each {
	a.Set("useMobileUI", value)
	return a
}

func (a *Each) Visible(value interface{}) *Each {
	a.Set("visible", value)
	return a
}

func (a *Each) VisibleOn(value interface{}) *Each {
	a.Set("visibleOn", value)
	return a
}
