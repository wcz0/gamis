package renderers

type IFrame struct {
	*BaseRenderer
}

func NewIFrame() *IFrame {
	a := &IFrame{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "iframe")
	return a
}

func (a *IFrame) Set(name string, value interface{}) *IFrame {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *IFrame) Allow(value interface{}) *IFrame {
	a.Set("allow", value)
	return a
}

func (a *IFrame) ClassName(value interface{}) *IFrame {
	a.Set("className", value)
	return a
}

func (a *IFrame) Disabled(value interface{}) *IFrame {
	a.Set("disabled", value)
	return a
}

func (a *IFrame) DisabledOn(value interface{}) *IFrame {
	a.Set("disabledOn", value)
	return a
}

func (a *IFrame) EditorSetting(value interface{}) *IFrame {
	a.Set("editorSetting", value)
	return a
}

func (a *IFrame) Events(value interface{}) *IFrame {
	a.Set("events", value)
	return a
}

func (a *IFrame) Height(value interface{}) *IFrame {
	a.Set("height", value)
	return a
}

func (a *IFrame) Hidden(value interface{}) *IFrame {
	a.Set("hidden", value)
	return a
}

func (a *IFrame) HiddenOn(value interface{}) *IFrame {
	a.Set("hiddenOn", value)
	return a
}

func (a *IFrame) Id(value interface{}) *IFrame {
	a.Set("id", value)
	return a
}

func (a *IFrame) Name(value interface{}) *IFrame {
	a.Set("name", value)
	return a
}

func (a *IFrame) OnEvent(value interface{}) *IFrame {
	a.Set("onEvent", value)
	return a
}

func (a *IFrame) Referrerpolicy(value interface{}) *IFrame {
	a.Set("referrerpolicy", value)
	return a
}

func (a *IFrame) Sandbox(value interface{}) *IFrame {
	a.Set("sandbox", value)
	return a
}

func (a *IFrame) Src(value interface{}) *IFrame {
	a.Set("src", value)
	return a
}

func (a *IFrame) Static(value interface{}) *IFrame {
	a.Set("static", value)
	return a
}

func (a *IFrame) StaticClassName(value interface{}) *IFrame {
	a.Set("staticClassName", value)
	return a
}

func (a *IFrame) StaticInputClassName(value interface{}) *IFrame {
	a.Set("staticInputClassName", value)
	return a
}

func (a *IFrame) StaticLabelClassName(value interface{}) *IFrame {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *IFrame) StaticOn(value interface{}) *IFrame {
	a.Set("staticOn", value)
	return a
}

func (a *IFrame) StaticPlaceholder(value interface{}) *IFrame {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *IFrame) StaticSchema(value interface{}) *IFrame {
	a.Set("staticSchema", value)
	return a
}

func (a *IFrame) Style(value interface{}) *IFrame {
	a.Set("style", value)
	return a
}

func (a *IFrame) TestIdBuilder(value interface{}) *IFrame {
	a.Set("testIdBuilder", value)
	return a
}

func (a *IFrame) Testid(value interface{}) *IFrame {
	a.Set("testid", value)
	return a
}

func (a *IFrame) Type(value interface{}) *IFrame {
	a.Set("type", value)
	return a
}

func (a *IFrame) UseMobileUI(value interface{}) *IFrame {
	a.Set("useMobileUI", value)
	return a
}

func (a *IFrame) Visible(value interface{}) *IFrame {
	a.Set("visible", value)
	return a
}

func (a *IFrame) VisibleOn(value interface{}) *IFrame {
	a.Set("visibleOn", value)
	return a
}

func (a *IFrame) Width(value interface{}) *IFrame {
	a.Set("width", value)
	return a
}
