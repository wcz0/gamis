package renderers

type SubFormControl struct {
	*BaseRenderer
}

func NewSubFormControl() *SubFormControl {
	a := &SubFormControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "input-sub-form")
	return a
}

func (a *SubFormControl) Set(name string, value interface{}) *SubFormControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *SubFormControl) AddButtonClassName(value interface{}) *SubFormControl {
	a.Set("addButtonClassName", value)
	return a
}

func (a *SubFormControl) AddButtonText(value interface{}) *SubFormControl {
	a.Set("addButtonText", value)
	return a
}

func (a *SubFormControl) Addable(value interface{}) *SubFormControl {
	a.Set("addable", value)
	return a
}

func (a *SubFormControl) AutoFill(value interface{}) *SubFormControl {
	a.Set("autoFill", value)
	return a
}

func (a *SubFormControl) BtnLabel(value interface{}) *SubFormControl {
	a.Set("btnLabel", value)
	return a
}

func (a *SubFormControl) ClassName(value interface{}) *SubFormControl {
	a.Set("className", value)
	return a
}

func (a *SubFormControl) ClearValueOnHidden(value interface{}) *SubFormControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *SubFormControl) Desc(value interface{}) *SubFormControl {
	a.Set("desc", value)
	return a
}

func (a *SubFormControl) Description(value interface{}) *SubFormControl {
	a.Set("description", value)
	return a
}

func (a *SubFormControl) DescriptionClassName(value interface{}) *SubFormControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *SubFormControl) Disabled(value interface{}) *SubFormControl {
	a.Set("disabled", value)
	return a
}

func (a *SubFormControl) DisabledOn(value interface{}) *SubFormControl {
	a.Set("disabledOn", value)
	return a
}

func (a *SubFormControl) Draggable(value interface{}) *SubFormControl {
	a.Set("draggable", value)
	return a
}

func (a *SubFormControl) DraggableTip(value interface{}) *SubFormControl {
	a.Set("draggableTip", value)
	return a
}

func (a *SubFormControl) EditorSetting(value interface{}) *SubFormControl {
	a.Set("editorSetting", value)
	return a
}

func (a *SubFormControl) ExtraName(value interface{}) *SubFormControl {
	a.Set("extraName", value)
	return a
}

func (a *SubFormControl) Form(value interface{}) *SubFormControl {
	a.Set("form", value)
	return a
}

func (a *SubFormControl) Hidden(value interface{}) *SubFormControl {
	a.Set("hidden", value)
	return a
}

func (a *SubFormControl) HiddenOn(value interface{}) *SubFormControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *SubFormControl) Hint(value interface{}) *SubFormControl {
	a.Set("hint", value)
	return a
}

func (a *SubFormControl) Horizontal(value interface{}) *SubFormControl {
	a.Set("horizontal", value)
	return a
}

func (a *SubFormControl) Id(value interface{}) *SubFormControl {
	a.Set("id", value)
	return a
}

func (a *SubFormControl) InitAutoFill(value interface{}) *SubFormControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *SubFormControl) Inline(value interface{}) *SubFormControl {
	a.Set("inline", value)
	return a
}

func (a *SubFormControl) InputClassName(value interface{}) *SubFormControl {
	a.Set("inputClassName", value)
	return a
}

func (a *SubFormControl) ItemClassName(value interface{}) *SubFormControl {
	a.Set("itemClassName", value)
	return a
}

func (a *SubFormControl) ItemsClassName(value interface{}) *SubFormControl {
	a.Set("itemsClassName", value)
	return a
}

func (a *SubFormControl) Label(value interface{}) *SubFormControl {
	a.Set("label", value)
	return a
}

func (a *SubFormControl) LabelAlign(value interface{}) *SubFormControl {
	a.Set("labelAlign", value)
	return a
}

func (a *SubFormControl) LabelClassName(value interface{}) *SubFormControl {
	a.Set("labelClassName", value)
	return a
}

func (a *SubFormControl) LabelField(value interface{}) *SubFormControl {
	a.Set("labelField", value)
	return a
}

func (a *SubFormControl) LabelRemark(value interface{}) *SubFormControl {
	a.Set("labelRemark", value)
	return a
}

func (a *SubFormControl) LabelWidth(value interface{}) *SubFormControl {
	a.Set("labelWidth", value)
	return a
}

func (a *SubFormControl) MaxLength(value interface{}) *SubFormControl {
	a.Set("maxLength", value)
	return a
}

func (a *SubFormControl) MinLength(value interface{}) *SubFormControl {
	a.Set("minLength", value)
	return a
}

func (a *SubFormControl) Mode(value interface{}) *SubFormControl {
	a.Set("mode", value)
	return a
}

func (a *SubFormControl) Multiple(value interface{}) *SubFormControl {
	a.Set("multiple", value)
	return a
}

func (a *SubFormControl) Name(value interface{}) *SubFormControl {
	a.Set("name", value)
	return a
}

func (a *SubFormControl) OnEvent(value interface{}) *SubFormControl {
	a.Set("onEvent", value)
	return a
}

func (a *SubFormControl) Placeholder(value interface{}) *SubFormControl {
	a.Set("placeholder", value)
	return a
}

func (a *SubFormControl) ReadOnly(value interface{}) *SubFormControl {
	a.Set("readOnly", value)
	return a
}

func (a *SubFormControl) ReadOnlyOn(value interface{}) *SubFormControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *SubFormControl) Remark(value interface{}) *SubFormControl {
	a.Set("remark", value)
	return a
}

func (a *SubFormControl) Removable(value interface{}) *SubFormControl {
	a.Set("removable", value)
	return a
}

func (a *SubFormControl) Required(value interface{}) *SubFormControl {
	a.Set("required", value)
	return a
}

func (a *SubFormControl) Row(value interface{}) *SubFormControl {
	a.Set("row", value)
	return a
}

func (a *SubFormControl) SaveImmediately(value interface{}) *SubFormControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *SubFormControl) Scaffold(value interface{}) *SubFormControl {
	a.Set("scaffold", value)
	return a
}

func (a *SubFormControl) ShowErrorMsg(value interface{}) *SubFormControl {
	a.Set("showErrorMsg", value)
	return a
}

func (a *SubFormControl) Size(value interface{}) *SubFormControl {
	a.Set("size", value)
	return a
}

func (a *SubFormControl) Static(value interface{}) *SubFormControl {
	a.Set("static", value)
	return a
}

func (a *SubFormControl) StaticClassName(value interface{}) *SubFormControl {
	a.Set("staticClassName", value)
	return a
}

func (a *SubFormControl) StaticInputClassName(value interface{}) *SubFormControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *SubFormControl) StaticLabelClassName(value interface{}) *SubFormControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *SubFormControl) StaticOn(value interface{}) *SubFormControl {
	a.Set("staticOn", value)
	return a
}

func (a *SubFormControl) StaticPlaceholder(value interface{}) *SubFormControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *SubFormControl) StaticSchema(value interface{}) *SubFormControl {
	a.Set("staticSchema", value)
	return a
}

func (a *SubFormControl) Style(value interface{}) *SubFormControl {
	a.Set("style", value)
	return a
}

func (a *SubFormControl) SubmitOnChange(value interface{}) *SubFormControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *SubFormControl) TestIdBuilder(value interface{}) *SubFormControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *SubFormControl) Type(value interface{}) *SubFormControl {
	a.Set("type", value)
	return a
}

func (a *SubFormControl) UseMobileUI(value interface{}) *SubFormControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *SubFormControl) ValidateApi(value interface{}) *SubFormControl {
	a.Set("validateApi", value)
	return a
}

func (a *SubFormControl) ValidateOnChange(value interface{}) *SubFormControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *SubFormControl) ValidationErrors(value interface{}) *SubFormControl {
	a.Set("validationErrors", value)
	return a
}

func (a *SubFormControl) Validations(value interface{}) *SubFormControl {
	a.Set("validations", value)
	return a
}

func (a *SubFormControl) Value(value interface{}) *SubFormControl {
	a.Set("value", value)
	return a
}

func (a *SubFormControl) Visible(value interface{}) *SubFormControl {
	a.Set("visible", value)
	return a
}

func (a *SubFormControl) VisibleOn(value interface{}) *SubFormControl {
	a.Set("visibleOn", value)
	return a
}

func (a *SubFormControl) Width(value interface{}) *SubFormControl {
	a.Set("width", value)
	return a
}
