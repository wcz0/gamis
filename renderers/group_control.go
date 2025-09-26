package renderers

type GroupControl struct {
	*BaseRenderer
}

func NewGroupControl() *GroupControl {
	a := &GroupControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "group")
	return a
}

func (a *GroupControl) Set(name string, value interface{}) *GroupControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *GroupControl) AutoFill(value interface{}) *GroupControl {
	a.Set("autoFill", value)
	return a
}

func (a *GroupControl) Body(value interface{}) *GroupControl {
	a.Set("body", value)
	return a
}

func (a *GroupControl) ClassName(value interface{}) *GroupControl {
	a.Set("className", value)
	return a
}

func (a *GroupControl) ClearValueOnHidden(value interface{}) *GroupControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *GroupControl) Desc(value interface{}) *GroupControl {
	a.Set("desc", value)
	return a
}

func (a *GroupControl) Description(value interface{}) *GroupControl {
	a.Set("description", value)
	return a
}

func (a *GroupControl) DescriptionClassName(value interface{}) *GroupControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *GroupControl) Direction(value interface{}) *GroupControl {
	a.Set("direction", value)
	return a
}

func (a *GroupControl) Disabled(value interface{}) *GroupControl {
	a.Set("disabled", value)
	return a
}

func (a *GroupControl) DisabledOn(value interface{}) *GroupControl {
	a.Set("disabledOn", value)
	return a
}

func (a *GroupControl) EditorSetting(value interface{}) *GroupControl {
	a.Set("editorSetting", value)
	return a
}

func (a *GroupControl) ExtraName(value interface{}) *GroupControl {
	a.Set("extraName", value)
	return a
}

func (a *GroupControl) Gap(value interface{}) *GroupControl {
	a.Set("gap", value)
	return a
}

func (a *GroupControl) Hidden(value interface{}) *GroupControl {
	a.Set("hidden", value)
	return a
}

func (a *GroupControl) HiddenOn(value interface{}) *GroupControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *GroupControl) Hint(value interface{}) *GroupControl {
	a.Set("hint", value)
	return a
}

func (a *GroupControl) Horizontal(value interface{}) *GroupControl {
	a.Set("horizontal", value)
	return a
}

func (a *GroupControl) Id(value interface{}) *GroupControl {
	a.Set("id", value)
	return a
}

func (a *GroupControl) InitAutoFill(value interface{}) *GroupControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *GroupControl) Inline(value interface{}) *GroupControl {
	a.Set("inline", value)
	return a
}

func (a *GroupControl) InputClassName(value interface{}) *GroupControl {
	a.Set("inputClassName", value)
	return a
}

func (a *GroupControl) Label(value interface{}) *GroupControl {
	a.Set("label", value)
	return a
}

func (a *GroupControl) LabelAlign(value interface{}) *GroupControl {
	a.Set("labelAlign", value)
	return a
}

func (a *GroupControl) LabelClassName(value interface{}) *GroupControl {
	a.Set("labelClassName", value)
	return a
}

func (a *GroupControl) LabelRemark(value interface{}) *GroupControl {
	a.Set("labelRemark", value)
	return a
}

func (a *GroupControl) LabelWidth(value interface{}) *GroupControl {
	a.Set("labelWidth", value)
	return a
}

func (a *GroupControl) Mode(value interface{}) *GroupControl {
	a.Set("mode", value)
	return a
}

func (a *GroupControl) Name(value interface{}) *GroupControl {
	a.Set("name", value)
	return a
}

func (a *GroupControl) OnEvent(value interface{}) *GroupControl {
	a.Set("onEvent", value)
	return a
}

func (a *GroupControl) Placeholder(value interface{}) *GroupControl {
	a.Set("placeholder", value)
	return a
}

func (a *GroupControl) ReadOnly(value interface{}) *GroupControl {
	a.Set("readOnly", value)
	return a
}

func (a *GroupControl) ReadOnlyOn(value interface{}) *GroupControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *GroupControl) Remark(value interface{}) *GroupControl {
	a.Set("remark", value)
	return a
}

func (a *GroupControl) Required(value interface{}) *GroupControl {
	a.Set("required", value)
	return a
}

func (a *GroupControl) Row(value interface{}) *GroupControl {
	a.Set("row", value)
	return a
}

func (a *GroupControl) SaveImmediately(value interface{}) *GroupControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *GroupControl) Size(value interface{}) *GroupControl {
	a.Set("size", value)
	return a
}

func (a *GroupControl) Static(value interface{}) *GroupControl {
	a.Set("static", value)
	return a
}

func (a *GroupControl) StaticClassName(value interface{}) *GroupControl {
	a.Set("staticClassName", value)
	return a
}

func (a *GroupControl) StaticInputClassName(value interface{}) *GroupControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *GroupControl) StaticLabelClassName(value interface{}) *GroupControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *GroupControl) StaticOn(value interface{}) *GroupControl {
	a.Set("staticOn", value)
	return a
}

func (a *GroupControl) StaticPlaceholder(value interface{}) *GroupControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *GroupControl) StaticSchema(value interface{}) *GroupControl {
	a.Set("staticSchema", value)
	return a
}

func (a *GroupControl) Style(value interface{}) *GroupControl {
	a.Set("style", value)
	return a
}

func (a *GroupControl) SubFormHorizontal(value interface{}) *GroupControl {
	a.Set("subFormHorizontal", value)
	return a
}

func (a *GroupControl) SubFormMode(value interface{}) *GroupControl {
	a.Set("subFormMode", value)
	return a
}

func (a *GroupControl) SubmitOnChange(value interface{}) *GroupControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *GroupControl) TestIdBuilder(value interface{}) *GroupControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *GroupControl) Type(value interface{}) *GroupControl {
	a.Set("type", value)
	return a
}

func (a *GroupControl) UseMobileUI(value interface{}) *GroupControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *GroupControl) ValidateApi(value interface{}) *GroupControl {
	a.Set("validateApi", value)
	return a
}

func (a *GroupControl) ValidateOnChange(value interface{}) *GroupControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *GroupControl) ValidationErrors(value interface{}) *GroupControl {
	a.Set("validationErrors", value)
	return a
}

func (a *GroupControl) Validations(value interface{}) *GroupControl {
	a.Set("validations", value)
	return a
}

func (a *GroupControl) Value(value interface{}) *GroupControl {
	a.Set("value", value)
	return a
}

func (a *GroupControl) Visible(value interface{}) *GroupControl {
	a.Set("visible", value)
	return a
}

func (a *GroupControl) VisibleOn(value interface{}) *GroupControl {
	a.Set("visibleOn", value)
	return a
}

func (a *GroupControl) Width(value interface{}) *GroupControl {
	a.Set("width", value)
	return a
}
