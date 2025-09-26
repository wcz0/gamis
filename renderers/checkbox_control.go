package renderers


/**
 * Checkbox 勾选框。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/checkbox

 * @author wcz0
 * @version 6.2.2
 */
type CheckboxControl struct {
	*BaseRenderer
}

func NewCheckboxControl() *CheckboxControl {
    a := &CheckboxControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "checkbox")
    return a
}


func (a *CheckboxControl) Set(name string, value interface{}) *CheckboxControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * inline
 */
func (a *CheckboxControl) Inline(value interface{}) *CheckboxControl {
    a.Set("inline", value)
    return a
}

/**
 * size
 */
func (a *CheckboxControl) Size(value interface{}) *CheckboxControl {
    a.Set("size", value)
    return a
}

/**
 * labelAlign
 */
func (a *CheckboxControl) LabelAlign(value interface{}) *CheckboxControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * labelOverflow
 */
func (a *CheckboxControl) LabelOverflow(value interface{}) *CheckboxControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * extraName
 */
func (a *CheckboxControl) ExtraName(value interface{}) *CheckboxControl {
    a.Set("extraName", value)
    return a
}

/**
 * inputClassName
 */
func (a *CheckboxControl) InputClassName(value interface{}) *CheckboxControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * validateApi
 */
func (a *CheckboxControl) ValidateApi(value interface{}) *CheckboxControl {
    a.Set("validateApi", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *CheckboxControl) ClearValueOnHidden(value interface{}) *CheckboxControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * partial
 */
func (a *CheckboxControl) Partial(value interface{}) *CheckboxControl {
    a.Set("partial", value)
    return a
}

/**
 * labelClassName
 */
func (a *CheckboxControl) LabelClassName(value interface{}) *CheckboxControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * horizontal
 */
func (a *CheckboxControl) Horizontal(value interface{}) *CheckboxControl {
    a.Set("horizontal", value)
    return a
}

/**
 * disabled
 */
func (a *CheckboxControl) Disabled(value interface{}) *CheckboxControl {
    a.Set("disabled", value)
    return a
}

/**
 * disabledOn
 */
func (a *CheckboxControl) DisabledOn(value interface{}) *CheckboxControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * id
 */
func (a *CheckboxControl) Id(value interface{}) *CheckboxControl {
    a.Set("id", value)
    return a
}

/**
 * useMobileUI
 */
func (a *CheckboxControl) UseMobileUI(value interface{}) *CheckboxControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * testIdBuilder
 */
func (a *CheckboxControl) TestIdBuilder(value interface{}) *CheckboxControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * readOnly
 */
func (a *CheckboxControl) ReadOnly(value interface{}) *CheckboxControl {
    a.Set("readOnly", value)
    return a
}

/**
 * validateOnChange
 */
func (a *CheckboxControl) ValidateOnChange(value interface{}) *CheckboxControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * className
 */
func (a *CheckboxControl) ClassName(value interface{}) *CheckboxControl {
    a.Set("className", value)
    return a
}

/**
 * staticOn
 */
func (a *CheckboxControl) StaticOn(value interface{}) *CheckboxControl {
    a.Set("staticOn", value)
    return a
}

/**
 * staticClassName
 */
func (a *CheckboxControl) StaticClassName(value interface{}) *CheckboxControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *CheckboxControl) StaticSchema(value interface{}) *CheckboxControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *CheckboxControl) Width(value interface{}) *CheckboxControl {
    a.Set("width", value)
    return a
}

/**
 * visible
 */
func (a *CheckboxControl) Visible(value interface{}) *CheckboxControl {
    a.Set("visible", value)
    return a
}

/**
 * hidden
 */
func (a *CheckboxControl) Hidden(value interface{}) *CheckboxControl {
    a.Set("hidden", value)
    return a
}

/**
 * onEvent
 */
func (a *CheckboxControl) OnEvent(value interface{}) *CheckboxControl {
    a.Set("onEvent", value)
    return a
}

/**
 * editorSetting
 */
func (a *CheckboxControl) EditorSetting(value interface{}) *CheckboxControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * labelRemark
 */
func (a *CheckboxControl) LabelRemark(value interface{}) *CheckboxControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * visibleOn
 */
func (a *CheckboxControl) VisibleOn(value interface{}) *CheckboxControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *CheckboxControl) StaticPlaceholder(value interface{}) *CheckboxControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *CheckboxControl) ReadOnlyOn(value interface{}) *CheckboxControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * name
 */
func (a *CheckboxControl) Name(value interface{}) *CheckboxControl {
    a.Set("name", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *CheckboxControl) DescriptionClassName(value interface{}) *CheckboxControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *CheckboxControl) StaticLabelClassName(value interface{}) *CheckboxControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * falseValue
 */
func (a *CheckboxControl) FalseValue(value interface{}) *CheckboxControl {
    a.Set("falseValue", value)
    return a
}

/**
 * option
 */
func (a *CheckboxControl) Option(value interface{}) *CheckboxControl {
    a.Set("option", value)
    return a
}

/**
 * description
 */
func (a *CheckboxControl) Description(value interface{}) *CheckboxControl {
    a.Set("description", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *CheckboxControl) StaticInputClassName(value interface{}) *CheckboxControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * checked
 */
func (a *CheckboxControl) Checked(value interface{}) *CheckboxControl {
    a.Set("checked", value)
    return a
}

/**
 * labelWidth
 */
func (a *CheckboxControl) LabelWidth(value interface{}) *CheckboxControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * submitOnChange
 */
func (a *CheckboxControl) SubmitOnChange(value interface{}) *CheckboxControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * required
 */
func (a *CheckboxControl) Required(value interface{}) *CheckboxControl {
    a.Set("required", value)
    return a
}

/**
 */
func (a *CheckboxControl) Type(value interface{}) *CheckboxControl {
    a.Set("type", value)
    return a
}

/**
 * optionType
 */
func (a *CheckboxControl) OptionType(value interface{}) *CheckboxControl {
    a.Set("optionType", value)
    return a
}

/**
 * label
 */
func (a *CheckboxControl) Label(value interface{}) *CheckboxControl {
    a.Set("label", value)
    return a
}

/**
 * desc
 */
func (a *CheckboxControl) Desc(value interface{}) *CheckboxControl {
    a.Set("desc", value)
    return a
}

/**
 * autoFill
 */
func (a *CheckboxControl) AutoFill(value interface{}) *CheckboxControl {
    a.Set("autoFill", value)
    return a
}

/**
 * trueValue
 */
func (a *CheckboxControl) TrueValue(value interface{}) *CheckboxControl {
    a.Set("trueValue", value)
    return a
}

/**
 * validationErrors
 */
func (a *CheckboxControl) ValidationErrors(value interface{}) *CheckboxControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * validations
 */
func (a *CheckboxControl) Validations(value interface{}) *CheckboxControl {
    a.Set("validations", value)
    return a
}

/**
 * placeholder
 */
func (a *CheckboxControl) Placeholder(value interface{}) *CheckboxControl {
    a.Set("placeholder", value)
    return a
}

/**
 * value
 */
func (a *CheckboxControl) Value(value interface{}) *CheckboxControl {
    a.Set("value", value)
    return a
}

/**
 * initAutoFill
 */
func (a *CheckboxControl) InitAutoFill(value interface{}) *CheckboxControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * hiddenOn
 */
func (a *CheckboxControl) HiddenOn(value interface{}) *CheckboxControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * static
 */
func (a *CheckboxControl) Static(value interface{}) *CheckboxControl {
    a.Set("static", value)
    return a
}

/**
 * style
 */
func (a *CheckboxControl) Style(value interface{}) *CheckboxControl {
    a.Set("style", value)
    return a
}

/**
 * badge
 */
func (a *CheckboxControl) Badge(value interface{}) *CheckboxControl {
    a.Set("badge", value)
    return a
}

/**
 * mode
 */
func (a *CheckboxControl) Mode(value interface{}) *CheckboxControl {
    a.Set("mode", value)
    return a
}

/**
 * row
 */
func (a *CheckboxControl) Row(value interface{}) *CheckboxControl {
    a.Set("row", value)
    return a
}

/**
 * remark
 */
func (a *CheckboxControl) Remark(value interface{}) *CheckboxControl {
    a.Set("remark", value)
    return a
}

/**
 * hint
 */
func (a *CheckboxControl) Hint(value interface{}) *CheckboxControl {
    a.Set("hint", value)
    return a
}
