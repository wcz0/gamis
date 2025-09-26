package renderers

type SchemaCopyable struct {
	*BaseRenderer
}

func NewSchemaCopyable() *SchemaCopyable {
	a := &SchemaCopyable{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *SchemaCopyable) Set(name string, value interface{}) *SchemaCopyable {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *SchemaCopyable) Content(value interface{}) *SchemaCopyable {
	a.Set("content", value)
	return a
}

func (a *SchemaCopyable) Icon(value interface{}) *SchemaCopyable {
	a.Set("icon", value)
	return a
}

func (a *SchemaCopyable) Tooltip(value interface{}) *SchemaCopyable {
	a.Set("tooltip", value)
	return a
}
