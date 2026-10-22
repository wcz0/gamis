package renderers

type SchemaQuickEdit struct {
	*BaseRenderer
}

func NewSchemaQuickEdit() *SchemaQuickEdit {
	a := &SchemaQuickEdit{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *SchemaQuickEdit) Set(name string, value interface{}) *SchemaQuickEdit {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *SchemaQuickEdit) Body(value interface{}) *SchemaQuickEdit {
	a.Set("body", value)
	return a
}

func (a *SchemaQuickEdit) Icon(value interface{}) *SchemaQuickEdit {
	a.Set("icon", value)
	return a
}

func (a *SchemaQuickEdit) Mode(value interface{}) *SchemaQuickEdit {
	a.Set("mode", value)
	return a
}

func (a *SchemaQuickEdit) Reload(value interface{}) *SchemaQuickEdit {
	a.Set("reload", value)
	return a
}

func (a *SchemaQuickEdit) ResetOnFailed(value interface{}) *SchemaQuickEdit {
	a.Set("resetOnFailed", value)
	return a
}

func (a *SchemaQuickEdit) SaveImmediately(value interface{}) *SchemaQuickEdit {
	a.Set("saveImmediately", value)
	return a
}
