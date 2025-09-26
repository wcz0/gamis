package renderers

type SchemaPopOver struct {
	*BaseRenderer
}

func NewSchemaPopOver() *SchemaPopOver {
	a := &SchemaPopOver{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *SchemaPopOver) Set(name string, value interface{}) *SchemaPopOver {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *SchemaPopOver) Body(value interface{}) *SchemaPopOver {
	a.Set("body", value)
	return a
}

func (a *SchemaPopOver) ClassName(value interface{}) *SchemaPopOver {
	a.Set("className", value)
	return a
}

func (a *SchemaPopOver) Mode(value interface{}) *SchemaPopOver {
	a.Set("mode", value)
	return a
}

func (a *SchemaPopOver) Offset(value interface{}) *SchemaPopOver {
	a.Set("offset", value)
	return a
}

func (a *SchemaPopOver) PopOverClassName(value interface{}) *SchemaPopOver {
	a.Set("popOverClassName", value)
	return a
}

func (a *SchemaPopOver) PopOverEnableOn(value interface{}) *SchemaPopOver {
	a.Set("popOverEnableOn", value)
	return a
}

func (a *SchemaPopOver) Position(value interface{}) *SchemaPopOver {
	a.Set("position", value)
	return a
}

func (a *SchemaPopOver) ShowIcon(value interface{}) *SchemaPopOver {
	a.Set("showIcon", value)
	return a
}

func (a *SchemaPopOver) Size(value interface{}) *SchemaPopOver {
	a.Set("size", value)
	return a
}

func (a *SchemaPopOver) Title(value interface{}) *SchemaPopOver {
	a.Set("title", value)
	return a
}

func (a *SchemaPopOver) Trigger(value interface{}) *SchemaPopOver {
	a.Set("trigger", value)
	return a
}
