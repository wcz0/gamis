package renderers

type Alert struct {
	*BaseRenderer
}

func NewAlert() *Alert {
	a := &Alert{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "alert")
	return a
}

func (a *Alert) Set(name string, value interface{}) *Alert {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Alert) Actions(value interface{}) *Alert {
	a.Set("actions", value)
	return a
}

func (a *Alert) Body(value interface{}) *Alert {
	a.Set("body", value)
	return a
}

func (a *Alert) ClassName(value interface{}) *Alert {
	a.Set("className", value)
	return a
}

func (a *Alert) CloseButtonClassName(value interface{}) *Alert {
	a.Set("closeButtonClassName", value)
	return a
}

func (a *Alert) Disabled(value interface{}) *Alert {
	a.Set("disabled", value)
	return a
}

func (a *Alert) DisabledOn(value interface{}) *Alert {
	a.Set("disabledOn", value)
	return a
}

func (a *Alert) EditorSetting(value interface{}) *Alert {
	a.Set("editorSetting", value)
	return a
}

func (a *Alert) Hidden(value interface{}) *Alert {
	a.Set("hidden", value)
	return a
}

func (a *Alert) HiddenOn(value interface{}) *Alert {
	a.Set("hiddenOn", value)
	return a
}

func (a *Alert) Icon(value interface{}) *Alert {
	a.Set("icon", value)
	return a
}

func (a *Alert) IconClassName(value interface{}) *Alert {
	a.Set("iconClassName", value)
	return a
}

func (a *Alert) Id(value interface{}) *Alert {
	a.Set("id", value)
	return a
}

func (a *Alert) Level(value interface{}) *Alert {
	a.Set("level", value)
	return a
}

func (a *Alert) OnEvent(value interface{}) *Alert {
	a.Set("onEvent", value)
	return a
}

func (a *Alert) ShowCloseButton(value interface{}) *Alert {
	a.Set("showCloseButton", value)
	return a
}

func (a *Alert) ShowIcon(value interface{}) *Alert {
	a.Set("showIcon", value)
	return a
}

func (a *Alert) Static(value interface{}) *Alert {
	a.Set("static", value)
	return a
}

func (a *Alert) StaticClassName(value interface{}) *Alert {
	a.Set("staticClassName", value)
	return a
}

func (a *Alert) StaticInputClassName(value interface{}) *Alert {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Alert) StaticLabelClassName(value interface{}) *Alert {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Alert) StaticOn(value interface{}) *Alert {
	a.Set("staticOn", value)
	return a
}

func (a *Alert) StaticPlaceholder(value interface{}) *Alert {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Alert) StaticSchema(value interface{}) *Alert {
	a.Set("staticSchema", value)
	return a
}

func (a *Alert) Style(value interface{}) *Alert {
	a.Set("style", value)
	return a
}

func (a *Alert) TestIdBuilder(value interface{}) *Alert {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Alert) Testid(value interface{}) *Alert {
	a.Set("testid", value)
	return a
}

func (a *Alert) Title(value interface{}) *Alert {
	a.Set("title", value)
	return a
}

func (a *Alert) Type(value interface{}) *Alert {
	a.Set("type", value)
	return a
}

func (a *Alert) UseMobileUI(value interface{}) *Alert {
	a.Set("useMobileUI", value)
	return a
}

func (a *Alert) Visible(value interface{}) *Alert {
	a.Set("visible", value)
	return a
}

func (a *Alert) VisibleOn(value interface{}) *Alert {
	a.Set("visibleOn", value)
	return a
}
