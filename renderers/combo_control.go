package renderers

type ComboControl struct {
	*BaseRenderer
}

func NewComboControl() *ComboControl {
	a := &ComboControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "combo")
	return a
}

func (a *ComboControl) Set(name string, value interface{}) *ComboControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *ComboControl) AddButtonClassName(value interface{}) *ComboControl {
	a.Set("addButtonClassName", value)
	return a
}

func (a *ComboControl) AddButtonText(value interface{}) *ComboControl {
	a.Set("addButtonText", value)
	return a
}

func (a *ComboControl) Addable(value interface{}) *ComboControl {
	a.Set("addable", value)
	return a
}

func (a *ComboControl) Addattop(value interface{}) *ComboControl {
	a.Set("addattop", value)
	return a
}

func (a *ComboControl) AutoFill(value interface{}) *ComboControl {
	a.Set("autoFill", value)
	return a
}

func (a *ComboControl) CanAccessSuperData(value interface{}) *ComboControl {
	a.Set("canAccessSuperData", value)
	return a
}

func (a *ComboControl) ClassName(value interface{}) *ComboControl {
	a.Set("className", value)
	return a
}

func (a *ComboControl) ClearValueOnHidden(value interface{}) *ComboControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *ComboControl) Conditions(value interface{}) *ComboControl {
	a.Set("conditions", value)
	return a
}

func (a *ComboControl) DeleteApi(value interface{}) *ComboControl {
	a.Set("deleteApi", value)
	return a
}

func (a *ComboControl) DeleteConfirmText(value interface{}) *ComboControl {
	a.Set("deleteConfirmText", value)
	return a
}

func (a *ComboControl) Delimiter(value interface{}) *ComboControl {
	a.Set("delimiter", value)
	return a
}

func (a *ComboControl) Desc(value interface{}) *ComboControl {
	a.Set("desc", value)
	return a
}

func (a *ComboControl) Description(value interface{}) *ComboControl {
	a.Set("description", value)
	return a
}

func (a *ComboControl) DescriptionClassName(value interface{}) *ComboControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *ComboControl) Disabled(value interface{}) *ComboControl {
	a.Set("disabled", value)
	return a
}

func (a *ComboControl) DisabledOn(value interface{}) *ComboControl {
	a.Set("disabledOn", value)
	return a
}

func (a *ComboControl) Draggable(value interface{}) *ComboControl {
	a.Set("draggable", value)
	return a
}

func (a *ComboControl) DraggableTip(value interface{}) *ComboControl {
	a.Set("draggableTip", value)
	return a
}

func (a *ComboControl) EditorSetting(value interface{}) *ComboControl {
	a.Set("editorSetting", value)
	return a
}

func (a *ComboControl) ExtraName(value interface{}) *ComboControl {
	a.Set("extraName", value)
	return a
}

func (a *ComboControl) Flat(value interface{}) *ComboControl {
	a.Set("flat", value)
	return a
}

func (a *ComboControl) FormClassName(value interface{}) *ComboControl {
	a.Set("formClassName", value)
	return a
}

func (a *ComboControl) Hidden(value interface{}) *ComboControl {
	a.Set("hidden", value)
	return a
}

func (a *ComboControl) HiddenOn(value interface{}) *ComboControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *ComboControl) Hint(value interface{}) *ComboControl {
	a.Set("hint", value)
	return a
}

func (a *ComboControl) Horizontal(value interface{}) *ComboControl {
	a.Set("horizontal", value)
	return a
}

func (a *ComboControl) Id(value interface{}) *ComboControl {
	a.Set("id", value)
	return a
}

func (a *ComboControl) InitAutoFill(value interface{}) *ComboControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *ComboControl) Inline(value interface{}) *ComboControl {
	a.Set("inline", value)
	return a
}

func (a *ComboControl) InputClassName(value interface{}) *ComboControl {
	a.Set("inputClassName", value)
	return a
}

func (a *ComboControl) Items(value interface{}) *ComboControl {
	a.Set("items", value)
	return a
}

func (a *ComboControl) JoinValues(value interface{}) *ComboControl {
	a.Set("joinValues", value)
	return a
}

func (a *ComboControl) Label(value interface{}) *ComboControl {
	a.Set("label", value)
	return a
}

func (a *ComboControl) LabelAlign(value interface{}) *ComboControl {
	a.Set("labelAlign", value)
	return a
}

func (a *ComboControl) LabelClassName(value interface{}) *ComboControl {
	a.Set("labelClassName", value)
	return a
}

func (a *ComboControl) LabelRemark(value interface{}) *ComboControl {
	a.Set("labelRemark", value)
	return a
}

func (a *ComboControl) LabelWidth(value interface{}) *ComboControl {
	a.Set("labelWidth", value)
	return a
}

func (a *ComboControl) LazyLoad(value interface{}) *ComboControl {
	a.Set("lazyLoad", value)
	return a
}

func (a *ComboControl) MaxLength(value interface{}) *ComboControl {
	a.Set("maxLength", value)
	return a
}

func (a *ComboControl) Messages(value interface{}) *ComboControl {
	a.Set("messages", value)
	return a
}

func (a *ComboControl) MinLength(value interface{}) *ComboControl {
	a.Set("minLength", value)
	return a
}

func (a *ComboControl) Mode(value interface{}) *ComboControl {
	a.Set("mode", value)
	return a
}

func (a *ComboControl) MultiLine(value interface{}) *ComboControl {
	a.Set("multiLine", value)
	return a
}

func (a *ComboControl) Multiple(value interface{}) *ComboControl {
	a.Set("multiple", value)
	return a
}

func (a *ComboControl) Name(value interface{}) *ComboControl {
	a.Set("name", value)
	return a
}

func (a *ComboControl) NoBorder(value interface{}) *ComboControl {
	a.Set("noBorder", value)
	return a
}

func (a *ComboControl) Nullable(value interface{}) *ComboControl {
	a.Set("nullable", value)
	return a
}

func (a *ComboControl) OnEvent(value interface{}) *ComboControl {
	a.Set("onEvent", value)
	return a
}

func (a *ComboControl) Placeholder(value interface{}) *ComboControl {
	a.Set("placeholder", value)
	return a
}

func (a *ComboControl) ReadOnly(value interface{}) *ComboControl {
	a.Set("readOnly", value)
	return a
}

func (a *ComboControl) ReadOnlyOn(value interface{}) *ComboControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *ComboControl) Remark(value interface{}) *ComboControl {
	a.Set("remark", value)
	return a
}

func (a *ComboControl) Removable(value interface{}) *ComboControl {
	a.Set("removable", value)
	return a
}

func (a *ComboControl) Required(value interface{}) *ComboControl {
	a.Set("required", value)
	return a
}

func (a *ComboControl) Row(value interface{}) *ComboControl {
	a.Set("row", value)
	return a
}

func (a *ComboControl) SaveImmediately(value interface{}) *ComboControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *ComboControl) Scaffold(value interface{}) *ComboControl {
	a.Set("scaffold", value)
	return a
}

func (a *ComboControl) Size(value interface{}) *ComboControl {
	a.Set("size", value)
	return a
}

func (a *ComboControl) Static(value interface{}) *ComboControl {
	a.Set("static", value)
	return a
}

func (a *ComboControl) StaticClassName(value interface{}) *ComboControl {
	a.Set("staticClassName", value)
	return a
}

func (a *ComboControl) StaticInputClassName(value interface{}) *ComboControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *ComboControl) StaticLabelClassName(value interface{}) *ComboControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *ComboControl) StaticOn(value interface{}) *ComboControl {
	a.Set("staticOn", value)
	return a
}

func (a *ComboControl) StaticPlaceholder(value interface{}) *ComboControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *ComboControl) StaticSchema(value interface{}) *ComboControl {
	a.Set("staticSchema", value)
	return a
}

func (a *ComboControl) StrictMode(value interface{}) *ComboControl {
	a.Set("strictMode", value)
	return a
}

func (a *ComboControl) Style(value interface{}) *ComboControl {
	a.Set("style", value)
	return a
}

func (a *ComboControl) SubFormHorizontal(value interface{}) *ComboControl {
	a.Set("subFormHorizontal", value)
	return a
}

func (a *ComboControl) SubFormMode(value interface{}) *ComboControl {
	a.Set("subFormMode", value)
	return a
}

func (a *ComboControl) SubmitOnChange(value interface{}) *ComboControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *ComboControl) SyncFields(value interface{}) *ComboControl {
	a.Set("syncFields", value)
	return a
}

func (a *ComboControl) TabsLabelTpl(value interface{}) *ComboControl {
	a.Set("tabsLabelTpl", value)
	return a
}

func (a *ComboControl) TabsMode(value interface{}) *ComboControl {
	a.Set("tabsMode", value)
	return a
}

func (a *ComboControl) TabsStyle(value interface{}) *ComboControl {
	a.Set("tabsStyle", value)
	return a
}

func (a *ComboControl) TestIdBuilder(value interface{}) *ComboControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *ComboControl) Type(value interface{}) *ComboControl {
	a.Set("type", value)
	return a
}

func (a *ComboControl) TypeSwitchable(value interface{}) *ComboControl {
	a.Set("typeSwitchable", value)
	return a
}

func (a *ComboControl) UpdatePristineAfterStoreDataReInit(value interface{}) *ComboControl {
	a.Set("updatePristineAfterStoreDataReInit", value)
	return a
}

func (a *ComboControl) UseMobileUI(value interface{}) *ComboControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *ComboControl) ValidateApi(value interface{}) *ComboControl {
	a.Set("validateApi", value)
	return a
}

func (a *ComboControl) ValidateOnChange(value interface{}) *ComboControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *ComboControl) ValidationErrors(value interface{}) *ComboControl {
	a.Set("validationErrors", value)
	return a
}

func (a *ComboControl) Validations(value interface{}) *ComboControl {
	a.Set("validations", value)
	return a
}

func (a *ComboControl) Value(value interface{}) *ComboControl {
	a.Set("value", value)
	return a
}

func (a *ComboControl) Visible(value interface{}) *ComboControl {
	a.Set("visible", value)
	return a
}

func (a *ComboControl) VisibleOn(value interface{}) *ComboControl {
	a.Set("visibleOn", value)
	return a
}

func (a *ComboControl) Width(value interface{}) *ComboControl {
	a.Set("width", value)
	return a
}
