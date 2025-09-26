package renderers

type TextareaControl struct {
	*BaseRenderer
}

func NewTextareaControl() *TextareaControl {
	a := &TextareaControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "textarea")
	return a
}

func (a *TextareaControl) Set(name string, value interface{}) *TextareaControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *TextareaControl) AutoFill(value interface{}) *TextareaControl {
	a.Set("autoFill", value)
	return a
}

func (a *TextareaControl) BorderMode(value interface{}) *TextareaControl {
	a.Set("borderMode", value)
	return a
}

func (a *TextareaControl) ClassName(value interface{}) *TextareaControl {
	a.Set("className", value)
	return a
}

func (a *TextareaControl) ClearValueOnHidden(value interface{}) *TextareaControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *TextareaControl) Clearable(value interface{}) *TextareaControl {
	a.Set("clearable", value)
	return a
}

func (a *TextareaControl) Desc(value interface{}) *TextareaControl {
	a.Set("desc", value)
	return a
}

func (a *TextareaControl) Description(value interface{}) *TextareaControl {
	a.Set("description", value)
	return a
}

func (a *TextareaControl) DescriptionClassName(value interface{}) *TextareaControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *TextareaControl) Disabled(value interface{}) *TextareaControl {
	a.Set("disabled", value)
	return a
}

func (a *TextareaControl) DisabledOn(value interface{}) *TextareaControl {
	a.Set("disabledOn", value)
	return a
}

func (a *TextareaControl) EditorSetting(value interface{}) *TextareaControl {
	a.Set("editorSetting", value)
	return a
}

func (a *TextareaControl) ExtraName(value interface{}) *TextareaControl {
	a.Set("extraName", value)
	return a
}

func (a *TextareaControl) Hidden(value interface{}) *TextareaControl {
	a.Set("hidden", value)
	return a
}

func (a *TextareaControl) HiddenOn(value interface{}) *TextareaControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *TextareaControl) Hint(value interface{}) *TextareaControl {
	a.Set("hint", value)
	return a
}

func (a *TextareaControl) Horizontal(value interface{}) *TextareaControl {
	a.Set("horizontal", value)
	return a
}

func (a *TextareaControl) Id(value interface{}) *TextareaControl {
	a.Set("id", value)
	return a
}

func (a *TextareaControl) InitAutoFill(value interface{}) *TextareaControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *TextareaControl) Inline(value interface{}) *TextareaControl {
	a.Set("inline", value)
	return a
}

func (a *TextareaControl) InputClassName(value interface{}) *TextareaControl {
	a.Set("inputClassName", value)
	return a
}

func (a *TextareaControl) Label(value interface{}) *TextareaControl {
	a.Set("label", value)
	return a
}

func (a *TextareaControl) LabelAlign(value interface{}) *TextareaControl {
	a.Set("labelAlign", value)
	return a
}

func (a *TextareaControl) LabelClassName(value interface{}) *TextareaControl {
	a.Set("labelClassName", value)
	return a
}

func (a *TextareaControl) LabelRemark(value interface{}) *TextareaControl {
	a.Set("labelRemark", value)
	return a
}

func (a *TextareaControl) LabelWidth(value interface{}) *TextareaControl {
	a.Set("labelWidth", value)
	return a
}

func (a *TextareaControl) MaxLength(value interface{}) *TextareaControl {
	a.Set("maxLength", value)
	return a
}

func (a *TextareaControl) MaxRows(value interface{}) *TextareaControl {
	a.Set("maxRows", value)
	return a
}

func (a *TextareaControl) MinRows(value interface{}) *TextareaControl {
	a.Set("minRows", value)
	return a
}

func (a *TextareaControl) Mode(value interface{}) *TextareaControl {
	a.Set("mode", value)
	return a
}

func (a *TextareaControl) Name(value interface{}) *TextareaControl {
	a.Set("name", value)
	return a
}

func (a *TextareaControl) OnEvent(value interface{}) *TextareaControl {
	a.Set("onEvent", value)
	return a
}

func (a *TextareaControl) Placeholder(value interface{}) *TextareaControl {
	a.Set("placeholder", value)
	return a
}

func (a *TextareaControl) ReadOnly(value interface{}) *TextareaControl {
	a.Set("readOnly", value)
	return a
}

func (a *TextareaControl) ReadOnlyOn(value interface{}) *TextareaControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *TextareaControl) Remark(value interface{}) *TextareaControl {
	a.Set("remark", value)
	return a
}

func (a *TextareaControl) Required(value interface{}) *TextareaControl {
	a.Set("required", value)
	return a
}

func (a *TextareaControl) ResetValue(value interface{}) *TextareaControl {
	a.Set("resetValue", value)
	return a
}

func (a *TextareaControl) Row(value interface{}) *TextareaControl {
	a.Set("row", value)
	return a
}

func (a *TextareaControl) SaveImmediately(value interface{}) *TextareaControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *TextareaControl) ShowCounter(value interface{}) *TextareaControl {
	a.Set("showCounter", value)
	return a
}

func (a *TextareaControl) Size(value interface{}) *TextareaControl {
	a.Set("size", value)
	return a
}

func (a *TextareaControl) Static(value interface{}) *TextareaControl {
	a.Set("static", value)
	return a
}

func (a *TextareaControl) StaticClassName(value interface{}) *TextareaControl {
	a.Set("staticClassName", value)
	return a
}

func (a *TextareaControl) StaticInputClassName(value interface{}) *TextareaControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *TextareaControl) StaticLabelClassName(value interface{}) *TextareaControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *TextareaControl) StaticOn(value interface{}) *TextareaControl {
	a.Set("staticOn", value)
	return a
}

func (a *TextareaControl) StaticPlaceholder(value interface{}) *TextareaControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *TextareaControl) StaticSchema(value interface{}) *TextareaControl {
	a.Set("staticSchema", value)
	return a
}

func (a *TextareaControl) Style(value interface{}) *TextareaControl {
	a.Set("style", value)
	return a
}

func (a *TextareaControl) SubmitOnChange(value interface{}) *TextareaControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *TextareaControl) TestIdBuilder(value interface{}) *TextareaControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *TextareaControl) Type(value interface{}) *TextareaControl {
	a.Set("type", value)
	return a
}

func (a *TextareaControl) UseMobileUI(value interface{}) *TextareaControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *TextareaControl) ValidateApi(value interface{}) *TextareaControl {
	a.Set("validateApi", value)
	return a
}

func (a *TextareaControl) ValidateOnChange(value interface{}) *TextareaControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *TextareaControl) ValidationErrors(value interface{}) *TextareaControl {
	a.Set("validationErrors", value)
	return a
}

func (a *TextareaControl) Validations(value interface{}) *TextareaControl {
	a.Set("validations", value)
	return a
}

func (a *TextareaControl) Value(value interface{}) *TextareaControl {
	a.Set("value", value)
	return a
}

func (a *TextareaControl) Visible(value interface{}) *TextareaControl {
	a.Set("visible", value)
	return a
}

func (a *TextareaControl) VisibleOn(value interface{}) *TextareaControl {
	a.Set("visibleOn", value)
	return a
}

func (a *TextareaControl) Width(value interface{}) *TextareaControl {
	a.Set("width", value)
	return a
}
