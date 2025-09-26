package renderers

type RichTextControl struct {
	*BaseRenderer
}

func NewRichTextControl() *RichTextControl {
	a := &RichTextControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "input-rich-text")
	return a
}

func (a *RichTextControl) Set(name string, value interface{}) *RichTextControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *RichTextControl) AutoFill(value interface{}) *RichTextControl {
	a.Set("autoFill", value)
	return a
}

func (a *RichTextControl) BorderMode(value interface{}) *RichTextControl {
	a.Set("borderMode", value)
	return a
}

func (a *RichTextControl) ClassName(value interface{}) *RichTextControl {
	a.Set("className", value)
	return a
}

func (a *RichTextControl) ClearValueOnHidden(value interface{}) *RichTextControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *RichTextControl) Desc(value interface{}) *RichTextControl {
	a.Set("desc", value)
	return a
}

func (a *RichTextControl) Description(value interface{}) *RichTextControl {
	a.Set("description", value)
	return a
}

func (a *RichTextControl) DescriptionClassName(value interface{}) *RichTextControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *RichTextControl) Disabled(value interface{}) *RichTextControl {
	a.Set("disabled", value)
	return a
}

func (a *RichTextControl) DisabledOn(value interface{}) *RichTextControl {
	a.Set("disabledOn", value)
	return a
}

func (a *RichTextControl) EditorSetting(value interface{}) *RichTextControl {
	a.Set("editorSetting", value)
	return a
}

func (a *RichTextControl) ExtraName(value interface{}) *RichTextControl {
	a.Set("extraName", value)
	return a
}

func (a *RichTextControl) FileField(value interface{}) *RichTextControl {
	a.Set("fileField", value)
	return a
}

func (a *RichTextControl) Hidden(value interface{}) *RichTextControl {
	a.Set("hidden", value)
	return a
}

func (a *RichTextControl) HiddenOn(value interface{}) *RichTextControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *RichTextControl) Hint(value interface{}) *RichTextControl {
	a.Set("hint", value)
	return a
}

func (a *RichTextControl) Horizontal(value interface{}) *RichTextControl {
	a.Set("horizontal", value)
	return a
}

func (a *RichTextControl) Id(value interface{}) *RichTextControl {
	a.Set("id", value)
	return a
}

func (a *RichTextControl) InitAutoFill(value interface{}) *RichTextControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *RichTextControl) Inline(value interface{}) *RichTextControl {
	a.Set("inline", value)
	return a
}

func (a *RichTextControl) InputClassName(value interface{}) *RichTextControl {
	a.Set("inputClassName", value)
	return a
}

func (a *RichTextControl) Label(value interface{}) *RichTextControl {
	a.Set("label", value)
	return a
}

func (a *RichTextControl) LabelAlign(value interface{}) *RichTextControl {
	a.Set("labelAlign", value)
	return a
}

func (a *RichTextControl) LabelClassName(value interface{}) *RichTextControl {
	a.Set("labelClassName", value)
	return a
}

func (a *RichTextControl) LabelRemark(value interface{}) *RichTextControl {
	a.Set("labelRemark", value)
	return a
}

func (a *RichTextControl) LabelWidth(value interface{}) *RichTextControl {
	a.Set("labelWidth", value)
	return a
}

func (a *RichTextControl) Mode(value interface{}) *RichTextControl {
	a.Set("mode", value)
	return a
}

func (a *RichTextControl) Name(value interface{}) *RichTextControl {
	a.Set("name", value)
	return a
}

func (a *RichTextControl) OnEvent(value interface{}) *RichTextControl {
	a.Set("onEvent", value)
	return a
}

func (a *RichTextControl) Options(value interface{}) *RichTextControl {
	a.Set("options", value)
	return a
}

func (a *RichTextControl) Placeholder(value interface{}) *RichTextControl {
	a.Set("placeholder", value)
	return a
}

func (a *RichTextControl) ReadOnly(value interface{}) *RichTextControl {
	a.Set("readOnly", value)
	return a
}

func (a *RichTextControl) ReadOnlyOn(value interface{}) *RichTextControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *RichTextControl) Receiver(value interface{}) *RichTextControl {
	a.Set("receiver", value)
	return a
}

func (a *RichTextControl) Remark(value interface{}) *RichTextControl {
	a.Set("remark", value)
	return a
}

func (a *RichTextControl) Required(value interface{}) *RichTextControl {
	a.Set("required", value)
	return a
}

func (a *RichTextControl) Row(value interface{}) *RichTextControl {
	a.Set("row", value)
	return a
}

func (a *RichTextControl) SaveImmediately(value interface{}) *RichTextControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *RichTextControl) Size(value interface{}) *RichTextControl {
	a.Set("size", value)
	return a
}

func (a *RichTextControl) Static(value interface{}) *RichTextControl {
	a.Set("static", value)
	return a
}

func (a *RichTextControl) StaticClassName(value interface{}) *RichTextControl {
	a.Set("staticClassName", value)
	return a
}

func (a *RichTextControl) StaticInputClassName(value interface{}) *RichTextControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *RichTextControl) StaticLabelClassName(value interface{}) *RichTextControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *RichTextControl) StaticOn(value interface{}) *RichTextControl {
	a.Set("staticOn", value)
	return a
}

func (a *RichTextControl) StaticPlaceholder(value interface{}) *RichTextControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *RichTextControl) StaticSchema(value interface{}) *RichTextControl {
	a.Set("staticSchema", value)
	return a
}

func (a *RichTextControl) Style(value interface{}) *RichTextControl {
	a.Set("style", value)
	return a
}

func (a *RichTextControl) SubmitOnChange(value interface{}) *RichTextControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *RichTextControl) TestIdBuilder(value interface{}) *RichTextControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *RichTextControl) Type(value interface{}) *RichTextControl {
	a.Set("type", value)
	return a
}

func (a *RichTextControl) UseMobileUI(value interface{}) *RichTextControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *RichTextControl) ValidateApi(value interface{}) *RichTextControl {
	a.Set("validateApi", value)
	return a
}

func (a *RichTextControl) ValidateOnChange(value interface{}) *RichTextControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *RichTextControl) ValidationErrors(value interface{}) *RichTextControl {
	a.Set("validationErrors", value)
	return a
}

func (a *RichTextControl) Validations(value interface{}) *RichTextControl {
	a.Set("validations", value)
	return a
}

func (a *RichTextControl) Value(value interface{}) *RichTextControl {
	a.Set("value", value)
	return a
}

func (a *RichTextControl) Vendor(value interface{}) *RichTextControl {
	a.Set("vendor", value)
	return a
}

func (a *RichTextControl) VideoReceiver(value interface{}) *RichTextControl {
	a.Set("videoReceiver", value)
	return a
}

func (a *RichTextControl) Visible(value interface{}) *RichTextControl {
	a.Set("visible", value)
	return a
}

func (a *RichTextControl) VisibleOn(value interface{}) *RichTextControl {
	a.Set("visibleOn", value)
	return a
}

func (a *RichTextControl) Width(value interface{}) *RichTextControl {
	a.Set("width", value)
	return a
}
