package renderers

type DiffControl struct {
	*BaseRenderer
}

func NewDiffControl() *DiffControl {
	a := &DiffControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "diff-editor")
	return a
}

func (a *DiffControl) Set(name string, value interface{}) *DiffControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *DiffControl) AutoFill(value interface{}) *DiffControl {
	a.Set("autoFill", value)
	return a
}

func (a *DiffControl) ClassName(value interface{}) *DiffControl {
	a.Set("className", value)
	return a
}

func (a *DiffControl) ClearValueOnHidden(value interface{}) *DiffControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *DiffControl) Desc(value interface{}) *DiffControl {
	a.Set("desc", value)
	return a
}

func (a *DiffControl) Description(value interface{}) *DiffControl {
	a.Set("description", value)
	return a
}

func (a *DiffControl) DescriptionClassName(value interface{}) *DiffControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *DiffControl) DiffValue(value interface{}) *DiffControl {
	a.Set("diffValue", value)
	return a
}

func (a *DiffControl) Disabled(value interface{}) *DiffControl {
	a.Set("disabled", value)
	return a
}

func (a *DiffControl) DisabledOn(value interface{}) *DiffControl {
	a.Set("disabledOn", value)
	return a
}

func (a *DiffControl) EditorSetting(value interface{}) *DiffControl {
	a.Set("editorSetting", value)
	return a
}

func (a *DiffControl) ExtraName(value interface{}) *DiffControl {
	a.Set("extraName", value)
	return a
}

func (a *DiffControl) Hidden(value interface{}) *DiffControl {
	a.Set("hidden", value)
	return a
}

func (a *DiffControl) HiddenOn(value interface{}) *DiffControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *DiffControl) Hint(value interface{}) *DiffControl {
	a.Set("hint", value)
	return a
}

func (a *DiffControl) Horizontal(value interface{}) *DiffControl {
	a.Set("horizontal", value)
	return a
}

func (a *DiffControl) Id(value interface{}) *DiffControl {
	a.Set("id", value)
	return a
}

func (a *DiffControl) InitAutoFill(value interface{}) *DiffControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *DiffControl) Inline(value interface{}) *DiffControl {
	a.Set("inline", value)
	return a
}

func (a *DiffControl) InputClassName(value interface{}) *DiffControl {
	a.Set("inputClassName", value)
	return a
}

func (a *DiffControl) Label(value interface{}) *DiffControl {
	a.Set("label", value)
	return a
}

func (a *DiffControl) LabelAlign(value interface{}) *DiffControl {
	a.Set("labelAlign", value)
	return a
}

func (a *DiffControl) LabelClassName(value interface{}) *DiffControl {
	a.Set("labelClassName", value)
	return a
}

func (a *DiffControl) LabelRemark(value interface{}) *DiffControl {
	a.Set("labelRemark", value)
	return a
}

func (a *DiffControl) LabelWidth(value interface{}) *DiffControl {
	a.Set("labelWidth", value)
	return a
}

func (a *DiffControl) Language(value interface{}) *DiffControl {
	a.Set("language", value)
	return a
}

func (a *DiffControl) Mode(value interface{}) *DiffControl {
	a.Set("mode", value)
	return a
}

func (a *DiffControl) Name(value interface{}) *DiffControl {
	a.Set("name", value)
	return a
}

func (a *DiffControl) OnEvent(value interface{}) *DiffControl {
	a.Set("onEvent", value)
	return a
}

func (a *DiffControl) Options(value interface{}) *DiffControl {
	a.Set("options", value)
	return a
}

func (a *DiffControl) Placeholder(value interface{}) *DiffControl {
	a.Set("placeholder", value)
	return a
}

func (a *DiffControl) ReadOnly(value interface{}) *DiffControl {
	a.Set("readOnly", value)
	return a
}

func (a *DiffControl) ReadOnlyOn(value interface{}) *DiffControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *DiffControl) Remark(value interface{}) *DiffControl {
	a.Set("remark", value)
	return a
}

func (a *DiffControl) Required(value interface{}) *DiffControl {
	a.Set("required", value)
	return a
}

func (a *DiffControl) Row(value interface{}) *DiffControl {
	a.Set("row", value)
	return a
}

func (a *DiffControl) SaveImmediately(value interface{}) *DiffControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *DiffControl) Size(value interface{}) *DiffControl {
	a.Set("size", value)
	return a
}

func (a *DiffControl) Static(value interface{}) *DiffControl {
	a.Set("static", value)
	return a
}

func (a *DiffControl) StaticClassName(value interface{}) *DiffControl {
	a.Set("staticClassName", value)
	return a
}

func (a *DiffControl) StaticInputClassName(value interface{}) *DiffControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *DiffControl) StaticLabelClassName(value interface{}) *DiffControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *DiffControl) StaticOn(value interface{}) *DiffControl {
	a.Set("staticOn", value)
	return a
}

func (a *DiffControl) StaticPlaceholder(value interface{}) *DiffControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *DiffControl) StaticSchema(value interface{}) *DiffControl {
	a.Set("staticSchema", value)
	return a
}

func (a *DiffControl) Style(value interface{}) *DiffControl {
	a.Set("style", value)
	return a
}

func (a *DiffControl) SubmitOnChange(value interface{}) *DiffControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *DiffControl) TestIdBuilder(value interface{}) *DiffControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *DiffControl) Type(value interface{}) *DiffControl {
	a.Set("type", value)
	return a
}

func (a *DiffControl) UseMobileUI(value interface{}) *DiffControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *DiffControl) ValidateApi(value interface{}) *DiffControl {
	a.Set("validateApi", value)
	return a
}

func (a *DiffControl) ValidateOnChange(value interface{}) *DiffControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *DiffControl) ValidationErrors(value interface{}) *DiffControl {
	a.Set("validationErrors", value)
	return a
}

func (a *DiffControl) Validations(value interface{}) *DiffControl {
	a.Set("validations", value)
	return a
}

func (a *DiffControl) Value(value interface{}) *DiffControl {
	a.Set("value", value)
	return a
}

func (a *DiffControl) Visible(value interface{}) *DiffControl {
	a.Set("visible", value)
	return a
}

func (a *DiffControl) VisibleOn(value interface{}) *DiffControl {
	a.Set("visibleOn", value)
	return a
}

func (a *DiffControl) Width(value interface{}) *DiffControl {
	a.Set("width", value)
	return a
}
