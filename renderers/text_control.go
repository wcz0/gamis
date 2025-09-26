package renderers

type TextControl struct {
	*BaseRenderer
}

func NewTextControl() *TextControl {
	a := &TextControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "input-text")
	return a
}

func (a *TextControl) Set(name string, value interface{}) *TextControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *TextControl) AddApi(value interface{}) *TextControl {
	a.Set("addApi", value)
	return a
}

func (a *TextControl) AddControls(value interface{}) *TextControl {
	a.Set("addControls", value)
	return a
}

func (a *TextControl) AddDialog(value interface{}) *TextControl {
	a.Set("addDialog", value)
	return a
}

func (a *TextControl) AddOn(value interface{}) *TextControl {
	a.Set("addOn", value)
	return a
}

func (a *TextControl) AutoComplete(value interface{}) *TextControl {
	a.Set("autoComplete", value)
	return a
}

func (a *TextControl) AutoFill(value interface{}) *TextControl {
	a.Set("autoFill", value)
	return a
}

func (a *TextControl) BorderMode(value interface{}) *TextControl {
	a.Set("borderMode", value)
	return a
}

func (a *TextControl) ClassName(value interface{}) *TextControl {
	a.Set("className", value)
	return a
}

func (a *TextControl) ClearValueOnEmpty(value interface{}) *TextControl {
	a.Set("clearValueOnEmpty", value)
	return a
}

func (a *TextControl) ClearValueOnHidden(value interface{}) *TextControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *TextControl) ClearValueOnSourceChange(value interface{}) *TextControl {
	a.Set("clearValueOnSourceChange", value)
	return a
}

func (a *TextControl) Clearable(value interface{}) *TextControl {
	a.Set("clearable", value)
	return a
}

func (a *TextControl) Creatable(value interface{}) *TextControl {
	a.Set("creatable", value)
	return a
}

func (a *TextControl) CreateBtnLabel(value interface{}) *TextControl {
	a.Set("createBtnLabel", value)
	return a
}

func (a *TextControl) DeferApi(value interface{}) *TextControl {
	a.Set("deferApi", value)
	return a
}

func (a *TextControl) DeferField(value interface{}) *TextControl {
	a.Set("deferField", value)
	return a
}

func (a *TextControl) DeleteApi(value interface{}) *TextControl {
	a.Set("deleteApi", value)
	return a
}

func (a *TextControl) DeleteConfirmText(value interface{}) *TextControl {
	a.Set("deleteConfirmText", value)
	return a
}

func (a *TextControl) Delimiter(value interface{}) *TextControl {
	a.Set("delimiter", value)
	return a
}

func (a *TextControl) Desc(value interface{}) *TextControl {
	a.Set("desc", value)
	return a
}

func (a *TextControl) Description(value interface{}) *TextControl {
	a.Set("description", value)
	return a
}

func (a *TextControl) DescriptionClassName(value interface{}) *TextControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *TextControl) Disabled(value interface{}) *TextControl {
	a.Set("disabled", value)
	return a
}

func (a *TextControl) DisabledOn(value interface{}) *TextControl {
	a.Set("disabledOn", value)
	return a
}

func (a *TextControl) EditApi(value interface{}) *TextControl {
	a.Set("editApi", value)
	return a
}

func (a *TextControl) EditControls(value interface{}) *TextControl {
	a.Set("editControls", value)
	return a
}

func (a *TextControl) EditDialog(value interface{}) *TextControl {
	a.Set("editDialog", value)
	return a
}

func (a *TextControl) Editable(value interface{}) *TextControl {
	a.Set("editable", value)
	return a
}

func (a *TextControl) EditorSetting(value interface{}) *TextControl {
	a.Set("editorSetting", value)
	return a
}

func (a *TextControl) ExtraName(value interface{}) *TextControl {
	a.Set("extraName", value)
	return a
}

func (a *TextControl) ExtractValue(value interface{}) *TextControl {
	a.Set("extractValue", value)
	return a
}

func (a *TextControl) Hidden(value interface{}) *TextControl {
	a.Set("hidden", value)
	return a
}

func (a *TextControl) HiddenOn(value interface{}) *TextControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *TextControl) Hint(value interface{}) *TextControl {
	a.Set("hint", value)
	return a
}

func (a *TextControl) Horizontal(value interface{}) *TextControl {
	a.Set("horizontal", value)
	return a
}

func (a *TextControl) Id(value interface{}) *TextControl {
	a.Set("id", value)
	return a
}

func (a *TextControl) InitAutoFill(value interface{}) *TextControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *TextControl) InitFetch(value interface{}) *TextControl {
	a.Set("initFetch", value)
	return a
}

func (a *TextControl) InitFetchOn(value interface{}) *TextControl {
	a.Set("initFetchOn", value)
	return a
}

func (a *TextControl) Inline(value interface{}) *TextControl {
	a.Set("inline", value)
	return a
}

func (a *TextControl) InputClassName(value interface{}) *TextControl {
	a.Set("inputClassName", value)
	return a
}

func (a *TextControl) InputControlClassName(value interface{}) *TextControl {
	a.Set("inputControlClassName", value)
	return a
}

func (a *TextControl) JoinValues(value interface{}) *TextControl {
	a.Set("joinValues", value)
	return a
}

func (a *TextControl) Label(value interface{}) *TextControl {
	a.Set("label", value)
	return a
}

func (a *TextControl) LabelAlign(value interface{}) *TextControl {
	a.Set("labelAlign", value)
	return a
}

func (a *TextControl) LabelClassName(value interface{}) *TextControl {
	a.Set("labelClassName", value)
	return a
}

func (a *TextControl) LabelRemark(value interface{}) *TextControl {
	a.Set("labelRemark", value)
	return a
}

func (a *TextControl) LabelWidth(value interface{}) *TextControl {
	a.Set("labelWidth", value)
	return a
}

func (a *TextControl) MaxLength(value interface{}) *TextControl {
	a.Set("maxLength", value)
	return a
}

func (a *TextControl) MinLength(value interface{}) *TextControl {
	a.Set("minLength", value)
	return a
}

func (a *TextControl) Mode(value interface{}) *TextControl {
	a.Set("mode", value)
	return a
}

func (a *TextControl) Multiple(value interface{}) *TextControl {
	a.Set("multiple", value)
	return a
}

func (a *TextControl) Name(value interface{}) *TextControl {
	a.Set("name", value)
	return a
}

func (a *TextControl) NativeAutoComplete(value interface{}) *TextControl {
	a.Set("nativeAutoComplete", value)
	return a
}

func (a *TextControl) NativeInputClassName(value interface{}) *TextControl {
	a.Set("nativeInputClassName", value)
	return a
}

func (a *TextControl) OnEvent(value interface{}) *TextControl {
	a.Set("onEvent", value)
	return a
}

func (a *TextControl) Options(value interface{}) *TextControl {
	a.Set("options", value)
	return a
}

func (a *TextControl) Placeholder(value interface{}) *TextControl {
	a.Set("placeholder", value)
	return a
}

func (a *TextControl) Prefix(value interface{}) *TextControl {
	a.Set("prefix", value)
	return a
}

func (a *TextControl) ReadOnly(value interface{}) *TextControl {
	a.Set("readOnly", value)
	return a
}

func (a *TextControl) ReadOnlyOn(value interface{}) *TextControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *TextControl) Remark(value interface{}) *TextControl {
	a.Set("remark", value)
	return a
}

func (a *TextControl) Removable(value interface{}) *TextControl {
	a.Set("removable", value)
	return a
}

func (a *TextControl) Required(value interface{}) *TextControl {
	a.Set("required", value)
	return a
}

func (a *TextControl) ResetValue(value interface{}) *TextControl {
	a.Set("resetValue", value)
	return a
}

func (a *TextControl) Row(value interface{}) *TextControl {
	a.Set("row", value)
	return a
}

func (a *TextControl) SaveImmediately(value interface{}) *TextControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *TextControl) SelectFirst(value interface{}) *TextControl {
	a.Set("selectFirst", value)
	return a
}

func (a *TextControl) ShowCounter(value interface{}) *TextControl {
	a.Set("showCounter", value)
	return a
}

func (a *TextControl) Size(value interface{}) *TextControl {
	a.Set("size", value)
	return a
}

func (a *TextControl) Source(value interface{}) *TextControl {
	a.Set("source", value)
	return a
}

func (a *TextControl) Static(value interface{}) *TextControl {
	a.Set("static", value)
	return a
}

func (a *TextControl) StaticClassName(value interface{}) *TextControl {
	a.Set("staticClassName", value)
	return a
}

func (a *TextControl) StaticInputClassName(value interface{}) *TextControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *TextControl) StaticLabelClassName(value interface{}) *TextControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *TextControl) StaticOn(value interface{}) *TextControl {
	a.Set("staticOn", value)
	return a
}

func (a *TextControl) StaticPlaceholder(value interface{}) *TextControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *TextControl) StaticSchema(value interface{}) *TextControl {
	a.Set("staticSchema", value)
	return a
}

func (a *TextControl) Style(value interface{}) *TextControl {
	a.Set("style", value)
	return a
}

func (a *TextControl) SubmitOnChange(value interface{}) *TextControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *TextControl) Suffix(value interface{}) *TextControl {
	a.Set("suffix", value)
	return a
}

func (a *TextControl) TestIdBuilder(value interface{}) *TextControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *TextControl) Transform(value interface{}) *TextControl {
	a.Set("transform", value)
	return a
}

func (a *TextControl) TrimContents(value interface{}) *TextControl {
	a.Set("trimContents", value)
	return a
}

func (a *TextControl) Type(value interface{}) *TextControl {
	a.Set("type", value)
	return a
}

func (a *TextControl) UseMobileUI(value interface{}) *TextControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *TextControl) ValidateApi(value interface{}) *TextControl {
	a.Set("validateApi", value)
	return a
}

func (a *TextControl) ValidateOnChange(value interface{}) *TextControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *TextControl) ValidationErrors(value interface{}) *TextControl {
	a.Set("validationErrors", value)
	return a
}

func (a *TextControl) Validations(value interface{}) *TextControl {
	a.Set("validations", value)
	return a
}

func (a *TextControl) Value(value interface{}) *TextControl {
	a.Set("value", value)
	return a
}

func (a *TextControl) ValuesNoWrap(value interface{}) *TextControl {
	a.Set("valuesNoWrap", value)
	return a
}

func (a *TextControl) Visible(value interface{}) *TextControl {
	a.Set("visible", value)
	return a
}

func (a *TextControl) VisibleOn(value interface{}) *TextControl {
	a.Set("visibleOn", value)
	return a
}

func (a *TextControl) Width(value interface{}) *TextControl {
	a.Set("width", value)
	return a
}
