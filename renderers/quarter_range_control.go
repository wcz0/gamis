package renderers

type QuarterRangeControl struct {
	*BaseRenderer
}

func NewQuarterRangeControl() *QuarterRangeControl {
	a := &QuarterRangeControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "input-quarter-range")
	return a
}

func (a *QuarterRangeControl) Set(name string, value interface{}) *QuarterRangeControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *QuarterRangeControl) Animation(value interface{}) *QuarterRangeControl {
	a.Set("animation", value)
	return a
}

func (a *QuarterRangeControl) AutoFill(value interface{}) *QuarterRangeControl {
	a.Set("autoFill", value)
	return a
}

func (a *QuarterRangeControl) BorderMode(value interface{}) *QuarterRangeControl {
	a.Set("borderMode", value)
	return a
}

func (a *QuarterRangeControl) ClassName(value interface{}) *QuarterRangeControl {
	a.Set("className", value)
	return a
}

func (a *QuarterRangeControl) ClearValueOnHidden(value interface{}) *QuarterRangeControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *QuarterRangeControl) Delimiter(value interface{}) *QuarterRangeControl {
	a.Set("delimiter", value)
	return a
}

func (a *QuarterRangeControl) Desc(value interface{}) *QuarterRangeControl {
	a.Set("desc", value)
	return a
}

func (a *QuarterRangeControl) Description(value interface{}) *QuarterRangeControl {
	a.Set("description", value)
	return a
}

func (a *QuarterRangeControl) DescriptionClassName(value interface{}) *QuarterRangeControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *QuarterRangeControl) Disabled(value interface{}) *QuarterRangeControl {
	a.Set("disabled", value)
	return a
}

func (a *QuarterRangeControl) DisabledOn(value interface{}) *QuarterRangeControl {
	a.Set("disabledOn", value)
	return a
}

func (a *QuarterRangeControl) DisplayFormat(value interface{}) *QuarterRangeControl {
	a.Set("displayFormat", value)
	return a
}

func (a *QuarterRangeControl) EditorSetting(value interface{}) *QuarterRangeControl {
	a.Set("editorSetting", value)
	return a
}

func (a *QuarterRangeControl) Embed(value interface{}) *QuarterRangeControl {
	a.Set("embed", value)
	return a
}

func (a *QuarterRangeControl) EndPlaceholder(value interface{}) *QuarterRangeControl {
	a.Set("endPlaceholder", value)
	return a
}

func (a *QuarterRangeControl) ExtraName(value interface{}) *QuarterRangeControl {
	a.Set("extraName", value)
	return a
}

func (a *QuarterRangeControl) Format(value interface{}) *QuarterRangeControl {
	a.Set("format", value)
	return a
}

func (a *QuarterRangeControl) Hidden(value interface{}) *QuarterRangeControl {
	a.Set("hidden", value)
	return a
}

func (a *QuarterRangeControl) HiddenOn(value interface{}) *QuarterRangeControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *QuarterRangeControl) Hint(value interface{}) *QuarterRangeControl {
	a.Set("hint", value)
	return a
}

func (a *QuarterRangeControl) Horizontal(value interface{}) *QuarterRangeControl {
	a.Set("horizontal", value)
	return a
}

func (a *QuarterRangeControl) Id(value interface{}) *QuarterRangeControl {
	a.Set("id", value)
	return a
}

func (a *QuarterRangeControl) InitAutoFill(value interface{}) *QuarterRangeControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *QuarterRangeControl) Inline(value interface{}) *QuarterRangeControl {
	a.Set("inline", value)
	return a
}

func (a *QuarterRangeControl) InputClassName(value interface{}) *QuarterRangeControl {
	a.Set("inputClassName", value)
	return a
}

func (a *QuarterRangeControl) InputFormat(value interface{}) *QuarterRangeControl {
	a.Set("inputFormat", value)
	return a
}

func (a *QuarterRangeControl) JoinValues(value interface{}) *QuarterRangeControl {
	a.Set("joinValues", value)
	return a
}

func (a *QuarterRangeControl) Label(value interface{}) *QuarterRangeControl {
	a.Set("label", value)
	return a
}

func (a *QuarterRangeControl) LabelAlign(value interface{}) *QuarterRangeControl {
	a.Set("labelAlign", value)
	return a
}

func (a *QuarterRangeControl) LabelClassName(value interface{}) *QuarterRangeControl {
	a.Set("labelClassName", value)
	return a
}

func (a *QuarterRangeControl) LabelRemark(value interface{}) *QuarterRangeControl {
	a.Set("labelRemark", value)
	return a
}

func (a *QuarterRangeControl) LabelWidth(value interface{}) *QuarterRangeControl {
	a.Set("labelWidth", value)
	return a
}

func (a *QuarterRangeControl) MaxDate(value interface{}) *QuarterRangeControl {
	a.Set("maxDate", value)
	return a
}

func (a *QuarterRangeControl) MaxDuration(value interface{}) *QuarterRangeControl {
	a.Set("maxDuration", value)
	return a
}

func (a *QuarterRangeControl) MinDate(value interface{}) *QuarterRangeControl {
	a.Set("minDate", value)
	return a
}

func (a *QuarterRangeControl) MinDuration(value interface{}) *QuarterRangeControl {
	a.Set("minDuration", value)
	return a
}

func (a *QuarterRangeControl) Mode(value interface{}) *QuarterRangeControl {
	a.Set("mode", value)
	return a
}

func (a *QuarterRangeControl) Name(value interface{}) *QuarterRangeControl {
	a.Set("name", value)
	return a
}

func (a *QuarterRangeControl) OnEvent(value interface{}) *QuarterRangeControl {
	a.Set("onEvent", value)
	return a
}

func (a *QuarterRangeControl) Placeholder(value interface{}) *QuarterRangeControl {
	a.Set("placeholder", value)
	return a
}

func (a *QuarterRangeControl) PopOverContainerSelector(value interface{}) *QuarterRangeControl {
	a.Set("popOverContainerSelector", value)
	return a
}

func (a *QuarterRangeControl) Ranges(value interface{}) *QuarterRangeControl {
	a.Set("ranges", value)
	return a
}

func (a *QuarterRangeControl) ReadOnly(value interface{}) *QuarterRangeControl {
	a.Set("readOnly", value)
	return a
}

func (a *QuarterRangeControl) ReadOnlyOn(value interface{}) *QuarterRangeControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *QuarterRangeControl) Remark(value interface{}) *QuarterRangeControl {
	a.Set("remark", value)
	return a
}

func (a *QuarterRangeControl) Required(value interface{}) *QuarterRangeControl {
	a.Set("required", value)
	return a
}

func (a *QuarterRangeControl) Row(value interface{}) *QuarterRangeControl {
	a.Set("row", value)
	return a
}

func (a *QuarterRangeControl) SaveImmediately(value interface{}) *QuarterRangeControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *QuarterRangeControl) Shortcuts(value interface{}) *QuarterRangeControl {
	a.Set("shortcuts", value)
	return a
}

func (a *QuarterRangeControl) Size(value interface{}) *QuarterRangeControl {
	a.Set("size", value)
	return a
}

func (a *QuarterRangeControl) StartPlaceholder(value interface{}) *QuarterRangeControl {
	a.Set("startPlaceholder", value)
	return a
}

func (a *QuarterRangeControl) Static(value interface{}) *QuarterRangeControl {
	a.Set("static", value)
	return a
}

func (a *QuarterRangeControl) StaticClassName(value interface{}) *QuarterRangeControl {
	a.Set("staticClassName", value)
	return a
}

func (a *QuarterRangeControl) StaticInputClassName(value interface{}) *QuarterRangeControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *QuarterRangeControl) StaticLabelClassName(value interface{}) *QuarterRangeControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *QuarterRangeControl) StaticOn(value interface{}) *QuarterRangeControl {
	a.Set("staticOn", value)
	return a
}

func (a *QuarterRangeControl) StaticPlaceholder(value interface{}) *QuarterRangeControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *QuarterRangeControl) StaticSchema(value interface{}) *QuarterRangeControl {
	a.Set("staticSchema", value)
	return a
}

func (a *QuarterRangeControl) Style(value interface{}) *QuarterRangeControl {
	a.Set("style", value)
	return a
}

func (a *QuarterRangeControl) SubmitOnChange(value interface{}) *QuarterRangeControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *QuarterRangeControl) TestIdBuilder(value interface{}) *QuarterRangeControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *QuarterRangeControl) Transform(value interface{}) *QuarterRangeControl {
	a.Set("transform", value)
	return a
}

func (a *QuarterRangeControl) Type(value interface{}) *QuarterRangeControl {
	a.Set("type", value)
	return a
}

func (a *QuarterRangeControl) UseMobileUI(value interface{}) *QuarterRangeControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *QuarterRangeControl) ValidateApi(value interface{}) *QuarterRangeControl {
	a.Set("validateApi", value)
	return a
}

func (a *QuarterRangeControl) ValidateOnChange(value interface{}) *QuarterRangeControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *QuarterRangeControl) ValidationErrors(value interface{}) *QuarterRangeControl {
	a.Set("validationErrors", value)
	return a
}

func (a *QuarterRangeControl) Validations(value interface{}) *QuarterRangeControl {
	a.Set("validations", value)
	return a
}

func (a *QuarterRangeControl) Value(value interface{}) *QuarterRangeControl {
	a.Set("value", value)
	return a
}

func (a *QuarterRangeControl) ValueFormat(value interface{}) *QuarterRangeControl {
	a.Set("valueFormat", value)
	return a
}

func (a *QuarterRangeControl) Visible(value interface{}) *QuarterRangeControl {
	a.Set("visible", value)
	return a
}

func (a *QuarterRangeControl) VisibleOn(value interface{}) *QuarterRangeControl {
	a.Set("visibleOn", value)
	return a
}

func (a *QuarterRangeControl) Width(value interface{}) *QuarterRangeControl {
	a.Set("width", value)
	return a
}
