package renderers

type TransferPickerControl struct {
	*BaseRenderer
}

func NewTransferPickerControl() *TransferPickerControl {
	a := &TransferPickerControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "transfer-picker")
	return a
}

func (a *TransferPickerControl) Set(name string, value interface{}) *TransferPickerControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *TransferPickerControl) AddApi(value interface{}) *TransferPickerControl {
	a.Set("addApi", value)
	return a
}

func (a *TransferPickerControl) AddControls(value interface{}) *TransferPickerControl {
	a.Set("addControls", value)
	return a
}

func (a *TransferPickerControl) AddDialog(value interface{}) *TransferPickerControl {
	a.Set("addDialog", value)
	return a
}

func (a *TransferPickerControl) AutoCheckChildren(value interface{}) *TransferPickerControl {
	a.Set("autoCheckChildren", value)
	return a
}

func (a *TransferPickerControl) AutoFill(value interface{}) *TransferPickerControl {
	a.Set("autoFill", value)
	return a
}

func (a *TransferPickerControl) BorderMode(value interface{}) *TransferPickerControl {
	a.Set("borderMode", value)
	return a
}

func (a *TransferPickerControl) ClassName(value interface{}) *TransferPickerControl {
	a.Set("className", value)
	return a
}

func (a *TransferPickerControl) ClearValueOnHidden(value interface{}) *TransferPickerControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *TransferPickerControl) ClearValueOnSourceChange(value interface{}) *TransferPickerControl {
	a.Set("clearValueOnSourceChange", value)
	return a
}

func (a *TransferPickerControl) Clearable(value interface{}) *TransferPickerControl {
	a.Set("clearable", value)
	return a
}

func (a *TransferPickerControl) Columns(value interface{}) *TransferPickerControl {
	a.Set("columns", value)
	return a
}

func (a *TransferPickerControl) Creatable(value interface{}) *TransferPickerControl {
	a.Set("creatable", value)
	return a
}

func (a *TransferPickerControl) CreateBtnLabel(value interface{}) *TransferPickerControl {
	a.Set("createBtnLabel", value)
	return a
}

func (a *TransferPickerControl) DeferApi(value interface{}) *TransferPickerControl {
	a.Set("deferApi", value)
	return a
}

func (a *TransferPickerControl) DeferField(value interface{}) *TransferPickerControl {
	a.Set("deferField", value)
	return a
}

func (a *TransferPickerControl) DeleteApi(value interface{}) *TransferPickerControl {
	a.Set("deleteApi", value)
	return a
}

func (a *TransferPickerControl) DeleteConfirmText(value interface{}) *TransferPickerControl {
	a.Set("deleteConfirmText", value)
	return a
}

func (a *TransferPickerControl) Delimiter(value interface{}) *TransferPickerControl {
	a.Set("delimiter", value)
	return a
}

func (a *TransferPickerControl) Desc(value interface{}) *TransferPickerControl {
	a.Set("desc", value)
	return a
}

func (a *TransferPickerControl) Description(value interface{}) *TransferPickerControl {
	a.Set("description", value)
	return a
}

func (a *TransferPickerControl) DescriptionClassName(value interface{}) *TransferPickerControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *TransferPickerControl) Disabled(value interface{}) *TransferPickerControl {
	a.Set("disabled", value)
	return a
}

func (a *TransferPickerControl) DisabledOn(value interface{}) *TransferPickerControl {
	a.Set("disabledOn", value)
	return a
}

func (a *TransferPickerControl) EditApi(value interface{}) *TransferPickerControl {
	a.Set("editApi", value)
	return a
}

func (a *TransferPickerControl) EditControls(value interface{}) *TransferPickerControl {
	a.Set("editControls", value)
	return a
}

func (a *TransferPickerControl) EditDialog(value interface{}) *TransferPickerControl {
	a.Set("editDialog", value)
	return a
}

func (a *TransferPickerControl) Editable(value interface{}) *TransferPickerControl {
	a.Set("editable", value)
	return a
}

func (a *TransferPickerControl) EditorSetting(value interface{}) *TransferPickerControl {
	a.Set("editorSetting", value)
	return a
}

func (a *TransferPickerControl) ExtraName(value interface{}) *TransferPickerControl {
	a.Set("extraName", value)
	return a
}

func (a *TransferPickerControl) ExtractValue(value interface{}) *TransferPickerControl {
	a.Set("extractValue", value)
	return a
}

func (a *TransferPickerControl) Hidden(value interface{}) *TransferPickerControl {
	a.Set("hidden", value)
	return a
}

func (a *TransferPickerControl) HiddenOn(value interface{}) *TransferPickerControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *TransferPickerControl) Hint(value interface{}) *TransferPickerControl {
	a.Set("hint", value)
	return a
}

func (a *TransferPickerControl) Horizontal(value interface{}) *TransferPickerControl {
	a.Set("horizontal", value)
	return a
}

func (a *TransferPickerControl) Id(value interface{}) *TransferPickerControl {
	a.Set("id", value)
	return a
}

func (a *TransferPickerControl) InitAutoFill(value interface{}) *TransferPickerControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *TransferPickerControl) InitFetch(value interface{}) *TransferPickerControl {
	a.Set("initFetch", value)
	return a
}

func (a *TransferPickerControl) InitFetchOn(value interface{}) *TransferPickerControl {
	a.Set("initFetchOn", value)
	return a
}

func (a *TransferPickerControl) InitiallyOpen(value interface{}) *TransferPickerControl {
	a.Set("initiallyOpen", value)
	return a
}

func (a *TransferPickerControl) Inline(value interface{}) *TransferPickerControl {
	a.Set("inline", value)
	return a
}

func (a *TransferPickerControl) InputClassName(value interface{}) *TransferPickerControl {
	a.Set("inputClassName", value)
	return a
}

func (a *TransferPickerControl) ItemHeight(value interface{}) *TransferPickerControl {
	a.Set("itemHeight", value)
	return a
}

func (a *TransferPickerControl) JoinValues(value interface{}) *TransferPickerControl {
	a.Set("joinValues", value)
	return a
}

func (a *TransferPickerControl) Label(value interface{}) *TransferPickerControl {
	a.Set("label", value)
	return a
}

func (a *TransferPickerControl) LabelAlign(value interface{}) *TransferPickerControl {
	a.Set("labelAlign", value)
	return a
}

func (a *TransferPickerControl) LabelClassName(value interface{}) *TransferPickerControl {
	a.Set("labelClassName", value)
	return a
}

func (a *TransferPickerControl) LabelRemark(value interface{}) *TransferPickerControl {
	a.Set("labelRemark", value)
	return a
}

func (a *TransferPickerControl) LabelWidth(value interface{}) *TransferPickerControl {
	a.Set("labelWidth", value)
	return a
}

func (a *TransferPickerControl) LeftMode(value interface{}) *TransferPickerControl {
	a.Set("leftMode", value)
	return a
}

func (a *TransferPickerControl) LeftOptions(value interface{}) *TransferPickerControl {
	a.Set("leftOptions", value)
	return a
}

func (a *TransferPickerControl) LoadingConfig(value interface{}) *TransferPickerControl {
	a.Set("loadingConfig", value)
	return a
}

func (a *TransferPickerControl) MenuTpl(value interface{}) *TransferPickerControl {
	a.Set("menuTpl", value)
	return a
}

func (a *TransferPickerControl) Mode(value interface{}) *TransferPickerControl {
	a.Set("mode", value)
	return a
}

func (a *TransferPickerControl) Multiple(value interface{}) *TransferPickerControl {
	a.Set("multiple", value)
	return a
}

func (a *TransferPickerControl) Name(value interface{}) *TransferPickerControl {
	a.Set("name", value)
	return a
}

func (a *TransferPickerControl) OnEvent(value interface{}) *TransferPickerControl {
	a.Set("onEvent", value)
	return a
}

func (a *TransferPickerControl) OnlyChildren(value interface{}) *TransferPickerControl {
	a.Set("onlyChildren", value)
	return a
}

func (a *TransferPickerControl) Options(value interface{}) *TransferPickerControl {
	a.Set("options", value)
	return a
}

func (a *TransferPickerControl) Pagination(value interface{}) *TransferPickerControl {
	a.Set("pagination", value)
	return a
}

func (a *TransferPickerControl) PickerSize(value interface{}) *TransferPickerControl {
	a.Set("pickerSize", value)
	return a
}

func (a *TransferPickerControl) Placeholder(value interface{}) *TransferPickerControl {
	a.Set("placeholder", value)
	return a
}

func (a *TransferPickerControl) ReadOnly(value interface{}) *TransferPickerControl {
	a.Set("readOnly", value)
	return a
}

func (a *TransferPickerControl) ReadOnlyOn(value interface{}) *TransferPickerControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *TransferPickerControl) Remark(value interface{}) *TransferPickerControl {
	a.Set("remark", value)
	return a
}

func (a *TransferPickerControl) Removable(value interface{}) *TransferPickerControl {
	a.Set("removable", value)
	return a
}

func (a *TransferPickerControl) Required(value interface{}) *TransferPickerControl {
	a.Set("required", value)
	return a
}

func (a *TransferPickerControl) ResetValue(value interface{}) *TransferPickerControl {
	a.Set("resetValue", value)
	return a
}

func (a *TransferPickerControl) ResultListModeFollowSelect(value interface{}) *TransferPickerControl {
	a.Set("resultListModeFollowSelect", value)
	return a
}

func (a *TransferPickerControl) ResultSearchPlaceholder(value interface{}) *TransferPickerControl {
	a.Set("resultSearchPlaceholder", value)
	return a
}

func (a *TransferPickerControl) ResultSearchable(value interface{}) *TransferPickerControl {
	a.Set("resultSearchable", value)
	return a
}

func (a *TransferPickerControl) ResultTitle(value interface{}) *TransferPickerControl {
	a.Set("resultTitle", value)
	return a
}

func (a *TransferPickerControl) RightMode(value interface{}) *TransferPickerControl {
	a.Set("rightMode", value)
	return a
}

func (a *TransferPickerControl) Row(value interface{}) *TransferPickerControl {
	a.Set("row", value)
	return a
}

func (a *TransferPickerControl) SaveImmediately(value interface{}) *TransferPickerControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *TransferPickerControl) SearchApi(value interface{}) *TransferPickerControl {
	a.Set("searchApi", value)
	return a
}

func (a *TransferPickerControl) SearchPlaceholder(value interface{}) *TransferPickerControl {
	a.Set("searchPlaceholder", value)
	return a
}

func (a *TransferPickerControl) SearchResultColumns(value interface{}) *TransferPickerControl {
	a.Set("searchResultColumns", value)
	return a
}

func (a *TransferPickerControl) SearchResultMode(value interface{}) *TransferPickerControl {
	a.Set("searchResultMode", value)
	return a
}

func (a *TransferPickerControl) Searchable(value interface{}) *TransferPickerControl {
	a.Set("searchable", value)
	return a
}

func (a *TransferPickerControl) SelectFirst(value interface{}) *TransferPickerControl {
	a.Set("selectFirst", value)
	return a
}

func (a *TransferPickerControl) SelectMode(value interface{}) *TransferPickerControl {
	a.Set("selectMode", value)
	return a
}

func (a *TransferPickerControl) SelectTitle(value interface{}) *TransferPickerControl {
	a.Set("selectTitle", value)
	return a
}

func (a *TransferPickerControl) ShowArrow(value interface{}) *TransferPickerControl {
	a.Set("showArrow", value)
	return a
}

func (a *TransferPickerControl) ShowInvalidMatch(value interface{}) *TransferPickerControl {
	a.Set("showInvalidMatch", value)
	return a
}

func (a *TransferPickerControl) Size(value interface{}) *TransferPickerControl {
	a.Set("size", value)
	return a
}

func (a *TransferPickerControl) Sortable(value interface{}) *TransferPickerControl {
	a.Set("sortable", value)
	return a
}

func (a *TransferPickerControl) Source(value interface{}) *TransferPickerControl {
	a.Set("source", value)
	return a
}

func (a *TransferPickerControl) Static(value interface{}) *TransferPickerControl {
	a.Set("static", value)
	return a
}

func (a *TransferPickerControl) StaticClassName(value interface{}) *TransferPickerControl {
	a.Set("staticClassName", value)
	return a
}

func (a *TransferPickerControl) StaticInputClassName(value interface{}) *TransferPickerControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *TransferPickerControl) StaticLabelClassName(value interface{}) *TransferPickerControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *TransferPickerControl) StaticOn(value interface{}) *TransferPickerControl {
	a.Set("staticOn", value)
	return a
}

func (a *TransferPickerControl) StaticPlaceholder(value interface{}) *TransferPickerControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *TransferPickerControl) StaticSchema(value interface{}) *TransferPickerControl {
	a.Set("staticSchema", value)
	return a
}

func (a *TransferPickerControl) Statistics(value interface{}) *TransferPickerControl {
	a.Set("statistics", value)
	return a
}

func (a *TransferPickerControl) Style(value interface{}) *TransferPickerControl {
	a.Set("style", value)
	return a
}

func (a *TransferPickerControl) SubmitOnChange(value interface{}) *TransferPickerControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *TransferPickerControl) TestIdBuilder(value interface{}) *TransferPickerControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *TransferPickerControl) Type(value interface{}) *TransferPickerControl {
	a.Set("type", value)
	return a
}

func (a *TransferPickerControl) UseMobileUI(value interface{}) *TransferPickerControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *TransferPickerControl) ValidateApi(value interface{}) *TransferPickerControl {
	a.Set("validateApi", value)
	return a
}

func (a *TransferPickerControl) ValidateOnChange(value interface{}) *TransferPickerControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *TransferPickerControl) ValidationErrors(value interface{}) *TransferPickerControl {
	a.Set("validationErrors", value)
	return a
}

func (a *TransferPickerControl) Validations(value interface{}) *TransferPickerControl {
	a.Set("validations", value)
	return a
}

func (a *TransferPickerControl) Value(value interface{}) *TransferPickerControl {
	a.Set("value", value)
	return a
}

func (a *TransferPickerControl) ValueTpl(value interface{}) *TransferPickerControl {
	a.Set("valueTpl", value)
	return a
}

func (a *TransferPickerControl) ValuesNoWrap(value interface{}) *TransferPickerControl {
	a.Set("valuesNoWrap", value)
	return a
}

func (a *TransferPickerControl) VirtualThreshold(value interface{}) *TransferPickerControl {
	a.Set("virtualThreshold", value)
	return a
}

func (a *TransferPickerControl) Visible(value interface{}) *TransferPickerControl {
	a.Set("visible", value)
	return a
}

func (a *TransferPickerControl) VisibleOn(value interface{}) *TransferPickerControl {
	a.Set("visibleOn", value)
	return a
}

func (a *TransferPickerControl) Width(value interface{}) *TransferPickerControl {
	a.Set("width", value)
	return a
}
