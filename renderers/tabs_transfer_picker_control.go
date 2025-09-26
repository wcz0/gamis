package renderers

type TabsTransferPickerControl struct {
	*BaseRenderer
}

func NewTabsTransferPickerControl() *TabsTransferPickerControl {
	a := &TabsTransferPickerControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "tabs-transfer-picker")
	return a
}

func (a *TabsTransferPickerControl) Set(name string, value interface{}) *TabsTransferPickerControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *TabsTransferPickerControl) AddApi(value interface{}) *TabsTransferPickerControl {
	a.Set("addApi", value)
	return a
}

func (a *TabsTransferPickerControl) AddControls(value interface{}) *TabsTransferPickerControl {
	a.Set("addControls", value)
	return a
}

func (a *TabsTransferPickerControl) AddDialog(value interface{}) *TabsTransferPickerControl {
	a.Set("addDialog", value)
	return a
}

func (a *TabsTransferPickerControl) AutoCheckChildren(value interface{}) *TabsTransferPickerControl {
	a.Set("autoCheckChildren", value)
	return a
}

func (a *TabsTransferPickerControl) AutoFill(value interface{}) *TabsTransferPickerControl {
	a.Set("autoFill", value)
	return a
}

func (a *TabsTransferPickerControl) ClassName(value interface{}) *TabsTransferPickerControl {
	a.Set("className", value)
	return a
}

func (a *TabsTransferPickerControl) ClearValueOnHidden(value interface{}) *TabsTransferPickerControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *TabsTransferPickerControl) ClearValueOnSourceChange(value interface{}) *TabsTransferPickerControl {
	a.Set("clearValueOnSourceChange", value)
	return a
}

func (a *TabsTransferPickerControl) Clearable(value interface{}) *TabsTransferPickerControl {
	a.Set("clearable", value)
	return a
}

func (a *TabsTransferPickerControl) Columns(value interface{}) *TabsTransferPickerControl {
	a.Set("columns", value)
	return a
}

func (a *TabsTransferPickerControl) Creatable(value interface{}) *TabsTransferPickerControl {
	a.Set("creatable", value)
	return a
}

func (a *TabsTransferPickerControl) CreateBtnLabel(value interface{}) *TabsTransferPickerControl {
	a.Set("createBtnLabel", value)
	return a
}

func (a *TabsTransferPickerControl) DeferApi(value interface{}) *TabsTransferPickerControl {
	a.Set("deferApi", value)
	return a
}

func (a *TabsTransferPickerControl) DeferField(value interface{}) *TabsTransferPickerControl {
	a.Set("deferField", value)
	return a
}

func (a *TabsTransferPickerControl) DeleteApi(value interface{}) *TabsTransferPickerControl {
	a.Set("deleteApi", value)
	return a
}

func (a *TabsTransferPickerControl) DeleteConfirmText(value interface{}) *TabsTransferPickerControl {
	a.Set("deleteConfirmText", value)
	return a
}

func (a *TabsTransferPickerControl) Delimiter(value interface{}) *TabsTransferPickerControl {
	a.Set("delimiter", value)
	return a
}

func (a *TabsTransferPickerControl) Desc(value interface{}) *TabsTransferPickerControl {
	a.Set("desc", value)
	return a
}

func (a *TabsTransferPickerControl) Description(value interface{}) *TabsTransferPickerControl {
	a.Set("description", value)
	return a
}

func (a *TabsTransferPickerControl) DescriptionClassName(value interface{}) *TabsTransferPickerControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *TabsTransferPickerControl) Disabled(value interface{}) *TabsTransferPickerControl {
	a.Set("disabled", value)
	return a
}

func (a *TabsTransferPickerControl) DisabledOn(value interface{}) *TabsTransferPickerControl {
	a.Set("disabledOn", value)
	return a
}

func (a *TabsTransferPickerControl) EditApi(value interface{}) *TabsTransferPickerControl {
	a.Set("editApi", value)
	return a
}

func (a *TabsTransferPickerControl) EditControls(value interface{}) *TabsTransferPickerControl {
	a.Set("editControls", value)
	return a
}

func (a *TabsTransferPickerControl) EditDialog(value interface{}) *TabsTransferPickerControl {
	a.Set("editDialog", value)
	return a
}

func (a *TabsTransferPickerControl) Editable(value interface{}) *TabsTransferPickerControl {
	a.Set("editable", value)
	return a
}

func (a *TabsTransferPickerControl) EditorSetting(value interface{}) *TabsTransferPickerControl {
	a.Set("editorSetting", value)
	return a
}

func (a *TabsTransferPickerControl) ExtraName(value interface{}) *TabsTransferPickerControl {
	a.Set("extraName", value)
	return a
}

func (a *TabsTransferPickerControl) ExtractValue(value interface{}) *TabsTransferPickerControl {
	a.Set("extractValue", value)
	return a
}

func (a *TabsTransferPickerControl) Hidden(value interface{}) *TabsTransferPickerControl {
	a.Set("hidden", value)
	return a
}

func (a *TabsTransferPickerControl) HiddenOn(value interface{}) *TabsTransferPickerControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *TabsTransferPickerControl) Hint(value interface{}) *TabsTransferPickerControl {
	a.Set("hint", value)
	return a
}

func (a *TabsTransferPickerControl) Horizontal(value interface{}) *TabsTransferPickerControl {
	a.Set("horizontal", value)
	return a
}

func (a *TabsTransferPickerControl) Id(value interface{}) *TabsTransferPickerControl {
	a.Set("id", value)
	return a
}

func (a *TabsTransferPickerControl) InitAutoFill(value interface{}) *TabsTransferPickerControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *TabsTransferPickerControl) InitFetch(value interface{}) *TabsTransferPickerControl {
	a.Set("initFetch", value)
	return a
}

func (a *TabsTransferPickerControl) InitFetchOn(value interface{}) *TabsTransferPickerControl {
	a.Set("initFetchOn", value)
	return a
}

func (a *TabsTransferPickerControl) InitiallyOpen(value interface{}) *TabsTransferPickerControl {
	a.Set("initiallyOpen", value)
	return a
}

func (a *TabsTransferPickerControl) Inline(value interface{}) *TabsTransferPickerControl {
	a.Set("inline", value)
	return a
}

func (a *TabsTransferPickerControl) InputClassName(value interface{}) *TabsTransferPickerControl {
	a.Set("inputClassName", value)
	return a
}

func (a *TabsTransferPickerControl) ItemHeight(value interface{}) *TabsTransferPickerControl {
	a.Set("itemHeight", value)
	return a
}

func (a *TabsTransferPickerControl) JoinValues(value interface{}) *TabsTransferPickerControl {
	a.Set("joinValues", value)
	return a
}

func (a *TabsTransferPickerControl) Label(value interface{}) *TabsTransferPickerControl {
	a.Set("label", value)
	return a
}

func (a *TabsTransferPickerControl) LabelAlign(value interface{}) *TabsTransferPickerControl {
	a.Set("labelAlign", value)
	return a
}

func (a *TabsTransferPickerControl) LabelClassName(value interface{}) *TabsTransferPickerControl {
	a.Set("labelClassName", value)
	return a
}

func (a *TabsTransferPickerControl) LabelRemark(value interface{}) *TabsTransferPickerControl {
	a.Set("labelRemark", value)
	return a
}

func (a *TabsTransferPickerControl) LabelWidth(value interface{}) *TabsTransferPickerControl {
	a.Set("labelWidth", value)
	return a
}

func (a *TabsTransferPickerControl) LeftMode(value interface{}) *TabsTransferPickerControl {
	a.Set("leftMode", value)
	return a
}

func (a *TabsTransferPickerControl) LeftOptions(value interface{}) *TabsTransferPickerControl {
	a.Set("leftOptions", value)
	return a
}

func (a *TabsTransferPickerControl) LoadingConfig(value interface{}) *TabsTransferPickerControl {
	a.Set("loadingConfig", value)
	return a
}

func (a *TabsTransferPickerControl) MenuTpl(value interface{}) *TabsTransferPickerControl {
	a.Set("menuTpl", value)
	return a
}

func (a *TabsTransferPickerControl) Mode(value interface{}) *TabsTransferPickerControl {
	a.Set("mode", value)
	return a
}

func (a *TabsTransferPickerControl) Multiple(value interface{}) *TabsTransferPickerControl {
	a.Set("multiple", value)
	return a
}

func (a *TabsTransferPickerControl) Name(value interface{}) *TabsTransferPickerControl {
	a.Set("name", value)
	return a
}

func (a *TabsTransferPickerControl) OnEvent(value interface{}) *TabsTransferPickerControl {
	a.Set("onEvent", value)
	return a
}

func (a *TabsTransferPickerControl) OnlyChildren(value interface{}) *TabsTransferPickerControl {
	a.Set("onlyChildren", value)
	return a
}

func (a *TabsTransferPickerControl) Options(value interface{}) *TabsTransferPickerControl {
	a.Set("options", value)
	return a
}

func (a *TabsTransferPickerControl) Pagination(value interface{}) *TabsTransferPickerControl {
	a.Set("pagination", value)
	return a
}

func (a *TabsTransferPickerControl) Placeholder(value interface{}) *TabsTransferPickerControl {
	a.Set("placeholder", value)
	return a
}

func (a *TabsTransferPickerControl) ReadOnly(value interface{}) *TabsTransferPickerControl {
	a.Set("readOnly", value)
	return a
}

func (a *TabsTransferPickerControl) ReadOnlyOn(value interface{}) *TabsTransferPickerControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *TabsTransferPickerControl) Remark(value interface{}) *TabsTransferPickerControl {
	a.Set("remark", value)
	return a
}

func (a *TabsTransferPickerControl) Removable(value interface{}) *TabsTransferPickerControl {
	a.Set("removable", value)
	return a
}

func (a *TabsTransferPickerControl) Required(value interface{}) *TabsTransferPickerControl {
	a.Set("required", value)
	return a
}

func (a *TabsTransferPickerControl) ResetValue(value interface{}) *TabsTransferPickerControl {
	a.Set("resetValue", value)
	return a
}

func (a *TabsTransferPickerControl) ResultListModeFollowSelect(value interface{}) *TabsTransferPickerControl {
	a.Set("resultListModeFollowSelect", value)
	return a
}

func (a *TabsTransferPickerControl) ResultSearchPlaceholder(value interface{}) *TabsTransferPickerControl {
	a.Set("resultSearchPlaceholder", value)
	return a
}

func (a *TabsTransferPickerControl) ResultSearchable(value interface{}) *TabsTransferPickerControl {
	a.Set("resultSearchable", value)
	return a
}

func (a *TabsTransferPickerControl) ResultTitle(value interface{}) *TabsTransferPickerControl {
	a.Set("resultTitle", value)
	return a
}

func (a *TabsTransferPickerControl) RightMode(value interface{}) *TabsTransferPickerControl {
	a.Set("rightMode", value)
	return a
}

func (a *TabsTransferPickerControl) Row(value interface{}) *TabsTransferPickerControl {
	a.Set("row", value)
	return a
}

func (a *TabsTransferPickerControl) SaveImmediately(value interface{}) *TabsTransferPickerControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *TabsTransferPickerControl) SearchApi(value interface{}) *TabsTransferPickerControl {
	a.Set("searchApi", value)
	return a
}

func (a *TabsTransferPickerControl) SearchPlaceholder(value interface{}) *TabsTransferPickerControl {
	a.Set("searchPlaceholder", value)
	return a
}

func (a *TabsTransferPickerControl) SearchResultColumns(value interface{}) *TabsTransferPickerControl {
	a.Set("searchResultColumns", value)
	return a
}

func (a *TabsTransferPickerControl) SearchResultMode(value interface{}) *TabsTransferPickerControl {
	a.Set("searchResultMode", value)
	return a
}

func (a *TabsTransferPickerControl) Searchable(value interface{}) *TabsTransferPickerControl {
	a.Set("searchable", value)
	return a
}

func (a *TabsTransferPickerControl) SelectFirst(value interface{}) *TabsTransferPickerControl {
	a.Set("selectFirst", value)
	return a
}

func (a *TabsTransferPickerControl) SelectMode(value interface{}) *TabsTransferPickerControl {
	a.Set("selectMode", value)
	return a
}

func (a *TabsTransferPickerControl) SelectTitle(value interface{}) *TabsTransferPickerControl {
	a.Set("selectTitle", value)
	return a
}

func (a *TabsTransferPickerControl) ShowArrow(value interface{}) *TabsTransferPickerControl {
	a.Set("showArrow", value)
	return a
}

func (a *TabsTransferPickerControl) ShowInvalidMatch(value interface{}) *TabsTransferPickerControl {
	a.Set("showInvalidMatch", value)
	return a
}

func (a *TabsTransferPickerControl) Size(value interface{}) *TabsTransferPickerControl {
	a.Set("size", value)
	return a
}

func (a *TabsTransferPickerControl) Sortable(value interface{}) *TabsTransferPickerControl {
	a.Set("sortable", value)
	return a
}

func (a *TabsTransferPickerControl) Source(value interface{}) *TabsTransferPickerControl {
	a.Set("source", value)
	return a
}

func (a *TabsTransferPickerControl) Static(value interface{}) *TabsTransferPickerControl {
	a.Set("static", value)
	return a
}

func (a *TabsTransferPickerControl) StaticClassName(value interface{}) *TabsTransferPickerControl {
	a.Set("staticClassName", value)
	return a
}

func (a *TabsTransferPickerControl) StaticInputClassName(value interface{}) *TabsTransferPickerControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *TabsTransferPickerControl) StaticLabelClassName(value interface{}) *TabsTransferPickerControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *TabsTransferPickerControl) StaticOn(value interface{}) *TabsTransferPickerControl {
	a.Set("staticOn", value)
	return a
}

func (a *TabsTransferPickerControl) StaticPlaceholder(value interface{}) *TabsTransferPickerControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *TabsTransferPickerControl) StaticSchema(value interface{}) *TabsTransferPickerControl {
	a.Set("staticSchema", value)
	return a
}

func (a *TabsTransferPickerControl) Statistics(value interface{}) *TabsTransferPickerControl {
	a.Set("statistics", value)
	return a
}

func (a *TabsTransferPickerControl) Style(value interface{}) *TabsTransferPickerControl {
	a.Set("style", value)
	return a
}

func (a *TabsTransferPickerControl) SubmitOnChange(value interface{}) *TabsTransferPickerControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *TabsTransferPickerControl) TestIdBuilder(value interface{}) *TabsTransferPickerControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *TabsTransferPickerControl) Type(value interface{}) *TabsTransferPickerControl {
	a.Set("type", value)
	return a
}

func (a *TabsTransferPickerControl) UseMobileUI(value interface{}) *TabsTransferPickerControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *TabsTransferPickerControl) ValidateApi(value interface{}) *TabsTransferPickerControl {
	a.Set("validateApi", value)
	return a
}

func (a *TabsTransferPickerControl) ValidateOnChange(value interface{}) *TabsTransferPickerControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *TabsTransferPickerControl) ValidationErrors(value interface{}) *TabsTransferPickerControl {
	a.Set("validationErrors", value)
	return a
}

func (a *TabsTransferPickerControl) Validations(value interface{}) *TabsTransferPickerControl {
	a.Set("validations", value)
	return a
}

func (a *TabsTransferPickerControl) Value(value interface{}) *TabsTransferPickerControl {
	a.Set("value", value)
	return a
}

func (a *TabsTransferPickerControl) ValueTpl(value interface{}) *TabsTransferPickerControl {
	a.Set("valueTpl", value)
	return a
}

func (a *TabsTransferPickerControl) ValuesNoWrap(value interface{}) *TabsTransferPickerControl {
	a.Set("valuesNoWrap", value)
	return a
}

func (a *TabsTransferPickerControl) VirtualThreshold(value interface{}) *TabsTransferPickerControl {
	a.Set("virtualThreshold", value)
	return a
}

func (a *TabsTransferPickerControl) Visible(value interface{}) *TabsTransferPickerControl {
	a.Set("visible", value)
	return a
}

func (a *TabsTransferPickerControl) VisibleOn(value interface{}) *TabsTransferPickerControl {
	a.Set("visibleOn", value)
	return a
}

func (a *TabsTransferPickerControl) Width(value interface{}) *TabsTransferPickerControl {
	a.Set("width", value)
	return a
}
