package renderers

type RepeatControl struct {
	*BaseRenderer
}

func NewRepeatControl() *RepeatControl {
	a := &RepeatControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "input-repeat")
	return a
}

func (a *RepeatControl) Set(name string, value interface{}) *RepeatControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *RepeatControl) AutoFill(value interface{}) *RepeatControl {
	a.Set("autoFill", value)
	return a
}

func (a *RepeatControl) ClassName(value interface{}) *RepeatControl {
	a.Set("className", value)
	return a
}

func (a *RepeatControl) ClearValueOnHidden(value interface{}) *RepeatControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *RepeatControl) Desc(value interface{}) *RepeatControl {
	a.Set("desc", value)
	return a
}

func (a *RepeatControl) Description(value interface{}) *RepeatControl {
	a.Set("description", value)
	return a
}

func (a *RepeatControl) DescriptionClassName(value interface{}) *RepeatControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *RepeatControl) Disabled(value interface{}) *RepeatControl {
	a.Set("disabled", value)
	return a
}

func (a *RepeatControl) DisabledOn(value interface{}) *RepeatControl {
	a.Set("disabledOn", value)
	return a
}

func (a *RepeatControl) EditorSetting(value interface{}) *RepeatControl {
	a.Set("editorSetting", value)
	return a
}

func (a *RepeatControl) ExtraName(value interface{}) *RepeatControl {
	a.Set("extraName", value)
	return a
}

func (a *RepeatControl) Hidden(value interface{}) *RepeatControl {
	a.Set("hidden", value)
	return a
}

func (a *RepeatControl) HiddenOn(value interface{}) *RepeatControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *RepeatControl) Hint(value interface{}) *RepeatControl {
	a.Set("hint", value)
	return a
}

func (a *RepeatControl) Horizontal(value interface{}) *RepeatControl {
	a.Set("horizontal", value)
	return a
}

func (a *RepeatControl) Id(value interface{}) *RepeatControl {
	a.Set("id", value)
	return a
}

func (a *RepeatControl) InitAutoFill(value interface{}) *RepeatControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *RepeatControl) Inline(value interface{}) *RepeatControl {
	a.Set("inline", value)
	return a
}

func (a *RepeatControl) InputClassName(value interface{}) *RepeatControl {
	a.Set("inputClassName", value)
	return a
}

func (a *RepeatControl) Label(value interface{}) *RepeatControl {
	a.Set("label", value)
	return a
}

func (a *RepeatControl) LabelAlign(value interface{}) *RepeatControl {
	a.Set("labelAlign", value)
	return a
}

func (a *RepeatControl) LabelClassName(value interface{}) *RepeatControl {
	a.Set("labelClassName", value)
	return a
}

func (a *RepeatControl) LabelRemark(value interface{}) *RepeatControl {
	a.Set("labelRemark", value)
	return a
}

func (a *RepeatControl) LabelWidth(value interface{}) *RepeatControl {
	a.Set("labelWidth", value)
	return a
}

func (a *RepeatControl) Mode(value interface{}) *RepeatControl {
	a.Set("mode", value)
	return a
}

func (a *RepeatControl) Name(value interface{}) *RepeatControl {
	a.Set("name", value)
	return a
}

func (a *RepeatControl) OnEvent(value interface{}) *RepeatControl {
	a.Set("onEvent", value)
	return a
}

func (a *RepeatControl) Options(value interface{}) *RepeatControl {
	a.Set("options", value)
	return a
}

func (a *RepeatControl) Placeholder(value interface{}) *RepeatControl {
	a.Set("placeholder", value)
	return a
}

func (a *RepeatControl) ReadOnly(value interface{}) *RepeatControl {
	a.Set("readOnly", value)
	return a
}

func (a *RepeatControl) ReadOnlyOn(value interface{}) *RepeatControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *RepeatControl) Remark(value interface{}) *RepeatControl {
	a.Set("remark", value)
	return a
}

func (a *RepeatControl) Required(value interface{}) *RepeatControl {
	a.Set("required", value)
	return a
}

func (a *RepeatControl) Row(value interface{}) *RepeatControl {
	a.Set("row", value)
	return a
}

func (a *RepeatControl) SaveImmediately(value interface{}) *RepeatControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *RepeatControl) Size(value interface{}) *RepeatControl {
	a.Set("size", value)
	return a
}

func (a *RepeatControl) Static(value interface{}) *RepeatControl {
	a.Set("static", value)
	return a
}

func (a *RepeatControl) StaticClassName(value interface{}) *RepeatControl {
	a.Set("staticClassName", value)
	return a
}

func (a *RepeatControl) StaticInputClassName(value interface{}) *RepeatControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *RepeatControl) StaticLabelClassName(value interface{}) *RepeatControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *RepeatControl) StaticOn(value interface{}) *RepeatControl {
	a.Set("staticOn", value)
	return a
}

func (a *RepeatControl) StaticPlaceholder(value interface{}) *RepeatControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *RepeatControl) StaticSchema(value interface{}) *RepeatControl {
	a.Set("staticSchema", value)
	return a
}

func (a *RepeatControl) Style(value interface{}) *RepeatControl {
	a.Set("style", value)
	return a
}

func (a *RepeatControl) SubmitOnChange(value interface{}) *RepeatControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *RepeatControl) TestIdBuilder(value interface{}) *RepeatControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *RepeatControl) Type(value interface{}) *RepeatControl {
	a.Set("type", value)
	return a
}

func (a *RepeatControl) UseMobileUI(value interface{}) *RepeatControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *RepeatControl) ValidateApi(value interface{}) *RepeatControl {
	a.Set("validateApi", value)
	return a
}

func (a *RepeatControl) ValidateOnChange(value interface{}) *RepeatControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *RepeatControl) ValidationErrors(value interface{}) *RepeatControl {
	a.Set("validationErrors", value)
	return a
}

func (a *RepeatControl) Validations(value interface{}) *RepeatControl {
	a.Set("validations", value)
	return a
}

func (a *RepeatControl) Value(value interface{}) *RepeatControl {
	a.Set("value", value)
	return a
}

func (a *RepeatControl) Visible(value interface{}) *RepeatControl {
	a.Set("visible", value)
	return a
}

func (a *RepeatControl) VisibleOn(value interface{}) *RepeatControl {
	a.Set("visibleOn", value)
	return a
}

func (a *RepeatControl) Width(value interface{}) *RepeatControl {
	a.Set("width", value)
	return a
}
