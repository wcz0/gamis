package renderers

type InputSignature struct {
	*BaseRenderer
}

func NewInputSignature() *InputSignature {
	a := &InputSignature{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "input-signature")
	return a
}

func (a *InputSignature) Set(name string, value interface{}) *InputSignature {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *InputSignature) AutoFill(value interface{}) *InputSignature {
	a.Set("autoFill", value)
	return a
}

func (a *InputSignature) BgColor(value interface{}) *InputSignature {
	a.Set("bgColor", value)
	return a
}

func (a *InputSignature) ClassName(value interface{}) *InputSignature {
	a.Set("className", value)
	return a
}

func (a *InputSignature) ClearBtnIcon(value interface{}) *InputSignature {
	a.Set("clearBtnIcon", value)
	return a
}

func (a *InputSignature) ClearBtnLabel(value interface{}) *InputSignature {
	a.Set("clearBtnLabel", value)
	return a
}

func (a *InputSignature) ClearValueOnHidden(value interface{}) *InputSignature {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *InputSignature) Color(value interface{}) *InputSignature {
	a.Set("color", value)
	return a
}

func (a *InputSignature) ConfirmBtnIcon(value interface{}) *InputSignature {
	a.Set("confirmBtnIcon", value)
	return a
}

func (a *InputSignature) ConfirmBtnLabel(value interface{}) *InputSignature {
	a.Set("confirmBtnLabel", value)
	return a
}

func (a *InputSignature) Desc(value interface{}) *InputSignature {
	a.Set("desc", value)
	return a
}

func (a *InputSignature) Description(value interface{}) *InputSignature {
	a.Set("description", value)
	return a
}

func (a *InputSignature) DescriptionClassName(value interface{}) *InputSignature {
	a.Set("descriptionClassName", value)
	return a
}

func (a *InputSignature) Disabled(value interface{}) *InputSignature {
	a.Set("disabled", value)
	return a
}

func (a *InputSignature) DisabledOn(value interface{}) *InputSignature {
	a.Set("disabledOn", value)
	return a
}

func (a *InputSignature) EbmedCancelIcon(value interface{}) *InputSignature {
	a.Set("ebmedCancelIcon", value)
	return a
}

func (a *InputSignature) EbmedCancelLabel(value interface{}) *InputSignature {
	a.Set("ebmedCancelLabel", value)
	return a
}

func (a *InputSignature) EditorSetting(value interface{}) *InputSignature {
	a.Set("editorSetting", value)
	return a
}

func (a *InputSignature) Embed(value interface{}) *InputSignature {
	a.Set("embed", value)
	return a
}

func (a *InputSignature) EmbedBtnIcon(value interface{}) *InputSignature {
	a.Set("embedBtnIcon", value)
	return a
}

func (a *InputSignature) EmbedBtnLabel(value interface{}) *InputSignature {
	a.Set("embedBtnLabel", value)
	return a
}

func (a *InputSignature) EmbedConfirmIcon(value interface{}) *InputSignature {
	a.Set("embedConfirmIcon", value)
	return a
}

func (a *InputSignature) EmbedConfirmLabel(value interface{}) *InputSignature {
	a.Set("embedConfirmLabel", value)
	return a
}

func (a *InputSignature) ExtraName(value interface{}) *InputSignature {
	a.Set("extraName", value)
	return a
}

func (a *InputSignature) Height(value interface{}) *InputSignature {
	a.Set("height", value)
	return a
}

func (a *InputSignature) Hidden(value interface{}) *InputSignature {
	a.Set("hidden", value)
	return a
}

func (a *InputSignature) HiddenOn(value interface{}) *InputSignature {
	a.Set("hiddenOn", value)
	return a
}

func (a *InputSignature) Hint(value interface{}) *InputSignature {
	a.Set("hint", value)
	return a
}

func (a *InputSignature) Horizontal(value interface{}) *InputSignature {
	a.Set("horizontal", value)
	return a
}

func (a *InputSignature) Id(value interface{}) *InputSignature {
	a.Set("id", value)
	return a
}

func (a *InputSignature) InitAutoFill(value interface{}) *InputSignature {
	a.Set("initAutoFill", value)
	return a
}

func (a *InputSignature) Inline(value interface{}) *InputSignature {
	a.Set("inline", value)
	return a
}

func (a *InputSignature) InputClassName(value interface{}) *InputSignature {
	a.Set("inputClassName", value)
	return a
}

func (a *InputSignature) Label(value interface{}) *InputSignature {
	a.Set("label", value)
	return a
}

func (a *InputSignature) LabelAlign(value interface{}) *InputSignature {
	a.Set("labelAlign", value)
	return a
}

func (a *InputSignature) LabelClassName(value interface{}) *InputSignature {
	a.Set("labelClassName", value)
	return a
}

func (a *InputSignature) LabelRemark(value interface{}) *InputSignature {
	a.Set("labelRemark", value)
	return a
}

func (a *InputSignature) LabelWidth(value interface{}) *InputSignature {
	a.Set("labelWidth", value)
	return a
}

func (a *InputSignature) Mode(value interface{}) *InputSignature {
	a.Set("mode", value)
	return a
}

func (a *InputSignature) Name(value interface{}) *InputSignature {
	a.Set("name", value)
	return a
}

func (a *InputSignature) OnEvent(value interface{}) *InputSignature {
	a.Set("onEvent", value)
	return a
}

func (a *InputSignature) Placeholder(value interface{}) *InputSignature {
	a.Set("placeholder", value)
	return a
}

func (a *InputSignature) ReadOnly(value interface{}) *InputSignature {
	a.Set("readOnly", value)
	return a
}

func (a *InputSignature) ReadOnlyOn(value interface{}) *InputSignature {
	a.Set("readOnlyOn", value)
	return a
}

func (a *InputSignature) Remark(value interface{}) *InputSignature {
	a.Set("remark", value)
	return a
}

func (a *InputSignature) Required(value interface{}) *InputSignature {
	a.Set("required", value)
	return a
}

func (a *InputSignature) Row(value interface{}) *InputSignature {
	a.Set("row", value)
	return a
}

func (a *InputSignature) Size(value interface{}) *InputSignature {
	a.Set("size", value)
	return a
}

func (a *InputSignature) Static(value interface{}) *InputSignature {
	a.Set("static", value)
	return a
}

func (a *InputSignature) StaticClassName(value interface{}) *InputSignature {
	a.Set("staticClassName", value)
	return a
}

func (a *InputSignature) StaticInputClassName(value interface{}) *InputSignature {
	a.Set("staticInputClassName", value)
	return a
}

func (a *InputSignature) StaticLabelClassName(value interface{}) *InputSignature {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *InputSignature) StaticOn(value interface{}) *InputSignature {
	a.Set("staticOn", value)
	return a
}

func (a *InputSignature) StaticPlaceholder(value interface{}) *InputSignature {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *InputSignature) StaticSchema(value interface{}) *InputSignature {
	a.Set("staticSchema", value)
	return a
}

func (a *InputSignature) Style(value interface{}) *InputSignature {
	a.Set("style", value)
	return a
}

func (a *InputSignature) SubmitOnChange(value interface{}) *InputSignature {
	a.Set("submitOnChange", value)
	return a
}

func (a *InputSignature) TestIdBuilder(value interface{}) *InputSignature {
	a.Set("testIdBuilder", value)
	return a
}

func (a *InputSignature) Type(value interface{}) *InputSignature {
	a.Set("type", value)
	return a
}

func (a *InputSignature) UndoBtnIcon(value interface{}) *InputSignature {
	a.Set("undoBtnIcon", value)
	return a
}

func (a *InputSignature) UndoBtnLabel(value interface{}) *InputSignature {
	a.Set("undoBtnLabel", value)
	return a
}

func (a *InputSignature) UploadApi(value interface{}) *InputSignature {
	a.Set("uploadApi", value)
	return a
}

func (a *InputSignature) UseMobileUI(value interface{}) *InputSignature {
	a.Set("useMobileUI", value)
	return a
}

func (a *InputSignature) ValidateApi(value interface{}) *InputSignature {
	a.Set("validateApi", value)
	return a
}

func (a *InputSignature) ValidateOnChange(value interface{}) *InputSignature {
	a.Set("validateOnChange", value)
	return a
}

func (a *InputSignature) ValidationErrors(value interface{}) *InputSignature {
	a.Set("validationErrors", value)
	return a
}

func (a *InputSignature) Validations(value interface{}) *InputSignature {
	a.Set("validations", value)
	return a
}

func (a *InputSignature) Value(value interface{}) *InputSignature {
	a.Set("value", value)
	return a
}

func (a *InputSignature) Visible(value interface{}) *InputSignature {
	a.Set("visible", value)
	return a
}

func (a *InputSignature) VisibleOn(value interface{}) *InputSignature {
	a.Set("visibleOn", value)
	return a
}

func (a *InputSignature) Width(value interface{}) *InputSignature {
	a.Set("width", value)
	return a
}
