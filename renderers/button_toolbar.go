package renderers

type ButtonToolbar struct {
	*BaseRenderer
}

func NewButtonToolbar() *ButtonToolbar {
	a := &ButtonToolbar{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "button-toolbar")
	return a
}

func (a *ButtonToolbar) Set(name string, value interface{}) *ButtonToolbar {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *ButtonToolbar) AutoFill(value interface{}) *ButtonToolbar {
	a.Set("autoFill", value)
	return a
}

func (a *ButtonToolbar) Buttons(value interface{}) *ButtonToolbar {
	a.Set("buttons", value)
	return a
}

func (a *ButtonToolbar) ClassName(value interface{}) *ButtonToolbar {
	a.Set("className", value)
	return a
}

func (a *ButtonToolbar) ClearValueOnHidden(value interface{}) *ButtonToolbar {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *ButtonToolbar) Desc(value interface{}) *ButtonToolbar {
	a.Set("desc", value)
	return a
}

func (a *ButtonToolbar) Description(value interface{}) *ButtonToolbar {
	a.Set("description", value)
	return a
}

func (a *ButtonToolbar) DescriptionClassName(value interface{}) *ButtonToolbar {
	a.Set("descriptionClassName", value)
	return a
}

func (a *ButtonToolbar) Disabled(value interface{}) *ButtonToolbar {
	a.Set("disabled", value)
	return a
}

func (a *ButtonToolbar) DisabledOn(value interface{}) *ButtonToolbar {
	a.Set("disabledOn", value)
	return a
}

func (a *ButtonToolbar) EditorSetting(value interface{}) *ButtonToolbar {
	a.Set("editorSetting", value)
	return a
}

func (a *ButtonToolbar) ExtraName(value interface{}) *ButtonToolbar {
	a.Set("extraName", value)
	return a
}

func (a *ButtonToolbar) Hidden(value interface{}) *ButtonToolbar {
	a.Set("hidden", value)
	return a
}

func (a *ButtonToolbar) HiddenOn(value interface{}) *ButtonToolbar {
	a.Set("hiddenOn", value)
	return a
}

func (a *ButtonToolbar) Hint(value interface{}) *ButtonToolbar {
	a.Set("hint", value)
	return a
}

func (a *ButtonToolbar) Horizontal(value interface{}) *ButtonToolbar {
	a.Set("horizontal", value)
	return a
}

func (a *ButtonToolbar) Id(value interface{}) *ButtonToolbar {
	a.Set("id", value)
	return a
}

func (a *ButtonToolbar) InitAutoFill(value interface{}) *ButtonToolbar {
	a.Set("initAutoFill", value)
	return a
}

func (a *ButtonToolbar) Inline(value interface{}) *ButtonToolbar {
	a.Set("inline", value)
	return a
}

func (a *ButtonToolbar) InputClassName(value interface{}) *ButtonToolbar {
	a.Set("inputClassName", value)
	return a
}

func (a *ButtonToolbar) Label(value interface{}) *ButtonToolbar {
	a.Set("label", value)
	return a
}

func (a *ButtonToolbar) LabelAlign(value interface{}) *ButtonToolbar {
	a.Set("labelAlign", value)
	return a
}

func (a *ButtonToolbar) LabelClassName(value interface{}) *ButtonToolbar {
	a.Set("labelClassName", value)
	return a
}

func (a *ButtonToolbar) LabelOverflow(value interface{}) *ButtonToolbar {
	a.Set("labelOverflow", value)
	return a
}

func (a *ButtonToolbar) LabelRemark(value interface{}) *ButtonToolbar {
	a.Set("labelRemark", value)
	return a
}

func (a *ButtonToolbar) LabelWidth(value interface{}) *ButtonToolbar {
	a.Set("labelWidth", value)
	return a
}

func (a *ButtonToolbar) Mode(value interface{}) *ButtonToolbar {
	a.Set("mode", value)
	return a
}

func (a *ButtonToolbar) Name(value interface{}) *ButtonToolbar {
	a.Set("name", value)
	return a
}

func (a *ButtonToolbar) OnEvent(value interface{}) *ButtonToolbar {
	a.Set("onEvent", value)
	return a
}

func (a *ButtonToolbar) Placeholder(value interface{}) *ButtonToolbar {
	a.Set("placeholder", value)
	return a
}

func (a *ButtonToolbar) ReadOnly(value interface{}) *ButtonToolbar {
	a.Set("readOnly", value)
	return a
}

func (a *ButtonToolbar) ReadOnlyOn(value interface{}) *ButtonToolbar {
	a.Set("readOnlyOn", value)
	return a
}

func (a *ButtonToolbar) Remark(value interface{}) *ButtonToolbar {
	a.Set("remark", value)
	return a
}

func (a *ButtonToolbar) Required(value interface{}) *ButtonToolbar {
	a.Set("required", value)
	return a
}

func (a *ButtonToolbar) Row(value interface{}) *ButtonToolbar {
	a.Set("row", value)
	return a
}

func (a *ButtonToolbar) Static(value interface{}) *ButtonToolbar {
	a.Set("static", value)
	return a
}

func (a *ButtonToolbar) StaticClassName(value interface{}) *ButtonToolbar {
	a.Set("staticClassName", value)
	return a
}

func (a *ButtonToolbar) StaticInputClassName(value interface{}) *ButtonToolbar {
	a.Set("staticInputClassName", value)
	return a
}

func (a *ButtonToolbar) StaticLabelClassName(value interface{}) *ButtonToolbar {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *ButtonToolbar) StaticOn(value interface{}) *ButtonToolbar {
	a.Set("staticOn", value)
	return a
}

func (a *ButtonToolbar) StaticPlaceholder(value interface{}) *ButtonToolbar {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *ButtonToolbar) StaticSchema(value interface{}) *ButtonToolbar {
	a.Set("staticSchema", value)
	return a
}

func (a *ButtonToolbar) Style(value interface{}) *ButtonToolbar {
	a.Set("style", value)
	return a
}

func (a *ButtonToolbar) SubmitOnChange(value interface{}) *ButtonToolbar {
	a.Set("submitOnChange", value)
	return a
}

func (a *ButtonToolbar) TestIdBuilder(value interface{}) *ButtonToolbar {
	a.Set("testIdBuilder", value)
	return a
}

func (a *ButtonToolbar) Testid(value interface{}) *ButtonToolbar {
	a.Set("testid", value)
	return a
}

func (a *ButtonToolbar) Type(value interface{}) *ButtonToolbar {
	a.Set("type", value)
	return a
}

func (a *ButtonToolbar) UseMobileUI(value interface{}) *ButtonToolbar {
	a.Set("useMobileUI", value)
	return a
}

func (a *ButtonToolbar) ValidateApi(value interface{}) *ButtonToolbar {
	a.Set("validateApi", value)
	return a
}

func (a *ButtonToolbar) ValidateOnChange(value interface{}) *ButtonToolbar {
	a.Set("validateOnChange", value)
	return a
}

func (a *ButtonToolbar) ValidationErrors(value interface{}) *ButtonToolbar {
	a.Set("validationErrors", value)
	return a
}

func (a *ButtonToolbar) Validations(value interface{}) *ButtonToolbar {
	a.Set("validations", value)
	return a
}

func (a *ButtonToolbar) Value(value interface{}) *ButtonToolbar {
	a.Set("value", value)
	return a
}

func (a *ButtonToolbar) Visible(value interface{}) *ButtonToolbar {
	a.Set("visible", value)
	return a
}

func (a *ButtonToolbar) VisibleOn(value interface{}) *ButtonToolbar {
	a.Set("visibleOn", value)
	return a
}
