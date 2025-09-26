package renderers

type InputTimeRange struct {
	*BaseRenderer
}

func NewInputTimeRange() *InputTimeRange {
	a := &InputTimeRange{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "input-time-range")
	return a
}

func (a *InputTimeRange) Set(name string, value interface{}) *InputTimeRange {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *InputTimeRange) AutoFill(value interface{}) *InputTimeRange {
	a.Set("autoFill", value)
	return a
}

func (a *InputTimeRange) ClassName(value interface{}) *InputTimeRange {
	a.Set("className", value)
	return a
}

func (a *InputTimeRange) Clearable(value interface{}) *InputTimeRange {
	a.Set("clearable", value)
	return a
}

func (a *InputTimeRange) Description(value interface{}) *InputTimeRange {
	a.Set("description", value)
	return a
}

func (a *InputTimeRange) Disabled(value interface{}) *InputTimeRange {
	a.Set("disabled", value)
	return a
}

func (a *InputTimeRange) DisabledOn(value interface{}) *InputTimeRange {
	a.Set("disabledOn", value)
	return a
}

func (a *InputTimeRange) Embed(value interface{}) *InputTimeRange {
	a.Set("embed", value)
	return a
}

func (a *InputTimeRange) Format(value interface{}) *InputTimeRange {
	a.Set("format", value)
	return a
}

func (a *InputTimeRange) Inline(value interface{}) *InputTimeRange {
	a.Set("inline", value)
	return a
}

func (a *InputTimeRange) InputClassName(value interface{}) *InputTimeRange {
	a.Set("inputClassName", value)
	return a
}

func (a *InputTimeRange) InputFormat(value interface{}) *InputTimeRange {
	a.Set("inputFormat", value)
	return a
}

func (a *InputTimeRange) Label(value interface{}) *InputTimeRange {
	a.Set("label", value)
	return a
}

func (a *InputTimeRange) LabelAlign(value interface{}) *InputTimeRange {
	a.Set("labelAlign", value)
	return a
}

func (a *InputTimeRange) LabelClassName(value interface{}) *InputTimeRange {
	a.Set("labelClassName", value)
	return a
}

func (a *InputTimeRange) LabelRemark(value interface{}) *InputTimeRange {
	a.Set("labelRemark", value)
	return a
}

func (a *InputTimeRange) Name(value interface{}) *InputTimeRange {
	a.Set("name", value)
	return a
}

func (a *InputTimeRange) Placeholder(value interface{}) *InputTimeRange {
	a.Set("placeholder", value)
	return a
}

func (a *InputTimeRange) Required(value interface{}) *InputTimeRange {
	a.Set("required", value)
	return a
}

func (a *InputTimeRange) RequiredOn(value interface{}) *InputTimeRange {
	a.Set("requiredOn", value)
	return a
}

func (a *InputTimeRange) SubmitOnChange(value interface{}) *InputTimeRange {
	a.Set("submitOnChange", value)
	return a
}

func (a *InputTimeRange) TimeFormat(value interface{}) *InputTimeRange {
	a.Set("timeFormat", value)
	return a
}

func (a *InputTimeRange) Type(value interface{}) *InputTimeRange {
	a.Set("type", value)
	return a
}

func (a *InputTimeRange) ValidateApi(value interface{}) *InputTimeRange {
	a.Set("validateApi", value)
	return a
}

func (a *InputTimeRange) Validations(value interface{}) *InputTimeRange {
	a.Set("validations", value)
	return a
}

func (a *InputTimeRange) Value(value interface{}) *InputTimeRange {
	a.Set("value", value)
	return a
}

func (a *InputTimeRange) Visible(value interface{}) *InputTimeRange {
	a.Set("visible", value)
	return a
}

func (a *InputTimeRange) VisibleOn(value interface{}) *InputTimeRange {
	a.Set("visibleOn", value)
	return a
}
