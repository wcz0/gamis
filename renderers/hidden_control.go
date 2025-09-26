package renderers

type HiddenControl struct {
	*BaseRenderer
}

func NewHiddenControl() *HiddenControl {
	a := &HiddenControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "hidden")
	return a
}

func (a *HiddenControl) Set(name string, value interface{}) *HiddenControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *HiddenControl) AutoFill(value interface{}) *HiddenControl {
	a.Set("autoFill", value)
	return a
}

func (a *HiddenControl) ClassName(value interface{}) *HiddenControl {
	a.Set("className", value)
	return a
}

func (a *HiddenControl) ClearValueOnHidden(value interface{}) *HiddenControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *HiddenControl) Desc(value interface{}) *HiddenControl {
	a.Set("desc", value)
	return a
}

func (a *HiddenControl) Description(value interface{}) *HiddenControl {
	a.Set("description", value)
	return a
}

func (a *HiddenControl) DescriptionClassName(value interface{}) *HiddenControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *HiddenControl) Disabled(value interface{}) *HiddenControl {
	a.Set("disabled", value)
	return a
}

func (a *HiddenControl) DisabledOn(value interface{}) *HiddenControl {
	a.Set("disabledOn", value)
	return a
}

func (a *HiddenControl) EditorSetting(value interface{}) *HiddenControl {
	a.Set("editorSetting", value)
	return a
}

func (a *HiddenControl) ExtraName(value interface{}) *HiddenControl {
	a.Set("extraName", value)
	return a
}

func (a *HiddenControl) Hidden(value interface{}) *HiddenControl {
	a.Set("hidden", value)
	return a
}

func (a *HiddenControl) HiddenOn(value interface{}) *HiddenControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *HiddenControl) Hint(value interface{}) *HiddenControl {
	a.Set("hint", value)
	return a
}

func (a *HiddenControl) Horizontal(value interface{}) *HiddenControl {
	a.Set("horizontal", value)
	return a
}

func (a *HiddenControl) Id(value interface{}) *HiddenControl {
	a.Set("id", value)
	return a
}

func (a *HiddenControl) InitAutoFill(value interface{}) *HiddenControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *HiddenControl) Inline(value interface{}) *HiddenControl {
	a.Set("inline", value)
	return a
}

func (a *HiddenControl) InputClassName(value interface{}) *HiddenControl {
	a.Set("inputClassName", value)
	return a
}

func (a *HiddenControl) Label(value interface{}) *HiddenControl {
	a.Set("label", value)
	return a
}

func (a *HiddenControl) LabelAlign(value interface{}) *HiddenControl {
	a.Set("labelAlign", value)
	return a
}

func (a *HiddenControl) LabelClassName(value interface{}) *HiddenControl {
	a.Set("labelClassName", value)
	return a
}

func (a *HiddenControl) LabelRemark(value interface{}) *HiddenControl {
	a.Set("labelRemark", value)
	return a
}

func (a *HiddenControl) LabelWidth(value interface{}) *HiddenControl {
	a.Set("labelWidth", value)
	return a
}

func (a *HiddenControl) Mode(value interface{}) *HiddenControl {
	a.Set("mode", value)
	return a
}

func (a *HiddenControl) Name(value interface{}) *HiddenControl {
	a.Set("name", value)
	return a
}

func (a *HiddenControl) OnEvent(value interface{}) *HiddenControl {
	a.Set("onEvent", value)
	return a
}

func (a *HiddenControl) Placeholder(value interface{}) *HiddenControl {
	a.Set("placeholder", value)
	return a
}

func (a *HiddenControl) ReadOnly(value interface{}) *HiddenControl {
	a.Set("readOnly", value)
	return a
}

func (a *HiddenControl) ReadOnlyOn(value interface{}) *HiddenControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *HiddenControl) Remark(value interface{}) *HiddenControl {
	a.Set("remark", value)
	return a
}

func (a *HiddenControl) Required(value interface{}) *HiddenControl {
	a.Set("required", value)
	return a
}

func (a *HiddenControl) Row(value interface{}) *HiddenControl {
	a.Set("row", value)
	return a
}

func (a *HiddenControl) SaveImmediately(value interface{}) *HiddenControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *HiddenControl) Size(value interface{}) *HiddenControl {
	a.Set("size", value)
	return a
}

func (a *HiddenControl) Static(value interface{}) *HiddenControl {
	a.Set("static", value)
	return a
}

func (a *HiddenControl) StaticClassName(value interface{}) *HiddenControl {
	a.Set("staticClassName", value)
	return a
}

func (a *HiddenControl) StaticInputClassName(value interface{}) *HiddenControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *HiddenControl) StaticLabelClassName(value interface{}) *HiddenControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *HiddenControl) StaticOn(value interface{}) *HiddenControl {
	a.Set("staticOn", value)
	return a
}

func (a *HiddenControl) StaticPlaceholder(value interface{}) *HiddenControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *HiddenControl) StaticSchema(value interface{}) *HiddenControl {
	a.Set("staticSchema", value)
	return a
}

func (a *HiddenControl) Style(value interface{}) *HiddenControl {
	a.Set("style", value)
	return a
}

func (a *HiddenControl) SubmitOnChange(value interface{}) *HiddenControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *HiddenControl) TestIdBuilder(value interface{}) *HiddenControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *HiddenControl) Type(value interface{}) *HiddenControl {
	a.Set("type", value)
	return a
}

func (a *HiddenControl) UseMobileUI(value interface{}) *HiddenControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *HiddenControl) ValidateApi(value interface{}) *HiddenControl {
	a.Set("validateApi", value)
	return a
}

func (a *HiddenControl) ValidateOnChange(value interface{}) *HiddenControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *HiddenControl) ValidationErrors(value interface{}) *HiddenControl {
	a.Set("validationErrors", value)
	return a
}

func (a *HiddenControl) Validations(value interface{}) *HiddenControl {
	a.Set("validations", value)
	return a
}

func (a *HiddenControl) Value(value interface{}) *HiddenControl {
	a.Set("value", value)
	return a
}

func (a *HiddenControl) Visible(value interface{}) *HiddenControl {
	a.Set("visible", value)
	return a
}

func (a *HiddenControl) VisibleOn(value interface{}) *HiddenControl {
	a.Set("visibleOn", value)
	return a
}

func (a *HiddenControl) Width(value interface{}) *HiddenControl {
	a.Set("width", value)
	return a
}
