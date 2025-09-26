package renderers

type TimeControl struct {
	*BaseRenderer
}

func NewTimeControl() *TimeControl {
	a := &TimeControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "input-time")
	return a
}

func (a *TimeControl) Set(name string, value interface{}) *TimeControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *TimeControl) AutoFill(value interface{}) *TimeControl {
	a.Set("autoFill", value)
	return a
}

func (a *TimeControl) BorderMode(value interface{}) *TimeControl {
	a.Set("borderMode", value)
	return a
}

func (a *TimeControl) ClassName(value interface{}) *TimeControl {
	a.Set("className", value)
	return a
}

func (a *TimeControl) ClearValueOnHidden(value interface{}) *TimeControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *TimeControl) Clearable(value interface{}) *TimeControl {
	a.Set("clearable", value)
	return a
}

func (a *TimeControl) Desc(value interface{}) *TimeControl {
	a.Set("desc", value)
	return a
}

func (a *TimeControl) Description(value interface{}) *TimeControl {
	a.Set("description", value)
	return a
}

func (a *TimeControl) DescriptionClassName(value interface{}) *TimeControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *TimeControl) Disabled(value interface{}) *TimeControl {
	a.Set("disabled", value)
	return a
}

func (a *TimeControl) DisabledDate(value interface{}) *TimeControl {
	a.Set("disabledDate", value)
	return a
}

func (a *TimeControl) DisabledOn(value interface{}) *TimeControl {
	a.Set("disabledOn", value)
	return a
}

func (a *TimeControl) DisplayFormat(value interface{}) *TimeControl {
	a.Set("displayFormat", value)
	return a
}

func (a *TimeControl) EditorSetting(value interface{}) *TimeControl {
	a.Set("editorSetting", value)
	return a
}

func (a *TimeControl) Emebed(value interface{}) *TimeControl {
	a.Set("emebed", value)
	return a
}

func (a *TimeControl) ExtraName(value interface{}) *TimeControl {
	a.Set("extraName", value)
	return a
}

func (a *TimeControl) Format(value interface{}) *TimeControl {
	a.Set("format", value)
	return a
}

func (a *TimeControl) Hidden(value interface{}) *TimeControl {
	a.Set("hidden", value)
	return a
}

func (a *TimeControl) HiddenOn(value interface{}) *TimeControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *TimeControl) Hint(value interface{}) *TimeControl {
	a.Set("hint", value)
	return a
}

func (a *TimeControl) Horizontal(value interface{}) *TimeControl {
	a.Set("horizontal", value)
	return a
}

func (a *TimeControl) Id(value interface{}) *TimeControl {
	a.Set("id", value)
	return a
}

func (a *TimeControl) InitAutoFill(value interface{}) *TimeControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *TimeControl) Inline(value interface{}) *TimeControl {
	a.Set("inline", value)
	return a
}

func (a *TimeControl) InputClassName(value interface{}) *TimeControl {
	a.Set("inputClassName", value)
	return a
}

func (a *TimeControl) InputForbid(value interface{}) *TimeControl {
	a.Set("inputForbid", value)
	return a
}

func (a *TimeControl) InputFormat(value interface{}) *TimeControl {
	a.Set("inputFormat", value)
	return a
}

func (a *TimeControl) Label(value interface{}) *TimeControl {
	a.Set("label", value)
	return a
}

func (a *TimeControl) LabelAlign(value interface{}) *TimeControl {
	a.Set("labelAlign", value)
	return a
}

func (a *TimeControl) LabelClassName(value interface{}) *TimeControl {
	a.Set("labelClassName", value)
	return a
}

func (a *TimeControl) LabelRemark(value interface{}) *TimeControl {
	a.Set("labelRemark", value)
	return a
}

func (a *TimeControl) LabelWidth(value interface{}) *TimeControl {
	a.Set("labelWidth", value)
	return a
}

func (a *TimeControl) Mode(value interface{}) *TimeControl {
	a.Set("mode", value)
	return a
}

func (a *TimeControl) Name(value interface{}) *TimeControl {
	a.Set("name", value)
	return a
}

func (a *TimeControl) OnEvent(value interface{}) *TimeControl {
	a.Set("onEvent", value)
	return a
}

func (a *TimeControl) Placeholder(value interface{}) *TimeControl {
	a.Set("placeholder", value)
	return a
}

func (a *TimeControl) ReadOnly(value interface{}) *TimeControl {
	a.Set("readOnly", value)
	return a
}

func (a *TimeControl) ReadOnlyOn(value interface{}) *TimeControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *TimeControl) Remark(value interface{}) *TimeControl {
	a.Set("remark", value)
	return a
}

func (a *TimeControl) Required(value interface{}) *TimeControl {
	a.Set("required", value)
	return a
}

func (a *TimeControl) Row(value interface{}) *TimeControl {
	a.Set("row", value)
	return a
}

func (a *TimeControl) SaveImmediately(value interface{}) *TimeControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *TimeControl) Shortcuts(value interface{}) *TimeControl {
	a.Set("shortcuts", value)
	return a
}

func (a *TimeControl) Size(value interface{}) *TimeControl {
	a.Set("size", value)
	return a
}

func (a *TimeControl) Static(value interface{}) *TimeControl {
	a.Set("static", value)
	return a
}

func (a *TimeControl) StaticClassName(value interface{}) *TimeControl {
	a.Set("staticClassName", value)
	return a
}

func (a *TimeControl) StaticInputClassName(value interface{}) *TimeControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *TimeControl) StaticLabelClassName(value interface{}) *TimeControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *TimeControl) StaticOn(value interface{}) *TimeControl {
	a.Set("staticOn", value)
	return a
}

func (a *TimeControl) StaticPlaceholder(value interface{}) *TimeControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *TimeControl) StaticSchema(value interface{}) *TimeControl {
	a.Set("staticSchema", value)
	return a
}

func (a *TimeControl) Style(value interface{}) *TimeControl {
	a.Set("style", value)
	return a
}

func (a *TimeControl) SubmitOnChange(value interface{}) *TimeControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *TimeControl) TestIdBuilder(value interface{}) *TimeControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *TimeControl) TimeConstraints(value interface{}) *TimeControl {
	a.Set("timeConstraints", value)
	return a
}

func (a *TimeControl) TimeFormat(value interface{}) *TimeControl {
	a.Set("timeFormat", value)
	return a
}

func (a *TimeControl) Type(value interface{}) *TimeControl {
	a.Set("type", value)
	return a
}

func (a *TimeControl) UseMobileUI(value interface{}) *TimeControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *TimeControl) Utc(value interface{}) *TimeControl {
	a.Set("utc", value)
	return a
}

func (a *TimeControl) ValidateApi(value interface{}) *TimeControl {
	a.Set("validateApi", value)
	return a
}

func (a *TimeControl) ValidateOnChange(value interface{}) *TimeControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *TimeControl) ValidationErrors(value interface{}) *TimeControl {
	a.Set("validationErrors", value)
	return a
}

func (a *TimeControl) Validations(value interface{}) *TimeControl {
	a.Set("validations", value)
	return a
}

func (a *TimeControl) Value(value interface{}) *TimeControl {
	a.Set("value", value)
	return a
}

func (a *TimeControl) ValueFormat(value interface{}) *TimeControl {
	a.Set("valueFormat", value)
	return a
}

func (a *TimeControl) Visible(value interface{}) *TimeControl {
	a.Set("visible", value)
	return a
}

func (a *TimeControl) VisibleOn(value interface{}) *TimeControl {
	a.Set("visibleOn", value)
	return a
}

func (a *TimeControl) Width(value interface{}) *TimeControl {
	a.Set("width", value)
	return a
}
