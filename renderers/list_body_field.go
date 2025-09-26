package renderers

type ListBodyField struct {
	*BaseRenderer
}

func NewListBodyField() *ListBodyField {
	a := &ListBodyField{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *ListBodyField) Set(name string, value interface{}) *ListBodyField {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *ListBodyField) Copyable(value interface{}) *ListBodyField {
	a.Set("copyable", value)
	return a
}

func (a *ListBodyField) InnerClassName(value interface{}) *ListBodyField {
	a.Set("innerClassName", value)
	return a
}

func (a *ListBodyField) Label(value interface{}) *ListBodyField {
	a.Set("label", value)
	return a
}

func (a *ListBodyField) LabelClassName(value interface{}) *ListBodyField {
	a.Set("labelClassName", value)
	return a
}

func (a *ListBodyField) Name(value interface{}) *ListBodyField {
	a.Set("name", value)
	return a
}

func (a *ListBodyField) PopOver(value interface{}) *ListBodyField {
	a.Set("popOver", value)
	return a
}

func (a *ListBodyField) QuickEdit(value interface{}) *ListBodyField {
	a.Set("quickEdit", value)
	return a
}
