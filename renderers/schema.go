package renderers

type Schema struct {
	*BaseRenderer
}

func NewSchema() *Schema {
	a := &Schema{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *Schema) Set(name string, value interface{}) *Schema {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}
