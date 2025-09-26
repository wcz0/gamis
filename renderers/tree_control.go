package renderers

type TreeControl struct {
	*BaseRenderer
}

func NewTreeControl() *TreeControl {
	a := &TreeControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "input-tree")
	return a
}

func (a *TreeControl) Set(name string, value interface{}) *TreeControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *TreeControl) AddApi(value interface{}) *TreeControl {
	a.Set("addApi", value)
	return a
}

func (a *TreeControl) AddControls(value interface{}) *TreeControl {
	a.Set("addControls", value)
	return a
}

func (a *TreeControl) AddDialog(value interface{}) *TreeControl {
	a.Set("addDialog", value)
	return a
}

func (a *TreeControl) AutoCancelParent(value interface{}) *TreeControl {
	a.Set("autoCancelParent", value)
	return a
}

func (a *TreeControl) AutoCheckChildren(value interface{}) *TreeControl {
	a.Set("autoCheckChildren", value)
	return a
}

func (a *TreeControl) AutoFill(value interface{}) *TreeControl {
	a.Set("autoFill", value)
	return a
}

func (a *TreeControl) Cascade(value interface{}) *TreeControl {
	a.Set("cascade", value)
	return a
}

func (a *TreeControl) ClassName(value interface{}) *TreeControl {
	a.Set("className", value)
	return a
}

func (a *TreeControl) ClearValueOnHidden(value interface{}) *TreeControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *TreeControl) ClearValueOnSourceChange(value interface{}) *TreeControl {
	a.Set("clearValueOnSourceChange", value)
	return a
}

func (a *TreeControl) Clearable(value interface{}) *TreeControl {
	a.Set("clearable", value)
	return a
}

func (a *TreeControl) Creatable(value interface{}) *TreeControl {
	a.Set("creatable", value)
	return a
}

func (a *TreeControl) CreateBtnLabel(value interface{}) *TreeControl {
	a.Set("createBtnLabel", value)
	return a
}

func (a *TreeControl) DeferApi(value interface{}) *TreeControl {
	a.Set("deferApi", value)
	return a
}

func (a *TreeControl) DeferField(value interface{}) *TreeControl {
	a.Set("deferField", value)
	return a
}

func (a *TreeControl) DeleteApi(value interface{}) *TreeControl {
	a.Set("deleteApi", value)
	return a
}

func (a *TreeControl) DeleteConfirmText(value interface{}) *TreeControl {
	a.Set("deleteConfirmText", value)
	return a
}

func (a *TreeControl) Delimiter(value interface{}) *TreeControl {
	a.Set("delimiter", value)
	return a
}

func (a *TreeControl) Desc(value interface{}) *TreeControl {
	a.Set("desc", value)
	return a
}

func (a *TreeControl) Description(value interface{}) *TreeControl {
	a.Set("description", value)
	return a
}

func (a *TreeControl) DescriptionClassName(value interface{}) *TreeControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *TreeControl) Disabled(value interface{}) *TreeControl {
	a.Set("disabled", value)
	return a
}

func (a *TreeControl) DisabledOn(value interface{}) *TreeControl {
	a.Set("disabledOn", value)
	return a
}

func (a *TreeControl) EditApi(value interface{}) *TreeControl {
	a.Set("editApi", value)
	return a
}

func (a *TreeControl) EditControls(value interface{}) *TreeControl {
	a.Set("editControls", value)
	return a
}

func (a *TreeControl) EditDialog(value interface{}) *TreeControl {
	a.Set("editDialog", value)
	return a
}

func (a *TreeControl) Editable(value interface{}) *TreeControl {
	a.Set("editable", value)
	return a
}

func (a *TreeControl) EditorSetting(value interface{}) *TreeControl {
	a.Set("editorSetting", value)
	return a
}

func (a *TreeControl) EnableDefaultIcon(value interface{}) *TreeControl {
	a.Set("enableDefaultIcon", value)
	return a
}

func (a *TreeControl) EnableNodePath(value interface{}) *TreeControl {
	a.Set("enableNodePath", value)
	return a
}

func (a *TreeControl) ExtraName(value interface{}) *TreeControl {
	a.Set("extraName", value)
	return a
}

func (a *TreeControl) ExtractValue(value interface{}) *TreeControl {
	a.Set("extractValue", value)
	return a
}

func (a *TreeControl) HeightAuto(value interface{}) *TreeControl {
	a.Set("heightAuto", value)
	return a
}

func (a *TreeControl) Hidden(value interface{}) *TreeControl {
	a.Set("hidden", value)
	return a
}

func (a *TreeControl) HiddenOn(value interface{}) *TreeControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *TreeControl) HideRoot(value interface{}) *TreeControl {
	a.Set("hideRoot", value)
	return a
}

func (a *TreeControl) HighlightTxt(value interface{}) *TreeControl {
	a.Set("highlightTxt", value)
	return a
}

func (a *TreeControl) Hint(value interface{}) *TreeControl {
	a.Set("hint", value)
	return a
}

func (a *TreeControl) Horizontal(value interface{}) *TreeControl {
	a.Set("horizontal", value)
	return a
}

func (a *TreeControl) Id(value interface{}) *TreeControl {
	a.Set("id", value)
	return a
}

func (a *TreeControl) InitAutoFill(value interface{}) *TreeControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *TreeControl) InitFetch(value interface{}) *TreeControl {
	a.Set("initFetch", value)
	return a
}

func (a *TreeControl) InitFetchOn(value interface{}) *TreeControl {
	a.Set("initFetchOn", value)
	return a
}

func (a *TreeControl) Inline(value interface{}) *TreeControl {
	a.Set("inline", value)
	return a
}

func (a *TreeControl) InputClassName(value interface{}) *TreeControl {
	a.Set("inputClassName", value)
	return a
}

func (a *TreeControl) ItemActions(value interface{}) *TreeControl {
	a.Set("itemActions", value)
	return a
}

func (a *TreeControl) JoinValues(value interface{}) *TreeControl {
	a.Set("joinValues", value)
	return a
}

func (a *TreeControl) Label(value interface{}) *TreeControl {
	a.Set("label", value)
	return a
}

func (a *TreeControl) LabelAlign(value interface{}) *TreeControl {
	a.Set("labelAlign", value)
	return a
}

func (a *TreeControl) LabelClassName(value interface{}) *TreeControl {
	a.Set("labelClassName", value)
	return a
}

func (a *TreeControl) LabelField(value interface{}) *TreeControl {
	a.Set("labelField", value)
	return a
}

func (a *TreeControl) LabelRemark(value interface{}) *TreeControl {
	a.Set("labelRemark", value)
	return a
}

func (a *TreeControl) LabelWidth(value interface{}) *TreeControl {
	a.Set("labelWidth", value)
	return a
}

func (a *TreeControl) Mode(value interface{}) *TreeControl {
	a.Set("mode", value)
	return a
}

func (a *TreeControl) Multiple(value interface{}) *TreeControl {
	a.Set("multiple", value)
	return a
}

func (a *TreeControl) Name(value interface{}) *TreeControl {
	a.Set("name", value)
	return a
}

func (a *TreeControl) NodeBehavior(value interface{}) *TreeControl {
	a.Set("nodeBehavior", value)
	return a
}

func (a *TreeControl) OnEvent(value interface{}) *TreeControl {
	a.Set("onEvent", value)
	return a
}

func (a *TreeControl) OnlyChildren(value interface{}) *TreeControl {
	a.Set("onlyChildren", value)
	return a
}

func (a *TreeControl) OnlyLeaf(value interface{}) *TreeControl {
	a.Set("onlyLeaf", value)
	return a
}

func (a *TreeControl) Options(value interface{}) *TreeControl {
	a.Set("options", value)
	return a
}

func (a *TreeControl) PathSeparator(value interface{}) *TreeControl {
	a.Set("pathSeparator", value)
	return a
}

func (a *TreeControl) Placeholder(value interface{}) *TreeControl {
	a.Set("placeholder", value)
	return a
}

func (a *TreeControl) ReadOnly(value interface{}) *TreeControl {
	a.Set("readOnly", value)
	return a
}

func (a *TreeControl) ReadOnlyOn(value interface{}) *TreeControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *TreeControl) Remark(value interface{}) *TreeControl {
	a.Set("remark", value)
	return a
}

func (a *TreeControl) Removable(value interface{}) *TreeControl {
	a.Set("removable", value)
	return a
}

func (a *TreeControl) Required(value interface{}) *TreeControl {
	a.Set("required", value)
	return a
}

func (a *TreeControl) ResetValue(value interface{}) *TreeControl {
	a.Set("resetValue", value)
	return a
}

func (a *TreeControl) RootCreatable(value interface{}) *TreeControl {
	a.Set("rootCreatable", value)
	return a
}

func (a *TreeControl) RootLabel(value interface{}) *TreeControl {
	a.Set("rootLabel", value)
	return a
}

func (a *TreeControl) RootValue(value interface{}) *TreeControl {
	a.Set("rootValue", value)
	return a
}

func (a *TreeControl) Row(value interface{}) *TreeControl {
	a.Set("row", value)
	return a
}

func (a *TreeControl) SaveImmediately(value interface{}) *TreeControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *TreeControl) SearchApi(value interface{}) *TreeControl {
	a.Set("searchApi", value)
	return a
}

func (a *TreeControl) SearchConfig(value interface{}) *TreeControl {
	a.Set("searchConfig", value)
	return a
}

func (a *TreeControl) Searchable(value interface{}) *TreeControl {
	a.Set("searchable", value)
	return a
}

func (a *TreeControl) SelectFirst(value interface{}) *TreeControl {
	a.Set("selectFirst", value)
	return a
}

func (a *TreeControl) ShowIcon(value interface{}) *TreeControl {
	a.Set("showIcon", value)
	return a
}

func (a *TreeControl) ShowOutline(value interface{}) *TreeControl {
	a.Set("showOutline", value)
	return a
}

func (a *TreeControl) Size(value interface{}) *TreeControl {
	a.Set("size", value)
	return a
}

func (a *TreeControl) Source(value interface{}) *TreeControl {
	a.Set("source", value)
	return a
}

func (a *TreeControl) Static(value interface{}) *TreeControl {
	a.Set("static", value)
	return a
}

func (a *TreeControl) StaticClassName(value interface{}) *TreeControl {
	a.Set("staticClassName", value)
	return a
}

func (a *TreeControl) StaticInputClassName(value interface{}) *TreeControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *TreeControl) StaticLabelClassName(value interface{}) *TreeControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *TreeControl) StaticOn(value interface{}) *TreeControl {
	a.Set("staticOn", value)
	return a
}

func (a *TreeControl) StaticPlaceholder(value interface{}) *TreeControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *TreeControl) StaticSchema(value interface{}) *TreeControl {
	a.Set("staticSchema", value)
	return a
}

func (a *TreeControl) Style(value interface{}) *TreeControl {
	a.Set("style", value)
	return a
}

func (a *TreeControl) SubmitOnChange(value interface{}) *TreeControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *TreeControl) TestIdBuilder(value interface{}) *TreeControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *TreeControl) Type(value interface{}) *TreeControl {
	a.Set("type", value)
	return a
}

func (a *TreeControl) UseMobileUI(value interface{}) *TreeControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *TreeControl) ValidateApi(value interface{}) *TreeControl {
	a.Set("validateApi", value)
	return a
}

func (a *TreeControl) ValidateOnChange(value interface{}) *TreeControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *TreeControl) ValidationErrors(value interface{}) *TreeControl {
	a.Set("validationErrors", value)
	return a
}

func (a *TreeControl) Validations(value interface{}) *TreeControl {
	a.Set("validations", value)
	return a
}

func (a *TreeControl) Value(value interface{}) *TreeControl {
	a.Set("value", value)
	return a
}

func (a *TreeControl) ValueField(value interface{}) *TreeControl {
	a.Set("valueField", value)
	return a
}

func (a *TreeControl) ValuesNoWrap(value interface{}) *TreeControl {
	a.Set("valuesNoWrap", value)
	return a
}

func (a *TreeControl) Visible(value interface{}) *TreeControl {
	a.Set("visible", value)
	return a
}

func (a *TreeControl) VisibleOn(value interface{}) *TreeControl {
	a.Set("visibleOn", value)
	return a
}

func (a *TreeControl) Width(value interface{}) *TreeControl {
	a.Set("width", value)
	return a
}

func (a *TreeControl) WithChildren(value interface{}) *TreeControl {
	a.Set("withChildren", value)
	return a
}
