package renderers

type TabsTransferControl struct {
	*BaseRenderer
}

func NewTabsTransferControl() *TabsTransferControl {
	a := &TabsTransferControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "tabs-transfer")
	return a
}

func (a *TabsTransferControl) Set(name string, value interface{}) *TabsTransferControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *TabsTransferControl) AddApi(value interface{}) *TabsTransferControl {
	a.Set("addApi", value)
	return a
}

func (a *TabsTransferControl) AddControls(value interface{}) *TabsTransferControl {
	a.Set("addControls", value)
	return a
}

func (a *TabsTransferControl) AddDialog(value interface{}) *TabsTransferControl {
	a.Set("addDialog", value)
	return a
}

func (a *TabsTransferControl) AutoCheckChildren(value interface{}) *TabsTransferControl {
	a.Set("autoCheckChildren", value)
	return a
}

func (a *TabsTransferControl) AutoFill(value interface{}) *TabsTransferControl {
	a.Set("autoFill", value)
	return a
}

func (a *TabsTransferControl) ClassName(value interface{}) *TabsTransferControl {
	a.Set("className", value)
	return a
}

func (a *TabsTransferControl) ClearValueOnHidden(value interface{}) *TabsTransferControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *TabsTransferControl) ClearValueOnSourceChange(value interface{}) *TabsTransferControl {
	a.Set("clearValueOnSourceChange", value)
	return a
}

func (a *TabsTransferControl) Clearable(value interface{}) *TabsTransferControl {
	a.Set("clearable", value)
	return a
}

func (a *TabsTransferControl) Columns(value interface{}) *TabsTransferControl {
	a.Set("columns", value)
	return a
}

func (a *TabsTransferControl) Creatable(value interface{}) *TabsTransferControl {
	a.Set("creatable", value)
	return a
}

func (a *TabsTransferControl) CreateBtnLabel(value interface{}) *TabsTransferControl {
	a.Set("createBtnLabel", value)
	return a
}

func (a *TabsTransferControl) DeferApi(value interface{}) *TabsTransferControl {
	a.Set("deferApi", value)
	return a
}

func (a *TabsTransferControl) DeferField(value interface{}) *TabsTransferControl {
	a.Set("deferField", value)
	return a
}

func (a *TabsTransferControl) DeleteApi(value interface{}) *TabsTransferControl {
	a.Set("deleteApi", value)
	return a
}

func (a *TabsTransferControl) DeleteConfirmText(value interface{}) *TabsTransferControl {
	a.Set("deleteConfirmText", value)
	return a
}

func (a *TabsTransferControl) Delimiter(value interface{}) *TabsTransferControl {
	a.Set("delimiter", value)
	return a
}

func (a *TabsTransferControl) Desc(value interface{}) *TabsTransferControl {
	a.Set("desc", value)
	return a
}

func (a *TabsTransferControl) Description(value interface{}) *TabsTransferControl {
	a.Set("description", value)
	return a
}

func (a *TabsTransferControl) DescriptionClassName(value interface{}) *TabsTransferControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *TabsTransferControl) Disabled(value interface{}) *TabsTransferControl {
	a.Set("disabled", value)
	return a
}

func (a *TabsTransferControl) DisabledOn(value interface{}) *TabsTransferControl {
	a.Set("disabledOn", value)
	return a
}

func (a *TabsTransferControl) EditApi(value interface{}) *TabsTransferControl {
	a.Set("editApi", value)
	return a
}

func (a *TabsTransferControl) EditControls(value interface{}) *TabsTransferControl {
	a.Set("editControls", value)
	return a
}

func (a *TabsTransferControl) EditDialog(value interface{}) *TabsTransferControl {
	a.Set("editDialog", value)
	return a
}

func (a *TabsTransferControl) Editable(value interface{}) *TabsTransferControl {
	a.Set("editable", value)
	return a
}

func (a *TabsTransferControl) EditorSetting(value interface{}) *TabsTransferControl {
	a.Set("editorSetting", value)
	return a
}

func (a *TabsTransferControl) ExtraName(value interface{}) *TabsTransferControl {
	a.Set("extraName", value)
	return a
}

func (a *TabsTransferControl) ExtractValue(value interface{}) *TabsTransferControl {
	a.Set("extractValue", value)
	return a
}

func (a *TabsTransferControl) Hidden(value interface{}) *TabsTransferControl {
	a.Set("hidden", value)
	return a
}

func (a *TabsTransferControl) HiddenOn(value interface{}) *TabsTransferControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *TabsTransferControl) Hint(value interface{}) *TabsTransferControl {
	a.Set("hint", value)
	return a
}

func (a *TabsTransferControl) Horizontal(value interface{}) *TabsTransferControl {
	a.Set("horizontal", value)
	return a
}

func (a *TabsTransferControl) Id(value interface{}) *TabsTransferControl {
	a.Set("id", value)
	return a
}

func (a *TabsTransferControl) InitAutoFill(value interface{}) *TabsTransferControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *TabsTransferControl) InitFetch(value interface{}) *TabsTransferControl {
	a.Set("initFetch", value)
	return a
}

func (a *TabsTransferControl) InitFetchOn(value interface{}) *TabsTransferControl {
	a.Set("initFetchOn", value)
	return a
}

func (a *TabsTransferControl) InitiallyOpen(value interface{}) *TabsTransferControl {
	a.Set("initiallyOpen", value)
	return a
}

func (a *TabsTransferControl) Inline(value interface{}) *TabsTransferControl {
	a.Set("inline", value)
	return a
}

func (a *TabsTransferControl) InputClassName(value interface{}) *TabsTransferControl {
	a.Set("inputClassName", value)
	return a
}

func (a *TabsTransferControl) ItemHeight(value interface{}) *TabsTransferControl {
	a.Set("itemHeight", value)
	return a
}

func (a *TabsTransferControl) JoinValues(value interface{}) *TabsTransferControl {
	a.Set("joinValues", value)
	return a
}

func (a *TabsTransferControl) Label(value interface{}) *TabsTransferControl {
	a.Set("label", value)
	return a
}

func (a *TabsTransferControl) LabelAlign(value interface{}) *TabsTransferControl {
	a.Set("labelAlign", value)
	return a
}

func (a *TabsTransferControl) LabelClassName(value interface{}) *TabsTransferControl {
	a.Set("labelClassName", value)
	return a
}

func (a *TabsTransferControl) LabelRemark(value interface{}) *TabsTransferControl {
	a.Set("labelRemark", value)
	return a
}

func (a *TabsTransferControl) LabelWidth(value interface{}) *TabsTransferControl {
	a.Set("labelWidth", value)
	return a
}

func (a *TabsTransferControl) LeftMode(value interface{}) *TabsTransferControl {
	a.Set("leftMode", value)
	return a
}

func (a *TabsTransferControl) LeftOptions(value interface{}) *TabsTransferControl {
	a.Set("leftOptions", value)
	return a
}

func (a *TabsTransferControl) LoadingConfig(value interface{}) *TabsTransferControl {
	a.Set("loadingConfig", value)
	return a
}

func (a *TabsTransferControl) MenuTpl(value interface{}) *TabsTransferControl {
	a.Set("menuTpl", value)
	return a
}

func (a *TabsTransferControl) Mode(value interface{}) *TabsTransferControl {
	a.Set("mode", value)
	return a
}

func (a *TabsTransferControl) Multiple(value interface{}) *TabsTransferControl {
	a.Set("multiple", value)
	return a
}

func (a *TabsTransferControl) Name(value interface{}) *TabsTransferControl {
	a.Set("name", value)
	return a
}

func (a *TabsTransferControl) OnEvent(value interface{}) *TabsTransferControl {
	a.Set("onEvent", value)
	return a
}

func (a *TabsTransferControl) OnlyChildren(value interface{}) *TabsTransferControl {
	a.Set("onlyChildren", value)
	return a
}

func (a *TabsTransferControl) Options(value interface{}) *TabsTransferControl {
	a.Set("options", value)
	return a
}

func (a *TabsTransferControl) Pagination(value interface{}) *TabsTransferControl {
	a.Set("pagination", value)
	return a
}

func (a *TabsTransferControl) Placeholder(value interface{}) *TabsTransferControl {
	a.Set("placeholder", value)
	return a
}

func (a *TabsTransferControl) ReadOnly(value interface{}) *TabsTransferControl {
	a.Set("readOnly", value)
	return a
}

func (a *TabsTransferControl) ReadOnlyOn(value interface{}) *TabsTransferControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *TabsTransferControl) Remark(value interface{}) *TabsTransferControl {
	a.Set("remark", value)
	return a
}

func (a *TabsTransferControl) Removable(value interface{}) *TabsTransferControl {
	a.Set("removable", value)
	return a
}

func (a *TabsTransferControl) Required(value interface{}) *TabsTransferControl {
	a.Set("required", value)
	return a
}

func (a *TabsTransferControl) ResetValue(value interface{}) *TabsTransferControl {
	a.Set("resetValue", value)
	return a
}

func (a *TabsTransferControl) ResultListModeFollowSelect(value interface{}) *TabsTransferControl {
	a.Set("resultListModeFollowSelect", value)
	return a
}

func (a *TabsTransferControl) ResultSearchPlaceholder(value interface{}) *TabsTransferControl {
	a.Set("resultSearchPlaceholder", value)
	return a
}

func (a *TabsTransferControl) ResultSearchable(value interface{}) *TabsTransferControl {
	a.Set("resultSearchable", value)
	return a
}

func (a *TabsTransferControl) ResultTitle(value interface{}) *TabsTransferControl {
	a.Set("resultTitle", value)
	return a
}

func (a *TabsTransferControl) RightMode(value interface{}) *TabsTransferControl {
	a.Set("rightMode", value)
	return a
}

func (a *TabsTransferControl) Row(value interface{}) *TabsTransferControl {
	a.Set("row", value)
	return a
}

func (a *TabsTransferControl) SaveImmediately(value interface{}) *TabsTransferControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *TabsTransferControl) SearchApi(value interface{}) *TabsTransferControl {
	a.Set("searchApi", value)
	return a
}

func (a *TabsTransferControl) SearchPlaceholder(value interface{}) *TabsTransferControl {
	a.Set("searchPlaceholder", value)
	return a
}

func (a *TabsTransferControl) SearchResultColumns(value interface{}) *TabsTransferControl {
	a.Set("searchResultColumns", value)
	return a
}

func (a *TabsTransferControl) SearchResultMode(value interface{}) *TabsTransferControl {
	a.Set("searchResultMode", value)
	return a
}

func (a *TabsTransferControl) Searchable(value interface{}) *TabsTransferControl {
	a.Set("searchable", value)
	return a
}

func (a *TabsTransferControl) SelectFirst(value interface{}) *TabsTransferControl {
	a.Set("selectFirst", value)
	return a
}

func (a *TabsTransferControl) SelectMode(value interface{}) *TabsTransferControl {
	a.Set("selectMode", value)
	return a
}

func (a *TabsTransferControl) SelectTitle(value interface{}) *TabsTransferControl {
	a.Set("selectTitle", value)
	return a
}

func (a *TabsTransferControl) ShowArrow(value interface{}) *TabsTransferControl {
	a.Set("showArrow", value)
	return a
}

func (a *TabsTransferControl) ShowInvalidMatch(value interface{}) *TabsTransferControl {
	a.Set("showInvalidMatch", value)
	return a
}

func (a *TabsTransferControl) Size(value interface{}) *TabsTransferControl {
	a.Set("size", value)
	return a
}

func (a *TabsTransferControl) Sortable(value interface{}) *TabsTransferControl {
	a.Set("sortable", value)
	return a
}

func (a *TabsTransferControl) Source(value interface{}) *TabsTransferControl {
	a.Set("source", value)
	return a
}

func (a *TabsTransferControl) Static(value interface{}) *TabsTransferControl {
	a.Set("static", value)
	return a
}

func (a *TabsTransferControl) StaticClassName(value interface{}) *TabsTransferControl {
	a.Set("staticClassName", value)
	return a
}

func (a *TabsTransferControl) StaticInputClassName(value interface{}) *TabsTransferControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *TabsTransferControl) StaticLabelClassName(value interface{}) *TabsTransferControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *TabsTransferControl) StaticOn(value interface{}) *TabsTransferControl {
	a.Set("staticOn", value)
	return a
}

func (a *TabsTransferControl) StaticPlaceholder(value interface{}) *TabsTransferControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *TabsTransferControl) StaticSchema(value interface{}) *TabsTransferControl {
	a.Set("staticSchema", value)
	return a
}

func (a *TabsTransferControl) Statistics(value interface{}) *TabsTransferControl {
	a.Set("statistics", value)
	return a
}

func (a *TabsTransferControl) Style(value interface{}) *TabsTransferControl {
	a.Set("style", value)
	return a
}

func (a *TabsTransferControl) SubmitOnChange(value interface{}) *TabsTransferControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *TabsTransferControl) TestIdBuilder(value interface{}) *TabsTransferControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *TabsTransferControl) Type(value interface{}) *TabsTransferControl {
	a.Set("type", value)
	return a
}

func (a *TabsTransferControl) UseMobileUI(value interface{}) *TabsTransferControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *TabsTransferControl) ValidateApi(value interface{}) *TabsTransferControl {
	a.Set("validateApi", value)
	return a
}

func (a *TabsTransferControl) ValidateOnChange(value interface{}) *TabsTransferControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *TabsTransferControl) ValidationErrors(value interface{}) *TabsTransferControl {
	a.Set("validationErrors", value)
	return a
}

func (a *TabsTransferControl) Validations(value interface{}) *TabsTransferControl {
	a.Set("validations", value)
	return a
}

func (a *TabsTransferControl) Value(value interface{}) *TabsTransferControl {
	a.Set("value", value)
	return a
}

func (a *TabsTransferControl) ValueTpl(value interface{}) *TabsTransferControl {
	a.Set("valueTpl", value)
	return a
}

func (a *TabsTransferControl) ValuesNoWrap(value interface{}) *TabsTransferControl {
	a.Set("valuesNoWrap", value)
	return a
}

func (a *TabsTransferControl) VirtualThreshold(value interface{}) *TabsTransferControl {
	a.Set("virtualThreshold", value)
	return a
}

func (a *TabsTransferControl) Visible(value interface{}) *TabsTransferControl {
	a.Set("visible", value)
	return a
}

func (a *TabsTransferControl) VisibleOn(value interface{}) *TabsTransferControl {
	a.Set("visibleOn", value)
	return a
}

func (a *TabsTransferControl) Width(value interface{}) *TabsTransferControl {
	a.Set("width", value)
	return a
}
