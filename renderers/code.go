package renderers

type Code struct {
	*BaseRenderer
}

func NewCode() *Code {
	a := &Code{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "code")
	return a
}

func (a *Code) Set(name string, value interface{}) *Code {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Code) ClassName(value interface{}) *Code {
	a.Set("className", value)
	return a
}

func (a *Code) EditorTheme(value interface{}) *Code {
	a.Set("editorTheme", value)
	return a
}

func (a *Code) Language(value interface{}) *Code {
	a.Set("language", value)
	return a
}

func (a *Code) Name(value interface{}) *Code {
	a.Set("name", value)
	return a
}

func (a *Code) TabSize(value interface{}) *Code {
	a.Set("tabSize", value)
	return a
}

func (a *Code) Type(value interface{}) *Code {
	a.Set("type", value)
	return a
}

func (a *Code) Value(value interface{}) *Code {
	a.Set("value", value)
	return a
}

func (a *Code) WordWrap(value interface{}) *Code {
	a.Set("wordWrap", value)
	return a
}
