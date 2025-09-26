package renderers

type IconChecked struct {
	*BaseRenderer
}

func NewIconChecked() *IconChecked {
	a := &IconChecked{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *IconChecked) Set(name string, value interface{}) *IconChecked {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *IconChecked) Id(value interface{}) *IconChecked {
	a.Set("id", value)
	return a
}

func (a *IconChecked) Name(value interface{}) *IconChecked {
	a.Set("name", value)
	return a
}

func (a *IconChecked) Svg(value interface{}) *IconChecked {
	a.Set("svg", value)
	return a
}
