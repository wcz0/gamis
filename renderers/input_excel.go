package renderers

type InputExcel struct {
	*BaseRenderer
}

func NewInputExcel() *InputExcel {
	a := &InputExcel{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "input-excel")
	return a
}

func (a *InputExcel) Set(name string, value interface{}) *InputExcel {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *InputExcel) AllSheets(value interface{}) *InputExcel {
	a.Set("allSheets", value)
	return a
}

func (a *InputExcel) AutoFill(value interface{}) *InputExcel {
	a.Set("autoFill", value)
	return a
}

func (a *InputExcel) ClassName(value interface{}) *InputExcel {
	a.Set("className", value)
	return a
}

func (a *InputExcel) Description(value interface{}) *InputExcel {
	a.Set("description", value)
	return a
}

func (a *InputExcel) Disabled(value interface{}) *InputExcel {
	a.Set("disabled", value)
	return a
}

func (a *InputExcel) DisabledOn(value interface{}) *InputExcel {
	a.Set("disabledOn", value)
	return a
}

func (a *InputExcel) IncludeEmpty(value interface{}) *InputExcel {
	a.Set("includeEmpty", value)
	return a
}

func (a *InputExcel) Inline(value interface{}) *InputExcel {
	a.Set("inline", value)
	return a
}

func (a *InputExcel) InputClassName(value interface{}) *InputExcel {
	a.Set("inputClassName", value)
	return a
}

func (a *InputExcel) Label(value interface{}) *InputExcel {
	a.Set("label", value)
	return a
}

func (a *InputExcel) LabelAlign(value interface{}) *InputExcel {
	a.Set("labelAlign", value)
	return a
}

func (a *InputExcel) LabelClassName(value interface{}) *InputExcel {
	a.Set("labelClassName", value)
	return a
}

func (a *InputExcel) LabelRemark(value interface{}) *InputExcel {
	a.Set("labelRemark", value)
	return a
}

func (a *InputExcel) Name(value interface{}) *InputExcel {
	a.Set("name", value)
	return a
}

func (a *InputExcel) ParseMode(value interface{}) *InputExcel {
	a.Set("parseMode", value)
	return a
}

func (a *InputExcel) Placeholder(value interface{}) *InputExcel {
	a.Set("placeholder", value)
	return a
}

func (a *InputExcel) PlainText(value interface{}) *InputExcel {
	a.Set("plainText", value)
	return a
}

func (a *InputExcel) Required(value interface{}) *InputExcel {
	a.Set("required", value)
	return a
}

func (a *InputExcel) RequiredOn(value interface{}) *InputExcel {
	a.Set("requiredOn", value)
	return a
}

func (a *InputExcel) SubmitOnChange(value interface{}) *InputExcel {
	a.Set("submitOnChange", value)
	return a
}

func (a *InputExcel) Type(value interface{}) *InputExcel {
	a.Set("type", value)
	return a
}

func (a *InputExcel) ValidateApi(value interface{}) *InputExcel {
	a.Set("validateApi", value)
	return a
}

func (a *InputExcel) Validations(value interface{}) *InputExcel {
	a.Set("validations", value)
	return a
}

func (a *InputExcel) Value(value interface{}) *InputExcel {
	a.Set("value", value)
	return a
}

func (a *InputExcel) Visible(value interface{}) *InputExcel {
	a.Set("visible", value)
	return a
}

func (a *InputExcel) VisibleOn(value interface{}) *InputExcel {
	a.Set("visibleOn", value)
	return a
}
