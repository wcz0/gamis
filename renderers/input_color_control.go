package renderers


/**
 * Color 颜色选择框 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/color

 * @author wcz0
 * @version 6.2.2
 */
type InputColorControl struct {
	*BaseRenderer
}

func NewInputColorControl() *InputColorControl {
    a := &InputColorControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-color")
    return a
}


func (a *InputColorControl) Set(name string, value interface{}) *InputColorControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * labelAlign
 */
func (a *InputColorControl) LabelAlign(value interface{}) *InputColorControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * hint
 */
func (a *InputColorControl) Hint(value interface{}) *InputColorControl {
    a.Set("hint", value)
    return a
}

/**
 * placeholder
 */
func (a *InputColorControl) Placeholder(value interface{}) *InputColorControl {
    a.Set("placeholder", value)
    return a
}

/**
 * validateApi
 */
func (a *InputColorControl) ValidateApi(value interface{}) *InputColorControl {
    a.Set("validateApi", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *InputColorControl) StaticLabelClassName(value interface{}) *InputColorControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * editorSetting
 */
func (a *InputColorControl) EditorSetting(value interface{}) *InputColorControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * inputClassName
 */
func (a *InputColorControl) InputClassName(value interface{}) *InputColorControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * remark
 */
func (a *InputColorControl) Remark(value interface{}) *InputColorControl {
    a.Set("remark", value)
    return a
}

/**
 * submitOnChange
 */
func (a *InputColorControl) SubmitOnChange(value interface{}) *InputColorControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * horizontal
 */
func (a *InputColorControl) Horizontal(value interface{}) *InputColorControl {
    a.Set("horizontal", value)
    return a
}

/**
 * inline
 */
func (a *InputColorControl) Inline(value interface{}) *InputColorControl {
    a.Set("inline", value)
    return a
}

/**
 * hiddenOn
 */
func (a *InputColorControl) HiddenOn(value interface{}) *InputColorControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * id
 */
func (a *InputColorControl) Id(value interface{}) *InputColorControl {
    a.Set("id", value)
    return a
}

/**
 * closeOnSelect
 */
func (a *InputColorControl) CloseOnSelect(value interface{}) *InputColorControl {
    a.Set("closeOnSelect", value)
    return a
}

/**
 * disabledOn
 */
func (a *InputColorControl) DisabledOn(value interface{}) *InputColorControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *InputColorControl) ClearValueOnHidden(value interface{}) *InputColorControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *InputColorControl) StaticInputClassName(value interface{}) *InputColorControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * name
 */
func (a *InputColorControl) Name(value interface{}) *InputColorControl {
    a.Set("name", value)
    return a
}

/**
 * desc
 */
func (a *InputColorControl) Desc(value interface{}) *InputColorControl {
    a.Set("desc", value)
    return a
}

/**
 * hidden
 */
func (a *InputColorControl) Hidden(value interface{}) *InputColorControl {
    a.Set("hidden", value)
    return a
}

/**
 * validations
 */
func (a *InputColorControl) Validations(value interface{}) *InputColorControl {
    a.Set("validations", value)
    return a
}

/**
 * visibleOn
 */
func (a *InputColorControl) VisibleOn(value interface{}) *InputColorControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticClassName
 */
func (a *InputColorControl) StaticClassName(value interface{}) *InputColorControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *InputColorControl) StaticSchema(value interface{}) *InputColorControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * labelRemark
 */
func (a *InputColorControl) LabelRemark(value interface{}) *InputColorControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * description
 */
func (a *InputColorControl) Description(value interface{}) *InputColorControl {
    a.Set("description", value)
    return a
}

/**
 * value
 */
func (a *InputColorControl) Value(value interface{}) *InputColorControl {
    a.Set("value", value)
    return a
}

/**
 * autoFill
 */
func (a *InputColorControl) AutoFill(value interface{}) *InputColorControl {
    a.Set("autoFill", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *InputColorControl) StaticPlaceholder(value interface{}) *InputColorControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * readOnly
 */
func (a *InputColorControl) ReadOnly(value interface{}) *InputColorControl {
    a.Set("readOnly", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *InputColorControl) ReadOnlyOn(value interface{}) *InputColorControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * row
 */
func (a *InputColorControl) Row(value interface{}) *InputColorControl {
    a.Set("row", value)
    return a
}

/**
 * onEvent
 */
func (a *InputColorControl) OnEvent(value interface{}) *InputColorControl {
    a.Set("onEvent", value)
    return a
}

/**
 * staticOn
 */
func (a *InputColorControl) StaticOn(value interface{}) *InputColorControl {
    a.Set("staticOn", value)
    return a
}

/**
 */
func (a *InputColorControl) Type(value interface{}) *InputColorControl {
    a.Set("type", value)
    return a
}

/**
 * labelClassName
 */
func (a *InputColorControl) LabelClassName(value interface{}) *InputColorControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * extraName
 */
func (a *InputColorControl) ExtraName(value interface{}) *InputColorControl {
    a.Set("extraName", value)
    return a
}

/**
 * visible
 */
func (a *InputColorControl) Visible(value interface{}) *InputColorControl {
    a.Set("visible", value)
    return a
}

/**
 * static
 */
func (a *InputColorControl) Static(value interface{}) *InputColorControl {
    a.Set("static", value)
    return a
}

/**
 * style
 */
func (a *InputColorControl) Style(value interface{}) *InputColorControl {
    a.Set("style", value)
    return a
}

/**
 * useMobileUI
 */
func (a *InputColorControl) UseMobileUI(value interface{}) *InputColorControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * presetColors
 */
func (a *InputColorControl) PresetColors(value interface{}) *InputColorControl {
    a.Set("presetColors", value)
    return a
}

/**
 * allowCustomColor
 */
func (a *InputColorControl) AllowCustomColor(value interface{}) *InputColorControl {
    a.Set("allowCustomColor", value)
    return a
}

/**
 * labelWidth
 */
func (a *InputColorControl) LabelWidth(value interface{}) *InputColorControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *InputColorControl) DescriptionClassName(value interface{}) *InputColorControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * required
 */
func (a *InputColorControl) Required(value interface{}) *InputColorControl {
    a.Set("required", value)
    return a
}

/**
 * className
 */
func (a *InputColorControl) ClassName(value interface{}) *InputColorControl {
    a.Set("className", value)
    return a
}

/**
 * format
 */
func (a *InputColorControl) Format(value interface{}) *InputColorControl {
    a.Set("format", value)
    return a
}

/**
 * popOverContainerSelector
 */
func (a *InputColorControl) PopOverContainerSelector(value interface{}) *InputColorControl {
    a.Set("popOverContainerSelector", value)
    return a
}

/**
 * labelOverflow
 */
func (a *InputColorControl) LabelOverflow(value interface{}) *InputColorControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * mode
 */
func (a *InputColorControl) Mode(value interface{}) *InputColorControl {
    a.Set("mode", value)
    return a
}

/**
 * size
 */
func (a *InputColorControl) Size(value interface{}) *InputColorControl {
    a.Set("size", value)
    return a
}

/**
 * clearable
 */
func (a *InputColorControl) Clearable(value interface{}) *InputColorControl {
    a.Set("clearable", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *InputColorControl) Width(value interface{}) *InputColorControl {
    a.Set("width", value)
    return a
}

/**
 * label
 */
func (a *InputColorControl) Label(value interface{}) *InputColorControl {
    a.Set("label", value)
    return a
}

/**
 * validateOnChange
 */
func (a *InputColorControl) ValidateOnChange(value interface{}) *InputColorControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * validationErrors
 */
func (a *InputColorControl) ValidationErrors(value interface{}) *InputColorControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * initAutoFill
 */
func (a *InputColorControl) InitAutoFill(value interface{}) *InputColorControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * disabled
 */
func (a *InputColorControl) Disabled(value interface{}) *InputColorControl {
    a.Set("disabled", value)
    return a
}
