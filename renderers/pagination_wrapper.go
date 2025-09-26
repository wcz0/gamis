package renderers

type PaginationWrapper struct {
	*BaseRenderer
}

func NewPaginationWrapper() *PaginationWrapper {
	a := &PaginationWrapper{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "pagination-wrapper")
	return a
}

func (a *PaginationWrapper) Set(name string, value interface{}) *PaginationWrapper {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *PaginationWrapper) Body(value interface{}) *PaginationWrapper {
	a.Set("body", value)
	return a
}

func (a *PaginationWrapper) ClassName(value interface{}) *PaginationWrapper {
	a.Set("className", value)
	return a
}

func (a *PaginationWrapper) Disabled(value interface{}) *PaginationWrapper {
	a.Set("disabled", value)
	return a
}

func (a *PaginationWrapper) DisabledOn(value interface{}) *PaginationWrapper {
	a.Set("disabledOn", value)
	return a
}

func (a *PaginationWrapper) EditorSetting(value interface{}) *PaginationWrapper {
	a.Set("editorSetting", value)
	return a
}

func (a *PaginationWrapper) Hidden(value interface{}) *PaginationWrapper {
	a.Set("hidden", value)
	return a
}

func (a *PaginationWrapper) HiddenOn(value interface{}) *PaginationWrapper {
	a.Set("hiddenOn", value)
	return a
}

func (a *PaginationWrapper) Id(value interface{}) *PaginationWrapper {
	a.Set("id", value)
	return a
}

func (a *PaginationWrapper) InputName(value interface{}) *PaginationWrapper {
	a.Set("inputName", value)
	return a
}

func (a *PaginationWrapper) MaxButtons(value interface{}) *PaginationWrapper {
	a.Set("maxButtons", value)
	return a
}

func (a *PaginationWrapper) OnEvent(value interface{}) *PaginationWrapper {
	a.Set("onEvent", value)
	return a
}

func (a *PaginationWrapper) OutputName(value interface{}) *PaginationWrapper {
	a.Set("outputName", value)
	return a
}

func (a *PaginationWrapper) PerPage(value interface{}) *PaginationWrapper {
	a.Set("perPage", value)
	return a
}

func (a *PaginationWrapper) Position(value interface{}) *PaginationWrapper {
	a.Set("position", value)
	return a
}

func (a *PaginationWrapper) ShowPageInput(value interface{}) *PaginationWrapper {
	a.Set("showPageInput", value)
	return a
}

func (a *PaginationWrapper) Static(value interface{}) *PaginationWrapper {
	a.Set("static", value)
	return a
}

func (a *PaginationWrapper) StaticClassName(value interface{}) *PaginationWrapper {
	a.Set("staticClassName", value)
	return a
}

func (a *PaginationWrapper) StaticInputClassName(value interface{}) *PaginationWrapper {
	a.Set("staticInputClassName", value)
	return a
}

func (a *PaginationWrapper) StaticLabelClassName(value interface{}) *PaginationWrapper {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *PaginationWrapper) StaticOn(value interface{}) *PaginationWrapper {
	a.Set("staticOn", value)
	return a
}

func (a *PaginationWrapper) StaticPlaceholder(value interface{}) *PaginationWrapper {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *PaginationWrapper) StaticSchema(value interface{}) *PaginationWrapper {
	a.Set("staticSchema", value)
	return a
}

func (a *PaginationWrapper) Style(value interface{}) *PaginationWrapper {
	a.Set("style", value)
	return a
}

func (a *PaginationWrapper) TestIdBuilder(value interface{}) *PaginationWrapper {
	a.Set("testIdBuilder", value)
	return a
}

func (a *PaginationWrapper) Testid(value interface{}) *PaginationWrapper {
	a.Set("testid", value)
	return a
}

func (a *PaginationWrapper) Type(value interface{}) *PaginationWrapper {
	a.Set("type", value)
	return a
}

func (a *PaginationWrapper) UseMobileUI(value interface{}) *PaginationWrapper {
	a.Set("useMobileUI", value)
	return a
}

func (a *PaginationWrapper) Visible(value interface{}) *PaginationWrapper {
	a.Set("visible", value)
	return a
}

func (a *PaginationWrapper) VisibleOn(value interface{}) *PaginationWrapper {
	a.Set("visibleOn", value)
	return a
}
