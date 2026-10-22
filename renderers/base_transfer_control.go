package renderers

type BaseTransferControl struct {
	*BaseRenderer
}

func NewBaseTransferControl() *BaseTransferControl {
	a := &BaseTransferControl{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *BaseTransferControl) Set(name string, value interface{}) *BaseTransferControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *BaseTransferControl) AddApi(value interface{}) *BaseTransferControl {
	a.Set("addApi", value)
	return a
}

func (a *BaseTransferControl) AddControls(value interface{}) *BaseTransferControl {
	a.Set("addControls", value)
	return a
}

func (a *BaseTransferControl) AddDialog(value interface{}) *BaseTransferControl {
	a.Set("addDialog", value)
	return a
}

func (a *BaseTransferControl) AutoCheckChildren(value interface{}) *BaseTransferControl {
	a.Set("autoCheckChildren", value)
	return a
}

func (a *BaseTransferControl) AutoFill(value interface{}) *BaseTransferControl {
	a.Set("autoFill", value)
	return a
}

func (a *BaseTransferControl) CheckAll(value interface{}) *BaseTransferControl {
	a.Set("checkAll", value)
	return a
}

func (a *BaseTransferControl) ClassName(value interface{}) *BaseTransferControl {
	a.Set("className", value)
	return a
}

func (a *BaseTransferControl) ClearValueOnHidden(value interface{}) *BaseTransferControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *BaseTransferControl) ClearValueOnSourceChange(value interface{}) *BaseTransferControl {
	a.Set("clearValueOnSourceChange", value)
	return a
}

func (a *BaseTransferControl) Clearable(value interface{}) *BaseTransferControl {
	a.Set("clearable", value)
	return a
}

func (a *BaseTransferControl) Columns(value interface{}) *BaseTransferControl {
	a.Set("columns", value)
	return a
}

func (a *BaseTransferControl) Creatable(value interface{}) *BaseTransferControl {
	a.Set("creatable", value)
	return a
}

func (a *BaseTransferControl) CreateBtnLabel(value interface{}) *BaseTransferControl {
	a.Set("createBtnLabel", value)
	return a
}

func (a *BaseTransferControl) DeferApi(value interface{}) *BaseTransferControl {
	a.Set("deferApi", value)
	return a
}

func (a *BaseTransferControl) DeferField(value interface{}) *BaseTransferControl {
	a.Set("deferField", value)
	return a
}

func (a *BaseTransferControl) DeleteApi(value interface{}) *BaseTransferControl {
	a.Set("deleteApi", value)
	return a
}

func (a *BaseTransferControl) DeleteConfirmText(value interface{}) *BaseTransferControl {
	a.Set("deleteConfirmText", value)
	return a
}

func (a *BaseTransferControl) Delimiter(value interface{}) *BaseTransferControl {
	a.Set("delimiter", value)
	return a
}

func (a *BaseTransferControl) Desc(value interface{}) *BaseTransferControl {
	a.Set("desc", value)
	return a
}

func (a *BaseTransferControl) Description(value interface{}) *BaseTransferControl {
	a.Set("description", value)
	return a
}

func (a *BaseTransferControl) DescriptionClassName(value interface{}) *BaseTransferControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *BaseTransferControl) Disabled(value interface{}) *BaseTransferControl {
	a.Set("disabled", value)
	return a
}

func (a *BaseTransferControl) DisabledOn(value interface{}) *BaseTransferControl {
	a.Set("disabledOn", value)
	return a
}

func (a *BaseTransferControl) EditApi(value interface{}) *BaseTransferControl {
	a.Set("editApi", value)
	return a
}

func (a *BaseTransferControl) EditControls(value interface{}) *BaseTransferControl {
	a.Set("editControls", value)
	return a
}

func (a *BaseTransferControl) EditDialog(value interface{}) *BaseTransferControl {
	a.Set("editDialog", value)
	return a
}

func (a *BaseTransferControl) Editable(value interface{}) *BaseTransferControl {
	a.Set("editable", value)
	return a
}

func (a *BaseTransferControl) EditorSetting(value interface{}) *BaseTransferControl {
	a.Set("editorSetting", value)
	return a
}

func (a *BaseTransferControl) ExtraName(value interface{}) *BaseTransferControl {
	a.Set("extraName", value)
	return a
}

func (a *BaseTransferControl) ExtractValue(value interface{}) *BaseTransferControl {
	a.Set("extractValue", value)
	return a
}

func (a *BaseTransferControl) Hidden(value interface{}) *BaseTransferControl {
	a.Set("hidden", value)
	return a
}

func (a *BaseTransferControl) HiddenOn(value interface{}) *BaseTransferControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *BaseTransferControl) Hint(value interface{}) *BaseTransferControl {
	a.Set("hint", value)
	return a
}

func (a *BaseTransferControl) Horizontal(value interface{}) *BaseTransferControl {
	a.Set("horizontal", value)
	return a
}

func (a *BaseTransferControl) Id(value interface{}) *BaseTransferControl {
	a.Set("id", value)
	return a
}

func (a *BaseTransferControl) InitAutoFill(value interface{}) *BaseTransferControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *BaseTransferControl) InitFetch(value interface{}) *BaseTransferControl {
	a.Set("initFetch", value)
	return a
}

func (a *BaseTransferControl) InitFetchOn(value interface{}) *BaseTransferControl {
	a.Set("initFetchOn", value)
	return a
}

func (a *BaseTransferControl) InitiallyOpen(value interface{}) *BaseTransferControl {
	a.Set("initiallyOpen", value)
	return a
}

func (a *BaseTransferControl) Inline(value interface{}) *BaseTransferControl {
	a.Set("inline", value)
	return a
}

func (a *BaseTransferControl) InputClassName(value interface{}) *BaseTransferControl {
	a.Set("inputClassName", value)
	return a
}

func (a *BaseTransferControl) ItemHeight(value interface{}) *BaseTransferControl {
	a.Set("itemHeight", value)
	return a
}

func (a *BaseTransferControl) JoinValues(value interface{}) *BaseTransferControl {
	a.Set("joinValues", value)
	return a
}

func (a *BaseTransferControl) Label(value interface{}) *BaseTransferControl {
	a.Set("label", value)
	return a
}

func (a *BaseTransferControl) LabelAlign(value interface{}) *BaseTransferControl {
	a.Set("labelAlign", value)
	return a
}

func (a *BaseTransferControl) LabelClassName(value interface{}) *BaseTransferControl {
	a.Set("labelClassName", value)
	return a
}

func (a *BaseTransferControl) LabelOverflow(value interface{}) *BaseTransferControl {
	a.Set("labelOverflow", value)
	return a
}

func (a *BaseTransferControl) LabelRemark(value interface{}) *BaseTransferControl {
	a.Set("labelRemark", value)
	return a
}

func (a *BaseTransferControl) LabelWidth(value interface{}) *BaseTransferControl {
	a.Set("labelWidth", value)
	return a
}

func (a *BaseTransferControl) LeftMode(value interface{}) *BaseTransferControl {
	a.Set("leftMode", value)
	return a
}

func (a *BaseTransferControl) LeftOptions(value interface{}) *BaseTransferControl {
	a.Set("leftOptions", value)
	return a
}

func (a *BaseTransferControl) LoadingConfig(value interface{}) *BaseTransferControl {
	a.Set("loadingConfig", value)
	return a
}

func (a *BaseTransferControl) MenuTpl(value interface{}) *BaseTransferControl {
	a.Set("menuTpl", value)
	return a
}

func (a *BaseTransferControl) Mode(value interface{}) *BaseTransferControl {
	a.Set("mode", value)
	return a
}

func (a *BaseTransferControl) Multiple(value interface{}) *BaseTransferControl {
	a.Set("multiple", value)
	return a
}

func (a *BaseTransferControl) Name(value interface{}) *BaseTransferControl {
	a.Set("name", value)
	return a
}

func (a *BaseTransferControl) OnEvent(value interface{}) *BaseTransferControl {
	a.Set("onEvent", value)
	return a
}

func (a *BaseTransferControl) OnlyChildren(value interface{}) *BaseTransferControl {
	a.Set("onlyChildren", value)
	return a
}

func (a *BaseTransferControl) Options(value interface{}) *BaseTransferControl {
	a.Set("options", value)
	return a
}

func (a *BaseTransferControl) Pagination(value interface{}) *BaseTransferControl {
	a.Set("pagination", value)
	return a
}

func (a *BaseTransferControl) Placeholder(value interface{}) *BaseTransferControl {
	a.Set("placeholder", value)
	return a
}

func (a *BaseTransferControl) ReadOnly(value interface{}) *BaseTransferControl {
	a.Set("readOnly", value)
	return a
}

func (a *BaseTransferControl) ReadOnlyOn(value interface{}) *BaseTransferControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *BaseTransferControl) Remark(value interface{}) *BaseTransferControl {
	a.Set("remark", value)
	return a
}

func (a *BaseTransferControl) Removable(value interface{}) *BaseTransferControl {
	a.Set("removable", value)
	return a
}

func (a *BaseTransferControl) Required(value interface{}) *BaseTransferControl {
	a.Set("required", value)
	return a
}

func (a *BaseTransferControl) ResetValue(value interface{}) *BaseTransferControl {
	a.Set("resetValue", value)
	return a
}

func (a *BaseTransferControl) ResultListModeFollowSelect(value interface{}) *BaseTransferControl {
	a.Set("resultListModeFollowSelect", value)
	return a
}

func (a *BaseTransferControl) ResultSearchPlaceholder(value interface{}) *BaseTransferControl {
	a.Set("resultSearchPlaceholder", value)
	return a
}

func (a *BaseTransferControl) ResultSearchable(value interface{}) *BaseTransferControl {
	a.Set("resultSearchable", value)
	return a
}

func (a *BaseTransferControl) ResultTitle(value interface{}) *BaseTransferControl {
	a.Set("resultTitle", value)
	return a
}

func (a *BaseTransferControl) RightMode(value interface{}) *BaseTransferControl {
	a.Set("rightMode", value)
	return a
}

func (a *BaseTransferControl) Row(value interface{}) *BaseTransferControl {
	a.Set("row", value)
	return a
}

func (a *BaseTransferControl) SaveImmediately(value interface{}) *BaseTransferControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *BaseTransferControl) SearchApi(value interface{}) *BaseTransferControl {
	a.Set("searchApi", value)
	return a
}

func (a *BaseTransferControl) SearchPlaceholder(value interface{}) *BaseTransferControl {
	a.Set("searchPlaceholder", value)
	return a
}

func (a *BaseTransferControl) SearchResultColumns(value interface{}) *BaseTransferControl {
	a.Set("searchResultColumns", value)
	return a
}

func (a *BaseTransferControl) SearchResultMode(value interface{}) *BaseTransferControl {
	a.Set("searchResultMode", value)
	return a
}

func (a *BaseTransferControl) Searchable(value interface{}) *BaseTransferControl {
	a.Set("searchable", value)
	return a
}

func (a *BaseTransferControl) SelectFirst(value interface{}) *BaseTransferControl {
	a.Set("selectFirst", value)
	return a
}

func (a *BaseTransferControl) SelectMode(value interface{}) *BaseTransferControl {
	a.Set("selectMode", value)
	return a
}

func (a *BaseTransferControl) SelectTitle(value interface{}) *BaseTransferControl {
	a.Set("selectTitle", value)
	return a
}

func (a *BaseTransferControl) ShowArrow(value interface{}) *BaseTransferControl {
	a.Set("showArrow", value)
	return a
}

func (a *BaseTransferControl) ShowInvalidMatch(value interface{}) *BaseTransferControl {
	a.Set("showInvalidMatch", value)
	return a
}

func (a *BaseTransferControl) Size(value interface{}) *BaseTransferControl {
	a.Set("size", value)
	return a
}

func (a *BaseTransferControl) Sortable(value interface{}) *BaseTransferControl {
	a.Set("sortable", value)
	return a
}

func (a *BaseTransferControl) Source(value interface{}) *BaseTransferControl {
	a.Set("source", value)
	return a
}

func (a *BaseTransferControl) Static(value interface{}) *BaseTransferControl {
	a.Set("static", value)
	return a
}

func (a *BaseTransferControl) StaticClassName(value interface{}) *BaseTransferControl {
	a.Set("staticClassName", value)
	return a
}

func (a *BaseTransferControl) StaticInputClassName(value interface{}) *BaseTransferControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *BaseTransferControl) StaticLabelClassName(value interface{}) *BaseTransferControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *BaseTransferControl) StaticOn(value interface{}) *BaseTransferControl {
	a.Set("staticOn", value)
	return a
}

func (a *BaseTransferControl) StaticPlaceholder(value interface{}) *BaseTransferControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *BaseTransferControl) StaticSchema(value interface{}) *BaseTransferControl {
	a.Set("staticSchema", value)
	return a
}

func (a *BaseTransferControl) Statistics(value interface{}) *BaseTransferControl {
	a.Set("statistics", value)
	return a
}

func (a *BaseTransferControl) Style(value interface{}) *BaseTransferControl {
	a.Set("style", value)
	return a
}

func (a *BaseTransferControl) SubmitOnChange(value interface{}) *BaseTransferControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *BaseTransferControl) Type(value interface{}) *BaseTransferControl {
	a.Set("type", value)
	return a
}

func (a *BaseTransferControl) UseMobileUI(value interface{}) *BaseTransferControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *BaseTransferControl) ValidateApi(value interface{}) *BaseTransferControl {
	a.Set("validateApi", value)
	return a
}

func (a *BaseTransferControl) ValidateOnChange(value interface{}) *BaseTransferControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *BaseTransferControl) ValidationErrors(value interface{}) *BaseTransferControl {
	a.Set("validationErrors", value)
	return a
}

func (a *BaseTransferControl) Validations(value interface{}) *BaseTransferControl {
	a.Set("validations", value)
	return a
}

func (a *BaseTransferControl) Value(value interface{}) *BaseTransferControl {
	a.Set("value", value)
	return a
}

func (a *BaseTransferControl) ValueTpl(value interface{}) *BaseTransferControl {
	a.Set("valueTpl", value)
	return a
}

func (a *BaseTransferControl) ValuesNoWrap(value interface{}) *BaseTransferControl {
	a.Set("valuesNoWrap", value)
	return a
}

func (a *BaseTransferControl) VirtualThreshold(value interface{}) *BaseTransferControl {
	a.Set("virtualThreshold", value)
	return a
}

func (a *BaseTransferControl) Visible(value interface{}) *BaseTransferControl {
	a.Set("visible", value)
	return a
}

func (a *BaseTransferControl) VisibleOn(value interface{}) *BaseTransferControl {
	a.Set("visibleOn", value)
	return a
}

func (a *BaseTransferControl) Width(value interface{}) *BaseTransferControl {
	a.Set("width", value)
	return a
}
