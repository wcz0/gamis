package renderers

type InputDatetimeRange struct {
	*BaseRenderer
}

func NewInputDatetimeRange() *InputDatetimeRange {
	a := &InputDatetimeRange{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "input-datetime-range")
	return a
}

func (a *InputDatetimeRange) Set(name string, value interface{}) *InputDatetimeRange {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *InputDatetimeRange) AutoFill(value interface{}) *InputDatetimeRange {
	a.Set("autoFill", value)
	return a
}

func (a *InputDatetimeRange) ClassName(value interface{}) *InputDatetimeRange {
	a.Set("className", value)
	return a
}

func (a *InputDatetimeRange) Clearable(value interface{}) *InputDatetimeRange {
	a.Set("clearable", value)
	return a
}

func (a *InputDatetimeRange) Description(value interface{}) *InputDatetimeRange {
	a.Set("description", value)
	return a
}

func (a *InputDatetimeRange) Disabled(value interface{}) *InputDatetimeRange {
	a.Set("disabled", value)
	return a
}

func (a *InputDatetimeRange) DisabledOn(value interface{}) *InputDatetimeRange {
	a.Set("disabledOn", value)
	return a
}

func (a *InputDatetimeRange) Format(value interface{}) *InputDatetimeRange {
	a.Set("format", value)
	return a
}

func (a *InputDatetimeRange) Inline(value interface{}) *InputDatetimeRange {
	a.Set("inline", value)
	return a
}

func (a *InputDatetimeRange) InputClassName(value interface{}) *InputDatetimeRange {
	a.Set("inputClassName", value)
	return a
}

func (a *InputDatetimeRange) InputFormat(value interface{}) *InputDatetimeRange {
	a.Set("inputFormat", value)
	return a
}

func (a *InputDatetimeRange) Label(value interface{}) *InputDatetimeRange {
	a.Set("label", value)
	return a
}

func (a *InputDatetimeRange) LabelAlign(value interface{}) *InputDatetimeRange {
	a.Set("labelAlign", value)
	return a
}

func (a *InputDatetimeRange) LabelClassName(value interface{}) *InputDatetimeRange {
	a.Set("labelClassName", value)
	return a
}

func (a *InputDatetimeRange) LabelRemark(value interface{}) *InputDatetimeRange {
	a.Set("labelRemark", value)
	return a
}

func (a *InputDatetimeRange) MaxDate(value interface{}) *InputDatetimeRange {
	a.Set("maxDate", value)
	return a
}

func (a *InputDatetimeRange) MinDate(value interface{}) *InputDatetimeRange {
	a.Set("minDate", value)
	return a
}

func (a *InputDatetimeRange) Name(value interface{}) *InputDatetimeRange {
	a.Set("name", value)
	return a
}

func (a *InputDatetimeRange) Placeholder(value interface{}) *InputDatetimeRange {
	a.Set("placeholder", value)
	return a
}

func (a *InputDatetimeRange) Ranges(value interface{}) *InputDatetimeRange {
	a.Set("ranges", value)
	return a
}

func (a *InputDatetimeRange) Required(value interface{}) *InputDatetimeRange {
	a.Set("required", value)
	return a
}

func (a *InputDatetimeRange) RequiredOn(value interface{}) *InputDatetimeRange {
	a.Set("requiredOn", value)
	return a
}

func (a *InputDatetimeRange) SubmitOnChange(value interface{}) *InputDatetimeRange {
	a.Set("submitOnChange", value)
	return a
}

func (a *InputDatetimeRange) Type(value interface{}) *InputDatetimeRange {
	a.Set("type", value)
	return a
}

func (a *InputDatetimeRange) Utc(value interface{}) *InputDatetimeRange {
	a.Set("utc", value)
	return a
}

func (a *InputDatetimeRange) ValidateApi(value interface{}) *InputDatetimeRange {
	a.Set("validateApi", value)
	return a
}

func (a *InputDatetimeRange) Validations(value interface{}) *InputDatetimeRange {
	a.Set("validations", value)
	return a
}

func (a *InputDatetimeRange) Value(value interface{}) *InputDatetimeRange {
	a.Set("value", value)
	return a
}

func (a *InputDatetimeRange) Visible(value interface{}) *InputDatetimeRange {
	a.Set("visible", value)
	return a
}

func (a *InputDatetimeRange) VisibleOn(value interface{}) *InputDatetimeRange {
	a.Set("visibleOn", value)
	return a
}
