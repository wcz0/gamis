package renderers

type InputCityControl struct {
	*BaseRenderer
}

func NewInputCityControl() *InputCityControl {
	a := &InputCityControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "input-city")
	return a
}

func (a *InputCityControl) Set(name string, value interface{}) *InputCityControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *InputCityControl) AllowCity(value interface{}) *InputCityControl {
	a.Set("allowCity", value)
	return a
}

func (a *InputCityControl) AllowDistrict(value interface{}) *InputCityControl {
	a.Set("allowDistrict", value)
	return a
}

func (a *InputCityControl) AllowStreet(value interface{}) *InputCityControl {
	a.Set("allowStreet", value)
	return a
}

func (a *InputCityControl) AutoFill(value interface{}) *InputCityControl {
	a.Set("autoFill", value)
	return a
}

func (a *InputCityControl) ClassName(value interface{}) *InputCityControl {
	a.Set("className", value)
	return a
}

func (a *InputCityControl) ClearValueOnHidden(value interface{}) *InputCityControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *InputCityControl) Delimiter(value interface{}) *InputCityControl {
	a.Set("delimiter", value)
	return a
}

func (a *InputCityControl) Desc(value interface{}) *InputCityControl {
	a.Set("desc", value)
	return a
}

func (a *InputCityControl) Description(value interface{}) *InputCityControl {
	a.Set("description", value)
	return a
}

func (a *InputCityControl) DescriptionClassName(value interface{}) *InputCityControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *InputCityControl) Disabled(value interface{}) *InputCityControl {
	a.Set("disabled", value)
	return a
}

func (a *InputCityControl) DisabledOn(value interface{}) *InputCityControl {
	a.Set("disabledOn", value)
	return a
}

func (a *InputCityControl) EditorSetting(value interface{}) *InputCityControl {
	a.Set("editorSetting", value)
	return a
}

func (a *InputCityControl) ExtraName(value interface{}) *InputCityControl {
	a.Set("extraName", value)
	return a
}

func (a *InputCityControl) ExtractValue(value interface{}) *InputCityControl {
	a.Set("extractValue", value)
	return a
}

func (a *InputCityControl) Hidden(value interface{}) *InputCityControl {
	a.Set("hidden", value)
	return a
}

func (a *InputCityControl) HiddenOn(value interface{}) *InputCityControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *InputCityControl) Hint(value interface{}) *InputCityControl {
	a.Set("hint", value)
	return a
}

func (a *InputCityControl) Horizontal(value interface{}) *InputCityControl {
	a.Set("horizontal", value)
	return a
}

func (a *InputCityControl) Id(value interface{}) *InputCityControl {
	a.Set("id", value)
	return a
}

func (a *InputCityControl) InitAutoFill(value interface{}) *InputCityControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *InputCityControl) Inline(value interface{}) *InputCityControl {
	a.Set("inline", value)
	return a
}

func (a *InputCityControl) InputClassName(value interface{}) *InputCityControl {
	a.Set("inputClassName", value)
	return a
}

func (a *InputCityControl) ItemClassName(value interface{}) *InputCityControl {
	a.Set("itemClassName", value)
	return a
}

func (a *InputCityControl) JoinValues(value interface{}) *InputCityControl {
	a.Set("joinValues", value)
	return a
}

func (a *InputCityControl) Label(value interface{}) *InputCityControl {
	a.Set("label", value)
	return a
}

func (a *InputCityControl) LabelAlign(value interface{}) *InputCityControl {
	a.Set("labelAlign", value)
	return a
}

func (a *InputCityControl) LabelClassName(value interface{}) *InputCityControl {
	a.Set("labelClassName", value)
	return a
}

func (a *InputCityControl) LabelRemark(value interface{}) *InputCityControl {
	a.Set("labelRemark", value)
	return a
}

func (a *InputCityControl) LabelWidth(value interface{}) *InputCityControl {
	a.Set("labelWidth", value)
	return a
}

func (a *InputCityControl) LoadingConfig(value interface{}) *InputCityControl {
	a.Set("loadingConfig", value)
	return a
}

func (a *InputCityControl) Mode(value interface{}) *InputCityControl {
	a.Set("mode", value)
	return a
}

func (a *InputCityControl) Name(value interface{}) *InputCityControl {
	a.Set("name", value)
	return a
}

func (a *InputCityControl) OnEvent(value interface{}) *InputCityControl {
	a.Set("onEvent", value)
	return a
}

func (a *InputCityControl) Placeholder(value interface{}) *InputCityControl {
	a.Set("placeholder", value)
	return a
}

func (a *InputCityControl) ReadOnly(value interface{}) *InputCityControl {
	a.Set("readOnly", value)
	return a
}

func (a *InputCityControl) ReadOnlyOn(value interface{}) *InputCityControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *InputCityControl) Remark(value interface{}) *InputCityControl {
	a.Set("remark", value)
	return a
}

func (a *InputCityControl) Required(value interface{}) *InputCityControl {
	a.Set("required", value)
	return a
}

func (a *InputCityControl) Row(value interface{}) *InputCityControl {
	a.Set("row", value)
	return a
}

func (a *InputCityControl) SaveImmediately(value interface{}) *InputCityControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *InputCityControl) Searchable(value interface{}) *InputCityControl {
	a.Set("searchable", value)
	return a
}

func (a *InputCityControl) Size(value interface{}) *InputCityControl {
	a.Set("size", value)
	return a
}

func (a *InputCityControl) Static(value interface{}) *InputCityControl {
	a.Set("static", value)
	return a
}

func (a *InputCityControl) StaticClassName(value interface{}) *InputCityControl {
	a.Set("staticClassName", value)
	return a
}

func (a *InputCityControl) StaticInputClassName(value interface{}) *InputCityControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *InputCityControl) StaticLabelClassName(value interface{}) *InputCityControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *InputCityControl) StaticOn(value interface{}) *InputCityControl {
	a.Set("staticOn", value)
	return a
}

func (a *InputCityControl) StaticPlaceholder(value interface{}) *InputCityControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *InputCityControl) StaticSchema(value interface{}) *InputCityControl {
	a.Set("staticSchema", value)
	return a
}

func (a *InputCityControl) Style(value interface{}) *InputCityControl {
	a.Set("style", value)
	return a
}

func (a *InputCityControl) SubmitOnChange(value interface{}) *InputCityControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *InputCityControl) TestIdBuilder(value interface{}) *InputCityControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *InputCityControl) Type(value interface{}) *InputCityControl {
	a.Set("type", value)
	return a
}

func (a *InputCityControl) UseMobileUI(value interface{}) *InputCityControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *InputCityControl) ValidateApi(value interface{}) *InputCityControl {
	a.Set("validateApi", value)
	return a
}

func (a *InputCityControl) ValidateOnChange(value interface{}) *InputCityControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *InputCityControl) ValidationErrors(value interface{}) *InputCityControl {
	a.Set("validationErrors", value)
	return a
}

func (a *InputCityControl) Validations(value interface{}) *InputCityControl {
	a.Set("validations", value)
	return a
}

func (a *InputCityControl) Value(value interface{}) *InputCityControl {
	a.Set("value", value)
	return a
}

func (a *InputCityControl) Visible(value interface{}) *InputCityControl {
	a.Set("visible", value)
	return a
}

func (a *InputCityControl) VisibleOn(value interface{}) *InputCityControl {
	a.Set("visibleOn", value)
	return a
}

func (a *InputCityControl) Width(value interface{}) *InputCityControl {
	a.Set("width", value)
	return a
}
