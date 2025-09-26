package renderers

type UserSelectControl struct {
	*BaseRenderer
}

func NewUserSelectControl() *UserSelectControl {
	a := &UserSelectControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "users-select")
	return a
}

func (a *UserSelectControl) Set(name string, value interface{}) *UserSelectControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *UserSelectControl) AddApi(value interface{}) *UserSelectControl {
	a.Set("addApi", value)
	return a
}

func (a *UserSelectControl) AddControls(value interface{}) *UserSelectControl {
	a.Set("addControls", value)
	return a
}

func (a *UserSelectControl) AddDialog(value interface{}) *UserSelectControl {
	a.Set("addDialog", value)
	return a
}

func (a *UserSelectControl) AutoFill(value interface{}) *UserSelectControl {
	a.Set("autoFill", value)
	return a
}

func (a *UserSelectControl) ClassName(value interface{}) *UserSelectControl {
	a.Set("className", value)
	return a
}

func (a *UserSelectControl) ClearValueOnHidden(value interface{}) *UserSelectControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *UserSelectControl) ClearValueOnSourceChange(value interface{}) *UserSelectControl {
	a.Set("clearValueOnSourceChange", value)
	return a
}

func (a *UserSelectControl) Clearable(value interface{}) *UserSelectControl {
	a.Set("clearable", value)
	return a
}

func (a *UserSelectControl) Creatable(value interface{}) *UserSelectControl {
	a.Set("creatable", value)
	return a
}

func (a *UserSelectControl) CreateBtnLabel(value interface{}) *UserSelectControl {
	a.Set("createBtnLabel", value)
	return a
}

func (a *UserSelectControl) DeferApi(value interface{}) *UserSelectControl {
	a.Set("deferApi", value)
	return a
}

func (a *UserSelectControl) DeferField(value interface{}) *UserSelectControl {
	a.Set("deferField", value)
	return a
}

func (a *UserSelectControl) DeleteApi(value interface{}) *UserSelectControl {
	a.Set("deleteApi", value)
	return a
}

func (a *UserSelectControl) DeleteConfirmText(value interface{}) *UserSelectControl {
	a.Set("deleteConfirmText", value)
	return a
}

func (a *UserSelectControl) Delimiter(value interface{}) *UserSelectControl {
	a.Set("delimiter", value)
	return a
}

func (a *UserSelectControl) Desc(value interface{}) *UserSelectControl {
	a.Set("desc", value)
	return a
}

func (a *UserSelectControl) Description(value interface{}) *UserSelectControl {
	a.Set("description", value)
	return a
}

func (a *UserSelectControl) DescriptionClassName(value interface{}) *UserSelectControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *UserSelectControl) Disabled(value interface{}) *UserSelectControl {
	a.Set("disabled", value)
	return a
}

func (a *UserSelectControl) DisabledOn(value interface{}) *UserSelectControl {
	a.Set("disabledOn", value)
	return a
}

func (a *UserSelectControl) EditApi(value interface{}) *UserSelectControl {
	a.Set("editApi", value)
	return a
}

func (a *UserSelectControl) EditControls(value interface{}) *UserSelectControl {
	a.Set("editControls", value)
	return a
}

func (a *UserSelectControl) EditDialog(value interface{}) *UserSelectControl {
	a.Set("editDialog", value)
	return a
}

func (a *UserSelectControl) Editable(value interface{}) *UserSelectControl {
	a.Set("editable", value)
	return a
}

func (a *UserSelectControl) EditorSetting(value interface{}) *UserSelectControl {
	a.Set("editorSetting", value)
	return a
}

func (a *UserSelectControl) ExtraName(value interface{}) *UserSelectControl {
	a.Set("extraName", value)
	return a
}

func (a *UserSelectControl) ExtractValue(value interface{}) *UserSelectControl {
	a.Set("extractValue", value)
	return a
}

func (a *UserSelectControl) Hidden(value interface{}) *UserSelectControl {
	a.Set("hidden", value)
	return a
}

func (a *UserSelectControl) HiddenOn(value interface{}) *UserSelectControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *UserSelectControl) Hint(value interface{}) *UserSelectControl {
	a.Set("hint", value)
	return a
}

func (a *UserSelectControl) Horizontal(value interface{}) *UserSelectControl {
	a.Set("horizontal", value)
	return a
}

func (a *UserSelectControl) Id(value interface{}) *UserSelectControl {
	a.Set("id", value)
	return a
}

func (a *UserSelectControl) InitAutoFill(value interface{}) *UserSelectControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *UserSelectControl) InitFetch(value interface{}) *UserSelectControl {
	a.Set("initFetch", value)
	return a
}

func (a *UserSelectControl) InitFetchOn(value interface{}) *UserSelectControl {
	a.Set("initFetchOn", value)
	return a
}

func (a *UserSelectControl) Inline(value interface{}) *UserSelectControl {
	a.Set("inline", value)
	return a
}

func (a *UserSelectControl) InputClassName(value interface{}) *UserSelectControl {
	a.Set("inputClassName", value)
	return a
}

func (a *UserSelectControl) JoinValues(value interface{}) *UserSelectControl {
	a.Set("joinValues", value)
	return a
}

func (a *UserSelectControl) Label(value interface{}) *UserSelectControl {
	a.Set("label", value)
	return a
}

func (a *UserSelectControl) LabelAlign(value interface{}) *UserSelectControl {
	a.Set("labelAlign", value)
	return a
}

func (a *UserSelectControl) LabelClassName(value interface{}) *UserSelectControl {
	a.Set("labelClassName", value)
	return a
}

func (a *UserSelectControl) LabelRemark(value interface{}) *UserSelectControl {
	a.Set("labelRemark", value)
	return a
}

func (a *UserSelectControl) LabelWidth(value interface{}) *UserSelectControl {
	a.Set("labelWidth", value)
	return a
}

func (a *UserSelectControl) Mode(value interface{}) *UserSelectControl {
	a.Set("mode", value)
	return a
}

func (a *UserSelectControl) Multiple(value interface{}) *UserSelectControl {
	a.Set("multiple", value)
	return a
}

func (a *UserSelectControl) Name(value interface{}) *UserSelectControl {
	a.Set("name", value)
	return a
}

func (a *UserSelectControl) OnEvent(value interface{}) *UserSelectControl {
	a.Set("onEvent", value)
	return a
}

func (a *UserSelectControl) Options(value interface{}) *UserSelectControl {
	a.Set("options", value)
	return a
}

func (a *UserSelectControl) Placeholder(value interface{}) *UserSelectControl {
	a.Set("placeholder", value)
	return a
}

func (a *UserSelectControl) ReadOnly(value interface{}) *UserSelectControl {
	a.Set("readOnly", value)
	return a
}

func (a *UserSelectControl) ReadOnlyOn(value interface{}) *UserSelectControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *UserSelectControl) Remark(value interface{}) *UserSelectControl {
	a.Set("remark", value)
	return a
}

func (a *UserSelectControl) Removable(value interface{}) *UserSelectControl {
	a.Set("removable", value)
	return a
}

func (a *UserSelectControl) Required(value interface{}) *UserSelectControl {
	a.Set("required", value)
	return a
}

func (a *UserSelectControl) ResetValue(value interface{}) *UserSelectControl {
	a.Set("resetValue", value)
	return a
}

func (a *UserSelectControl) Row(value interface{}) *UserSelectControl {
	a.Set("row", value)
	return a
}

func (a *UserSelectControl) SaveImmediately(value interface{}) *UserSelectControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *UserSelectControl) SelectFirst(value interface{}) *UserSelectControl {
	a.Set("selectFirst", value)
	return a
}

func (a *UserSelectControl) Size(value interface{}) *UserSelectControl {
	a.Set("size", value)
	return a
}

func (a *UserSelectControl) Source(value interface{}) *UserSelectControl {
	a.Set("source", value)
	return a
}

func (a *UserSelectControl) Static(value interface{}) *UserSelectControl {
	a.Set("static", value)
	return a
}

func (a *UserSelectControl) StaticClassName(value interface{}) *UserSelectControl {
	a.Set("staticClassName", value)
	return a
}

func (a *UserSelectControl) StaticInputClassName(value interface{}) *UserSelectControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *UserSelectControl) StaticLabelClassName(value interface{}) *UserSelectControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *UserSelectControl) StaticOn(value interface{}) *UserSelectControl {
	a.Set("staticOn", value)
	return a
}

func (a *UserSelectControl) StaticPlaceholder(value interface{}) *UserSelectControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *UserSelectControl) StaticSchema(value interface{}) *UserSelectControl {
	a.Set("staticSchema", value)
	return a
}

func (a *UserSelectControl) Style(value interface{}) *UserSelectControl {
	a.Set("style", value)
	return a
}

func (a *UserSelectControl) SubmitOnChange(value interface{}) *UserSelectControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *UserSelectControl) TestIdBuilder(value interface{}) *UserSelectControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *UserSelectControl) Type(value interface{}) *UserSelectControl {
	a.Set("type", value)
	return a
}

func (a *UserSelectControl) UseMobileUI(value interface{}) *UserSelectControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *UserSelectControl) ValidateApi(value interface{}) *UserSelectControl {
	a.Set("validateApi", value)
	return a
}

func (a *UserSelectControl) ValidateOnChange(value interface{}) *UserSelectControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *UserSelectControl) ValidationErrors(value interface{}) *UserSelectControl {
	a.Set("validationErrors", value)
	return a
}

func (a *UserSelectControl) Validations(value interface{}) *UserSelectControl {
	a.Set("validations", value)
	return a
}

func (a *UserSelectControl) Value(value interface{}) *UserSelectControl {
	a.Set("value", value)
	return a
}

func (a *UserSelectControl) ValuesNoWrap(value interface{}) *UserSelectControl {
	a.Set("valuesNoWrap", value)
	return a
}

func (a *UserSelectControl) Visible(value interface{}) *UserSelectControl {
	a.Set("visible", value)
	return a
}

func (a *UserSelectControl) VisibleOn(value interface{}) *UserSelectControl {
	a.Set("visibleOn", value)
	return a
}

func (a *UserSelectControl) Width(value interface{}) *UserSelectControl {
	a.Set("width", value)
	return a
}
