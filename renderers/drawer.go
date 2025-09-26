package renderers

type Drawer struct {
	*BaseRenderer
}

func NewDrawer() *Drawer {
	a := &Drawer{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "drawer")
	return a
}

func (a *Drawer) Set(name string, value interface{}) *Drawer {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Drawer) Actions(value interface{}) *Drawer {
	a.Set("actions", value)
	return a
}

func (a *Drawer) Body(value interface{}) *Drawer {
	a.Set("body", value)
	return a
}

func (a *Drawer) BodyClassName(value interface{}) *Drawer {
	a.Set("bodyClassName", value)
	return a
}

func (a *Drawer) ClassName(value interface{}) *Drawer {
	a.Set("className", value)
	return a
}

func (a *Drawer) CloseOnEsc(value interface{}) *Drawer {
	a.Set("closeOnEsc", value)
	return a
}

func (a *Drawer) CloseOnOutside(value interface{}) *Drawer {
	a.Set("closeOnOutside", value)
	return a
}

func (a *Drawer) Confirm(value interface{}) *Drawer {
	a.Set("confirm", value)
	return a
}

func (a *Drawer) Data(value interface{}) *Drawer {
	a.Set("data", value)
	return a
}

func (a *Drawer) Disabled(value interface{}) *Drawer {
	a.Set("disabled", value)
	return a
}

func (a *Drawer) DisabledOn(value interface{}) *Drawer {
	a.Set("disabledOn", value)
	return a
}

func (a *Drawer) EditorSetting(value interface{}) *Drawer {
	a.Set("editorSetting", value)
	return a
}

func (a *Drawer) Footer(value interface{}) *Drawer {
	a.Set("footer", value)
	return a
}

func (a *Drawer) FooterClassName(value interface{}) *Drawer {
	a.Set("footerClassName", value)
	return a
}

func (a *Drawer) Header(value interface{}) *Drawer {
	a.Set("header", value)
	return a
}

func (a *Drawer) HeaderClassName(value interface{}) *Drawer {
	a.Set("headerClassName", value)
	return a
}

func (a *Drawer) Height(value interface{}) *Drawer {
	a.Set("height", value)
	return a
}

func (a *Drawer) Hidden(value interface{}) *Drawer {
	a.Set("hidden", value)
	return a
}

func (a *Drawer) HiddenOn(value interface{}) *Drawer {
	a.Set("hiddenOn", value)
	return a
}

func (a *Drawer) Id(value interface{}) *Drawer {
	a.Set("id", value)
	return a
}

func (a *Drawer) InputParams(value interface{}) *Drawer {
	a.Set("inputParams", value)
	return a
}

func (a *Drawer) Name(value interface{}) *Drawer {
	a.Set("name", value)
	return a
}

func (a *Drawer) OnEvent(value interface{}) *Drawer {
	a.Set("onEvent", value)
	return a
}

func (a *Drawer) Overlay(value interface{}) *Drawer {
	a.Set("overlay", value)
	return a
}

func (a *Drawer) Position(value interface{}) *Drawer {
	a.Set("position", value)
	return a
}

func (a *Drawer) Resizable(value interface{}) *Drawer {
	a.Set("resizable", value)
	return a
}

func (a *Drawer) ShowCloseButton(value interface{}) *Drawer {
	a.Set("showCloseButton", value)
	return a
}

func (a *Drawer) ShowErrorMsg(value interface{}) *Drawer {
	a.Set("showErrorMsg", value)
	return a
}

func (a *Drawer) Size(value interface{}) *Drawer {
	a.Set("size", value)
	return a
}

func (a *Drawer) Static(value interface{}) *Drawer {
	a.Set("static", value)
	return a
}

func (a *Drawer) StaticClassName(value interface{}) *Drawer {
	a.Set("staticClassName", value)
	return a
}

func (a *Drawer) StaticInputClassName(value interface{}) *Drawer {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Drawer) StaticLabelClassName(value interface{}) *Drawer {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Drawer) StaticOn(value interface{}) *Drawer {
	a.Set("staticOn", value)
	return a
}

func (a *Drawer) StaticPlaceholder(value interface{}) *Drawer {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Drawer) StaticSchema(value interface{}) *Drawer {
	a.Set("staticSchema", value)
	return a
}

func (a *Drawer) Style(value interface{}) *Drawer {
	a.Set("style", value)
	return a
}

func (a *Drawer) TestIdBuilder(value interface{}) *Drawer {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Drawer) Testid(value interface{}) *Drawer {
	a.Set("testid", value)
	return a
}

func (a *Drawer) Title(value interface{}) *Drawer {
	a.Set("title", value)
	return a
}

func (a *Drawer) Type(value interface{}) *Drawer {
	a.Set("type", value)
	return a
}

func (a *Drawer) UseMobileUI(value interface{}) *Drawer {
	a.Set("useMobileUI", value)
	return a
}

func (a *Drawer) Visible(value interface{}) *Drawer {
	a.Set("visible", value)
	return a
}

func (a *Drawer) VisibleOn(value interface{}) *Drawer {
	a.Set("visibleOn", value)
	return a
}

func (a *Drawer) Width(value interface{}) *Drawer {
	a.Set("width", value)
	return a
}
