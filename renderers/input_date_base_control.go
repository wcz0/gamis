package renderers

type InputDateBaseControl struct {
	*BaseRenderer
}

func NewInputDateBaseControl() *InputDateBaseControl {
	a := &InputDateBaseControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "input-date")
	return a
}

func (a *InputDateBaseControl) Set(name string, value interface{}) *InputDateBaseControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *InputDateBaseControl) AutoFill(value interface{}) *InputDateBaseControl {
	a.Set("autoFill", value)
	return a
}

func (a *InputDateBaseControl) BorderMode(value interface{}) *InputDateBaseControl {
	a.Set("borderMode", value)
	return a
}

func (a *InputDateBaseControl) ClassName(value interface{}) *InputDateBaseControl {
	a.Set("className", value)
	return a
}

func (a *InputDateBaseControl) ClearValueOnHidden(value interface{}) *InputDateBaseControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *InputDateBaseControl) Clearable(value interface{}) *InputDateBaseControl {
	a.Set("clearable", value)
	return a
}

func (a *InputDateBaseControl) Desc(value interface{}) *InputDateBaseControl {
	a.Set("desc", value)
	return a
}

func (a *InputDateBaseControl) Description(value interface{}) *InputDateBaseControl {
	a.Set("description", value)
	return a
}

func (a *InputDateBaseControl) DescriptionClassName(value interface{}) *InputDateBaseControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *InputDateBaseControl) Disabled(value interface{}) *InputDateBaseControl {
	a.Set("disabled", value)
	return a
}

func (a *InputDateBaseControl) DisabledDate(value interface{}) *InputDateBaseControl {
	a.Set("disabledDate", value)
	return a
}

func (a *InputDateBaseControl) DisabledOn(value interface{}) *InputDateBaseControl {
	a.Set("disabledOn", value)
	return a
}

func (a *InputDateBaseControl) DisplayFormat(value interface{}) *InputDateBaseControl {
	a.Set("displayFormat", value)
	return a
}

func (a *InputDateBaseControl) EditorSetting(value interface{}) *InputDateBaseControl {
	a.Set("editorSetting", value)
	return a
}

func (a *InputDateBaseControl) Emebed(value interface{}) *InputDateBaseControl {
	a.Set("emebed", value)
	return a
}

func (a *InputDateBaseControl) ExtraName(value interface{}) *InputDateBaseControl {
	a.Set("extraName", value)
	return a
}

func (a *InputDateBaseControl) Format(value interface{}) *InputDateBaseControl {
	a.Set("format", value)
	return a
}

func (a *InputDateBaseControl) Hidden(value interface{}) *InputDateBaseControl {
	a.Set("hidden", value)
	return a
}

func (a *InputDateBaseControl) HiddenOn(value interface{}) *InputDateBaseControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *InputDateBaseControl) Hint(value interface{}) *InputDateBaseControl {
	a.Set("hint", value)
	return a
}

func (a *InputDateBaseControl) Horizontal(value interface{}) *InputDateBaseControl {
	a.Set("horizontal", value)
	return a
}

func (a *InputDateBaseControl) Id(value interface{}) *InputDateBaseControl {
	a.Set("id", value)
	return a
}

func (a *InputDateBaseControl) InitAutoFill(value interface{}) *InputDateBaseControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *InputDateBaseControl) Inline(value interface{}) *InputDateBaseControl {
	a.Set("inline", value)
	return a
}

func (a *InputDateBaseControl) InputClassName(value interface{}) *InputDateBaseControl {
	a.Set("inputClassName", value)
	return a
}

func (a *InputDateBaseControl) InputForbid(value interface{}) *InputDateBaseControl {
	a.Set("inputForbid", value)
	return a
}

func (a *InputDateBaseControl) InputFormat(value interface{}) *InputDateBaseControl {
	a.Set("inputFormat", value)
	return a
}

func (a *InputDateBaseControl) Label(value interface{}) *InputDateBaseControl {
	a.Set("label", value)
	return a
}

func (a *InputDateBaseControl) LabelAlign(value interface{}) *InputDateBaseControl {
	a.Set("labelAlign", value)
	return a
}

func (a *InputDateBaseControl) LabelClassName(value interface{}) *InputDateBaseControl {
	a.Set("labelClassName", value)
	return a
}

func (a *InputDateBaseControl) LabelOverflow(value interface{}) *InputDateBaseControl {
	a.Set("labelOverflow", value)
	return a
}

func (a *InputDateBaseControl) LabelRemark(value interface{}) *InputDateBaseControl {
	a.Set("labelRemark", value)
	return a
}

func (a *InputDateBaseControl) LabelWidth(value interface{}) *InputDateBaseControl {
	a.Set("labelWidth", value)
	return a
}

func (a *InputDateBaseControl) Mode(value interface{}) *InputDateBaseControl {
	a.Set("mode", value)
	return a
}

func (a *InputDateBaseControl) Name(value interface{}) *InputDateBaseControl {
	a.Set("name", value)
	return a
}

func (a *InputDateBaseControl) OnEvent(value interface{}) *InputDateBaseControl {
	a.Set("onEvent", value)
	return a
}

func (a *InputDateBaseControl) Placeholder(value interface{}) *InputDateBaseControl {
	a.Set("placeholder", value)
	return a
}

func (a *InputDateBaseControl) ReadOnly(value interface{}) *InputDateBaseControl {
	a.Set("readOnly", value)
	return a
}

func (a *InputDateBaseControl) ReadOnlyOn(value interface{}) *InputDateBaseControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *InputDateBaseControl) Remark(value interface{}) *InputDateBaseControl {
	a.Set("remark", value)
	return a
}

func (a *InputDateBaseControl) Required(value interface{}) *InputDateBaseControl {
	a.Set("required", value)
	return a
}

func (a *InputDateBaseControl) Row(value interface{}) *InputDateBaseControl {
	a.Set("row", value)
	return a
}

func (a *InputDateBaseControl) SaveImmediately(value interface{}) *InputDateBaseControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *InputDateBaseControl) Shortcuts(value interface{}) *InputDateBaseControl {
	a.Set("shortcuts", value)
	return a
}

func (a *InputDateBaseControl) Size(value interface{}) *InputDateBaseControl {
	a.Set("size", value)
	return a
}

func (a *InputDateBaseControl) Static(value interface{}) *InputDateBaseControl {
	a.Set("static", value)
	return a
}

func (a *InputDateBaseControl) StaticClassName(value interface{}) *InputDateBaseControl {
	a.Set("staticClassName", value)
	return a
}

func (a *InputDateBaseControl) StaticInputClassName(value interface{}) *InputDateBaseControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *InputDateBaseControl) StaticLabelClassName(value interface{}) *InputDateBaseControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *InputDateBaseControl) StaticOn(value interface{}) *InputDateBaseControl {
	a.Set("staticOn", value)
	return a
}

func (a *InputDateBaseControl) StaticPlaceholder(value interface{}) *InputDateBaseControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *InputDateBaseControl) StaticSchema(value interface{}) *InputDateBaseControl {
	a.Set("staticSchema", value)
	return a
}

func (a *InputDateBaseControl) Style(value interface{}) *InputDateBaseControl {
	a.Set("style", value)
	return a
}

func (a *InputDateBaseControl) SubmitOnChange(value interface{}) *InputDateBaseControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *InputDateBaseControl) Type(value interface{}) *InputDateBaseControl {
	a.Set("type", value)
	return a
}

func (a *InputDateBaseControl) UseMobileUI(value interface{}) *InputDateBaseControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *InputDateBaseControl) Utc(value interface{}) *InputDateBaseControl {
	a.Set("utc", value)
	return a
}

func (a *InputDateBaseControl) ValidateApi(value interface{}) *InputDateBaseControl {
	a.Set("validateApi", value)
	return a
}

func (a *InputDateBaseControl) ValidateOnChange(value interface{}) *InputDateBaseControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *InputDateBaseControl) ValidationErrors(value interface{}) *InputDateBaseControl {
	a.Set("validationErrors", value)
	return a
}

func (a *InputDateBaseControl) Validations(value interface{}) *InputDateBaseControl {
	a.Set("validations", value)
	return a
}

func (a *InputDateBaseControl) Value(value interface{}) *InputDateBaseControl {
	a.Set("value", value)
	return a
}

func (a *InputDateBaseControl) ValueFormat(value interface{}) *InputDateBaseControl {
	a.Set("valueFormat", value)
	return a
}

func (a *InputDateBaseControl) Visible(value interface{}) *InputDateBaseControl {
	a.Set("visible", value)
	return a
}

func (a *InputDateBaseControl) VisibleOn(value interface{}) *InputDateBaseControl {
	a.Set("visibleOn", value)
	return a
}

func (a *InputDateBaseControl) Width(value interface{}) *InputDateBaseControl {
	a.Set("width", value)
	return a
}
