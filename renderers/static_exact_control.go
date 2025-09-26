package renderers

type StaticExactControl struct {
	*BaseRenderer
}

func NewStaticExactControl() *StaticExactControl {
	a := &StaticExactControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "static")
	return a
}

func (a *StaticExactControl) Set(name string, value interface{}) *StaticExactControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *StaticExactControl) AutoFill(value interface{}) *StaticExactControl {
	a.Set("autoFill", value)
	return a
}

func (a *StaticExactControl) BorderMode(value interface{}) *StaticExactControl {
	a.Set("borderMode", value)
	return a
}

func (a *StaticExactControl) ClassName(value interface{}) *StaticExactControl {
	a.Set("className", value)
	return a
}

func (a *StaticExactControl) ClearValueOnHidden(value interface{}) *StaticExactControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *StaticExactControl) Copyable(value interface{}) *StaticExactControl {
	a.Set("copyable", value)
	return a
}

func (a *StaticExactControl) Desc(value interface{}) *StaticExactControl {
	a.Set("desc", value)
	return a
}

func (a *StaticExactControl) Description(value interface{}) *StaticExactControl {
	a.Set("description", value)
	return a
}

func (a *StaticExactControl) DescriptionClassName(value interface{}) *StaticExactControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *StaticExactControl) Disabled(value interface{}) *StaticExactControl {
	a.Set("disabled", value)
	return a
}

func (a *StaticExactControl) DisabledOn(value interface{}) *StaticExactControl {
	a.Set("disabledOn", value)
	return a
}

func (a *StaticExactControl) EditorSetting(value interface{}) *StaticExactControl {
	a.Set("editorSetting", value)
	return a
}

func (a *StaticExactControl) ExtraName(value interface{}) *StaticExactControl {
	a.Set("extraName", value)
	return a
}

func (a *StaticExactControl) Hidden(value interface{}) *StaticExactControl {
	a.Set("hidden", value)
	return a
}

func (a *StaticExactControl) HiddenOn(value interface{}) *StaticExactControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *StaticExactControl) Hint(value interface{}) *StaticExactControl {
	a.Set("hint", value)
	return a
}

func (a *StaticExactControl) Horizontal(value interface{}) *StaticExactControl {
	a.Set("horizontal", value)
	return a
}

func (a *StaticExactControl) Id(value interface{}) *StaticExactControl {
	a.Set("id", value)
	return a
}

func (a *StaticExactControl) InitAutoFill(value interface{}) *StaticExactControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *StaticExactControl) Inline(value interface{}) *StaticExactControl {
	a.Set("inline", value)
	return a
}

func (a *StaticExactControl) InputClassName(value interface{}) *StaticExactControl {
	a.Set("inputClassName", value)
	return a
}

func (a *StaticExactControl) Label(value interface{}) *StaticExactControl {
	a.Set("label", value)
	return a
}

func (a *StaticExactControl) LabelAlign(value interface{}) *StaticExactControl {
	a.Set("labelAlign", value)
	return a
}

func (a *StaticExactControl) LabelClassName(value interface{}) *StaticExactControl {
	a.Set("labelClassName", value)
	return a
}

func (a *StaticExactControl) LabelRemark(value interface{}) *StaticExactControl {
	a.Set("labelRemark", value)
	return a
}

func (a *StaticExactControl) LabelWidth(value interface{}) *StaticExactControl {
	a.Set("labelWidth", value)
	return a
}

func (a *StaticExactControl) Mode(value interface{}) *StaticExactControl {
	a.Set("mode", value)
	return a
}

func (a *StaticExactControl) Name(value interface{}) *StaticExactControl {
	a.Set("name", value)
	return a
}

func (a *StaticExactControl) OnEvent(value interface{}) *StaticExactControl {
	a.Set("onEvent", value)
	return a
}

func (a *StaticExactControl) Placeholder(value interface{}) *StaticExactControl {
	a.Set("placeholder", value)
	return a
}

func (a *StaticExactControl) PopOver(value interface{}) *StaticExactControl {
	a.Set("popOver", value)
	return a
}

func (a *StaticExactControl) QuickEdit(value interface{}) *StaticExactControl {
	a.Set("quickEdit", value)
	return a
}

func (a *StaticExactControl) ReadOnly(value interface{}) *StaticExactControl {
	a.Set("readOnly", value)
	return a
}

func (a *StaticExactControl) ReadOnlyOn(value interface{}) *StaticExactControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *StaticExactControl) Remark(value interface{}) *StaticExactControl {
	a.Set("remark", value)
	return a
}

func (a *StaticExactControl) Required(value interface{}) *StaticExactControl {
	a.Set("required", value)
	return a
}

func (a *StaticExactControl) Row(value interface{}) *StaticExactControl {
	a.Set("row", value)
	return a
}

func (a *StaticExactControl) SaveImmediately(value interface{}) *StaticExactControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *StaticExactControl) Size(value interface{}) *StaticExactControl {
	a.Set("size", value)
	return a
}

func (a *StaticExactControl) Static(value interface{}) *StaticExactControl {
	a.Set("static", value)
	return a
}

func (a *StaticExactControl) StaticClassName(value interface{}) *StaticExactControl {
	a.Set("staticClassName", value)
	return a
}

func (a *StaticExactControl) StaticInputClassName(value interface{}) *StaticExactControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *StaticExactControl) StaticLabelClassName(value interface{}) *StaticExactControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *StaticExactControl) StaticOn(value interface{}) *StaticExactControl {
	a.Set("staticOn", value)
	return a
}

func (a *StaticExactControl) StaticPlaceholder(value interface{}) *StaticExactControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *StaticExactControl) StaticSchema(value interface{}) *StaticExactControl {
	a.Set("staticSchema", value)
	return a
}

func (a *StaticExactControl) Style(value interface{}) *StaticExactControl {
	a.Set("style", value)
	return a
}

func (a *StaticExactControl) SubmitOnChange(value interface{}) *StaticExactControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *StaticExactControl) TestIdBuilder(value interface{}) *StaticExactControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *StaticExactControl) Text(value interface{}) *StaticExactControl {
	a.Set("text", value)
	return a
}

func (a *StaticExactControl) Tpl(value interface{}) *StaticExactControl {
	a.Set("tpl", value)
	return a
}

func (a *StaticExactControl) Type(value interface{}) *StaticExactControl {
	a.Set("type", value)
	return a
}

func (a *StaticExactControl) UseMobileUI(value interface{}) *StaticExactControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *StaticExactControl) ValidateApi(value interface{}) *StaticExactControl {
	a.Set("validateApi", value)
	return a
}

func (a *StaticExactControl) ValidateOnChange(value interface{}) *StaticExactControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *StaticExactControl) ValidationErrors(value interface{}) *StaticExactControl {
	a.Set("validationErrors", value)
	return a
}

func (a *StaticExactControl) Validations(value interface{}) *StaticExactControl {
	a.Set("validations", value)
	return a
}

func (a *StaticExactControl) Value(value interface{}) *StaticExactControl {
	a.Set("value", value)
	return a
}

func (a *StaticExactControl) Visible(value interface{}) *StaticExactControl {
	a.Set("visible", value)
	return a
}

func (a *StaticExactControl) VisibleOn(value interface{}) *StaticExactControl {
	a.Set("visibleOn", value)
	return a
}

func (a *StaticExactControl) Width(value interface{}) *StaticExactControl {
	a.Set("width", value)
	return a
}
