package renderers

type EditorControl struct {
	*BaseRenderer
}

func NewEditorControl() *EditorControl {
	a := &EditorControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "editor")
	return a
}

func (a *EditorControl) Set(name string, value interface{}) *EditorControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *EditorControl) AllowFullscreen(value interface{}) *EditorControl {
	a.Set("allowFullscreen", value)
	return a
}

func (a *EditorControl) AutoFill(value interface{}) *EditorControl {
	a.Set("autoFill", value)
	return a
}

func (a *EditorControl) ClassName(value interface{}) *EditorControl {
	a.Set("className", value)
	return a
}

func (a *EditorControl) ClearValueOnHidden(value interface{}) *EditorControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *EditorControl) Desc(value interface{}) *EditorControl {
	a.Set("desc", value)
	return a
}

func (a *EditorControl) Description(value interface{}) *EditorControl {
	a.Set("description", value)
	return a
}

func (a *EditorControl) DescriptionClassName(value interface{}) *EditorControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *EditorControl) Disabled(value interface{}) *EditorControl {
	a.Set("disabled", value)
	return a
}

func (a *EditorControl) DisabledOn(value interface{}) *EditorControl {
	a.Set("disabledOn", value)
	return a
}

func (a *EditorControl) EditorDidMount(value interface{}) *EditorControl {
	a.Set("editorDidMount", value)
	return a
}

func (a *EditorControl) EditorSetting(value interface{}) *EditorControl {
	a.Set("editorSetting", value)
	return a
}

func (a *EditorControl) ExtraName(value interface{}) *EditorControl {
	a.Set("extraName", value)
	return a
}

func (a *EditorControl) Hidden(value interface{}) *EditorControl {
	a.Set("hidden", value)
	return a
}

func (a *EditorControl) HiddenOn(value interface{}) *EditorControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *EditorControl) Hint(value interface{}) *EditorControl {
	a.Set("hint", value)
	return a
}

func (a *EditorControl) Horizontal(value interface{}) *EditorControl {
	a.Set("horizontal", value)
	return a
}

func (a *EditorControl) Id(value interface{}) *EditorControl {
	a.Set("id", value)
	return a
}

func (a *EditorControl) InitAutoFill(value interface{}) *EditorControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *EditorControl) Inline(value interface{}) *EditorControl {
	a.Set("inline", value)
	return a
}

func (a *EditorControl) InputClassName(value interface{}) *EditorControl {
	a.Set("inputClassName", value)
	return a
}

func (a *EditorControl) Label(value interface{}) *EditorControl {
	a.Set("label", value)
	return a
}

func (a *EditorControl) LabelAlign(value interface{}) *EditorControl {
	a.Set("labelAlign", value)
	return a
}

func (a *EditorControl) LabelClassName(value interface{}) *EditorControl {
	a.Set("labelClassName", value)
	return a
}

func (a *EditorControl) LabelRemark(value interface{}) *EditorControl {
	a.Set("labelRemark", value)
	return a
}

func (a *EditorControl) LabelWidth(value interface{}) *EditorControl {
	a.Set("labelWidth", value)
	return a
}

func (a *EditorControl) Language(value interface{}) *EditorControl {
	a.Set("language", value)
	return a
}

func (a *EditorControl) Mode(value interface{}) *EditorControl {
	a.Set("mode", value)
	return a
}

func (a *EditorControl) Name(value interface{}) *EditorControl {
	a.Set("name", value)
	return a
}

func (a *EditorControl) OnEvent(value interface{}) *EditorControl {
	a.Set("onEvent", value)
	return a
}

func (a *EditorControl) Options(value interface{}) *EditorControl {
	a.Set("options", value)
	return a
}

func (a *EditorControl) Placeholder(value interface{}) *EditorControl {
	a.Set("placeholder", value)
	return a
}

func (a *EditorControl) ReadOnly(value interface{}) *EditorControl {
	a.Set("readOnly", value)
	return a
}

func (a *EditorControl) ReadOnlyOn(value interface{}) *EditorControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *EditorControl) Remark(value interface{}) *EditorControl {
	a.Set("remark", value)
	return a
}

func (a *EditorControl) Required(value interface{}) *EditorControl {
	a.Set("required", value)
	return a
}

func (a *EditorControl) Row(value interface{}) *EditorControl {
	a.Set("row", value)
	return a
}

func (a *EditorControl) SaveImmediately(value interface{}) *EditorControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *EditorControl) Size(value interface{}) *EditorControl {
	a.Set("size", value)
	return a
}

func (a *EditorControl) Static(value interface{}) *EditorControl {
	a.Set("static", value)
	return a
}

func (a *EditorControl) StaticClassName(value interface{}) *EditorControl {
	a.Set("staticClassName", value)
	return a
}

func (a *EditorControl) StaticInputClassName(value interface{}) *EditorControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *EditorControl) StaticLabelClassName(value interface{}) *EditorControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *EditorControl) StaticOn(value interface{}) *EditorControl {
	a.Set("staticOn", value)
	return a
}

func (a *EditorControl) StaticPlaceholder(value interface{}) *EditorControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *EditorControl) StaticSchema(value interface{}) *EditorControl {
	a.Set("staticSchema", value)
	return a
}

func (a *EditorControl) Style(value interface{}) *EditorControl {
	a.Set("style", value)
	return a
}

func (a *EditorControl) SubmitOnChange(value interface{}) *EditorControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *EditorControl) TestIdBuilder(value interface{}) *EditorControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *EditorControl) Type(value interface{}) *EditorControl {
	a.Set("type", value)
	return a
}

func (a *EditorControl) UseMobileUI(value interface{}) *EditorControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *EditorControl) ValidateApi(value interface{}) *EditorControl {
	a.Set("validateApi", value)
	return a
}

func (a *EditorControl) ValidateOnChange(value interface{}) *EditorControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *EditorControl) ValidationErrors(value interface{}) *EditorControl {
	a.Set("validationErrors", value)
	return a
}

func (a *EditorControl) Validations(value interface{}) *EditorControl {
	a.Set("validations", value)
	return a
}

func (a *EditorControl) Value(value interface{}) *EditorControl {
	a.Set("value", value)
	return a
}

func (a *EditorControl) Visible(value interface{}) *EditorControl {
	a.Set("visible", value)
	return a
}

func (a *EditorControl) VisibleOn(value interface{}) *EditorControl {
	a.Set("visibleOn", value)
	return a
}

func (a *EditorControl) Width(value interface{}) *EditorControl {
	a.Set("width", value)
	return a
}
