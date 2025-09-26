package renderers

type ListControl struct {
	*BaseRenderer
}

func NewListControl() *ListControl {
	a := &ListControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "list-select")
	return a
}

func (a *ListControl) Set(name string, value interface{}) *ListControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *ListControl) ActiveItemSchema(value interface{}) *ListControl {
	a.Set("activeItemSchema", value)
	return a
}

func (a *ListControl) AddApi(value interface{}) *ListControl {
	a.Set("addApi", value)
	return a
}

func (a *ListControl) AddControls(value interface{}) *ListControl {
	a.Set("addControls", value)
	return a
}

func (a *ListControl) AddDialog(value interface{}) *ListControl {
	a.Set("addDialog", value)
	return a
}

func (a *ListControl) AutoFill(value interface{}) *ListControl {
	a.Set("autoFill", value)
	return a
}

func (a *ListControl) ClassName(value interface{}) *ListControl {
	a.Set("className", value)
	return a
}

func (a *ListControl) ClearValueOnHidden(value interface{}) *ListControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *ListControl) ClearValueOnSourceChange(value interface{}) *ListControl {
	a.Set("clearValueOnSourceChange", value)
	return a
}

func (a *ListControl) Clearable(value interface{}) *ListControl {
	a.Set("clearable", value)
	return a
}

func (a *ListControl) Creatable(value interface{}) *ListControl {
	a.Set("creatable", value)
	return a
}

func (a *ListControl) CreateBtnLabel(value interface{}) *ListControl {
	a.Set("createBtnLabel", value)
	return a
}

func (a *ListControl) DeferApi(value interface{}) *ListControl {
	a.Set("deferApi", value)
	return a
}

func (a *ListControl) DeferField(value interface{}) *ListControl {
	a.Set("deferField", value)
	return a
}

func (a *ListControl) DeleteApi(value interface{}) *ListControl {
	a.Set("deleteApi", value)
	return a
}

func (a *ListControl) DeleteConfirmText(value interface{}) *ListControl {
	a.Set("deleteConfirmText", value)
	return a
}

func (a *ListControl) Delimiter(value interface{}) *ListControl {
	a.Set("delimiter", value)
	return a
}

func (a *ListControl) Desc(value interface{}) *ListControl {
	a.Set("desc", value)
	return a
}

func (a *ListControl) Description(value interface{}) *ListControl {
	a.Set("description", value)
	return a
}

func (a *ListControl) DescriptionClassName(value interface{}) *ListControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *ListControl) Disabled(value interface{}) *ListControl {
	a.Set("disabled", value)
	return a
}

func (a *ListControl) DisabledOn(value interface{}) *ListControl {
	a.Set("disabledOn", value)
	return a
}

func (a *ListControl) EditApi(value interface{}) *ListControl {
	a.Set("editApi", value)
	return a
}

func (a *ListControl) EditControls(value interface{}) *ListControl {
	a.Set("editControls", value)
	return a
}

func (a *ListControl) EditDialog(value interface{}) *ListControl {
	a.Set("editDialog", value)
	return a
}

func (a *ListControl) Editable(value interface{}) *ListControl {
	a.Set("editable", value)
	return a
}

func (a *ListControl) EditorSetting(value interface{}) *ListControl {
	a.Set("editorSetting", value)
	return a
}

func (a *ListControl) ExtraName(value interface{}) *ListControl {
	a.Set("extraName", value)
	return a
}

func (a *ListControl) ExtractValue(value interface{}) *ListControl {
	a.Set("extractValue", value)
	return a
}

func (a *ListControl) Hidden(value interface{}) *ListControl {
	a.Set("hidden", value)
	return a
}

func (a *ListControl) HiddenOn(value interface{}) *ListControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *ListControl) Hint(value interface{}) *ListControl {
	a.Set("hint", value)
	return a
}

func (a *ListControl) Horizontal(value interface{}) *ListControl {
	a.Set("horizontal", value)
	return a
}

func (a *ListControl) Id(value interface{}) *ListControl {
	a.Set("id", value)
	return a
}

func (a *ListControl) ImageClassName(value interface{}) *ListControl {
	a.Set("imageClassName", value)
	return a
}

func (a *ListControl) InitAutoFill(value interface{}) *ListControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *ListControl) InitFetch(value interface{}) *ListControl {
	a.Set("initFetch", value)
	return a
}

func (a *ListControl) InitFetchOn(value interface{}) *ListControl {
	a.Set("initFetchOn", value)
	return a
}

func (a *ListControl) Inline(value interface{}) *ListControl {
	a.Set("inline", value)
	return a
}

func (a *ListControl) InputClassName(value interface{}) *ListControl {
	a.Set("inputClassName", value)
	return a
}

func (a *ListControl) ItemSchema(value interface{}) *ListControl {
	a.Set("itemSchema", value)
	return a
}

func (a *ListControl) JoinValues(value interface{}) *ListControl {
	a.Set("joinValues", value)
	return a
}

func (a *ListControl) Label(value interface{}) *ListControl {
	a.Set("label", value)
	return a
}

func (a *ListControl) LabelAlign(value interface{}) *ListControl {
	a.Set("labelAlign", value)
	return a
}

func (a *ListControl) LabelClassName(value interface{}) *ListControl {
	a.Set("labelClassName", value)
	return a
}

func (a *ListControl) LabelRemark(value interface{}) *ListControl {
	a.Set("labelRemark", value)
	return a
}

func (a *ListControl) LabelWidth(value interface{}) *ListControl {
	a.Set("labelWidth", value)
	return a
}

func (a *ListControl) ListClassName(value interface{}) *ListControl {
	a.Set("listClassName", value)
	return a
}

func (a *ListControl) Mode(value interface{}) *ListControl {
	a.Set("mode", value)
	return a
}

func (a *ListControl) Multiple(value interface{}) *ListControl {
	a.Set("multiple", value)
	return a
}

func (a *ListControl) Name(value interface{}) *ListControl {
	a.Set("name", value)
	return a
}

func (a *ListControl) OnEvent(value interface{}) *ListControl {
	a.Set("onEvent", value)
	return a
}

func (a *ListControl) Options(value interface{}) *ListControl {
	a.Set("options", value)
	return a
}

func (a *ListControl) Placeholder(value interface{}) *ListControl {
	a.Set("placeholder", value)
	return a
}

func (a *ListControl) ReadOnly(value interface{}) *ListControl {
	a.Set("readOnly", value)
	return a
}

func (a *ListControl) ReadOnlyOn(value interface{}) *ListControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *ListControl) Remark(value interface{}) *ListControl {
	a.Set("remark", value)
	return a
}

func (a *ListControl) Removable(value interface{}) *ListControl {
	a.Set("removable", value)
	return a
}

func (a *ListControl) Required(value interface{}) *ListControl {
	a.Set("required", value)
	return a
}

func (a *ListControl) ResetValue(value interface{}) *ListControl {
	a.Set("resetValue", value)
	return a
}

func (a *ListControl) Row(value interface{}) *ListControl {
	a.Set("row", value)
	return a
}

func (a *ListControl) SaveImmediately(value interface{}) *ListControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *ListControl) SelectFirst(value interface{}) *ListControl {
	a.Set("selectFirst", value)
	return a
}

func (a *ListControl) Size(value interface{}) *ListControl {
	a.Set("size", value)
	return a
}

func (a *ListControl) Source(value interface{}) *ListControl {
	a.Set("source", value)
	return a
}

func (a *ListControl) Static(value interface{}) *ListControl {
	a.Set("static", value)
	return a
}

func (a *ListControl) StaticClassName(value interface{}) *ListControl {
	a.Set("staticClassName", value)
	return a
}

func (a *ListControl) StaticInputClassName(value interface{}) *ListControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *ListControl) StaticLabelClassName(value interface{}) *ListControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *ListControl) StaticOn(value interface{}) *ListControl {
	a.Set("staticOn", value)
	return a
}

func (a *ListControl) StaticPlaceholder(value interface{}) *ListControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *ListControl) StaticSchema(value interface{}) *ListControl {
	a.Set("staticSchema", value)
	return a
}

func (a *ListControl) Style(value interface{}) *ListControl {
	a.Set("style", value)
	return a
}

func (a *ListControl) SubmitOnChange(value interface{}) *ListControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *ListControl) SubmitOnDBClick(value interface{}) *ListControl {
	a.Set("submitOnDBClick", value)
	return a
}

func (a *ListControl) TestIdBuilder(value interface{}) *ListControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *ListControl) Type(value interface{}) *ListControl {
	a.Set("type", value)
	return a
}

func (a *ListControl) UseMobileUI(value interface{}) *ListControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *ListControl) ValidateApi(value interface{}) *ListControl {
	a.Set("validateApi", value)
	return a
}

func (a *ListControl) ValidateOnChange(value interface{}) *ListControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *ListControl) ValidationErrors(value interface{}) *ListControl {
	a.Set("validationErrors", value)
	return a
}

func (a *ListControl) Validations(value interface{}) *ListControl {
	a.Set("validations", value)
	return a
}

func (a *ListControl) Value(value interface{}) *ListControl {
	a.Set("value", value)
	return a
}

func (a *ListControl) ValuesNoWrap(value interface{}) *ListControl {
	a.Set("valuesNoWrap", value)
	return a
}

func (a *ListControl) Visible(value interface{}) *ListControl {
	a.Set("visible", value)
	return a
}

func (a *ListControl) VisibleOn(value interface{}) *ListControl {
	a.Set("visibleOn", value)
	return a
}

func (a *ListControl) Width(value interface{}) *ListControl {
	a.Set("width", value)
	return a
}
