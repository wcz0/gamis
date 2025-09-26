package renderers

type CheckboxControl struct {
	*BaseRenderer
}

func NewCheckboxControl() *CheckboxControl {
	a := &CheckboxControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "checkbox")
	return a
}

func (a *CheckboxControl) Set(name string, value interface{}) *CheckboxControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *CheckboxControl) AutoFill(value interface{}) *CheckboxControl {
	a.Set("autoFill", value)
	return a
}

func (a *CheckboxControl) Badge(value interface{}) *CheckboxControl {
	a.Set("badge", value)
	return a
}

func (a *CheckboxControl) Checked(value interface{}) *CheckboxControl {
	a.Set("checked", value)
	return a
}

func (a *CheckboxControl) ClassName(value interface{}) *CheckboxControl {
	a.Set("className", value)
	return a
}

func (a *CheckboxControl) ClearValueOnHidden(value interface{}) *CheckboxControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *CheckboxControl) Desc(value interface{}) *CheckboxControl {
	a.Set("desc", value)
	return a
}

func (a *CheckboxControl) Description(value interface{}) *CheckboxControl {
	a.Set("description", value)
	return a
}

func (a *CheckboxControl) DescriptionClassName(value interface{}) *CheckboxControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *CheckboxControl) Disabled(value interface{}) *CheckboxControl {
	a.Set("disabled", value)
	return a
}

func (a *CheckboxControl) DisabledOn(value interface{}) *CheckboxControl {
	a.Set("disabledOn", value)
	return a
}

func (a *CheckboxControl) EditorSetting(value interface{}) *CheckboxControl {
	a.Set("editorSetting", value)
	return a
}

func (a *CheckboxControl) ExtraName(value interface{}) *CheckboxControl {
	a.Set("extraName", value)
	return a
}

func (a *CheckboxControl) FalseValue(value interface{}) *CheckboxControl {
	a.Set("falseValue", value)
	return a
}

func (a *CheckboxControl) Hidden(value interface{}) *CheckboxControl {
	a.Set("hidden", value)
	return a
}

func (a *CheckboxControl) HiddenOn(value interface{}) *CheckboxControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *CheckboxControl) Hint(value interface{}) *CheckboxControl {
	a.Set("hint", value)
	return a
}

func (a *CheckboxControl) Horizontal(value interface{}) *CheckboxControl {
	a.Set("horizontal", value)
	return a
}

func (a *CheckboxControl) Id(value interface{}) *CheckboxControl {
	a.Set("id", value)
	return a
}

func (a *CheckboxControl) InitAutoFill(value interface{}) *CheckboxControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *CheckboxControl) Inline(value interface{}) *CheckboxControl {
	a.Set("inline", value)
	return a
}

func (a *CheckboxControl) InputClassName(value interface{}) *CheckboxControl {
	a.Set("inputClassName", value)
	return a
}

func (a *CheckboxControl) Label(value interface{}) *CheckboxControl {
	a.Set("label", value)
	return a
}

func (a *CheckboxControl) LabelAlign(value interface{}) *CheckboxControl {
	a.Set("labelAlign", value)
	return a
}

func (a *CheckboxControl) LabelClassName(value interface{}) *CheckboxControl {
	a.Set("labelClassName", value)
	return a
}

func (a *CheckboxControl) LabelRemark(value interface{}) *CheckboxControl {
	a.Set("labelRemark", value)
	return a
}

func (a *CheckboxControl) LabelWidth(value interface{}) *CheckboxControl {
	a.Set("labelWidth", value)
	return a
}

func (a *CheckboxControl) Mode(value interface{}) *CheckboxControl {
	a.Set("mode", value)
	return a
}

func (a *CheckboxControl) Name(value interface{}) *CheckboxControl {
	a.Set("name", value)
	return a
}

func (a *CheckboxControl) OnEvent(value interface{}) *CheckboxControl {
	a.Set("onEvent", value)
	return a
}

func (a *CheckboxControl) Option(value interface{}) *CheckboxControl {
	a.Set("option", value)
	return a
}

func (a *CheckboxControl) OptionType(value interface{}) *CheckboxControl {
	a.Set("optionType", value)
	return a
}

func (a *CheckboxControl) Partial(value interface{}) *CheckboxControl {
	a.Set("partial", value)
	return a
}

func (a *CheckboxControl) Placeholder(value interface{}) *CheckboxControl {
	a.Set("placeholder", value)
	return a
}

func (a *CheckboxControl) ReadOnly(value interface{}) *CheckboxControl {
	a.Set("readOnly", value)
	return a
}

func (a *CheckboxControl) ReadOnlyOn(value interface{}) *CheckboxControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *CheckboxControl) Remark(value interface{}) *CheckboxControl {
	a.Set("remark", value)
	return a
}

func (a *CheckboxControl) Required(value interface{}) *CheckboxControl {
	a.Set("required", value)
	return a
}

func (a *CheckboxControl) Row(value interface{}) *CheckboxControl {
	a.Set("row", value)
	return a
}

func (a *CheckboxControl) SaveImmediately(value interface{}) *CheckboxControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *CheckboxControl) Size(value interface{}) *CheckboxControl {
	a.Set("size", value)
	return a
}

func (a *CheckboxControl) Static(value interface{}) *CheckboxControl {
	a.Set("static", value)
	return a
}

func (a *CheckboxControl) StaticClassName(value interface{}) *CheckboxControl {
	a.Set("staticClassName", value)
	return a
}

func (a *CheckboxControl) StaticInputClassName(value interface{}) *CheckboxControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *CheckboxControl) StaticLabelClassName(value interface{}) *CheckboxControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *CheckboxControl) StaticOn(value interface{}) *CheckboxControl {
	a.Set("staticOn", value)
	return a
}

func (a *CheckboxControl) StaticPlaceholder(value interface{}) *CheckboxControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *CheckboxControl) StaticSchema(value interface{}) *CheckboxControl {
	a.Set("staticSchema", value)
	return a
}

func (a *CheckboxControl) Style(value interface{}) *CheckboxControl {
	a.Set("style", value)
	return a
}

func (a *CheckboxControl) SubmitOnChange(value interface{}) *CheckboxControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *CheckboxControl) TestIdBuilder(value interface{}) *CheckboxControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *CheckboxControl) TrueValue(value interface{}) *CheckboxControl {
	a.Set("trueValue", value)
	return a
}

func (a *CheckboxControl) Type(value interface{}) *CheckboxControl {
	a.Set("type", value)
	return a
}

func (a *CheckboxControl) UseMobileUI(value interface{}) *CheckboxControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *CheckboxControl) ValidateApi(value interface{}) *CheckboxControl {
	a.Set("validateApi", value)
	return a
}

func (a *CheckboxControl) ValidateOnChange(value interface{}) *CheckboxControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *CheckboxControl) ValidationErrors(value interface{}) *CheckboxControl {
	a.Set("validationErrors", value)
	return a
}

func (a *CheckboxControl) Validations(value interface{}) *CheckboxControl {
	a.Set("validations", value)
	return a
}

func (a *CheckboxControl) Value(value interface{}) *CheckboxControl {
	a.Set("value", value)
	return a
}

func (a *CheckboxControl) Visible(value interface{}) *CheckboxControl {
	a.Set("visible", value)
	return a
}

func (a *CheckboxControl) VisibleOn(value interface{}) *CheckboxControl {
	a.Set("visibleOn", value)
	return a
}

func (a *CheckboxControl) Width(value interface{}) *CheckboxControl {
	a.Set("width", value)
	return a
}
