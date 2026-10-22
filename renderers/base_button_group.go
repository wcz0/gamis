package renderers

type BaseButtonGroup struct {
	*BaseRenderer
}

func NewBaseButtonGroup() *BaseButtonGroup {
	a := &BaseButtonGroup{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *BaseButtonGroup) Set(name string, value interface{}) *BaseButtonGroup {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *BaseButtonGroup) BtnActiveClassName(value interface{}) *BaseButtonGroup {
	a.Set("btnActiveClassName", value)
	return a
}

func (a *BaseButtonGroup) BtnActiveLevel(value interface{}) *BaseButtonGroup {
	a.Set("btnActiveLevel", value)
	return a
}

func (a *BaseButtonGroup) BtnClassName(value interface{}) *BaseButtonGroup {
	a.Set("btnClassName", value)
	return a
}

func (a *BaseButtonGroup) BtnLevel(value interface{}) *BaseButtonGroup {
	a.Set("btnLevel", value)
	return a
}

func (a *BaseButtonGroup) Buttons(value interface{}) *BaseButtonGroup {
	a.Set("buttons", value)
	return a
}

func (a *BaseButtonGroup) ClassName(value interface{}) *BaseButtonGroup {
	a.Set("className", value)
	return a
}

func (a *BaseButtonGroup) Disabled(value interface{}) *BaseButtonGroup {
	a.Set("disabled", value)
	return a
}

func (a *BaseButtonGroup) DisabledOn(value interface{}) *BaseButtonGroup {
	a.Set("disabledOn", value)
	return a
}

func (a *BaseButtonGroup) EditorSetting(value interface{}) *BaseButtonGroup {
	a.Set("editorSetting", value)
	return a
}

func (a *BaseButtonGroup) Hidden(value interface{}) *BaseButtonGroup {
	a.Set("hidden", value)
	return a
}

func (a *BaseButtonGroup) HiddenOn(value interface{}) *BaseButtonGroup {
	a.Set("hiddenOn", value)
	return a
}

func (a *BaseButtonGroup) Id(value interface{}) *BaseButtonGroup {
	a.Set("id", value)
	return a
}

func (a *BaseButtonGroup) OnEvent(value interface{}) *BaseButtonGroup {
	a.Set("onEvent", value)
	return a
}

func (a *BaseButtonGroup) Size(value interface{}) *BaseButtonGroup {
	a.Set("size", value)
	return a
}

func (a *BaseButtonGroup) Static(value interface{}) *BaseButtonGroup {
	a.Set("static", value)
	return a
}

func (a *BaseButtonGroup) StaticClassName(value interface{}) *BaseButtonGroup {
	a.Set("staticClassName", value)
	return a
}

func (a *BaseButtonGroup) StaticInputClassName(value interface{}) *BaseButtonGroup {
	a.Set("staticInputClassName", value)
	return a
}

func (a *BaseButtonGroup) StaticLabelClassName(value interface{}) *BaseButtonGroup {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *BaseButtonGroup) StaticOn(value interface{}) *BaseButtonGroup {
	a.Set("staticOn", value)
	return a
}

func (a *BaseButtonGroup) StaticPlaceholder(value interface{}) *BaseButtonGroup {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *BaseButtonGroup) StaticSchema(value interface{}) *BaseButtonGroup {
	a.Set("staticSchema", value)
	return a
}

func (a *BaseButtonGroup) Style(value interface{}) *BaseButtonGroup {
	a.Set("style", value)
	return a
}

func (a *BaseButtonGroup) Testid(value interface{}) *BaseButtonGroup {
	a.Set("testid", value)
	return a
}

func (a *BaseButtonGroup) Tiled(value interface{}) *BaseButtonGroup {
	a.Set("tiled", value)
	return a
}

func (a *BaseButtonGroup) Type(value interface{}) *BaseButtonGroup {
	a.Set("type", value)
	return a
}

func (a *BaseButtonGroup) UseMobileUI(value interface{}) *BaseButtonGroup {
	a.Set("useMobileUI", value)
	return a
}

func (a *BaseButtonGroup) Vertical(value interface{}) *BaseButtonGroup {
	a.Set("vertical", value)
	return a
}

func (a *BaseButtonGroup) Visible(value interface{}) *BaseButtonGroup {
	a.Set("visible", value)
	return a
}

func (a *BaseButtonGroup) VisibleOn(value interface{}) *BaseButtonGroup {
	a.Set("visibleOn", value)
	return a
}
