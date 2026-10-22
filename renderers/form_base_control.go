package renderers

type FormBaseControl struct {
	*BaseRenderer
}

func NewFormBaseControl() *FormBaseControl {
	a := &FormBaseControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "form")
	return a
}

func (a *FormBaseControl) Set(name string, value interface{}) *FormBaseControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *FormBaseControl) AutoFill(value interface{}) *FormBaseControl {
	a.Set("autoFill", value)
	return a
}

func (a *FormBaseControl) ClassName(value interface{}) *FormBaseControl {
	a.Set("className", value)
	return a
}

func (a *FormBaseControl) ClearValueOnHidden(value interface{}) *FormBaseControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *FormBaseControl) Desc(value interface{}) *FormBaseControl {
	a.Set("desc", value)
	return a
}

func (a *FormBaseControl) Description(value interface{}) *FormBaseControl {
	a.Set("description", value)
	return a
}

func (a *FormBaseControl) DescriptionClassName(value interface{}) *FormBaseControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *FormBaseControl) Disabled(value interface{}) *FormBaseControl {
	a.Set("disabled", value)
	return a
}

func (a *FormBaseControl) DisabledOn(value interface{}) *FormBaseControl {
	a.Set("disabledOn", value)
	return a
}

func (a *FormBaseControl) EditorSetting(value interface{}) *FormBaseControl {
	a.Set("editorSetting", value)
	return a
}

func (a *FormBaseControl) ExtraName(value interface{}) *FormBaseControl {
	a.Set("extraName", value)
	return a
}

func (a *FormBaseControl) Hidden(value interface{}) *FormBaseControl {
	a.Set("hidden", value)
	return a
}

func (a *FormBaseControl) HiddenOn(value interface{}) *FormBaseControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *FormBaseControl) Hint(value interface{}) *FormBaseControl {
	a.Set("hint", value)
	return a
}

func (a *FormBaseControl) Horizontal(value interface{}) *FormBaseControl {
	a.Set("horizontal", value)
	return a
}

func (a *FormBaseControl) Id(value interface{}) *FormBaseControl {
	a.Set("id", value)
	return a
}

func (a *FormBaseControl) InitAutoFill(value interface{}) *FormBaseControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *FormBaseControl) Inline(value interface{}) *FormBaseControl {
	a.Set("inline", value)
	return a
}

func (a *FormBaseControl) InputClassName(value interface{}) *FormBaseControl {
	a.Set("inputClassName", value)
	return a
}

func (a *FormBaseControl) Label(value interface{}) *FormBaseControl {
	a.Set("label", value)
	return a
}

func (a *FormBaseControl) LabelAlign(value interface{}) *FormBaseControl {
	a.Set("labelAlign", value)
	return a
}

func (a *FormBaseControl) LabelClassName(value interface{}) *FormBaseControl {
	a.Set("labelClassName", value)
	return a
}

func (a *FormBaseControl) LabelOverflow(value interface{}) *FormBaseControl {
	a.Set("labelOverflow", value)
	return a
}

func (a *FormBaseControl) LabelRemark(value interface{}) *FormBaseControl {
	a.Set("labelRemark", value)
	return a
}

func (a *FormBaseControl) LabelWidth(value interface{}) *FormBaseControl {
	a.Set("labelWidth", value)
	return a
}

func (a *FormBaseControl) Mode(value interface{}) *FormBaseControl {
	a.Set("mode", value)
	return a
}

func (a *FormBaseControl) Name(value interface{}) *FormBaseControl {
	a.Set("name", value)
	return a
}

func (a *FormBaseControl) OnEvent(value interface{}) *FormBaseControl {
	a.Set("onEvent", value)
	return a
}

func (a *FormBaseControl) Placeholder(value interface{}) *FormBaseControl {
	a.Set("placeholder", value)
	return a
}

func (a *FormBaseControl) ReadOnly(value interface{}) *FormBaseControl {
	a.Set("readOnly", value)
	return a
}

func (a *FormBaseControl) ReadOnlyOn(value interface{}) *FormBaseControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *FormBaseControl) Remark(value interface{}) *FormBaseControl {
	a.Set("remark", value)
	return a
}

func (a *FormBaseControl) Required(value interface{}) *FormBaseControl {
	a.Set("required", value)
	return a
}

func (a *FormBaseControl) Row(value interface{}) *FormBaseControl {
	a.Set("row", value)
	return a
}

func (a *FormBaseControl) SaveImmediately(value interface{}) *FormBaseControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *FormBaseControl) Size(value interface{}) *FormBaseControl {
	a.Set("size", value)
	return a
}

func (a *FormBaseControl) Static(value interface{}) *FormBaseControl {
	a.Set("static", value)
	return a
}

func (a *FormBaseControl) StaticClassName(value interface{}) *FormBaseControl {
	a.Set("staticClassName", value)
	return a
}

func (a *FormBaseControl) StaticInputClassName(value interface{}) *FormBaseControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *FormBaseControl) StaticLabelClassName(value interface{}) *FormBaseControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *FormBaseControl) StaticOn(value interface{}) *FormBaseControl {
	a.Set("staticOn", value)
	return a
}

func (a *FormBaseControl) StaticPlaceholder(value interface{}) *FormBaseControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *FormBaseControl) StaticSchema(value interface{}) *FormBaseControl {
	a.Set("staticSchema", value)
	return a
}

func (a *FormBaseControl) Style(value interface{}) *FormBaseControl {
	a.Set("style", value)
	return a
}

func (a *FormBaseControl) SubmitOnChange(value interface{}) *FormBaseControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *FormBaseControl) Type(value interface{}) *FormBaseControl {
	a.Set("type", value)
	return a
}

func (a *FormBaseControl) UseMobileUI(value interface{}) *FormBaseControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *FormBaseControl) ValidateApi(value interface{}) *FormBaseControl {
	a.Set("validateApi", value)
	return a
}

func (a *FormBaseControl) ValidateOnChange(value interface{}) *FormBaseControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *FormBaseControl) ValidationErrors(value interface{}) *FormBaseControl {
	a.Set("validationErrors", value)
	return a
}

func (a *FormBaseControl) Validations(value interface{}) *FormBaseControl {
	a.Set("validations", value)
	return a
}

func (a *FormBaseControl) Value(value interface{}) *FormBaseControl {
	a.Set("value", value)
	return a
}

func (a *FormBaseControl) Visible(value interface{}) *FormBaseControl {
	a.Set("visible", value)
	return a
}

func (a *FormBaseControl) VisibleOn(value interface{}) *FormBaseControl {
	a.Set("visibleOn", value)
	return a
}

func (a *FormBaseControl) Width(value interface{}) *FormBaseControl {
	a.Set("width", value)
	return a
}
