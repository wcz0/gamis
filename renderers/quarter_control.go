package renderers

type QuarterControl struct {
	*BaseRenderer
}

func NewQuarterControl() *QuarterControl {
	a := &QuarterControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "input-quarter")
	return a
}

func (a *QuarterControl) Set(name string, value interface{}) *QuarterControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *QuarterControl) AutoFill(value interface{}) *QuarterControl {
	a.Set("autoFill", value)
	return a
}

func (a *QuarterControl) BorderMode(value interface{}) *QuarterControl {
	a.Set("borderMode", value)
	return a
}

func (a *QuarterControl) ClassName(value interface{}) *QuarterControl {
	a.Set("className", value)
	return a
}

func (a *QuarterControl) ClearValueOnHidden(value interface{}) *QuarterControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *QuarterControl) Clearable(value interface{}) *QuarterControl {
	a.Set("clearable", value)
	return a
}

func (a *QuarterControl) Desc(value interface{}) *QuarterControl {
	a.Set("desc", value)
	return a
}

func (a *QuarterControl) Description(value interface{}) *QuarterControl {
	a.Set("description", value)
	return a
}

func (a *QuarterControl) DescriptionClassName(value interface{}) *QuarterControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *QuarterControl) Disabled(value interface{}) *QuarterControl {
	a.Set("disabled", value)
	return a
}

func (a *QuarterControl) DisabledDate(value interface{}) *QuarterControl {
	a.Set("disabledDate", value)
	return a
}

func (a *QuarterControl) DisabledOn(value interface{}) *QuarterControl {
	a.Set("disabledOn", value)
	return a
}

func (a *QuarterControl) DisplayFormat(value interface{}) *QuarterControl {
	a.Set("displayFormat", value)
	return a
}

func (a *QuarterControl) EditorSetting(value interface{}) *QuarterControl {
	a.Set("editorSetting", value)
	return a
}

func (a *QuarterControl) Emebed(value interface{}) *QuarterControl {
	a.Set("emebed", value)
	return a
}

func (a *QuarterControl) ExtraName(value interface{}) *QuarterControl {
	a.Set("extraName", value)
	return a
}

func (a *QuarterControl) Format(value interface{}) *QuarterControl {
	a.Set("format", value)
	return a
}

func (a *QuarterControl) Hidden(value interface{}) *QuarterControl {
	a.Set("hidden", value)
	return a
}

func (a *QuarterControl) HiddenOn(value interface{}) *QuarterControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *QuarterControl) Hint(value interface{}) *QuarterControl {
	a.Set("hint", value)
	return a
}

func (a *QuarterControl) Horizontal(value interface{}) *QuarterControl {
	a.Set("horizontal", value)
	return a
}

func (a *QuarterControl) Id(value interface{}) *QuarterControl {
	a.Set("id", value)
	return a
}

func (a *QuarterControl) InitAutoFill(value interface{}) *QuarterControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *QuarterControl) Inline(value interface{}) *QuarterControl {
	a.Set("inline", value)
	return a
}

func (a *QuarterControl) InputClassName(value interface{}) *QuarterControl {
	a.Set("inputClassName", value)
	return a
}

func (a *QuarterControl) InputForbid(value interface{}) *QuarterControl {
	a.Set("inputForbid", value)
	return a
}

func (a *QuarterControl) InputFormat(value interface{}) *QuarterControl {
	a.Set("inputFormat", value)
	return a
}

func (a *QuarterControl) Label(value interface{}) *QuarterControl {
	a.Set("label", value)
	return a
}

func (a *QuarterControl) LabelAlign(value interface{}) *QuarterControl {
	a.Set("labelAlign", value)
	return a
}

func (a *QuarterControl) LabelClassName(value interface{}) *QuarterControl {
	a.Set("labelClassName", value)
	return a
}

func (a *QuarterControl) LabelRemark(value interface{}) *QuarterControl {
	a.Set("labelRemark", value)
	return a
}

func (a *QuarterControl) LabelWidth(value interface{}) *QuarterControl {
	a.Set("labelWidth", value)
	return a
}

func (a *QuarterControl) Mode(value interface{}) *QuarterControl {
	a.Set("mode", value)
	return a
}

func (a *QuarterControl) Name(value interface{}) *QuarterControl {
	a.Set("name", value)
	return a
}

func (a *QuarterControl) OnEvent(value interface{}) *QuarterControl {
	a.Set("onEvent", value)
	return a
}

func (a *QuarterControl) Placeholder(value interface{}) *QuarterControl {
	a.Set("placeholder", value)
	return a
}

func (a *QuarterControl) ReadOnly(value interface{}) *QuarterControl {
	a.Set("readOnly", value)
	return a
}

func (a *QuarterControl) ReadOnlyOn(value interface{}) *QuarterControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *QuarterControl) Remark(value interface{}) *QuarterControl {
	a.Set("remark", value)
	return a
}

func (a *QuarterControl) Required(value interface{}) *QuarterControl {
	a.Set("required", value)
	return a
}

func (a *QuarterControl) Row(value interface{}) *QuarterControl {
	a.Set("row", value)
	return a
}

func (a *QuarterControl) SaveImmediately(value interface{}) *QuarterControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *QuarterControl) Shortcuts(value interface{}) *QuarterControl {
	a.Set("shortcuts", value)
	return a
}

func (a *QuarterControl) Size(value interface{}) *QuarterControl {
	a.Set("size", value)
	return a
}

func (a *QuarterControl) Static(value interface{}) *QuarterControl {
	a.Set("static", value)
	return a
}

func (a *QuarterControl) StaticClassName(value interface{}) *QuarterControl {
	a.Set("staticClassName", value)
	return a
}

func (a *QuarterControl) StaticInputClassName(value interface{}) *QuarterControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *QuarterControl) StaticLabelClassName(value interface{}) *QuarterControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *QuarterControl) StaticOn(value interface{}) *QuarterControl {
	a.Set("staticOn", value)
	return a
}

func (a *QuarterControl) StaticPlaceholder(value interface{}) *QuarterControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *QuarterControl) StaticSchema(value interface{}) *QuarterControl {
	a.Set("staticSchema", value)
	return a
}

func (a *QuarterControl) Style(value interface{}) *QuarterControl {
	a.Set("style", value)
	return a
}

func (a *QuarterControl) SubmitOnChange(value interface{}) *QuarterControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *QuarterControl) TestIdBuilder(value interface{}) *QuarterControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *QuarterControl) Type(value interface{}) *QuarterControl {
	a.Set("type", value)
	return a
}

func (a *QuarterControl) UseMobileUI(value interface{}) *QuarterControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *QuarterControl) Utc(value interface{}) *QuarterControl {
	a.Set("utc", value)
	return a
}

func (a *QuarterControl) ValidateApi(value interface{}) *QuarterControl {
	a.Set("validateApi", value)
	return a
}

func (a *QuarterControl) ValidateOnChange(value interface{}) *QuarterControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *QuarterControl) ValidationErrors(value interface{}) *QuarterControl {
	a.Set("validationErrors", value)
	return a
}

func (a *QuarterControl) Validations(value interface{}) *QuarterControl {
	a.Set("validations", value)
	return a
}

func (a *QuarterControl) Value(value interface{}) *QuarterControl {
	a.Set("value", value)
	return a
}

func (a *QuarterControl) ValueFormat(value interface{}) *QuarterControl {
	a.Set("valueFormat", value)
	return a
}

func (a *QuarterControl) Visible(value interface{}) *QuarterControl {
	a.Set("visible", value)
	return a
}

func (a *QuarterControl) VisibleOn(value interface{}) *QuarterControl {
	a.Set("visibleOn", value)
	return a
}

func (a *QuarterControl) Width(value interface{}) *QuarterControl {
	a.Set("width", value)
	return a
}
