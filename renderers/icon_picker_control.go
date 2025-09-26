package renderers

type IconPickerControl struct {
	*BaseRenderer
}

func NewIconPickerControl() *IconPickerControl {
	a := &IconPickerControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "icon-picker")
	return a
}

func (a *IconPickerControl) Set(name string, value interface{}) *IconPickerControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *IconPickerControl) AutoFill(value interface{}) *IconPickerControl {
	a.Set("autoFill", value)
	return a
}

func (a *IconPickerControl) ClassName(value interface{}) *IconPickerControl {
	a.Set("className", value)
	return a
}

func (a *IconPickerControl) ClearValueOnHidden(value interface{}) *IconPickerControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *IconPickerControl) Desc(value interface{}) *IconPickerControl {
	a.Set("desc", value)
	return a
}

func (a *IconPickerControl) Description(value interface{}) *IconPickerControl {
	a.Set("description", value)
	return a
}

func (a *IconPickerControl) DescriptionClassName(value interface{}) *IconPickerControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *IconPickerControl) Disabled(value interface{}) *IconPickerControl {
	a.Set("disabled", value)
	return a
}

func (a *IconPickerControl) DisabledOn(value interface{}) *IconPickerControl {
	a.Set("disabledOn", value)
	return a
}

func (a *IconPickerControl) EditorSetting(value interface{}) *IconPickerControl {
	a.Set("editorSetting", value)
	return a
}

func (a *IconPickerControl) ExtraName(value interface{}) *IconPickerControl {
	a.Set("extraName", value)
	return a
}

func (a *IconPickerControl) Hidden(value interface{}) *IconPickerControl {
	a.Set("hidden", value)
	return a
}

func (a *IconPickerControl) HiddenOn(value interface{}) *IconPickerControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *IconPickerControl) Hint(value interface{}) *IconPickerControl {
	a.Set("hint", value)
	return a
}

func (a *IconPickerControl) Horizontal(value interface{}) *IconPickerControl {
	a.Set("horizontal", value)
	return a
}

func (a *IconPickerControl) Id(value interface{}) *IconPickerControl {
	a.Set("id", value)
	return a
}

func (a *IconPickerControl) InitAutoFill(value interface{}) *IconPickerControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *IconPickerControl) Inline(value interface{}) *IconPickerControl {
	a.Set("inline", value)
	return a
}

func (a *IconPickerControl) InputClassName(value interface{}) *IconPickerControl {
	a.Set("inputClassName", value)
	return a
}

func (a *IconPickerControl) Label(value interface{}) *IconPickerControl {
	a.Set("label", value)
	return a
}

func (a *IconPickerControl) LabelAlign(value interface{}) *IconPickerControl {
	a.Set("labelAlign", value)
	return a
}

func (a *IconPickerControl) LabelClassName(value interface{}) *IconPickerControl {
	a.Set("labelClassName", value)
	return a
}

func (a *IconPickerControl) LabelRemark(value interface{}) *IconPickerControl {
	a.Set("labelRemark", value)
	return a
}

func (a *IconPickerControl) LabelWidth(value interface{}) *IconPickerControl {
	a.Set("labelWidth", value)
	return a
}

func (a *IconPickerControl) Mode(value interface{}) *IconPickerControl {
	a.Set("mode", value)
	return a
}

func (a *IconPickerControl) Name(value interface{}) *IconPickerControl {
	a.Set("name", value)
	return a
}

func (a *IconPickerControl) OnEvent(value interface{}) *IconPickerControl {
	a.Set("onEvent", value)
	return a
}

func (a *IconPickerControl) Placeholder(value interface{}) *IconPickerControl {
	a.Set("placeholder", value)
	return a
}

func (a *IconPickerControl) ReadOnly(value interface{}) *IconPickerControl {
	a.Set("readOnly", value)
	return a
}

func (a *IconPickerControl) ReadOnlyOn(value interface{}) *IconPickerControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *IconPickerControl) Remark(value interface{}) *IconPickerControl {
	a.Set("remark", value)
	return a
}

func (a *IconPickerControl) Required(value interface{}) *IconPickerControl {
	a.Set("required", value)
	return a
}

func (a *IconPickerControl) Row(value interface{}) *IconPickerControl {
	a.Set("row", value)
	return a
}

func (a *IconPickerControl) SaveImmediately(value interface{}) *IconPickerControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *IconPickerControl) Size(value interface{}) *IconPickerControl {
	a.Set("size", value)
	return a
}

func (a *IconPickerControl) Static(value interface{}) *IconPickerControl {
	a.Set("static", value)
	return a
}

func (a *IconPickerControl) StaticClassName(value interface{}) *IconPickerControl {
	a.Set("staticClassName", value)
	return a
}

func (a *IconPickerControl) StaticInputClassName(value interface{}) *IconPickerControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *IconPickerControl) StaticLabelClassName(value interface{}) *IconPickerControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *IconPickerControl) StaticOn(value interface{}) *IconPickerControl {
	a.Set("staticOn", value)
	return a
}

func (a *IconPickerControl) StaticPlaceholder(value interface{}) *IconPickerControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *IconPickerControl) StaticSchema(value interface{}) *IconPickerControl {
	a.Set("staticSchema", value)
	return a
}

func (a *IconPickerControl) Style(value interface{}) *IconPickerControl {
	a.Set("style", value)
	return a
}

func (a *IconPickerControl) SubmitOnChange(value interface{}) *IconPickerControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *IconPickerControl) TestIdBuilder(value interface{}) *IconPickerControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *IconPickerControl) Type(value interface{}) *IconPickerControl {
	a.Set("type", value)
	return a
}

func (a *IconPickerControl) UseMobileUI(value interface{}) *IconPickerControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *IconPickerControl) ValidateApi(value interface{}) *IconPickerControl {
	a.Set("validateApi", value)
	return a
}

func (a *IconPickerControl) ValidateOnChange(value interface{}) *IconPickerControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *IconPickerControl) ValidationErrors(value interface{}) *IconPickerControl {
	a.Set("validationErrors", value)
	return a
}

func (a *IconPickerControl) Validations(value interface{}) *IconPickerControl {
	a.Set("validations", value)
	return a
}

func (a *IconPickerControl) Value(value interface{}) *IconPickerControl {
	a.Set("value", value)
	return a
}

func (a *IconPickerControl) Visible(value interface{}) *IconPickerControl {
	a.Set("visible", value)
	return a
}

func (a *IconPickerControl) VisibleOn(value interface{}) *IconPickerControl {
	a.Set("visibleOn", value)
	return a
}

func (a *IconPickerControl) Width(value interface{}) *IconPickerControl {
	a.Set("width", value)
	return a
}
