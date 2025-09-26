package renderers

type Avatar struct {
	*BaseRenderer
}

func NewAvatar() *Avatar {
	a := &Avatar{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "avatar")
	return a
}

func (a *Avatar) Set(name string, value interface{}) *Avatar {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Avatar) Alt(value interface{}) *Avatar {
	a.Set("alt", value)
	return a
}

func (a *Avatar) Badge(value interface{}) *Avatar {
	a.Set("badge", value)
	return a
}

func (a *Avatar) ClassName(value interface{}) *Avatar {
	a.Set("className", value)
	return a
}

func (a *Avatar) CrossOrigin(value interface{}) *Avatar {
	a.Set("crossOrigin", value)
	return a
}

func (a *Avatar) DefaultAvatar(value interface{}) *Avatar {
	a.Set("defaultAvatar", value)
	return a
}

func (a *Avatar) Disabled(value interface{}) *Avatar {
	a.Set("disabled", value)
	return a
}

func (a *Avatar) DisabledOn(value interface{}) *Avatar {
	a.Set("disabledOn", value)
	return a
}

func (a *Avatar) Draggable(value interface{}) *Avatar {
	a.Set("draggable", value)
	return a
}

func (a *Avatar) EditorSetting(value interface{}) *Avatar {
	a.Set("editorSetting", value)
	return a
}

func (a *Avatar) Fit(value interface{}) *Avatar {
	a.Set("fit", value)
	return a
}

func (a *Avatar) Gap(value interface{}) *Avatar {
	a.Set("gap", value)
	return a
}

func (a *Avatar) Hidden(value interface{}) *Avatar {
	a.Set("hidden", value)
	return a
}

func (a *Avatar) HiddenOn(value interface{}) *Avatar {
	a.Set("hiddenOn", value)
	return a
}

func (a *Avatar) Icon(value interface{}) *Avatar {
	a.Set("icon", value)
	return a
}

func (a *Avatar) Id(value interface{}) *Avatar {
	a.Set("id", value)
	return a
}

func (a *Avatar) OnError(value interface{}) *Avatar {
	a.Set("onError", value)
	return a
}

func (a *Avatar) OnEvent(value interface{}) *Avatar {
	a.Set("onEvent", value)
	return a
}

func (a *Avatar) Shape(value interface{}) *Avatar {
	a.Set("shape", value)
	return a
}

func (a *Avatar) Size(value interface{}) *Avatar {
	a.Set("size", value)
	return a
}

func (a *Avatar) Src(value interface{}) *Avatar {
	a.Set("src", value)
	return a
}

func (a *Avatar) Static(value interface{}) *Avatar {
	a.Set("static", value)
	return a
}

func (a *Avatar) StaticClassName(value interface{}) *Avatar {
	a.Set("staticClassName", value)
	return a
}

func (a *Avatar) StaticInputClassName(value interface{}) *Avatar {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Avatar) StaticLabelClassName(value interface{}) *Avatar {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Avatar) StaticOn(value interface{}) *Avatar {
	a.Set("staticOn", value)
	return a
}

func (a *Avatar) StaticPlaceholder(value interface{}) *Avatar {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Avatar) StaticSchema(value interface{}) *Avatar {
	a.Set("staticSchema", value)
	return a
}

func (a *Avatar) Style(value interface{}) *Avatar {
	a.Set("style", value)
	return a
}

func (a *Avatar) TestIdBuilder(value interface{}) *Avatar {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Avatar) Testid(value interface{}) *Avatar {
	a.Set("testid", value)
	return a
}

func (a *Avatar) Text(value interface{}) *Avatar {
	a.Set("text", value)
	return a
}

func (a *Avatar) Type(value interface{}) *Avatar {
	a.Set("type", value)
	return a
}

func (a *Avatar) UseMobileUI(value interface{}) *Avatar {
	a.Set("useMobileUI", value)
	return a
}

func (a *Avatar) Visible(value interface{}) *Avatar {
	a.Set("visible", value)
	return a
}

func (a *Avatar) VisibleOn(value interface{}) *Avatar {
	a.Set("visibleOn", value)
	return a
}
