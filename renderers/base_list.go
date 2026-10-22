package renderers

type BaseList struct {
	*BaseRenderer
}

func NewBaseList() *BaseList {
	a := &BaseList{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *BaseList) Set(name string, value interface{}) *BaseList {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *BaseList) AffixFooter(value interface{}) *BaseList {
	a.Set("affixFooter", value)
	return a
}

func (a *BaseList) AffixHeader(value interface{}) *BaseList {
	a.Set("affixHeader", value)
	return a
}

func (a *BaseList) CheckOnItemClick(value interface{}) *BaseList {
	a.Set("checkOnItemClick", value)
	return a
}

func (a *BaseList) ClassName(value interface{}) *BaseList {
	a.Set("className", value)
	return a
}

func (a *BaseList) Disabled(value interface{}) *BaseList {
	a.Set("disabled", value)
	return a
}

func (a *BaseList) DisabledOn(value interface{}) *BaseList {
	a.Set("disabledOn", value)
	return a
}

func (a *BaseList) EditorSetting(value interface{}) *BaseList {
	a.Set("editorSetting", value)
	return a
}

func (a *BaseList) Footer(value interface{}) *BaseList {
	a.Set("footer", value)
	return a
}

func (a *BaseList) FooterClassName(value interface{}) *BaseList {
	a.Set("footerClassName", value)
	return a
}

func (a *BaseList) Header(value interface{}) *BaseList {
	a.Set("header", value)
	return a
}

func (a *BaseList) HeaderClassName(value interface{}) *BaseList {
	a.Set("headerClassName", value)
	return a
}

func (a *BaseList) Hidden(value interface{}) *BaseList {
	a.Set("hidden", value)
	return a
}

func (a *BaseList) HiddenOn(value interface{}) *BaseList {
	a.Set("hiddenOn", value)
	return a
}

func (a *BaseList) HideCheckToggler(value interface{}) *BaseList {
	a.Set("hideCheckToggler", value)
	return a
}

func (a *BaseList) Id(value interface{}) *BaseList {
	a.Set("id", value)
	return a
}

func (a *BaseList) IndexBarOffset(value interface{}) *BaseList {
	a.Set("indexBarOffset", value)
	return a
}

func (a *BaseList) IndexField(value interface{}) *BaseList {
	a.Set("indexField", value)
	return a
}

func (a *BaseList) ItemAction(value interface{}) *BaseList {
	a.Set("itemAction", value)
	return a
}

func (a *BaseList) ItemCheckableOn(value interface{}) *BaseList {
	a.Set("itemCheckableOn", value)
	return a
}

func (a *BaseList) ItemDraggableOn(value interface{}) *BaseList {
	a.Set("itemDraggableOn", value)
	return a
}

func (a *BaseList) ListItem(value interface{}) *BaseList {
	a.Set("listItem", value)
	return a
}

func (a *BaseList) OnEvent(value interface{}) *BaseList {
	a.Set("onEvent", value)
	return a
}

func (a *BaseList) Placeholder(value interface{}) *BaseList {
	a.Set("placeholder", value)
	return a
}

func (a *BaseList) ShowFooter(value interface{}) *BaseList {
	a.Set("showFooter", value)
	return a
}

func (a *BaseList) ShowHeader(value interface{}) *BaseList {
	a.Set("showHeader", value)
	return a
}

func (a *BaseList) ShowIndexBar(value interface{}) *BaseList {
	a.Set("showIndexBar", value)
	return a
}

func (a *BaseList) Size(value interface{}) *BaseList {
	a.Set("size", value)
	return a
}

func (a *BaseList) Source(value interface{}) *BaseList {
	a.Set("source", value)
	return a
}

func (a *BaseList) Static(value interface{}) *BaseList {
	a.Set("static", value)
	return a
}

func (a *BaseList) StaticClassName(value interface{}) *BaseList {
	a.Set("staticClassName", value)
	return a
}

func (a *BaseList) StaticInputClassName(value interface{}) *BaseList {
	a.Set("staticInputClassName", value)
	return a
}

func (a *BaseList) StaticLabelClassName(value interface{}) *BaseList {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *BaseList) StaticOn(value interface{}) *BaseList {
	a.Set("staticOn", value)
	return a
}

func (a *BaseList) StaticPlaceholder(value interface{}) *BaseList {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *BaseList) StaticSchema(value interface{}) *BaseList {
	a.Set("staticSchema", value)
	return a
}

func (a *BaseList) Style(value interface{}) *BaseList {
	a.Set("style", value)
	return a
}

func (a *BaseList) Testid(value interface{}) *BaseList {
	a.Set("testid", value)
	return a
}

func (a *BaseList) Title(value interface{}) *BaseList {
	a.Set("title", value)
	return a
}

func (a *BaseList) Type(value interface{}) *BaseList {
	a.Set("type", value)
	return a
}

func (a *BaseList) UseMobileUI(value interface{}) *BaseList {
	a.Set("useMobileUI", value)
	return a
}

func (a *BaseList) ValueField(value interface{}) *BaseList {
	a.Set("valueField", value)
	return a
}

func (a *BaseList) Visible(value interface{}) *BaseList {
	a.Set("visible", value)
	return a
}

func (a *BaseList) VisibleOn(value interface{}) *BaseList {
	a.Set("visibleOn", value)
	return a
}
