package renderers


/**
 * InputGroup 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/input-group

 * @author wcz0
 * @version 6.2.2
 */
type InputGroupControl struct {
	*BaseRenderer
}

func NewInputGroupControl() *InputGroupControl {
    a := &InputGroupControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-group")
    return a
}


func (a *InputGroupControl) Set(name string, value interface{}) *InputGroupControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * useMobileUI
 */
func (a *InputGroupControl) UseMobileUI(value interface{}) *InputGroupControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * hint
 */
func (a *InputGroupControl) Hint(value interface{}) *InputGroupControl {
    a.Set("hint", value)
    return a
}

/**
 * submitOnChange
 */
func (a *InputGroupControl) SubmitOnChange(value interface{}) *InputGroupControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * horizontal
 */
func (a *InputGroupControl) Horizontal(value interface{}) *InputGroupControl {
    a.Set("horizontal", value)
    return a
}

/**
 * validationConfig
 */
func (a *InputGroupControl) ValidationConfig(value interface{}) *InputGroupControl {
    a.Set("validationConfig", value)
    return a
}

/**
 * autoFill
 */
func (a *InputGroupControl) AutoFill(value interface{}) *InputGroupControl {
    a.Set("autoFill", value)
    return a
}

/**
 * disabledOn
 */
func (a *InputGroupControl) DisabledOn(value interface{}) *InputGroupControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * labelClassName
 */
func (a *InputGroupControl) LabelClassName(value interface{}) *InputGroupControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *InputGroupControl) StaticLabelClassName(value interface{}) *InputGroupControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * editorSetting
 */
func (a *InputGroupControl) EditorSetting(value interface{}) *InputGroupControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * labelWidth
 */
func (a *InputGroupControl) LabelWidth(value interface{}) *InputGroupControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * validateOnChange
 */
func (a *InputGroupControl) ValidateOnChange(value interface{}) *InputGroupControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * validationErrors
 */
func (a *InputGroupControl) ValidationErrors(value interface{}) *InputGroupControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * visible
 */
func (a *InputGroupControl) Visible(value interface{}) *InputGroupControl {
    a.Set("visible", value)
    return a
}

/**
 * size
 */
func (a *InputGroupControl) Size(value interface{}) *InputGroupControl {
    a.Set("size", value)
    return a
}

/**
 * validations
 */
func (a *InputGroupControl) Validations(value interface{}) *InputGroupControl {
    a.Set("validations", value)
    return a
}

/**
 * row
 */
func (a *InputGroupControl) Row(value interface{}) *InputGroupControl {
    a.Set("row", value)
    return a
}

/**
 * remark
 */
func (a *InputGroupControl) Remark(value interface{}) *InputGroupControl {
    a.Set("remark", value)
    return a
}

/**
 * readOnly
 */
func (a *InputGroupControl) ReadOnly(value interface{}) *InputGroupControl {
    a.Set("readOnly", value)
    return a
}

/**
 * description
 */
func (a *InputGroupControl) Description(value interface{}) *InputGroupControl {
    a.Set("description", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *InputGroupControl) ClearValueOnHidden(value interface{}) *InputGroupControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *InputGroupControl) Width(value interface{}) *InputGroupControl {
    a.Set("width", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *InputGroupControl) ReadOnlyOn(value interface{}) *InputGroupControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * placeholder
 */
func (a *InputGroupControl) Placeholder(value interface{}) *InputGroupControl {
    a.Set("placeholder", value)
    return a
}

/**
 * value
 */
func (a *InputGroupControl) Value(value interface{}) *InputGroupControl {
    a.Set("value", value)
    return a
}

/**
 * labelRemark
 */
func (a *InputGroupControl) LabelRemark(value interface{}) *InputGroupControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * body
 */
func (a *InputGroupControl) Body(value interface{}) *InputGroupControl {
    a.Set("body", value)
    return a
}

/**
 * validateApi
 */
func (a *InputGroupControl) ValidateApi(value interface{}) *InputGroupControl {
    a.Set("validateApi", value)
    return a
}

/**
 * visibleOn
 */
func (a *InputGroupControl) VisibleOn(value interface{}) *InputGroupControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * id
 */
func (a *InputGroupControl) Id(value interface{}) *InputGroupControl {
    a.Set("id", value)
    return a
}

/**
 * onEvent
 */
func (a *InputGroupControl) OnEvent(value interface{}) *InputGroupControl {
    a.Set("onEvent", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *InputGroupControl) StaticPlaceholder(value interface{}) *InputGroupControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *InputGroupControl) StaticInputClassName(value interface{}) *InputGroupControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *InputGroupControl) Type(value interface{}) *InputGroupControl {
    a.Set("type", value)
    return a
}

/**
 * labelOverflow
 */
func (a *InputGroupControl) LabelOverflow(value interface{}) *InputGroupControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * desc
 */
func (a *InputGroupControl) Desc(value interface{}) *InputGroupControl {
    a.Set("desc", value)
    return a
}

/**
 * required
 */
func (a *InputGroupControl) Required(value interface{}) *InputGroupControl {
    a.Set("required", value)
    return a
}

/**
 * static
 */
func (a *InputGroupControl) Static(value interface{}) *InputGroupControl {
    a.Set("static", value)
    return a
}

/**
 * name
 */
func (a *InputGroupControl) Name(value interface{}) *InputGroupControl {
    a.Set("name", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *InputGroupControl) DescriptionClassName(value interface{}) *InputGroupControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * initAutoFill
 */
func (a *InputGroupControl) InitAutoFill(value interface{}) *InputGroupControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * disabled
 */
func (a *InputGroupControl) Disabled(value interface{}) *InputGroupControl {
    a.Set("disabled", value)
    return a
}

/**
 * hiddenOn
 */
func (a *InputGroupControl) HiddenOn(value interface{}) *InputGroupControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * label
 */
func (a *InputGroupControl) Label(value interface{}) *InputGroupControl {
    a.Set("label", value)
    return a
}

/**
 * inline
 */
func (a *InputGroupControl) Inline(value interface{}) *InputGroupControl {
    a.Set("inline", value)
    return a
}

/**
 * staticOn
 */
func (a *InputGroupControl) StaticOn(value interface{}) *InputGroupControl {
    a.Set("staticOn", value)
    return a
}

/**
 * staticClassName
 */
func (a *InputGroupControl) StaticClassName(value interface{}) *InputGroupControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * style
 */
func (a *InputGroupControl) Style(value interface{}) *InputGroupControl {
    a.Set("style", value)
    return a
}

/**
 * extraName
 */
func (a *InputGroupControl) ExtraName(value interface{}) *InputGroupControl {
    a.Set("extraName", value)
    return a
}

/**
 * className
 */
func (a *InputGroupControl) ClassName(value interface{}) *InputGroupControl {
    a.Set("className", value)
    return a
}

/**
 * staticSchema
 */
func (a *InputGroupControl) StaticSchema(value interface{}) *InputGroupControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * labelAlign
 */
func (a *InputGroupControl) LabelAlign(value interface{}) *InputGroupControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * mode
 */
func (a *InputGroupControl) Mode(value interface{}) *InputGroupControl {
    a.Set("mode", value)
    return a
}

/**
 * inputClassName
 */
func (a *InputGroupControl) InputClassName(value interface{}) *InputGroupControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * hidden
 */
func (a *InputGroupControl) Hidden(value interface{}) *InputGroupControl {
    a.Set("hidden", value)
    return a
}
