package renderers

type FormulaControl struct {
	*BaseRenderer
}

func NewFormulaControl() *FormulaControl {
	a := &FormulaControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "formula")
	return a
}

func (a *FormulaControl) Set(name string, value interface{}) *FormulaControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *FormulaControl) AutoFill(value interface{}) *FormulaControl {
	a.Set("autoFill", value)
	return a
}

func (a *FormulaControl) AutoSet(value interface{}) *FormulaControl {
	a.Set("autoSet", value)
	return a
}

func (a *FormulaControl) ClassName(value interface{}) *FormulaControl {
	a.Set("className", value)
	return a
}

func (a *FormulaControl) ClearValueOnHidden(value interface{}) *FormulaControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *FormulaControl) Condition(value interface{}) *FormulaControl {
	a.Set("condition", value)
	return a
}

func (a *FormulaControl) Desc(value interface{}) *FormulaControl {
	a.Set("desc", value)
	return a
}

func (a *FormulaControl) Description(value interface{}) *FormulaControl {
	a.Set("description", value)
	return a
}

func (a *FormulaControl) DescriptionClassName(value interface{}) *FormulaControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *FormulaControl) Disabled(value interface{}) *FormulaControl {
	a.Set("disabled", value)
	return a
}

func (a *FormulaControl) DisabledOn(value interface{}) *FormulaControl {
	a.Set("disabledOn", value)
	return a
}

func (a *FormulaControl) EditorSetting(value interface{}) *FormulaControl {
	a.Set("editorSetting", value)
	return a
}

func (a *FormulaControl) ExtraName(value interface{}) *FormulaControl {
	a.Set("extraName", value)
	return a
}

func (a *FormulaControl) Formula(value interface{}) *FormulaControl {
	a.Set("formula", value)
	return a
}

func (a *FormulaControl) Hidden(value interface{}) *FormulaControl {
	a.Set("hidden", value)
	return a
}

func (a *FormulaControl) HiddenOn(value interface{}) *FormulaControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *FormulaControl) Hint(value interface{}) *FormulaControl {
	a.Set("hint", value)
	return a
}

func (a *FormulaControl) Horizontal(value interface{}) *FormulaControl {
	a.Set("horizontal", value)
	return a
}

func (a *FormulaControl) Id(value interface{}) *FormulaControl {
	a.Set("id", value)
	return a
}

func (a *FormulaControl) InitAutoFill(value interface{}) *FormulaControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *FormulaControl) InitSet(value interface{}) *FormulaControl {
	a.Set("initSet", value)
	return a
}

func (a *FormulaControl) Inline(value interface{}) *FormulaControl {
	a.Set("inline", value)
	return a
}

func (a *FormulaControl) InputClassName(value interface{}) *FormulaControl {
	a.Set("inputClassName", value)
	return a
}

func (a *FormulaControl) Label(value interface{}) *FormulaControl {
	a.Set("label", value)
	return a
}

func (a *FormulaControl) LabelAlign(value interface{}) *FormulaControl {
	a.Set("labelAlign", value)
	return a
}

func (a *FormulaControl) LabelClassName(value interface{}) *FormulaControl {
	a.Set("labelClassName", value)
	return a
}

func (a *FormulaControl) LabelRemark(value interface{}) *FormulaControl {
	a.Set("labelRemark", value)
	return a
}

func (a *FormulaControl) LabelWidth(value interface{}) *FormulaControl {
	a.Set("labelWidth", value)
	return a
}

func (a *FormulaControl) Mode(value interface{}) *FormulaControl {
	a.Set("mode", value)
	return a
}

func (a *FormulaControl) Name(value interface{}) *FormulaControl {
	a.Set("name", value)
	return a
}

func (a *FormulaControl) OnEvent(value interface{}) *FormulaControl {
	a.Set("onEvent", value)
	return a
}

func (a *FormulaControl) Placeholder(value interface{}) *FormulaControl {
	a.Set("placeholder", value)
	return a
}

func (a *FormulaControl) ReadOnly(value interface{}) *FormulaControl {
	a.Set("readOnly", value)
	return a
}

func (a *FormulaControl) ReadOnlyOn(value interface{}) *FormulaControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *FormulaControl) Remark(value interface{}) *FormulaControl {
	a.Set("remark", value)
	return a
}

func (a *FormulaControl) Required(value interface{}) *FormulaControl {
	a.Set("required", value)
	return a
}

func (a *FormulaControl) Row(value interface{}) *FormulaControl {
	a.Set("row", value)
	return a
}

func (a *FormulaControl) SaveImmediately(value interface{}) *FormulaControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *FormulaControl) Size(value interface{}) *FormulaControl {
	a.Set("size", value)
	return a
}

func (a *FormulaControl) Static(value interface{}) *FormulaControl {
	a.Set("static", value)
	return a
}

func (a *FormulaControl) StaticClassName(value interface{}) *FormulaControl {
	a.Set("staticClassName", value)
	return a
}

func (a *FormulaControl) StaticInputClassName(value interface{}) *FormulaControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *FormulaControl) StaticLabelClassName(value interface{}) *FormulaControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *FormulaControl) StaticOn(value interface{}) *FormulaControl {
	a.Set("staticOn", value)
	return a
}

func (a *FormulaControl) StaticPlaceholder(value interface{}) *FormulaControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *FormulaControl) StaticSchema(value interface{}) *FormulaControl {
	a.Set("staticSchema", value)
	return a
}

func (a *FormulaControl) Style(value interface{}) *FormulaControl {
	a.Set("style", value)
	return a
}

func (a *FormulaControl) SubmitOnChange(value interface{}) *FormulaControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *FormulaControl) TestIdBuilder(value interface{}) *FormulaControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *FormulaControl) Type(value interface{}) *FormulaControl {
	a.Set("type", value)
	return a
}

func (a *FormulaControl) UseMobileUI(value interface{}) *FormulaControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *FormulaControl) ValidateApi(value interface{}) *FormulaControl {
	a.Set("validateApi", value)
	return a
}

func (a *FormulaControl) ValidateOnChange(value interface{}) *FormulaControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *FormulaControl) ValidationErrors(value interface{}) *FormulaControl {
	a.Set("validationErrors", value)
	return a
}

func (a *FormulaControl) Validations(value interface{}) *FormulaControl {
	a.Set("validations", value)
	return a
}

func (a *FormulaControl) Value(value interface{}) *FormulaControl {
	a.Set("value", value)
	return a
}

func (a *FormulaControl) Visible(value interface{}) *FormulaControl {
	a.Set("visible", value)
	return a
}

func (a *FormulaControl) VisibleOn(value interface{}) *FormulaControl {
	a.Set("visibleOn", value)
	return a
}

func (a *FormulaControl) Width(value interface{}) *FormulaControl {
	a.Set("width", value)
	return a
}
