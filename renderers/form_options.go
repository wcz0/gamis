package renderers

type FormOptions struct {
	*BaseRenderer
}

func NewFormOptions() *FormOptions {
	a := &FormOptions{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *FormOptions) Set(name string, value interface{}) *FormOptions {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *FormOptions) AddApi(value interface{}) *FormOptions {
	a.Set("addApi", value)
	return a
}

func (a *FormOptions) AddControls(value interface{}) *FormOptions {
	a.Set("addControls", value)
	return a
}

func (a *FormOptions) AddDialog(value interface{}) *FormOptions {
	a.Set("addDialog", value)
	return a
}

func (a *FormOptions) AutoFill(value interface{}) *FormOptions {
	a.Set("autoFill", value)
	return a
}

func (a *FormOptions) CheckAll(value interface{}) *FormOptions {
	a.Set("checkAll", value)
	return a
}

func (a *FormOptions) ClassName(value interface{}) *FormOptions {
	a.Set("className", value)
	return a
}

func (a *FormOptions) ClearValueOnHidden(value interface{}) *FormOptions {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *FormOptions) ClearValueOnSourceChange(value interface{}) *FormOptions {
	a.Set("clearValueOnSourceChange", value)
	return a
}

func (a *FormOptions) Clearable(value interface{}) *FormOptions {
	a.Set("clearable", value)
	return a
}

func (a *FormOptions) Creatable(value interface{}) *FormOptions {
	a.Set("creatable", value)
	return a
}

func (a *FormOptions) CreateBtnLabel(value interface{}) *FormOptions {
	a.Set("createBtnLabel", value)
	return a
}

func (a *FormOptions) DeferApi(value interface{}) *FormOptions {
	a.Set("deferApi", value)
	return a
}

func (a *FormOptions) DeferField(value interface{}) *FormOptions {
	a.Set("deferField", value)
	return a
}

func (a *FormOptions) DeleteApi(value interface{}) *FormOptions {
	a.Set("deleteApi", value)
	return a
}

func (a *FormOptions) DeleteConfirmText(value interface{}) *FormOptions {
	a.Set("deleteConfirmText", value)
	return a
}

func (a *FormOptions) Delimiter(value interface{}) *FormOptions {
	a.Set("delimiter", value)
	return a
}

func (a *FormOptions) Desc(value interface{}) *FormOptions {
	a.Set("desc", value)
	return a
}

func (a *FormOptions) Description(value interface{}) *FormOptions {
	a.Set("description", value)
	return a
}

func (a *FormOptions) DescriptionClassName(value interface{}) *FormOptions {
	a.Set("descriptionClassName", value)
	return a
}

func (a *FormOptions) Disabled(value interface{}) *FormOptions {
	a.Set("disabled", value)
	return a
}

func (a *FormOptions) DisabledOn(value interface{}) *FormOptions {
	a.Set("disabledOn", value)
	return a
}

func (a *FormOptions) EditApi(value interface{}) *FormOptions {
	a.Set("editApi", value)
	return a
}

func (a *FormOptions) EditControls(value interface{}) *FormOptions {
	a.Set("editControls", value)
	return a
}

func (a *FormOptions) EditDialog(value interface{}) *FormOptions {
	a.Set("editDialog", value)
	return a
}

func (a *FormOptions) Editable(value interface{}) *FormOptions {
	a.Set("editable", value)
	return a
}

func (a *FormOptions) EditorSetting(value interface{}) *FormOptions {
	a.Set("editorSetting", value)
	return a
}

func (a *FormOptions) ExtraName(value interface{}) *FormOptions {
	a.Set("extraName", value)
	return a
}

func (a *FormOptions) ExtractValue(value interface{}) *FormOptions {
	a.Set("extractValue", value)
	return a
}

func (a *FormOptions) Hidden(value interface{}) *FormOptions {
	a.Set("hidden", value)
	return a
}

func (a *FormOptions) HiddenOn(value interface{}) *FormOptions {
	a.Set("hiddenOn", value)
	return a
}

func (a *FormOptions) Hint(value interface{}) *FormOptions {
	a.Set("hint", value)
	return a
}

func (a *FormOptions) Horizontal(value interface{}) *FormOptions {
	a.Set("horizontal", value)
	return a
}

func (a *FormOptions) Id(value interface{}) *FormOptions {
	a.Set("id", value)
	return a
}

func (a *FormOptions) InitAutoFill(value interface{}) *FormOptions {
	a.Set("initAutoFill", value)
	return a
}

func (a *FormOptions) InitFetch(value interface{}) *FormOptions {
	a.Set("initFetch", value)
	return a
}

func (a *FormOptions) InitFetchOn(value interface{}) *FormOptions {
	a.Set("initFetchOn", value)
	return a
}

func (a *FormOptions) Inline(value interface{}) *FormOptions {
	a.Set("inline", value)
	return a
}

func (a *FormOptions) InputClassName(value interface{}) *FormOptions {
	a.Set("inputClassName", value)
	return a
}

func (a *FormOptions) JoinValues(value interface{}) *FormOptions {
	a.Set("joinValues", value)
	return a
}

func (a *FormOptions) Label(value interface{}) *FormOptions {
	a.Set("label", value)
	return a
}

func (a *FormOptions) LabelAlign(value interface{}) *FormOptions {
	a.Set("labelAlign", value)
	return a
}

func (a *FormOptions) LabelClassName(value interface{}) *FormOptions {
	a.Set("labelClassName", value)
	return a
}

func (a *FormOptions) LabelOverflow(value interface{}) *FormOptions {
	a.Set("labelOverflow", value)
	return a
}

func (a *FormOptions) LabelRemark(value interface{}) *FormOptions {
	a.Set("labelRemark", value)
	return a
}

func (a *FormOptions) LabelWidth(value interface{}) *FormOptions {
	a.Set("labelWidth", value)
	return a
}

func (a *FormOptions) Mode(value interface{}) *FormOptions {
	a.Set("mode", value)
	return a
}

func (a *FormOptions) Multiple(value interface{}) *FormOptions {
	a.Set("multiple", value)
	return a
}

func (a *FormOptions) Name(value interface{}) *FormOptions {
	a.Set("name", value)
	return a
}

func (a *FormOptions) OnEvent(value interface{}) *FormOptions {
	a.Set("onEvent", value)
	return a
}

func (a *FormOptions) Options(value interface{}) *FormOptions {
	a.Set("options", value)
	return a
}

func (a *FormOptions) Placeholder(value interface{}) *FormOptions {
	a.Set("placeholder", value)
	return a
}

func (a *FormOptions) ReadOnly(value interface{}) *FormOptions {
	a.Set("readOnly", value)
	return a
}

func (a *FormOptions) ReadOnlyOn(value interface{}) *FormOptions {
	a.Set("readOnlyOn", value)
	return a
}

func (a *FormOptions) Remark(value interface{}) *FormOptions {
	a.Set("remark", value)
	return a
}

func (a *FormOptions) Removable(value interface{}) *FormOptions {
	a.Set("removable", value)
	return a
}

func (a *FormOptions) Required(value interface{}) *FormOptions {
	a.Set("required", value)
	return a
}

func (a *FormOptions) ResetValue(value interface{}) *FormOptions {
	a.Set("resetValue", value)
	return a
}

func (a *FormOptions) Row(value interface{}) *FormOptions {
	a.Set("row", value)
	return a
}

func (a *FormOptions) SelectFirst(value interface{}) *FormOptions {
	a.Set("selectFirst", value)
	return a
}

func (a *FormOptions) Size(value interface{}) *FormOptions {
	a.Set("size", value)
	return a
}

func (a *FormOptions) Source(value interface{}) *FormOptions {
	a.Set("source", value)
	return a
}

func (a *FormOptions) Static(value interface{}) *FormOptions {
	a.Set("static", value)
	return a
}

func (a *FormOptions) StaticClassName(value interface{}) *FormOptions {
	a.Set("staticClassName", value)
	return a
}

func (a *FormOptions) StaticInputClassName(value interface{}) *FormOptions {
	a.Set("staticInputClassName", value)
	return a
}

func (a *FormOptions) StaticLabelClassName(value interface{}) *FormOptions {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *FormOptions) StaticOn(value interface{}) *FormOptions {
	a.Set("staticOn", value)
	return a
}

func (a *FormOptions) StaticPlaceholder(value interface{}) *FormOptions {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *FormOptions) StaticSchema(value interface{}) *FormOptions {
	a.Set("staticSchema", value)
	return a
}

func (a *FormOptions) Style(value interface{}) *FormOptions {
	a.Set("style", value)
	return a
}

func (a *FormOptions) SubmitOnChange(value interface{}) *FormOptions {
	a.Set("submitOnChange", value)
	return a
}

func (a *FormOptions) Type(value interface{}) *FormOptions {
	a.Set("type", value)
	return a
}

func (a *FormOptions) UseMobileUI(value interface{}) *FormOptions {
	a.Set("useMobileUI", value)
	return a
}

func (a *FormOptions) ValidateApi(value interface{}) *FormOptions {
	a.Set("validateApi", value)
	return a
}

func (a *FormOptions) ValidateOnChange(value interface{}) *FormOptions {
	a.Set("validateOnChange", value)
	return a
}

func (a *FormOptions) ValidationErrors(value interface{}) *FormOptions {
	a.Set("validationErrors", value)
	return a
}

func (a *FormOptions) Validations(value interface{}) *FormOptions {
	a.Set("validations", value)
	return a
}

func (a *FormOptions) Value(value interface{}) *FormOptions {
	a.Set("value", value)
	return a
}

func (a *FormOptions) ValuesNoWrap(value interface{}) *FormOptions {
	a.Set("valuesNoWrap", value)
	return a
}

func (a *FormOptions) Visible(value interface{}) *FormOptions {
	a.Set("visible", value)
	return a
}

func (a *FormOptions) VisibleOn(value interface{}) *FormOptions {
	a.Set("visibleOn", value)
	return a
}
