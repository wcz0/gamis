package renderers

type InputKV struct {
	*BaseRenderer
}

func NewInputKV() *InputKV {
	a := &InputKV{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "input-kv")
	return a
}

func (a *InputKV) Set(name string, value interface{}) *InputKV {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *InputKV) AutoFill(value interface{}) *InputKV {
	a.Set("autoFill", value)
	return a
}

func (a *InputKV) ClassName(value interface{}) *InputKV {
	a.Set("className", value)
	return a
}

func (a *InputKV) DefaultValue(value interface{}) *InputKV {
	a.Set("defaultValue", value)
	return a
}

func (a *InputKV) Description(value interface{}) *InputKV {
	a.Set("description", value)
	return a
}

func (a *InputKV) Disabled(value interface{}) *InputKV {
	a.Set("disabled", value)
	return a
}

func (a *InputKV) DisabledOn(value interface{}) *InputKV {
	a.Set("disabledOn", value)
	return a
}

func (a *InputKV) Draggable(value interface{}) *InputKV {
	a.Set("draggable", value)
	return a
}

func (a *InputKV) Inline(value interface{}) *InputKV {
	a.Set("inline", value)
	return a
}

func (a *InputKV) InputClassName(value interface{}) *InputKV {
	a.Set("inputClassName", value)
	return a
}

func (a *InputKV) KeyPlaceholder(value interface{}) *InputKV {
	a.Set("keyPlaceholder", value)
	return a
}

func (a *InputKV) Label(value interface{}) *InputKV {
	a.Set("label", value)
	return a
}

func (a *InputKV) LabelAlign(value interface{}) *InputKV {
	a.Set("labelAlign", value)
	return a
}

func (a *InputKV) LabelClassName(value interface{}) *InputKV {
	a.Set("labelClassName", value)
	return a
}

func (a *InputKV) LabelRemark(value interface{}) *InputKV {
	a.Set("labelRemark", value)
	return a
}

func (a *InputKV) Name(value interface{}) *InputKV {
	a.Set("name", value)
	return a
}

func (a *InputKV) Placeholder(value interface{}) *InputKV {
	a.Set("placeholder", value)
	return a
}

func (a *InputKV) Required(value interface{}) *InputKV {
	a.Set("required", value)
	return a
}

func (a *InputKV) RequiredOn(value interface{}) *InputKV {
	a.Set("requiredOn", value)
	return a
}

func (a *InputKV) SubmitOnChange(value interface{}) *InputKV {
	a.Set("submitOnChange", value)
	return a
}

func (a *InputKV) Type(value interface{}) *InputKV {
	a.Set("type", value)
	return a
}

func (a *InputKV) ValidateApi(value interface{}) *InputKV {
	a.Set("validateApi", value)
	return a
}

func (a *InputKV) Validations(value interface{}) *InputKV {
	a.Set("validations", value)
	return a
}

func (a *InputKV) Value(value interface{}) *InputKV {
	a.Set("value", value)
	return a
}

func (a *InputKV) ValuePlaceholder(value interface{}) *InputKV {
	a.Set("valuePlaceholder", value)
	return a
}

func (a *InputKV) ValueType(value interface{}) *InputKV {
	a.Set("valueType", value)
	return a
}

func (a *InputKV) Visible(value interface{}) *InputKV {
	a.Set("visible", value)
	return a
}

func (a *InputKV) VisibleOn(value interface{}) *InputKV {
	a.Set("visibleOn", value)
	return a
}
