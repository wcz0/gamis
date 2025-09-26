package renderers

type NavOverflow struct {
	*BaseRenderer
}

func NewNavOverflow() *NavOverflow {
	a := &NavOverflow{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *NavOverflow) Set(name string, value interface{}) *NavOverflow {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *NavOverflow) Enable(value interface{}) *NavOverflow {
	a.Set("enable", value)
	return a
}

func (a *NavOverflow) ItemWidth(value interface{}) *NavOverflow {
	a.Set("itemWidth", value)
	return a
}

func (a *NavOverflow) MaxVisibleCount(value interface{}) *NavOverflow {
	a.Set("maxVisibleCount", value)
	return a
}

func (a *NavOverflow) Mode(value interface{}) *NavOverflow {
	a.Set("mode", value)
	return a
}

func (a *NavOverflow) OverflowClassName(value interface{}) *NavOverflow {
	a.Set("overflowClassName", value)
	return a
}

func (a *NavOverflow) OverflowIndicator(value interface{}) *NavOverflow {
	a.Set("overflowIndicator", value)
	return a
}

func (a *NavOverflow) OverflowLabel(value interface{}) *NavOverflow {
	a.Set("overflowLabel", value)
	return a
}

func (a *NavOverflow) OverflowListClassName(value interface{}) *NavOverflow {
	a.Set("overflowListClassName", value)
	return a
}

func (a *NavOverflow) OverflowPopoverClassName(value interface{}) *NavOverflow {
	a.Set("overflowPopoverClassName", value)
	return a
}

func (a *NavOverflow) OverflowSuffix(value interface{}) *NavOverflow {
	a.Set("overflowSuffix", value)
	return a
}

func (a *NavOverflow) Style(value interface{}) *NavOverflow {
	a.Set("style", value)
	return a
}

func (a *NavOverflow) WrapperComponent(value interface{}) *NavOverflow {
	a.Set("wrapperComponent", value)
	return a
}
