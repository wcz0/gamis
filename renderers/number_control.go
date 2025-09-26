package renderers


/**
 * 数字输入框 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/input-number

 * @author wcz0
 * @version 6.2.2
 */
type NumberControl struct {
	*BaseRenderer
}

func NewNumberControl() *NumberControl {
    a := &NumberControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-number")
    return a
}


func (a *NumberControl) Set(name string, value interface{}) *NumberControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * borderMode
 */
func (a *NumberControl) BorderMode(value interface{}) *NumberControl {
    a.Set("borderMode", value)
    return a
}

/**
 * description
 */
func (a *NumberControl) Description(value interface{}) *NumberControl {
    a.Set("description", value)
    return a
}

/**
 * required
 */
func (a *NumberControl) Required(value interface{}) *NumberControl {
    a.Set("required", value)
    return a
}

/**
 * submitOnChange
 */
func (a *NumberControl) SubmitOnChange(value interface{}) *NumberControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * mode
 */
func (a *NumberControl) Mode(value interface{}) *NumberControl {
    a.Set("mode", value)
    return a
}

/**
 * hidden
 */
func (a *NumberControl) Hidden(value interface{}) *NumberControl {
    a.Set("hidden", value)
    return a
}

/**
 * static
 */
func (a *NumberControl) Static(value interface{}) *NumberControl {
    a.Set("static", value)
    return a
}

/**
 * editorSetting
 */
func (a *NumberControl) EditorSetting(value interface{}) *NumberControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * step
 */
func (a *NumberControl) Step(value interface{}) *NumberControl {
    a.Set("step", value)
    return a
}

/**
 * unitOptions
 */
func (a *NumberControl) UnitOptions(value interface{}) *NumberControl {
    a.Set("unitOptions", value)
    return a
}

/**
 * keyboard
 */
func (a *NumberControl) Keyboard(value interface{}) *NumberControl {
    a.Set("keyboard", value)
    return a
}

/**
 * label
 */
func (a *NumberControl) Label(value interface{}) *NumberControl {
    a.Set("label", value)
    return a
}

/**
 * disabled
 */
func (a *NumberControl) Disabled(value interface{}) *NumberControl {
    a.Set("disabled", value)
    return a
}

/**
 * labelOverflow
 */
func (a *NumberControl) LabelOverflow(value interface{}) *NumberControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * remark
 */
func (a *NumberControl) Remark(value interface{}) *NumberControl {
    a.Set("remark", value)
    return a
}

/**
 * desc
 */
func (a *NumberControl) Desc(value interface{}) *NumberControl {
    a.Set("desc", value)
    return a
}

/**
 * validateApi
 */
func (a *NumberControl) ValidateApi(value interface{}) *NumberControl {
    a.Set("validateApi", value)
    return a
}

/**
 * visibleOn
 */
func (a *NumberControl) VisibleOn(value interface{}) *NumberControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * onEvent
 */
func (a *NumberControl) OnEvent(value interface{}) *NumberControl {
    a.Set("onEvent", value)
    return a
}

/**
 * staticClassName
 */
func (a *NumberControl) StaticClassName(value interface{}) *NumberControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *NumberControl) StaticSchema(value interface{}) *NumberControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * validateOnChange
 */
func (a *NumberControl) ValidateOnChange(value interface{}) *NumberControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * className
 */
func (a *NumberControl) ClassName(value interface{}) *NumberControl {
    a.Set("className", value)
    return a
}

/**
 * useMobileUI
 */
func (a *NumberControl) UseMobileUI(value interface{}) *NumberControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *NumberControl) Type(value interface{}) *NumberControl {
    a.Set("type", value)
    return a
}

/**
 * showSteps
 */
func (a *NumberControl) ShowSteps(value interface{}) *NumberControl {
    a.Set("showSteps", value)
    return a
}

/**
 * displayMode
 */
func (a *NumberControl) DisplayMode(value interface{}) *NumberControl {
    a.Set("displayMode", value)
    return a
}

/**
 * value
 */
func (a *NumberControl) Value(value interface{}) *NumberControl {
    a.Set("value", value)
    return a
}

/**
 * precision
 */
func (a *NumberControl) Precision(value interface{}) *NumberControl {
    a.Set("precision", value)
    return a
}

/**
 * name
 */
func (a *NumberControl) Name(value interface{}) *NumberControl {
    a.Set("name", value)
    return a
}

/**
 * extraName
 */
func (a *NumberControl) ExtraName(value interface{}) *NumberControl {
    a.Set("extraName", value)
    return a
}

/**
 * placeholder
 */
func (a *NumberControl) Placeholder(value interface{}) *NumberControl {
    a.Set("placeholder", value)
    return a
}

/**
 * kilobitSeparator
 */
func (a *NumberControl) KilobitSeparator(value interface{}) *NumberControl {
    a.Set("kilobitSeparator", value)
    return a
}

/**
 * disabledOn
 */
func (a *NumberControl) DisabledOn(value interface{}) *NumberControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * readOnly
 */
func (a *NumberControl) ReadOnly(value interface{}) *NumberControl {
    a.Set("readOnly", value)
    return a
}

/**
 * size
 */
func (a *NumberControl) Size(value interface{}) *NumberControl {
    a.Set("size", value)
    return a
}

/**
 * big
 */
func (a *NumberControl) Big(value interface{}) *NumberControl {
    a.Set("big", value)
    return a
}

/**
 * labelRemark
 */
func (a *NumberControl) LabelRemark(value interface{}) *NumberControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *NumberControl) StaticLabelClassName(value interface{}) *NumberControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * style
 */
func (a *NumberControl) Style(value interface{}) *NumberControl {
    a.Set("style", value)
    return a
}

/**
 * min
 */
func (a *NumberControl) Min(value interface{}) *NumberControl {
    a.Set("min", value)
    return a
}

/**
 * prefix
 */
func (a *NumberControl) Prefix(value interface{}) *NumberControl {
    a.Set("prefix", value)
    return a
}

/**
 * autoFill
 */
func (a *NumberControl) AutoFill(value interface{}) *NumberControl {
    a.Set("autoFill", value)
    return a
}

/**
 * visible
 */
func (a *NumberControl) Visible(value interface{}) *NumberControl {
    a.Set("visible", value)
    return a
}

/**
 * inputClassName
 */
func (a *NumberControl) InputClassName(value interface{}) *NumberControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * labelWidth
 */
func (a *NumberControl) LabelWidth(value interface{}) *NumberControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * initAutoFill
 */
func (a *NumberControl) InitAutoFill(value interface{}) *NumberControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * row
 */
func (a *NumberControl) Row(value interface{}) *NumberControl {
    a.Set("row", value)
    return a
}

/**
 * hiddenOn
 */
func (a *NumberControl) HiddenOn(value interface{}) *NumberControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * max
 */
func (a *NumberControl) Max(value interface{}) *NumberControl {
    a.Set("max", value)
    return a
}

/**
 * labelClassName
 */
func (a *NumberControl) LabelClassName(value interface{}) *NumberControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *NumberControl) ClearValueOnHidden(value interface{}) *NumberControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * showAsPercent
 */
func (a *NumberControl) ShowAsPercent(value interface{}) *NumberControl {
    a.Set("showAsPercent", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *NumberControl) Width(value interface{}) *NumberControl {
    a.Set("width", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *NumberControl) StaticInputClassName(value interface{}) *NumberControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *NumberControl) ReadOnlyOn(value interface{}) *NumberControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *NumberControl) DescriptionClassName(value interface{}) *NumberControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * horizontal
 */
func (a *NumberControl) Horizontal(value interface{}) *NumberControl {
    a.Set("horizontal", value)
    return a
}

/**
 * validations
 */
func (a *NumberControl) Validations(value interface{}) *NumberControl {
    a.Set("validations", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *NumberControl) StaticPlaceholder(value interface{}) *NumberControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * suffix
 */
func (a *NumberControl) Suffix(value interface{}) *NumberControl {
    a.Set("suffix", value)
    return a
}

/**
 * labelAlign
 */
func (a *NumberControl) LabelAlign(value interface{}) *NumberControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * hint
 */
func (a *NumberControl) Hint(value interface{}) *NumberControl {
    a.Set("hint", value)
    return a
}

/**
 * inline
 */
func (a *NumberControl) Inline(value interface{}) *NumberControl {
    a.Set("inline", value)
    return a
}

/**
 * validationErrors
 */
func (a *NumberControl) ValidationErrors(value interface{}) *NumberControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * id
 */
func (a *NumberControl) Id(value interface{}) *NumberControl {
    a.Set("id", value)
    return a
}

/**
 * staticOn
 */
func (a *NumberControl) StaticOn(value interface{}) *NumberControl {
    a.Set("staticOn", value)
    return a
}
