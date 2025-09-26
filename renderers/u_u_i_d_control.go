package renderers


/**
 * UUID 功能性组件 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/uuid

 * @author wcz0
 * @version 6.2.2
 */
type UUIDControl struct {
	*BaseRenderer
}

func NewUUIDControl() *UUIDControl {
    a := &UUIDControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "uuid")
    return a
}


func (a *UUIDControl) Set(name string, value interface{}) *UUIDControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * style
 */
func (a *UUIDControl) Style(value interface{}) *UUIDControl {
    a.Set("style", value)
    return a
}

/**
 * editorSetting
 */
func (a *UUIDControl) EditorSetting(value interface{}) *UUIDControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * length
 */
func (a *UUIDControl) Length(value interface{}) *UUIDControl {
    a.Set("length", value)
    return a
}

/**
 * extraName
 */
func (a *UUIDControl) ExtraName(value interface{}) *UUIDControl {
    a.Set("extraName", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *UUIDControl) DescriptionClassName(value interface{}) *UUIDControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * visibleOn
 */
func (a *UUIDControl) VisibleOn(value interface{}) *UUIDControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * static
 */
func (a *UUIDControl) Static(value interface{}) *UUIDControl {
    a.Set("static", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *UUIDControl) StaticLabelClassName(value interface{}) *UUIDControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * readOnly
 */
func (a *UUIDControl) ReadOnly(value interface{}) *UUIDControl {
    a.Set("readOnly", value)
    return a
}

/**
 * desc
 */
func (a *UUIDControl) Desc(value interface{}) *UUIDControl {
    a.Set("desc", value)
    return a
}

/**
 * horizontal
 */
func (a *UUIDControl) Horizontal(value interface{}) *UUIDControl {
    a.Set("horizontal", value)
    return a
}

/**
 * value
 */
func (a *UUIDControl) Value(value interface{}) *UUIDControl {
    a.Set("value", value)
    return a
}

/**
 * autoFill
 */
func (a *UUIDControl) AutoFill(value interface{}) *UUIDControl {
    a.Set("autoFill", value)
    return a
}

/**
 * visible
 */
func (a *UUIDControl) Visible(value interface{}) *UUIDControl {
    a.Set("visible", value)
    return a
}

/**
 * staticClassName
 */
func (a *UUIDControl) StaticClassName(value interface{}) *UUIDControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *UUIDControl) StaticInputClassName(value interface{}) *UUIDControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * labelOverflow
 */
func (a *UUIDControl) LabelOverflow(value interface{}) *UUIDControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * description
 */
func (a *UUIDControl) Description(value interface{}) *UUIDControl {
    a.Set("description", value)
    return a
}

/**
 * hidden
 */
func (a *UUIDControl) Hidden(value interface{}) *UUIDControl {
    a.Set("hidden", value)
    return a
}

/**
 * onEvent
 */
func (a *UUIDControl) OnEvent(value interface{}) *UUIDControl {
    a.Set("onEvent", value)
    return a
}

/**
 * staticOn
 */
func (a *UUIDControl) StaticOn(value interface{}) *UUIDControl {
    a.Set("staticOn", value)
    return a
}

/**
 * useMobileUI
 */
func (a *UUIDControl) UseMobileUI(value interface{}) *UUIDControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * size
 */
func (a *UUIDControl) Size(value interface{}) *UUIDControl {
    a.Set("size", value)
    return a
}

/**
 */
func (a *UUIDControl) Type(value interface{}) *UUIDControl {
    a.Set("type", value)
    return a
}

/**
 * label
 */
func (a *UUIDControl) Label(value interface{}) *UUIDControl {
    a.Set("label", value)
    return a
}

/**
 * labelAlign
 */
func (a *UUIDControl) LabelAlign(value interface{}) *UUIDControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * labelWidth
 */
func (a *UUIDControl) LabelWidth(value interface{}) *UUIDControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * labelClassName
 */
func (a *UUIDControl) LabelClassName(value interface{}) *UUIDControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * validateOnChange
 */
func (a *UUIDControl) ValidateOnChange(value interface{}) *UUIDControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * row
 */
func (a *UUIDControl) Row(value interface{}) *UUIDControl {
    a.Set("row", value)
    return a
}

/**
 * hiddenOn
 */
func (a *UUIDControl) HiddenOn(value interface{}) *UUIDControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * placeholder
 */
func (a *UUIDControl) Placeholder(value interface{}) *UUIDControl {
    a.Set("placeholder", value)
    return a
}

/**
 * validationErrors
 */
func (a *UUIDControl) ValidationErrors(value interface{}) *UUIDControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * className
 */
func (a *UUIDControl) ClassName(value interface{}) *UUIDControl {
    a.Set("className", value)
    return a
}

/**
 * disabled
 */
func (a *UUIDControl) Disabled(value interface{}) *UUIDControl {
    a.Set("disabled", value)
    return a
}

/**
 * staticSchema
 */
func (a *UUIDControl) StaticSchema(value interface{}) *UUIDControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *UUIDControl) Width(value interface{}) *UUIDControl {
    a.Set("width", value)
    return a
}

/**
 * name
 */
func (a *UUIDControl) Name(value interface{}) *UUIDControl {
    a.Set("name", value)
    return a
}

/**
 * remark
 */
func (a *UUIDControl) Remark(value interface{}) *UUIDControl {
    a.Set("remark", value)
    return a
}

/**
 * mode
 */
func (a *UUIDControl) Mode(value interface{}) *UUIDControl {
    a.Set("mode", value)
    return a
}

/**
 * inline
 */
func (a *UUIDControl) Inline(value interface{}) *UUIDControl {
    a.Set("inline", value)
    return a
}

/**
 * validations
 */
func (a *UUIDControl) Validations(value interface{}) *UUIDControl {
    a.Set("validations", value)
    return a
}

/**
 * validateApi
 */
func (a *UUIDControl) ValidateApi(value interface{}) *UUIDControl {
    a.Set("validateApi", value)
    return a
}

/**
 * disabledOn
 */
func (a *UUIDControl) DisabledOn(value interface{}) *UUIDControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * labelRemark
 */
func (a *UUIDControl) LabelRemark(value interface{}) *UUIDControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * hint
 */
func (a *UUIDControl) Hint(value interface{}) *UUIDControl {
    a.Set("hint", value)
    return a
}

/**
 * submitOnChange
 */
func (a *UUIDControl) SubmitOnChange(value interface{}) *UUIDControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * inputClassName
 */
func (a *UUIDControl) InputClassName(value interface{}) *UUIDControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * required
 */
func (a *UUIDControl) Required(value interface{}) *UUIDControl {
    a.Set("required", value)
    return a
}

/**
 * initAutoFill
 */
func (a *UUIDControl) InitAutoFill(value interface{}) *UUIDControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * id
 */
func (a *UUIDControl) Id(value interface{}) *UUIDControl {
    a.Set("id", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *UUIDControl) ReadOnlyOn(value interface{}) *UUIDControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *UUIDControl) ClearValueOnHidden(value interface{}) *UUIDControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *UUIDControl) StaticPlaceholder(value interface{}) *UUIDControl {
    a.Set("staticPlaceholder", value)
    return a
}
