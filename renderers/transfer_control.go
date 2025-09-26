package renderers

type TransferControl struct {
	*BaseRenderer
}

func NewTransferControl() *TransferControl {
	a := &TransferControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "transfer")
	return a
}

func (a *TransferControl) Set(name string, value interface{}) *TransferControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *TransferControl) AddApi(value interface{}) *TransferControl {
	a.Set("addApi", value)
	return a
}

func (a *TransferControl) AddControls(value interface{}) *TransferControl {
	a.Set("addControls", value)
	return a
}

func (a *TransferControl) AddDialog(value interface{}) *TransferControl {
	a.Set("addDialog", value)
	return a
}

func (a *TransferControl) AutoCheckChildren(value interface{}) *TransferControl {
	a.Set("autoCheckChildren", value)
	return a
}

func (a *TransferControl) AutoFill(value interface{}) *TransferControl {
	a.Set("autoFill", value)
	return a
}

func (a *TransferControl) ClassName(value interface{}) *TransferControl {
	a.Set("className", value)
	return a
}

func (a *TransferControl) ClearValueOnHidden(value interface{}) *TransferControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *TransferControl) ClearValueOnSourceChange(value interface{}) *TransferControl {
	a.Set("clearValueOnSourceChange", value)
	return a
}

func (a *TransferControl) Clearable(value interface{}) *TransferControl {
	a.Set("clearable", value)
	return a
}

func (a *TransferControl) Columns(value interface{}) *TransferControl {
	a.Set("columns", value)
	return a
}

func (a *TransferControl) Creatable(value interface{}) *TransferControl {
	a.Set("creatable", value)
	return a
}

func (a *TransferControl) CreateBtnLabel(value interface{}) *TransferControl {
	a.Set("createBtnLabel", value)
	return a
}

func (a *TransferControl) DeferApi(value interface{}) *TransferControl {
	a.Set("deferApi", value)
	return a
}

func (a *TransferControl) DeferField(value interface{}) *TransferControl {
	a.Set("deferField", value)
	return a
}

func (a *TransferControl) DeleteApi(value interface{}) *TransferControl {
	a.Set("deleteApi", value)
	return a
}

func (a *TransferControl) DeleteConfirmText(value interface{}) *TransferControl {
	a.Set("deleteConfirmText", value)
	return a
}

func (a *TransferControl) Delimiter(value interface{}) *TransferControl {
	a.Set("delimiter", value)
	return a
}

func (a *TransferControl) Desc(value interface{}) *TransferControl {
	a.Set("desc", value)
	return a
}

func (a *TransferControl) Description(value interface{}) *TransferControl {
	a.Set("description", value)
	return a
}

func (a *TransferControl) DescriptionClassName(value interface{}) *TransferControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *TransferControl) Disabled(value interface{}) *TransferControl {
	a.Set("disabled", value)
	return a
}

func (a *TransferControl) DisabledOn(value interface{}) *TransferControl {
	a.Set("disabledOn", value)
	return a
}

func (a *TransferControl) EditApi(value interface{}) *TransferControl {
	a.Set("editApi", value)
	return a
}

func (a *TransferControl) EditControls(value interface{}) *TransferControl {
	a.Set("editControls", value)
	return a
}

func (a *TransferControl) EditDialog(value interface{}) *TransferControl {
	a.Set("editDialog", value)
	return a
}

func (a *TransferControl) Editable(value interface{}) *TransferControl {
	a.Set("editable", value)
	return a
}

func (a *TransferControl) EditorSetting(value interface{}) *TransferControl {
	a.Set("editorSetting", value)
	return a
}

func (a *TransferControl) ExtraName(value interface{}) *TransferControl {
	a.Set("extraName", value)
	return a
}

func (a *TransferControl) ExtractValue(value interface{}) *TransferControl {
	a.Set("extractValue", value)
	return a
}

func (a *TransferControl) Hidden(value interface{}) *TransferControl {
	a.Set("hidden", value)
	return a
}

func (a *TransferControl) HiddenOn(value interface{}) *TransferControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *TransferControl) Hint(value interface{}) *TransferControl {
	a.Set("hint", value)
	return a
}

func (a *TransferControl) Horizontal(value interface{}) *TransferControl {
	a.Set("horizontal", value)
	return a
}

func (a *TransferControl) Id(value interface{}) *TransferControl {
	a.Set("id", value)
	return a
}

func (a *TransferControl) InitAutoFill(value interface{}) *TransferControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *TransferControl) InitFetch(value interface{}) *TransferControl {
	a.Set("initFetch", value)
	return a
}

func (a *TransferControl) InitFetchOn(value interface{}) *TransferControl {
	a.Set("initFetchOn", value)
	return a
}

func (a *TransferControl) InitiallyOpen(value interface{}) *TransferControl {
	a.Set("initiallyOpen", value)
	return a
}

func (a *TransferControl) Inline(value interface{}) *TransferControl {
	a.Set("inline", value)
	return a
}

func (a *TransferControl) InputClassName(value interface{}) *TransferControl {
	a.Set("inputClassName", value)
	return a
}

func (a *TransferControl) ItemHeight(value interface{}) *TransferControl {
	a.Set("itemHeight", value)
	return a
}

func (a *TransferControl) JoinValues(value interface{}) *TransferControl {
	a.Set("joinValues", value)
	return a
}

func (a *TransferControl) Label(value interface{}) *TransferControl {
	a.Set("label", value)
	return a
}

func (a *TransferControl) LabelAlign(value interface{}) *TransferControl {
	a.Set("labelAlign", value)
	return a
}

func (a *TransferControl) LabelClassName(value interface{}) *TransferControl {
	a.Set("labelClassName", value)
	return a
}

func (a *TransferControl) LabelRemark(value interface{}) *TransferControl {
	a.Set("labelRemark", value)
	return a
}

func (a *TransferControl) LabelWidth(value interface{}) *TransferControl {
	a.Set("labelWidth", value)
	return a
}

func (a *TransferControl) LeftMode(value interface{}) *TransferControl {
	a.Set("leftMode", value)
	return a
}

func (a *TransferControl) LeftOptions(value interface{}) *TransferControl {
	a.Set("leftOptions", value)
	return a
}

func (a *TransferControl) LoadingConfig(value interface{}) *TransferControl {
	a.Set("loadingConfig", value)
	return a
}

func (a *TransferControl) MenuTpl(value interface{}) *TransferControl {
	a.Set("menuTpl", value)
	return a
}

func (a *TransferControl) Mode(value interface{}) *TransferControl {
	a.Set("mode", value)
	return a
}

func (a *TransferControl) Multiple(value interface{}) *TransferControl {
	a.Set("multiple", value)
	return a
}

func (a *TransferControl) Name(value interface{}) *TransferControl {
	a.Set("name", value)
	return a
}

func (a *TransferControl) OnEvent(value interface{}) *TransferControl {
	a.Set("onEvent", value)
	return a
}

func (a *TransferControl) OnlyChildren(value interface{}) *TransferControl {
	a.Set("onlyChildren", value)
	return a
}

func (a *TransferControl) Options(value interface{}) *TransferControl {
	a.Set("options", value)
	return a
}

func (a *TransferControl) Pagination(value interface{}) *TransferControl {
	a.Set("pagination", value)
	return a
}

func (a *TransferControl) Placeholder(value interface{}) *TransferControl {
	a.Set("placeholder", value)
	return a
}

func (a *TransferControl) ReadOnly(value interface{}) *TransferControl {
	a.Set("readOnly", value)
	return a
}

func (a *TransferControl) ReadOnlyOn(value interface{}) *TransferControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *TransferControl) Remark(value interface{}) *TransferControl {
	a.Set("remark", value)
	return a
}

func (a *TransferControl) Removable(value interface{}) *TransferControl {
	a.Set("removable", value)
	return a
}

func (a *TransferControl) Required(value interface{}) *TransferControl {
	a.Set("required", value)
	return a
}

func (a *TransferControl) ResetValue(value interface{}) *TransferControl {
	a.Set("resetValue", value)
	return a
}

func (a *TransferControl) ResultListModeFollowSelect(value interface{}) *TransferControl {
	a.Set("resultListModeFollowSelect", value)
	return a
}

func (a *TransferControl) ResultSearchPlaceholder(value interface{}) *TransferControl {
	a.Set("resultSearchPlaceholder", value)
	return a
}

func (a *TransferControl) ResultSearchable(value interface{}) *TransferControl {
	a.Set("resultSearchable", value)
	return a
}

func (a *TransferControl) ResultTitle(value interface{}) *TransferControl {
	a.Set("resultTitle", value)
	return a
}

func (a *TransferControl) RightMode(value interface{}) *TransferControl {
	a.Set("rightMode", value)
	return a
}

func (a *TransferControl) Row(value interface{}) *TransferControl {
	a.Set("row", value)
	return a
}

func (a *TransferControl) SaveImmediately(value interface{}) *TransferControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *TransferControl) SearchApi(value interface{}) *TransferControl {
	a.Set("searchApi", value)
	return a
}

func (a *TransferControl) SearchPlaceholder(value interface{}) *TransferControl {
	a.Set("searchPlaceholder", value)
	return a
}

func (a *TransferControl) SearchResultColumns(value interface{}) *TransferControl {
	a.Set("searchResultColumns", value)
	return a
}

func (a *TransferControl) SearchResultMode(value interface{}) *TransferControl {
	a.Set("searchResultMode", value)
	return a
}

func (a *TransferControl) Searchable(value interface{}) *TransferControl {
	a.Set("searchable", value)
	return a
}

func (a *TransferControl) SelectFirst(value interface{}) *TransferControl {
	a.Set("selectFirst", value)
	return a
}

func (a *TransferControl) SelectMode(value interface{}) *TransferControl {
	a.Set("selectMode", value)
	return a
}

func (a *TransferControl) SelectTitle(value interface{}) *TransferControl {
	a.Set("selectTitle", value)
	return a
}

func (a *TransferControl) ShowArrow(value interface{}) *TransferControl {
	a.Set("showArrow", value)
	return a
}

func (a *TransferControl) ShowInvalidMatch(value interface{}) *TransferControl {
	a.Set("showInvalidMatch", value)
	return a
}

func (a *TransferControl) Size(value interface{}) *TransferControl {
	a.Set("size", value)
	return a
}

func (a *TransferControl) Sortable(value interface{}) *TransferControl {
	a.Set("sortable", value)
	return a
}

func (a *TransferControl) Source(value interface{}) *TransferControl {
	a.Set("source", value)
	return a
}

func (a *TransferControl) Static(value interface{}) *TransferControl {
	a.Set("static", value)
	return a
}

func (a *TransferControl) StaticClassName(value interface{}) *TransferControl {
	a.Set("staticClassName", value)
	return a
}

func (a *TransferControl) StaticInputClassName(value interface{}) *TransferControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *TransferControl) StaticLabelClassName(value interface{}) *TransferControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *TransferControl) StaticOn(value interface{}) *TransferControl {
	a.Set("staticOn", value)
	return a
}

func (a *TransferControl) StaticPlaceholder(value interface{}) *TransferControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *TransferControl) StaticSchema(value interface{}) *TransferControl {
	a.Set("staticSchema", value)
	return a
}

func (a *TransferControl) Statistics(value interface{}) *TransferControl {
	a.Set("statistics", value)
	return a
}

func (a *TransferControl) Style(value interface{}) *TransferControl {
	a.Set("style", value)
	return a
}

func (a *TransferControl) SubmitOnChange(value interface{}) *TransferControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *TransferControl) TestIdBuilder(value interface{}) *TransferControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *TransferControl) Type(value interface{}) *TransferControl {
	a.Set("type", value)
	return a
}

func (a *TransferControl) UseMobileUI(value interface{}) *TransferControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *TransferControl) ValidateApi(value interface{}) *TransferControl {
	a.Set("validateApi", value)
	return a
}

func (a *TransferControl) ValidateOnChange(value interface{}) *TransferControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *TransferControl) ValidationErrors(value interface{}) *TransferControl {
	a.Set("validationErrors", value)
	return a
}

func (a *TransferControl) Validations(value interface{}) *TransferControl {
	a.Set("validations", value)
	return a
}

func (a *TransferControl) Value(value interface{}) *TransferControl {
	a.Set("value", value)
	return a
}

func (a *TransferControl) ValueTpl(value interface{}) *TransferControl {
	a.Set("valueTpl", value)
	return a
}

func (a *TransferControl) ValuesNoWrap(value interface{}) *TransferControl {
	a.Set("valuesNoWrap", value)
	return a
}

func (a *TransferControl) VirtualThreshold(value interface{}) *TransferControl {
	a.Set("virtualThreshold", value)
	return a
}

func (a *TransferControl) Visible(value interface{}) *TransferControl {
	a.Set("visible", value)
	return a
}

func (a *TransferControl) VisibleOn(value interface{}) *TransferControl {
	a.Set("visibleOn", value)
	return a
}

func (a *TransferControl) Width(value interface{}) *TransferControl {
	a.Set("width", value)
	return a
}
