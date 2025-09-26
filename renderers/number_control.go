package renderers

type NumberControl struct {
	*BaseRenderer
}

func NewNumberControl() *NumberControl {
	a := &NumberControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "input-number")
	return a
}

func (a *NumberControl) Set(name string, value interface{}) *NumberControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *NumberControl) AutoFill(value interface{}) *NumberControl {
	a.Set("autoFill", value)
	return a
}

func (a *NumberControl) Big(value interface{}) *NumberControl {
	a.Set("big", value)
	return a
}

func (a *NumberControl) BorderMode(value interface{}) *NumberControl {
	a.Set("borderMode", value)
	return a
}

func (a *NumberControl) ClassName(value interface{}) *NumberControl {
	a.Set("className", value)
	return a
}

func (a *NumberControl) ClearValueOnHidden(value interface{}) *NumberControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *NumberControl) Desc(value interface{}) *NumberControl {
	a.Set("desc", value)
	return a
}

func (a *NumberControl) Description(value interface{}) *NumberControl {
	a.Set("description", value)
	return a
}

func (a *NumberControl) DescriptionClassName(value interface{}) *NumberControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *NumberControl) Disabled(value interface{}) *NumberControl {
	a.Set("disabled", value)
	return a
}

func (a *NumberControl) DisabledOn(value interface{}) *NumberControl {
	a.Set("disabledOn", value)
	return a
}

func (a *NumberControl) DisplayMode(value interface{}) *NumberControl {
	a.Set("displayMode", value)
	return a
}

func (a *NumberControl) EditorSetting(value interface{}) *NumberControl {
	a.Set("editorSetting", value)
	return a
}

func (a *NumberControl) ExtraName(value interface{}) *NumberControl {
	a.Set("extraName", value)
	return a
}

func (a *NumberControl) Hidden(value interface{}) *NumberControl {
	a.Set("hidden", value)
	return a
}

func (a *NumberControl) HiddenOn(value interface{}) *NumberControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *NumberControl) Hint(value interface{}) *NumberControl {
	a.Set("hint", value)
	return a
}

func (a *NumberControl) Horizontal(value interface{}) *NumberControl {
	a.Set("horizontal", value)
	return a
}

func (a *NumberControl) Id(value interface{}) *NumberControl {
	a.Set("id", value)
	return a
}

func (a *NumberControl) InitAutoFill(value interface{}) *NumberControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *NumberControl) Inline(value interface{}) *NumberControl {
	a.Set("inline", value)
	return a
}

func (a *NumberControl) InputClassName(value interface{}) *NumberControl {
	a.Set("inputClassName", value)
	return a
}

func (a *NumberControl) Keyboard(value interface{}) *NumberControl {
	a.Set("keyboard", value)
	return a
}

func (a *NumberControl) KilobitSeparator(value interface{}) *NumberControl {
	a.Set("kilobitSeparator", value)
	return a
}

func (a *NumberControl) Label(value interface{}) *NumberControl {
	a.Set("label", value)
	return a
}

func (a *NumberControl) LabelAlign(value interface{}) *NumberControl {
	a.Set("labelAlign", value)
	return a
}

func (a *NumberControl) LabelClassName(value interface{}) *NumberControl {
	a.Set("labelClassName", value)
	return a
}

func (a *NumberControl) LabelRemark(value interface{}) *NumberControl {
	a.Set("labelRemark", value)
	return a
}

func (a *NumberControl) LabelWidth(value interface{}) *NumberControl {
	a.Set("labelWidth", value)
	return a
}

func (a *NumberControl) Max(value interface{}) *NumberControl {
	a.Set("max", value)
	return a
}

func (a *NumberControl) Min(value interface{}) *NumberControl {
	a.Set("min", value)
	return a
}

func (a *NumberControl) Mode(value interface{}) *NumberControl {
	a.Set("mode", value)
	return a
}

func (a *NumberControl) Name(value interface{}) *NumberControl {
	a.Set("name", value)
	return a
}

func (a *NumberControl) OnEvent(value interface{}) *NumberControl {
	a.Set("onEvent", value)
	return a
}

func (a *NumberControl) Placeholder(value interface{}) *NumberControl {
	a.Set("placeholder", value)
	return a
}

func (a *NumberControl) Precision(value interface{}) *NumberControl {
	a.Set("precision", value)
	return a
}

func (a *NumberControl) Prefix(value interface{}) *NumberControl {
	a.Set("prefix", value)
	return a
}

func (a *NumberControl) ReadOnly(value interface{}) *NumberControl {
	a.Set("readOnly", value)
	return a
}

func (a *NumberControl) ReadOnlyOn(value interface{}) *NumberControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *NumberControl) Remark(value interface{}) *NumberControl {
	a.Set("remark", value)
	return a
}

func (a *NumberControl) Required(value interface{}) *NumberControl {
	a.Set("required", value)
	return a
}

func (a *NumberControl) Row(value interface{}) *NumberControl {
	a.Set("row", value)
	return a
}

func (a *NumberControl) SaveImmediately(value interface{}) *NumberControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *NumberControl) ShowAsPercent(value interface{}) *NumberControl {
	a.Set("showAsPercent", value)
	return a
}

func (a *NumberControl) ShowSteps(value interface{}) *NumberControl {
	a.Set("showSteps", value)
	return a
}

func (a *NumberControl) Size(value interface{}) *NumberControl {
	a.Set("size", value)
	return a
}

func (a *NumberControl) Static(value interface{}) *NumberControl {
	a.Set("static", value)
	return a
}

func (a *NumberControl) StaticClassName(value interface{}) *NumberControl {
	a.Set("staticClassName", value)
	return a
}

func (a *NumberControl) StaticInputClassName(value interface{}) *NumberControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *NumberControl) StaticLabelClassName(value interface{}) *NumberControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *NumberControl) StaticOn(value interface{}) *NumberControl {
	a.Set("staticOn", value)
	return a
}

func (a *NumberControl) StaticPlaceholder(value interface{}) *NumberControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *NumberControl) StaticSchema(value interface{}) *NumberControl {
	a.Set("staticSchema", value)
	return a
}

func (a *NumberControl) Step(value interface{}) *NumberControl {
	a.Set("step", value)
	return a
}

func (a *NumberControl) Style(value interface{}) *NumberControl {
	a.Set("style", value)
	return a
}

func (a *NumberControl) SubmitOnChange(value interface{}) *NumberControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *NumberControl) Suffix(value interface{}) *NumberControl {
	a.Set("suffix", value)
	return a
}

func (a *NumberControl) TestIdBuilder(value interface{}) *NumberControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *NumberControl) Type(value interface{}) *NumberControl {
	a.Set("type", value)
	return a
}

func (a *NumberControl) UnitOptions(value interface{}) *NumberControl {
	a.Set("unitOptions", value)
	return a
}

func (a *NumberControl) UseMobileUI(value interface{}) *NumberControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *NumberControl) ValidateApi(value interface{}) *NumberControl {
	a.Set("validateApi", value)
	return a
}

func (a *NumberControl) ValidateOnChange(value interface{}) *NumberControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *NumberControl) ValidationErrors(value interface{}) *NumberControl {
	a.Set("validationErrors", value)
	return a
}

func (a *NumberControl) Validations(value interface{}) *NumberControl {
	a.Set("validations", value)
	return a
}

func (a *NumberControl) Value(value interface{}) *NumberControl {
	a.Set("value", value)
	return a
}

func (a *NumberControl) Visible(value interface{}) *NumberControl {
	a.Set("visible", value)
	return a
}

func (a *NumberControl) VisibleOn(value interface{}) *NumberControl {
	a.Set("visibleOn", value)
	return a
}

func (a *NumberControl) Width(value interface{}) *NumberControl {
	a.Set("width", value)
	return a
}
