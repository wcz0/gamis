package renderers

type TreeSelectControl struct {
	*BaseRenderer
}

func NewTreeSelectControl() *TreeSelectControl {
	a := &TreeSelectControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "tree-select")
	return a
}

func (a *TreeSelectControl) Set(name string, value interface{}) *TreeSelectControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *TreeSelectControl) AddApi(value interface{}) *TreeSelectControl {
	a.Set("addApi", value)
	return a
}

func (a *TreeSelectControl) AddControls(value interface{}) *TreeSelectControl {
	a.Set("addControls", value)
	return a
}

func (a *TreeSelectControl) AddDialog(value interface{}) *TreeSelectControl {
	a.Set("addDialog", value)
	return a
}

func (a *TreeSelectControl) AutoCancelParent(value interface{}) *TreeSelectControl {
	a.Set("autoCancelParent", value)
	return a
}

func (a *TreeSelectControl) AutoCheckChildren(value interface{}) *TreeSelectControl {
	a.Set("autoCheckChildren", value)
	return a
}

func (a *TreeSelectControl) AutoFill(value interface{}) *TreeSelectControl {
	a.Set("autoFill", value)
	return a
}

func (a *TreeSelectControl) Cascade(value interface{}) *TreeSelectControl {
	a.Set("cascade", value)
	return a
}

func (a *TreeSelectControl) ClassName(value interface{}) *TreeSelectControl {
	a.Set("className", value)
	return a
}

func (a *TreeSelectControl) ClearValueOnHidden(value interface{}) *TreeSelectControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *TreeSelectControl) ClearValueOnSourceChange(value interface{}) *TreeSelectControl {
	a.Set("clearValueOnSourceChange", value)
	return a
}

func (a *TreeSelectControl) Clearable(value interface{}) *TreeSelectControl {
	a.Set("clearable", value)
	return a
}

func (a *TreeSelectControl) Creatable(value interface{}) *TreeSelectControl {
	a.Set("creatable", value)
	return a
}

func (a *TreeSelectControl) CreateBtnLabel(value interface{}) *TreeSelectControl {
	a.Set("createBtnLabel", value)
	return a
}

func (a *TreeSelectControl) DeferApi(value interface{}) *TreeSelectControl {
	a.Set("deferApi", value)
	return a
}

func (a *TreeSelectControl) DeferField(value interface{}) *TreeSelectControl {
	a.Set("deferField", value)
	return a
}

func (a *TreeSelectControl) DeleteApi(value interface{}) *TreeSelectControl {
	a.Set("deleteApi", value)
	return a
}

func (a *TreeSelectControl) DeleteConfirmText(value interface{}) *TreeSelectControl {
	a.Set("deleteConfirmText", value)
	return a
}

func (a *TreeSelectControl) Delimiter(value interface{}) *TreeSelectControl {
	a.Set("delimiter", value)
	return a
}

func (a *TreeSelectControl) Desc(value interface{}) *TreeSelectControl {
	a.Set("desc", value)
	return a
}

func (a *TreeSelectControl) Description(value interface{}) *TreeSelectControl {
	a.Set("description", value)
	return a
}

func (a *TreeSelectControl) DescriptionClassName(value interface{}) *TreeSelectControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *TreeSelectControl) Disabled(value interface{}) *TreeSelectControl {
	a.Set("disabled", value)
	return a
}

func (a *TreeSelectControl) DisabledOn(value interface{}) *TreeSelectControl {
	a.Set("disabledOn", value)
	return a
}

func (a *TreeSelectControl) EditApi(value interface{}) *TreeSelectControl {
	a.Set("editApi", value)
	return a
}

func (a *TreeSelectControl) EditControls(value interface{}) *TreeSelectControl {
	a.Set("editControls", value)
	return a
}

func (a *TreeSelectControl) EditDialog(value interface{}) *TreeSelectControl {
	a.Set("editDialog", value)
	return a
}

func (a *TreeSelectControl) Editable(value interface{}) *TreeSelectControl {
	a.Set("editable", value)
	return a
}

func (a *TreeSelectControl) EditorSetting(value interface{}) *TreeSelectControl {
	a.Set("editorSetting", value)
	return a
}

func (a *TreeSelectControl) EnableDefaultIcon(value interface{}) *TreeSelectControl {
	a.Set("enableDefaultIcon", value)
	return a
}

func (a *TreeSelectControl) EnableNodePath(value interface{}) *TreeSelectControl {
	a.Set("enableNodePath", value)
	return a
}

func (a *TreeSelectControl) ExtraName(value interface{}) *TreeSelectControl {
	a.Set("extraName", value)
	return a
}

func (a *TreeSelectControl) ExtractValue(value interface{}) *TreeSelectControl {
	a.Set("extractValue", value)
	return a
}

func (a *TreeSelectControl) Hidden(value interface{}) *TreeSelectControl {
	a.Set("hidden", value)
	return a
}

func (a *TreeSelectControl) HiddenOn(value interface{}) *TreeSelectControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *TreeSelectControl) HideNodePathLabel(value interface{}) *TreeSelectControl {
	a.Set("hideNodePathLabel", value)
	return a
}

func (a *TreeSelectControl) HideRoot(value interface{}) *TreeSelectControl {
	a.Set("hideRoot", value)
	return a
}

func (a *TreeSelectControl) Hint(value interface{}) *TreeSelectControl {
	a.Set("hint", value)
	return a
}

func (a *TreeSelectControl) Horizontal(value interface{}) *TreeSelectControl {
	a.Set("horizontal", value)
	return a
}

func (a *TreeSelectControl) Id(value interface{}) *TreeSelectControl {
	a.Set("id", value)
	return a
}

func (a *TreeSelectControl) InitAutoFill(value interface{}) *TreeSelectControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *TreeSelectControl) InitFetch(value interface{}) *TreeSelectControl {
	a.Set("initFetch", value)
	return a
}

func (a *TreeSelectControl) InitFetchOn(value interface{}) *TreeSelectControl {
	a.Set("initFetchOn", value)
	return a
}

func (a *TreeSelectControl) Inline(value interface{}) *TreeSelectControl {
	a.Set("inline", value)
	return a
}

func (a *TreeSelectControl) InputClassName(value interface{}) *TreeSelectControl {
	a.Set("inputClassName", value)
	return a
}

func (a *TreeSelectControl) ItemActions(value interface{}) *TreeSelectControl {
	a.Set("itemActions", value)
	return a
}

func (a *TreeSelectControl) JoinValues(value interface{}) *TreeSelectControl {
	a.Set("joinValues", value)
	return a
}

func (a *TreeSelectControl) Label(value interface{}) *TreeSelectControl {
	a.Set("label", value)
	return a
}

func (a *TreeSelectControl) LabelAlign(value interface{}) *TreeSelectControl {
	a.Set("labelAlign", value)
	return a
}

func (a *TreeSelectControl) LabelClassName(value interface{}) *TreeSelectControl {
	a.Set("labelClassName", value)
	return a
}

func (a *TreeSelectControl) LabelField(value interface{}) *TreeSelectControl {
	a.Set("labelField", value)
	return a
}

func (a *TreeSelectControl) LabelRemark(value interface{}) *TreeSelectControl {
	a.Set("labelRemark", value)
	return a
}

func (a *TreeSelectControl) LabelWidth(value interface{}) *TreeSelectControl {
	a.Set("labelWidth", value)
	return a
}

func (a *TreeSelectControl) MaxTagCount(value interface{}) *TreeSelectControl {
	a.Set("maxTagCount", value)
	return a
}

func (a *TreeSelectControl) MenuTpl(value interface{}) *TreeSelectControl {
	a.Set("menuTpl", value)
	return a
}

func (a *TreeSelectControl) Mode(value interface{}) *TreeSelectControl {
	a.Set("mode", value)
	return a
}

func (a *TreeSelectControl) Multiple(value interface{}) *TreeSelectControl {
	a.Set("multiple", value)
	return a
}

func (a *TreeSelectControl) Name(value interface{}) *TreeSelectControl {
	a.Set("name", value)
	return a
}

func (a *TreeSelectControl) NodeBehavior(value interface{}) *TreeSelectControl {
	a.Set("nodeBehavior", value)
	return a
}

func (a *TreeSelectControl) OnEvent(value interface{}) *TreeSelectControl {
	a.Set("onEvent", value)
	return a
}

func (a *TreeSelectControl) OnlyChildren(value interface{}) *TreeSelectControl {
	a.Set("onlyChildren", value)
	return a
}

func (a *TreeSelectControl) OnlyLeaf(value interface{}) *TreeSelectControl {
	a.Set("onlyLeaf", value)
	return a
}

func (a *TreeSelectControl) Options(value interface{}) *TreeSelectControl {
	a.Set("options", value)
	return a
}

func (a *TreeSelectControl) OverflowTagPopover(value interface{}) *TreeSelectControl {
	a.Set("overflowTagPopover", value)
	return a
}

func (a *TreeSelectControl) PathSeparator(value interface{}) *TreeSelectControl {
	a.Set("pathSeparator", value)
	return a
}

func (a *TreeSelectControl) Placeholder(value interface{}) *TreeSelectControl {
	a.Set("placeholder", value)
	return a
}

func (a *TreeSelectControl) ReadOnly(value interface{}) *TreeSelectControl {
	a.Set("readOnly", value)
	return a
}

func (a *TreeSelectControl) ReadOnlyOn(value interface{}) *TreeSelectControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *TreeSelectControl) Remark(value interface{}) *TreeSelectControl {
	a.Set("remark", value)
	return a
}

func (a *TreeSelectControl) Removable(value interface{}) *TreeSelectControl {
	a.Set("removable", value)
	return a
}

func (a *TreeSelectControl) Required(value interface{}) *TreeSelectControl {
	a.Set("required", value)
	return a
}

func (a *TreeSelectControl) ResetValue(value interface{}) *TreeSelectControl {
	a.Set("resetValue", value)
	return a
}

func (a *TreeSelectControl) RootCreatable(value interface{}) *TreeSelectControl {
	a.Set("rootCreatable", value)
	return a
}

func (a *TreeSelectControl) RootLabel(value interface{}) *TreeSelectControl {
	a.Set("rootLabel", value)
	return a
}

func (a *TreeSelectControl) RootValue(value interface{}) *TreeSelectControl {
	a.Set("rootValue", value)
	return a
}

func (a *TreeSelectControl) Row(value interface{}) *TreeSelectControl {
	a.Set("row", value)
	return a
}

func (a *TreeSelectControl) SaveImmediately(value interface{}) *TreeSelectControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *TreeSelectControl) Searchable(value interface{}) *TreeSelectControl {
	a.Set("searchable", value)
	return a
}

func (a *TreeSelectControl) SelectFirst(value interface{}) *TreeSelectControl {
	a.Set("selectFirst", value)
	return a
}

func (a *TreeSelectControl) ShowIcon(value interface{}) *TreeSelectControl {
	a.Set("showIcon", value)
	return a
}

func (a *TreeSelectControl) ShowOutline(value interface{}) *TreeSelectControl {
	a.Set("showOutline", value)
	return a
}

func (a *TreeSelectControl) Size(value interface{}) *TreeSelectControl {
	a.Set("size", value)
	return a
}

func (a *TreeSelectControl) Source(value interface{}) *TreeSelectControl {
	a.Set("source", value)
	return a
}

func (a *TreeSelectControl) Static(value interface{}) *TreeSelectControl {
	a.Set("static", value)
	return a
}

func (a *TreeSelectControl) StaticClassName(value interface{}) *TreeSelectControl {
	a.Set("staticClassName", value)
	return a
}

func (a *TreeSelectControl) StaticInputClassName(value interface{}) *TreeSelectControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *TreeSelectControl) StaticLabelClassName(value interface{}) *TreeSelectControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *TreeSelectControl) StaticOn(value interface{}) *TreeSelectControl {
	a.Set("staticOn", value)
	return a
}

func (a *TreeSelectControl) StaticPlaceholder(value interface{}) *TreeSelectControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *TreeSelectControl) StaticSchema(value interface{}) *TreeSelectControl {
	a.Set("staticSchema", value)
	return a
}

func (a *TreeSelectControl) Style(value interface{}) *TreeSelectControl {
	a.Set("style", value)
	return a
}

func (a *TreeSelectControl) SubmitOnChange(value interface{}) *TreeSelectControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *TreeSelectControl) TestIdBuilder(value interface{}) *TreeSelectControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *TreeSelectControl) Type(value interface{}) *TreeSelectControl {
	a.Set("type", value)
	return a
}

func (a *TreeSelectControl) UseMobileUI(value interface{}) *TreeSelectControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *TreeSelectControl) ValidateApi(value interface{}) *TreeSelectControl {
	a.Set("validateApi", value)
	return a
}

func (a *TreeSelectControl) ValidateOnChange(value interface{}) *TreeSelectControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *TreeSelectControl) ValidationErrors(value interface{}) *TreeSelectControl {
	a.Set("validationErrors", value)
	return a
}

func (a *TreeSelectControl) Validations(value interface{}) *TreeSelectControl {
	a.Set("validations", value)
	return a
}

func (a *TreeSelectControl) Value(value interface{}) *TreeSelectControl {
	a.Set("value", value)
	return a
}

func (a *TreeSelectControl) ValueField(value interface{}) *TreeSelectControl {
	a.Set("valueField", value)
	return a
}

func (a *TreeSelectControl) ValuesNoWrap(value interface{}) *TreeSelectControl {
	a.Set("valuesNoWrap", value)
	return a
}

func (a *TreeSelectControl) Visible(value interface{}) *TreeSelectControl {
	a.Set("visible", value)
	return a
}

func (a *TreeSelectControl) VisibleOn(value interface{}) *TreeSelectControl {
	a.Set("visibleOn", value)
	return a
}

func (a *TreeSelectControl) Width(value interface{}) *TreeSelectControl {
	a.Set("width", value)
	return a
}

func (a *TreeSelectControl) WithChildren(value interface{}) *TreeSelectControl {
	a.Set("withChildren", value)
	return a
}
