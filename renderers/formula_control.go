package renderers


/**
 * 公式功能控件。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/formula

 * @author wcz0
 * @version 6.2.2
 */
type FormulaControl struct {
	*BaseRenderer
}

func NewFormulaControl() *FormulaControl {
    a := &FormulaControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "formula")
    return a
}


func (a *FormulaControl) Set(name string, value interface{}) *FormulaControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * value
 */
func (a *FormulaControl) Value(value interface{}) *FormulaControl {
    a.Set("value", value)
    return a
}

/**
 * onEvent
 */
func (a *FormulaControl) OnEvent(value interface{}) *FormulaControl {
    a.Set("onEvent", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *FormulaControl) StaticLabelClassName(value interface{}) *FormulaControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * mode
 */
func (a *FormulaControl) Mode(value interface{}) *FormulaControl {
    a.Set("mode", value)
    return a
}

/**
 * autoFill
 */
func (a *FormulaControl) AutoFill(value interface{}) *FormulaControl {
    a.Set("autoFill", value)
    return a
}

/**
 * hidden
 */
func (a *FormulaControl) Hidden(value interface{}) *FormulaControl {
    a.Set("hidden", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *FormulaControl) StaticInputClassName(value interface{}) *FormulaControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * initSet
 */
func (a *FormulaControl) InitSet(value interface{}) *FormulaControl {
    a.Set("initSet", value)
    return a
}

/**
 * validateOnChange
 */
func (a *FormulaControl) ValidateOnChange(value interface{}) *FormulaControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * labelAlign
 */
func (a *FormulaControl) LabelAlign(value interface{}) *FormulaControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * labelWidth
 */
func (a *FormulaControl) LabelWidth(value interface{}) *FormulaControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * name
 */
func (a *FormulaControl) Name(value interface{}) *FormulaControl {
    a.Set("name", value)
    return a
}

/**
 * row
 */
func (a *FormulaControl) Row(value interface{}) *FormulaControl {
    a.Set("row", value)
    return a
}

/**
 * placeholder
 */
func (a *FormulaControl) Placeholder(value interface{}) *FormulaControl {
    a.Set("placeholder", value)
    return a
}

/**
 * useMobileUI
 */
func (a *FormulaControl) UseMobileUI(value interface{}) *FormulaControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *FormulaControl) ClearValueOnHidden(value interface{}) *FormulaControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * disabledOn
 */
func (a *FormulaControl) DisabledOn(value interface{}) *FormulaControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * size
 */
func (a *FormulaControl) Size(value interface{}) *FormulaControl {
    a.Set("size", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *FormulaControl) ReadOnlyOn(value interface{}) *FormulaControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * initAutoFill
 */
func (a *FormulaControl) InitAutoFill(value interface{}) *FormulaControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 */
func (a *FormulaControl) Type(value interface{}) *FormulaControl {
    a.Set("type", value)
    return a
}

/**
 * readOnly
 */
func (a *FormulaControl) ReadOnly(value interface{}) *FormulaControl {
    a.Set("readOnly", value)
    return a
}

/**
 * inputClassName
 */
func (a *FormulaControl) InputClassName(value interface{}) *FormulaControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * disabled
 */
func (a *FormulaControl) Disabled(value interface{}) *FormulaControl {
    a.Set("disabled", value)
    return a
}

/**
 * labelRemark
 */
func (a *FormulaControl) LabelRemark(value interface{}) *FormulaControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * hint
 */
func (a *FormulaControl) Hint(value interface{}) *FormulaControl {
    a.Set("hint", value)
    return a
}

/**
 * submitOnChange
 */
func (a *FormulaControl) SubmitOnChange(value interface{}) *FormulaControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * desc
 */
func (a *FormulaControl) Desc(value interface{}) *FormulaControl {
    a.Set("desc", value)
    return a
}

/**
 * validateApi
 */
func (a *FormulaControl) ValidateApi(value interface{}) *FormulaControl {
    a.Set("validateApi", value)
    return a
}

/**
 * label
 */
func (a *FormulaControl) Label(value interface{}) *FormulaControl {
    a.Set("label", value)
    return a
}

/**
 * extraName
 */
func (a *FormulaControl) ExtraName(value interface{}) *FormulaControl {
    a.Set("extraName", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *FormulaControl) DescriptionClassName(value interface{}) *FormulaControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * horizontal
 */
func (a *FormulaControl) Horizontal(value interface{}) *FormulaControl {
    a.Set("horizontal", value)
    return a
}

/**
 * inline
 */
func (a *FormulaControl) Inline(value interface{}) *FormulaControl {
    a.Set("inline", value)
    return a
}

/**
 * className
 */
func (a *FormulaControl) ClassName(value interface{}) *FormulaControl {
    a.Set("className", value)
    return a
}

/**
 * autoSet
 */
func (a *FormulaControl) AutoSet(value interface{}) *FormulaControl {
    a.Set("autoSet", value)
    return a
}

/**
 * labelOverflow
 */
func (a *FormulaControl) LabelOverflow(value interface{}) *FormulaControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * hiddenOn
 */
func (a *FormulaControl) HiddenOn(value interface{}) *FormulaControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * visibleOn
 */
func (a *FormulaControl) VisibleOn(value interface{}) *FormulaControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * id
 */
func (a *FormulaControl) Id(value interface{}) *FormulaControl {
    a.Set("id", value)
    return a
}

/**
 * staticOn
 */
func (a *FormulaControl) StaticOn(value interface{}) *FormulaControl {
    a.Set("staticOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *FormulaControl) StaticPlaceholder(value interface{}) *FormulaControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * condition
 */
func (a *FormulaControl) Condition(value interface{}) *FormulaControl {
    a.Set("condition", value)
    return a
}

/**
 * formula
 */
func (a *FormulaControl) Formula(value interface{}) *FormulaControl {
    a.Set("formula", value)
    return a
}

/**
 * visible
 */
func (a *FormulaControl) Visible(value interface{}) *FormulaControl {
    a.Set("visible", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *FormulaControl) Width(value interface{}) *FormulaControl {
    a.Set("width", value)
    return a
}

/**
 * description
 */
func (a *FormulaControl) Description(value interface{}) *FormulaControl {
    a.Set("description", value)
    return a
}

/**
 * static
 */
func (a *FormulaControl) Static(value interface{}) *FormulaControl {
    a.Set("static", value)
    return a
}

/**
 * style
 */
func (a *FormulaControl) Style(value interface{}) *FormulaControl {
    a.Set("style", value)
    return a
}

/**
 * remark
 */
func (a *FormulaControl) Remark(value interface{}) *FormulaControl {
    a.Set("remark", value)
    return a
}

/**
 * required
 */
func (a *FormulaControl) Required(value interface{}) *FormulaControl {
    a.Set("required", value)
    return a
}

/**
 * staticClassName
 */
func (a *FormulaControl) StaticClassName(value interface{}) *FormulaControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *FormulaControl) StaticSchema(value interface{}) *FormulaControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * labelClassName
 */
func (a *FormulaControl) LabelClassName(value interface{}) *FormulaControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * validationErrors
 */
func (a *FormulaControl) ValidationErrors(value interface{}) *FormulaControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * validations
 */
func (a *FormulaControl) Validations(value interface{}) *FormulaControl {
    a.Set("validations", value)
    return a
}

/**
 * editorSetting
 */
func (a *FormulaControl) EditorSetting(value interface{}) *FormulaControl {
    a.Set("editorSetting", value)
    return a
}
