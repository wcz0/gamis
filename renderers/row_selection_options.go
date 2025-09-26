package renderers

type RowSelectionOptions struct {
	*BaseRenderer
}

func NewRowSelectionOptions() *RowSelectionOptions {
	a := &RowSelectionOptions{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *RowSelectionOptions) Set(name string, value interface{}) *RowSelectionOptions {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *RowSelectionOptions) Key(value interface{}) *RowSelectionOptions {
	a.Set("key", value)
	return a
}

func (a *RowSelectionOptions) Text(value interface{}) *RowSelectionOptions {
	a.Set("text", value)
	return a
}
