package renderers

type Tag struct {
	*BaseRenderer
}

func NewTag() *Tag {
	a := &Tag{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "tag")
	return a
}

func (a *Tag) Set(name string, value interface{}) *Tag {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Tag) Checkable(value interface{}) *Tag {
	a.Set("checkable", value)
	return a
}

func (a *Tag) Checked(value interface{}) *Tag {
	a.Set("checked", value)
	return a
}

func (a *Tag) ClassName(value interface{}) *Tag {
	a.Set("className", value)
	return a
}

func (a *Tag) Closable(value interface{}) *Tag {
	a.Set("closable", value)
	return a
}

func (a *Tag) CloseIcon(value interface{}) *Tag {
	a.Set("closeIcon", value)
	return a
}

func (a *Tag) Color(value interface{}) *Tag {
	a.Set("color", value)
	return a
}

func (a *Tag) Disabled(value interface{}) *Tag {
	a.Set("disabled", value)
	return a
}

func (a *Tag) DisabledOn(value interface{}) *Tag {
	a.Set("disabledOn", value)
	return a
}

func (a *Tag) DisplayMode(value interface{}) *Tag {
	a.Set("displayMode", value)
	return a
}

func (a *Tag) EditorSetting(value interface{}) *Tag {
	a.Set("editorSetting", value)
	return a
}

func (a *Tag) Hidden(value interface{}) *Tag {
	a.Set("hidden", value)
	return a
}

func (a *Tag) HiddenOn(value interface{}) *Tag {
	a.Set("hiddenOn", value)
	return a
}

func (a *Tag) Icon(value interface{}) *Tag {
	a.Set("icon", value)
	return a
}

func (a *Tag) Id(value interface{}) *Tag {
	a.Set("id", value)
	return a
}

func (a *Tag) Label(value interface{}) *Tag {
	a.Set("label", value)
	return a
}

func (a *Tag) OnEvent(value interface{}) *Tag {
	a.Set("onEvent", value)
	return a
}

func (a *Tag) Static(value interface{}) *Tag {
	a.Set("static", value)
	return a
}

func (a *Tag) StaticClassName(value interface{}) *Tag {
	a.Set("staticClassName", value)
	return a
}

func (a *Tag) StaticInputClassName(value interface{}) *Tag {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Tag) StaticLabelClassName(value interface{}) *Tag {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Tag) StaticOn(value interface{}) *Tag {
	a.Set("staticOn", value)
	return a
}

func (a *Tag) StaticPlaceholder(value interface{}) *Tag {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Tag) StaticSchema(value interface{}) *Tag {
	a.Set("staticSchema", value)
	return a
}

func (a *Tag) Style(value interface{}) *Tag {
	a.Set("style", value)
	return a
}

func (a *Tag) TestIdBuilder(value interface{}) *Tag {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Tag) Testid(value interface{}) *Tag {
	a.Set("testid", value)
	return a
}

func (a *Tag) Type(value interface{}) *Tag {
	a.Set("type", value)
	return a
}

func (a *Tag) UseMobileUI(value interface{}) *Tag {
	a.Set("useMobileUI", value)
	return a
}

func (a *Tag) Visible(value interface{}) *Tag {
	a.Set("visible", value)
	return a
}

func (a *Tag) VisibleOn(value interface{}) *Tag {
	a.Set("visibleOn", value)
	return a
}
