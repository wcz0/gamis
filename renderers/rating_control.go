package renderers

type RatingControl struct {
	*BaseRenderer
}

func NewRatingControl() *RatingControl {
	a := &RatingControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "input-rating")
	return a
}

func (a *RatingControl) Set(name string, value interface{}) *RatingControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *RatingControl) AllowClear(value interface{}) *RatingControl {
	a.Set("allowClear", value)
	return a
}

func (a *RatingControl) AutoFill(value interface{}) *RatingControl {
	a.Set("autoFill", value)
	return a
}

func (a *RatingControl) Char(value interface{}) *RatingControl {
	a.Set("char", value)
	return a
}

func (a *RatingControl) CharClassName(value interface{}) *RatingControl {
	a.Set("charClassName", value)
	return a
}

func (a *RatingControl) ClassName(value interface{}) *RatingControl {
	a.Set("className", value)
	return a
}

func (a *RatingControl) ClearValueOnHidden(value interface{}) *RatingControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *RatingControl) Colors(value interface{}) *RatingControl {
	a.Set("colors", value)
	return a
}

func (a *RatingControl) Count(value interface{}) *RatingControl {
	a.Set("count", value)
	return a
}

func (a *RatingControl) Desc(value interface{}) *RatingControl {
	a.Set("desc", value)
	return a
}

func (a *RatingControl) Description(value interface{}) *RatingControl {
	a.Set("description", value)
	return a
}

func (a *RatingControl) DescriptionClassName(value interface{}) *RatingControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *RatingControl) Disabled(value interface{}) *RatingControl {
	a.Set("disabled", value)
	return a
}

func (a *RatingControl) DisabledOn(value interface{}) *RatingControl {
	a.Set("disabledOn", value)
	return a
}

func (a *RatingControl) EditorSetting(value interface{}) *RatingControl {
	a.Set("editorSetting", value)
	return a
}

func (a *RatingControl) ExtraName(value interface{}) *RatingControl {
	a.Set("extraName", value)
	return a
}

func (a *RatingControl) Half(value interface{}) *RatingControl {
	a.Set("half", value)
	return a
}

func (a *RatingControl) Hidden(value interface{}) *RatingControl {
	a.Set("hidden", value)
	return a
}

func (a *RatingControl) HiddenOn(value interface{}) *RatingControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *RatingControl) Hint(value interface{}) *RatingControl {
	a.Set("hint", value)
	return a
}

func (a *RatingControl) Horizontal(value interface{}) *RatingControl {
	a.Set("horizontal", value)
	return a
}

func (a *RatingControl) Id(value interface{}) *RatingControl {
	a.Set("id", value)
	return a
}

func (a *RatingControl) InactiveColor(value interface{}) *RatingControl {
	a.Set("inactiveColor", value)
	return a
}

func (a *RatingControl) InitAutoFill(value interface{}) *RatingControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *RatingControl) Inline(value interface{}) *RatingControl {
	a.Set("inline", value)
	return a
}

func (a *RatingControl) InputClassName(value interface{}) *RatingControl {
	a.Set("inputClassName", value)
	return a
}

func (a *RatingControl) Label(value interface{}) *RatingControl {
	a.Set("label", value)
	return a
}

func (a *RatingControl) LabelAlign(value interface{}) *RatingControl {
	a.Set("labelAlign", value)
	return a
}

func (a *RatingControl) LabelClassName(value interface{}) *RatingControl {
	a.Set("labelClassName", value)
	return a
}

func (a *RatingControl) LabelRemark(value interface{}) *RatingControl {
	a.Set("labelRemark", value)
	return a
}

func (a *RatingControl) LabelWidth(value interface{}) *RatingControl {
	a.Set("labelWidth", value)
	return a
}

func (a *RatingControl) Mode(value interface{}) *RatingControl {
	a.Set("mode", value)
	return a
}

func (a *RatingControl) Name(value interface{}) *RatingControl {
	a.Set("name", value)
	return a
}

func (a *RatingControl) OnEvent(value interface{}) *RatingControl {
	a.Set("onEvent", value)
	return a
}

func (a *RatingControl) Placeholder(value interface{}) *RatingControl {
	a.Set("placeholder", value)
	return a
}

func (a *RatingControl) ReadOnly(value interface{}) *RatingControl {
	a.Set("readOnly", value)
	return a
}

func (a *RatingControl) ReadOnlyOn(value interface{}) *RatingControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *RatingControl) Remark(value interface{}) *RatingControl {
	a.Set("remark", value)
	return a
}

func (a *RatingControl) Required(value interface{}) *RatingControl {
	a.Set("required", value)
	return a
}

func (a *RatingControl) Row(value interface{}) *RatingControl {
	a.Set("row", value)
	return a
}

func (a *RatingControl) SaveImmediately(value interface{}) *RatingControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *RatingControl) Size(value interface{}) *RatingControl {
	a.Set("size", value)
	return a
}

func (a *RatingControl) Static(value interface{}) *RatingControl {
	a.Set("static", value)
	return a
}

func (a *RatingControl) StaticClassName(value interface{}) *RatingControl {
	a.Set("staticClassName", value)
	return a
}

func (a *RatingControl) StaticInputClassName(value interface{}) *RatingControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *RatingControl) StaticLabelClassName(value interface{}) *RatingControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *RatingControl) StaticOn(value interface{}) *RatingControl {
	a.Set("staticOn", value)
	return a
}

func (a *RatingControl) StaticPlaceholder(value interface{}) *RatingControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *RatingControl) StaticSchema(value interface{}) *RatingControl {
	a.Set("staticSchema", value)
	return a
}

func (a *RatingControl) Style(value interface{}) *RatingControl {
	a.Set("style", value)
	return a
}

func (a *RatingControl) SubmitOnChange(value interface{}) *RatingControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *RatingControl) TestIdBuilder(value interface{}) *RatingControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *RatingControl) TextClassName(value interface{}) *RatingControl {
	a.Set("textClassName", value)
	return a
}

func (a *RatingControl) TextPosition(value interface{}) *RatingControl {
	a.Set("textPosition", value)
	return a
}

func (a *RatingControl) Texts(value interface{}) *RatingControl {
	a.Set("texts", value)
	return a
}

func (a *RatingControl) Type(value interface{}) *RatingControl {
	a.Set("type", value)
	return a
}

func (a *RatingControl) UseMobileUI(value interface{}) *RatingControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *RatingControl) ValidateApi(value interface{}) *RatingControl {
	a.Set("validateApi", value)
	return a
}

func (a *RatingControl) ValidateOnChange(value interface{}) *RatingControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *RatingControl) ValidationErrors(value interface{}) *RatingControl {
	a.Set("validationErrors", value)
	return a
}

func (a *RatingControl) Validations(value interface{}) *RatingControl {
	a.Set("validations", value)
	return a
}

func (a *RatingControl) Value(value interface{}) *RatingControl {
	a.Set("value", value)
	return a
}

func (a *RatingControl) Visible(value interface{}) *RatingControl {
	a.Set("visible", value)
	return a
}

func (a *RatingControl) VisibleOn(value interface{}) *RatingControl {
	a.Set("visibleOn", value)
	return a
}

func (a *RatingControl) Width(value interface{}) *RatingControl {
	a.Set("width", value)
	return a
}
