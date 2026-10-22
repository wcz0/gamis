package renderers

type InputTextAddOn struct {
	*BaseRenderer
}

func NewInputTextAddOn() *InputTextAddOn {
	a := &InputTextAddOn{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *InputTextAddOn) Set(name string, value interface{}) *InputTextAddOn {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *InputTextAddOn) Position(value interface{}) *InputTextAddOn {
	a.Set("position", value)
	return a
}
