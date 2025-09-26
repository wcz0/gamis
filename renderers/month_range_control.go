package renderers

type MonthRangeControl struct {
	*BaseRenderer
}

func NewMonthRangeControl() *MonthRangeControl {
	a := &MonthRangeControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "input-month-range")
	return a
}

func (a *MonthRangeControl) Set(name string, value interface{}) *MonthRangeControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *MonthRangeControl) Animation(value interface{}) *MonthRangeControl {
	a.Set("animation", value)
	return a
}

func (a *MonthRangeControl) AutoFill(value interface{}) *MonthRangeControl {
	a.Set("autoFill", value)
	return a
}

func (a *MonthRangeControl) BorderMode(value interface{}) *MonthRangeControl {
	a.Set("borderMode", value)
	return a
}

func (a *MonthRangeControl) ClassName(value interface{}) *MonthRangeControl {
	a.Set("className", value)
	return a
}

func (a *MonthRangeControl) ClearValueOnHidden(value interface{}) *MonthRangeControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *MonthRangeControl) Delimiter(value interface{}) *MonthRangeControl {
	a.Set("delimiter", value)
	return a
}

func (a *MonthRangeControl) Desc(value interface{}) *MonthRangeControl {
	a.Set("desc", value)
	return a
}

func (a *MonthRangeControl) Description(value interface{}) *MonthRangeControl {
	a.Set("description", value)
	return a
}

func (a *MonthRangeControl) DescriptionClassName(value interface{}) *MonthRangeControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *MonthRangeControl) Disabled(value interface{}) *MonthRangeControl {
	a.Set("disabled", value)
	return a
}

func (a *MonthRangeControl) DisabledOn(value interface{}) *MonthRangeControl {
	a.Set("disabledOn", value)
	return a
}

func (a *MonthRangeControl) DisplayFormat(value interface{}) *MonthRangeControl {
	a.Set("displayFormat", value)
	return a
}

func (a *MonthRangeControl) EditorSetting(value interface{}) *MonthRangeControl {
	a.Set("editorSetting", value)
	return a
}

func (a *MonthRangeControl) Embed(value interface{}) *MonthRangeControl {
	a.Set("embed", value)
	return a
}

func (a *MonthRangeControl) EndPlaceholder(value interface{}) *MonthRangeControl {
	a.Set("endPlaceholder", value)
	return a
}

func (a *MonthRangeControl) ExtraName(value interface{}) *MonthRangeControl {
	a.Set("extraName", value)
	return a
}

func (a *MonthRangeControl) Format(value interface{}) *MonthRangeControl {
	a.Set("format", value)
	return a
}

func (a *MonthRangeControl) Hidden(value interface{}) *MonthRangeControl {
	a.Set("hidden", value)
	return a
}

func (a *MonthRangeControl) HiddenOn(value interface{}) *MonthRangeControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *MonthRangeControl) Hint(value interface{}) *MonthRangeControl {
	a.Set("hint", value)
	return a
}

func (a *MonthRangeControl) Horizontal(value interface{}) *MonthRangeControl {
	a.Set("horizontal", value)
	return a
}

func (a *MonthRangeControl) Id(value interface{}) *MonthRangeControl {
	a.Set("id", value)
	return a
}

func (a *MonthRangeControl) InitAutoFill(value interface{}) *MonthRangeControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *MonthRangeControl) Inline(value interface{}) *MonthRangeControl {
	a.Set("inline", value)
	return a
}

func (a *MonthRangeControl) InputClassName(value interface{}) *MonthRangeControl {
	a.Set("inputClassName", value)
	return a
}

func (a *MonthRangeControl) InputFormat(value interface{}) *MonthRangeControl {
	a.Set("inputFormat", value)
	return a
}

func (a *MonthRangeControl) JoinValues(value interface{}) *MonthRangeControl {
	a.Set("joinValues", value)
	return a
}

func (a *MonthRangeControl) Label(value interface{}) *MonthRangeControl {
	a.Set("label", value)
	return a
}

func (a *MonthRangeControl) LabelAlign(value interface{}) *MonthRangeControl {
	a.Set("labelAlign", value)
	return a
}

func (a *MonthRangeControl) LabelClassName(value interface{}) *MonthRangeControl {
	a.Set("labelClassName", value)
	return a
}

func (a *MonthRangeControl) LabelRemark(value interface{}) *MonthRangeControl {
	a.Set("labelRemark", value)
	return a
}

func (a *MonthRangeControl) LabelWidth(value interface{}) *MonthRangeControl {
	a.Set("labelWidth", value)
	return a
}

func (a *MonthRangeControl) MaxDate(value interface{}) *MonthRangeControl {
	a.Set("maxDate", value)
	return a
}

func (a *MonthRangeControl) MaxDuration(value interface{}) *MonthRangeControl {
	a.Set("maxDuration", value)
	return a
}

func (a *MonthRangeControl) MinDate(value interface{}) *MonthRangeControl {
	a.Set("minDate", value)
	return a
}

func (a *MonthRangeControl) MinDuration(value interface{}) *MonthRangeControl {
	a.Set("minDuration", value)
	return a
}

func (a *MonthRangeControl) Mode(value interface{}) *MonthRangeControl {
	a.Set("mode", value)
	return a
}

func (a *MonthRangeControl) Name(value interface{}) *MonthRangeControl {
	a.Set("name", value)
	return a
}

func (a *MonthRangeControl) OnEvent(value interface{}) *MonthRangeControl {
	a.Set("onEvent", value)
	return a
}

func (a *MonthRangeControl) Placeholder(value interface{}) *MonthRangeControl {
	a.Set("placeholder", value)
	return a
}

func (a *MonthRangeControl) PopOverContainerSelector(value interface{}) *MonthRangeControl {
	a.Set("popOverContainerSelector", value)
	return a
}

func (a *MonthRangeControl) Ranges(value interface{}) *MonthRangeControl {
	a.Set("ranges", value)
	return a
}

func (a *MonthRangeControl) ReadOnly(value interface{}) *MonthRangeControl {
	a.Set("readOnly", value)
	return a
}

func (a *MonthRangeControl) ReadOnlyOn(value interface{}) *MonthRangeControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *MonthRangeControl) Remark(value interface{}) *MonthRangeControl {
	a.Set("remark", value)
	return a
}

func (a *MonthRangeControl) Required(value interface{}) *MonthRangeControl {
	a.Set("required", value)
	return a
}

func (a *MonthRangeControl) Row(value interface{}) *MonthRangeControl {
	a.Set("row", value)
	return a
}

func (a *MonthRangeControl) SaveImmediately(value interface{}) *MonthRangeControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *MonthRangeControl) Shortcuts(value interface{}) *MonthRangeControl {
	a.Set("shortcuts", value)
	return a
}

func (a *MonthRangeControl) Size(value interface{}) *MonthRangeControl {
	a.Set("size", value)
	return a
}

func (a *MonthRangeControl) StartPlaceholder(value interface{}) *MonthRangeControl {
	a.Set("startPlaceholder", value)
	return a
}

func (a *MonthRangeControl) Static(value interface{}) *MonthRangeControl {
	a.Set("static", value)
	return a
}

func (a *MonthRangeControl) StaticClassName(value interface{}) *MonthRangeControl {
	a.Set("staticClassName", value)
	return a
}

func (a *MonthRangeControl) StaticInputClassName(value interface{}) *MonthRangeControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *MonthRangeControl) StaticLabelClassName(value interface{}) *MonthRangeControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *MonthRangeControl) StaticOn(value interface{}) *MonthRangeControl {
	a.Set("staticOn", value)
	return a
}

func (a *MonthRangeControl) StaticPlaceholder(value interface{}) *MonthRangeControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *MonthRangeControl) StaticSchema(value interface{}) *MonthRangeControl {
	a.Set("staticSchema", value)
	return a
}

func (a *MonthRangeControl) Style(value interface{}) *MonthRangeControl {
	a.Set("style", value)
	return a
}

func (a *MonthRangeControl) SubmitOnChange(value interface{}) *MonthRangeControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *MonthRangeControl) TestIdBuilder(value interface{}) *MonthRangeControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *MonthRangeControl) Transform(value interface{}) *MonthRangeControl {
	a.Set("transform", value)
	return a
}

func (a *MonthRangeControl) Type(value interface{}) *MonthRangeControl {
	a.Set("type", value)
	return a
}

func (a *MonthRangeControl) UseMobileUI(value interface{}) *MonthRangeControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *MonthRangeControl) ValidateApi(value interface{}) *MonthRangeControl {
	a.Set("validateApi", value)
	return a
}

func (a *MonthRangeControl) ValidateOnChange(value interface{}) *MonthRangeControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *MonthRangeControl) ValidationErrors(value interface{}) *MonthRangeControl {
	a.Set("validationErrors", value)
	return a
}

func (a *MonthRangeControl) Validations(value interface{}) *MonthRangeControl {
	a.Set("validations", value)
	return a
}

func (a *MonthRangeControl) Value(value interface{}) *MonthRangeControl {
	a.Set("value", value)
	return a
}

func (a *MonthRangeControl) ValueFormat(value interface{}) *MonthRangeControl {
	a.Set("valueFormat", value)
	return a
}

func (a *MonthRangeControl) Visible(value interface{}) *MonthRangeControl {
	a.Set("visible", value)
	return a
}

func (a *MonthRangeControl) VisibleOn(value interface{}) *MonthRangeControl {
	a.Set("visibleOn", value)
	return a
}

func (a *MonthRangeControl) Width(value interface{}) *MonthRangeControl {
	a.Set("width", value)
	return a
}
