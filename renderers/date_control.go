package renderers

type DateControl struct {
	*BaseRenderer
}

func NewDateControl() *DateControl {
	a := &DateControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "input-date")
	return a
}

func (a *DateControl) Set(name string, value interface{}) *DateControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *DateControl) AutoFill(value interface{}) *DateControl {
	a.Set("autoFill", value)
	return a
}

func (a *DateControl) BorderMode(value interface{}) *DateControl {
	a.Set("borderMode", value)
	return a
}

func (a *DateControl) ClassName(value interface{}) *DateControl {
	a.Set("className", value)
	return a
}

func (a *DateControl) ClearValueOnHidden(value interface{}) *DateControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *DateControl) Clearable(value interface{}) *DateControl {
	a.Set("clearable", value)
	return a
}

func (a *DateControl) CloseOnSelect(value interface{}) *DateControl {
	a.Set("closeOnSelect", value)
	return a
}

func (a *DateControl) Desc(value interface{}) *DateControl {
	a.Set("desc", value)
	return a
}

func (a *DateControl) Description(value interface{}) *DateControl {
	a.Set("description", value)
	return a
}

func (a *DateControl) DescriptionClassName(value interface{}) *DateControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *DateControl) Disabled(value interface{}) *DateControl {
	a.Set("disabled", value)
	return a
}

func (a *DateControl) DisabledDate(value interface{}) *DateControl {
	a.Set("disabledDate", value)
	return a
}

func (a *DateControl) DisabledOn(value interface{}) *DateControl {
	a.Set("disabledOn", value)
	return a
}

func (a *DateControl) DisplayFormat(value interface{}) *DateControl {
	a.Set("displayFormat", value)
	return a
}

func (a *DateControl) EditorSetting(value interface{}) *DateControl {
	a.Set("editorSetting", value)
	return a
}

func (a *DateControl) Emebed(value interface{}) *DateControl {
	a.Set("emebed", value)
	return a
}

func (a *DateControl) ExtraName(value interface{}) *DateControl {
	a.Set("extraName", value)
	return a
}

func (a *DateControl) Format(value interface{}) *DateControl {
	a.Set("format", value)
	return a
}

func (a *DateControl) Hidden(value interface{}) *DateControl {
	a.Set("hidden", value)
	return a
}

func (a *DateControl) HiddenOn(value interface{}) *DateControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *DateControl) Hint(value interface{}) *DateControl {
	a.Set("hint", value)
	return a
}

func (a *DateControl) Horizontal(value interface{}) *DateControl {
	a.Set("horizontal", value)
	return a
}

func (a *DateControl) Id(value interface{}) *DateControl {
	a.Set("id", value)
	return a
}

func (a *DateControl) InitAutoFill(value interface{}) *DateControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *DateControl) Inline(value interface{}) *DateControl {
	a.Set("inline", value)
	return a
}

func (a *DateControl) InputClassName(value interface{}) *DateControl {
	a.Set("inputClassName", value)
	return a
}

func (a *DateControl) InputForbid(value interface{}) *DateControl {
	a.Set("inputForbid", value)
	return a
}

func (a *DateControl) InputFormat(value interface{}) *DateControl {
	a.Set("inputFormat", value)
	return a
}

func (a *DateControl) Label(value interface{}) *DateControl {
	a.Set("label", value)
	return a
}

func (a *DateControl) LabelAlign(value interface{}) *DateControl {
	a.Set("labelAlign", value)
	return a
}

func (a *DateControl) LabelClassName(value interface{}) *DateControl {
	a.Set("labelClassName", value)
	return a
}

func (a *DateControl) LabelRemark(value interface{}) *DateControl {
	a.Set("labelRemark", value)
	return a
}

func (a *DateControl) LabelWidth(value interface{}) *DateControl {
	a.Set("labelWidth", value)
	return a
}

func (a *DateControl) MaxDate(value interface{}) *DateControl {
	a.Set("maxDate", value)
	return a
}

func (a *DateControl) MinDate(value interface{}) *DateControl {
	a.Set("minDate", value)
	return a
}

func (a *DateControl) Mode(value interface{}) *DateControl {
	a.Set("mode", value)
	return a
}

func (a *DateControl) Name(value interface{}) *DateControl {
	a.Set("name", value)
	return a
}

func (a *DateControl) OnEvent(value interface{}) *DateControl {
	a.Set("onEvent", value)
	return a
}

func (a *DateControl) Placeholder(value interface{}) *DateControl {
	a.Set("placeholder", value)
	return a
}

func (a *DateControl) PopOverContainerSelector(value interface{}) *DateControl {
	a.Set("popOverContainerSelector", value)
	return a
}

func (a *DateControl) ReadOnly(value interface{}) *DateControl {
	a.Set("readOnly", value)
	return a
}

func (a *DateControl) ReadOnlyOn(value interface{}) *DateControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *DateControl) Remark(value interface{}) *DateControl {
	a.Set("remark", value)
	return a
}

func (a *DateControl) Required(value interface{}) *DateControl {
	a.Set("required", value)
	return a
}

func (a *DateControl) Row(value interface{}) *DateControl {
	a.Set("row", value)
	return a
}

func (a *DateControl) SaveImmediately(value interface{}) *DateControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *DateControl) Shortcuts(value interface{}) *DateControl {
	a.Set("shortcuts", value)
	return a
}

func (a *DateControl) Size(value interface{}) *DateControl {
	a.Set("size", value)
	return a
}

func (a *DateControl) Static(value interface{}) *DateControl {
	a.Set("static", value)
	return a
}

func (a *DateControl) StaticClassName(value interface{}) *DateControl {
	a.Set("staticClassName", value)
	return a
}

func (a *DateControl) StaticInputClassName(value interface{}) *DateControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *DateControl) StaticLabelClassName(value interface{}) *DateControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *DateControl) StaticOn(value interface{}) *DateControl {
	a.Set("staticOn", value)
	return a
}

func (a *DateControl) StaticPlaceholder(value interface{}) *DateControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *DateControl) StaticSchema(value interface{}) *DateControl {
	a.Set("staticSchema", value)
	return a
}

func (a *DateControl) Style(value interface{}) *DateControl {
	a.Set("style", value)
	return a
}

func (a *DateControl) SubmitOnChange(value interface{}) *DateControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *DateControl) TestIdBuilder(value interface{}) *DateControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *DateControl) Type(value interface{}) *DateControl {
	a.Set("type", value)
	return a
}

func (a *DateControl) UseMobileUI(value interface{}) *DateControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *DateControl) Utc(value interface{}) *DateControl {
	a.Set("utc", value)
	return a
}

func (a *DateControl) ValidateApi(value interface{}) *DateControl {
	a.Set("validateApi", value)
	return a
}

func (a *DateControl) ValidateOnChange(value interface{}) *DateControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *DateControl) ValidationErrors(value interface{}) *DateControl {
	a.Set("validationErrors", value)
	return a
}

func (a *DateControl) Validations(value interface{}) *DateControl {
	a.Set("validations", value)
	return a
}

func (a *DateControl) Value(value interface{}) *DateControl {
	a.Set("value", value)
	return a
}

func (a *DateControl) ValueFormat(value interface{}) *DateControl {
	a.Set("valueFormat", value)
	return a
}

func (a *DateControl) Visible(value interface{}) *DateControl {
	a.Set("visible", value)
	return a
}

func (a *DateControl) VisibleOn(value interface{}) *DateControl {
	a.Set("visibleOn", value)
	return a
}

func (a *DateControl) Width(value interface{}) *DateControl {
	a.Set("width", value)
	return a
}
