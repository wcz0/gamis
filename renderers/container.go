package renderers

type Container struct {
	*BaseRenderer
}

func NewContainer() *Container {
	a := &Container{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "container")
	return a
}

func (a *Container) Set(name string, value interface{}) *Container {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Container) Body(value interface{}) *Container {
	a.Set("body", value)
	return a
}

func (a *Container) BodyClassName(value interface{}) *Container {
	a.Set("bodyClassName", value)
	return a
}

func (a *Container) ClassName(value interface{}) *Container {
	a.Set("className", value)
	return a
}

func (a *Container) Disabled(value interface{}) *Container {
	a.Set("disabled", value)
	return a
}

func (a *Container) DisabledOn(value interface{}) *Container {
	a.Set("disabledOn", value)
	return a
}

func (a *Container) Draggable(value interface{}) *Container {
	a.Set("draggable", value)
	return a
}

func (a *Container) DraggableConfig(value interface{}) *Container {
	a.Set("draggableConfig", value)
	return a
}

func (a *Container) EditorSetting(value interface{}) *Container {
	a.Set("editorSetting", value)
	return a
}

func (a *Container) Hidden(value interface{}) *Container {
	a.Set("hidden", value)
	return a
}

func (a *Container) HiddenOn(value interface{}) *Container {
	a.Set("hiddenOn", value)
	return a
}

func (a *Container) Id(value interface{}) *Container {
	a.Set("id", value)
	return a
}

func (a *Container) OnEvent(value interface{}) *Container {
	a.Set("onEvent", value)
	return a
}

func (a *Container) Static(value interface{}) *Container {
	a.Set("static", value)
	return a
}

func (a *Container) StaticClassName(value interface{}) *Container {
	a.Set("staticClassName", value)
	return a
}

func (a *Container) StaticInputClassName(value interface{}) *Container {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Container) StaticLabelClassName(value interface{}) *Container {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Container) StaticOn(value interface{}) *Container {
	a.Set("staticOn", value)
	return a
}

func (a *Container) StaticPlaceholder(value interface{}) *Container {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Container) StaticSchema(value interface{}) *Container {
	a.Set("staticSchema", value)
	return a
}

func (a *Container) Style(value interface{}) *Container {
	a.Set("style", value)
	return a
}

func (a *Container) TestIdBuilder(value interface{}) *Container {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Container) Testid(value interface{}) *Container {
	a.Set("testid", value)
	return a
}

func (a *Container) Type(value interface{}) *Container {
	a.Set("type", value)
	return a
}

func (a *Container) UseMobileUI(value interface{}) *Container {
	a.Set("useMobileUI", value)
	return a
}

func (a *Container) Visible(value interface{}) *Container {
	a.Set("visible", value)
	return a
}

func (a *Container) VisibleOn(value interface{}) *Container {
	a.Set("visibleOn", value)
	return a
}

func (a *Container) WrapperBody(value interface{}) *Container {
	a.Set("wrapperBody", value)
	return a
}

func (a *Container) WrapperComponent(value interface{}) *Container {
	a.Set("wrapperComponent", value)
	return a
}
