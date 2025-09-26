package renderers

type DropdownButton struct {
	*BaseRenderer
}

func NewDropdownButton() *DropdownButton {
	a := &DropdownButton{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "dropdown-button")
	return a
}

func (a *DropdownButton) Set(name string, value interface{}) *DropdownButton {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *DropdownButton) Align(value interface{}) *DropdownButton {
	a.Set("align", value)
	return a
}

func (a *DropdownButton) Block(value interface{}) *DropdownButton {
	a.Set("block", value)
	return a
}

func (a *DropdownButton) Body(value interface{}) *DropdownButton {
	a.Set("body", value)
	return a
}

func (a *DropdownButton) BtnClassName(value interface{}) *DropdownButton {
	a.Set("btnClassName", value)
	return a
}

func (a *DropdownButton) Buttons(value interface{}) *DropdownButton {
	a.Set("buttons", value)
	return a
}

func (a *DropdownButton) ClassName(value interface{}) *DropdownButton {
	a.Set("className", value)
	return a
}

func (a *DropdownButton) CloseOnClick(value interface{}) *DropdownButton {
	a.Set("closeOnClick", value)
	return a
}

func (a *DropdownButton) CloseOnOutside(value interface{}) *DropdownButton {
	a.Set("closeOnOutside", value)
	return a
}

func (a *DropdownButton) Disabled(value interface{}) *DropdownButton {
	a.Set("disabled", value)
	return a
}

func (a *DropdownButton) DisabledOn(value interface{}) *DropdownButton {
	a.Set("disabledOn", value)
	return a
}

func (a *DropdownButton) EditorSetting(value interface{}) *DropdownButton {
	a.Set("editorSetting", value)
	return a
}

func (a *DropdownButton) Hidden(value interface{}) *DropdownButton {
	a.Set("hidden", value)
	return a
}

func (a *DropdownButton) HiddenOn(value interface{}) *DropdownButton {
	a.Set("hiddenOn", value)
	return a
}

func (a *DropdownButton) HideCaret(value interface{}) *DropdownButton {
	a.Set("hideCaret", value)
	return a
}

func (a *DropdownButton) IconOnly(value interface{}) *DropdownButton {
	a.Set("iconOnly", value)
	return a
}

func (a *DropdownButton) Id(value interface{}) *DropdownButton {
	a.Set("id", value)
	return a
}

func (a *DropdownButton) Label(value interface{}) *DropdownButton {
	a.Set("label", value)
	return a
}

func (a *DropdownButton) Level(value interface{}) *DropdownButton {
	a.Set("level", value)
	return a
}

func (a *DropdownButton) MenuClassName(value interface{}) *DropdownButton {
	a.Set("menuClassName", value)
	return a
}

func (a *DropdownButton) OnEvent(value interface{}) *DropdownButton {
	a.Set("onEvent", value)
	return a
}

func (a *DropdownButton) OverlayPlacement(value interface{}) *DropdownButton {
	a.Set("overlayPlacement", value)
	return a
}

func (a *DropdownButton) RightIcon(value interface{}) *DropdownButton {
	a.Set("rightIcon", value)
	return a
}

func (a *DropdownButton) Size(value interface{}) *DropdownButton {
	a.Set("size", value)
	return a
}

func (a *DropdownButton) Static(value interface{}) *DropdownButton {
	a.Set("static", value)
	return a
}

func (a *DropdownButton) StaticClassName(value interface{}) *DropdownButton {
	a.Set("staticClassName", value)
	return a
}

func (a *DropdownButton) StaticInputClassName(value interface{}) *DropdownButton {
	a.Set("staticInputClassName", value)
	return a
}

func (a *DropdownButton) StaticLabelClassName(value interface{}) *DropdownButton {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *DropdownButton) StaticOn(value interface{}) *DropdownButton {
	a.Set("staticOn", value)
	return a
}

func (a *DropdownButton) StaticPlaceholder(value interface{}) *DropdownButton {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *DropdownButton) StaticSchema(value interface{}) *DropdownButton {
	a.Set("staticSchema", value)
	return a
}

func (a *DropdownButton) Style(value interface{}) *DropdownButton {
	a.Set("style", value)
	return a
}

func (a *DropdownButton) TestIdBuilder(value interface{}) *DropdownButton {
	a.Set("testIdBuilder", value)
	return a
}

func (a *DropdownButton) Testid(value interface{}) *DropdownButton {
	a.Set("testid", value)
	return a
}

func (a *DropdownButton) Trigger(value interface{}) *DropdownButton {
	a.Set("trigger", value)
	return a
}

func (a *DropdownButton) Type(value interface{}) *DropdownButton {
	a.Set("type", value)
	return a
}

func (a *DropdownButton) UseMobileUI(value interface{}) *DropdownButton {
	a.Set("useMobileUI", value)
	return a
}

func (a *DropdownButton) Visible(value interface{}) *DropdownButton {
	a.Set("visible", value)
	return a
}

func (a *DropdownButton) VisibleOn(value interface{}) *DropdownButton {
	a.Set("visibleOn", value)
	return a
}
