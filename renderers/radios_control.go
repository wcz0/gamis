package renderers

type RadiosControl struct {
	*BaseRenderer
}

func NewRadiosControl() *RadiosControl {
	a := &RadiosControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "radios")
	return a
}

func (a *RadiosControl) Set(name string, value interface{}) *RadiosControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *RadiosControl) AddApi(value interface{}) *RadiosControl {
	a.Set("addApi", value)
	return a
}

func (a *RadiosControl) AddControls(value interface{}) *RadiosControl {
	a.Set("addControls", value)
	return a
}

func (a *RadiosControl) AddDialog(value interface{}) *RadiosControl {
	a.Set("addDialog", value)
	return a
}

func (a *RadiosControl) AutoFill(value interface{}) *RadiosControl {
	a.Set("autoFill", value)
	return a
}

func (a *RadiosControl) ClassName(value interface{}) *RadiosControl {
	a.Set("className", value)
	return a
}

func (a *RadiosControl) ClearValueOnHidden(value interface{}) *RadiosControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *RadiosControl) ClearValueOnSourceChange(value interface{}) *RadiosControl {
	a.Set("clearValueOnSourceChange", value)
	return a
}

func (a *RadiosControl) Clearable(value interface{}) *RadiosControl {
	a.Set("clearable", value)
	return a
}

func (a *RadiosControl) ColumnsCount(value interface{}) *RadiosControl {
	a.Set("columnsCount", value)
	return a
}

func (a *RadiosControl) Creatable(value interface{}) *RadiosControl {
	a.Set("creatable", value)
	return a
}

func (a *RadiosControl) CreateBtnLabel(value interface{}) *RadiosControl {
	a.Set("createBtnLabel", value)
	return a
}

func (a *RadiosControl) DeferApi(value interface{}) *RadiosControl {
	a.Set("deferApi", value)
	return a
}

func (a *RadiosControl) DeferField(value interface{}) *RadiosControl {
	a.Set("deferField", value)
	return a
}

func (a *RadiosControl) DeleteApi(value interface{}) *RadiosControl {
	a.Set("deleteApi", value)
	return a
}

func (a *RadiosControl) DeleteConfirmText(value interface{}) *RadiosControl {
	a.Set("deleteConfirmText", value)
	return a
}

func (a *RadiosControl) Delimiter(value interface{}) *RadiosControl {
	a.Set("delimiter", value)
	return a
}

func (a *RadiosControl) Desc(value interface{}) *RadiosControl {
	a.Set("desc", value)
	return a
}

func (a *RadiosControl) Description(value interface{}) *RadiosControl {
	a.Set("description", value)
	return a
}

func (a *RadiosControl) DescriptionClassName(value interface{}) *RadiosControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *RadiosControl) Disabled(value interface{}) *RadiosControl {
	a.Set("disabled", value)
	return a
}

func (a *RadiosControl) DisabledOn(value interface{}) *RadiosControl {
	a.Set("disabledOn", value)
	return a
}

func (a *RadiosControl) EditApi(value interface{}) *RadiosControl {
	a.Set("editApi", value)
	return a
}

func (a *RadiosControl) EditControls(value interface{}) *RadiosControl {
	a.Set("editControls", value)
	return a
}

func (a *RadiosControl) EditDialog(value interface{}) *RadiosControl {
	a.Set("editDialog", value)
	return a
}

func (a *RadiosControl) Editable(value interface{}) *RadiosControl {
	a.Set("editable", value)
	return a
}

func (a *RadiosControl) EditorSetting(value interface{}) *RadiosControl {
	a.Set("editorSetting", value)
	return a
}

func (a *RadiosControl) ExtraName(value interface{}) *RadiosControl {
	a.Set("extraName", value)
	return a
}

func (a *RadiosControl) ExtractValue(value interface{}) *RadiosControl {
	a.Set("extractValue", value)
	return a
}

func (a *RadiosControl) Hidden(value interface{}) *RadiosControl {
	a.Set("hidden", value)
	return a
}

func (a *RadiosControl) HiddenOn(value interface{}) *RadiosControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *RadiosControl) Hint(value interface{}) *RadiosControl {
	a.Set("hint", value)
	return a
}

func (a *RadiosControl) Horizontal(value interface{}) *RadiosControl {
	a.Set("horizontal", value)
	return a
}

func (a *RadiosControl) Id(value interface{}) *RadiosControl {
	a.Set("id", value)
	return a
}

func (a *RadiosControl) InitAutoFill(value interface{}) *RadiosControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *RadiosControl) InitFetch(value interface{}) *RadiosControl {
	a.Set("initFetch", value)
	return a
}

func (a *RadiosControl) InitFetchOn(value interface{}) *RadiosControl {
	a.Set("initFetchOn", value)
	return a
}

func (a *RadiosControl) Inline(value interface{}) *RadiosControl {
	a.Set("inline", value)
	return a
}

func (a *RadiosControl) InputClassName(value interface{}) *RadiosControl {
	a.Set("inputClassName", value)
	return a
}

func (a *RadiosControl) JoinValues(value interface{}) *RadiosControl {
	a.Set("joinValues", value)
	return a
}

func (a *RadiosControl) Label(value interface{}) *RadiosControl {
	a.Set("label", value)
	return a
}

func (a *RadiosControl) LabelAlign(value interface{}) *RadiosControl {
	a.Set("labelAlign", value)
	return a
}

func (a *RadiosControl) LabelClassName(value interface{}) *RadiosControl {
	a.Set("labelClassName", value)
	return a
}

func (a *RadiosControl) LabelRemark(value interface{}) *RadiosControl {
	a.Set("labelRemark", value)
	return a
}

func (a *RadiosControl) LabelWidth(value interface{}) *RadiosControl {
	a.Set("labelWidth", value)
	return a
}

func (a *RadiosControl) Mode(value interface{}) *RadiosControl {
	a.Set("mode", value)
	return a
}

func (a *RadiosControl) Multiple(value interface{}) *RadiosControl {
	a.Set("multiple", value)
	return a
}

func (a *RadiosControl) Name(value interface{}) *RadiosControl {
	a.Set("name", value)
	return a
}

func (a *RadiosControl) OnEvent(value interface{}) *RadiosControl {
	a.Set("onEvent", value)
	return a
}

func (a *RadiosControl) Options(value interface{}) *RadiosControl {
	a.Set("options", value)
	return a
}

func (a *RadiosControl) Placeholder(value interface{}) *RadiosControl {
	a.Set("placeholder", value)
	return a
}

func (a *RadiosControl) ReadOnly(value interface{}) *RadiosControl {
	a.Set("readOnly", value)
	return a
}

func (a *RadiosControl) ReadOnlyOn(value interface{}) *RadiosControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *RadiosControl) Remark(value interface{}) *RadiosControl {
	a.Set("remark", value)
	return a
}

func (a *RadiosControl) Removable(value interface{}) *RadiosControl {
	a.Set("removable", value)
	return a
}

func (a *RadiosControl) Required(value interface{}) *RadiosControl {
	a.Set("required", value)
	return a
}

func (a *RadiosControl) ResetValue(value interface{}) *RadiosControl {
	a.Set("resetValue", value)
	return a
}

func (a *RadiosControl) Row(value interface{}) *RadiosControl {
	a.Set("row", value)
	return a
}

func (a *RadiosControl) SaveImmediately(value interface{}) *RadiosControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *RadiosControl) SelectFirst(value interface{}) *RadiosControl {
	a.Set("selectFirst", value)
	return a
}

func (a *RadiosControl) Size(value interface{}) *RadiosControl {
	a.Set("size", value)
	return a
}

func (a *RadiosControl) Source(value interface{}) *RadiosControl {
	a.Set("source", value)
	return a
}

func (a *RadiosControl) Static(value interface{}) *RadiosControl {
	a.Set("static", value)
	return a
}

func (a *RadiosControl) StaticClassName(value interface{}) *RadiosControl {
	a.Set("staticClassName", value)
	return a
}

func (a *RadiosControl) StaticInputClassName(value interface{}) *RadiosControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *RadiosControl) StaticLabelClassName(value interface{}) *RadiosControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *RadiosControl) StaticOn(value interface{}) *RadiosControl {
	a.Set("staticOn", value)
	return a
}

func (a *RadiosControl) StaticPlaceholder(value interface{}) *RadiosControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *RadiosControl) StaticSchema(value interface{}) *RadiosControl {
	a.Set("staticSchema", value)
	return a
}

func (a *RadiosControl) Style(value interface{}) *RadiosControl {
	a.Set("style", value)
	return a
}

func (a *RadiosControl) SubmitOnChange(value interface{}) *RadiosControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *RadiosControl) TestIdBuilder(value interface{}) *RadiosControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *RadiosControl) Type(value interface{}) *RadiosControl {
	a.Set("type", value)
	return a
}

func (a *RadiosControl) UseMobileUI(value interface{}) *RadiosControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *RadiosControl) ValidateApi(value interface{}) *RadiosControl {
	a.Set("validateApi", value)
	return a
}

func (a *RadiosControl) ValidateOnChange(value interface{}) *RadiosControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *RadiosControl) ValidationErrors(value interface{}) *RadiosControl {
	a.Set("validationErrors", value)
	return a
}

func (a *RadiosControl) Validations(value interface{}) *RadiosControl {
	a.Set("validations", value)
	return a
}

func (a *RadiosControl) Value(value interface{}) *RadiosControl {
	a.Set("value", value)
	return a
}

func (a *RadiosControl) ValuesNoWrap(value interface{}) *RadiosControl {
	a.Set("valuesNoWrap", value)
	return a
}

func (a *RadiosControl) Visible(value interface{}) *RadiosControl {
	a.Set("visible", value)
	return a
}

func (a *RadiosControl) VisibleOn(value interface{}) *RadiosControl {
	a.Set("visibleOn", value)
	return a
}

func (a *RadiosControl) Width(value interface{}) *RadiosControl {
	a.Set("width", value)
	return a
}
