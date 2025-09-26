package renderers

type Component struct {
	*BaseRenderer
}

func NewComponent() *Component {
	a := &Component{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *Component) Set(name string, value interface{}) *Component {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Component) SetType(value interface{}) *Component {
	a.Set("type", value)
	return a
}
