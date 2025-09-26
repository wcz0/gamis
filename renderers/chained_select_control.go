package renderers

type ChainedSelectControl struct {
	*BaseRenderer
}

func NewChainedSelectControl() *ChainedSelectControl {
	a := &ChainedSelectControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "chained-select")
	return a
}

func (a *ChainedSelectControl) Set(name string, value interface{}) *ChainedSelectControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *ChainedSelectControl) AddApi(value interface{}) *ChainedSelectControl {
	a.Set("addApi", value)
	return a
}

func (a *ChainedSelectControl) AddControls(value interface{}) *ChainedSelectControl {
	a.Set("addControls", value)
	return a
}

func (a *ChainedSelectControl) AddDialog(value interface{}) *ChainedSelectControl {
	a.Set("addDialog", value)
	return a
}

func (a *ChainedSelectControl) AutoFill(value interface{}) *ChainedSelectControl {
	a.Set("autoFill", value)
	return a
}

func (a *ChainedSelectControl) ClassName(value interface{}) *ChainedSelectControl {
	a.Set("className", value)
	return a
}

func (a *ChainedSelectControl) ClearValueOnHidden(value interface{}) *ChainedSelectControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *ChainedSelectControl) ClearValueOnSourceChange(value interface{}) *ChainedSelectControl {
	a.Set("clearValueOnSourceChange", value)
	return a
}

func (a *ChainedSelectControl) Clearable(value interface{}) *ChainedSelectControl {
	a.Set("clearable", value)
	return a
}

func (a *ChainedSelectControl) Creatable(value interface{}) *ChainedSelectControl {
	a.Set("creatable", value)
	return a
}

func (a *ChainedSelectControl) CreateBtnLabel(value interface{}) *ChainedSelectControl {
	a.Set("createBtnLabel", value)
	return a
}

func (a *ChainedSelectControl) DeferApi(value interface{}) *ChainedSelectControl {
	a.Set("deferApi", value)
	return a
}

func (a *ChainedSelectControl) DeferField(value interface{}) *ChainedSelectControl {
	a.Set("deferField", value)
	return a
}

func (a *ChainedSelectControl) DeleteApi(value interface{}) *ChainedSelectControl {
	a.Set("deleteApi", value)
	return a
}

func (a *ChainedSelectControl) DeleteConfirmText(value interface{}) *ChainedSelectControl {
	a.Set("deleteConfirmText", value)
	return a
}

func (a *ChainedSelectControl) Delimiter(value interface{}) *ChainedSelectControl {
	a.Set("delimiter", value)
	return a
}

func (a *ChainedSelectControl) Desc(value interface{}) *ChainedSelectControl {
	a.Set("desc", value)
	return a
}

func (a *ChainedSelectControl) Description(value interface{}) *ChainedSelectControl {
	a.Set("description", value)
	return a
}

func (a *ChainedSelectControl) DescriptionClassName(value interface{}) *ChainedSelectControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *ChainedSelectControl) Disabled(value interface{}) *ChainedSelectControl {
	a.Set("disabled", value)
	return a
}

func (a *ChainedSelectControl) DisabledOn(value interface{}) *ChainedSelectControl {
	a.Set("disabledOn", value)
	return a
}

func (a *ChainedSelectControl) EditApi(value interface{}) *ChainedSelectControl {
	a.Set("editApi", value)
	return a
}

func (a *ChainedSelectControl) EditControls(value interface{}) *ChainedSelectControl {
	a.Set("editControls", value)
	return a
}

func (a *ChainedSelectControl) EditDialog(value interface{}) *ChainedSelectControl {
	a.Set("editDialog", value)
	return a
}

func (a *ChainedSelectControl) Editable(value interface{}) *ChainedSelectControl {
	a.Set("editable", value)
	return a
}

func (a *ChainedSelectControl) EditorSetting(value interface{}) *ChainedSelectControl {
	a.Set("editorSetting", value)
	return a
}

func (a *ChainedSelectControl) ExtraName(value interface{}) *ChainedSelectControl {
	a.Set("extraName", value)
	return a
}

func (a *ChainedSelectControl) ExtractValue(value interface{}) *ChainedSelectControl {
	a.Set("extractValue", value)
	return a
}

func (a *ChainedSelectControl) Hidden(value interface{}) *ChainedSelectControl {
	a.Set("hidden", value)
	return a
}

func (a *ChainedSelectControl) HiddenOn(value interface{}) *ChainedSelectControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *ChainedSelectControl) Hint(value interface{}) *ChainedSelectControl {
	a.Set("hint", value)
	return a
}

func (a *ChainedSelectControl) Horizontal(value interface{}) *ChainedSelectControl {
	a.Set("horizontal", value)
	return a
}

func (a *ChainedSelectControl) Id(value interface{}) *ChainedSelectControl {
	a.Set("id", value)
	return a
}

func (a *ChainedSelectControl) InitAutoFill(value interface{}) *ChainedSelectControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *ChainedSelectControl) InitFetch(value interface{}) *ChainedSelectControl {
	a.Set("initFetch", value)
	return a
}

func (a *ChainedSelectControl) InitFetchOn(value interface{}) *ChainedSelectControl {
	a.Set("initFetchOn", value)
	return a
}

func (a *ChainedSelectControl) Inline(value interface{}) *ChainedSelectControl {
	a.Set("inline", value)
	return a
}

func (a *ChainedSelectControl) InputClassName(value interface{}) *ChainedSelectControl {
	a.Set("inputClassName", value)
	return a
}

func (a *ChainedSelectControl) JoinValues(value interface{}) *ChainedSelectControl {
	a.Set("joinValues", value)
	return a
}

func (a *ChainedSelectControl) Label(value interface{}) *ChainedSelectControl {
	a.Set("label", value)
	return a
}

func (a *ChainedSelectControl) LabelAlign(value interface{}) *ChainedSelectControl {
	a.Set("labelAlign", value)
	return a
}

func (a *ChainedSelectControl) LabelClassName(value interface{}) *ChainedSelectControl {
	a.Set("labelClassName", value)
	return a
}

func (a *ChainedSelectControl) LabelRemark(value interface{}) *ChainedSelectControl {
	a.Set("labelRemark", value)
	return a
}

func (a *ChainedSelectControl) LabelWidth(value interface{}) *ChainedSelectControl {
	a.Set("labelWidth", value)
	return a
}

func (a *ChainedSelectControl) Mode(value interface{}) *ChainedSelectControl {
	a.Set("mode", value)
	return a
}

func (a *ChainedSelectControl) Multiple(value interface{}) *ChainedSelectControl {
	a.Set("multiple", value)
	return a
}

func (a *ChainedSelectControl) Name(value interface{}) *ChainedSelectControl {
	a.Set("name", value)
	return a
}

func (a *ChainedSelectControl) OnEvent(value interface{}) *ChainedSelectControl {
	a.Set("onEvent", value)
	return a
}

func (a *ChainedSelectControl) Options(value interface{}) *ChainedSelectControl {
	a.Set("options", value)
	return a
}

func (a *ChainedSelectControl) Placeholder(value interface{}) *ChainedSelectControl {
	a.Set("placeholder", value)
	return a
}

func (a *ChainedSelectControl) ReadOnly(value interface{}) *ChainedSelectControl {
	a.Set("readOnly", value)
	return a
}

func (a *ChainedSelectControl) ReadOnlyOn(value interface{}) *ChainedSelectControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *ChainedSelectControl) Remark(value interface{}) *ChainedSelectControl {
	a.Set("remark", value)
	return a
}

func (a *ChainedSelectControl) Removable(value interface{}) *ChainedSelectControl {
	a.Set("removable", value)
	return a
}

func (a *ChainedSelectControl) Required(value interface{}) *ChainedSelectControl {
	a.Set("required", value)
	return a
}

func (a *ChainedSelectControl) ResetValue(value interface{}) *ChainedSelectControl {
	a.Set("resetValue", value)
	return a
}

func (a *ChainedSelectControl) Row(value interface{}) *ChainedSelectControl {
	a.Set("row", value)
	return a
}

func (a *ChainedSelectControl) SaveImmediately(value interface{}) *ChainedSelectControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *ChainedSelectControl) SelectFirst(value interface{}) *ChainedSelectControl {
	a.Set("selectFirst", value)
	return a
}

func (a *ChainedSelectControl) Size(value interface{}) *ChainedSelectControl {
	a.Set("size", value)
	return a
}

func (a *ChainedSelectControl) Source(value interface{}) *ChainedSelectControl {
	a.Set("source", value)
	return a
}

func (a *ChainedSelectControl) Static(value interface{}) *ChainedSelectControl {
	a.Set("static", value)
	return a
}

func (a *ChainedSelectControl) StaticClassName(value interface{}) *ChainedSelectControl {
	a.Set("staticClassName", value)
	return a
}

func (a *ChainedSelectControl) StaticInputClassName(value interface{}) *ChainedSelectControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *ChainedSelectControl) StaticLabelClassName(value interface{}) *ChainedSelectControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *ChainedSelectControl) StaticOn(value interface{}) *ChainedSelectControl {
	a.Set("staticOn", value)
	return a
}

func (a *ChainedSelectControl) StaticPlaceholder(value interface{}) *ChainedSelectControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *ChainedSelectControl) StaticSchema(value interface{}) *ChainedSelectControl {
	a.Set("staticSchema", value)
	return a
}

func (a *ChainedSelectControl) Style(value interface{}) *ChainedSelectControl {
	a.Set("style", value)
	return a
}

func (a *ChainedSelectControl) SubmitOnChange(value interface{}) *ChainedSelectControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *ChainedSelectControl) TestIdBuilder(value interface{}) *ChainedSelectControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *ChainedSelectControl) Type(value interface{}) *ChainedSelectControl {
	a.Set("type", value)
	return a
}

func (a *ChainedSelectControl) UseMobileUI(value interface{}) *ChainedSelectControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *ChainedSelectControl) ValidateApi(value interface{}) *ChainedSelectControl {
	a.Set("validateApi", value)
	return a
}

func (a *ChainedSelectControl) ValidateOnChange(value interface{}) *ChainedSelectControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *ChainedSelectControl) ValidationErrors(value interface{}) *ChainedSelectControl {
	a.Set("validationErrors", value)
	return a
}

func (a *ChainedSelectControl) Validations(value interface{}) *ChainedSelectControl {
	a.Set("validations", value)
	return a
}

func (a *ChainedSelectControl) Value(value interface{}) *ChainedSelectControl {
	a.Set("value", value)
	return a
}

func (a *ChainedSelectControl) ValuesNoWrap(value interface{}) *ChainedSelectControl {
	a.Set("valuesNoWrap", value)
	return a
}

func (a *ChainedSelectControl) Visible(value interface{}) *ChainedSelectControl {
	a.Set("visible", value)
	return a
}

func (a *ChainedSelectControl) VisibleOn(value interface{}) *ChainedSelectControl {
	a.Set("visibleOn", value)
	return a
}

func (a *ChainedSelectControl) Width(value interface{}) *ChainedSelectControl {
	a.Set("width", value)
	return a
}
