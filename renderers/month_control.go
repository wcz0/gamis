package renderers

type MonthControl struct {
	*BaseRenderer
}

func NewMonthControl() *MonthControl {
	a := &MonthControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "input-month")
	return a
}

func (a *MonthControl) Set(name string, value interface{}) *MonthControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *MonthControl) AutoFill(value interface{}) *MonthControl {
	a.Set("autoFill", value)
	return a
}

func (a *MonthControl) BorderMode(value interface{}) *MonthControl {
	a.Set("borderMode", value)
	return a
}

func (a *MonthControl) ClassName(value interface{}) *MonthControl {
	a.Set("className", value)
	return a
}

func (a *MonthControl) ClearValueOnHidden(value interface{}) *MonthControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *MonthControl) Clearable(value interface{}) *MonthControl {
	a.Set("clearable", value)
	return a
}

func (a *MonthControl) Desc(value interface{}) *MonthControl {
	a.Set("desc", value)
	return a
}

func (a *MonthControl) Description(value interface{}) *MonthControl {
	a.Set("description", value)
	return a
}

func (a *MonthControl) DescriptionClassName(value interface{}) *MonthControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *MonthControl) Disabled(value interface{}) *MonthControl {
	a.Set("disabled", value)
	return a
}

func (a *MonthControl) DisabledDate(value interface{}) *MonthControl {
	a.Set("disabledDate", value)
	return a
}

func (a *MonthControl) DisabledOn(value interface{}) *MonthControl {
	a.Set("disabledOn", value)
	return a
}

func (a *MonthControl) DisplayFormat(value interface{}) *MonthControl {
	a.Set("displayFormat", value)
	return a
}

func (a *MonthControl) EditorSetting(value interface{}) *MonthControl {
	a.Set("editorSetting", value)
	return a
}

func (a *MonthControl) Emebed(value interface{}) *MonthControl {
	a.Set("emebed", value)
	return a
}

func (a *MonthControl) ExtraName(value interface{}) *MonthControl {
	a.Set("extraName", value)
	return a
}

func (a *MonthControl) Format(value interface{}) *MonthControl {
	a.Set("format", value)
	return a
}

func (a *MonthControl) Hidden(value interface{}) *MonthControl {
	a.Set("hidden", value)
	return a
}

func (a *MonthControl) HiddenOn(value interface{}) *MonthControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *MonthControl) Hint(value interface{}) *MonthControl {
	a.Set("hint", value)
	return a
}

func (a *MonthControl) Horizontal(value interface{}) *MonthControl {
	a.Set("horizontal", value)
	return a
}

func (a *MonthControl) Id(value interface{}) *MonthControl {
	a.Set("id", value)
	return a
}

func (a *MonthControl) InitAutoFill(value interface{}) *MonthControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *MonthControl) Inline(value interface{}) *MonthControl {
	a.Set("inline", value)
	return a
}

func (a *MonthControl) InputClassName(value interface{}) *MonthControl {
	a.Set("inputClassName", value)
	return a
}

func (a *MonthControl) InputForbid(value interface{}) *MonthControl {
	a.Set("inputForbid", value)
	return a
}

func (a *MonthControl) InputFormat(value interface{}) *MonthControl {
	a.Set("inputFormat", value)
	return a
}

func (a *MonthControl) Label(value interface{}) *MonthControl {
	a.Set("label", value)
	return a
}

func (a *MonthControl) LabelAlign(value interface{}) *MonthControl {
	a.Set("labelAlign", value)
	return a
}

func (a *MonthControl) LabelClassName(value interface{}) *MonthControl {
	a.Set("labelClassName", value)
	return a
}

func (a *MonthControl) LabelRemark(value interface{}) *MonthControl {
	a.Set("labelRemark", value)
	return a
}

func (a *MonthControl) LabelWidth(value interface{}) *MonthControl {
	a.Set("labelWidth", value)
	return a
}

func (a *MonthControl) Mode(value interface{}) *MonthControl {
	a.Set("mode", value)
	return a
}

func (a *MonthControl) Name(value interface{}) *MonthControl {
	a.Set("name", value)
	return a
}

func (a *MonthControl) OnEvent(value interface{}) *MonthControl {
	a.Set("onEvent", value)
	return a
}

func (a *MonthControl) Placeholder(value interface{}) *MonthControl {
	a.Set("placeholder", value)
	return a
}

func (a *MonthControl) ReadOnly(value interface{}) *MonthControl {
	a.Set("readOnly", value)
	return a
}

func (a *MonthControl) ReadOnlyOn(value interface{}) *MonthControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *MonthControl) Remark(value interface{}) *MonthControl {
	a.Set("remark", value)
	return a
}

func (a *MonthControl) Required(value interface{}) *MonthControl {
	a.Set("required", value)
	return a
}

func (a *MonthControl) Row(value interface{}) *MonthControl {
	a.Set("row", value)
	return a
}

func (a *MonthControl) SaveImmediately(value interface{}) *MonthControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *MonthControl) Shortcuts(value interface{}) *MonthControl {
	a.Set("shortcuts", value)
	return a
}

func (a *MonthControl) Size(value interface{}) *MonthControl {
	a.Set("size", value)
	return a
}

func (a *MonthControl) Static(value interface{}) *MonthControl {
	a.Set("static", value)
	return a
}

func (a *MonthControl) StaticClassName(value interface{}) *MonthControl {
	a.Set("staticClassName", value)
	return a
}

func (a *MonthControl) StaticInputClassName(value interface{}) *MonthControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *MonthControl) StaticLabelClassName(value interface{}) *MonthControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *MonthControl) StaticOn(value interface{}) *MonthControl {
	a.Set("staticOn", value)
	return a
}

func (a *MonthControl) StaticPlaceholder(value interface{}) *MonthControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *MonthControl) StaticSchema(value interface{}) *MonthControl {
	a.Set("staticSchema", value)
	return a
}

func (a *MonthControl) Style(value interface{}) *MonthControl {
	a.Set("style", value)
	return a
}

func (a *MonthControl) SubmitOnChange(value interface{}) *MonthControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *MonthControl) TestIdBuilder(value interface{}) *MonthControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *MonthControl) Type(value interface{}) *MonthControl {
	a.Set("type", value)
	return a
}

func (a *MonthControl) UseMobileUI(value interface{}) *MonthControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *MonthControl) Utc(value interface{}) *MonthControl {
	a.Set("utc", value)
	return a
}

func (a *MonthControl) ValidateApi(value interface{}) *MonthControl {
	a.Set("validateApi", value)
	return a
}

func (a *MonthControl) ValidateOnChange(value interface{}) *MonthControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *MonthControl) ValidationErrors(value interface{}) *MonthControl {
	a.Set("validationErrors", value)
	return a
}

func (a *MonthControl) Validations(value interface{}) *MonthControl {
	a.Set("validations", value)
	return a
}

func (a *MonthControl) Value(value interface{}) *MonthControl {
	a.Set("value", value)
	return a
}

func (a *MonthControl) ValueFormat(value interface{}) *MonthControl {
	a.Set("valueFormat", value)
	return a
}

func (a *MonthControl) Visible(value interface{}) *MonthControl {
	a.Set("visible", value)
	return a
}

func (a *MonthControl) VisibleOn(value interface{}) *MonthControl {
	a.Set("visibleOn", value)
	return a
}

func (a *MonthControl) Width(value interface{}) *MonthControl {
	a.Set("width", value)
	return a
}
