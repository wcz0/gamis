package renderers

type Mapping struct {
	*BaseRenderer
}

func NewMapping() *Mapping {
	a := &Mapping{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "map")
	return a
}

func (a *Mapping) Set(name string, value interface{}) *Mapping {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Mapping) ClassName(value interface{}) *Mapping {
	a.Set("className", value)
	return a
}

func (a *Mapping) Disabled(value interface{}) *Mapping {
	a.Set("disabled", value)
	return a
}

func (a *Mapping) DisabledOn(value interface{}) *Mapping {
	a.Set("disabledOn", value)
	return a
}

func (a *Mapping) EditorSetting(value interface{}) *Mapping {
	a.Set("editorSetting", value)
	return a
}

func (a *Mapping) Hidden(value interface{}) *Mapping {
	a.Set("hidden", value)
	return a
}

func (a *Mapping) HiddenOn(value interface{}) *Mapping {
	a.Set("hiddenOn", value)
	return a
}

func (a *Mapping) Id(value interface{}) *Mapping {
	a.Set("id", value)
	return a
}

func (a *Mapping) ItemSchema(value interface{}) *Mapping {
	a.Set("itemSchema", value)
	return a
}

func (a *Mapping) LabelField(value interface{}) *Mapping {
	a.Set("labelField", value)
	return a
}

func (a *Mapping) Map(value interface{}) *Mapping {
	a.Set("map", value)
	return a
}

func (a *Mapping) Name(value interface{}) *Mapping {
	a.Set("name", value)
	return a
}

func (a *Mapping) OnEvent(value interface{}) *Mapping {
	a.Set("onEvent", value)
	return a
}

func (a *Mapping) Placeholder(value interface{}) *Mapping {
	a.Set("placeholder", value)
	return a
}

func (a *Mapping) Source(value interface{}) *Mapping {
	a.Set("source", value)
	return a
}

func (a *Mapping) Static(value interface{}) *Mapping {
	a.Set("static", value)
	return a
}

func (a *Mapping) StaticClassName(value interface{}) *Mapping {
	a.Set("staticClassName", value)
	return a
}

func (a *Mapping) StaticInputClassName(value interface{}) *Mapping {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Mapping) StaticLabelClassName(value interface{}) *Mapping {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Mapping) StaticOn(value interface{}) *Mapping {
	a.Set("staticOn", value)
	return a
}

func (a *Mapping) StaticPlaceholder(value interface{}) *Mapping {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Mapping) StaticSchema(value interface{}) *Mapping {
	a.Set("staticSchema", value)
	return a
}

func (a *Mapping) Style(value interface{}) *Mapping {
	a.Set("style", value)
	return a
}

func (a *Mapping) TestIdBuilder(value interface{}) *Mapping {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Mapping) Testid(value interface{}) *Mapping {
	a.Set("testid", value)
	return a
}

func (a *Mapping) Type(value interface{}) *Mapping {
	a.Set("type", value)
	return a
}

func (a *Mapping) UseMobileUI(value interface{}) *Mapping {
	a.Set("useMobileUI", value)
	return a
}

func (a *Mapping) ValueField(value interface{}) *Mapping {
	a.Set("valueField", value)
	return a
}

func (a *Mapping) Visible(value interface{}) *Mapping {
	a.Set("visible", value)
	return a
}

func (a *Mapping) VisibleOn(value interface{}) *Mapping {
	a.Set("visibleOn", value)
	return a
}
