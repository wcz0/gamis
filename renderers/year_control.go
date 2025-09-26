package renderers

type YearControl struct {
	*BaseRenderer
}

func NewYearControl() *YearControl {
	a := &YearControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "input-year")
	return a
}

func (a *YearControl) Set(name string, value interface{}) *YearControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *YearControl) AutoFill(value interface{}) *YearControl {
	a.Set("autoFill", value)
	return a
}

func (a *YearControl) BorderMode(value interface{}) *YearControl {
	a.Set("borderMode", value)
	return a
}

func (a *YearControl) ClassName(value interface{}) *YearControl {
	a.Set("className", value)
	return a
}

func (a *YearControl) ClearValueOnHidden(value interface{}) *YearControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *YearControl) Clearable(value interface{}) *YearControl {
	a.Set("clearable", value)
	return a
}

func (a *YearControl) Desc(value interface{}) *YearControl {
	a.Set("desc", value)
	return a
}

func (a *YearControl) Description(value interface{}) *YearControl {
	a.Set("description", value)
	return a
}

func (a *YearControl) DescriptionClassName(value interface{}) *YearControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *YearControl) Disabled(value interface{}) *YearControl {
	a.Set("disabled", value)
	return a
}

func (a *YearControl) DisabledDate(value interface{}) *YearControl {
	a.Set("disabledDate", value)
	return a
}

func (a *YearControl) DisabledOn(value interface{}) *YearControl {
	a.Set("disabledOn", value)
	return a
}

func (a *YearControl) DisplayFormat(value interface{}) *YearControl {
	a.Set("displayFormat", value)
	return a
}

func (a *YearControl) EditorSetting(value interface{}) *YearControl {
	a.Set("editorSetting", value)
	return a
}

func (a *YearControl) Emebed(value interface{}) *YearControl {
	a.Set("emebed", value)
	return a
}

func (a *YearControl) ExtraName(value interface{}) *YearControl {
	a.Set("extraName", value)
	return a
}

func (a *YearControl) Format(value interface{}) *YearControl {
	a.Set("format", value)
	return a
}

func (a *YearControl) Hidden(value interface{}) *YearControl {
	a.Set("hidden", value)
	return a
}

func (a *YearControl) HiddenOn(value interface{}) *YearControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *YearControl) Hint(value interface{}) *YearControl {
	a.Set("hint", value)
	return a
}

func (a *YearControl) Horizontal(value interface{}) *YearControl {
	a.Set("horizontal", value)
	return a
}

func (a *YearControl) Id(value interface{}) *YearControl {
	a.Set("id", value)
	return a
}

func (a *YearControl) InitAutoFill(value interface{}) *YearControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *YearControl) Inline(value interface{}) *YearControl {
	a.Set("inline", value)
	return a
}

func (a *YearControl) InputClassName(value interface{}) *YearControl {
	a.Set("inputClassName", value)
	return a
}

func (a *YearControl) InputForbid(value interface{}) *YearControl {
	a.Set("inputForbid", value)
	return a
}

func (a *YearControl) InputFormat(value interface{}) *YearControl {
	a.Set("inputFormat", value)
	return a
}

func (a *YearControl) Label(value interface{}) *YearControl {
	a.Set("label", value)
	return a
}

func (a *YearControl) LabelAlign(value interface{}) *YearControl {
	a.Set("labelAlign", value)
	return a
}

func (a *YearControl) LabelClassName(value interface{}) *YearControl {
	a.Set("labelClassName", value)
	return a
}

func (a *YearControl) LabelRemark(value interface{}) *YearControl {
	a.Set("labelRemark", value)
	return a
}

func (a *YearControl) LabelWidth(value interface{}) *YearControl {
	a.Set("labelWidth", value)
	return a
}

func (a *YearControl) Mode(value interface{}) *YearControl {
	a.Set("mode", value)
	return a
}

func (a *YearControl) Name(value interface{}) *YearControl {
	a.Set("name", value)
	return a
}

func (a *YearControl) OnEvent(value interface{}) *YearControl {
	a.Set("onEvent", value)
	return a
}

func (a *YearControl) Placeholder(value interface{}) *YearControl {
	a.Set("placeholder", value)
	return a
}

func (a *YearControl) ReadOnly(value interface{}) *YearControl {
	a.Set("readOnly", value)
	return a
}

func (a *YearControl) ReadOnlyOn(value interface{}) *YearControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *YearControl) Remark(value interface{}) *YearControl {
	a.Set("remark", value)
	return a
}

func (a *YearControl) Required(value interface{}) *YearControl {
	a.Set("required", value)
	return a
}

func (a *YearControl) Row(value interface{}) *YearControl {
	a.Set("row", value)
	return a
}

func (a *YearControl) SaveImmediately(value interface{}) *YearControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *YearControl) Shortcuts(value interface{}) *YearControl {
	a.Set("shortcuts", value)
	return a
}

func (a *YearControl) Size(value interface{}) *YearControl {
	a.Set("size", value)
	return a
}

func (a *YearControl) Static(value interface{}) *YearControl {
	a.Set("static", value)
	return a
}

func (a *YearControl) StaticClassName(value interface{}) *YearControl {
	a.Set("staticClassName", value)
	return a
}

func (a *YearControl) StaticInputClassName(value interface{}) *YearControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *YearControl) StaticLabelClassName(value interface{}) *YearControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *YearControl) StaticOn(value interface{}) *YearControl {
	a.Set("staticOn", value)
	return a
}

func (a *YearControl) StaticPlaceholder(value interface{}) *YearControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *YearControl) StaticSchema(value interface{}) *YearControl {
	a.Set("staticSchema", value)
	return a
}

func (a *YearControl) Style(value interface{}) *YearControl {
	a.Set("style", value)
	return a
}

func (a *YearControl) SubmitOnChange(value interface{}) *YearControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *YearControl) TestIdBuilder(value interface{}) *YearControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *YearControl) Type(value interface{}) *YearControl {
	a.Set("type", value)
	return a
}

func (a *YearControl) UseMobileUI(value interface{}) *YearControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *YearControl) Utc(value interface{}) *YearControl {
	a.Set("utc", value)
	return a
}

func (a *YearControl) ValidateApi(value interface{}) *YearControl {
	a.Set("validateApi", value)
	return a
}

func (a *YearControl) ValidateOnChange(value interface{}) *YearControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *YearControl) ValidationErrors(value interface{}) *YearControl {
	a.Set("validationErrors", value)
	return a
}

func (a *YearControl) Validations(value interface{}) *YearControl {
	a.Set("validations", value)
	return a
}

func (a *YearControl) Value(value interface{}) *YearControl {
	a.Set("value", value)
	return a
}

func (a *YearControl) ValueFormat(value interface{}) *YearControl {
	a.Set("valueFormat", value)
	return a
}

func (a *YearControl) Visible(value interface{}) *YearControl {
	a.Set("visible", value)
	return a
}

func (a *YearControl) VisibleOn(value interface{}) *YearControl {
	a.Set("visibleOn", value)
	return a
}

func (a *YearControl) Width(value interface{}) *YearControl {
	a.Set("width", value)
	return a
}
