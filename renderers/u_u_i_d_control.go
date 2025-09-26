package renderers

type UUIDControl struct {
	*BaseRenderer
}

func NewUUIDControl() *UUIDControl {
	a := &UUIDControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "uuid")
	return a
}

func (a *UUIDControl) Set(name string, value interface{}) *UUIDControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *UUIDControl) AutoFill(value interface{}) *UUIDControl {
	a.Set("autoFill", value)
	return a
}

func (a *UUIDControl) ClassName(value interface{}) *UUIDControl {
	a.Set("className", value)
	return a
}

func (a *UUIDControl) ClearValueOnHidden(value interface{}) *UUIDControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *UUIDControl) Desc(value interface{}) *UUIDControl {
	a.Set("desc", value)
	return a
}

func (a *UUIDControl) Description(value interface{}) *UUIDControl {
	a.Set("description", value)
	return a
}

func (a *UUIDControl) DescriptionClassName(value interface{}) *UUIDControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *UUIDControl) Disabled(value interface{}) *UUIDControl {
	a.Set("disabled", value)
	return a
}

func (a *UUIDControl) DisabledOn(value interface{}) *UUIDControl {
	a.Set("disabledOn", value)
	return a
}

func (a *UUIDControl) EditorSetting(value interface{}) *UUIDControl {
	a.Set("editorSetting", value)
	return a
}

func (a *UUIDControl) ExtraName(value interface{}) *UUIDControl {
	a.Set("extraName", value)
	return a
}

func (a *UUIDControl) Hidden(value interface{}) *UUIDControl {
	a.Set("hidden", value)
	return a
}

func (a *UUIDControl) HiddenOn(value interface{}) *UUIDControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *UUIDControl) Hint(value interface{}) *UUIDControl {
	a.Set("hint", value)
	return a
}

func (a *UUIDControl) Horizontal(value interface{}) *UUIDControl {
	a.Set("horizontal", value)
	return a
}

func (a *UUIDControl) Id(value interface{}) *UUIDControl {
	a.Set("id", value)
	return a
}

func (a *UUIDControl) InitAutoFill(value interface{}) *UUIDControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *UUIDControl) Inline(value interface{}) *UUIDControl {
	a.Set("inline", value)
	return a
}

func (a *UUIDControl) InputClassName(value interface{}) *UUIDControl {
	a.Set("inputClassName", value)
	return a
}

func (a *UUIDControl) Label(value interface{}) *UUIDControl {
	a.Set("label", value)
	return a
}

func (a *UUIDControl) LabelAlign(value interface{}) *UUIDControl {
	a.Set("labelAlign", value)
	return a
}

func (a *UUIDControl) LabelClassName(value interface{}) *UUIDControl {
	a.Set("labelClassName", value)
	return a
}

func (a *UUIDControl) LabelRemark(value interface{}) *UUIDControl {
	a.Set("labelRemark", value)
	return a
}

func (a *UUIDControl) LabelWidth(value interface{}) *UUIDControl {
	a.Set("labelWidth", value)
	return a
}

func (a *UUIDControl) Length(value interface{}) *UUIDControl {
	a.Set("length", value)
	return a
}

func (a *UUIDControl) Mode(value interface{}) *UUIDControl {
	a.Set("mode", value)
	return a
}

func (a *UUIDControl) Name(value interface{}) *UUIDControl {
	a.Set("name", value)
	return a
}

func (a *UUIDControl) OnEvent(value interface{}) *UUIDControl {
	a.Set("onEvent", value)
	return a
}

func (a *UUIDControl) Placeholder(value interface{}) *UUIDControl {
	a.Set("placeholder", value)
	return a
}

func (a *UUIDControl) ReadOnly(value interface{}) *UUIDControl {
	a.Set("readOnly", value)
	return a
}

func (a *UUIDControl) ReadOnlyOn(value interface{}) *UUIDControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *UUIDControl) Remark(value interface{}) *UUIDControl {
	a.Set("remark", value)
	return a
}

func (a *UUIDControl) Required(value interface{}) *UUIDControl {
	a.Set("required", value)
	return a
}

func (a *UUIDControl) Row(value interface{}) *UUIDControl {
	a.Set("row", value)
	return a
}

func (a *UUIDControl) SaveImmediately(value interface{}) *UUIDControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *UUIDControl) Size(value interface{}) *UUIDControl {
	a.Set("size", value)
	return a
}

func (a *UUIDControl) Static(value interface{}) *UUIDControl {
	a.Set("static", value)
	return a
}

func (a *UUIDControl) StaticClassName(value interface{}) *UUIDControl {
	a.Set("staticClassName", value)
	return a
}

func (a *UUIDControl) StaticInputClassName(value interface{}) *UUIDControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *UUIDControl) StaticLabelClassName(value interface{}) *UUIDControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *UUIDControl) StaticOn(value interface{}) *UUIDControl {
	a.Set("staticOn", value)
	return a
}

func (a *UUIDControl) StaticPlaceholder(value interface{}) *UUIDControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *UUIDControl) StaticSchema(value interface{}) *UUIDControl {
	a.Set("staticSchema", value)
	return a
}

func (a *UUIDControl) Style(value interface{}) *UUIDControl {
	a.Set("style", value)
	return a
}

func (a *UUIDControl) SubmitOnChange(value interface{}) *UUIDControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *UUIDControl) TestIdBuilder(value interface{}) *UUIDControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *UUIDControl) Type(value interface{}) *UUIDControl {
	a.Set("type", value)
	return a
}

func (a *UUIDControl) UseMobileUI(value interface{}) *UUIDControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *UUIDControl) ValidateApi(value interface{}) *UUIDControl {
	a.Set("validateApi", value)
	return a
}

func (a *UUIDControl) ValidateOnChange(value interface{}) *UUIDControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *UUIDControl) ValidationErrors(value interface{}) *UUIDControl {
	a.Set("validationErrors", value)
	return a
}

func (a *UUIDControl) Validations(value interface{}) *UUIDControl {
	a.Set("validations", value)
	return a
}

func (a *UUIDControl) Value(value interface{}) *UUIDControl {
	a.Set("value", value)
	return a
}

func (a *UUIDControl) Visible(value interface{}) *UUIDControl {
	a.Set("visible", value)
	return a
}

func (a *UUIDControl) VisibleOn(value interface{}) *UUIDControl {
	a.Set("visibleOn", value)
	return a
}

func (a *UUIDControl) Width(value interface{}) *UUIDControl {
	a.Set("width", value)
	return a
}
