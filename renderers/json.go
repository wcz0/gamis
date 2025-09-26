package renderers

type Json struct {
	*BaseRenderer
}

func NewJson() *Json {
	a := &Json{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "json")
	return a
}

func (a *Json) Set(name string, value interface{}) *Json {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Json) ClassName(value interface{}) *Json {
	a.Set("className", value)
	return a
}

func (a *Json) Disabled(value interface{}) *Json {
	a.Set("disabled", value)
	return a
}

func (a *Json) DisabledOn(value interface{}) *Json {
	a.Set("disabledOn", value)
	return a
}

func (a *Json) DisplayDataTypes(value interface{}) *Json {
	a.Set("displayDataTypes", value)
	return a
}

func (a *Json) EditorSetting(value interface{}) *Json {
	a.Set("editorSetting", value)
	return a
}

func (a *Json) EllipsisThreshold(value interface{}) *Json {
	a.Set("ellipsisThreshold", value)
	return a
}

func (a *Json) EnableClipboard(value interface{}) *Json {
	a.Set("enableClipboard", value)
	return a
}

func (a *Json) Hidden(value interface{}) *Json {
	a.Set("hidden", value)
	return a
}

func (a *Json) HiddenOn(value interface{}) *Json {
	a.Set("hiddenOn", value)
	return a
}

func (a *Json) IconStyle(value interface{}) *Json {
	a.Set("iconStyle", value)
	return a
}

func (a *Json) Id(value interface{}) *Json {
	a.Set("id", value)
	return a
}

func (a *Json) LevelExpand(value interface{}) *Json {
	a.Set("levelExpand", value)
	return a
}

func (a *Json) Mutable(value interface{}) *Json {
	a.Set("mutable", value)
	return a
}

func (a *Json) OnEvent(value interface{}) *Json {
	a.Set("onEvent", value)
	return a
}

func (a *Json) QuotesOnKeys(value interface{}) *Json {
	a.Set("quotesOnKeys", value)
	return a
}

func (a *Json) SortKeys(value interface{}) *Json {
	a.Set("sortKeys", value)
	return a
}

func (a *Json) Source(value interface{}) *Json {
	a.Set("source", value)
	return a
}

func (a *Json) Static(value interface{}) *Json {
	a.Set("static", value)
	return a
}

func (a *Json) StaticClassName(value interface{}) *Json {
	a.Set("staticClassName", value)
	return a
}

func (a *Json) StaticInputClassName(value interface{}) *Json {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Json) StaticLabelClassName(value interface{}) *Json {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Json) StaticOn(value interface{}) *Json {
	a.Set("staticOn", value)
	return a
}

func (a *Json) StaticPlaceholder(value interface{}) *Json {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Json) StaticSchema(value interface{}) *Json {
	a.Set("staticSchema", value)
	return a
}

func (a *Json) Style(value interface{}) *Json {
	a.Set("style", value)
	return a
}

func (a *Json) TestIdBuilder(value interface{}) *Json {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Json) Testid(value interface{}) *Json {
	a.Set("testid", value)
	return a
}

func (a *Json) Type(value interface{}) *Json {
	a.Set("type", value)
	return a
}

func (a *Json) UseMobileUI(value interface{}) *Json {
	a.Set("useMobileUI", value)
	return a
}

func (a *Json) Value(value interface{}) *Json {
	a.Set("value", value)
	return a
}

func (a *Json) Visible(value interface{}) *Json {
	a.Set("visible", value)
	return a
}

func (a *Json) VisibleOn(value interface{}) *Json {
	a.Set("visibleOn", value)
	return a
}
