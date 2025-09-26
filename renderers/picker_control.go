package renderers

type PickerControl struct {
	*BaseRenderer
}

func NewPickerControl() *PickerControl {
	a := &PickerControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "picker")
	return a
}

func (a *PickerControl) Set(name string, value interface{}) *PickerControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *PickerControl) AddApi(value interface{}) *PickerControl {
	a.Set("addApi", value)
	return a
}

func (a *PickerControl) AddControls(value interface{}) *PickerControl {
	a.Set("addControls", value)
	return a
}

func (a *PickerControl) AddDialog(value interface{}) *PickerControl {
	a.Set("addDialog", value)
	return a
}

func (a *PickerControl) AutoFill(value interface{}) *PickerControl {
	a.Set("autoFill", value)
	return a
}

func (a *PickerControl) ClassName(value interface{}) *PickerControl {
	a.Set("className", value)
	return a
}

func (a *PickerControl) ClearValueOnHidden(value interface{}) *PickerControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *PickerControl) ClearValueOnSourceChange(value interface{}) *PickerControl {
	a.Set("clearValueOnSourceChange", value)
	return a
}

func (a *PickerControl) Clearable(value interface{}) *PickerControl {
	a.Set("clearable", value)
	return a
}

func (a *PickerControl) Creatable(value interface{}) *PickerControl {
	a.Set("creatable", value)
	return a
}

func (a *PickerControl) CreateBtnLabel(value interface{}) *PickerControl {
	a.Set("createBtnLabel", value)
	return a
}

func (a *PickerControl) DeferApi(value interface{}) *PickerControl {
	a.Set("deferApi", value)
	return a
}

func (a *PickerControl) DeferField(value interface{}) *PickerControl {
	a.Set("deferField", value)
	return a
}

func (a *PickerControl) DeleteApi(value interface{}) *PickerControl {
	a.Set("deleteApi", value)
	return a
}

func (a *PickerControl) DeleteConfirmText(value interface{}) *PickerControl {
	a.Set("deleteConfirmText", value)
	return a
}

func (a *PickerControl) Delimiter(value interface{}) *PickerControl {
	a.Set("delimiter", value)
	return a
}

func (a *PickerControl) Desc(value interface{}) *PickerControl {
	a.Set("desc", value)
	return a
}

func (a *PickerControl) Description(value interface{}) *PickerControl {
	a.Set("description", value)
	return a
}

func (a *PickerControl) DescriptionClassName(value interface{}) *PickerControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *PickerControl) Disabled(value interface{}) *PickerControl {
	a.Set("disabled", value)
	return a
}

func (a *PickerControl) DisabledOn(value interface{}) *PickerControl {
	a.Set("disabledOn", value)
	return a
}

func (a *PickerControl) EditApi(value interface{}) *PickerControl {
	a.Set("editApi", value)
	return a
}

func (a *PickerControl) EditControls(value interface{}) *PickerControl {
	a.Set("editControls", value)
	return a
}

func (a *PickerControl) EditDialog(value interface{}) *PickerControl {
	a.Set("editDialog", value)
	return a
}

func (a *PickerControl) Editable(value interface{}) *PickerControl {
	a.Set("editable", value)
	return a
}

func (a *PickerControl) EditorSetting(value interface{}) *PickerControl {
	a.Set("editorSetting", value)
	return a
}

func (a *PickerControl) Embed(value interface{}) *PickerControl {
	a.Set("embed", value)
	return a
}

func (a *PickerControl) ExtraName(value interface{}) *PickerControl {
	a.Set("extraName", value)
	return a
}

func (a *PickerControl) ExtractValue(value interface{}) *PickerControl {
	a.Set("extractValue", value)
	return a
}

func (a *PickerControl) Hidden(value interface{}) *PickerControl {
	a.Set("hidden", value)
	return a
}

func (a *PickerControl) HiddenOn(value interface{}) *PickerControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *PickerControl) Hint(value interface{}) *PickerControl {
	a.Set("hint", value)
	return a
}

func (a *PickerControl) Horizontal(value interface{}) *PickerControl {
	a.Set("horizontal", value)
	return a
}

func (a *PickerControl) Id(value interface{}) *PickerControl {
	a.Set("id", value)
	return a
}

func (a *PickerControl) InitAutoFill(value interface{}) *PickerControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *PickerControl) InitFetch(value interface{}) *PickerControl {
	a.Set("initFetch", value)
	return a
}

func (a *PickerControl) InitFetchOn(value interface{}) *PickerControl {
	a.Set("initFetchOn", value)
	return a
}

func (a *PickerControl) Inline(value interface{}) *PickerControl {
	a.Set("inline", value)
	return a
}

func (a *PickerControl) InputClassName(value interface{}) *PickerControl {
	a.Set("inputClassName", value)
	return a
}

func (a *PickerControl) ItemClearable(value interface{}) *PickerControl {
	a.Set("itemClearable", value)
	return a
}

func (a *PickerControl) JoinValues(value interface{}) *PickerControl {
	a.Set("joinValues", value)
	return a
}

func (a *PickerControl) Label(value interface{}) *PickerControl {
	a.Set("label", value)
	return a
}

func (a *PickerControl) LabelAlign(value interface{}) *PickerControl {
	a.Set("labelAlign", value)
	return a
}

func (a *PickerControl) LabelClassName(value interface{}) *PickerControl {
	a.Set("labelClassName", value)
	return a
}

func (a *PickerControl) LabelField(value interface{}) *PickerControl {
	a.Set("labelField", value)
	return a
}

func (a *PickerControl) LabelRemark(value interface{}) *PickerControl {
	a.Set("labelRemark", value)
	return a
}

func (a *PickerControl) LabelTpl(value interface{}) *PickerControl {
	a.Set("labelTpl", value)
	return a
}

func (a *PickerControl) LabelWidth(value interface{}) *PickerControl {
	a.Set("labelWidth", value)
	return a
}

func (a *PickerControl) ModalMode(value interface{}) *PickerControl {
	a.Set("modalMode", value)
	return a
}

func (a *PickerControl) ModalSize(value interface{}) *PickerControl {
	a.Set("modalSize", value)
	return a
}

func (a *PickerControl) ModalTitle(value interface{}) *PickerControl {
	a.Set("modalTitle", value)
	return a
}

func (a *PickerControl) Mode(value interface{}) *PickerControl {
	a.Set("mode", value)
	return a
}

func (a *PickerControl) Multiple(value interface{}) *PickerControl {
	a.Set("multiple", value)
	return a
}

func (a *PickerControl) Name(value interface{}) *PickerControl {
	a.Set("name", value)
	return a
}

func (a *PickerControl) OnEvent(value interface{}) *PickerControl {
	a.Set("onEvent", value)
	return a
}

func (a *PickerControl) Options(value interface{}) *PickerControl {
	a.Set("options", value)
	return a
}

func (a *PickerControl) OverflowConfig(value interface{}) *PickerControl {
	a.Set("overflowConfig", value)
	return a
}

func (a *PickerControl) PickerSchema(value interface{}) *PickerControl {
	a.Set("pickerSchema", value)
	return a
}

func (a *PickerControl) Placeholder(value interface{}) *PickerControl {
	a.Set("placeholder", value)
	return a
}

func (a *PickerControl) ReadOnly(value interface{}) *PickerControl {
	a.Set("readOnly", value)
	return a
}

func (a *PickerControl) ReadOnlyOn(value interface{}) *PickerControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *PickerControl) Remark(value interface{}) *PickerControl {
	a.Set("remark", value)
	return a
}

func (a *PickerControl) Removable(value interface{}) *PickerControl {
	a.Set("removable", value)
	return a
}

func (a *PickerControl) Required(value interface{}) *PickerControl {
	a.Set("required", value)
	return a
}

func (a *PickerControl) ResetValue(value interface{}) *PickerControl {
	a.Set("resetValue", value)
	return a
}

func (a *PickerControl) Row(value interface{}) *PickerControl {
	a.Set("row", value)
	return a
}

func (a *PickerControl) SaveImmediately(value interface{}) *PickerControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *PickerControl) SelectFirst(value interface{}) *PickerControl {
	a.Set("selectFirst", value)
	return a
}

func (a *PickerControl) Size(value interface{}) *PickerControl {
	a.Set("size", value)
	return a
}

func (a *PickerControl) Source(value interface{}) *PickerControl {
	a.Set("source", value)
	return a
}

func (a *PickerControl) Static(value interface{}) *PickerControl {
	a.Set("static", value)
	return a
}

func (a *PickerControl) StaticClassName(value interface{}) *PickerControl {
	a.Set("staticClassName", value)
	return a
}

func (a *PickerControl) StaticInputClassName(value interface{}) *PickerControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *PickerControl) StaticLabelClassName(value interface{}) *PickerControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *PickerControl) StaticOn(value interface{}) *PickerControl {
	a.Set("staticOn", value)
	return a
}

func (a *PickerControl) StaticPlaceholder(value interface{}) *PickerControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *PickerControl) StaticSchema(value interface{}) *PickerControl {
	a.Set("staticSchema", value)
	return a
}

func (a *PickerControl) Style(value interface{}) *PickerControl {
	a.Set("style", value)
	return a
}

func (a *PickerControl) SubmitOnChange(value interface{}) *PickerControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *PickerControl) TestIdBuilder(value interface{}) *PickerControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *PickerControl) Type(value interface{}) *PickerControl {
	a.Set("type", value)
	return a
}

func (a *PickerControl) UseMobileUI(value interface{}) *PickerControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *PickerControl) ValidateApi(value interface{}) *PickerControl {
	a.Set("validateApi", value)
	return a
}

func (a *PickerControl) ValidateOnChange(value interface{}) *PickerControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *PickerControl) ValidationErrors(value interface{}) *PickerControl {
	a.Set("validationErrors", value)
	return a
}

func (a *PickerControl) Validations(value interface{}) *PickerControl {
	a.Set("validations", value)
	return a
}

func (a *PickerControl) Value(value interface{}) *PickerControl {
	a.Set("value", value)
	return a
}

func (a *PickerControl) ValueField(value interface{}) *PickerControl {
	a.Set("valueField", value)
	return a
}

func (a *PickerControl) ValuesNoWrap(value interface{}) *PickerControl {
	a.Set("valuesNoWrap", value)
	return a
}

func (a *PickerControl) Visible(value interface{}) *PickerControl {
	a.Set("visible", value)
	return a
}

func (a *PickerControl) VisibleOn(value interface{}) *PickerControl {
	a.Set("visibleOn", value)
	return a
}

func (a *PickerControl) Width(value interface{}) *PickerControl {
	a.Set("width", value)
	return a
}
