package renderers

type Options struct {
	*BaseRenderer
}

func NewOptions() *Options {
	a := &Options{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *Options) Set(name string, value interface{}) *Options {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Options) AutoFill(value interface{}) *Options {
	a.Set("autoFill", value)
	return a
}

func (a *Options) ClassName(value interface{}) *Options {
	a.Set("className", value)
	return a
}

func (a *Options) Description(value interface{}) *Options {
	a.Set("description", value)
	return a
}

func (a *Options) Disabled(value interface{}) *Options {
	a.Set("disabled", value)
	return a
}

func (a *Options) DisabledOn(value interface{}) *Options {
	a.Set("disabledOn", value)
	return a
}

func (a *Options) ExtractValue(value interface{}) *Options {
	a.Set("extractValue", value)
	return a
}

func (a *Options) Inline(value interface{}) *Options {
	a.Set("inline", value)
	return a
}

func (a *Options) InputClassName(value interface{}) *Options {
	a.Set("inputClassName", value)
	return a
}

func (a *Options) ItemHeight(value interface{}) *Options {
	a.Set("itemHeight", value)
	return a
}

func (a *Options) JoinValues(value interface{}) *Options {
	a.Set("joinValues", value)
	return a
}

func (a *Options) Label(value interface{}) *Options {
	a.Set("label", value)
	return a
}

func (a *Options) LabelAlign(value interface{}) *Options {
	a.Set("labelAlign", value)
	return a
}

func (a *Options) LabelClassName(value interface{}) *Options {
	a.Set("labelClassName", value)
	return a
}

func (a *Options) LabelField(value interface{}) *Options {
	a.Set("labelField", value)
	return a
}

func (a *Options) LabelRemark(value interface{}) *Options {
	a.Set("labelRemark", value)
	return a
}

func (a *Options) Multiple(value interface{}) *Options {
	a.Set("multiple", value)
	return a
}

func (a *Options) Name(value interface{}) *Options {
	a.Set("name", value)
	return a
}

func (a *Options) Options(value interface{}) *Options {
	a.Set("options", value)
	return a
}

func (a *Options) Required(value interface{}) *Options {
	a.Set("required", value)
	return a
}

func (a *Options) RequiredOn(value interface{}) *Options {
	a.Set("requiredOn", value)
	return a
}

func (a *Options) Source(value interface{}) *Options {
	a.Set("source", value)
	return a
}

func (a *Options) SubmitOnChange(value interface{}) *Options {
	a.Set("submitOnChange", value)
	return a
}

func (a *Options) ValidateApi(value interface{}) *Options {
	a.Set("validateApi", value)
	return a
}

func (a *Options) Validations(value interface{}) *Options {
	a.Set("validations", value)
	return a
}

func (a *Options) Value(value interface{}) *Options {
	a.Set("value", value)
	return a
}

func (a *Options) ValueField(value interface{}) *Options {
	a.Set("valueField", value)
	return a
}

func (a *Options) ValuesNoWrap(value interface{}) *Options {
	a.Set("valuesNoWrap", value)
	return a
}

func (a *Options) VirtualThreshold(value interface{}) *Options {
	a.Set("virtualThreshold", value)
	return a
}

func (a *Options) Visible(value interface{}) *Options {
	a.Set("visible", value)
	return a
}

func (a *Options) VisibleOn(value interface{}) *Options {
	a.Set("visibleOn", value)
	return a
}
