package renderers

type InputYearRange struct {
	*BaseRenderer
}

func NewInputYearRange() *InputYearRange {
	a := &InputYearRange{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "input-year-range")
	return a
}

func (a *InputYearRange) Set(name string, value interface{}) *InputYearRange {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *InputYearRange) AutoFill(value interface{}) *InputYearRange {
	a.Set("autoFill", value)
	return a
}

func (a *InputYearRange) ClassName(value interface{}) *InputYearRange {
	a.Set("className", value)
	return a
}

func (a *InputYearRange) Clearable(value interface{}) *InputYearRange {
	a.Set("clearable", value)
	return a
}

func (a *InputYearRange) Description(value interface{}) *InputYearRange {
	a.Set("description", value)
	return a
}

func (a *InputYearRange) Disabled(value interface{}) *InputYearRange {
	a.Set("disabled", value)
	return a
}

func (a *InputYearRange) DisabledOn(value interface{}) *InputYearRange {
	a.Set("disabledOn", value)
	return a
}

func (a *InputYearRange) Embed(value interface{}) *InputYearRange {
	a.Set("embed", value)
	return a
}

func (a *InputYearRange) Format(value interface{}) *InputYearRange {
	a.Set("format", value)
	return a
}

func (a *InputYearRange) Inline(value interface{}) *InputYearRange {
	a.Set("inline", value)
	return a
}

func (a *InputYearRange) InputClassName(value interface{}) *InputYearRange {
	a.Set("inputClassName", value)
	return a
}

func (a *InputYearRange) InputFormat(value interface{}) *InputYearRange {
	a.Set("inputFormat", value)
	return a
}

func (a *InputYearRange) Label(value interface{}) *InputYearRange {
	a.Set("label", value)
	return a
}

func (a *InputYearRange) LabelAlign(value interface{}) *InputYearRange {
	a.Set("labelAlign", value)
	return a
}

func (a *InputYearRange) LabelClassName(value interface{}) *InputYearRange {
	a.Set("labelClassName", value)
	return a
}

func (a *InputYearRange) LabelRemark(value interface{}) *InputYearRange {
	a.Set("labelRemark", value)
	return a
}

func (a *InputYearRange) MaxDate(value interface{}) *InputYearRange {
	a.Set("maxDate", value)
	return a
}

func (a *InputYearRange) MaxDuration(value interface{}) *InputYearRange {
	a.Set("maxDuration", value)
	return a
}

func (a *InputYearRange) MinDate(value interface{}) *InputYearRange {
	a.Set("minDate", value)
	return a
}

func (a *InputYearRange) MinDuration(value interface{}) *InputYearRange {
	a.Set("minDuration", value)
	return a
}

func (a *InputYearRange) Name(value interface{}) *InputYearRange {
	a.Set("name", value)
	return a
}

func (a *InputYearRange) Placeholder(value interface{}) *InputYearRange {
	a.Set("placeholder", value)
	return a
}

func (a *InputYearRange) Required(value interface{}) *InputYearRange {
	a.Set("required", value)
	return a
}

func (a *InputYearRange) RequiredOn(value interface{}) *InputYearRange {
	a.Set("requiredOn", value)
	return a
}

func (a *InputYearRange) SubmitOnChange(value interface{}) *InputYearRange {
	a.Set("submitOnChange", value)
	return a
}

func (a *InputYearRange) Type(value interface{}) *InputYearRange {
	a.Set("type", value)
	return a
}

func (a *InputYearRange) Utc(value interface{}) *InputYearRange {
	a.Set("utc", value)
	return a
}

func (a *InputYearRange) ValidateApi(value interface{}) *InputYearRange {
	a.Set("validateApi", value)
	return a
}

func (a *InputYearRange) Validations(value interface{}) *InputYearRange {
	a.Set("validations", value)
	return a
}

func (a *InputYearRange) Value(value interface{}) *InputYearRange {
	a.Set("value", value)
	return a
}

func (a *InputYearRange) Visible(value interface{}) *InputYearRange {
	a.Set("visible", value)
	return a
}

func (a *InputYearRange) VisibleOn(value interface{}) *InputYearRange {
	a.Set("visibleOn", value)
	return a
}
