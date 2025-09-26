package renderers

type Custom struct {
	*BaseRenderer
}

func NewCustom() *Custom {
	a := &Custom{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "custom")
	return a
}

func (a *Custom) Set(name string, value interface{}) *Custom {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Custom) ClassName(value interface{}) *Custom {
	a.Set("className", value)
	return a
}

func (a *Custom) Html(value interface{}) *Custom {
	a.Set("html", value)
	return a
}

func (a *Custom) Id(value interface{}) *Custom {
	a.Set("id", value)
	return a
}

func (a *Custom) Inline(value interface{}) *Custom {
	a.Set("inline", value)
	return a
}

func (a *Custom) Name(value interface{}) *Custom {
	a.Set("name", value)
	return a
}

func (a *Custom) OnMount(value interface{}) *Custom {
	a.Set("onMount", value)
	return a
}

func (a *Custom) OnUnmount(value interface{}) *Custom {
	a.Set("onUnmount", value)
	return a
}

func (a *Custom) OnUpdate(value interface{}) *Custom {
	a.Set("onUpdate", value)
	return a
}

func (a *Custom) Type(value interface{}) *Custom {
	a.Set("type", value)
	return a
}
