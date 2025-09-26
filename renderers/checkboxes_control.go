package renderers

type CheckboxesControl struct {
	*BaseRenderer
}

func NewCheckboxesControl() *CheckboxesControl {
	a := &CheckboxesControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "checkboxes")
	return a
}

func (a *CheckboxesControl) Set(name string, value interface{}) *CheckboxesControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *CheckboxesControl) AddApi(value interface{}) *CheckboxesControl {
	a.Set("addApi", value)
	return a
}

func (a *CheckboxesControl) AddControls(value interface{}) *CheckboxesControl {
	a.Set("addControls", value)
	return a
}

func (a *CheckboxesControl) AddDialog(value interface{}) *CheckboxesControl {
	a.Set("addDialog", value)
	return a
}

func (a *CheckboxesControl) AutoFill(value interface{}) *CheckboxesControl {
	a.Set("autoFill", value)
	return a
}

func (a *CheckboxesControl) CheckAll(value interface{}) *CheckboxesControl {
	a.Set("checkAll", value)
	return a
}

func (a *CheckboxesControl) CheckAllText(value interface{}) *CheckboxesControl {
	a.Set("checkAllText", value)
	return a
}

func (a *CheckboxesControl) ClassName(value interface{}) *CheckboxesControl {
	a.Set("className", value)
	return a
}

func (a *CheckboxesControl) ClearValueOnHidden(value interface{}) *CheckboxesControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *CheckboxesControl) ClearValueOnSourceChange(value interface{}) *CheckboxesControl {
	a.Set("clearValueOnSourceChange", value)
	return a
}

func (a *CheckboxesControl) Clearable(value interface{}) *CheckboxesControl {
	a.Set("clearable", value)
	return a
}

func (a *CheckboxesControl) ColumnsCount(value interface{}) *CheckboxesControl {
	a.Set("columnsCount", value)
	return a
}

func (a *CheckboxesControl) Creatable(value interface{}) *CheckboxesControl {
	a.Set("creatable", value)
	return a
}

func (a *CheckboxesControl) CreateBtnLabel(value interface{}) *CheckboxesControl {
	a.Set("createBtnLabel", value)
	return a
}

func (a *CheckboxesControl) DefaultCheckAll(value interface{}) *CheckboxesControl {
	a.Set("defaultCheckAll", value)
	return a
}

func (a *CheckboxesControl) DeferApi(value interface{}) *CheckboxesControl {
	a.Set("deferApi", value)
	return a
}

func (a *CheckboxesControl) DeferField(value interface{}) *CheckboxesControl {
	a.Set("deferField", value)
	return a
}

func (a *CheckboxesControl) DeleteApi(value interface{}) *CheckboxesControl {
	a.Set("deleteApi", value)
	return a
}

func (a *CheckboxesControl) DeleteConfirmText(value interface{}) *CheckboxesControl {
	a.Set("deleteConfirmText", value)
	return a
}

func (a *CheckboxesControl) Delimiter(value interface{}) *CheckboxesControl {
	a.Set("delimiter", value)
	return a
}

func (a *CheckboxesControl) Desc(value interface{}) *CheckboxesControl {
	a.Set("desc", value)
	return a
}

func (a *CheckboxesControl) Description(value interface{}) *CheckboxesControl {
	a.Set("description", value)
	return a
}

func (a *CheckboxesControl) DescriptionClassName(value interface{}) *CheckboxesControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *CheckboxesControl) Disabled(value interface{}) *CheckboxesControl {
	a.Set("disabled", value)
	return a
}

func (a *CheckboxesControl) DisabledOn(value interface{}) *CheckboxesControl {
	a.Set("disabledOn", value)
	return a
}

func (a *CheckboxesControl) EditApi(value interface{}) *CheckboxesControl {
	a.Set("editApi", value)
	return a
}

func (a *CheckboxesControl) EditControls(value interface{}) *CheckboxesControl {
	a.Set("editControls", value)
	return a
}

func (a *CheckboxesControl) EditDialog(value interface{}) *CheckboxesControl {
	a.Set("editDialog", value)
	return a
}

func (a *CheckboxesControl) Editable(value interface{}) *CheckboxesControl {
	a.Set("editable", value)
	return a
}

func (a *CheckboxesControl) EditorSetting(value interface{}) *CheckboxesControl {
	a.Set("editorSetting", value)
	return a
}

func (a *CheckboxesControl) ExtraName(value interface{}) *CheckboxesControl {
	a.Set("extraName", value)
	return a
}

func (a *CheckboxesControl) ExtractValue(value interface{}) *CheckboxesControl {
	a.Set("extractValue", value)
	return a
}

func (a *CheckboxesControl) Hidden(value interface{}) *CheckboxesControl {
	a.Set("hidden", value)
	return a
}

func (a *CheckboxesControl) HiddenOn(value interface{}) *CheckboxesControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *CheckboxesControl) Hint(value interface{}) *CheckboxesControl {
	a.Set("hint", value)
	return a
}

func (a *CheckboxesControl) Horizontal(value interface{}) *CheckboxesControl {
	a.Set("horizontal", value)
	return a
}

func (a *CheckboxesControl) Id(value interface{}) *CheckboxesControl {
	a.Set("id", value)
	return a
}

func (a *CheckboxesControl) InitAutoFill(value interface{}) *CheckboxesControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *CheckboxesControl) InitFetch(value interface{}) *CheckboxesControl {
	a.Set("initFetch", value)
	return a
}

func (a *CheckboxesControl) InitFetchOn(value interface{}) *CheckboxesControl {
	a.Set("initFetchOn", value)
	return a
}

func (a *CheckboxesControl) Inline(value interface{}) *CheckboxesControl {
	a.Set("inline", value)
	return a
}

func (a *CheckboxesControl) InputClassName(value interface{}) *CheckboxesControl {
	a.Set("inputClassName", value)
	return a
}

func (a *CheckboxesControl) JoinValues(value interface{}) *CheckboxesControl {
	a.Set("joinValues", value)
	return a
}

func (a *CheckboxesControl) Label(value interface{}) *CheckboxesControl {
	a.Set("label", value)
	return a
}

func (a *CheckboxesControl) LabelAlign(value interface{}) *CheckboxesControl {
	a.Set("labelAlign", value)
	return a
}

func (a *CheckboxesControl) LabelClassName(value interface{}) *CheckboxesControl {
	a.Set("labelClassName", value)
	return a
}

func (a *CheckboxesControl) LabelRemark(value interface{}) *CheckboxesControl {
	a.Set("labelRemark", value)
	return a
}

func (a *CheckboxesControl) LabelWidth(value interface{}) *CheckboxesControl {
	a.Set("labelWidth", value)
	return a
}

func (a *CheckboxesControl) MenuTpl(value interface{}) *CheckboxesControl {
	a.Set("menuTpl", value)
	return a
}

func (a *CheckboxesControl) Mode(value interface{}) *CheckboxesControl {
	a.Set("mode", value)
	return a
}

func (a *CheckboxesControl) Multiple(value interface{}) *CheckboxesControl {
	a.Set("multiple", value)
	return a
}

func (a *CheckboxesControl) Name(value interface{}) *CheckboxesControl {
	a.Set("name", value)
	return a
}

func (a *CheckboxesControl) OnEvent(value interface{}) *CheckboxesControl {
	a.Set("onEvent", value)
	return a
}

func (a *CheckboxesControl) Options(value interface{}) *CheckboxesControl {
	a.Set("options", value)
	return a
}

func (a *CheckboxesControl) Placeholder(value interface{}) *CheckboxesControl {
	a.Set("placeholder", value)
	return a
}

func (a *CheckboxesControl) ReadOnly(value interface{}) *CheckboxesControl {
	a.Set("readOnly", value)
	return a
}

func (a *CheckboxesControl) ReadOnlyOn(value interface{}) *CheckboxesControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *CheckboxesControl) Remark(value interface{}) *CheckboxesControl {
	a.Set("remark", value)
	return a
}

func (a *CheckboxesControl) Removable(value interface{}) *CheckboxesControl {
	a.Set("removable", value)
	return a
}

func (a *CheckboxesControl) Required(value interface{}) *CheckboxesControl {
	a.Set("required", value)
	return a
}

func (a *CheckboxesControl) ResetValue(value interface{}) *CheckboxesControl {
	a.Set("resetValue", value)
	return a
}

func (a *CheckboxesControl) Row(value interface{}) *CheckboxesControl {
	a.Set("row", value)
	return a
}

func (a *CheckboxesControl) SaveImmediately(value interface{}) *CheckboxesControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *CheckboxesControl) SelectFirst(value interface{}) *CheckboxesControl {
	a.Set("selectFirst", value)
	return a
}

func (a *CheckboxesControl) Size(value interface{}) *CheckboxesControl {
	a.Set("size", value)
	return a
}

func (a *CheckboxesControl) Source(value interface{}) *CheckboxesControl {
	a.Set("source", value)
	return a
}

func (a *CheckboxesControl) Static(value interface{}) *CheckboxesControl {
	a.Set("static", value)
	return a
}

func (a *CheckboxesControl) StaticClassName(value interface{}) *CheckboxesControl {
	a.Set("staticClassName", value)
	return a
}

func (a *CheckboxesControl) StaticInputClassName(value interface{}) *CheckboxesControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *CheckboxesControl) StaticLabelClassName(value interface{}) *CheckboxesControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *CheckboxesControl) StaticOn(value interface{}) *CheckboxesControl {
	a.Set("staticOn", value)
	return a
}

func (a *CheckboxesControl) StaticPlaceholder(value interface{}) *CheckboxesControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *CheckboxesControl) StaticSchema(value interface{}) *CheckboxesControl {
	a.Set("staticSchema", value)
	return a
}

func (a *CheckboxesControl) Style(value interface{}) *CheckboxesControl {
	a.Set("style", value)
	return a
}

func (a *CheckboxesControl) SubmitOnChange(value interface{}) *CheckboxesControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *CheckboxesControl) TestIdBuilder(value interface{}) *CheckboxesControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *CheckboxesControl) Type(value interface{}) *CheckboxesControl {
	a.Set("type", value)
	return a
}

func (a *CheckboxesControl) UseMobileUI(value interface{}) *CheckboxesControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *CheckboxesControl) ValidateApi(value interface{}) *CheckboxesControl {
	a.Set("validateApi", value)
	return a
}

func (a *CheckboxesControl) ValidateOnChange(value interface{}) *CheckboxesControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *CheckboxesControl) ValidationErrors(value interface{}) *CheckboxesControl {
	a.Set("validationErrors", value)
	return a
}

func (a *CheckboxesControl) Validations(value interface{}) *CheckboxesControl {
	a.Set("validations", value)
	return a
}

func (a *CheckboxesControl) Value(value interface{}) *CheckboxesControl {
	a.Set("value", value)
	return a
}

func (a *CheckboxesControl) ValuesNoWrap(value interface{}) *CheckboxesControl {
	a.Set("valuesNoWrap", value)
	return a
}

func (a *CheckboxesControl) Visible(value interface{}) *CheckboxesControl {
	a.Set("visible", value)
	return a
}

func (a *CheckboxesControl) VisibleOn(value interface{}) *CheckboxesControl {
	a.Set("visibleOn", value)
	return a
}

func (a *CheckboxesControl) Width(value interface{}) *CheckboxesControl {
	a.Set("width", value)
	return a
}
