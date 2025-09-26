package renderers

type SchemaMessage struct {
	*BaseRenderer
}

func NewSchemaMessage() *SchemaMessage {
	a := &SchemaMessage{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *SchemaMessage) Set(name string, value interface{}) *SchemaMessage {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *SchemaMessage) FetchFailed(value interface{}) *SchemaMessage {
	a.Set("fetchFailed", value)
	return a
}

func (a *SchemaMessage) FetchSuccess(value interface{}) *SchemaMessage {
	a.Set("fetchSuccess", value)
	return a
}

func (a *SchemaMessage) SaveFailed(value interface{}) *SchemaMessage {
	a.Set("saveFailed", value)
	return a
}

func (a *SchemaMessage) SaveSuccess(value interface{}) *SchemaMessage {
	a.Set("saveSuccess", value)
	return a
}
