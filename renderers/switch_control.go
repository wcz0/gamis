package renderers


/**
 * Switch 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/switch

 * @author wcz0
 * @version 6.2.2
 */
type SwitchControl struct {
	*BaseRenderer
}

func NewSwitchControl() *SwitchControl {
    a := &SwitchControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "switch")
    return a
}


func (a *SwitchControl) Set(name string, value interface{}) *SwitchControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * offText
 */
func (a *SwitchControl) OffText(value interface{}) *SwitchControl {
    a.Set("offText", value)
    return a
}

/**
 * labelClassName
 */
func (a *SwitchControl) LabelClassName(value interface{}) *SwitchControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * id
 */
func (a *SwitchControl) Id(value interface{}) *SwitchControl {
    a.Set("id", value)
    return a
}

/**
 * onEvent
 */
func (a *SwitchControl) OnEvent(value interface{}) *SwitchControl {
    a.Set("onEvent", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *SwitchControl) StaticInputClassName(value interface{}) *SwitchControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * mode
 */
func (a *SwitchControl) Mode(value interface{}) *SwitchControl {
    a.Set("mode", value)
    return a
}

/**
 * required
 */
func (a *SwitchControl) Required(value interface{}) *SwitchControl {
    a.Set("required", value)
    return a
}

/**
 * validations
 */
func (a *SwitchControl) Validations(value interface{}) *SwitchControl {
    a.Set("validations", value)
    return a
}

/**
 * option
 */
func (a *SwitchControl) Option(value interface{}) *SwitchControl {
    a.Set("option", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *SwitchControl) ClearValueOnHidden(value interface{}) *SwitchControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * className
 */
func (a *SwitchControl) ClassName(value interface{}) *SwitchControl {
    a.Set("className", value)
    return a
}

/**
 * disabled
 */
func (a *SwitchControl) Disabled(value interface{}) *SwitchControl {
    a.Set("disabled", value)
    return a
}

/**
 * hidden
 */
func (a *SwitchControl) Hidden(value interface{}) *SwitchControl {
    a.Set("hidden", value)
    return a
}

/**
 * visibleOn
 */
func (a *SwitchControl) VisibleOn(value interface{}) *SwitchControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * static
 */
func (a *SwitchControl) Static(value interface{}) *SwitchControl {
    a.Set("static", value)
    return a
}

/**
 * labelWidth
 */
func (a *SwitchControl) LabelWidth(value interface{}) *SwitchControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * style
 */
func (a *SwitchControl) Style(value interface{}) *SwitchControl {
    a.Set("style", value)
    return a
}

/**
 * editorSetting
 */
func (a *SwitchControl) EditorSetting(value interface{}) *SwitchControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * labelAlign
 */
func (a *SwitchControl) LabelAlign(value interface{}) *SwitchControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *SwitchControl) StaticPlaceholder(value interface{}) *SwitchControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *SwitchControl) StaticLabelClassName(value interface{}) *SwitchControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * autoFill
 */
func (a *SwitchControl) AutoFill(value interface{}) *SwitchControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *SwitchControl) Width(value interface{}) *SwitchControl {
    a.Set("width", value)
    return a
}

/**
 * value
 */
func (a *SwitchControl) Value(value interface{}) *SwitchControl {
    a.Set("value", value)
    return a
}

/**
 * extraName
 */
func (a *SwitchControl) ExtraName(value interface{}) *SwitchControl {
    a.Set("extraName", value)
    return a
}

/**
 * inline
 */
func (a *SwitchControl) Inline(value interface{}) *SwitchControl {
    a.Set("inline", value)
    return a
}

/**
 * visible
 */
func (a *SwitchControl) Visible(value interface{}) *SwitchControl {
    a.Set("visible", value)
    return a
}

/**
 * useMobileUI
 */
func (a *SwitchControl) UseMobileUI(value interface{}) *SwitchControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * readOnly
 */
func (a *SwitchControl) ReadOnly(value interface{}) *SwitchControl {
    a.Set("readOnly", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *SwitchControl) DescriptionClassName(value interface{}) *SwitchControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *SwitchControl) StaticSchema(value interface{}) *SwitchControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * desc
 */
func (a *SwitchControl) Desc(value interface{}) *SwitchControl {
    a.Set("desc", value)
    return a
}

/**
 * staticOn
 */
func (a *SwitchControl) StaticOn(value interface{}) *SwitchControl {
    a.Set("staticOn", value)
    return a
}

/**
 * falseValue
 */
func (a *SwitchControl) FalseValue(value interface{}) *SwitchControl {
    a.Set("falseValue", value)
    return a
}

/**
 * label
 */
func (a *SwitchControl) Label(value interface{}) *SwitchControl {
    a.Set("label", value)
    return a
}

/**
 * horizontal
 */
func (a *SwitchControl) Horizontal(value interface{}) *SwitchControl {
    a.Set("horizontal", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *SwitchControl) ReadOnlyOn(value interface{}) *SwitchControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * validationErrors
 */
func (a *SwitchControl) ValidationErrors(value interface{}) *SwitchControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * initAutoFill
 */
func (a *SwitchControl) InitAutoFill(value interface{}) *SwitchControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * staticClassName
 */
func (a *SwitchControl) StaticClassName(value interface{}) *SwitchControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * size
 */
func (a *SwitchControl) Size(value interface{}) *SwitchControl {
    a.Set("size", value)
    return a
}

/**
 * hint
 */
func (a *SwitchControl) Hint(value interface{}) *SwitchControl {
    a.Set("hint", value)
    return a
}

/**
 * submitOnChange
 */
func (a *SwitchControl) SubmitOnChange(value interface{}) *SwitchControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * validateOnChange
 */
func (a *SwitchControl) ValidateOnChange(value interface{}) *SwitchControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * disabledOn
 */
func (a *SwitchControl) DisabledOn(value interface{}) *SwitchControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * inputClassName
 */
func (a *SwitchControl) InputClassName(value interface{}) *SwitchControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * remark
 */
func (a *SwitchControl) Remark(value interface{}) *SwitchControl {
    a.Set("remark", value)
    return a
}

/**
 * labelRemark
 */
func (a *SwitchControl) LabelRemark(value interface{}) *SwitchControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * description
 */
func (a *SwitchControl) Description(value interface{}) *SwitchControl {
    a.Set("description", value)
    return a
}

/**
 * validateApi
 */
func (a *SwitchControl) ValidateApi(value interface{}) *SwitchControl {
    a.Set("validateApi", value)
    return a
}

/**
 * row
 */
func (a *SwitchControl) Row(value interface{}) *SwitchControl {
    a.Set("row", value)
    return a
}

/**
 */
func (a *SwitchControl) Type(value interface{}) *SwitchControl {
    a.Set("type", value)
    return a
}

/**
 * loading
 */
func (a *SwitchControl) Loading(value interface{}) *SwitchControl {
    a.Set("loading", value)
    return a
}

/**
 * trueValue
 */
func (a *SwitchControl) TrueValue(value interface{}) *SwitchControl {
    a.Set("trueValue", value)
    return a
}

/**
 * labelOverflow
 */
func (a *SwitchControl) LabelOverflow(value interface{}) *SwitchControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * name
 */
func (a *SwitchControl) Name(value interface{}) *SwitchControl {
    a.Set("name", value)
    return a
}

/**
 * placeholder
 */
func (a *SwitchControl) Placeholder(value interface{}) *SwitchControl {
    a.Set("placeholder", value)
    return a
}

/**
 * hiddenOn
 */
func (a *SwitchControl) HiddenOn(value interface{}) *SwitchControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * onText
 */
func (a *SwitchControl) OnText(value interface{}) *SwitchControl {
    a.Set("onText", value)
    return a
}
