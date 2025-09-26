package renderers

type NestedSelectControl struct {
	*BaseRenderer
}

func NewNestedSelectControl() *NestedSelectControl {
	a := &NestedSelectControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "nested-select")
	return a
}

func (a *NestedSelectControl) Set(name string, value interface{}) *NestedSelectControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *NestedSelectControl) AddApi(value interface{}) *NestedSelectControl {
	a.Set("addApi", value)
	return a
}

func (a *NestedSelectControl) AddControls(value interface{}) *NestedSelectControl {
	a.Set("addControls", value)
	return a
}

func (a *NestedSelectControl) AddDialog(value interface{}) *NestedSelectControl {
	a.Set("addDialog", value)
	return a
}

func (a *NestedSelectControl) AutoFill(value interface{}) *NestedSelectControl {
	a.Set("autoFill", value)
	return a
}

func (a *NestedSelectControl) BorderMode(value interface{}) *NestedSelectControl {
	a.Set("borderMode", value)
	return a
}

func (a *NestedSelectControl) Cascade(value interface{}) *NestedSelectControl {
	a.Set("cascade", value)
	return a
}

func (a *NestedSelectControl) ClassName(value interface{}) *NestedSelectControl {
	a.Set("className", value)
	return a
}

func (a *NestedSelectControl) ClearValueOnHidden(value interface{}) *NestedSelectControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *NestedSelectControl) ClearValueOnSourceChange(value interface{}) *NestedSelectControl {
	a.Set("clearValueOnSourceChange", value)
	return a
}

func (a *NestedSelectControl) Clearable(value interface{}) *NestedSelectControl {
	a.Set("clearable", value)
	return a
}

func (a *NestedSelectControl) Creatable(value interface{}) *NestedSelectControl {
	a.Set("creatable", value)
	return a
}

func (a *NestedSelectControl) CreateBtnLabel(value interface{}) *NestedSelectControl {
	a.Set("createBtnLabel", value)
	return a
}

func (a *NestedSelectControl) DeferApi(value interface{}) *NestedSelectControl {
	a.Set("deferApi", value)
	return a
}

func (a *NestedSelectControl) DeferField(value interface{}) *NestedSelectControl {
	a.Set("deferField", value)
	return a
}

func (a *NestedSelectControl) DeleteApi(value interface{}) *NestedSelectControl {
	a.Set("deleteApi", value)
	return a
}

func (a *NestedSelectControl) DeleteConfirmText(value interface{}) *NestedSelectControl {
	a.Set("deleteConfirmText", value)
	return a
}

func (a *NestedSelectControl) Delimiter(value interface{}) *NestedSelectControl {
	a.Set("delimiter", value)
	return a
}

func (a *NestedSelectControl) Desc(value interface{}) *NestedSelectControl {
	a.Set("desc", value)
	return a
}

func (a *NestedSelectControl) Description(value interface{}) *NestedSelectControl {
	a.Set("description", value)
	return a
}

func (a *NestedSelectControl) DescriptionClassName(value interface{}) *NestedSelectControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *NestedSelectControl) Disabled(value interface{}) *NestedSelectControl {
	a.Set("disabled", value)
	return a
}

func (a *NestedSelectControl) DisabledOn(value interface{}) *NestedSelectControl {
	a.Set("disabledOn", value)
	return a
}

func (a *NestedSelectControl) EditApi(value interface{}) *NestedSelectControl {
	a.Set("editApi", value)
	return a
}

func (a *NestedSelectControl) EditControls(value interface{}) *NestedSelectControl {
	a.Set("editControls", value)
	return a
}

func (a *NestedSelectControl) EditDialog(value interface{}) *NestedSelectControl {
	a.Set("editDialog", value)
	return a
}

func (a *NestedSelectControl) Editable(value interface{}) *NestedSelectControl {
	a.Set("editable", value)
	return a
}

func (a *NestedSelectControl) EditorSetting(value interface{}) *NestedSelectControl {
	a.Set("editorSetting", value)
	return a
}

func (a *NestedSelectControl) ExtraName(value interface{}) *NestedSelectControl {
	a.Set("extraName", value)
	return a
}

func (a *NestedSelectControl) ExtractValue(value interface{}) *NestedSelectControl {
	a.Set("extractValue", value)
	return a
}

func (a *NestedSelectControl) Hidden(value interface{}) *NestedSelectControl {
	a.Set("hidden", value)
	return a
}

func (a *NestedSelectControl) HiddenOn(value interface{}) *NestedSelectControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *NestedSelectControl) HideNodePathLabel(value interface{}) *NestedSelectControl {
	a.Set("hideNodePathLabel", value)
	return a
}

func (a *NestedSelectControl) Hint(value interface{}) *NestedSelectControl {
	a.Set("hint", value)
	return a
}

func (a *NestedSelectControl) Horizontal(value interface{}) *NestedSelectControl {
	a.Set("horizontal", value)
	return a
}

func (a *NestedSelectControl) Id(value interface{}) *NestedSelectControl {
	a.Set("id", value)
	return a
}

func (a *NestedSelectControl) InitAutoFill(value interface{}) *NestedSelectControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *NestedSelectControl) InitFetch(value interface{}) *NestedSelectControl {
	a.Set("initFetch", value)
	return a
}

func (a *NestedSelectControl) InitFetchOn(value interface{}) *NestedSelectControl {
	a.Set("initFetchOn", value)
	return a
}

func (a *NestedSelectControl) Inline(value interface{}) *NestedSelectControl {
	a.Set("inline", value)
	return a
}

func (a *NestedSelectControl) InputClassName(value interface{}) *NestedSelectControl {
	a.Set("inputClassName", value)
	return a
}

func (a *NestedSelectControl) JoinValues(value interface{}) *NestedSelectControl {
	a.Set("joinValues", value)
	return a
}

func (a *NestedSelectControl) Label(value interface{}) *NestedSelectControl {
	a.Set("label", value)
	return a
}

func (a *NestedSelectControl) LabelAlign(value interface{}) *NestedSelectControl {
	a.Set("labelAlign", value)
	return a
}

func (a *NestedSelectControl) LabelClassName(value interface{}) *NestedSelectControl {
	a.Set("labelClassName", value)
	return a
}

func (a *NestedSelectControl) LabelRemark(value interface{}) *NestedSelectControl {
	a.Set("labelRemark", value)
	return a
}

func (a *NestedSelectControl) LabelWidth(value interface{}) *NestedSelectControl {
	a.Set("labelWidth", value)
	return a
}

func (a *NestedSelectControl) MaxTagCount(value interface{}) *NestedSelectControl {
	a.Set("maxTagCount", value)
	return a
}

func (a *NestedSelectControl) MenuClassName(value interface{}) *NestedSelectControl {
	a.Set("menuClassName", value)
	return a
}

func (a *NestedSelectControl) Mode(value interface{}) *NestedSelectControl {
	a.Set("mode", value)
	return a
}

func (a *NestedSelectControl) Multiple(value interface{}) *NestedSelectControl {
	a.Set("multiple", value)
	return a
}

func (a *NestedSelectControl) Name(value interface{}) *NestedSelectControl {
	a.Set("name", value)
	return a
}

func (a *NestedSelectControl) OnEvent(value interface{}) *NestedSelectControl {
	a.Set("onEvent", value)
	return a
}

func (a *NestedSelectControl) OnlyChildren(value interface{}) *NestedSelectControl {
	a.Set("onlyChildren", value)
	return a
}

func (a *NestedSelectControl) OnlyLeaf(value interface{}) *NestedSelectControl {
	a.Set("onlyLeaf", value)
	return a
}

func (a *NestedSelectControl) Options(value interface{}) *NestedSelectControl {
	a.Set("options", value)
	return a
}

func (a *NestedSelectControl) OverflowTagPopover(value interface{}) *NestedSelectControl {
	a.Set("overflowTagPopover", value)
	return a
}

func (a *NestedSelectControl) Placeholder(value interface{}) *NestedSelectControl {
	a.Set("placeholder", value)
	return a
}

func (a *NestedSelectControl) ReadOnly(value interface{}) *NestedSelectControl {
	a.Set("readOnly", value)
	return a
}

func (a *NestedSelectControl) ReadOnlyOn(value interface{}) *NestedSelectControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *NestedSelectControl) Remark(value interface{}) *NestedSelectControl {
	a.Set("remark", value)
	return a
}

func (a *NestedSelectControl) Removable(value interface{}) *NestedSelectControl {
	a.Set("removable", value)
	return a
}

func (a *NestedSelectControl) Required(value interface{}) *NestedSelectControl {
	a.Set("required", value)
	return a
}

func (a *NestedSelectControl) ResetValue(value interface{}) *NestedSelectControl {
	a.Set("resetValue", value)
	return a
}

func (a *NestedSelectControl) Row(value interface{}) *NestedSelectControl {
	a.Set("row", value)
	return a
}

func (a *NestedSelectControl) SaveImmediately(value interface{}) *NestedSelectControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *NestedSelectControl) SelectFirst(value interface{}) *NestedSelectControl {
	a.Set("selectFirst", value)
	return a
}

func (a *NestedSelectControl) Size(value interface{}) *NestedSelectControl {
	a.Set("size", value)
	return a
}

func (a *NestedSelectControl) Source(value interface{}) *NestedSelectControl {
	a.Set("source", value)
	return a
}

func (a *NestedSelectControl) Static(value interface{}) *NestedSelectControl {
	a.Set("static", value)
	return a
}

func (a *NestedSelectControl) StaticClassName(value interface{}) *NestedSelectControl {
	a.Set("staticClassName", value)
	return a
}

func (a *NestedSelectControl) StaticInputClassName(value interface{}) *NestedSelectControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *NestedSelectControl) StaticLabelClassName(value interface{}) *NestedSelectControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *NestedSelectControl) StaticOn(value interface{}) *NestedSelectControl {
	a.Set("staticOn", value)
	return a
}

func (a *NestedSelectControl) StaticPlaceholder(value interface{}) *NestedSelectControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *NestedSelectControl) StaticSchema(value interface{}) *NestedSelectControl {
	a.Set("staticSchema", value)
	return a
}

func (a *NestedSelectControl) Style(value interface{}) *NestedSelectControl {
	a.Set("style", value)
	return a
}

func (a *NestedSelectControl) SubmitOnChange(value interface{}) *NestedSelectControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *NestedSelectControl) TestIdBuilder(value interface{}) *NestedSelectControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *NestedSelectControl) Type(value interface{}) *NestedSelectControl {
	a.Set("type", value)
	return a
}

func (a *NestedSelectControl) UseMobileUI(value interface{}) *NestedSelectControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *NestedSelectControl) ValidateApi(value interface{}) *NestedSelectControl {
	a.Set("validateApi", value)
	return a
}

func (a *NestedSelectControl) ValidateOnChange(value interface{}) *NestedSelectControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *NestedSelectControl) ValidationErrors(value interface{}) *NestedSelectControl {
	a.Set("validationErrors", value)
	return a
}

func (a *NestedSelectControl) Validations(value interface{}) *NestedSelectControl {
	a.Set("validations", value)
	return a
}

func (a *NestedSelectControl) Value(value interface{}) *NestedSelectControl {
	a.Set("value", value)
	return a
}

func (a *NestedSelectControl) ValuesNoWrap(value interface{}) *NestedSelectControl {
	a.Set("valuesNoWrap", value)
	return a
}

func (a *NestedSelectControl) Visible(value interface{}) *NestedSelectControl {
	a.Set("visible", value)
	return a
}

func (a *NestedSelectControl) VisibleOn(value interface{}) *NestedSelectControl {
	a.Set("visibleOn", value)
	return a
}

func (a *NestedSelectControl) Width(value interface{}) *NestedSelectControl {
	a.Set("width", value)
	return a
}

func (a *NestedSelectControl) WithChildren(value interface{}) *NestedSelectControl {
	a.Set("withChildren", value)
	return a
}
