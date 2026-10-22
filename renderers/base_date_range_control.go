package renderers

type BaseDateRangeControl struct {
	*BaseRenderer
}

func NewBaseDateRangeControl() *BaseDateRangeControl {
	a := &BaseDateRangeControl{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *BaseDateRangeControl) Set(name string, value interface{}) *BaseDateRangeControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *BaseDateRangeControl) Animation(value interface{}) *BaseDateRangeControl {
	a.Set("animation", value)
	return a
}

func (a *BaseDateRangeControl) AutoFill(value interface{}) *BaseDateRangeControl {
	a.Set("autoFill", value)
	return a
}

func (a *BaseDateRangeControl) BorderMode(value interface{}) *BaseDateRangeControl {
	a.Set("borderMode", value)
	return a
}

func (a *BaseDateRangeControl) ClassName(value interface{}) *BaseDateRangeControl {
	a.Set("className", value)
	return a
}

func (a *BaseDateRangeControl) ClearValueOnHidden(value interface{}) *BaseDateRangeControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *BaseDateRangeControl) Delimiter(value interface{}) *BaseDateRangeControl {
	a.Set("delimiter", value)
	return a
}

func (a *BaseDateRangeControl) Desc(value interface{}) *BaseDateRangeControl {
	a.Set("desc", value)
	return a
}

func (a *BaseDateRangeControl) Description(value interface{}) *BaseDateRangeControl {
	a.Set("description", value)
	return a
}

func (a *BaseDateRangeControl) DescriptionClassName(value interface{}) *BaseDateRangeControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *BaseDateRangeControl) Disabled(value interface{}) *BaseDateRangeControl {
	a.Set("disabled", value)
	return a
}

func (a *BaseDateRangeControl) DisabledOn(value interface{}) *BaseDateRangeControl {
	a.Set("disabledOn", value)
	return a
}

func (a *BaseDateRangeControl) DisplayFormat(value interface{}) *BaseDateRangeControl {
	a.Set("displayFormat", value)
	return a
}

func (a *BaseDateRangeControl) EditorSetting(value interface{}) *BaseDateRangeControl {
	a.Set("editorSetting", value)
	return a
}

func (a *BaseDateRangeControl) Embed(value interface{}) *BaseDateRangeControl {
	a.Set("embed", value)
	return a
}

func (a *BaseDateRangeControl) EndPlaceholder(value interface{}) *BaseDateRangeControl {
	a.Set("endPlaceholder", value)
	return a
}

func (a *BaseDateRangeControl) ExtraName(value interface{}) *BaseDateRangeControl {
	a.Set("extraName", value)
	return a
}

func (a *BaseDateRangeControl) Format(value interface{}) *BaseDateRangeControl {
	a.Set("format", value)
	return a
}

func (a *BaseDateRangeControl) Hidden(value interface{}) *BaseDateRangeControl {
	a.Set("hidden", value)
	return a
}

func (a *BaseDateRangeControl) HiddenOn(value interface{}) *BaseDateRangeControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *BaseDateRangeControl) Hint(value interface{}) *BaseDateRangeControl {
	a.Set("hint", value)
	return a
}

func (a *BaseDateRangeControl) Horizontal(value interface{}) *BaseDateRangeControl {
	a.Set("horizontal", value)
	return a
}

func (a *BaseDateRangeControl) Id(value interface{}) *BaseDateRangeControl {
	a.Set("id", value)
	return a
}

func (a *BaseDateRangeControl) InitAutoFill(value interface{}) *BaseDateRangeControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *BaseDateRangeControl) Inline(value interface{}) *BaseDateRangeControl {
	a.Set("inline", value)
	return a
}

func (a *BaseDateRangeControl) InputClassName(value interface{}) *BaseDateRangeControl {
	a.Set("inputClassName", value)
	return a
}

func (a *BaseDateRangeControl) InputFormat(value interface{}) *BaseDateRangeControl {
	a.Set("inputFormat", value)
	return a
}

func (a *BaseDateRangeControl) JoinValues(value interface{}) *BaseDateRangeControl {
	a.Set("joinValues", value)
	return a
}

func (a *BaseDateRangeControl) Label(value interface{}) *BaseDateRangeControl {
	a.Set("label", value)
	return a
}

func (a *BaseDateRangeControl) LabelAlign(value interface{}) *BaseDateRangeControl {
	a.Set("labelAlign", value)
	return a
}

func (a *BaseDateRangeControl) LabelClassName(value interface{}) *BaseDateRangeControl {
	a.Set("labelClassName", value)
	return a
}

func (a *BaseDateRangeControl) LabelOverflow(value interface{}) *BaseDateRangeControl {
	a.Set("labelOverflow", value)
	return a
}

func (a *BaseDateRangeControl) LabelRemark(value interface{}) *BaseDateRangeControl {
	a.Set("labelRemark", value)
	return a
}

func (a *BaseDateRangeControl) LabelWidth(value interface{}) *BaseDateRangeControl {
	a.Set("labelWidth", value)
	return a
}

func (a *BaseDateRangeControl) MaxDate(value interface{}) *BaseDateRangeControl {
	a.Set("maxDate", value)
	return a
}

func (a *BaseDateRangeControl) MaxDuration(value interface{}) *BaseDateRangeControl {
	a.Set("maxDuration", value)
	return a
}

func (a *BaseDateRangeControl) MinDate(value interface{}) *BaseDateRangeControl {
	a.Set("minDate", value)
	return a
}

func (a *BaseDateRangeControl) MinDuration(value interface{}) *BaseDateRangeControl {
	a.Set("minDuration", value)
	return a
}

func (a *BaseDateRangeControl) Mode(value interface{}) *BaseDateRangeControl {
	a.Set("mode", value)
	return a
}

func (a *BaseDateRangeControl) Name(value interface{}) *BaseDateRangeControl {
	a.Set("name", value)
	return a
}

func (a *BaseDateRangeControl) OnEvent(value interface{}) *BaseDateRangeControl {
	a.Set("onEvent", value)
	return a
}

func (a *BaseDateRangeControl) Placeholder(value interface{}) *BaseDateRangeControl {
	a.Set("placeholder", value)
	return a
}

func (a *BaseDateRangeControl) PopOverContainerSelector(value interface{}) *BaseDateRangeControl {
	a.Set("popOverContainerSelector", value)
	return a
}

func (a *BaseDateRangeControl) Ranges(value interface{}) *BaseDateRangeControl {
	a.Set("ranges", value)
	return a
}

func (a *BaseDateRangeControl) ReadOnly(value interface{}) *BaseDateRangeControl {
	a.Set("readOnly", value)
	return a
}

func (a *BaseDateRangeControl) ReadOnlyOn(value interface{}) *BaseDateRangeControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *BaseDateRangeControl) Remark(value interface{}) *BaseDateRangeControl {
	a.Set("remark", value)
	return a
}

func (a *BaseDateRangeControl) Required(value interface{}) *BaseDateRangeControl {
	a.Set("required", value)
	return a
}

func (a *BaseDateRangeControl) Row(value interface{}) *BaseDateRangeControl {
	a.Set("row", value)
	return a
}

func (a *BaseDateRangeControl) SaveImmediately(value interface{}) *BaseDateRangeControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *BaseDateRangeControl) Shortcuts(value interface{}) *BaseDateRangeControl {
	a.Set("shortcuts", value)
	return a
}

func (a *BaseDateRangeControl) Size(value interface{}) *BaseDateRangeControl {
	a.Set("size", value)
	return a
}

func (a *BaseDateRangeControl) StartPlaceholder(value interface{}) *BaseDateRangeControl {
	a.Set("startPlaceholder", value)
	return a
}

func (a *BaseDateRangeControl) Static(value interface{}) *BaseDateRangeControl {
	a.Set("static", value)
	return a
}

func (a *BaseDateRangeControl) StaticClassName(value interface{}) *BaseDateRangeControl {
	a.Set("staticClassName", value)
	return a
}

func (a *BaseDateRangeControl) StaticInputClassName(value interface{}) *BaseDateRangeControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *BaseDateRangeControl) StaticLabelClassName(value interface{}) *BaseDateRangeControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *BaseDateRangeControl) StaticOn(value interface{}) *BaseDateRangeControl {
	a.Set("staticOn", value)
	return a
}

func (a *BaseDateRangeControl) StaticPlaceholder(value interface{}) *BaseDateRangeControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *BaseDateRangeControl) StaticSchema(value interface{}) *BaseDateRangeControl {
	a.Set("staticSchema", value)
	return a
}

func (a *BaseDateRangeControl) Style(value interface{}) *BaseDateRangeControl {
	a.Set("style", value)
	return a
}

func (a *BaseDateRangeControl) SubmitOnChange(value interface{}) *BaseDateRangeControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *BaseDateRangeControl) Transform(value interface{}) *BaseDateRangeControl {
	a.Set("transform", value)
	return a
}

func (a *BaseDateRangeControl) Type(value interface{}) *BaseDateRangeControl {
	a.Set("type", value)
	return a
}

func (a *BaseDateRangeControl) UseMobileUI(value interface{}) *BaseDateRangeControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *BaseDateRangeControl) ValidateApi(value interface{}) *BaseDateRangeControl {
	a.Set("validateApi", value)
	return a
}

func (a *BaseDateRangeControl) ValidateOnChange(value interface{}) *BaseDateRangeControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *BaseDateRangeControl) ValidationErrors(value interface{}) *BaseDateRangeControl {
	a.Set("validationErrors", value)
	return a
}

func (a *BaseDateRangeControl) Validations(value interface{}) *BaseDateRangeControl {
	a.Set("validations", value)
	return a
}

func (a *BaseDateRangeControl) Value(value interface{}) *BaseDateRangeControl {
	a.Set("value", value)
	return a
}

func (a *BaseDateRangeControl) ValueFormat(value interface{}) *BaseDateRangeControl {
	a.Set("valueFormat", value)
	return a
}

func (a *BaseDateRangeControl) Visible(value interface{}) *BaseDateRangeControl {
	a.Set("visible", value)
	return a
}

func (a *BaseDateRangeControl) VisibleOn(value interface{}) *BaseDateRangeControl {
	a.Set("visibleOn", value)
	return a
}

func (a *BaseDateRangeControl) Width(value interface{}) *BaseDateRangeControl {
	a.Set("width", value)
	return a
}
