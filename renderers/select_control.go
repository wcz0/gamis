package renderers

type SelectControl struct {
	*BaseRenderer
}

func NewSelectControl() *SelectControl {
	a := &SelectControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "select")
	return a
}

func (a *SelectControl) Set(name string, value interface{}) *SelectControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *SelectControl) AddApi(value interface{}) *SelectControl {
	a.Set("addApi", value)
	return a
}

func (a *SelectControl) AddControls(value interface{}) *SelectControl {
	a.Set("addControls", value)
	return a
}

func (a *SelectControl) AddDialog(value interface{}) *SelectControl {
	a.Set("addDialog", value)
	return a
}

func (a *SelectControl) AutoCheckChildren(value interface{}) *SelectControl {
	a.Set("autoCheckChildren", value)
	return a
}

func (a *SelectControl) AutoComplete(value interface{}) *SelectControl {
	a.Set("autoComplete", value)
	return a
}

func (a *SelectControl) AutoFill(value interface{}) *SelectControl {
	a.Set("autoFill", value)
	return a
}

func (a *SelectControl) BorderMode(value interface{}) *SelectControl {
	a.Set("borderMode", value)
	return a
}

func (a *SelectControl) CheckAll(value interface{}) *SelectControl {
	a.Set("checkAll", value)
	return a
}

func (a *SelectControl) CheckAllLabel(value interface{}) *SelectControl {
	a.Set("checkAllLabel", value)
	return a
}

func (a *SelectControl) ClassName(value interface{}) *SelectControl {
	a.Set("className", value)
	return a
}

func (a *SelectControl) ClearValueOnHidden(value interface{}) *SelectControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *SelectControl) ClearValueOnSourceChange(value interface{}) *SelectControl {
	a.Set("clearValueOnSourceChange", value)
	return a
}

func (a *SelectControl) Clearable(value interface{}) *SelectControl {
	a.Set("clearable", value)
	return a
}

func (a *SelectControl) Columns(value interface{}) *SelectControl {
	a.Set("columns", value)
	return a
}

func (a *SelectControl) Creatable(value interface{}) *SelectControl {
	a.Set("creatable", value)
	return a
}

func (a *SelectControl) CreateBtnLabel(value interface{}) *SelectControl {
	a.Set("createBtnLabel", value)
	return a
}

func (a *SelectControl) DefaultCheckAll(value interface{}) *SelectControl {
	a.Set("defaultCheckAll", value)
	return a
}

func (a *SelectControl) DeferApi(value interface{}) *SelectControl {
	a.Set("deferApi", value)
	return a
}

func (a *SelectControl) DeferField(value interface{}) *SelectControl {
	a.Set("deferField", value)
	return a
}

func (a *SelectControl) DeleteApi(value interface{}) *SelectControl {
	a.Set("deleteApi", value)
	return a
}

func (a *SelectControl) DeleteConfirmText(value interface{}) *SelectControl {
	a.Set("deleteConfirmText", value)
	return a
}

func (a *SelectControl) Delimiter(value interface{}) *SelectControl {
	a.Set("delimiter", value)
	return a
}

func (a *SelectControl) Desc(value interface{}) *SelectControl {
	a.Set("desc", value)
	return a
}

func (a *SelectControl) Description(value interface{}) *SelectControl {
	a.Set("description", value)
	return a
}

func (a *SelectControl) DescriptionClassName(value interface{}) *SelectControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *SelectControl) Disabled(value interface{}) *SelectControl {
	a.Set("disabled", value)
	return a
}

func (a *SelectControl) DisabledOn(value interface{}) *SelectControl {
	a.Set("disabledOn", value)
	return a
}

func (a *SelectControl) EditApi(value interface{}) *SelectControl {
	a.Set("editApi", value)
	return a
}

func (a *SelectControl) EditControls(value interface{}) *SelectControl {
	a.Set("editControls", value)
	return a
}

func (a *SelectControl) EditDialog(value interface{}) *SelectControl {
	a.Set("editDialog", value)
	return a
}

func (a *SelectControl) Editable(value interface{}) *SelectControl {
	a.Set("editable", value)
	return a
}

func (a *SelectControl) EditorSetting(value interface{}) *SelectControl {
	a.Set("editorSetting", value)
	return a
}

func (a *SelectControl) ExtraName(value interface{}) *SelectControl {
	a.Set("extraName", value)
	return a
}

func (a *SelectControl) ExtractValue(value interface{}) *SelectControl {
	a.Set("extractValue", value)
	return a
}

func (a *SelectControl) Hidden(value interface{}) *SelectControl {
	a.Set("hidden", value)
	return a
}

func (a *SelectControl) HiddenOn(value interface{}) *SelectControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *SelectControl) Hint(value interface{}) *SelectControl {
	a.Set("hint", value)
	return a
}

func (a *SelectControl) Horizontal(value interface{}) *SelectControl {
	a.Set("horizontal", value)
	return a
}

func (a *SelectControl) Id(value interface{}) *SelectControl {
	a.Set("id", value)
	return a
}

func (a *SelectControl) InitAutoFill(value interface{}) *SelectControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *SelectControl) InitFetch(value interface{}) *SelectControl {
	a.Set("initFetch", value)
	return a
}

func (a *SelectControl) InitFetchOn(value interface{}) *SelectControl {
	a.Set("initFetchOn", value)
	return a
}

func (a *SelectControl) Inline(value interface{}) *SelectControl {
	a.Set("inline", value)
	return a
}

func (a *SelectControl) InputClassName(value interface{}) *SelectControl {
	a.Set("inputClassName", value)
	return a
}

func (a *SelectControl) ItemHeight(value interface{}) *SelectControl {
	a.Set("itemHeight", value)
	return a
}

func (a *SelectControl) JoinValues(value interface{}) *SelectControl {
	a.Set("joinValues", value)
	return a
}

func (a *SelectControl) Label(value interface{}) *SelectControl {
	a.Set("label", value)
	return a
}

func (a *SelectControl) LabelAlign(value interface{}) *SelectControl {
	a.Set("labelAlign", value)
	return a
}

func (a *SelectControl) LabelClassName(value interface{}) *SelectControl {
	a.Set("labelClassName", value)
	return a
}

func (a *SelectControl) LabelField(value interface{}) *SelectControl {
	a.Set("labelField", value)
	return a
}

func (a *SelectControl) LabelRemark(value interface{}) *SelectControl {
	a.Set("labelRemark", value)
	return a
}

func (a *SelectControl) LabelWidth(value interface{}) *SelectControl {
	a.Set("labelWidth", value)
	return a
}

func (a *SelectControl) LeftMode(value interface{}) *SelectControl {
	a.Set("leftMode", value)
	return a
}

func (a *SelectControl) LeftOptions(value interface{}) *SelectControl {
	a.Set("leftOptions", value)
	return a
}

func (a *SelectControl) LoadingConfig(value interface{}) *SelectControl {
	a.Set("loadingConfig", value)
	return a
}

func (a *SelectControl) MaxTagCount(value interface{}) *SelectControl {
	a.Set("maxTagCount", value)
	return a
}

func (a *SelectControl) MenuTpl(value interface{}) *SelectControl {
	a.Set("menuTpl", value)
	return a
}

func (a *SelectControl) Mode(value interface{}) *SelectControl {
	a.Set("mode", value)
	return a
}

func (a *SelectControl) Multiple(value interface{}) *SelectControl {
	a.Set("multiple", value)
	return a
}

func (a *SelectControl) Name(value interface{}) *SelectControl {
	a.Set("name", value)
	return a
}

func (a *SelectControl) OnEvent(value interface{}) *SelectControl {
	a.Set("onEvent", value)
	return a
}

func (a *SelectControl) OptionClassName(value interface{}) *SelectControl {
	a.Set("optionClassName", value)
	return a
}

func (a *SelectControl) Options(value interface{}) *SelectControl {
	a.Set("options", value)
	return a
}

func (a *SelectControl) OverflowTagPopover(value interface{}) *SelectControl {
	a.Set("overflowTagPopover", value)
	return a
}

func (a *SelectControl) Overlay(value interface{}) *SelectControl {
	a.Set("overlay", value)
	return a
}

func (a *SelectControl) Placeholder(value interface{}) *SelectControl {
	a.Set("placeholder", value)
	return a
}

func (a *SelectControl) ReadOnly(value interface{}) *SelectControl {
	a.Set("readOnly", value)
	return a
}

func (a *SelectControl) ReadOnlyOn(value interface{}) *SelectControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *SelectControl) Remark(value interface{}) *SelectControl {
	a.Set("remark", value)
	return a
}

func (a *SelectControl) Removable(value interface{}) *SelectControl {
	a.Set("removable", value)
	return a
}

func (a *SelectControl) Required(value interface{}) *SelectControl {
	a.Set("required", value)
	return a
}

func (a *SelectControl) ResetValue(value interface{}) *SelectControl {
	a.Set("resetValue", value)
	return a
}

func (a *SelectControl) RightMode(value interface{}) *SelectControl {
	a.Set("rightMode", value)
	return a
}

func (a *SelectControl) Row(value interface{}) *SelectControl {
	a.Set("row", value)
	return a
}

func (a *SelectControl) SaveImmediately(value interface{}) *SelectControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *SelectControl) SearchApi(value interface{}) *SelectControl {
	a.Set("searchApi", value)
	return a
}

func (a *SelectControl) SearchResultColumns(value interface{}) *SelectControl {
	a.Set("searchResultColumns", value)
	return a
}

func (a *SelectControl) SearchResultMode(value interface{}) *SelectControl {
	a.Set("searchResultMode", value)
	return a
}

func (a *SelectControl) Searchable(value interface{}) *SelectControl {
	a.Set("searchable", value)
	return a
}

func (a *SelectControl) SelectFirst(value interface{}) *SelectControl {
	a.Set("selectFirst", value)
	return a
}

func (a *SelectControl) SelectMode(value interface{}) *SelectControl {
	a.Set("selectMode", value)
	return a
}

func (a *SelectControl) ShowInvalidMatch(value interface{}) *SelectControl {
	a.Set("showInvalidMatch", value)
	return a
}

func (a *SelectControl) Size(value interface{}) *SelectControl {
	a.Set("size", value)
	return a
}

func (a *SelectControl) Source(value interface{}) *SelectControl {
	a.Set("source", value)
	return a
}

func (a *SelectControl) Static(value interface{}) *SelectControl {
	a.Set("static", value)
	return a
}

func (a *SelectControl) StaticClassName(value interface{}) *SelectControl {
	a.Set("staticClassName", value)
	return a
}

func (a *SelectControl) StaticInputClassName(value interface{}) *SelectControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *SelectControl) StaticLabelClassName(value interface{}) *SelectControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *SelectControl) StaticOn(value interface{}) *SelectControl {
	a.Set("staticOn", value)
	return a
}

func (a *SelectControl) StaticPlaceholder(value interface{}) *SelectControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *SelectControl) StaticSchema(value interface{}) *SelectControl {
	a.Set("staticSchema", value)
	return a
}

func (a *SelectControl) Style(value interface{}) *SelectControl {
	a.Set("style", value)
	return a
}

func (a *SelectControl) SubmitOnChange(value interface{}) *SelectControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *SelectControl) TestIdBuilder(value interface{}) *SelectControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *SelectControl) Type(value interface{}) *SelectControl {
	a.Set("type", value)
	return a
}

func (a *SelectControl) UseMobileUI(value interface{}) *SelectControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *SelectControl) ValidateApi(value interface{}) *SelectControl {
	a.Set("validateApi", value)
	return a
}

func (a *SelectControl) ValidateOnChange(value interface{}) *SelectControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *SelectControl) ValidationErrors(value interface{}) *SelectControl {
	a.Set("validationErrors", value)
	return a
}

func (a *SelectControl) Validations(value interface{}) *SelectControl {
	a.Set("validations", value)
	return a
}

func (a *SelectControl) Value(value interface{}) *SelectControl {
	a.Set("value", value)
	return a
}

func (a *SelectControl) ValueField(value interface{}) *SelectControl {
	a.Set("valueField", value)
	return a
}

func (a *SelectControl) ValuesNoWrap(value interface{}) *SelectControl {
	a.Set("valuesNoWrap", value)
	return a
}

func (a *SelectControl) VirtualThreshold(value interface{}) *SelectControl {
	a.Set("virtualThreshold", value)
	return a
}

func (a *SelectControl) Visible(value interface{}) *SelectControl {
	a.Set("visible", value)
	return a
}

func (a *SelectControl) VisibleOn(value interface{}) *SelectControl {
	a.Set("visibleOn", value)
	return a
}

func (a *SelectControl) Width(value interface{}) *SelectControl {
	a.Set("width", value)
	return a
}
