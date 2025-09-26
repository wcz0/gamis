package renderers

type RadioControl struct {
	*BaseRenderer
}

func NewRadioControl() *RadioControl {
	a := &RadioControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "radio")
	return a
}

func (a *RadioControl) Set(name string, value interface{}) *RadioControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *RadioControl) AutoFill(value interface{}) *RadioControl {
	a.Set("autoFill", value)
	return a
}

func (a *RadioControl) Badge(value interface{}) *RadioControl {
	a.Set("badge", value)
	return a
}

func (a *RadioControl) ClassName(value interface{}) *RadioControl {
	a.Set("className", value)
	return a
}

func (a *RadioControl) ClearValueOnHidden(value interface{}) *RadioControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *RadioControl) Desc(value interface{}) *RadioControl {
	a.Set("desc", value)
	return a
}

func (a *RadioControl) Description(value interface{}) *RadioControl {
	a.Set("description", value)
	return a
}

func (a *RadioControl) DescriptionClassName(value interface{}) *RadioControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *RadioControl) Disabled(value interface{}) *RadioControl {
	a.Set("disabled", value)
	return a
}

func (a *RadioControl) DisabledOn(value interface{}) *RadioControl {
	a.Set("disabledOn", value)
	return a
}

func (a *RadioControl) EditorSetting(value interface{}) *RadioControl {
	a.Set("editorSetting", value)
	return a
}

func (a *RadioControl) ExtraName(value interface{}) *RadioControl {
	a.Set("extraName", value)
	return a
}

func (a *RadioControl) FalseValue(value interface{}) *RadioControl {
	a.Set("falseValue", value)
	return a
}

func (a *RadioControl) Hidden(value interface{}) *RadioControl {
	a.Set("hidden", value)
	return a
}

func (a *RadioControl) HiddenOn(value interface{}) *RadioControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *RadioControl) Hint(value interface{}) *RadioControl {
	a.Set("hint", value)
	return a
}

func (a *RadioControl) Horizontal(value interface{}) *RadioControl {
	a.Set("horizontal", value)
	return a
}

func (a *RadioControl) Id(value interface{}) *RadioControl {
	a.Set("id", value)
	return a
}

func (a *RadioControl) InitAutoFill(value interface{}) *RadioControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *RadioControl) Inline(value interface{}) *RadioControl {
	a.Set("inline", value)
	return a
}

func (a *RadioControl) InputClassName(value interface{}) *RadioControl {
	a.Set("inputClassName", value)
	return a
}

func (a *RadioControl) Label(value interface{}) *RadioControl {
	a.Set("label", value)
	return a
}

func (a *RadioControl) LabelAlign(value interface{}) *RadioControl {
	a.Set("labelAlign", value)
	return a
}

func (a *RadioControl) LabelClassName(value interface{}) *RadioControl {
	a.Set("labelClassName", value)
	return a
}

func (a *RadioControl) LabelRemark(value interface{}) *RadioControl {
	a.Set("labelRemark", value)
	return a
}

func (a *RadioControl) LabelWidth(value interface{}) *RadioControl {
	a.Set("labelWidth", value)
	return a
}

func (a *RadioControl) Mode(value interface{}) *RadioControl {
	a.Set("mode", value)
	return a
}

func (a *RadioControl) Name(value interface{}) *RadioControl {
	a.Set("name", value)
	return a
}

func (a *RadioControl) OnEvent(value interface{}) *RadioControl {
	a.Set("onEvent", value)
	return a
}

func (a *RadioControl) Option(value interface{}) *RadioControl {
	a.Set("option", value)
	return a
}

func (a *RadioControl) OptionType(value interface{}) *RadioControl {
	a.Set("optionType", value)
	return a
}

func (a *RadioControl) Partial(value interface{}) *RadioControl {
	a.Set("partial", value)
	return a
}

func (a *RadioControl) Placeholder(value interface{}) *RadioControl {
	a.Set("placeholder", value)
	return a
}

func (a *RadioControl) ReadOnly(value interface{}) *RadioControl {
	a.Set("readOnly", value)
	return a
}

func (a *RadioControl) ReadOnlyOn(value interface{}) *RadioControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *RadioControl) Remark(value interface{}) *RadioControl {
	a.Set("remark", value)
	return a
}

func (a *RadioControl) Required(value interface{}) *RadioControl {
	a.Set("required", value)
	return a
}

func (a *RadioControl) Row(value interface{}) *RadioControl {
	a.Set("row", value)
	return a
}

func (a *RadioControl) SaveImmediately(value interface{}) *RadioControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *RadioControl) Size(value interface{}) *RadioControl {
	a.Set("size", value)
	return a
}

func (a *RadioControl) Static(value interface{}) *RadioControl {
	a.Set("static", value)
	return a
}

func (a *RadioControl) StaticClassName(value interface{}) *RadioControl {
	a.Set("staticClassName", value)
	return a
}

func (a *RadioControl) StaticInputClassName(value interface{}) *RadioControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *RadioControl) StaticLabelClassName(value interface{}) *RadioControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *RadioControl) StaticOn(value interface{}) *RadioControl {
	a.Set("staticOn", value)
	return a
}

func (a *RadioControl) StaticPlaceholder(value interface{}) *RadioControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *RadioControl) StaticSchema(value interface{}) *RadioControl {
	a.Set("staticSchema", value)
	return a
}

func (a *RadioControl) Style(value interface{}) *RadioControl {
	a.Set("style", value)
	return a
}

func (a *RadioControl) SubmitOnChange(value interface{}) *RadioControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *RadioControl) TestIdBuilder(value interface{}) *RadioControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *RadioControl) TrueValue(value interface{}) *RadioControl {
	a.Set("trueValue", value)
	return a
}

func (a *RadioControl) Type(value interface{}) *RadioControl {
	a.Set("type", value)
	return a
}

func (a *RadioControl) UseMobileUI(value interface{}) *RadioControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *RadioControl) ValidateApi(value interface{}) *RadioControl {
	a.Set("validateApi", value)
	return a
}

func (a *RadioControl) ValidateOnChange(value interface{}) *RadioControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *RadioControl) ValidationErrors(value interface{}) *RadioControl {
	a.Set("validationErrors", value)
	return a
}

func (a *RadioControl) Validations(value interface{}) *RadioControl {
	a.Set("validations", value)
	return a
}

func (a *RadioControl) Value(value interface{}) *RadioControl {
	a.Set("value", value)
	return a
}

func (a *RadioControl) Visible(value interface{}) *RadioControl {
	a.Set("visible", value)
	return a
}

func (a *RadioControl) VisibleOn(value interface{}) *RadioControl {
	a.Set("visibleOn", value)
	return a
}

func (a *RadioControl) Width(value interface{}) *RadioControl {
	a.Set("width", value)
	return a
}
