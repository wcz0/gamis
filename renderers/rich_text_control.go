package renderers


/**
 * RichText 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/input-rich-text

 * @author wcz0
 * @version 6.2.2
 */
type RichTextControl struct {
	*BaseRenderer
}

func NewRichTextControl() *RichTextControl {
    a := &RichTextControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-rich-text")
    return a
}


func (a *RichTextControl) Set(name string, value interface{}) *RichTextControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * validations
 */
func (a *RichTextControl) Validations(value interface{}) *RichTextControl {
    a.Set("validations", value)
    return a
}

/**
 * visibleOn
 */
func (a *RichTextControl) VisibleOn(value interface{}) *RichTextControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * onEvent
 */
func (a *RichTextControl) OnEvent(value interface{}) *RichTextControl {
    a.Set("onEvent", value)
    return a
}

/**
 * value
 */
func (a *RichTextControl) Value(value interface{}) *RichTextControl {
    a.Set("value", value)
    return a
}

/**
 * size
 */
func (a *RichTextControl) Size(value interface{}) *RichTextControl {
    a.Set("size", value)
    return a
}

/**
 * borderMode
 */
func (a *RichTextControl) BorderMode(value interface{}) *RichTextControl {
    a.Set("borderMode", value)
    return a
}

/**
 * validateOnChange
 */
func (a *RichTextControl) ValidateOnChange(value interface{}) *RichTextControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * validationErrors
 */
func (a *RichTextControl) ValidationErrors(value interface{}) *RichTextControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * id
 */
func (a *RichTextControl) Id(value interface{}) *RichTextControl {
    a.Set("id", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *RichTextControl) StaticPlaceholder(value interface{}) *RichTextControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticClassName
 */
func (a *RichTextControl) StaticClassName(value interface{}) *RichTextControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *RichTextControl) StaticInputClassName(value interface{}) *RichTextControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * style
 */
func (a *RichTextControl) Style(value interface{}) *RichTextControl {
    a.Set("style", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *RichTextControl) ClearValueOnHidden(value interface{}) *RichTextControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * row
 */
func (a *RichTextControl) Row(value interface{}) *RichTextControl {
    a.Set("row", value)
    return a
}

/**
 * disabled
 */
func (a *RichTextControl) Disabled(value interface{}) *RichTextControl {
    a.Set("disabled", value)
    return a
}

/**
 * labelRemark
 */
func (a *RichTextControl) LabelRemark(value interface{}) *RichTextControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * inline
 */
func (a *RichTextControl) Inline(value interface{}) *RichTextControl {
    a.Set("inline", value)
    return a
}

/**
 * required
 */
func (a *RichTextControl) Required(value interface{}) *RichTextControl {
    a.Set("required", value)
    return a
}

/**
 * remark
 */
func (a *RichTextControl) Remark(value interface{}) *RichTextControl {
    a.Set("remark", value)
    return a
}

/**
 * hidden
 */
func (a *RichTextControl) Hidden(value interface{}) *RichTextControl {
    a.Set("hidden", value)
    return a
}

/**
 * staticSchema
 */
func (a *RichTextControl) StaticSchema(value interface{}) *RichTextControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * fileField
 */
func (a *RichTextControl) FileField(value interface{}) *RichTextControl {
    a.Set("fileField", value)
    return a
}

/**
 * labelAlign
 */
func (a *RichTextControl) LabelAlign(value interface{}) *RichTextControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * staticOn
 */
func (a *RichTextControl) StaticOn(value interface{}) *RichTextControl {
    a.Set("staticOn", value)
    return a
}

/**
 * disabledOn
 */
func (a *RichTextControl) DisabledOn(value interface{}) *RichTextControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * receiver
 */
func (a *RichTextControl) Receiver(value interface{}) *RichTextControl {
    a.Set("receiver", value)
    return a
}

/**
 * submitOnChange
 */
func (a *RichTextControl) SubmitOnChange(value interface{}) *RichTextControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * inputClassName
 */
func (a *RichTextControl) InputClassName(value interface{}) *RichTextControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * static
 */
func (a *RichTextControl) Static(value interface{}) *RichTextControl {
    a.Set("static", value)
    return a
}

/**
 * labelOverflow
 */
func (a *RichTextControl) LabelOverflow(value interface{}) *RichTextControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * mode
 */
func (a *RichTextControl) Mode(value interface{}) *RichTextControl {
    a.Set("mode", value)
    return a
}

/**
 * className
 */
func (a *RichTextControl) ClassName(value interface{}) *RichTextControl {
    a.Set("className", value)
    return a
}

/**
 * visible
 */
func (a *RichTextControl) Visible(value interface{}) *RichTextControl {
    a.Set("visible", value)
    return a
}

/**
 * useMobileUI
 */
func (a *RichTextControl) UseMobileUI(value interface{}) *RichTextControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *RichTextControl) DescriptionClassName(value interface{}) *RichTextControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *RichTextControl) StaticLabelClassName(value interface{}) *RichTextControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * videoReceiver
 */
func (a *RichTextControl) VideoReceiver(value interface{}) *RichTextControl {
    a.Set("videoReceiver", value)
    return a
}

/**
 * labelWidth
 */
func (a *RichTextControl) LabelWidth(value interface{}) *RichTextControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * extraName
 */
func (a *RichTextControl) ExtraName(value interface{}) *RichTextControl {
    a.Set("extraName", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *RichTextControl) ReadOnlyOn(value interface{}) *RichTextControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * hint
 */
func (a *RichTextControl) Hint(value interface{}) *RichTextControl {
    a.Set("hint", value)
    return a
}

/**
 * readOnly
 */
func (a *RichTextControl) ReadOnly(value interface{}) *RichTextControl {
    a.Set("readOnly", value)
    return a
}

/**
 * autoFill
 */
func (a *RichTextControl) AutoFill(value interface{}) *RichTextControl {
    a.Set("autoFill", value)
    return a
}

/**
 * hiddenOn
 */
func (a *RichTextControl) HiddenOn(value interface{}) *RichTextControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 */
func (a *RichTextControl) Type(value interface{}) *RichTextControl {
    a.Set("type", value)
    return a
}

/**
 * vendor
 */
func (a *RichTextControl) Vendor(value interface{}) *RichTextControl {
    a.Set("vendor", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *RichTextControl) Width(value interface{}) *RichTextControl {
    a.Set("width", value)
    return a
}

/**
 * validateApi
 */
func (a *RichTextControl) ValidateApi(value interface{}) *RichTextControl {
    a.Set("validateApi", value)
    return a
}

/**
 * description
 */
func (a *RichTextControl) Description(value interface{}) *RichTextControl {
    a.Set("description", value)
    return a
}

/**
 * editorSetting
 */
func (a *RichTextControl) EditorSetting(value interface{}) *RichTextControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * labelClassName
 */
func (a *RichTextControl) LabelClassName(value interface{}) *RichTextControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * name
 */
func (a *RichTextControl) Name(value interface{}) *RichTextControl {
    a.Set("name", value)
    return a
}

/**
 * initAutoFill
 */
func (a *RichTextControl) InitAutoFill(value interface{}) *RichTextControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * placeholder
 */
func (a *RichTextControl) Placeholder(value interface{}) *RichTextControl {
    a.Set("placeholder", value)
    return a
}

/**
 * options
 */
func (a *RichTextControl) Options(value interface{}) *RichTextControl {
    a.Set("options", value)
    return a
}

/**
 * label
 */
func (a *RichTextControl) Label(value interface{}) *RichTextControl {
    a.Set("label", value)
    return a
}

/**
 * desc
 */
func (a *RichTextControl) Desc(value interface{}) *RichTextControl {
    a.Set("desc", value)
    return a
}

/**
 * horizontal
 */
func (a *RichTextControl) Horizontal(value interface{}) *RichTextControl {
    a.Set("horizontal", value)
    return a
}
