package renderers

type FieldSetControl struct {
	*BaseRenderer
}

func NewFieldSetControl() *FieldSetControl {
	a := &FieldSetControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "fieldset")
	return a
}

func (a *FieldSetControl) Set(name string, value interface{}) *FieldSetControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *FieldSetControl) AutoFill(value interface{}) *FieldSetControl {
	a.Set("autoFill", value)
	return a
}

func (a *FieldSetControl) Body(value interface{}) *FieldSetControl {
	a.Set("body", value)
	return a
}

func (a *FieldSetControl) BodyClassName(value interface{}) *FieldSetControl {
	a.Set("bodyClassName", value)
	return a
}

func (a *FieldSetControl) ClassName(value interface{}) *FieldSetControl {
	a.Set("className", value)
	return a
}

func (a *FieldSetControl) ClearValueOnHidden(value interface{}) *FieldSetControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *FieldSetControl) Collapsable(value interface{}) *FieldSetControl {
	a.Set("collapsable", value)
	return a
}

func (a *FieldSetControl) CollapseHeader(value interface{}) *FieldSetControl {
	a.Set("collapseHeader", value)
	return a
}

func (a *FieldSetControl) CollapseTitle(value interface{}) *FieldSetControl {
	a.Set("collapseTitle", value)
	return a
}

func (a *FieldSetControl) Collapsed(value interface{}) *FieldSetControl {
	a.Set("collapsed", value)
	return a
}

func (a *FieldSetControl) Desc(value interface{}) *FieldSetControl {
	a.Set("desc", value)
	return a
}

func (a *FieldSetControl) Description(value interface{}) *FieldSetControl {
	a.Set("description", value)
	return a
}

func (a *FieldSetControl) DescriptionClassName(value interface{}) *FieldSetControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *FieldSetControl) Disabled(value interface{}) *FieldSetControl {
	a.Set("disabled", value)
	return a
}

func (a *FieldSetControl) DisabledOn(value interface{}) *FieldSetControl {
	a.Set("disabledOn", value)
	return a
}

func (a *FieldSetControl) DivideLine(value interface{}) *FieldSetControl {
	a.Set("divideLine", value)
	return a
}

func (a *FieldSetControl) EditorSetting(value interface{}) *FieldSetControl {
	a.Set("editorSetting", value)
	return a
}

func (a *FieldSetControl) ExpandIcon(value interface{}) *FieldSetControl {
	a.Set("expandIcon", value)
	return a
}

func (a *FieldSetControl) ExtraName(value interface{}) *FieldSetControl {
	a.Set("extraName", value)
	return a
}

func (a *FieldSetControl) Header(value interface{}) *FieldSetControl {
	a.Set("header", value)
	return a
}

func (a *FieldSetControl) HeaderPosition(value interface{}) *FieldSetControl {
	a.Set("headerPosition", value)
	return a
}

func (a *FieldSetControl) HeadingClassName(value interface{}) *FieldSetControl {
	a.Set("headingClassName", value)
	return a
}

func (a *FieldSetControl) Hidden(value interface{}) *FieldSetControl {
	a.Set("hidden", value)
	return a
}

func (a *FieldSetControl) HiddenOn(value interface{}) *FieldSetControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *FieldSetControl) Hint(value interface{}) *FieldSetControl {
	a.Set("hint", value)
	return a
}

func (a *FieldSetControl) Horizontal(value interface{}) *FieldSetControl {
	a.Set("horizontal", value)
	return a
}

func (a *FieldSetControl) Id(value interface{}) *FieldSetControl {
	a.Set("id", value)
	return a
}

func (a *FieldSetControl) InitAutoFill(value interface{}) *FieldSetControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *FieldSetControl) Inline(value interface{}) *FieldSetControl {
	a.Set("inline", value)
	return a
}

func (a *FieldSetControl) InputClassName(value interface{}) *FieldSetControl {
	a.Set("inputClassName", value)
	return a
}

func (a *FieldSetControl) Key(value interface{}) *FieldSetControl {
	a.Set("key", value)
	return a
}

func (a *FieldSetControl) Label(value interface{}) *FieldSetControl {
	a.Set("label", value)
	return a
}

func (a *FieldSetControl) LabelAlign(value interface{}) *FieldSetControl {
	a.Set("labelAlign", value)
	return a
}

func (a *FieldSetControl) LabelClassName(value interface{}) *FieldSetControl {
	a.Set("labelClassName", value)
	return a
}

func (a *FieldSetControl) LabelRemark(value interface{}) *FieldSetControl {
	a.Set("labelRemark", value)
	return a
}

func (a *FieldSetControl) LabelWidth(value interface{}) *FieldSetControl {
	a.Set("labelWidth", value)
	return a
}

func (a *FieldSetControl) Mode(value interface{}) *FieldSetControl {
	a.Set("mode", value)
	return a
}

func (a *FieldSetControl) MountOnEnter(value interface{}) *FieldSetControl {
	a.Set("mountOnEnter", value)
	return a
}

func (a *FieldSetControl) Name(value interface{}) *FieldSetControl {
	a.Set("name", value)
	return a
}

func (a *FieldSetControl) OnEvent(value interface{}) *FieldSetControl {
	a.Set("onEvent", value)
	return a
}

func (a *FieldSetControl) Placeholder(value interface{}) *FieldSetControl {
	a.Set("placeholder", value)
	return a
}

func (a *FieldSetControl) ReadOnly(value interface{}) *FieldSetControl {
	a.Set("readOnly", value)
	return a
}

func (a *FieldSetControl) ReadOnlyOn(value interface{}) *FieldSetControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *FieldSetControl) Remark(value interface{}) *FieldSetControl {
	a.Set("remark", value)
	return a
}

func (a *FieldSetControl) Required(value interface{}) *FieldSetControl {
	a.Set("required", value)
	return a
}

func (a *FieldSetControl) Row(value interface{}) *FieldSetControl {
	a.Set("row", value)
	return a
}

func (a *FieldSetControl) SaveImmediately(value interface{}) *FieldSetControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *FieldSetControl) ShowArrow(value interface{}) *FieldSetControl {
	a.Set("showArrow", value)
	return a
}

func (a *FieldSetControl) Size(value interface{}) *FieldSetControl {
	a.Set("size", value)
	return a
}

func (a *FieldSetControl) Static(value interface{}) *FieldSetControl {
	a.Set("static", value)
	return a
}

func (a *FieldSetControl) StaticClassName(value interface{}) *FieldSetControl {
	a.Set("staticClassName", value)
	return a
}

func (a *FieldSetControl) StaticInputClassName(value interface{}) *FieldSetControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *FieldSetControl) StaticLabelClassName(value interface{}) *FieldSetControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *FieldSetControl) StaticOn(value interface{}) *FieldSetControl {
	a.Set("staticOn", value)
	return a
}

func (a *FieldSetControl) StaticPlaceholder(value interface{}) *FieldSetControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *FieldSetControl) StaticSchema(value interface{}) *FieldSetControl {
	a.Set("staticSchema", value)
	return a
}

func (a *FieldSetControl) Style(value interface{}) *FieldSetControl {
	a.Set("style", value)
	return a
}

func (a *FieldSetControl) SubFormHorizontal(value interface{}) *FieldSetControl {
	a.Set("subFormHorizontal", value)
	return a
}

func (a *FieldSetControl) SubFormMode(value interface{}) *FieldSetControl {
	a.Set("subFormMode", value)
	return a
}

func (a *FieldSetControl) SubmitOnChange(value interface{}) *FieldSetControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *FieldSetControl) TestIdBuilder(value interface{}) *FieldSetControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *FieldSetControl) Testid(value interface{}) *FieldSetControl {
	a.Set("testid", value)
	return a
}

func (a *FieldSetControl) Title(value interface{}) *FieldSetControl {
	a.Set("title", value)
	return a
}

func (a *FieldSetControl) TitlePosition(value interface{}) *FieldSetControl {
	a.Set("titlePosition", value)
	return a
}

func (a *FieldSetControl) Type(value interface{}) *FieldSetControl {
	a.Set("type", value)
	return a
}

func (a *FieldSetControl) UnmountOnExit(value interface{}) *FieldSetControl {
	a.Set("unmountOnExit", value)
	return a
}

func (a *FieldSetControl) UseMobileUI(value interface{}) *FieldSetControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *FieldSetControl) ValidateApi(value interface{}) *FieldSetControl {
	a.Set("validateApi", value)
	return a
}

func (a *FieldSetControl) ValidateOnChange(value interface{}) *FieldSetControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *FieldSetControl) ValidationErrors(value interface{}) *FieldSetControl {
	a.Set("validationErrors", value)
	return a
}

func (a *FieldSetControl) Validations(value interface{}) *FieldSetControl {
	a.Set("validations", value)
	return a
}

func (a *FieldSetControl) Value(value interface{}) *FieldSetControl {
	a.Set("value", value)
	return a
}

func (a *FieldSetControl) Visible(value interface{}) *FieldSetControl {
	a.Set("visible", value)
	return a
}

func (a *FieldSetControl) VisibleOn(value interface{}) *FieldSetControl {
	a.Set("visibleOn", value)
	return a
}

func (a *FieldSetControl) Width(value interface{}) *FieldSetControl {
	a.Set("width", value)
	return a
}
