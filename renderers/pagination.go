package renderers

type Pagination struct {
	*BaseRenderer
}

func NewPagination() *Pagination {
	a := &Pagination{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "pagination")
	return a
}

func (a *Pagination) Set(name string, value interface{}) *Pagination {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Pagination) ActivePage(value interface{}) *Pagination {
	a.Set("activePage", value)
	return a
}

func (a *Pagination) ClassName(value interface{}) *Pagination {
	a.Set("className", value)
	return a
}

func (a *Pagination) Disabled(value interface{}) *Pagination {
	a.Set("disabled", value)
	return a
}

func (a *Pagination) DisabledOn(value interface{}) *Pagination {
	a.Set("disabledOn", value)
	return a
}

func (a *Pagination) EditorSetting(value interface{}) *Pagination {
	a.Set("editorSetting", value)
	return a
}

func (a *Pagination) HasNext(value interface{}) *Pagination {
	a.Set("hasNext", value)
	return a
}

func (a *Pagination) Hidden(value interface{}) *Pagination {
	a.Set("hidden", value)
	return a
}

func (a *Pagination) HiddenOn(value interface{}) *Pagination {
	a.Set("hiddenOn", value)
	return a
}

func (a *Pagination) Id(value interface{}) *Pagination {
	a.Set("id", value)
	return a
}

func (a *Pagination) Layout(value interface{}) *Pagination {
	a.Set("layout", value)
	return a
}

func (a *Pagination) MaxButtons(value interface{}) *Pagination {
	a.Set("maxButtons", value)
	return a
}

func (a *Pagination) Mode(value interface{}) *Pagination {
	a.Set("mode", value)
	return a
}

func (a *Pagination) OnEvent(value interface{}) *Pagination {
	a.Set("onEvent", value)
	return a
}

func (a *Pagination) PerPage(value interface{}) *Pagination {
	a.Set("perPage", value)
	return a
}

func (a *Pagination) PerPageAvailable(value interface{}) *Pagination {
	a.Set("perPageAvailable", value)
	return a
}

func (a *Pagination) PopOverContainerSelector(value interface{}) *Pagination {
	a.Set("popOverContainerSelector", value)
	return a
}

func (a *Pagination) ShowPageInput(value interface{}) *Pagination {
	a.Set("showPageInput", value)
	return a
}

func (a *Pagination) ShowPerPage(value interface{}) *Pagination {
	a.Set("showPerPage", value)
	return a
}

func (a *Pagination) Static(value interface{}) *Pagination {
	a.Set("static", value)
	return a
}

func (a *Pagination) StaticClassName(value interface{}) *Pagination {
	a.Set("staticClassName", value)
	return a
}

func (a *Pagination) StaticInputClassName(value interface{}) *Pagination {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Pagination) StaticLabelClassName(value interface{}) *Pagination {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Pagination) StaticOn(value interface{}) *Pagination {
	a.Set("staticOn", value)
	return a
}

func (a *Pagination) StaticPlaceholder(value interface{}) *Pagination {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Pagination) StaticSchema(value interface{}) *Pagination {
	a.Set("staticSchema", value)
	return a
}

func (a *Pagination) Style(value interface{}) *Pagination {
	a.Set("style", value)
	return a
}

func (a *Pagination) TestIdBuilder(value interface{}) *Pagination {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Pagination) Testid(value interface{}) *Pagination {
	a.Set("testid", value)
	return a
}

func (a *Pagination) Total(value interface{}) *Pagination {
	a.Set("total", value)
	return a
}

func (a *Pagination) Type(value interface{}) *Pagination {
	a.Set("type", value)
	return a
}

func (a *Pagination) UseMobileUI(value interface{}) *Pagination {
	a.Set("useMobileUI", value)
	return a
}

func (a *Pagination) Visible(value interface{}) *Pagination {
	a.Set("visible", value)
	return a
}

func (a *Pagination) VisibleOn(value interface{}) *Pagination {
	a.Set("visibleOn", value)
	return a
}
