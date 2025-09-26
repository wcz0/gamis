package renderers


/**
 * TextArea 多行文本输入框。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/textarea

 * @author wcz0
 * @version 6.2.2
 */
type TextareaControl struct {
	*BaseRenderer
}

func NewTextareaControl() *TextareaControl {
    a := &TextareaControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "textarea")
    return a
}


func (a *TextareaControl) Set(name string, value interface{}) *TextareaControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * label
 */
func (a *TextareaControl) Label(value interface{}) *TextareaControl {
    a.Set("label", value)
    return a
}

/**
 * description
 */
func (a *TextareaControl) Description(value interface{}) *TextareaControl {
    a.Set("description", value)
    return a
}

/**
 * autoFill
 */
func (a *TextareaControl) AutoFill(value interface{}) *TextareaControl {
    a.Set("autoFill", value)
    return a
}

/**
 * hidden
 */
func (a *TextareaControl) Hidden(value interface{}) *TextareaControl {
    a.Set("hidden", value)
    return a
}

/**
 * maxRows
 */
func (a *TextareaControl) MaxRows(value interface{}) *TextareaControl {
    a.Set("maxRows", value)
    return a
}

/**
 * name
 */
func (a *TextareaControl) Name(value interface{}) *TextareaControl {
    a.Set("name", value)
    return a
}

/**
 * initAutoFill
 */
func (a *TextareaControl) InitAutoFill(value interface{}) *TextareaControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * static
 */
func (a *TextareaControl) Static(value interface{}) *TextareaControl {
    a.Set("static", value)
    return a
}

/**
 * size
 */
func (a *TextareaControl) Size(value interface{}) *TextareaControl {
    a.Set("size", value)
    return a
}

/**
 * showCounter
 */
func (a *TextareaControl) ShowCounter(value interface{}) *TextareaControl {
    a.Set("showCounter", value)
    return a
}

/**
 * mode
 */
func (a *TextareaControl) Mode(value interface{}) *TextareaControl {
    a.Set("mode", value)
    return a
}

/**
 * inline
 */
func (a *TextareaControl) Inline(value interface{}) *TextareaControl {
    a.Set("inline", value)
    return a
}

/**
 * hiddenOn
 */
func (a *TextareaControl) HiddenOn(value interface{}) *TextareaControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * validateApi
 */
func (a *TextareaControl) ValidateApi(value interface{}) *TextareaControl {
    a.Set("validateApi", value)
    return a
}

/**
 * disabled
 */
func (a *TextareaControl) Disabled(value interface{}) *TextareaControl {
    a.Set("disabled", value)
    return a
}

/**
 * id
 */
func (a *TextareaControl) Id(value interface{}) *TextareaControl {
    a.Set("id", value)
    return a
}

/**
 * resetValue
 */
func (a *TextareaControl) ResetValue(value interface{}) *TextareaControl {
    a.Set("resetValue", value)
    return a
}

/**
 * inputClassName
 */
func (a *TextareaControl) InputClassName(value interface{}) *TextareaControl {
    a.Set("inputClassName", value)
    return a
}

/**
 */
func (a *TextareaControl) Type(value interface{}) *TextareaControl {
    a.Set("type", value)
    return a
}

/**
 * maxLength
 */
func (a *TextareaControl) MaxLength(value interface{}) *TextareaControl {
    a.Set("maxLength", value)
    return a
}

/**
 * labelRemark
 */
func (a *TextareaControl) LabelRemark(value interface{}) *TextareaControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * horizontal
 */
func (a *TextareaControl) Horizontal(value interface{}) *TextareaControl {
    a.Set("horizontal", value)
    return a
}

/**
 * staticSchema
 */
func (a *TextareaControl) StaticSchema(value interface{}) *TextareaControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * labelAlign
 */
func (a *TextareaControl) LabelAlign(value interface{}) *TextareaControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * hint
 */
func (a *TextareaControl) Hint(value interface{}) *TextareaControl {
    a.Set("hint", value)
    return a
}

/**
 * value
 */
func (a *TextareaControl) Value(value interface{}) *TextareaControl {
    a.Set("value", value)
    return a
}

/**
 * borderMode
 */
func (a *TextareaControl) BorderMode(value interface{}) *TextareaControl {
    a.Set("borderMode", value)
    return a
}

/**
 * extraName
 */
func (a *TextareaControl) ExtraName(value interface{}) *TextareaControl {
    a.Set("extraName", value)
    return a
}

/**
 * placeholder
 */
func (a *TextareaControl) Placeholder(value interface{}) *TextareaControl {
    a.Set("placeholder", value)
    return a
}

/**
 * validationErrors
 */
func (a *TextareaControl) ValidationErrors(value interface{}) *TextareaControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * staticOn
 */
func (a *TextareaControl) StaticOn(value interface{}) *TextareaControl {
    a.Set("staticOn", value)
    return a
}

/**
 * staticClassName
 */
func (a *TextareaControl) StaticClassName(value interface{}) *TextareaControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * useMobileUI
 */
func (a *TextareaControl) UseMobileUI(value interface{}) *TextareaControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * validations
 */
func (a *TextareaControl) Validations(value interface{}) *TextareaControl {
    a.Set("validations", value)
    return a
}

/**
 * row
 */
func (a *TextareaControl) Row(value interface{}) *TextareaControl {
    a.Set("row", value)
    return a
}

/**
 * className
 */
func (a *TextareaControl) ClassName(value interface{}) *TextareaControl {
    a.Set("className", value)
    return a
}

/**
 * validateOnChange
 */
func (a *TextareaControl) ValidateOnChange(value interface{}) *TextareaControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * labelOverflow
 */
func (a *TextareaControl) LabelOverflow(value interface{}) *TextareaControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *TextareaControl) StaticPlaceholder(value interface{}) *TextareaControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *TextareaControl) StaticLabelClassName(value interface{}) *TextareaControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * minRows
 */
func (a *TextareaControl) MinRows(value interface{}) *TextareaControl {
    a.Set("minRows", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *TextareaControl) Width(value interface{}) *TextareaControl {
    a.Set("width", value)
    return a
}

/**
 * labelClassName
 */
func (a *TextareaControl) LabelClassName(value interface{}) *TextareaControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *TextareaControl) ReadOnlyOn(value interface{}) *TextareaControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *TextareaControl) DescriptionClassName(value interface{}) *TextareaControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * visible
 */
func (a *TextareaControl) Visible(value interface{}) *TextareaControl {
    a.Set("visible", value)
    return a
}

/**
 * visibleOn
 */
func (a *TextareaControl) VisibleOn(value interface{}) *TextareaControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * onEvent
 */
func (a *TextareaControl) OnEvent(value interface{}) *TextareaControl {
    a.Set("onEvent", value)
    return a
}

/**
 * remark
 */
func (a *TextareaControl) Remark(value interface{}) *TextareaControl {
    a.Set("remark", value)
    return a
}

/**
 * submitOnChange
 */
func (a *TextareaControl) SubmitOnChange(value interface{}) *TextareaControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *TextareaControl) ClearValueOnHidden(value interface{}) *TextareaControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * desc
 */
func (a *TextareaControl) Desc(value interface{}) *TextareaControl {
    a.Set("desc", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *TextareaControl) StaticInputClassName(value interface{}) *TextareaControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * labelWidth
 */
func (a *TextareaControl) LabelWidth(value interface{}) *TextareaControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * readOnly
 */
func (a *TextareaControl) ReadOnly(value interface{}) *TextareaControl {
    a.Set("readOnly", value)
    return a
}

/**
 * required
 */
func (a *TextareaControl) Required(value interface{}) *TextareaControl {
    a.Set("required", value)
    return a
}

/**
 * disabledOn
 */
func (a *TextareaControl) DisabledOn(value interface{}) *TextareaControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * style
 */
func (a *TextareaControl) Style(value interface{}) *TextareaControl {
    a.Set("style", value)
    return a
}

/**
 * editorSetting
 */
func (a *TextareaControl) EditorSetting(value interface{}) *TextareaControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * clearable
 */
func (a *TextareaControl) Clearable(value interface{}) *TextareaControl {
    a.Set("clearable", value)
    return a
}
