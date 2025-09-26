package renderers

type MatrixControl struct {
	*BaseRenderer
}

func NewMatrixControl() *MatrixControl {
	a := &MatrixControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "matrix-checkboxes")
	return a
}

func (a *MatrixControl) Set(name string, value interface{}) *MatrixControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *MatrixControl) AutoFill(value interface{}) *MatrixControl {
	a.Set("autoFill", value)
	return a
}

func (a *MatrixControl) ClassName(value interface{}) *MatrixControl {
	a.Set("className", value)
	return a
}

func (a *MatrixControl) ClearValueOnHidden(value interface{}) *MatrixControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *MatrixControl) Columns(value interface{}) *MatrixControl {
	a.Set("columns", value)
	return a
}

func (a *MatrixControl) Desc(value interface{}) *MatrixControl {
	a.Set("desc", value)
	return a
}

func (a *MatrixControl) Description(value interface{}) *MatrixControl {
	a.Set("description", value)
	return a
}

func (a *MatrixControl) DescriptionClassName(value interface{}) *MatrixControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *MatrixControl) Disabled(value interface{}) *MatrixControl {
	a.Set("disabled", value)
	return a
}

func (a *MatrixControl) DisabledOn(value interface{}) *MatrixControl {
	a.Set("disabledOn", value)
	return a
}

func (a *MatrixControl) EditorSetting(value interface{}) *MatrixControl {
	a.Set("editorSetting", value)
	return a
}

func (a *MatrixControl) ExtraName(value interface{}) *MatrixControl {
	a.Set("extraName", value)
	return a
}

func (a *MatrixControl) Hidden(value interface{}) *MatrixControl {
	a.Set("hidden", value)
	return a
}

func (a *MatrixControl) HiddenOn(value interface{}) *MatrixControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *MatrixControl) Hint(value interface{}) *MatrixControl {
	a.Set("hint", value)
	return a
}

func (a *MatrixControl) Horizontal(value interface{}) *MatrixControl {
	a.Set("horizontal", value)
	return a
}

func (a *MatrixControl) Id(value interface{}) *MatrixControl {
	a.Set("id", value)
	return a
}

func (a *MatrixControl) InitAutoFill(value interface{}) *MatrixControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *MatrixControl) Inline(value interface{}) *MatrixControl {
	a.Set("inline", value)
	return a
}

func (a *MatrixControl) InputClassName(value interface{}) *MatrixControl {
	a.Set("inputClassName", value)
	return a
}

func (a *MatrixControl) Label(value interface{}) *MatrixControl {
	a.Set("label", value)
	return a
}

func (a *MatrixControl) LabelAlign(value interface{}) *MatrixControl {
	a.Set("labelAlign", value)
	return a
}

func (a *MatrixControl) LabelClassName(value interface{}) *MatrixControl {
	a.Set("labelClassName", value)
	return a
}

func (a *MatrixControl) LabelRemark(value interface{}) *MatrixControl {
	a.Set("labelRemark", value)
	return a
}

func (a *MatrixControl) LabelWidth(value interface{}) *MatrixControl {
	a.Set("labelWidth", value)
	return a
}

func (a *MatrixControl) Mode(value interface{}) *MatrixControl {
	a.Set("mode", value)
	return a
}

func (a *MatrixControl) Multiple(value interface{}) *MatrixControl {
	a.Set("multiple", value)
	return a
}

func (a *MatrixControl) Name(value interface{}) *MatrixControl {
	a.Set("name", value)
	return a
}

func (a *MatrixControl) OnEvent(value interface{}) *MatrixControl {
	a.Set("onEvent", value)
	return a
}

func (a *MatrixControl) Placeholder(value interface{}) *MatrixControl {
	a.Set("placeholder", value)
	return a
}

func (a *MatrixControl) ReadOnly(value interface{}) *MatrixControl {
	a.Set("readOnly", value)
	return a
}

func (a *MatrixControl) ReadOnlyOn(value interface{}) *MatrixControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *MatrixControl) Remark(value interface{}) *MatrixControl {
	a.Set("remark", value)
	return a
}

func (a *MatrixControl) Required(value interface{}) *MatrixControl {
	a.Set("required", value)
	return a
}

func (a *MatrixControl) Row(value interface{}) *MatrixControl {
	a.Set("row", value)
	return a
}

func (a *MatrixControl) RowLabel(value interface{}) *MatrixControl {
	a.Set("rowLabel", value)
	return a
}

func (a *MatrixControl) Rows(value interface{}) *MatrixControl {
	a.Set("rows", value)
	return a
}

func (a *MatrixControl) SaveImmediately(value interface{}) *MatrixControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *MatrixControl) SingleSelectMode(value interface{}) *MatrixControl {
	a.Set("singleSelectMode", value)
	return a
}

func (a *MatrixControl) Size(value interface{}) *MatrixControl {
	a.Set("size", value)
	return a
}

func (a *MatrixControl) Source(value interface{}) *MatrixControl {
	a.Set("source", value)
	return a
}

func (a *MatrixControl) Static(value interface{}) *MatrixControl {
	a.Set("static", value)
	return a
}

func (a *MatrixControl) StaticClassName(value interface{}) *MatrixControl {
	a.Set("staticClassName", value)
	return a
}

func (a *MatrixControl) StaticInputClassName(value interface{}) *MatrixControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *MatrixControl) StaticLabelClassName(value interface{}) *MatrixControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *MatrixControl) StaticOn(value interface{}) *MatrixControl {
	a.Set("staticOn", value)
	return a
}

func (a *MatrixControl) StaticPlaceholder(value interface{}) *MatrixControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *MatrixControl) StaticSchema(value interface{}) *MatrixControl {
	a.Set("staticSchema", value)
	return a
}

func (a *MatrixControl) Style(value interface{}) *MatrixControl {
	a.Set("style", value)
	return a
}

func (a *MatrixControl) SubmitOnChange(value interface{}) *MatrixControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *MatrixControl) TestIdBuilder(value interface{}) *MatrixControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *MatrixControl) Type(value interface{}) *MatrixControl {
	a.Set("type", value)
	return a
}

func (a *MatrixControl) UseMobileUI(value interface{}) *MatrixControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *MatrixControl) ValidateApi(value interface{}) *MatrixControl {
	a.Set("validateApi", value)
	return a
}

func (a *MatrixControl) ValidateOnChange(value interface{}) *MatrixControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *MatrixControl) ValidationErrors(value interface{}) *MatrixControl {
	a.Set("validationErrors", value)
	return a
}

func (a *MatrixControl) Validations(value interface{}) *MatrixControl {
	a.Set("validations", value)
	return a
}

func (a *MatrixControl) Value(value interface{}) *MatrixControl {
	a.Set("value", value)
	return a
}

func (a *MatrixControl) Visible(value interface{}) *MatrixControl {
	a.Set("visible", value)
	return a
}

func (a *MatrixControl) VisibleOn(value interface{}) *MatrixControl {
	a.Set("visibleOn", value)
	return a
}

func (a *MatrixControl) Width(value interface{}) *MatrixControl {
	a.Set("width", value)
	return a
}
