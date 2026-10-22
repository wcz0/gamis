package renderers

type BaseCards struct {
	*BaseRenderer
}

func NewBaseCards() *BaseCards {
	a := &BaseCards{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *BaseCards) Set(name string, value interface{}) *BaseCards {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *BaseCards) AffixFooter(value interface{}) *BaseCards {
	a.Set("affixFooter", value)
	return a
}

func (a *BaseCards) AffixHeader(value interface{}) *BaseCards {
	a.Set("affixHeader", value)
	return a
}

func (a *BaseCards) Card(value interface{}) *BaseCards {
	a.Set("card", value)
	return a
}

func (a *BaseCards) CheckOnItemClick(value interface{}) *BaseCards {
	a.Set("checkOnItemClick", value)
	return a
}

func (a *BaseCards) ClassName(value interface{}) *BaseCards {
	a.Set("className", value)
	return a
}

func (a *BaseCards) Disabled(value interface{}) *BaseCards {
	a.Set("disabled", value)
	return a
}

func (a *BaseCards) DisabledOn(value interface{}) *BaseCards {
	a.Set("disabledOn", value)
	return a
}

func (a *BaseCards) EditorSetting(value interface{}) *BaseCards {
	a.Set("editorSetting", value)
	return a
}

func (a *BaseCards) Footer(value interface{}) *BaseCards {
	a.Set("footer", value)
	return a
}

func (a *BaseCards) FooterClassName(value interface{}) *BaseCards {
	a.Set("footerClassName", value)
	return a
}

func (a *BaseCards) Header(value interface{}) *BaseCards {
	a.Set("header", value)
	return a
}

func (a *BaseCards) HeaderClassName(value interface{}) *BaseCards {
	a.Set("headerClassName", value)
	return a
}

func (a *BaseCards) Hidden(value interface{}) *BaseCards {
	a.Set("hidden", value)
	return a
}

func (a *BaseCards) HiddenOn(value interface{}) *BaseCards {
	a.Set("hiddenOn", value)
	return a
}

func (a *BaseCards) HideCheckToggler(value interface{}) *BaseCards {
	a.Set("hideCheckToggler", value)
	return a
}

func (a *BaseCards) Id(value interface{}) *BaseCards {
	a.Set("id", value)
	return a
}

func (a *BaseCards) ItemCheckableOn(value interface{}) *BaseCards {
	a.Set("itemCheckableOn", value)
	return a
}

func (a *BaseCards) ItemClassName(value interface{}) *BaseCards {
	a.Set("itemClassName", value)
	return a
}

func (a *BaseCards) ItemDraggableOn(value interface{}) *BaseCards {
	a.Set("itemDraggableOn", value)
	return a
}

func (a *BaseCards) LoadingConfig(value interface{}) *BaseCards {
	a.Set("loadingConfig", value)
	return a
}

func (a *BaseCards) MasonryLayout(value interface{}) *BaseCards {
	a.Set("masonryLayout", value)
	return a
}

func (a *BaseCards) OnEvent(value interface{}) *BaseCards {
	a.Set("onEvent", value)
	return a
}

func (a *BaseCards) Placeholder(value interface{}) *BaseCards {
	a.Set("placeholder", value)
	return a
}

func (a *BaseCards) ShowFooter(value interface{}) *BaseCards {
	a.Set("showFooter", value)
	return a
}

func (a *BaseCards) ShowHeader(value interface{}) *BaseCards {
	a.Set("showHeader", value)
	return a
}

func (a *BaseCards) Source(value interface{}) *BaseCards {
	a.Set("source", value)
	return a
}

func (a *BaseCards) Static(value interface{}) *BaseCards {
	a.Set("static", value)
	return a
}

func (a *BaseCards) StaticClassName(value interface{}) *BaseCards {
	a.Set("staticClassName", value)
	return a
}

func (a *BaseCards) StaticInputClassName(value interface{}) *BaseCards {
	a.Set("staticInputClassName", value)
	return a
}

func (a *BaseCards) StaticLabelClassName(value interface{}) *BaseCards {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *BaseCards) StaticOn(value interface{}) *BaseCards {
	a.Set("staticOn", value)
	return a
}

func (a *BaseCards) StaticPlaceholder(value interface{}) *BaseCards {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *BaseCards) StaticSchema(value interface{}) *BaseCards {
	a.Set("staticSchema", value)
	return a
}

func (a *BaseCards) Style(value interface{}) *BaseCards {
	a.Set("style", value)
	return a
}

func (a *BaseCards) Testid(value interface{}) *BaseCards {
	a.Set("testid", value)
	return a
}

func (a *BaseCards) Title(value interface{}) *BaseCards {
	a.Set("title", value)
	return a
}

func (a *BaseCards) Type(value interface{}) *BaseCards {
	a.Set("type", value)
	return a
}

func (a *BaseCards) UseMobileUI(value interface{}) *BaseCards {
	a.Set("useMobileUI", value)
	return a
}

func (a *BaseCards) ValueField(value interface{}) *BaseCards {
	a.Set("valueField", value)
	return a
}

func (a *BaseCards) Visible(value interface{}) *BaseCards {
	a.Set("visible", value)
	return a
}

func (a *BaseCards) VisibleOn(value interface{}) *BaseCards {
	a.Set("visibleOn", value)
	return a
}
