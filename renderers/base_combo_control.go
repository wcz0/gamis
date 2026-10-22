package renderers

type BaseComboControl struct {
	*BaseRenderer
}

func NewBaseComboControl() *BaseComboControl {
	a := &BaseComboControl{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *BaseComboControl) Set(name string, value interface{}) *BaseComboControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *BaseComboControl) AddButtonClassName(value interface{}) *BaseComboControl {
	a.Set("addButtonClassName", value)
	return a
}

func (a *BaseComboControl) AddButtonText(value interface{}) *BaseComboControl {
	a.Set("addButtonText", value)
	return a
}

func (a *BaseComboControl) Addable(value interface{}) *BaseComboControl {
	a.Set("addable", value)
	return a
}

func (a *BaseComboControl) Addattop(value interface{}) *BaseComboControl {
	a.Set("addattop", value)
	return a
}

func (a *BaseComboControl) AutoFill(value interface{}) *BaseComboControl {
	a.Set("autoFill", value)
	return a
}

func (a *BaseComboControl) CanAccessSuperData(value interface{}) *BaseComboControl {
	a.Set("canAccessSuperData", value)
	return a
}

func (a *BaseComboControl) ClassName(value interface{}) *BaseComboControl {
	a.Set("className", value)
	return a
}

func (a *BaseComboControl) ClearValueOnHidden(value interface{}) *BaseComboControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *BaseComboControl) Conditions(value interface{}) *BaseComboControl {
	a.Set("conditions", value)
	return a
}

func (a *BaseComboControl) DeleteApi(value interface{}) *BaseComboControl {
	a.Set("deleteApi", value)
	return a
}

func (a *BaseComboControl) DeleteConfirmText(value interface{}) *BaseComboControl {
	a.Set("deleteConfirmText", value)
	return a
}

func (a *BaseComboControl) Delimiter(value interface{}) *BaseComboControl {
	a.Set("delimiter", value)
	return a
}

func (a *BaseComboControl) Desc(value interface{}) *BaseComboControl {
	a.Set("desc", value)
	return a
}

func (a *BaseComboControl) Description(value interface{}) *BaseComboControl {
	a.Set("description", value)
	return a
}

func (a *BaseComboControl) DescriptionClassName(value interface{}) *BaseComboControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *BaseComboControl) Disabled(value interface{}) *BaseComboControl {
	a.Set("disabled", value)
	return a
}

func (a *BaseComboControl) DisabledOn(value interface{}) *BaseComboControl {
	a.Set("disabledOn", value)
	return a
}

func (a *BaseComboControl) Draggable(value interface{}) *BaseComboControl {
	a.Set("draggable", value)
	return a
}

func (a *BaseComboControl) DraggableTip(value interface{}) *BaseComboControl {
	a.Set("draggableTip", value)
	return a
}

func (a *BaseComboControl) EditorSetting(value interface{}) *BaseComboControl {
	a.Set("editorSetting", value)
	return a
}

func (a *BaseComboControl) ExtraName(value interface{}) *BaseComboControl {
	a.Set("extraName", value)
	return a
}

func (a *BaseComboControl) Flat(value interface{}) *BaseComboControl {
	a.Set("flat", value)
	return a
}

func (a *BaseComboControl) FormClassName(value interface{}) *BaseComboControl {
	a.Set("formClassName", value)
	return a
}

func (a *BaseComboControl) Hidden(value interface{}) *BaseComboControl {
	a.Set("hidden", value)
	return a
}

func (a *BaseComboControl) HiddenOn(value interface{}) *BaseComboControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *BaseComboControl) Hint(value interface{}) *BaseComboControl {
	a.Set("hint", value)
	return a
}

func (a *BaseComboControl) Horizontal(value interface{}) *BaseComboControl {
	a.Set("horizontal", value)
	return a
}

func (a *BaseComboControl) Id(value interface{}) *BaseComboControl {
	a.Set("id", value)
	return a
}

func (a *BaseComboControl) InitAutoFill(value interface{}) *BaseComboControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *BaseComboControl) Inline(value interface{}) *BaseComboControl {
	a.Set("inline", value)
	return a
}

func (a *BaseComboControl) InputClassName(value interface{}) *BaseComboControl {
	a.Set("inputClassName", value)
	return a
}

func (a *BaseComboControl) Items(value interface{}) *BaseComboControl {
	a.Set("items", value)
	return a
}

func (a *BaseComboControl) JoinValues(value interface{}) *BaseComboControl {
	a.Set("joinValues", value)
	return a
}

func (a *BaseComboControl) Label(value interface{}) *BaseComboControl {
	a.Set("label", value)
	return a
}

func (a *BaseComboControl) LabelAlign(value interface{}) *BaseComboControl {
	a.Set("labelAlign", value)
	return a
}

func (a *BaseComboControl) LabelClassName(value interface{}) *BaseComboControl {
	a.Set("labelClassName", value)
	return a
}

func (a *BaseComboControl) LabelOverflow(value interface{}) *BaseComboControl {
	a.Set("labelOverflow", value)
	return a
}

func (a *BaseComboControl) LabelRemark(value interface{}) *BaseComboControl {
	a.Set("labelRemark", value)
	return a
}

func (a *BaseComboControl) LabelWidth(value interface{}) *BaseComboControl {
	a.Set("labelWidth", value)
	return a
}

func (a *BaseComboControl) LazyLoad(value interface{}) *BaseComboControl {
	a.Set("lazyLoad", value)
	return a
}

func (a *BaseComboControl) MaxLength(value interface{}) *BaseComboControl {
	a.Set("maxLength", value)
	return a
}

func (a *BaseComboControl) Messages(value interface{}) *BaseComboControl {
	a.Set("messages", value)
	return a
}

func (a *BaseComboControl) MinLength(value interface{}) *BaseComboControl {
	a.Set("minLength", value)
	return a
}

func (a *BaseComboControl) Mode(value interface{}) *BaseComboControl {
	a.Set("mode", value)
	return a
}

func (a *BaseComboControl) MultiLine(value interface{}) *BaseComboControl {
	a.Set("multiLine", value)
	return a
}

func (a *BaseComboControl) Multiple(value interface{}) *BaseComboControl {
	a.Set("multiple", value)
	return a
}

func (a *BaseComboControl) Name(value interface{}) *BaseComboControl {
	a.Set("name", value)
	return a
}

func (a *BaseComboControl) NoBorder(value interface{}) *BaseComboControl {
	a.Set("noBorder", value)
	return a
}

func (a *BaseComboControl) Nullable(value interface{}) *BaseComboControl {
	a.Set("nullable", value)
	return a
}

func (a *BaseComboControl) OnEvent(value interface{}) *BaseComboControl {
	a.Set("onEvent", value)
	return a
}

func (a *BaseComboControl) PerPage(value interface{}) *BaseComboControl {
	a.Set("perPage", value)
	return a
}

func (a *BaseComboControl) Placeholder(value interface{}) *BaseComboControl {
	a.Set("placeholder", value)
	return a
}

func (a *BaseComboControl) ReadOnly(value interface{}) *BaseComboControl {
	a.Set("readOnly", value)
	return a
}

func (a *BaseComboControl) ReadOnlyOn(value interface{}) *BaseComboControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *BaseComboControl) Remark(value interface{}) *BaseComboControl {
	a.Set("remark", value)
	return a
}

func (a *BaseComboControl) Removable(value interface{}) *BaseComboControl {
	a.Set("removable", value)
	return a
}

func (a *BaseComboControl) Required(value interface{}) *BaseComboControl {
	a.Set("required", value)
	return a
}

func (a *BaseComboControl) Row(value interface{}) *BaseComboControl {
	a.Set("row", value)
	return a
}

func (a *BaseComboControl) SaveImmediately(value interface{}) *BaseComboControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *BaseComboControl) Size(value interface{}) *BaseComboControl {
	a.Set("size", value)
	return a
}

func (a *BaseComboControl) Static(value interface{}) *BaseComboControl {
	a.Set("static", value)
	return a
}

func (a *BaseComboControl) StaticClassName(value interface{}) *BaseComboControl {
	a.Set("staticClassName", value)
	return a
}

func (a *BaseComboControl) StaticInputClassName(value interface{}) *BaseComboControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *BaseComboControl) StaticLabelClassName(value interface{}) *BaseComboControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *BaseComboControl) StaticOn(value interface{}) *BaseComboControl {
	a.Set("staticOn", value)
	return a
}

func (a *BaseComboControl) StaticPlaceholder(value interface{}) *BaseComboControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *BaseComboControl) StaticSchema(value interface{}) *BaseComboControl {
	a.Set("staticSchema", value)
	return a
}

func (a *BaseComboControl) StrictMode(value interface{}) *BaseComboControl {
	a.Set("strictMode", value)
	return a
}

func (a *BaseComboControl) Style(value interface{}) *BaseComboControl {
	a.Set("style", value)
	return a
}

func (a *BaseComboControl) SubFormHorizontal(value interface{}) *BaseComboControl {
	a.Set("subFormHorizontal", value)
	return a
}

func (a *BaseComboControl) SubFormMode(value interface{}) *BaseComboControl {
	a.Set("subFormMode", value)
	return a
}

func (a *BaseComboControl) SubmitOnChange(value interface{}) *BaseComboControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *BaseComboControl) SyncFields(value interface{}) *BaseComboControl {
	a.Set("syncFields", value)
	return a
}

func (a *BaseComboControl) TabsLabelTpl(value interface{}) *BaseComboControl {
	a.Set("tabsLabelTpl", value)
	return a
}

func (a *BaseComboControl) TabsMode(value interface{}) *BaseComboControl {
	a.Set("tabsMode", value)
	return a
}

func (a *BaseComboControl) TabsStyle(value interface{}) *BaseComboControl {
	a.Set("tabsStyle", value)
	return a
}

func (a *BaseComboControl) TestIdBuilder(value interface{}) *BaseComboControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *BaseComboControl) Type(value interface{}) *BaseComboControl {
	a.Set("type", value)
	return a
}

func (a *BaseComboControl) TypeSwitchable(value interface{}) *BaseComboControl {
	a.Set("typeSwitchable", value)
	return a
}

func (a *BaseComboControl) UpdatePristineAfterStoreDataReInit(value interface{}) *BaseComboControl {
	a.Set("updatePristineAfterStoreDataReInit", value)
	return a
}

func (a *BaseComboControl) UseMobileUI(value interface{}) *BaseComboControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *BaseComboControl) ValidateApi(value interface{}) *BaseComboControl {
	a.Set("validateApi", value)
	return a
}

func (a *BaseComboControl) ValidateOnChange(value interface{}) *BaseComboControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *BaseComboControl) ValidationErrors(value interface{}) *BaseComboControl {
	a.Set("validationErrors", value)
	return a
}

func (a *BaseComboControl) Validations(value interface{}) *BaseComboControl {
	a.Set("validations", value)
	return a
}

func (a *BaseComboControl) Value(value interface{}) *BaseComboControl {
	a.Set("value", value)
	return a
}

func (a *BaseComboControl) Visible(value interface{}) *BaseComboControl {
	a.Set("visible", value)
	return a
}

func (a *BaseComboControl) VisibleOn(value interface{}) *BaseComboControl {
	a.Set("visibleOn", value)
	return a
}

func (a *BaseComboControl) Width(value interface{}) *BaseComboControl {
	a.Set("width", value)
	return a
}
