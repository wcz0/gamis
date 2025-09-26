package renderers

type Panel struct {
	*BaseRenderer
}

func NewPanel() *Panel {
	a := &Panel{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "panel")
	return a
}

func (a *Panel) Set(name string, value interface{}) *Panel {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Panel) Actions(value interface{}) *Panel {
	a.Set("actions", value)
	return a
}

func (a *Panel) ActionsClassName(value interface{}) *Panel {
	a.Set("actionsClassName", value)
	return a
}

func (a *Panel) ActionsControlClassName(value interface{}) *Panel {
	a.Set("actionsControlClassName", value)
	return a
}

func (a *Panel) AffixFooter(value interface{}) *Panel {
	a.Set("affixFooter", value)
	return a
}

func (a *Panel) Body(value interface{}) *Panel {
	a.Set("body", value)
	return a
}

func (a *Panel) BodyClassName(value interface{}) *Panel {
	a.Set("bodyClassName", value)
	return a
}

func (a *Panel) BodyControlClassName(value interface{}) *Panel {
	a.Set("bodyControlClassName", value)
	return a
}

func (a *Panel) ClassName(value interface{}) *Panel {
	a.Set("className", value)
	return a
}

func (a *Panel) Collapsible(value interface{}) *Panel {
	a.Set("collapsible", value)
	return a
}

func (a *Panel) Disabled(value interface{}) *Panel {
	a.Set("disabled", value)
	return a
}

func (a *Panel) DisabledOn(value interface{}) *Panel {
	a.Set("disabledOn", value)
	return a
}

func (a *Panel) EditorSetting(value interface{}) *Panel {
	a.Set("editorSetting", value)
	return a
}

func (a *Panel) Footer(value interface{}) *Panel {
	a.Set("footer", value)
	return a
}

func (a *Panel) FooterClassName(value interface{}) *Panel {
	a.Set("footerClassName", value)
	return a
}

func (a *Panel) FooterWrapClassName(value interface{}) *Panel {
	a.Set("footerWrapClassName", value)
	return a
}

func (a *Panel) Header(value interface{}) *Panel {
	a.Set("header", value)
	return a
}

func (a *Panel) HeaderClassName(value interface{}) *Panel {
	a.Set("headerClassName", value)
	return a
}

func (a *Panel) HeaderControlClassName(value interface{}) *Panel {
	a.Set("headerControlClassName", value)
	return a
}

func (a *Panel) Hidden(value interface{}) *Panel {
	a.Set("hidden", value)
	return a
}

func (a *Panel) HiddenOn(value interface{}) *Panel {
	a.Set("hiddenOn", value)
	return a
}

func (a *Panel) Id(value interface{}) *Panel {
	a.Set("id", value)
	return a
}

func (a *Panel) OnEvent(value interface{}) *Panel {
	a.Set("onEvent", value)
	return a
}

func (a *Panel) Static(value interface{}) *Panel {
	a.Set("static", value)
	return a
}

func (a *Panel) StaticClassName(value interface{}) *Panel {
	a.Set("staticClassName", value)
	return a
}

func (a *Panel) StaticInputClassName(value interface{}) *Panel {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Panel) StaticLabelClassName(value interface{}) *Panel {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Panel) StaticOn(value interface{}) *Panel {
	a.Set("staticOn", value)
	return a
}

func (a *Panel) StaticPlaceholder(value interface{}) *Panel {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Panel) StaticSchema(value interface{}) *Panel {
	a.Set("staticSchema", value)
	return a
}

func (a *Panel) Style(value interface{}) *Panel {
	a.Set("style", value)
	return a
}

func (a *Panel) SubFormHorizontal(value interface{}) *Panel {
	a.Set("subFormHorizontal", value)
	return a
}

func (a *Panel) SubFormMode(value interface{}) *Panel {
	a.Set("subFormMode", value)
	return a
}

func (a *Panel) TestIdBuilder(value interface{}) *Panel {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Panel) Testid(value interface{}) *Panel {
	a.Set("testid", value)
	return a
}

func (a *Panel) Title(value interface{}) *Panel {
	a.Set("title", value)
	return a
}

func (a *Panel) Type(value interface{}) *Panel {
	a.Set("type", value)
	return a
}

func (a *Panel) UseMobileUI(value interface{}) *Panel {
	a.Set("useMobileUI", value)
	return a
}

func (a *Panel) Visible(value interface{}) *Panel {
	a.Set("visible", value)
	return a
}

func (a *Panel) VisibleOn(value interface{}) *Panel {
	a.Set("visibleOn", value)
	return a
}
