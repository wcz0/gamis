package renderers

type Option struct {
	*BaseRenderer
}

func NewOption() *Option {
	a := &Option{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *Option) Set(name string, value interface{}) *Option {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Option) Children(value interface{}) *Option {
	a.Set("children", value)
	return a
}

func (a *Option) Defer(value interface{}) *Option {
	a.Set("defer", value)
	return a
}

func (a *Option) DeferApi(value interface{}) *Option {
	a.Set("deferApi", value)
	return a
}

func (a *Option) Description(value interface{}) *Option {
	a.Set("description", value)
	return a
}

func (a *Option) Disabled(value interface{}) *Option {
	a.Set("disabled", value)
	return a
}

func (a *Option) Hidden(value interface{}) *Option {
	a.Set("hidden", value)
	return a
}

func (a *Option) Label(value interface{}) *Option {
	a.Set("label", value)
	return a
}

func (a *Option) Loaded(value interface{}) *Option {
	a.Set("loaded", value)
	return a
}

func (a *Option) Loading(value interface{}) *Option {
	a.Set("loading", value)
	return a
}

func (a *Option) ScopeLabel(value interface{}) *Option {
	a.Set("scopeLabel", value)
	return a
}

func (a *Option) Value(value interface{}) *Option {
	a.Set("value", value)
	return a
}

func (a *Option) Visible(value interface{}) *Option {
	a.Set("visible", value)
	return a
}
