package renderers

type BaseCollapse struct {
	*BaseRenderer
}

func NewBaseCollapse() *BaseCollapse {
	a := &BaseCollapse{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *BaseCollapse) Set(name string, value interface{}) *BaseCollapse {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *BaseCollapse) BodyClassName(value interface{}) *BaseCollapse {
	a.Set("bodyClassName", value)
	return a
}

func (a *BaseCollapse) ClassName(value interface{}) *BaseCollapse {
	a.Set("className", value)
	return a
}

func (a *BaseCollapse) Collapsable(value interface{}) *BaseCollapse {
	a.Set("collapsable", value)
	return a
}

func (a *BaseCollapse) CollapseHeader(value interface{}) *BaseCollapse {
	a.Set("collapseHeader", value)
	return a
}

func (a *BaseCollapse) Collapsed(value interface{}) *BaseCollapse {
	a.Set("collapsed", value)
	return a
}

func (a *BaseCollapse) Disabled(value interface{}) *BaseCollapse {
	a.Set("disabled", value)
	return a
}

func (a *BaseCollapse) DisabledOn(value interface{}) *BaseCollapse {
	a.Set("disabledOn", value)
	return a
}

func (a *BaseCollapse) DivideLine(value interface{}) *BaseCollapse {
	a.Set("divideLine", value)
	return a
}

func (a *BaseCollapse) EditorSetting(value interface{}) *BaseCollapse {
	a.Set("editorSetting", value)
	return a
}

func (a *BaseCollapse) ExpandIcon(value interface{}) *BaseCollapse {
	a.Set("expandIcon", value)
	return a
}

func (a *BaseCollapse) Header(value interface{}) *BaseCollapse {
	a.Set("header", value)
	return a
}

func (a *BaseCollapse) HeaderPosition(value interface{}) *BaseCollapse {
	a.Set("headerPosition", value)
	return a
}

func (a *BaseCollapse) HeadingClassName(value interface{}) *BaseCollapse {
	a.Set("headingClassName", value)
	return a
}

func (a *BaseCollapse) Hidden(value interface{}) *BaseCollapse {
	a.Set("hidden", value)
	return a
}

func (a *BaseCollapse) HiddenOn(value interface{}) *BaseCollapse {
	a.Set("hiddenOn", value)
	return a
}

func (a *BaseCollapse) Id(value interface{}) *BaseCollapse {
	a.Set("id", value)
	return a
}

func (a *BaseCollapse) Key(value interface{}) *BaseCollapse {
	a.Set("key", value)
	return a
}

func (a *BaseCollapse) MountOnEnter(value interface{}) *BaseCollapse {
	a.Set("mountOnEnter", value)
	return a
}

func (a *BaseCollapse) OnEvent(value interface{}) *BaseCollapse {
	a.Set("onEvent", value)
	return a
}

func (a *BaseCollapse) ShowArrow(value interface{}) *BaseCollapse {
	a.Set("showArrow", value)
	return a
}

func (a *BaseCollapse) Size(value interface{}) *BaseCollapse {
	a.Set("size", value)
	return a
}

func (a *BaseCollapse) Static(value interface{}) *BaseCollapse {
	a.Set("static", value)
	return a
}

func (a *BaseCollapse) StaticClassName(value interface{}) *BaseCollapse {
	a.Set("staticClassName", value)
	return a
}

func (a *BaseCollapse) StaticInputClassName(value interface{}) *BaseCollapse {
	a.Set("staticInputClassName", value)
	return a
}

func (a *BaseCollapse) StaticLabelClassName(value interface{}) *BaseCollapse {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *BaseCollapse) StaticOn(value interface{}) *BaseCollapse {
	a.Set("staticOn", value)
	return a
}

func (a *BaseCollapse) StaticPlaceholder(value interface{}) *BaseCollapse {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *BaseCollapse) StaticSchema(value interface{}) *BaseCollapse {
	a.Set("staticSchema", value)
	return a
}

func (a *BaseCollapse) Style(value interface{}) *BaseCollapse {
	a.Set("style", value)
	return a
}

func (a *BaseCollapse) Testid(value interface{}) *BaseCollapse {
	a.Set("testid", value)
	return a
}

func (a *BaseCollapse) Type(value interface{}) *BaseCollapse {
	a.Set("type", value)
	return a
}

func (a *BaseCollapse) UnmountOnExit(value interface{}) *BaseCollapse {
	a.Set("unmountOnExit", value)
	return a
}

func (a *BaseCollapse) UseMobileUI(value interface{}) *BaseCollapse {
	a.Set("useMobileUI", value)
	return a
}

func (a *BaseCollapse) Visible(value interface{}) *BaseCollapse {
	a.Set("visible", value)
	return a
}

func (a *BaseCollapse) VisibleOn(value interface{}) *BaseCollapse {
	a.Set("visibleOn", value)
	return a
}
