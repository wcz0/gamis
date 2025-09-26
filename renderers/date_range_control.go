package renderers

type DateRangeControl struct {
	*BaseRenderer
}

func NewDateRangeControl() *DateRangeControl {
	a := &DateRangeControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "input-date-range")
	return a
}

func (a *DateRangeControl) Set(name string, value interface{}) *DateRangeControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *DateRangeControl) Animation(value interface{}) *DateRangeControl {
	a.Set("animation", value)
	return a
}

func (a *DateRangeControl) AutoFill(value interface{}) *DateRangeControl {
	a.Set("autoFill", value)
	return a
}

func (a *DateRangeControl) BorderMode(value interface{}) *DateRangeControl {
	a.Set("borderMode", value)
	return a
}

func (a *DateRangeControl) ClassName(value interface{}) *DateRangeControl {
	a.Set("className", value)
	return a
}

func (a *DateRangeControl) ClearValueOnHidden(value interface{}) *DateRangeControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *DateRangeControl) Delimiter(value interface{}) *DateRangeControl {
	a.Set("delimiter", value)
	return a
}

func (a *DateRangeControl) Desc(value interface{}) *DateRangeControl {
	a.Set("desc", value)
	return a
}

func (a *DateRangeControl) Description(value interface{}) *DateRangeControl {
	a.Set("description", value)
	return a
}

func (a *DateRangeControl) DescriptionClassName(value interface{}) *DateRangeControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *DateRangeControl) Disabled(value interface{}) *DateRangeControl {
	a.Set("disabled", value)
	return a
}

func (a *DateRangeControl) DisabledOn(value interface{}) *DateRangeControl {
	a.Set("disabledOn", value)
	return a
}

func (a *DateRangeControl) DisplayFormat(value interface{}) *DateRangeControl {
	a.Set("displayFormat", value)
	return a
}

func (a *DateRangeControl) EditorSetting(value interface{}) *DateRangeControl {
	a.Set("editorSetting", value)
	return a
}

func (a *DateRangeControl) Embed(value interface{}) *DateRangeControl {
	a.Set("embed", value)
	return a
}

func (a *DateRangeControl) EndPlaceholder(value interface{}) *DateRangeControl {
	a.Set("endPlaceholder", value)
	return a
}

func (a *DateRangeControl) ExtraName(value interface{}) *DateRangeControl {
	a.Set("extraName", value)
	return a
}

func (a *DateRangeControl) Format(value interface{}) *DateRangeControl {
	a.Set("format", value)
	return a
}

func (a *DateRangeControl) Hidden(value interface{}) *DateRangeControl {
	a.Set("hidden", value)
	return a
}

func (a *DateRangeControl) HiddenOn(value interface{}) *DateRangeControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *DateRangeControl) Hint(value interface{}) *DateRangeControl {
	a.Set("hint", value)
	return a
}

func (a *DateRangeControl) Horizontal(value interface{}) *DateRangeControl {
	a.Set("horizontal", value)
	return a
}

func (a *DateRangeControl) Id(value interface{}) *DateRangeControl {
	a.Set("id", value)
	return a
}

func (a *DateRangeControl) InitAutoFill(value interface{}) *DateRangeControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *DateRangeControl) Inline(value interface{}) *DateRangeControl {
	a.Set("inline", value)
	return a
}

func (a *DateRangeControl) InputClassName(value interface{}) *DateRangeControl {
	a.Set("inputClassName", value)
	return a
}

func (a *DateRangeControl) InputFormat(value interface{}) *DateRangeControl {
	a.Set("inputFormat", value)
	return a
}

func (a *DateRangeControl) JoinValues(value interface{}) *DateRangeControl {
	a.Set("joinValues", value)
	return a
}

func (a *DateRangeControl) Label(value interface{}) *DateRangeControl {
	a.Set("label", value)
	return a
}

func (a *DateRangeControl) LabelAlign(value interface{}) *DateRangeControl {
	a.Set("labelAlign", value)
	return a
}

func (a *DateRangeControl) LabelClassName(value interface{}) *DateRangeControl {
	a.Set("labelClassName", value)
	return a
}

func (a *DateRangeControl) LabelRemark(value interface{}) *DateRangeControl {
	a.Set("labelRemark", value)
	return a
}

func (a *DateRangeControl) LabelWidth(value interface{}) *DateRangeControl {
	a.Set("labelWidth", value)
	return a
}

func (a *DateRangeControl) MaxDate(value interface{}) *DateRangeControl {
	a.Set("maxDate", value)
	return a
}

func (a *DateRangeControl) MaxDuration(value interface{}) *DateRangeControl {
	a.Set("maxDuration", value)
	return a
}

func (a *DateRangeControl) MinDate(value interface{}) *DateRangeControl {
	a.Set("minDate", value)
	return a
}

func (a *DateRangeControl) MinDuration(value interface{}) *DateRangeControl {
	a.Set("minDuration", value)
	return a
}

func (a *DateRangeControl) Mode(value interface{}) *DateRangeControl {
	a.Set("mode", value)
	return a
}

func (a *DateRangeControl) Name(value interface{}) *DateRangeControl {
	a.Set("name", value)
	return a
}

func (a *DateRangeControl) OnEvent(value interface{}) *DateRangeControl {
	a.Set("onEvent", value)
	return a
}

func (a *DateRangeControl) Placeholder(value interface{}) *DateRangeControl {
	a.Set("placeholder", value)
	return a
}

func (a *DateRangeControl) PopOverContainerSelector(value interface{}) *DateRangeControl {
	a.Set("popOverContainerSelector", value)
	return a
}

func (a *DateRangeControl) Ranges(value interface{}) *DateRangeControl {
	a.Set("ranges", value)
	return a
}

func (a *DateRangeControl) ReadOnly(value interface{}) *DateRangeControl {
	a.Set("readOnly", value)
	return a
}

func (a *DateRangeControl) ReadOnlyOn(value interface{}) *DateRangeControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *DateRangeControl) Remark(value interface{}) *DateRangeControl {
	a.Set("remark", value)
	return a
}

func (a *DateRangeControl) Required(value interface{}) *DateRangeControl {
	a.Set("required", value)
	return a
}

func (a *DateRangeControl) Row(value interface{}) *DateRangeControl {
	a.Set("row", value)
	return a
}

func (a *DateRangeControl) SaveImmediately(value interface{}) *DateRangeControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *DateRangeControl) Shortcuts(value interface{}) *DateRangeControl {
	a.Set("shortcuts", value)
	return a
}

func (a *DateRangeControl) Size(value interface{}) *DateRangeControl {
	a.Set("size", value)
	return a
}

func (a *DateRangeControl) StartPlaceholder(value interface{}) *DateRangeControl {
	a.Set("startPlaceholder", value)
	return a
}

func (a *DateRangeControl) Static(value interface{}) *DateRangeControl {
	a.Set("static", value)
	return a
}

func (a *DateRangeControl) StaticClassName(value interface{}) *DateRangeControl {
	a.Set("staticClassName", value)
	return a
}

func (a *DateRangeControl) StaticInputClassName(value interface{}) *DateRangeControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *DateRangeControl) StaticLabelClassName(value interface{}) *DateRangeControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *DateRangeControl) StaticOn(value interface{}) *DateRangeControl {
	a.Set("staticOn", value)
	return a
}

func (a *DateRangeControl) StaticPlaceholder(value interface{}) *DateRangeControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *DateRangeControl) StaticSchema(value interface{}) *DateRangeControl {
	a.Set("staticSchema", value)
	return a
}

func (a *DateRangeControl) Style(value interface{}) *DateRangeControl {
	a.Set("style", value)
	return a
}

func (a *DateRangeControl) SubmitOnChange(value interface{}) *DateRangeControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *DateRangeControl) TestIdBuilder(value interface{}) *DateRangeControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *DateRangeControl) Transform(value interface{}) *DateRangeControl {
	a.Set("transform", value)
	return a
}

func (a *DateRangeControl) Type(value interface{}) *DateRangeControl {
	a.Set("type", value)
	return a
}

func (a *DateRangeControl) UseMobileUI(value interface{}) *DateRangeControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *DateRangeControl) ValidateApi(value interface{}) *DateRangeControl {
	a.Set("validateApi", value)
	return a
}

func (a *DateRangeControl) ValidateOnChange(value interface{}) *DateRangeControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *DateRangeControl) ValidationErrors(value interface{}) *DateRangeControl {
	a.Set("validationErrors", value)
	return a
}

func (a *DateRangeControl) Validations(value interface{}) *DateRangeControl {
	a.Set("validations", value)
	return a
}

func (a *DateRangeControl) Value(value interface{}) *DateRangeControl {
	a.Set("value", value)
	return a
}

func (a *DateRangeControl) ValueFormat(value interface{}) *DateRangeControl {
	a.Set("valueFormat", value)
	return a
}

func (a *DateRangeControl) Visible(value interface{}) *DateRangeControl {
	a.Set("visible", value)
	return a
}

func (a *DateRangeControl) VisibleOn(value interface{}) *DateRangeControl {
	a.Set("visibleOn", value)
	return a
}

func (a *DateRangeControl) Width(value interface{}) *DateRangeControl {
	a.Set("width", value)
	return a
}
