package renderers

type ListRenderer struct {
	*BaseRenderer
}

func NewListRenderer() *ListRenderer {
	a := &ListRenderer{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "list")
	return a
}

func (a *ListRenderer) Set(name string, value interface{}) *ListRenderer {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *ListRenderer) AffixFooter(value interface{}) *ListRenderer {
	a.Set("affixFooter", value)
	return a
}

func (a *ListRenderer) AffixHeader(value interface{}) *ListRenderer {
	a.Set("affixHeader", value)
	return a
}

func (a *ListRenderer) CheckOnItemClick(value interface{}) *ListRenderer {
	a.Set("checkOnItemClick", value)
	return a
}

func (a *ListRenderer) ClassName(value interface{}) *ListRenderer {
	a.Set("className", value)
	return a
}

func (a *ListRenderer) Disabled(value interface{}) *ListRenderer {
	a.Set("disabled", value)
	return a
}

func (a *ListRenderer) DisabledOn(value interface{}) *ListRenderer {
	a.Set("disabledOn", value)
	return a
}

func (a *ListRenderer) EditorSetting(value interface{}) *ListRenderer {
	a.Set("editorSetting", value)
	return a
}

func (a *ListRenderer) Footer(value interface{}) *ListRenderer {
	a.Set("footer", value)
	return a
}

func (a *ListRenderer) FooterClassName(value interface{}) *ListRenderer {
	a.Set("footerClassName", value)
	return a
}

func (a *ListRenderer) Header(value interface{}) *ListRenderer {
	a.Set("header", value)
	return a
}

func (a *ListRenderer) HeaderClassName(value interface{}) *ListRenderer {
	a.Set("headerClassName", value)
	return a
}

func (a *ListRenderer) Hidden(value interface{}) *ListRenderer {
	a.Set("hidden", value)
	return a
}

func (a *ListRenderer) HiddenOn(value interface{}) *ListRenderer {
	a.Set("hiddenOn", value)
	return a
}

func (a *ListRenderer) HideCheckToggler(value interface{}) *ListRenderer {
	a.Set("hideCheckToggler", value)
	return a
}

func (a *ListRenderer) Id(value interface{}) *ListRenderer {
	a.Set("id", value)
	return a
}

func (a *ListRenderer) IndexBarOffset(value interface{}) *ListRenderer {
	a.Set("indexBarOffset", value)
	return a
}

func (a *ListRenderer) IndexField(value interface{}) *ListRenderer {
	a.Set("indexField", value)
	return a
}

func (a *ListRenderer) ItemAction(value interface{}) *ListRenderer {
	a.Set("itemAction", value)
	return a
}

func (a *ListRenderer) ItemCheckableOn(value interface{}) *ListRenderer {
	a.Set("itemCheckableOn", value)
	return a
}

func (a *ListRenderer) ItemDraggableOn(value interface{}) *ListRenderer {
	a.Set("itemDraggableOn", value)
	return a
}

func (a *ListRenderer) ListItem(value interface{}) *ListRenderer {
	a.Set("listItem", value)
	return a
}

func (a *ListRenderer) OnEvent(value interface{}) *ListRenderer {
	a.Set("onEvent", value)
	return a
}

func (a *ListRenderer) Placeholder(value interface{}) *ListRenderer {
	a.Set("placeholder", value)
	return a
}

func (a *ListRenderer) ShowFooter(value interface{}) *ListRenderer {
	a.Set("showFooter", value)
	return a
}

func (a *ListRenderer) ShowHeader(value interface{}) *ListRenderer {
	a.Set("showHeader", value)
	return a
}

func (a *ListRenderer) ShowIndexBar(value interface{}) *ListRenderer {
	a.Set("showIndexBar", value)
	return a
}

func (a *ListRenderer) Size(value interface{}) *ListRenderer {
	a.Set("size", value)
	return a
}

func (a *ListRenderer) Source(value interface{}) *ListRenderer {
	a.Set("source", value)
	return a
}

func (a *ListRenderer) Static(value interface{}) *ListRenderer {
	a.Set("static", value)
	return a
}

func (a *ListRenderer) StaticClassName(value interface{}) *ListRenderer {
	a.Set("staticClassName", value)
	return a
}

func (a *ListRenderer) StaticInputClassName(value interface{}) *ListRenderer {
	a.Set("staticInputClassName", value)
	return a
}

func (a *ListRenderer) StaticLabelClassName(value interface{}) *ListRenderer {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *ListRenderer) StaticOn(value interface{}) *ListRenderer {
	a.Set("staticOn", value)
	return a
}

func (a *ListRenderer) StaticPlaceholder(value interface{}) *ListRenderer {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *ListRenderer) StaticSchema(value interface{}) *ListRenderer {
	a.Set("staticSchema", value)
	return a
}

func (a *ListRenderer) Style(value interface{}) *ListRenderer {
	a.Set("style", value)
	return a
}

func (a *ListRenderer) TestIdBuilder(value interface{}) *ListRenderer {
	a.Set("testIdBuilder", value)
	return a
}

func (a *ListRenderer) Testid(value interface{}) *ListRenderer {
	a.Set("testid", value)
	return a
}

func (a *ListRenderer) Title(value interface{}) *ListRenderer {
	a.Set("title", value)
	return a
}

func (a *ListRenderer) Type(value interface{}) *ListRenderer {
	a.Set("type", value)
	return a
}

func (a *ListRenderer) UseMobileUI(value interface{}) *ListRenderer {
	a.Set("useMobileUI", value)
	return a
}

func (a *ListRenderer) ValueField(value interface{}) *ListRenderer {
	a.Set("valueField", value)
	return a
}

func (a *ListRenderer) Visible(value interface{}) *ListRenderer {
	a.Set("visible", value)
	return a
}

func (a *ListRenderer) VisibleOn(value interface{}) *ListRenderer {
	a.Set("visibleOn", value)
	return a
}
