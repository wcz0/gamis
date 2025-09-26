package renderers

type ConditionBuilderControl struct {
	*BaseRenderer
}

func NewConditionBuilderControl() *ConditionBuilderControl {
	a := &ConditionBuilderControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "condition-builder")
	return a
}

func (a *ConditionBuilderControl) Set(name string, value interface{}) *ConditionBuilderControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *ConditionBuilderControl) AddBtnVisibleOn(value interface{}) *ConditionBuilderControl {
	a.Set("addBtnVisibleOn", value)
	return a
}

func (a *ConditionBuilderControl) AddGroupBtnVisibleOn(value interface{}) *ConditionBuilderControl {
	a.Set("addGroupBtnVisibleOn", value)
	return a
}

func (a *ConditionBuilderControl) AutoFill(value interface{}) *ConditionBuilderControl {
	a.Set("autoFill", value)
	return a
}

func (a *ConditionBuilderControl) BuilderMode(value interface{}) *ConditionBuilderControl {
	a.Set("builderMode", value)
	return a
}

func (a *ConditionBuilderControl) ClassName(value interface{}) *ConditionBuilderControl {
	a.Set("className", value)
	return a
}

func (a *ConditionBuilderControl) ClearValueOnHidden(value interface{}) *ConditionBuilderControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *ConditionBuilderControl) Config(value interface{}) *ConditionBuilderControl {
	a.Set("config", value)
	return a
}

func (a *ConditionBuilderControl) Desc(value interface{}) *ConditionBuilderControl {
	a.Set("desc", value)
	return a
}

func (a *ConditionBuilderControl) Description(value interface{}) *ConditionBuilderControl {
	a.Set("description", value)
	return a
}

func (a *ConditionBuilderControl) DescriptionClassName(value interface{}) *ConditionBuilderControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *ConditionBuilderControl) Disabled(value interface{}) *ConditionBuilderControl {
	a.Set("disabled", value)
	return a
}

func (a *ConditionBuilderControl) DisabledOn(value interface{}) *ConditionBuilderControl {
	a.Set("disabledOn", value)
	return a
}

func (a *ConditionBuilderControl) Draggable(value interface{}) *ConditionBuilderControl {
	a.Set("draggable", value)
	return a
}

func (a *ConditionBuilderControl) EditorSetting(value interface{}) *ConditionBuilderControl {
	a.Set("editorSetting", value)
	return a
}

func (a *ConditionBuilderControl) Embed(value interface{}) *ConditionBuilderControl {
	a.Set("embed", value)
	return a
}

func (a *ConditionBuilderControl) ExtraName(value interface{}) *ConditionBuilderControl {
	a.Set("extraName", value)
	return a
}

func (a *ConditionBuilderControl) Fields(value interface{}) *ConditionBuilderControl {
	a.Set("fields", value)
	return a
}

func (a *ConditionBuilderControl) Formula(value interface{}) *ConditionBuilderControl {
	a.Set("formula", value)
	return a
}

func (a *ConditionBuilderControl) FormulaForIf(value interface{}) *ConditionBuilderControl {
	a.Set("formulaForIf", value)
	return a
}

func (a *ConditionBuilderControl) Funcs(value interface{}) *ConditionBuilderControl {
	a.Set("funcs", value)
	return a
}

func (a *ConditionBuilderControl) Hidden(value interface{}) *ConditionBuilderControl {
	a.Set("hidden", value)
	return a
}

func (a *ConditionBuilderControl) HiddenOn(value interface{}) *ConditionBuilderControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *ConditionBuilderControl) Hint(value interface{}) *ConditionBuilderControl {
	a.Set("hint", value)
	return a
}

func (a *ConditionBuilderControl) Horizontal(value interface{}) *ConditionBuilderControl {
	a.Set("horizontal", value)
	return a
}

func (a *ConditionBuilderControl) Id(value interface{}) *ConditionBuilderControl {
	a.Set("id", value)
	return a
}

func (a *ConditionBuilderControl) InitAutoFill(value interface{}) *ConditionBuilderControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *ConditionBuilderControl) Inline(value interface{}) *ConditionBuilderControl {
	a.Set("inline", value)
	return a
}

func (a *ConditionBuilderControl) InputClassName(value interface{}) *ConditionBuilderControl {
	a.Set("inputClassName", value)
	return a
}

func (a *ConditionBuilderControl) Label(value interface{}) *ConditionBuilderControl {
	a.Set("label", value)
	return a
}

func (a *ConditionBuilderControl) LabelAlign(value interface{}) *ConditionBuilderControl {
	a.Set("labelAlign", value)
	return a
}

func (a *ConditionBuilderControl) LabelClassName(value interface{}) *ConditionBuilderControl {
	a.Set("labelClassName", value)
	return a
}

func (a *ConditionBuilderControl) LabelRemark(value interface{}) *ConditionBuilderControl {
	a.Set("labelRemark", value)
	return a
}

func (a *ConditionBuilderControl) LabelWidth(value interface{}) *ConditionBuilderControl {
	a.Set("labelWidth", value)
	return a
}

func (a *ConditionBuilderControl) Mode(value interface{}) *ConditionBuilderControl {
	a.Set("mode", value)
	return a
}

func (a *ConditionBuilderControl) Name(value interface{}) *ConditionBuilderControl {
	a.Set("name", value)
	return a
}

func (a *ConditionBuilderControl) OnEvent(value interface{}) *ConditionBuilderControl {
	a.Set("onEvent", value)
	return a
}

func (a *ConditionBuilderControl) PickerIcon(value interface{}) *ConditionBuilderControl {
	a.Set("pickerIcon", value)
	return a
}

func (a *ConditionBuilderControl) Placeholder(value interface{}) *ConditionBuilderControl {
	a.Set("placeholder", value)
	return a
}

func (a *ConditionBuilderControl) ReadOnly(value interface{}) *ConditionBuilderControl {
	a.Set("readOnly", value)
	return a
}

func (a *ConditionBuilderControl) ReadOnlyOn(value interface{}) *ConditionBuilderControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *ConditionBuilderControl) Remark(value interface{}) *ConditionBuilderControl {
	a.Set("remark", value)
	return a
}

func (a *ConditionBuilderControl) Required(value interface{}) *ConditionBuilderControl {
	a.Set("required", value)
	return a
}

func (a *ConditionBuilderControl) Row(value interface{}) *ConditionBuilderControl {
	a.Set("row", value)
	return a
}

func (a *ConditionBuilderControl) SaveImmediately(value interface{}) *ConditionBuilderControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *ConditionBuilderControl) ShowANDOR(value interface{}) *ConditionBuilderControl {
	a.Set("showANDOR", value)
	return a
}

func (a *ConditionBuilderControl) Size(value interface{}) *ConditionBuilderControl {
	a.Set("size", value)
	return a
}

func (a *ConditionBuilderControl) Source(value interface{}) *ConditionBuilderControl {
	a.Set("source", value)
	return a
}

func (a *ConditionBuilderControl) Static(value interface{}) *ConditionBuilderControl {
	a.Set("static", value)
	return a
}

func (a *ConditionBuilderControl) StaticClassName(value interface{}) *ConditionBuilderControl {
	a.Set("staticClassName", value)
	return a
}

func (a *ConditionBuilderControl) StaticInputClassName(value interface{}) *ConditionBuilderControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *ConditionBuilderControl) StaticLabelClassName(value interface{}) *ConditionBuilderControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *ConditionBuilderControl) StaticOn(value interface{}) *ConditionBuilderControl {
	a.Set("staticOn", value)
	return a
}

func (a *ConditionBuilderControl) StaticPlaceholder(value interface{}) *ConditionBuilderControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *ConditionBuilderControl) StaticSchema(value interface{}) *ConditionBuilderControl {
	a.Set("staticSchema", value)
	return a
}

func (a *ConditionBuilderControl) Style(value interface{}) *ConditionBuilderControl {
	a.Set("style", value)
	return a
}

func (a *ConditionBuilderControl) SubmitOnChange(value interface{}) *ConditionBuilderControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *ConditionBuilderControl) TestIdBuilder(value interface{}) *ConditionBuilderControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *ConditionBuilderControl) Type(value interface{}) *ConditionBuilderControl {
	a.Set("type", value)
	return a
}

func (a *ConditionBuilderControl) UseMobileUI(value interface{}) *ConditionBuilderControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *ConditionBuilderControl) ValidateApi(value interface{}) *ConditionBuilderControl {
	a.Set("validateApi", value)
	return a
}

func (a *ConditionBuilderControl) ValidateOnChange(value interface{}) *ConditionBuilderControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *ConditionBuilderControl) ValidationErrors(value interface{}) *ConditionBuilderControl {
	a.Set("validationErrors", value)
	return a
}

func (a *ConditionBuilderControl) Validations(value interface{}) *ConditionBuilderControl {
	a.Set("validations", value)
	return a
}

func (a *ConditionBuilderControl) Value(value interface{}) *ConditionBuilderControl {
	a.Set("value", value)
	return a
}

func (a *ConditionBuilderControl) Visible(value interface{}) *ConditionBuilderControl {
	a.Set("visible", value)
	return a
}

func (a *ConditionBuilderControl) VisibleOn(value interface{}) *ConditionBuilderControl {
	a.Set("visibleOn", value)
	return a
}

func (a *ConditionBuilderControl) Width(value interface{}) *ConditionBuilderControl {
	a.Set("width", value)
	return a
}
