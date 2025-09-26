package renderers


/**
 * 图标选择器 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/icon-picker

 * @author wcz0
 * @version 6.2.2
 */
type IconPickerControl struct {
	*BaseRenderer
}

func NewIconPickerControl() *IconPickerControl {
    a := &IconPickerControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "icon-picker")
    return a
}


func (a *IconPickerControl) Set(name string, value interface{}) *IconPickerControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * onEvent
 */
func (a *IconPickerControl) OnEvent(value interface{}) *IconPickerControl {
    a.Set("onEvent", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *IconPickerControl) StaticPlaceholder(value interface{}) *IconPickerControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * labelAlign
 */
func (a *IconPickerControl) LabelAlign(value interface{}) *IconPickerControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * labelWidth
 */
func (a *IconPickerControl) LabelWidth(value interface{}) *IconPickerControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * inline
 */
func (a *IconPickerControl) Inline(value interface{}) *IconPickerControl {
    a.Set("inline", value)
    return a
}

/**
 * placeholder
 */
func (a *IconPickerControl) Placeholder(value interface{}) *IconPickerControl {
    a.Set("placeholder", value)
    return a
}

/**
 * required
 */
func (a *IconPickerControl) Required(value interface{}) *IconPickerControl {
    a.Set("required", value)
    return a
}

/**
 * static
 */
func (a *IconPickerControl) Static(value interface{}) *IconPickerControl {
    a.Set("static", value)
    return a
}

/**
 * staticSchema
 */
func (a *IconPickerControl) StaticSchema(value interface{}) *IconPickerControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * size
 */
func (a *IconPickerControl) Size(value interface{}) *IconPickerControl {
    a.Set("size", value)
    return a
}

/**
 * remark
 */
func (a *IconPickerControl) Remark(value interface{}) *IconPickerControl {
    a.Set("remark", value)
    return a
}

/**
 * submitOnChange
 */
func (a *IconPickerControl) SubmitOnChange(value interface{}) *IconPickerControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * horizontal
 */
func (a *IconPickerControl) Horizontal(value interface{}) *IconPickerControl {
    a.Set("horizontal", value)
    return a
}

/**
 * visible
 */
func (a *IconPickerControl) Visible(value interface{}) *IconPickerControl {
    a.Set("visible", value)
    return a
}

/**
 * visibleOn
 */
func (a *IconPickerControl) VisibleOn(value interface{}) *IconPickerControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticOn
 */
func (a *IconPickerControl) StaticOn(value interface{}) *IconPickerControl {
    a.Set("staticOn", value)
    return a
}

/**
 * style
 */
func (a *IconPickerControl) Style(value interface{}) *IconPickerControl {
    a.Set("style", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *IconPickerControl) Width(value interface{}) *IconPickerControl {
    a.Set("width", value)
    return a
}

/**
 * labelRemark
 */
func (a *IconPickerControl) LabelRemark(value interface{}) *IconPickerControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * inputClassName
 */
func (a *IconPickerControl) InputClassName(value interface{}) *IconPickerControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * value
 */
func (a *IconPickerControl) Value(value interface{}) *IconPickerControl {
    a.Set("value", value)
    return a
}

/**
 * validateApi
 */
func (a *IconPickerControl) ValidateApi(value interface{}) *IconPickerControl {
    a.Set("validateApi", value)
    return a
}

/**
 * row
 */
func (a *IconPickerControl) Row(value interface{}) *IconPickerControl {
    a.Set("row", value)
    return a
}

/**
 * id
 */
func (a *IconPickerControl) Id(value interface{}) *IconPickerControl {
    a.Set("id", value)
    return a
}

/**
 * validateOnChange
 */
func (a *IconPickerControl) ValidateOnChange(value interface{}) *IconPickerControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *IconPickerControl) DescriptionClassName(value interface{}) *IconPickerControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * validations
 */
func (a *IconPickerControl) Validations(value interface{}) *IconPickerControl {
    a.Set("validations", value)
    return a
}

/**
 * initAutoFill
 */
func (a *IconPickerControl) InitAutoFill(value interface{}) *IconPickerControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * disabledOn
 */
func (a *IconPickerControl) DisabledOn(value interface{}) *IconPickerControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *IconPickerControl) StaticInputClassName(value interface{}) *IconPickerControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *IconPickerControl) Type(value interface{}) *IconPickerControl {
    a.Set("type", value)
    return a
}

/**
 * extraName
 */
func (a *IconPickerControl) ExtraName(value interface{}) *IconPickerControl {
    a.Set("extraName", value)
    return a
}

/**
 * desc
 */
func (a *IconPickerControl) Desc(value interface{}) *IconPickerControl {
    a.Set("desc", value)
    return a
}

/**
 * validationErrors
 */
func (a *IconPickerControl) ValidationErrors(value interface{}) *IconPickerControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *IconPickerControl) ClearValueOnHidden(value interface{}) *IconPickerControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * className
 */
func (a *IconPickerControl) ClassName(value interface{}) *IconPickerControl {
    a.Set("className", value)
    return a
}

/**
 * staticClassName
 */
func (a *IconPickerControl) StaticClassName(value interface{}) *IconPickerControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * editorSetting
 */
func (a *IconPickerControl) EditorSetting(value interface{}) *IconPickerControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * readOnly
 */
func (a *IconPickerControl) ReadOnly(value interface{}) *IconPickerControl {
    a.Set("readOnly", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *IconPickerControl) ReadOnlyOn(value interface{}) *IconPickerControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * description
 */
func (a *IconPickerControl) Description(value interface{}) *IconPickerControl {
    a.Set("description", value)
    return a
}

/**
 * hiddenOn
 */
func (a *IconPickerControl) HiddenOn(value interface{}) *IconPickerControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *IconPickerControl) StaticLabelClassName(value interface{}) *IconPickerControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * labelOverflow
 */
func (a *IconPickerControl) LabelOverflow(value interface{}) *IconPickerControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * hint
 */
func (a *IconPickerControl) Hint(value interface{}) *IconPickerControl {
    a.Set("hint", value)
    return a
}

/**
 * disabled
 */
func (a *IconPickerControl) Disabled(value interface{}) *IconPickerControl {
    a.Set("disabled", value)
    return a
}

/**
 * useMobileUI
 */
func (a *IconPickerControl) UseMobileUI(value interface{}) *IconPickerControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * label
 */
func (a *IconPickerControl) Label(value interface{}) *IconPickerControl {
    a.Set("label", value)
    return a
}

/**
 * labelClassName
 */
func (a *IconPickerControl) LabelClassName(value interface{}) *IconPickerControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * name
 */
func (a *IconPickerControl) Name(value interface{}) *IconPickerControl {
    a.Set("name", value)
    return a
}

/**
 * mode
 */
func (a *IconPickerControl) Mode(value interface{}) *IconPickerControl {
    a.Set("mode", value)
    return a
}

/**
 * autoFill
 */
func (a *IconPickerControl) AutoFill(value interface{}) *IconPickerControl {
    a.Set("autoFill", value)
    return a
}

/**
 * hidden
 */
func (a *IconPickerControl) Hidden(value interface{}) *IconPickerControl {
    a.Set("hidden", value)
    return a
}
