package renderers

type ListItem struct {
	*BaseRenderer
}

func NewListItem() *ListItem {
	a := &ListItem{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *ListItem) Set(name string, value interface{}) *ListItem {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *ListItem) Actions(value interface{}) *ListItem {
	a.Set("actions", value)
	return a
}

func (a *ListItem) ActionsPosition(value interface{}) *ListItem {
	a.Set("actionsPosition", value)
	return a
}

func (a *ListItem) Avatar(value interface{}) *ListItem {
	a.Set("avatar", value)
	return a
}

func (a *ListItem) Body(value interface{}) *ListItem {
	a.Set("body", value)
	return a
}

func (a *ListItem) ClassName(value interface{}) *ListItem {
	a.Set("className", value)
	return a
}

func (a *ListItem) Desc(value interface{}) *ListItem {
	a.Set("desc", value)
	return a
}

func (a *ListItem) Disabled(value interface{}) *ListItem {
	a.Set("disabled", value)
	return a
}

func (a *ListItem) DisabledOn(value interface{}) *ListItem {
	a.Set("disabledOn", value)
	return a
}

func (a *ListItem) EditorSetting(value interface{}) *ListItem {
	a.Set("editorSetting", value)
	return a
}

func (a *ListItem) Hidden(value interface{}) *ListItem {
	a.Set("hidden", value)
	return a
}

func (a *ListItem) HiddenOn(value interface{}) *ListItem {
	a.Set("hiddenOn", value)
	return a
}

func (a *ListItem) Id(value interface{}) *ListItem {
	a.Set("id", value)
	return a
}

func (a *ListItem) OnEvent(value interface{}) *ListItem {
	a.Set("onEvent", value)
	return a
}

func (a *ListItem) Remark(value interface{}) *ListItem {
	a.Set("remark", value)
	return a
}

func (a *ListItem) Static(value interface{}) *ListItem {
	a.Set("static", value)
	return a
}

func (a *ListItem) StaticClassName(value interface{}) *ListItem {
	a.Set("staticClassName", value)
	return a
}

func (a *ListItem) StaticInputClassName(value interface{}) *ListItem {
	a.Set("staticInputClassName", value)
	return a
}

func (a *ListItem) StaticLabelClassName(value interface{}) *ListItem {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *ListItem) StaticOn(value interface{}) *ListItem {
	a.Set("staticOn", value)
	return a
}

func (a *ListItem) StaticPlaceholder(value interface{}) *ListItem {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *ListItem) StaticSchema(value interface{}) *ListItem {
	a.Set("staticSchema", value)
	return a
}

func (a *ListItem) Style(value interface{}) *ListItem {
	a.Set("style", value)
	return a
}

func (a *ListItem) SubTitle(value interface{}) *ListItem {
	a.Set("subTitle", value)
	return a
}

func (a *ListItem) TestIdBuilder(value interface{}) *ListItem {
	a.Set("testIdBuilder", value)
	return a
}

func (a *ListItem) Testid(value interface{}) *ListItem {
	a.Set("testid", value)
	return a
}

func (a *ListItem) Title(value interface{}) *ListItem {
	a.Set("title", value)
	return a
}

func (a *ListItem) UseMobileUI(value interface{}) *ListItem {
	a.Set("useMobileUI", value)
	return a
}

func (a *ListItem) Visible(value interface{}) *ListItem {
	a.Set("visible", value)
	return a
}

func (a *ListItem) VisibleOn(value interface{}) *ListItem {
	a.Set("visibleOn", value)
	return a
}
