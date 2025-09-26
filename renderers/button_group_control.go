package renderers

type ButtonGroupControl struct {
	*BaseRenderer
}

func NewButtonGroupControl() *ButtonGroupControl {
	a := &ButtonGroupControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "button-group-select")
	return a
}

func (a *ButtonGroupControl) Set(name string, value interface{}) *ButtonGroupControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *ButtonGroupControl) AddApi(value interface{}) *ButtonGroupControl {
	a.Set("addApi", value)
	return a
}

func (a *ButtonGroupControl) AddControls(value interface{}) *ButtonGroupControl {
	a.Set("addControls", value)
	return a
}

func (a *ButtonGroupControl) AddDialog(value interface{}) *ButtonGroupControl {
	a.Set("addDialog", value)
	return a
}

func (a *ButtonGroupControl) AutoFill(value interface{}) *ButtonGroupControl {
	a.Set("autoFill", value)
	return a
}

func (a *ButtonGroupControl) BtnActiveClassName(value interface{}) *ButtonGroupControl {
	a.Set("btnActiveClassName", value)
	return a
}

func (a *ButtonGroupControl) BtnActiveLevel(value interface{}) *ButtonGroupControl {
	a.Set("btnActiveLevel", value)
	return a
}

func (a *ButtonGroupControl) BtnClassName(value interface{}) *ButtonGroupControl {
	a.Set("btnClassName", value)
	return a
}

func (a *ButtonGroupControl) BtnLevel(value interface{}) *ButtonGroupControl {
	a.Set("btnLevel", value)
	return a
}

func (a *ButtonGroupControl) Buttons(value interface{}) *ButtonGroupControl {
	a.Set("buttons", value)
	return a
}

func (a *ButtonGroupControl) ClassName(value interface{}) *ButtonGroupControl {
	a.Set("className", value)
	return a
}

func (a *ButtonGroupControl) ClearValueOnHidden(value interface{}) *ButtonGroupControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *ButtonGroupControl) ClearValueOnSourceChange(value interface{}) *ButtonGroupControl {
	a.Set("clearValueOnSourceChange", value)
	return a
}

func (a *ButtonGroupControl) Clearable(value interface{}) *ButtonGroupControl {
	a.Set("clearable", value)
	return a
}

func (a *ButtonGroupControl) Creatable(value interface{}) *ButtonGroupControl {
	a.Set("creatable", value)
	return a
}

func (a *ButtonGroupControl) CreateBtnLabel(value interface{}) *ButtonGroupControl {
	a.Set("createBtnLabel", value)
	return a
}

func (a *ButtonGroupControl) DeferApi(value interface{}) *ButtonGroupControl {
	a.Set("deferApi", value)
	return a
}

func (a *ButtonGroupControl) DeferField(value interface{}) *ButtonGroupControl {
	a.Set("deferField", value)
	return a
}

func (a *ButtonGroupControl) DeleteApi(value interface{}) *ButtonGroupControl {
	a.Set("deleteApi", value)
	return a
}

func (a *ButtonGroupControl) DeleteConfirmText(value interface{}) *ButtonGroupControl {
	a.Set("deleteConfirmText", value)
	return a
}

func (a *ButtonGroupControl) Delimiter(value interface{}) *ButtonGroupControl {
	a.Set("delimiter", value)
	return a
}

func (a *ButtonGroupControl) Desc(value interface{}) *ButtonGroupControl {
	a.Set("desc", value)
	return a
}

func (a *ButtonGroupControl) Description(value interface{}) *ButtonGroupControl {
	a.Set("description", value)
	return a
}

func (a *ButtonGroupControl) DescriptionClassName(value interface{}) *ButtonGroupControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *ButtonGroupControl) Disabled(value interface{}) *ButtonGroupControl {
	a.Set("disabled", value)
	return a
}

func (a *ButtonGroupControl) DisabledOn(value interface{}) *ButtonGroupControl {
	a.Set("disabledOn", value)
	return a
}

func (a *ButtonGroupControl) EditApi(value interface{}) *ButtonGroupControl {
	a.Set("editApi", value)
	return a
}

func (a *ButtonGroupControl) EditControls(value interface{}) *ButtonGroupControl {
	a.Set("editControls", value)
	return a
}

func (a *ButtonGroupControl) EditDialog(value interface{}) *ButtonGroupControl {
	a.Set("editDialog", value)
	return a
}

func (a *ButtonGroupControl) Editable(value interface{}) *ButtonGroupControl {
	a.Set("editable", value)
	return a
}

func (a *ButtonGroupControl) EditorSetting(value interface{}) *ButtonGroupControl {
	a.Set("editorSetting", value)
	return a
}

func (a *ButtonGroupControl) ExtraName(value interface{}) *ButtonGroupControl {
	a.Set("extraName", value)
	return a
}

func (a *ButtonGroupControl) ExtractValue(value interface{}) *ButtonGroupControl {
	a.Set("extractValue", value)
	return a
}

func (a *ButtonGroupControl) Hidden(value interface{}) *ButtonGroupControl {
	a.Set("hidden", value)
	return a
}

func (a *ButtonGroupControl) HiddenOn(value interface{}) *ButtonGroupControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *ButtonGroupControl) Hint(value interface{}) *ButtonGroupControl {
	a.Set("hint", value)
	return a
}

func (a *ButtonGroupControl) Horizontal(value interface{}) *ButtonGroupControl {
	a.Set("horizontal", value)
	return a
}

func (a *ButtonGroupControl) Id(value interface{}) *ButtonGroupControl {
	a.Set("id", value)
	return a
}

func (a *ButtonGroupControl) InitAutoFill(value interface{}) *ButtonGroupControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *ButtonGroupControl) InitFetch(value interface{}) *ButtonGroupControl {
	a.Set("initFetch", value)
	return a
}

func (a *ButtonGroupControl) InitFetchOn(value interface{}) *ButtonGroupControl {
	a.Set("initFetchOn", value)
	return a
}

func (a *ButtonGroupControl) Inline(value interface{}) *ButtonGroupControl {
	a.Set("inline", value)
	return a
}

func (a *ButtonGroupControl) InputClassName(value interface{}) *ButtonGroupControl {
	a.Set("inputClassName", value)
	return a
}

func (a *ButtonGroupControl) JoinValues(value interface{}) *ButtonGroupControl {
	a.Set("joinValues", value)
	return a
}

func (a *ButtonGroupControl) Label(value interface{}) *ButtonGroupControl {
	a.Set("label", value)
	return a
}

func (a *ButtonGroupControl) LabelAlign(value interface{}) *ButtonGroupControl {
	a.Set("labelAlign", value)
	return a
}

func (a *ButtonGroupControl) LabelClassName(value interface{}) *ButtonGroupControl {
	a.Set("labelClassName", value)
	return a
}

func (a *ButtonGroupControl) LabelRemark(value interface{}) *ButtonGroupControl {
	a.Set("labelRemark", value)
	return a
}

func (a *ButtonGroupControl) LabelWidth(value interface{}) *ButtonGroupControl {
	a.Set("labelWidth", value)
	return a
}

func (a *ButtonGroupControl) Mode(value interface{}) *ButtonGroupControl {
	a.Set("mode", value)
	return a
}

func (a *ButtonGroupControl) Multiple(value interface{}) *ButtonGroupControl {
	a.Set("multiple", value)
	return a
}

func (a *ButtonGroupControl) Name(value interface{}) *ButtonGroupControl {
	a.Set("name", value)
	return a
}

func (a *ButtonGroupControl) OnEvent(value interface{}) *ButtonGroupControl {
	a.Set("onEvent", value)
	return a
}

func (a *ButtonGroupControl) Options(value interface{}) *ButtonGroupControl {
	a.Set("options", value)
	return a
}

func (a *ButtonGroupControl) Placeholder(value interface{}) *ButtonGroupControl {
	a.Set("placeholder", value)
	return a
}

func (a *ButtonGroupControl) ReadOnly(value interface{}) *ButtonGroupControl {
	a.Set("readOnly", value)
	return a
}

func (a *ButtonGroupControl) ReadOnlyOn(value interface{}) *ButtonGroupControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *ButtonGroupControl) Remark(value interface{}) *ButtonGroupControl {
	a.Set("remark", value)
	return a
}

func (a *ButtonGroupControl) Removable(value interface{}) *ButtonGroupControl {
	a.Set("removable", value)
	return a
}

func (a *ButtonGroupControl) Required(value interface{}) *ButtonGroupControl {
	a.Set("required", value)
	return a
}

func (a *ButtonGroupControl) ResetValue(value interface{}) *ButtonGroupControl {
	a.Set("resetValue", value)
	return a
}

func (a *ButtonGroupControl) Row(value interface{}) *ButtonGroupControl {
	a.Set("row", value)
	return a
}

func (a *ButtonGroupControl) SaveImmediately(value interface{}) *ButtonGroupControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *ButtonGroupControl) SelectFirst(value interface{}) *ButtonGroupControl {
	a.Set("selectFirst", value)
	return a
}

func (a *ButtonGroupControl) Size(value interface{}) *ButtonGroupControl {
	a.Set("size", value)
	return a
}

func (a *ButtonGroupControl) Source(value interface{}) *ButtonGroupControl {
	a.Set("source", value)
	return a
}

func (a *ButtonGroupControl) Static(value interface{}) *ButtonGroupControl {
	a.Set("static", value)
	return a
}

func (a *ButtonGroupControl) StaticClassName(value interface{}) *ButtonGroupControl {
	a.Set("staticClassName", value)
	return a
}

func (a *ButtonGroupControl) StaticInputClassName(value interface{}) *ButtonGroupControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *ButtonGroupControl) StaticLabelClassName(value interface{}) *ButtonGroupControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *ButtonGroupControl) StaticOn(value interface{}) *ButtonGroupControl {
	a.Set("staticOn", value)
	return a
}

func (a *ButtonGroupControl) StaticPlaceholder(value interface{}) *ButtonGroupControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *ButtonGroupControl) StaticSchema(value interface{}) *ButtonGroupControl {
	a.Set("staticSchema", value)
	return a
}

func (a *ButtonGroupControl) Style(value interface{}) *ButtonGroupControl {
	a.Set("style", value)
	return a
}

func (a *ButtonGroupControl) SubmitOnChange(value interface{}) *ButtonGroupControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *ButtonGroupControl) TestIdBuilder(value interface{}) *ButtonGroupControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *ButtonGroupControl) Testid(value interface{}) *ButtonGroupControl {
	a.Set("testid", value)
	return a
}

func (a *ButtonGroupControl) Tiled(value interface{}) *ButtonGroupControl {
	a.Set("tiled", value)
	return a
}

func (a *ButtonGroupControl) Type(value interface{}) *ButtonGroupControl {
	a.Set("type", value)
	return a
}

func (a *ButtonGroupControl) UseMobileUI(value interface{}) *ButtonGroupControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *ButtonGroupControl) ValidateApi(value interface{}) *ButtonGroupControl {
	a.Set("validateApi", value)
	return a
}

func (a *ButtonGroupControl) ValidateOnChange(value interface{}) *ButtonGroupControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *ButtonGroupControl) ValidationErrors(value interface{}) *ButtonGroupControl {
	a.Set("validationErrors", value)
	return a
}

func (a *ButtonGroupControl) Validations(value interface{}) *ButtonGroupControl {
	a.Set("validations", value)
	return a
}

func (a *ButtonGroupControl) Value(value interface{}) *ButtonGroupControl {
	a.Set("value", value)
	return a
}

func (a *ButtonGroupControl) ValuesNoWrap(value interface{}) *ButtonGroupControl {
	a.Set("valuesNoWrap", value)
	return a
}

func (a *ButtonGroupControl) Vertical(value interface{}) *ButtonGroupControl {
	a.Set("vertical", value)
	return a
}

func (a *ButtonGroupControl) Visible(value interface{}) *ButtonGroupControl {
	a.Set("visible", value)
	return a
}

func (a *ButtonGroupControl) VisibleOn(value interface{}) *ButtonGroupControl {
	a.Set("visibleOn", value)
	return a
}

func (a *ButtonGroupControl) Width(value interface{}) *ButtonGroupControl {
	a.Set("width", value)
	return a
}
