package renderers

type InputKVS struct {
	*BaseRenderer
}

func NewInputKVS() *InputKVS {
	a := &InputKVS{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "input-kvs")
	return a
}

func (a *InputKVS) Set(name string, value interface{}) *InputKVS {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *InputKVS) AutoFill(value interface{}) *InputKVS {
	a.Set("autoFill", value)
	return a
}

func (a *InputKVS) ClassName(value interface{}) *InputKVS {
	a.Set("className", value)
	return a
}

func (a *InputKVS) Description(value interface{}) *InputKVS {
	a.Set("description", value)
	return a
}

func (a *InputKVS) Disabled(value interface{}) *InputKVS {
	a.Set("disabled", value)
	return a
}

func (a *InputKVS) DisabledOn(value interface{}) *InputKVS {
	a.Set("disabledOn", value)
	return a
}

func (a *InputKVS) Inline(value interface{}) *InputKVS {
	a.Set("inline", value)
	return a
}

func (a *InputKVS) InputClassName(value interface{}) *InputKVS {
	a.Set("inputClassName", value)
	return a
}

func (a *InputKVS) Label(value interface{}) *InputKVS {
	a.Set("label", value)
	return a
}

func (a *InputKVS) LabelAlign(value interface{}) *InputKVS {
	a.Set("labelAlign", value)
	return a
}

func (a *InputKVS) LabelClassName(value interface{}) *InputKVS {
	a.Set("labelClassName", value)
	return a
}

func (a *InputKVS) LabelRemark(value interface{}) *InputKVS {
	a.Set("labelRemark", value)
	return a
}

func (a *InputKVS) Name(value interface{}) *InputKVS {
	a.Set("name", value)
	return a
}

func (a *InputKVS) Placeholder(value interface{}) *InputKVS {
	a.Set("placeholder", value)
	return a
}

func (a *InputKVS) Required(value interface{}) *InputKVS {
	a.Set("required", value)
	return a
}

func (a *InputKVS) RequiredOn(value interface{}) *InputKVS {
	a.Set("requiredOn", value)
	return a
}

func (a *InputKVS) SubmitOnChange(value interface{}) *InputKVS {
	a.Set("submitOnChange", value)
	return a
}

func (a *InputKVS) Type(value interface{}) *InputKVS {
	a.Set("type", value)
	return a
}

func (a *InputKVS) ValidateApi(value interface{}) *InputKVS {
	a.Set("validateApi", value)
	return a
}

func (a *InputKVS) Validations(value interface{}) *InputKVS {
	a.Set("validations", value)
	return a
}

func (a *InputKVS) Value(value interface{}) *InputKVS {
	a.Set("value", value)
	return a
}

func (a *InputKVS) Visible(value interface{}) *InputKVS {
	a.Set("visible", value)
	return a
}

func (a *InputKVS) VisibleOn(value interface{}) *InputKVS {
	a.Set("visibleOn", value)
	return a
}
