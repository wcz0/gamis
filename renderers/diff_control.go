package renderers


/**
 * Diff 编辑器 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/diff

 * @author wcz0
 * @version 6.2.2
 */
type DiffControl struct {
	*BaseRenderer
}

func NewDiffControl() *DiffControl {
    a := &DiffControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "diff-editor")
    return a
}


func (a *DiffControl) Set(name string, value interface{}) *DiffControl {
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
func (a *DiffControl) LabelAlign(value interface{}) *DiffControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *DiffControl) ClearValueOnHidden(value interface{}) *DiffControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * initAutoFill
 */
func (a *DiffControl) InitAutoFill(value interface{}) *DiffControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * hiddenOn
 */
func (a *DiffControl) HiddenOn(value interface{}) *DiffControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *DiffControl) StaticPlaceholder(value interface{}) *DiffControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticSchema
 */
func (a *DiffControl) StaticSchema(value interface{}) *DiffControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * label
 */
func (a *DiffControl) Label(value interface{}) *DiffControl {
    a.Set("label", value)
    return a
}

/**
 * desc
 */
func (a *DiffControl) Desc(value interface{}) *DiffControl {
    a.Set("desc", value)
    return a
}

/**
 * mode
 */
func (a *DiffControl) Mode(value interface{}) *DiffControl {
    a.Set("mode", value)
    return a
}

/**
 * useMobileUI
 */
func (a *DiffControl) UseMobileUI(value interface{}) *DiffControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * hint
 */
func (a *DiffControl) Hint(value interface{}) *DiffControl {
    a.Set("hint", value)
    return a
}

/**
 * style
 */
func (a *DiffControl) Style(value interface{}) *DiffControl {
    a.Set("style", value)
    return a
}

/**
 * editorSetting
 */
func (a *DiffControl) EditorSetting(value interface{}) *DiffControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * name
 */
func (a *DiffControl) Name(value interface{}) *DiffControl {
    a.Set("name", value)
    return a
}

/**
 * labelRemark
 */
func (a *DiffControl) LabelRemark(value interface{}) *DiffControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * static
 */
func (a *DiffControl) Static(value interface{}) *DiffControl {
    a.Set("static", value)
    return a
}

/**
 * visible
 */
func (a *DiffControl) Visible(value interface{}) *DiffControl {
    a.Set("visible", value)
    return a
}

/**
 * validateOnChange
 */
func (a *DiffControl) ValidateOnChange(value interface{}) *DiffControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * value
 */
func (a *DiffControl) Value(value interface{}) *DiffControl {
    a.Set("value", value)
    return a
}

/**
 * disabledOn
 */
func (a *DiffControl) DisabledOn(value interface{}) *DiffControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * diffValue
 */
func (a *DiffControl) DiffValue(value interface{}) *DiffControl {
    a.Set("diffValue", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *DiffControl) Width(value interface{}) *DiffControl {
    a.Set("width", value)
    return a
}

/**
 * hidden
 */
func (a *DiffControl) Hidden(value interface{}) *DiffControl {
    a.Set("hidden", value)
    return a
}

/**
 * staticClassName
 */
func (a *DiffControl) StaticClassName(value interface{}) *DiffControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * extraName
 */
func (a *DiffControl) ExtraName(value interface{}) *DiffControl {
    a.Set("extraName", value)
    return a
}

/**
 * horizontal
 */
func (a *DiffControl) Horizontal(value interface{}) *DiffControl {
    a.Set("horizontal", value)
    return a
}

/**
 * placeholder
 */
func (a *DiffControl) Placeholder(value interface{}) *DiffControl {
    a.Set("placeholder", value)
    return a
}

/**
 * validations
 */
func (a *DiffControl) Validations(value interface{}) *DiffControl {
    a.Set("validations", value)
    return a
}

/**
 * staticOn
 */
func (a *DiffControl) StaticOn(value interface{}) *DiffControl {
    a.Set("staticOn", value)
    return a
}

/**
 * size
 */
func (a *DiffControl) Size(value interface{}) *DiffControl {
    a.Set("size", value)
    return a
}

/**
 * remark
 */
func (a *DiffControl) Remark(value interface{}) *DiffControl {
    a.Set("remark", value)
    return a
}

/**
 * autoFill
 */
func (a *DiffControl) AutoFill(value interface{}) *DiffControl {
    a.Set("autoFill", value)
    return a
}

/**
 * onEvent
 */
func (a *DiffControl) OnEvent(value interface{}) *DiffControl {
    a.Set("onEvent", value)
    return a
}

/**
 * options
 */
func (a *DiffControl) Options(value interface{}) *DiffControl {
    a.Set("options", value)
    return a
}

/**
 * labelClassName
 */
func (a *DiffControl) LabelClassName(value interface{}) *DiffControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *DiffControl) ReadOnlyOn(value interface{}) *DiffControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * visibleOn
 */
func (a *DiffControl) VisibleOn(value interface{}) *DiffControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *DiffControl) StaticInputClassName(value interface{}) *DiffControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *DiffControl) DescriptionClassName(value interface{}) *DiffControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * validationErrors
 */
func (a *DiffControl) ValidationErrors(value interface{}) *DiffControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * language
 */
func (a *DiffControl) Language(value interface{}) *DiffControl {
    a.Set("language", value)
    return a
}

/**
 * labelWidth
 */
func (a *DiffControl) LabelWidth(value interface{}) *DiffControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * labelOverflow
 */
func (a *DiffControl) LabelOverflow(value interface{}) *DiffControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * description
 */
func (a *DiffControl) Description(value interface{}) *DiffControl {
    a.Set("description", value)
    return a
}

/**
 * required
 */
func (a *DiffControl) Required(value interface{}) *DiffControl {
    a.Set("required", value)
    return a
}

/**
 * disabled
 */
func (a *DiffControl) Disabled(value interface{}) *DiffControl {
    a.Set("disabled", value)
    return a
}

/**
 * submitOnChange
 */
func (a *DiffControl) SubmitOnChange(value interface{}) *DiffControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * inputClassName
 */
func (a *DiffControl) InputClassName(value interface{}) *DiffControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * validateApi
 */
func (a *DiffControl) ValidateApi(value interface{}) *DiffControl {
    a.Set("validateApi", value)
    return a
}

/**
 * id
 */
func (a *DiffControl) Id(value interface{}) *DiffControl {
    a.Set("id", value)
    return a
}

/**
 */
func (a *DiffControl) Type(value interface{}) *DiffControl {
    a.Set("type", value)
    return a
}

/**
 * readOnly
 */
func (a *DiffControl) ReadOnly(value interface{}) *DiffControl {
    a.Set("readOnly", value)
    return a
}

/**
 * inline
 */
func (a *DiffControl) Inline(value interface{}) *DiffControl {
    a.Set("inline", value)
    return a
}

/**
 * row
 */
func (a *DiffControl) Row(value interface{}) *DiffControl {
    a.Set("row", value)
    return a
}

/**
 * className
 */
func (a *DiffControl) ClassName(value interface{}) *DiffControl {
    a.Set("className", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *DiffControl) StaticLabelClassName(value interface{}) *DiffControl {
    a.Set("staticLabelClassName", value)
    return a
}
