package renderers

type FormControl struct {
	*BaseRenderer
}

func NewFormControl() *FormControl {
	a := &FormControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "control")
	return a
}

func (a *FormControl) Set(name string, value interface{}) *FormControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *FormControl) AutoFill(value interface{}) *FormControl {
	a.Set("autoFill", value)
	return a
}

func (a *FormControl) Body(value interface{}) *FormControl {
	a.Set("body", value)
	return a
}

func (a *FormControl) ClassName(value interface{}) *FormControl {
	a.Set("className", value)
	return a
}

func (a *FormControl) ClearValueOnHidden(value interface{}) *FormControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *FormControl) Desc(value interface{}) *FormControl {
	a.Set("desc", value)
	return a
}

func (a *FormControl) Description(value interface{}) *FormControl {
	a.Set("description", value)
	return a
}

func (a *FormControl) DescriptionClassName(value interface{}) *FormControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *FormControl) Disabled(value interface{}) *FormControl {
	a.Set("disabled", value)
	return a
}

func (a *FormControl) DisabledOn(value interface{}) *FormControl {
	a.Set("disabledOn", value)
	return a
}

func (a *FormControl) EditorSetting(value interface{}) *FormControl {
	a.Set("editorSetting", value)
	return a
}

func (a *FormControl) ExtraName(value interface{}) *FormControl {
	a.Set("extraName", value)
	return a
}

func (a *FormControl) Hidden(value interface{}) *FormControl {
	a.Set("hidden", value)
	return a
}

func (a *FormControl) HiddenOn(value interface{}) *FormControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *FormControl) Hint(value interface{}) *FormControl {
	a.Set("hint", value)
	return a
}

func (a *FormControl) Horizontal(value interface{}) *FormControl {
	a.Set("horizontal", value)
	return a
}

func (a *FormControl) Id(value interface{}) *FormControl {
	a.Set("id", value)
	return a
}

func (a *FormControl) InitAutoFill(value interface{}) *FormControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *FormControl) Inline(value interface{}) *FormControl {
	a.Set("inline", value)
	return a
}

func (a *FormControl) InputClassName(value interface{}) *FormControl {
	a.Set("inputClassName", value)
	return a
}

func (a *FormControl) Label(value interface{}) *FormControl {
	a.Set("label", value)
	return a
}

func (a *FormControl) LabelAlign(value interface{}) *FormControl {
	a.Set("labelAlign", value)
	return a
}

func (a *FormControl) LabelClassName(value interface{}) *FormControl {
	a.Set("labelClassName", value)
	return a
}

func (a *FormControl) LabelRemark(value interface{}) *FormControl {
	a.Set("labelRemark", value)
	return a
}

func (a *FormControl) LabelWidth(value interface{}) *FormControl {
	a.Set("labelWidth", value)
	return a
}

func (a *FormControl) Mode(value interface{}) *FormControl {
	a.Set("mode", value)
	return a
}

func (a *FormControl) Name(value interface{}) *FormControl {
	a.Set("name", value)
	return a
}

func (a *FormControl) OnEvent(value interface{}) *FormControl {
	a.Set("onEvent", value)
	return a
}

func (a *FormControl) Placeholder(value interface{}) *FormControl {
	a.Set("placeholder", value)
	return a
}

func (a *FormControl) ReadOnly(value interface{}) *FormControl {
	a.Set("readOnly", value)
	return a
}

func (a *FormControl) ReadOnlyOn(value interface{}) *FormControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *FormControl) Remark(value interface{}) *FormControl {
	a.Set("remark", value)
	return a
}

func (a *FormControl) Required(value interface{}) *FormControl {
	a.Set("required", value)
	return a
}

func (a *FormControl) Row(value interface{}) *FormControl {
	a.Set("row", value)
	return a
}

func (a *FormControl) SaveImmediately(value interface{}) *FormControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *FormControl) Size(value interface{}) *FormControl {
	a.Set("size", value)
	return a
}

func (a *FormControl) Static(value interface{}) *FormControl {
	a.Set("static", value)
	return a
}

func (a *FormControl) StaticClassName(value interface{}) *FormControl {
	a.Set("staticClassName", value)
	return a
}

func (a *FormControl) StaticInputClassName(value interface{}) *FormControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *FormControl) StaticLabelClassName(value interface{}) *FormControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *FormControl) StaticOn(value interface{}) *FormControl {
	a.Set("staticOn", value)
	return a
}

func (a *FormControl) StaticPlaceholder(value interface{}) *FormControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *FormControl) StaticSchema(value interface{}) *FormControl {
	a.Set("staticSchema", value)
	return a
}

func (a *FormControl) Style(value interface{}) *FormControl {
	a.Set("style", value)
	return a
}

func (a *FormControl) SubmitOnChange(value interface{}) *FormControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *FormControl) TestIdBuilder(value interface{}) *FormControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *FormControl) Type(value interface{}) *FormControl {
	a.Set("type", value)
	return a
}

func (a *FormControl) UseMobileUI(value interface{}) *FormControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *FormControl) ValidateApi(value interface{}) *FormControl {
	a.Set("validateApi", value)
	return a
}

func (a *FormControl) ValidateOnChange(value interface{}) *FormControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *FormControl) ValidationErrors(value interface{}) *FormControl {
	a.Set("validationErrors", value)
	return a
}

func (a *FormControl) Validations(value interface{}) *FormControl {
	a.Set("validations", value)
	return a
}

func (a *FormControl) Value(value interface{}) *FormControl {
	a.Set("value", value)
	return a
}

func (a *FormControl) Visible(value interface{}) *FormControl {
	a.Set("visible", value)
	return a
}

func (a *FormControl) VisibleOn(value interface{}) *FormControl {
	a.Set("visibleOn", value)
	return a
}

func (a *FormControl) Width(value interface{}) *FormControl {
	a.Set("width", value)
	return a
}
