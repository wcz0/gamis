package renderers

type Log struct {
	*BaseRenderer
}

func NewLog() *Log {
	a := &Log{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "log")
	return a
}

func (a *Log) Set(name string, value interface{}) *Log {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Log) AutoScroll(value interface{}) *Log {
	a.Set("autoScroll", value)
	return a
}

func (a *Log) ClassName(value interface{}) *Log {
	a.Set("className", value)
	return a
}

func (a *Log) DisableColor(value interface{}) *Log {
	a.Set("disableColor", value)
	return a
}

func (a *Log) Encoding(value interface{}) *Log {
	a.Set("encoding", value)
	return a
}

func (a *Log) Height(value interface{}) *Log {
	a.Set("height", value)
	return a
}

func (a *Log) MaxLength(value interface{}) *Log {
	a.Set("maxLength", value)
	return a
}

func (a *Log) Operation(value interface{}) *Log {
	a.Set("operation", value)
	return a
}

func (a *Log) Placeholder(value interface{}) *Log {
	a.Set("placeholder", value)
	return a
}

func (a *Log) RowHeight(value interface{}) *Log {
	a.Set("rowHeight", value)
	return a
}

func (a *Log) Source(value interface{}) *Log {
	a.Set("source", value)
	return a
}

func (a *Log) Type(value interface{}) *Log {
	a.Set("type", value)
	return a
}
