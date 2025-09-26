package renderers

type ArrayControl struct {
	*BaseRenderer
}

func NewArrayControl() *ArrayControl {
	a := &ArrayControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "input-array")
	return a
}

func (a *ArrayControl) Set(name string, value interface{}) *ArrayControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *ArrayControl) AddButtonClassName(value interface{}) *ArrayControl {
	a.Set("addButtonClassName", value)
	return a
}

func (a *ArrayControl) AddButtonText(value interface{}) *ArrayControl {
	a.Set("addButtonText", value)
	return a
}

func (a *ArrayControl) Addable(value interface{}) *ArrayControl {
	a.Set("addable", value)
	return a
}

func (a *ArrayControl) Addattop(value interface{}) *ArrayControl {
	a.Set("addattop", value)
	return a
}

func (a *ArrayControl) AutoFill(value interface{}) *ArrayControl {
	a.Set("autoFill", value)
	return a
}

func (a *ArrayControl) CanAccessSuperData(value interface{}) *ArrayControl {
	a.Set("canAccessSuperData", value)
	return a
}

func (a *ArrayControl) ClassName(value interface{}) *ArrayControl {
	a.Set("className", value)
	return a
}

func (a *ArrayControl) ClearValueOnHidden(value interface{}) *ArrayControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *ArrayControl) DeleteApi(value interface{}) *ArrayControl {
	a.Set("deleteApi", value)
	return a
}

func (a *ArrayControl) DeleteConfirmText(value interface{}) *ArrayControl {
	a.Set("deleteConfirmText", value)
	return a
}

func (a *ArrayControl) Delimiter(value interface{}) *ArrayControl {
	a.Set("delimiter", value)
	return a
}

func (a *ArrayControl) Desc(value interface{}) *ArrayControl {
	a.Set("desc", value)
	return a
}

func (a *ArrayControl) Description(value interface{}) *ArrayControl {
	a.Set("description", value)
	return a
}

func (a *ArrayControl) DescriptionClassName(value interface{}) *ArrayControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *ArrayControl) Disabled(value interface{}) *ArrayControl {
	a.Set("disabled", value)
	return a
}

func (a *ArrayControl) DisabledOn(value interface{}) *ArrayControl {
	a.Set("disabledOn", value)
	return a
}

func (a *ArrayControl) Draggable(value interface{}) *ArrayControl {
	a.Set("draggable", value)
	return a
}

func (a *ArrayControl) DraggableTip(value interface{}) *ArrayControl {
	a.Set("draggableTip", value)
	return a
}

func (a *ArrayControl) EditorSetting(value interface{}) *ArrayControl {
	a.Set("editorSetting", value)
	return a
}

func (a *ArrayControl) ExtraName(value interface{}) *ArrayControl {
	a.Set("extraName", value)
	return a
}

func (a *ArrayControl) Flat(value interface{}) *ArrayControl {
	a.Set("flat", value)
	return a
}

func (a *ArrayControl) FormClassName(value interface{}) *ArrayControl {
	a.Set("formClassName", value)
	return a
}

func (a *ArrayControl) Hidden(value interface{}) *ArrayControl {
	a.Set("hidden", value)
	return a
}

func (a *ArrayControl) HiddenOn(value interface{}) *ArrayControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *ArrayControl) Hint(value interface{}) *ArrayControl {
	a.Set("hint", value)
	return a
}

func (a *ArrayControl) Horizontal(value interface{}) *ArrayControl {
	a.Set("horizontal", value)
	return a
}

func (a *ArrayControl) Id(value interface{}) *ArrayControl {
	a.Set("id", value)
	return a
}

func (a *ArrayControl) InitAutoFill(value interface{}) *ArrayControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *ArrayControl) Inline(value interface{}) *ArrayControl {
	a.Set("inline", value)
	return a
}

func (a *ArrayControl) InputClassName(value interface{}) *ArrayControl {
	a.Set("inputClassName", value)
	return a
}

func (a *ArrayControl) Items(value interface{}) *ArrayControl {
	a.Set("items", value)
	return a
}

func (a *ArrayControl) JoinValues(value interface{}) *ArrayControl {
	a.Set("joinValues", value)
	return a
}

func (a *ArrayControl) Label(value interface{}) *ArrayControl {
	a.Set("label", value)
	return a
}

func (a *ArrayControl) LabelAlign(value interface{}) *ArrayControl {
	a.Set("labelAlign", value)
	return a
}

func (a *ArrayControl) LabelClassName(value interface{}) *ArrayControl {
	a.Set("labelClassName", value)
	return a
}

func (a *ArrayControl) LabelRemark(value interface{}) *ArrayControl {
	a.Set("labelRemark", value)
	return a
}

func (a *ArrayControl) LabelWidth(value interface{}) *ArrayControl {
	a.Set("labelWidth", value)
	return a
}

func (a *ArrayControl) LazyLoad(value interface{}) *ArrayControl {
	a.Set("lazyLoad", value)
	return a
}

func (a *ArrayControl) MaxLength(value interface{}) *ArrayControl {
	a.Set("maxLength", value)
	return a
}

func (a *ArrayControl) Messages(value interface{}) *ArrayControl {
	a.Set("messages", value)
	return a
}

func (a *ArrayControl) MinLength(value interface{}) *ArrayControl {
	a.Set("minLength", value)
	return a
}

func (a *ArrayControl) Mode(value interface{}) *ArrayControl {
	a.Set("mode", value)
	return a
}

func (a *ArrayControl) MultiLine(value interface{}) *ArrayControl {
	a.Set("multiLine", value)
	return a
}

func (a *ArrayControl) Multiple(value interface{}) *ArrayControl {
	a.Set("multiple", value)
	return a
}

func (a *ArrayControl) Name(value interface{}) *ArrayControl {
	a.Set("name", value)
	return a
}

func (a *ArrayControl) NoBorder(value interface{}) *ArrayControl {
	a.Set("noBorder", value)
	return a
}

func (a *ArrayControl) Nullable(value interface{}) *ArrayControl {
	a.Set("nullable", value)
	return a
}

func (a *ArrayControl) OnEvent(value interface{}) *ArrayControl {
	a.Set("onEvent", value)
	return a
}

func (a *ArrayControl) Placeholder(value interface{}) *ArrayControl {
	a.Set("placeholder", value)
	return a
}

func (a *ArrayControl) ReadOnly(value interface{}) *ArrayControl {
	a.Set("readOnly", value)
	return a
}

func (a *ArrayControl) ReadOnlyOn(value interface{}) *ArrayControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *ArrayControl) Remark(value interface{}) *ArrayControl {
	a.Set("remark", value)
	return a
}

func (a *ArrayControl) Removable(value interface{}) *ArrayControl {
	a.Set("removable", value)
	return a
}

func (a *ArrayControl) Required(value interface{}) *ArrayControl {
	a.Set("required", value)
	return a
}

func (a *ArrayControl) Row(value interface{}) *ArrayControl {
	a.Set("row", value)
	return a
}

func (a *ArrayControl) SaveImmediately(value interface{}) *ArrayControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *ArrayControl) Scaffold(value interface{}) *ArrayControl {
	a.Set("scaffold", value)
	return a
}

func (a *ArrayControl) Size(value interface{}) *ArrayControl {
	a.Set("size", value)
	return a
}

func (a *ArrayControl) Static(value interface{}) *ArrayControl {
	a.Set("static", value)
	return a
}

func (a *ArrayControl) StaticClassName(value interface{}) *ArrayControl {
	a.Set("staticClassName", value)
	return a
}

func (a *ArrayControl) StaticInputClassName(value interface{}) *ArrayControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *ArrayControl) StaticLabelClassName(value interface{}) *ArrayControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *ArrayControl) StaticOn(value interface{}) *ArrayControl {
	a.Set("staticOn", value)
	return a
}

func (a *ArrayControl) StaticPlaceholder(value interface{}) *ArrayControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *ArrayControl) StaticSchema(value interface{}) *ArrayControl {
	a.Set("staticSchema", value)
	return a
}

func (a *ArrayControl) StrictMode(value interface{}) *ArrayControl {
	a.Set("strictMode", value)
	return a
}

func (a *ArrayControl) Style(value interface{}) *ArrayControl {
	a.Set("style", value)
	return a
}

func (a *ArrayControl) SubFormHorizontal(value interface{}) *ArrayControl {
	a.Set("subFormHorizontal", value)
	return a
}

func (a *ArrayControl) SubFormMode(value interface{}) *ArrayControl {
	a.Set("subFormMode", value)
	return a
}

func (a *ArrayControl) SubmitOnChange(value interface{}) *ArrayControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *ArrayControl) SyncFields(value interface{}) *ArrayControl {
	a.Set("syncFields", value)
	return a
}

func (a *ArrayControl) TabsLabelTpl(value interface{}) *ArrayControl {
	a.Set("tabsLabelTpl", value)
	return a
}

func (a *ArrayControl) TabsMode(value interface{}) *ArrayControl {
	a.Set("tabsMode", value)
	return a
}

func (a *ArrayControl) TabsStyle(value interface{}) *ArrayControl {
	a.Set("tabsStyle", value)
	return a
}

func (a *ArrayControl) TestIdBuilder(value interface{}) *ArrayControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *ArrayControl) Type(value interface{}) *ArrayControl {
	a.Set("type", value)
	return a
}

func (a *ArrayControl) TypeSwitchable(value interface{}) *ArrayControl {
	a.Set("typeSwitchable", value)
	return a
}

func (a *ArrayControl) UpdatePristineAfterStoreDataReInit(value interface{}) *ArrayControl {
	a.Set("updatePristineAfterStoreDataReInit", value)
	return a
}

func (a *ArrayControl) UseMobileUI(value interface{}) *ArrayControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *ArrayControl) ValidateApi(value interface{}) *ArrayControl {
	a.Set("validateApi", value)
	return a
}

func (a *ArrayControl) ValidateOnChange(value interface{}) *ArrayControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *ArrayControl) ValidationErrors(value interface{}) *ArrayControl {
	a.Set("validationErrors", value)
	return a
}

func (a *ArrayControl) Validations(value interface{}) *ArrayControl {
	a.Set("validations", value)
	return a
}

func (a *ArrayControl) Value(value interface{}) *ArrayControl {
	a.Set("value", value)
	return a
}

func (a *ArrayControl) Visible(value interface{}) *ArrayControl {
	a.Set("visible", value)
	return a
}

func (a *ArrayControl) VisibleOn(value interface{}) *ArrayControl {
	a.Set("visibleOn", value)
	return a
}

func (a *ArrayControl) Width(value interface{}) *ArrayControl {
	a.Set("width", value)
	return a
}
