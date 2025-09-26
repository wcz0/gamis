package renderers

type DateTimeControl struct {
	*BaseRenderer
}

func NewDateTimeControl() *DateTimeControl {
	a := &DateTimeControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "input-datetime")
	return a
}

func (a *DateTimeControl) Set(name string, value interface{}) *DateTimeControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *DateTimeControl) AutoFill(value interface{}) *DateTimeControl {
	a.Set("autoFill", value)
	return a
}

func (a *DateTimeControl) BorderMode(value interface{}) *DateTimeControl {
	a.Set("borderMode", value)
	return a
}

func (a *DateTimeControl) ClassName(value interface{}) *DateTimeControl {
	a.Set("className", value)
	return a
}

func (a *DateTimeControl) ClearValueOnHidden(value interface{}) *DateTimeControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *DateTimeControl) Clearable(value interface{}) *DateTimeControl {
	a.Set("clearable", value)
	return a
}

func (a *DateTimeControl) Desc(value interface{}) *DateTimeControl {
	a.Set("desc", value)
	return a
}

func (a *DateTimeControl) Description(value interface{}) *DateTimeControl {
	a.Set("description", value)
	return a
}

func (a *DateTimeControl) DescriptionClassName(value interface{}) *DateTimeControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *DateTimeControl) Disabled(value interface{}) *DateTimeControl {
	a.Set("disabled", value)
	return a
}

func (a *DateTimeControl) DisabledDate(value interface{}) *DateTimeControl {
	a.Set("disabledDate", value)
	return a
}

func (a *DateTimeControl) DisabledOn(value interface{}) *DateTimeControl {
	a.Set("disabledOn", value)
	return a
}

func (a *DateTimeControl) DisplayFormat(value interface{}) *DateTimeControl {
	a.Set("displayFormat", value)
	return a
}

func (a *DateTimeControl) EditorSetting(value interface{}) *DateTimeControl {
	a.Set("editorSetting", value)
	return a
}

func (a *DateTimeControl) Emebed(value interface{}) *DateTimeControl {
	a.Set("emebed", value)
	return a
}

func (a *DateTimeControl) ExtraName(value interface{}) *DateTimeControl {
	a.Set("extraName", value)
	return a
}

func (a *DateTimeControl) Format(value interface{}) *DateTimeControl {
	a.Set("format", value)
	return a
}

func (a *DateTimeControl) Hidden(value interface{}) *DateTimeControl {
	a.Set("hidden", value)
	return a
}

func (a *DateTimeControl) HiddenOn(value interface{}) *DateTimeControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *DateTimeControl) Hint(value interface{}) *DateTimeControl {
	a.Set("hint", value)
	return a
}

func (a *DateTimeControl) Horizontal(value interface{}) *DateTimeControl {
	a.Set("horizontal", value)
	return a
}

func (a *DateTimeControl) Id(value interface{}) *DateTimeControl {
	a.Set("id", value)
	return a
}

func (a *DateTimeControl) InitAutoFill(value interface{}) *DateTimeControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *DateTimeControl) Inline(value interface{}) *DateTimeControl {
	a.Set("inline", value)
	return a
}

func (a *DateTimeControl) InputClassName(value interface{}) *DateTimeControl {
	a.Set("inputClassName", value)
	return a
}

func (a *DateTimeControl) InputForbid(value interface{}) *DateTimeControl {
	a.Set("inputForbid", value)
	return a
}

func (a *DateTimeControl) InputFormat(value interface{}) *DateTimeControl {
	a.Set("inputFormat", value)
	return a
}

func (a *DateTimeControl) IsEndDate(value interface{}) *DateTimeControl {
	a.Set("isEndDate", value)
	return a
}

func (a *DateTimeControl) Label(value interface{}) *DateTimeControl {
	a.Set("label", value)
	return a
}

func (a *DateTimeControl) LabelAlign(value interface{}) *DateTimeControl {
	a.Set("labelAlign", value)
	return a
}

func (a *DateTimeControl) LabelClassName(value interface{}) *DateTimeControl {
	a.Set("labelClassName", value)
	return a
}

func (a *DateTimeControl) LabelRemark(value interface{}) *DateTimeControl {
	a.Set("labelRemark", value)
	return a
}

func (a *DateTimeControl) LabelWidth(value interface{}) *DateTimeControl {
	a.Set("labelWidth", value)
	return a
}

func (a *DateTimeControl) MaxDate(value interface{}) *DateTimeControl {
	a.Set("maxDate", value)
	return a
}

func (a *DateTimeControl) MinDate(value interface{}) *DateTimeControl {
	a.Set("minDate", value)
	return a
}

func (a *DateTimeControl) Mode(value interface{}) *DateTimeControl {
	a.Set("mode", value)
	return a
}

func (a *DateTimeControl) Name(value interface{}) *DateTimeControl {
	a.Set("name", value)
	return a
}

func (a *DateTimeControl) OnEvent(value interface{}) *DateTimeControl {
	a.Set("onEvent", value)
	return a
}

func (a *DateTimeControl) Placeholder(value interface{}) *DateTimeControl {
	a.Set("placeholder", value)
	return a
}

func (a *DateTimeControl) ReadOnly(value interface{}) *DateTimeControl {
	a.Set("readOnly", value)
	return a
}

func (a *DateTimeControl) ReadOnlyOn(value interface{}) *DateTimeControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *DateTimeControl) Remark(value interface{}) *DateTimeControl {
	a.Set("remark", value)
	return a
}

func (a *DateTimeControl) Required(value interface{}) *DateTimeControl {
	a.Set("required", value)
	return a
}

func (a *DateTimeControl) Row(value interface{}) *DateTimeControl {
	a.Set("row", value)
	return a
}

func (a *DateTimeControl) SaveImmediately(value interface{}) *DateTimeControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *DateTimeControl) Shortcuts(value interface{}) *DateTimeControl {
	a.Set("shortcuts", value)
	return a
}

func (a *DateTimeControl) Size(value interface{}) *DateTimeControl {
	a.Set("size", value)
	return a
}

func (a *DateTimeControl) Static(value interface{}) *DateTimeControl {
	a.Set("static", value)
	return a
}

func (a *DateTimeControl) StaticClassName(value interface{}) *DateTimeControl {
	a.Set("staticClassName", value)
	return a
}

func (a *DateTimeControl) StaticInputClassName(value interface{}) *DateTimeControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *DateTimeControl) StaticLabelClassName(value interface{}) *DateTimeControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *DateTimeControl) StaticOn(value interface{}) *DateTimeControl {
	a.Set("staticOn", value)
	return a
}

func (a *DateTimeControl) StaticPlaceholder(value interface{}) *DateTimeControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *DateTimeControl) StaticSchema(value interface{}) *DateTimeControl {
	a.Set("staticSchema", value)
	return a
}

func (a *DateTimeControl) Style(value interface{}) *DateTimeControl {
	a.Set("style", value)
	return a
}

func (a *DateTimeControl) SubmitOnChange(value interface{}) *DateTimeControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *DateTimeControl) TestIdBuilder(value interface{}) *DateTimeControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *DateTimeControl) TimeConstraints(value interface{}) *DateTimeControl {
	a.Set("timeConstraints", value)
	return a
}

func (a *DateTimeControl) TimeFormat(value interface{}) *DateTimeControl {
	a.Set("timeFormat", value)
	return a
}

func (a *DateTimeControl) Type(value interface{}) *DateTimeControl {
	a.Set("type", value)
	return a
}

func (a *DateTimeControl) UseMobileUI(value interface{}) *DateTimeControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *DateTimeControl) Utc(value interface{}) *DateTimeControl {
	a.Set("utc", value)
	return a
}

func (a *DateTimeControl) ValidateApi(value interface{}) *DateTimeControl {
	a.Set("validateApi", value)
	return a
}

func (a *DateTimeControl) ValidateOnChange(value interface{}) *DateTimeControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *DateTimeControl) ValidationErrors(value interface{}) *DateTimeControl {
	a.Set("validationErrors", value)
	return a
}

func (a *DateTimeControl) Validations(value interface{}) *DateTimeControl {
	a.Set("validations", value)
	return a
}

func (a *DateTimeControl) Value(value interface{}) *DateTimeControl {
	a.Set("value", value)
	return a
}

func (a *DateTimeControl) ValueFormat(value interface{}) *DateTimeControl {
	a.Set("valueFormat", value)
	return a
}

func (a *DateTimeControl) Visible(value interface{}) *DateTimeControl {
	a.Set("visible", value)
	return a
}

func (a *DateTimeControl) VisibleOn(value interface{}) *DateTimeControl {
	a.Set("visibleOn", value)
	return a
}

func (a *DateTimeControl) Width(value interface{}) *DateTimeControl {
	a.Set("width", value)
	return a
}
