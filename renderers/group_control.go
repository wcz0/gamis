package renderers


/**
 * Group 表单集合渲染器，能让多个表单在一行显示 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/group

 * @author wcz0
 * @version 6.2.2
 */
type GroupControl struct {
	*BaseRenderer
}

func NewGroupControl() *GroupControl {
    a := &GroupControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "group")
    return a
}


func (a *GroupControl) Set(name string, value interface{}) *GroupControl {
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
func (a *GroupControl) LabelAlign(value interface{}) *GroupControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * name
 */
func (a *GroupControl) Name(value interface{}) *GroupControl {
    a.Set("name", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *GroupControl) ReadOnlyOn(value interface{}) *GroupControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * desc
 */
func (a *GroupControl) Desc(value interface{}) *GroupControl {
    a.Set("desc", value)
    return a
}

/**
 * visibleOn
 */
func (a *GroupControl) VisibleOn(value interface{}) *GroupControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * mode
 */
func (a *GroupControl) Mode(value interface{}) *GroupControl {
    a.Set("mode", value)
    return a
}

/**
 * value
 */
func (a *GroupControl) Value(value interface{}) *GroupControl {
    a.Set("value", value)
    return a
}

/**
 * static
 */
func (a *GroupControl) Static(value interface{}) *GroupControl {
    a.Set("static", value)
    return a
}

/**
 * validationErrors
 */
func (a *GroupControl) ValidationErrors(value interface{}) *GroupControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * visible
 */
func (a *GroupControl) Visible(value interface{}) *GroupControl {
    a.Set("visible", value)
    return a
}

/**
 */
func (a *GroupControl) Type(value interface{}) *GroupControl {
    a.Set("type", value)
    return a
}

/**
 * labelOverflow
 */
func (a *GroupControl) LabelOverflow(value interface{}) *GroupControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * extraName
 */
func (a *GroupControl) ExtraName(value interface{}) *GroupControl {
    a.Set("extraName", value)
    return a
}

/**
 * remark
 */
func (a *GroupControl) Remark(value interface{}) *GroupControl {
    a.Set("remark", value)
    return a
}

/**
 * readOnly
 */
func (a *GroupControl) ReadOnly(value interface{}) *GroupControl {
    a.Set("readOnly", value)
    return a
}

/**
 * inputClassName
 */
func (a *GroupControl) InputClassName(value interface{}) *GroupControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * row
 */
func (a *GroupControl) Row(value interface{}) *GroupControl {
    a.Set("row", value)
    return a
}

/**
 * useMobileUI
 */
func (a *GroupControl) UseMobileUI(value interface{}) *GroupControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * direction
 */
func (a *GroupControl) Direction(value interface{}) *GroupControl {
    a.Set("direction", value)
    return a
}

/**
 * inline
 */
func (a *GroupControl) Inline(value interface{}) *GroupControl {
    a.Set("inline", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *GroupControl) StaticInputClassName(value interface{}) *GroupControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *GroupControl) StaticSchema(value interface{}) *GroupControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * body
 */
func (a *GroupControl) Body(value interface{}) *GroupControl {
    a.Set("body", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *GroupControl) Width(value interface{}) *GroupControl {
    a.Set("width", value)
    return a
}

/**
 * placeholder
 */
func (a *GroupControl) Placeholder(value interface{}) *GroupControl {
    a.Set("placeholder", value)
    return a
}

/**
 * autoFill
 */
func (a *GroupControl) AutoFill(value interface{}) *GroupControl {
    a.Set("autoFill", value)
    return a
}

/**
 * horizontal
 */
func (a *GroupControl) Horizontal(value interface{}) *GroupControl {
    a.Set("horizontal", value)
    return a
}

/**
 * validateApi
 */
func (a *GroupControl) ValidateApi(value interface{}) *GroupControl {
    a.Set("validateApi", value)
    return a
}

/**
 * subFormMode
 */
func (a *GroupControl) SubFormMode(value interface{}) *GroupControl {
    a.Set("subFormMode", value)
    return a
}

/**
 * labelWidth
 */
func (a *GroupControl) LabelWidth(value interface{}) *GroupControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *GroupControl) DescriptionClassName(value interface{}) *GroupControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * onEvent
 */
func (a *GroupControl) OnEvent(value interface{}) *GroupControl {
    a.Set("onEvent", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *GroupControl) StaticPlaceholder(value interface{}) *GroupControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * description
 */
func (a *GroupControl) Description(value interface{}) *GroupControl {
    a.Set("description", value)
    return a
}

/**
 * required
 */
func (a *GroupControl) Required(value interface{}) *GroupControl {
    a.Set("required", value)
    return a
}

/**
 * validations
 */
func (a *GroupControl) Validations(value interface{}) *GroupControl {
    a.Set("validations", value)
    return a
}

/**
 * staticClassName
 */
func (a *GroupControl) StaticClassName(value interface{}) *GroupControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * subFormHorizontal
 */
func (a *GroupControl) SubFormHorizontal(value interface{}) *GroupControl {
    a.Set("subFormHorizontal", value)
    return a
}

/**
 * className
 */
func (a *GroupControl) ClassName(value interface{}) *GroupControl {
    a.Set("className", value)
    return a
}

/**
 * disabledOn
 */
func (a *GroupControl) DisabledOn(value interface{}) *GroupControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * hidden
 */
func (a *GroupControl) Hidden(value interface{}) *GroupControl {
    a.Set("hidden", value)
    return a
}

/**
 * labelClassName
 */
func (a *GroupControl) LabelClassName(value interface{}) *GroupControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * validateOnChange
 */
func (a *GroupControl) ValidateOnChange(value interface{}) *GroupControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *GroupControl) StaticLabelClassName(value interface{}) *GroupControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * gap
 */
func (a *GroupControl) Gap(value interface{}) *GroupControl {
    a.Set("gap", value)
    return a
}

/**
 * hiddenOn
 */
func (a *GroupControl) HiddenOn(value interface{}) *GroupControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * style
 */
func (a *GroupControl) Style(value interface{}) *GroupControl {
    a.Set("style", value)
    return a
}

/**
 * editorSetting
 */
func (a *GroupControl) EditorSetting(value interface{}) *GroupControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * disabled
 */
func (a *GroupControl) Disabled(value interface{}) *GroupControl {
    a.Set("disabled", value)
    return a
}

/**
 * labelRemark
 */
func (a *GroupControl) LabelRemark(value interface{}) *GroupControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *GroupControl) ClearValueOnHidden(value interface{}) *GroupControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * id
 */
func (a *GroupControl) Id(value interface{}) *GroupControl {
    a.Set("id", value)
    return a
}

/**
 * staticOn
 */
func (a *GroupControl) StaticOn(value interface{}) *GroupControl {
    a.Set("staticOn", value)
    return a
}

/**
 * label
 */
func (a *GroupControl) Label(value interface{}) *GroupControl {
    a.Set("label", value)
    return a
}

/**
 * hint
 */
func (a *GroupControl) Hint(value interface{}) *GroupControl {
    a.Set("hint", value)
    return a
}

/**
 * submitOnChange
 */
func (a *GroupControl) SubmitOnChange(value interface{}) *GroupControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * initAutoFill
 */
func (a *GroupControl) InitAutoFill(value interface{}) *GroupControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * size
 */
func (a *GroupControl) Size(value interface{}) *GroupControl {
    a.Set("size", value)
    return a
}
