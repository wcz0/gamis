package renderers

type TagControl struct {
	*BaseRenderer
}

func NewTagControl() *TagControl {
	a := &TagControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "input-tag")
	return a
}

func (a *TagControl) Set(name string, value interface{}) *TagControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *TagControl) AddApi(value interface{}) *TagControl {
	a.Set("addApi", value)
	return a
}

func (a *TagControl) AddControls(value interface{}) *TagControl {
	a.Set("addControls", value)
	return a
}

func (a *TagControl) AddDialog(value interface{}) *TagControl {
	a.Set("addDialog", value)
	return a
}

func (a *TagControl) AutoFill(value interface{}) *TagControl {
	a.Set("autoFill", value)
	return a
}

func (a *TagControl) ClassName(value interface{}) *TagControl {
	a.Set("className", value)
	return a
}

func (a *TagControl) ClearValueOnHidden(value interface{}) *TagControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *TagControl) ClearValueOnSourceChange(value interface{}) *TagControl {
	a.Set("clearValueOnSourceChange", value)
	return a
}

func (a *TagControl) Clearable(value interface{}) *TagControl {
	a.Set("clearable", value)
	return a
}

func (a *TagControl) Creatable(value interface{}) *TagControl {
	a.Set("creatable", value)
	return a
}

func (a *TagControl) CreateBtnLabel(value interface{}) *TagControl {
	a.Set("createBtnLabel", value)
	return a
}

func (a *TagControl) DeferApi(value interface{}) *TagControl {
	a.Set("deferApi", value)
	return a
}

func (a *TagControl) DeferField(value interface{}) *TagControl {
	a.Set("deferField", value)
	return a
}

func (a *TagControl) DeleteApi(value interface{}) *TagControl {
	a.Set("deleteApi", value)
	return a
}

func (a *TagControl) DeleteConfirmText(value interface{}) *TagControl {
	a.Set("deleteConfirmText", value)
	return a
}

func (a *TagControl) Delimiter(value interface{}) *TagControl {
	a.Set("delimiter", value)
	return a
}

func (a *TagControl) Desc(value interface{}) *TagControl {
	a.Set("desc", value)
	return a
}

func (a *TagControl) Description(value interface{}) *TagControl {
	a.Set("description", value)
	return a
}

func (a *TagControl) DescriptionClassName(value interface{}) *TagControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *TagControl) Disabled(value interface{}) *TagControl {
	a.Set("disabled", value)
	return a
}

func (a *TagControl) DisabledOn(value interface{}) *TagControl {
	a.Set("disabledOn", value)
	return a
}

func (a *TagControl) Dropdown(value interface{}) *TagControl {
	a.Set("dropdown", value)
	return a
}

func (a *TagControl) EditApi(value interface{}) *TagControl {
	a.Set("editApi", value)
	return a
}

func (a *TagControl) EditControls(value interface{}) *TagControl {
	a.Set("editControls", value)
	return a
}

func (a *TagControl) EditDialog(value interface{}) *TagControl {
	a.Set("editDialog", value)
	return a
}

func (a *TagControl) Editable(value interface{}) *TagControl {
	a.Set("editable", value)
	return a
}

func (a *TagControl) EditorSetting(value interface{}) *TagControl {
	a.Set("editorSetting", value)
	return a
}

func (a *TagControl) EnableBatchAdd(value interface{}) *TagControl {
	a.Set("enableBatchAdd", value)
	return a
}

func (a *TagControl) ExtraName(value interface{}) *TagControl {
	a.Set("extraName", value)
	return a
}

func (a *TagControl) ExtractValue(value interface{}) *TagControl {
	a.Set("extractValue", value)
	return a
}

func (a *TagControl) Hidden(value interface{}) *TagControl {
	a.Set("hidden", value)
	return a
}

func (a *TagControl) HiddenOn(value interface{}) *TagControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *TagControl) Hint(value interface{}) *TagControl {
	a.Set("hint", value)
	return a
}

func (a *TagControl) Horizontal(value interface{}) *TagControl {
	a.Set("horizontal", value)
	return a
}

func (a *TagControl) Id(value interface{}) *TagControl {
	a.Set("id", value)
	return a
}

func (a *TagControl) InitAutoFill(value interface{}) *TagControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *TagControl) InitFetch(value interface{}) *TagControl {
	a.Set("initFetch", value)
	return a
}

func (a *TagControl) InitFetchOn(value interface{}) *TagControl {
	a.Set("initFetchOn", value)
	return a
}

func (a *TagControl) Inline(value interface{}) *TagControl {
	a.Set("inline", value)
	return a
}

func (a *TagControl) InputClassName(value interface{}) *TagControl {
	a.Set("inputClassName", value)
	return a
}

func (a *TagControl) JoinValues(value interface{}) *TagControl {
	a.Set("joinValues", value)
	return a
}

func (a *TagControl) Label(value interface{}) *TagControl {
	a.Set("label", value)
	return a
}

func (a *TagControl) LabelAlign(value interface{}) *TagControl {
	a.Set("labelAlign", value)
	return a
}

func (a *TagControl) LabelClassName(value interface{}) *TagControl {
	a.Set("labelClassName", value)
	return a
}

func (a *TagControl) LabelRemark(value interface{}) *TagControl {
	a.Set("labelRemark", value)
	return a
}

func (a *TagControl) LabelWidth(value interface{}) *TagControl {
	a.Set("labelWidth", value)
	return a
}

func (a *TagControl) Max(value interface{}) *TagControl {
	a.Set("max", value)
	return a
}

func (a *TagControl) MaxTagCount(value interface{}) *TagControl {
	a.Set("maxTagCount", value)
	return a
}

func (a *TagControl) MaxTagLength(value interface{}) *TagControl {
	a.Set("maxTagLength", value)
	return a
}

func (a *TagControl) Mode(value interface{}) *TagControl {
	a.Set("mode", value)
	return a
}

func (a *TagControl) Multiple(value interface{}) *TagControl {
	a.Set("multiple", value)
	return a
}

func (a *TagControl) Name(value interface{}) *TagControl {
	a.Set("name", value)
	return a
}

func (a *TagControl) OnEvent(value interface{}) *TagControl {
	a.Set("onEvent", value)
	return a
}

func (a *TagControl) Options(value interface{}) *TagControl {
	a.Set("options", value)
	return a
}

func (a *TagControl) OptionsTip(value interface{}) *TagControl {
	a.Set("optionsTip", value)
	return a
}

func (a *TagControl) OverflowTagPopover(value interface{}) *TagControl {
	a.Set("overflowTagPopover", value)
	return a
}

func (a *TagControl) Placeholder(value interface{}) *TagControl {
	a.Set("placeholder", value)
	return a
}

func (a *TagControl) ReadOnly(value interface{}) *TagControl {
	a.Set("readOnly", value)
	return a
}

func (a *TagControl) ReadOnlyOn(value interface{}) *TagControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *TagControl) Remark(value interface{}) *TagControl {
	a.Set("remark", value)
	return a
}

func (a *TagControl) Removable(value interface{}) *TagControl {
	a.Set("removable", value)
	return a
}

func (a *TagControl) Required(value interface{}) *TagControl {
	a.Set("required", value)
	return a
}

func (a *TagControl) ResetValue(value interface{}) *TagControl {
	a.Set("resetValue", value)
	return a
}

func (a *TagControl) Row(value interface{}) *TagControl {
	a.Set("row", value)
	return a
}

func (a *TagControl) SaveImmediately(value interface{}) *TagControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *TagControl) SelectFirst(value interface{}) *TagControl {
	a.Set("selectFirst", value)
	return a
}

func (a *TagControl) Separator(value interface{}) *TagControl {
	a.Set("separator", value)
	return a
}

func (a *TagControl) Size(value interface{}) *TagControl {
	a.Set("size", value)
	return a
}

func (a *TagControl) Source(value interface{}) *TagControl {
	a.Set("source", value)
	return a
}

func (a *TagControl) Static(value interface{}) *TagControl {
	a.Set("static", value)
	return a
}

func (a *TagControl) StaticClassName(value interface{}) *TagControl {
	a.Set("staticClassName", value)
	return a
}

func (a *TagControl) StaticInputClassName(value interface{}) *TagControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *TagControl) StaticLabelClassName(value interface{}) *TagControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *TagControl) StaticOn(value interface{}) *TagControl {
	a.Set("staticOn", value)
	return a
}

func (a *TagControl) StaticPlaceholder(value interface{}) *TagControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *TagControl) StaticSchema(value interface{}) *TagControl {
	a.Set("staticSchema", value)
	return a
}

func (a *TagControl) Style(value interface{}) *TagControl {
	a.Set("style", value)
	return a
}

func (a *TagControl) SubmitOnChange(value interface{}) *TagControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *TagControl) TestIdBuilder(value interface{}) *TagControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *TagControl) Type(value interface{}) *TagControl {
	a.Set("type", value)
	return a
}

func (a *TagControl) UseMobileUI(value interface{}) *TagControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *TagControl) ValidateApi(value interface{}) *TagControl {
	a.Set("validateApi", value)
	return a
}

func (a *TagControl) ValidateOnChange(value interface{}) *TagControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *TagControl) ValidationErrors(value interface{}) *TagControl {
	a.Set("validationErrors", value)
	return a
}

func (a *TagControl) Validations(value interface{}) *TagControl {
	a.Set("validations", value)
	return a
}

func (a *TagControl) Value(value interface{}) *TagControl {
	a.Set("value", value)
	return a
}

func (a *TagControl) ValuesNoWrap(value interface{}) *TagControl {
	a.Set("valuesNoWrap", value)
	return a
}

func (a *TagControl) Visible(value interface{}) *TagControl {
	a.Set("visible", value)
	return a
}

func (a *TagControl) VisibleOn(value interface{}) *TagControl {
	a.Set("visibleOn", value)
	return a
}

func (a *TagControl) Width(value interface{}) *TagControl {
	a.Set("width", value)
	return a
}
