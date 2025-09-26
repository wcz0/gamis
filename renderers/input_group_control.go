package renderers

type InputGroupControl struct {
	*BaseRenderer
}

func NewInputGroupControl() *InputGroupControl {
	a := &InputGroupControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "input-group")
	return a
}

func (a *InputGroupControl) Set(name string, value interface{}) *InputGroupControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *InputGroupControl) AutoFill(value interface{}) *InputGroupControl {
	a.Set("autoFill", value)
	return a
}

func (a *InputGroupControl) Body(value interface{}) *InputGroupControl {
	a.Set("body", value)
	return a
}

func (a *InputGroupControl) ClassName(value interface{}) *InputGroupControl {
	a.Set("className", value)
	return a
}

func (a *InputGroupControl) ClearValueOnHidden(value interface{}) *InputGroupControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *InputGroupControl) Desc(value interface{}) *InputGroupControl {
	a.Set("desc", value)
	return a
}

func (a *InputGroupControl) Description(value interface{}) *InputGroupControl {
	a.Set("description", value)
	return a
}

func (a *InputGroupControl) DescriptionClassName(value interface{}) *InputGroupControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *InputGroupControl) Disabled(value interface{}) *InputGroupControl {
	a.Set("disabled", value)
	return a
}

func (a *InputGroupControl) DisabledOn(value interface{}) *InputGroupControl {
	a.Set("disabledOn", value)
	return a
}

func (a *InputGroupControl) EditorSetting(value interface{}) *InputGroupControl {
	a.Set("editorSetting", value)
	return a
}

func (a *InputGroupControl) ExtraName(value interface{}) *InputGroupControl {
	a.Set("extraName", value)
	return a
}

func (a *InputGroupControl) Hidden(value interface{}) *InputGroupControl {
	a.Set("hidden", value)
	return a
}

func (a *InputGroupControl) HiddenOn(value interface{}) *InputGroupControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *InputGroupControl) Hint(value interface{}) *InputGroupControl {
	a.Set("hint", value)
	return a
}

func (a *InputGroupControl) Horizontal(value interface{}) *InputGroupControl {
	a.Set("horizontal", value)
	return a
}

func (a *InputGroupControl) Id(value interface{}) *InputGroupControl {
	a.Set("id", value)
	return a
}

func (a *InputGroupControl) InitAutoFill(value interface{}) *InputGroupControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *InputGroupControl) Inline(value interface{}) *InputGroupControl {
	a.Set("inline", value)
	return a
}

func (a *InputGroupControl) InputClassName(value interface{}) *InputGroupControl {
	a.Set("inputClassName", value)
	return a
}

func (a *InputGroupControl) Label(value interface{}) *InputGroupControl {
	a.Set("label", value)
	return a
}

func (a *InputGroupControl) LabelAlign(value interface{}) *InputGroupControl {
	a.Set("labelAlign", value)
	return a
}

func (a *InputGroupControl) LabelClassName(value interface{}) *InputGroupControl {
	a.Set("labelClassName", value)
	return a
}

func (a *InputGroupControl) LabelRemark(value interface{}) *InputGroupControl {
	a.Set("labelRemark", value)
	return a
}

func (a *InputGroupControl) LabelWidth(value interface{}) *InputGroupControl {
	a.Set("labelWidth", value)
	return a
}

func (a *InputGroupControl) Mode(value interface{}) *InputGroupControl {
	a.Set("mode", value)
	return a
}

func (a *InputGroupControl) Name(value interface{}) *InputGroupControl {
	a.Set("name", value)
	return a
}

func (a *InputGroupControl) OnEvent(value interface{}) *InputGroupControl {
	a.Set("onEvent", value)
	return a
}

func (a *InputGroupControl) Placeholder(value interface{}) *InputGroupControl {
	a.Set("placeholder", value)
	return a
}

func (a *InputGroupControl) ReadOnly(value interface{}) *InputGroupControl {
	a.Set("readOnly", value)
	return a
}

func (a *InputGroupControl) ReadOnlyOn(value interface{}) *InputGroupControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *InputGroupControl) Remark(value interface{}) *InputGroupControl {
	a.Set("remark", value)
	return a
}

func (a *InputGroupControl) Required(value interface{}) *InputGroupControl {
	a.Set("required", value)
	return a
}

func (a *InputGroupControl) Row(value interface{}) *InputGroupControl {
	a.Set("row", value)
	return a
}

func (a *InputGroupControl) SaveImmediately(value interface{}) *InputGroupControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *InputGroupControl) Size(value interface{}) *InputGroupControl {
	a.Set("size", value)
	return a
}

func (a *InputGroupControl) Static(value interface{}) *InputGroupControl {
	a.Set("static", value)
	return a
}

func (a *InputGroupControl) StaticClassName(value interface{}) *InputGroupControl {
	a.Set("staticClassName", value)
	return a
}

func (a *InputGroupControl) StaticInputClassName(value interface{}) *InputGroupControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *InputGroupControl) StaticLabelClassName(value interface{}) *InputGroupControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *InputGroupControl) StaticOn(value interface{}) *InputGroupControl {
	a.Set("staticOn", value)
	return a
}

func (a *InputGroupControl) StaticPlaceholder(value interface{}) *InputGroupControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *InputGroupControl) StaticSchema(value interface{}) *InputGroupControl {
	a.Set("staticSchema", value)
	return a
}

func (a *InputGroupControl) Style(value interface{}) *InputGroupControl {
	a.Set("style", value)
	return a
}

func (a *InputGroupControl) SubmitOnChange(value interface{}) *InputGroupControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *InputGroupControl) TestIdBuilder(value interface{}) *InputGroupControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *InputGroupControl) Type(value interface{}) *InputGroupControl {
	a.Set("type", value)
	return a
}

func (a *InputGroupControl) UseMobileUI(value interface{}) *InputGroupControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *InputGroupControl) ValidateApi(value interface{}) *InputGroupControl {
	a.Set("validateApi", value)
	return a
}

func (a *InputGroupControl) ValidateOnChange(value interface{}) *InputGroupControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *InputGroupControl) ValidationConfig(value interface{}) *InputGroupControl {
	a.Set("validationConfig", value)
	return a
}

func (a *InputGroupControl) ValidationErrors(value interface{}) *InputGroupControl {
	a.Set("validationErrors", value)
	return a
}

func (a *InputGroupControl) Validations(value interface{}) *InputGroupControl {
	a.Set("validations", value)
	return a
}

func (a *InputGroupControl) Value(value interface{}) *InputGroupControl {
	a.Set("value", value)
	return a
}

func (a *InputGroupControl) Visible(value interface{}) *InputGroupControl {
	a.Set("visible", value)
	return a
}

func (a *InputGroupControl) VisibleOn(value interface{}) *InputGroupControl {
	a.Set("visibleOn", value)
	return a
}

func (a *InputGroupControl) Width(value interface{}) *InputGroupControl {
	a.Set("width", value)
	return a
}
