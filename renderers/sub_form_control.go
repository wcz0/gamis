package renderers


/**
 * SubForm 子表单 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/subform

 * @author wcz0
 * @version 6.2.2
 */
type SubFormControl struct {
	*BaseRenderer
}

func NewSubFormControl() *SubFormControl {
    a := &SubFormControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-sub-form")
    return a
}


func (a *SubFormControl) Set(name string, value interface{}) *SubFormControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 */
func (a *SubFormControl) Type(value interface{}) *SubFormControl {
    a.Set("type", value)
    return a
}

/**
 * multiple
 */
func (a *SubFormControl) Multiple(value interface{}) *SubFormControl {
    a.Set("multiple", value)
    return a
}

/**
 * itemsClassName
 */
func (a *SubFormControl) ItemsClassName(value interface{}) *SubFormControl {
    a.Set("itemsClassName", value)
    return a
}

/**
 * name
 */
func (a *SubFormControl) Name(value interface{}) *SubFormControl {
    a.Set("name", value)
    return a
}

/**
 * submitOnChange
 */
func (a *SubFormControl) SubmitOnChange(value interface{}) *SubFormControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * className
 */
func (a *SubFormControl) ClassName(value interface{}) *SubFormControl {
    a.Set("className", value)
    return a
}

/**
 * labelField
 */
func (a *SubFormControl) LabelField(value interface{}) *SubFormControl {
    a.Set("labelField", value)
    return a
}

/**
 * validations
 */
func (a *SubFormControl) Validations(value interface{}) *SubFormControl {
    a.Set("validations", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *SubFormControl) StaticLabelClassName(value interface{}) *SubFormControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * addButtonText
 */
func (a *SubFormControl) AddButtonText(value interface{}) *SubFormControl {
    a.Set("addButtonText", value)
    return a
}

/**
 * labelOverflow
 */
func (a *SubFormControl) LabelOverflow(value interface{}) *SubFormControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * placeholder
 */
func (a *SubFormControl) Placeholder(value interface{}) *SubFormControl {
    a.Set("placeholder", value)
    return a
}

/**
 * validateApi
 */
func (a *SubFormControl) ValidateApi(value interface{}) *SubFormControl {
    a.Set("validateApi", value)
    return a
}

/**
 * style
 */
func (a *SubFormControl) Style(value interface{}) *SubFormControl {
    a.Set("style", value)
    return a
}

/**
 * size
 */
func (a *SubFormControl) Size(value interface{}) *SubFormControl {
    a.Set("size", value)
    return a
}

/**
 * labelClassName
 */
func (a *SubFormControl) LabelClassName(value interface{}) *SubFormControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * labelRemark
 */
func (a *SubFormControl) LabelRemark(value interface{}) *SubFormControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * description
 */
func (a *SubFormControl) Description(value interface{}) *SubFormControl {
    a.Set("description", value)
    return a
}

/**
 * desc
 */
func (a *SubFormControl) Desc(value interface{}) *SubFormControl {
    a.Set("desc", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *SubFormControl) DescriptionClassName(value interface{}) *SubFormControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * hiddenOn
 */
func (a *SubFormControl) HiddenOn(value interface{}) *SubFormControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *SubFormControl) StaticPlaceholder(value interface{}) *SubFormControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *SubFormControl) ReadOnlyOn(value interface{}) *SubFormControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * initAutoFill
 */
func (a *SubFormControl) InitAutoFill(value interface{}) *SubFormControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * row
 */
func (a *SubFormControl) Row(value interface{}) *SubFormControl {
    a.Set("row", value)
    return a
}

/**
 * disabledOn
 */
func (a *SubFormControl) DisabledOn(value interface{}) *SubFormControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * onEvent
 */
func (a *SubFormControl) OnEvent(value interface{}) *SubFormControl {
    a.Set("onEvent", value)
    return a
}

/**
 * static
 */
func (a *SubFormControl) Static(value interface{}) *SubFormControl {
    a.Set("static", value)
    return a
}

/**
 * maxLength
 */
func (a *SubFormControl) MaxLength(value interface{}) *SubFormControl {
    a.Set("maxLength", value)
    return a
}

/**
 * inputClassName
 */
func (a *SubFormControl) InputClassName(value interface{}) *SubFormControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * validationErrors
 */
func (a *SubFormControl) ValidationErrors(value interface{}) *SubFormControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * hidden
 */
func (a *SubFormControl) Hidden(value interface{}) *SubFormControl {
    a.Set("hidden", value)
    return a
}

/**
 * draggable
 */
func (a *SubFormControl) Draggable(value interface{}) *SubFormControl {
    a.Set("draggable", value)
    return a
}

/**
 * draggableTip
 */
func (a *SubFormControl) DraggableTip(value interface{}) *SubFormControl {
    a.Set("draggableTip", value)
    return a
}

/**
 * horizontal
 */
func (a *SubFormControl) Horizontal(value interface{}) *SubFormControl {
    a.Set("horizontal", value)
    return a
}

/**
 * autoFill
 */
func (a *SubFormControl) AutoFill(value interface{}) *SubFormControl {
    a.Set("autoFill", value)
    return a
}

/**
 * itemClassName
 */
func (a *SubFormControl) ItemClassName(value interface{}) *SubFormControl {
    a.Set("itemClassName", value)
    return a
}

/**
 * showErrorMsg
 */
func (a *SubFormControl) ShowErrorMsg(value interface{}) *SubFormControl {
    a.Set("showErrorMsg", value)
    return a
}

/**
 * labelAlign
 */
func (a *SubFormControl) LabelAlign(value interface{}) *SubFormControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * remark
 */
func (a *SubFormControl) Remark(value interface{}) *SubFormControl {
    a.Set("remark", value)
    return a
}

/**
 * disabled
 */
func (a *SubFormControl) Disabled(value interface{}) *SubFormControl {
    a.Set("disabled", value)
    return a
}

/**
 * staticOn
 */
func (a *SubFormControl) StaticOn(value interface{}) *SubFormControl {
    a.Set("staticOn", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *SubFormControl) ClearValueOnHidden(value interface{}) *SubFormControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * addable
 */
func (a *SubFormControl) Addable(value interface{}) *SubFormControl {
    a.Set("addable", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *SubFormControl) StaticInputClassName(value interface{}) *SubFormControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * useMobileUI
 */
func (a *SubFormControl) UseMobileUI(value interface{}) *SubFormControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * btnLabel
 */
func (a *SubFormControl) BtnLabel(value interface{}) *SubFormControl {
    a.Set("btnLabel", value)
    return a
}

/**
 * scaffold
 */
func (a *SubFormControl) Scaffold(value interface{}) *SubFormControl {
    a.Set("scaffold", value)
    return a
}

/**
 * mode
 */
func (a *SubFormControl) Mode(value interface{}) *SubFormControl {
    a.Set("mode", value)
    return a
}

/**
 * visible
 */
func (a *SubFormControl) Visible(value interface{}) *SubFormControl {
    a.Set("visible", value)
    return a
}

/**
 * removable
 */
func (a *SubFormControl) Removable(value interface{}) *SubFormControl {
    a.Set("removable", value)
    return a
}

/**
 * addButtonClassName
 */
func (a *SubFormControl) AddButtonClassName(value interface{}) *SubFormControl {
    a.Set("addButtonClassName", value)
    return a
}

/**
 * form
 */
func (a *SubFormControl) Form(value interface{}) *SubFormControl {
    a.Set("form", value)
    return a
}

/**
 * hint
 */
func (a *SubFormControl) Hint(value interface{}) *SubFormControl {
    a.Set("hint", value)
    return a
}

/**
 * required
 */
func (a *SubFormControl) Required(value interface{}) *SubFormControl {
    a.Set("required", value)
    return a
}

/**
 * value
 */
func (a *SubFormControl) Value(value interface{}) *SubFormControl {
    a.Set("value", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *SubFormControl) Width(value interface{}) *SubFormControl {
    a.Set("width", value)
    return a
}

/**
 * labelWidth
 */
func (a *SubFormControl) LabelWidth(value interface{}) *SubFormControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * validateOnChange
 */
func (a *SubFormControl) ValidateOnChange(value interface{}) *SubFormControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * visibleOn
 */
func (a *SubFormControl) VisibleOn(value interface{}) *SubFormControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticClassName
 */
func (a *SubFormControl) StaticClassName(value interface{}) *SubFormControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *SubFormControl) StaticSchema(value interface{}) *SubFormControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * editorSetting
 */
func (a *SubFormControl) EditorSetting(value interface{}) *SubFormControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * extraName
 */
func (a *SubFormControl) ExtraName(value interface{}) *SubFormControl {
    a.Set("extraName", value)
    return a
}

/**
 * id
 */
func (a *SubFormControl) Id(value interface{}) *SubFormControl {
    a.Set("id", value)
    return a
}

/**
 * minLength
 */
func (a *SubFormControl) MinLength(value interface{}) *SubFormControl {
    a.Set("minLength", value)
    return a
}

/**
 * label
 */
func (a *SubFormControl) Label(value interface{}) *SubFormControl {
    a.Set("label", value)
    return a
}

/**
 * readOnly
 */
func (a *SubFormControl) ReadOnly(value interface{}) *SubFormControl {
    a.Set("readOnly", value)
    return a
}

/**
 * inline
 */
func (a *SubFormControl) Inline(value interface{}) *SubFormControl {
    a.Set("inline", value)
    return a
}
