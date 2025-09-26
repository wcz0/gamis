package renderers


/**
 * InputFormula 公式编辑器 文档：https://baidu.gitee.io/amis/zh-CN/components/form/input-formula

 * @author wcz0
 * @version 6.2.2
 */
type BaseInputFormulaControl struct {
	*BaseRenderer
}

func NewBaseInputFormulaControl() *BaseInputFormulaControl {
    a := &BaseInputFormulaControl{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}


func (a *BaseInputFormulaControl) Set(name string, value interface{}) *BaseInputFormulaControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * visibleOn
 */
func (a *BaseInputFormulaControl) VisibleOn(value interface{}) *BaseInputFormulaControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * variables
 */
func (a *BaseInputFormulaControl) Variables(value interface{}) *BaseInputFormulaControl {
    a.Set("variables", value)
    return a
}

/**
 * level
 */
func (a *BaseInputFormulaControl) Level(value interface{}) *BaseInputFormulaControl {
    a.Set("level", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *BaseInputFormulaControl) Width(value interface{}) *BaseInputFormulaControl {
    a.Set("width", value)
    return a
}

/**
 * desc
 */
func (a *BaseInputFormulaControl) Desc(value interface{}) *BaseInputFormulaControl {
    a.Set("desc", value)
    return a
}

/**
 * inline
 */
func (a *BaseInputFormulaControl) Inline(value interface{}) *BaseInputFormulaControl {
    a.Set("inline", value)
    return a
}

/**
 * validations
 */
func (a *BaseInputFormulaControl) Validations(value interface{}) *BaseInputFormulaControl {
    a.Set("validations", value)
    return a
}

/**
 * disabledOn
 */
func (a *BaseInputFormulaControl) DisabledOn(value interface{}) *BaseInputFormulaControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * id
 */
func (a *BaseInputFormulaControl) Id(value interface{}) *BaseInputFormulaControl {
    a.Set("id", value)
    return a
}

/**
 * static
 */
func (a *BaseInputFormulaControl) Static(value interface{}) *BaseInputFormulaControl {
    a.Set("static", value)
    return a
}

/**
 * icon
 */
func (a *BaseInputFormulaControl) Icon(value interface{}) *BaseInputFormulaControl {
    a.Set("icon", value)
    return a
}

/**
 * hidden
 */
func (a *BaseInputFormulaControl) Hidden(value interface{}) *BaseInputFormulaControl {
    a.Set("hidden", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *BaseInputFormulaControl) StaticPlaceholder(value interface{}) *BaseInputFormulaControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * size
 */
func (a *BaseInputFormulaControl) Size(value interface{}) *BaseInputFormulaControl {
    a.Set("size", value)
    return a
}

/**
 * functions
 */
func (a *BaseInputFormulaControl) Functions(value interface{}) *BaseInputFormulaControl {
    a.Set("functions", value)
    return a
}

/**
 * btnLabel
 */
func (a *BaseInputFormulaControl) BtnLabel(value interface{}) *BaseInputFormulaControl {
    a.Set("btnLabel", value)
    return a
}

/**
 * name
 */
func (a *BaseInputFormulaControl) Name(value interface{}) *BaseInputFormulaControl {
    a.Set("name", value)
    return a
}

/**
 * required
 */
func (a *BaseInputFormulaControl) Required(value interface{}) *BaseInputFormulaControl {
    a.Set("required", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *BaseInputFormulaControl) ClearValueOnHidden(value interface{}) *BaseInputFormulaControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * staticOn
 */
func (a *BaseInputFormulaControl) StaticOn(value interface{}) *BaseInputFormulaControl {
    a.Set("staticOn", value)
    return a
}

/**
 * title
 */
func (a *BaseInputFormulaControl) Title(value interface{}) *BaseInputFormulaControl {
    a.Set("title", value)
    return a
}

/**
 * description
 */
func (a *BaseInputFormulaControl) Description(value interface{}) *BaseInputFormulaControl {
    a.Set("description", value)
    return a
}

/**
 * disabled
 */
func (a *BaseInputFormulaControl) Disabled(value interface{}) *BaseInputFormulaControl {
    a.Set("disabled", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *BaseInputFormulaControl) StaticLabelClassName(value interface{}) *BaseInputFormulaControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * selfVariableName
 */
func (a *BaseInputFormulaControl) SelfVariableName(value interface{}) *BaseInputFormulaControl {
    a.Set("selfVariableName", value)
    return a
}

/**
 * submitOnChange
 */
func (a *BaseInputFormulaControl) SubmitOnChange(value interface{}) *BaseInputFormulaControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * hiddenOn
 */
func (a *BaseInputFormulaControl) HiddenOn(value interface{}) *BaseInputFormulaControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * editorSetting
 */
func (a *BaseInputFormulaControl) EditorSetting(value interface{}) *BaseInputFormulaControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * evalMode
 */
func (a *BaseInputFormulaControl) EvalMode(value interface{}) *BaseInputFormulaControl {
    a.Set("evalMode", value)
    return a
}

/**
 * hint
 */
func (a *BaseInputFormulaControl) Hint(value interface{}) *BaseInputFormulaControl {
    a.Set("hint", value)
    return a
}

/**
 * placeholder
 */
func (a *BaseInputFormulaControl) Placeholder(value interface{}) *BaseInputFormulaControl {
    a.Set("placeholder", value)
    return a
}

/**
 * className
 */
func (a *BaseInputFormulaControl) ClassName(value interface{}) *BaseInputFormulaControl {
    a.Set("className", value)
    return a
}

/**
 * visible
 */
func (a *BaseInputFormulaControl) Visible(value interface{}) *BaseInputFormulaControl {
    a.Set("visible", value)
    return a
}

/**
 * inputMode
 */
func (a *BaseInputFormulaControl) InputMode(value interface{}) *BaseInputFormulaControl {
    a.Set("inputMode", value)
    return a
}

/**
 * variableClassName
 */
func (a *BaseInputFormulaControl) VariableClassName(value interface{}) *BaseInputFormulaControl {
    a.Set("variableClassName", value)
    return a
}

/**
 * inputClassName
 */
func (a *BaseInputFormulaControl) InputClassName(value interface{}) *BaseInputFormulaControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * onEvent
 */
func (a *BaseInputFormulaControl) OnEvent(value interface{}) *BaseInputFormulaControl {
    a.Set("onEvent", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *BaseInputFormulaControl) StaticInputClassName(value interface{}) *BaseInputFormulaControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * type
 */
func (a *BaseInputFormulaControl) Type(value interface{}) *BaseInputFormulaControl {
    a.Set("type", value)
    return a
}

/**
 * labelOverflow
 */
func (a *BaseInputFormulaControl) LabelOverflow(value interface{}) *BaseInputFormulaControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *BaseInputFormulaControl) ReadOnlyOn(value interface{}) *BaseInputFormulaControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * header
 */
func (a *BaseInputFormulaControl) Header(value interface{}) *BaseInputFormulaControl {
    a.Set("header", value)
    return a
}

/**
 * btnSize
 */
func (a *BaseInputFormulaControl) BtnSize(value interface{}) *BaseInputFormulaControl {
    a.Set("btnSize", value)
    return a
}

/**
 * labelRemark
 */
func (a *BaseInputFormulaControl) LabelRemark(value interface{}) *BaseInputFormulaControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * style
 */
func (a *BaseInputFormulaControl) Style(value interface{}) *BaseInputFormulaControl {
    a.Set("style", value)
    return a
}

/**
 * variableMode
 */
func (a *BaseInputFormulaControl) VariableMode(value interface{}) *BaseInputFormulaControl {
    a.Set("variableMode", value)
    return a
}

/**
 * allowInput
 */
func (a *BaseInputFormulaControl) AllowInput(value interface{}) *BaseInputFormulaControl {
    a.Set("allowInput", value)
    return a
}

/**
 * labelAlign
 */
func (a *BaseInputFormulaControl) LabelAlign(value interface{}) *BaseInputFormulaControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * mode
 */
func (a *BaseInputFormulaControl) Mode(value interface{}) *BaseInputFormulaControl {
    a.Set("mode", value)
    return a
}

/**
 * horizontal
 */
func (a *BaseInputFormulaControl) Horizontal(value interface{}) *BaseInputFormulaControl {
    a.Set("horizontal", value)
    return a
}

/**
 * validateApi
 */
func (a *BaseInputFormulaControl) ValidateApi(value interface{}) *BaseInputFormulaControl {
    a.Set("validateApi", value)
    return a
}

/**
 * staticSchema
 */
func (a *BaseInputFormulaControl) StaticSchema(value interface{}) *BaseInputFormulaControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * mixedMode
 */
func (a *BaseInputFormulaControl) MixedMode(value interface{}) *BaseInputFormulaControl {
    a.Set("mixedMode", value)
    return a
}

/**
 * inputSettings
 */
func (a *BaseInputFormulaControl) InputSettings(value interface{}) *BaseInputFormulaControl {
    a.Set("inputSettings", value)
    return a
}

/**
 * autoFill
 */
func (a *BaseInputFormulaControl) AutoFill(value interface{}) *BaseInputFormulaControl {
    a.Set("autoFill", value)
    return a
}

/**
 * row
 */
func (a *BaseInputFormulaControl) Row(value interface{}) *BaseInputFormulaControl {
    a.Set("row", value)
    return a
}

/**
 * useMobileUI
 */
func (a *BaseInputFormulaControl) UseMobileUI(value interface{}) *BaseInputFormulaControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * functionClassName
 */
func (a *BaseInputFormulaControl) FunctionClassName(value interface{}) *BaseInputFormulaControl {
    a.Set("functionClassName", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *BaseInputFormulaControl) DescriptionClassName(value interface{}) *BaseInputFormulaControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * validateOnChange
 */
func (a *BaseInputFormulaControl) ValidateOnChange(value interface{}) *BaseInputFormulaControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * labelClassName
 */
func (a *BaseInputFormulaControl) LabelClassName(value interface{}) *BaseInputFormulaControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * remark
 */
func (a *BaseInputFormulaControl) Remark(value interface{}) *BaseInputFormulaControl {
    a.Set("remark", value)
    return a
}

/**
 * initAutoFill
 */
func (a *BaseInputFormulaControl) InitAutoFill(value interface{}) *BaseInputFormulaControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * borderMode
 */
func (a *BaseInputFormulaControl) BorderMode(value interface{}) *BaseInputFormulaControl {
    a.Set("borderMode", value)
    return a
}

/**
 * label
 */
func (a *BaseInputFormulaControl) Label(value interface{}) *BaseInputFormulaControl {
    a.Set("label", value)
    return a
}

/**
 * extraName
 */
func (a *BaseInputFormulaControl) ExtraName(value interface{}) *BaseInputFormulaControl {
    a.Set("extraName", value)
    return a
}

/**
 * value
 */
func (a *BaseInputFormulaControl) Value(value interface{}) *BaseInputFormulaControl {
    a.Set("value", value)
    return a
}

/**
 * staticClassName
 */
func (a *BaseInputFormulaControl) StaticClassName(value interface{}) *BaseInputFormulaControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * labelWidth
 */
func (a *BaseInputFormulaControl) LabelWidth(value interface{}) *BaseInputFormulaControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * readOnly
 */
func (a *BaseInputFormulaControl) ReadOnly(value interface{}) *BaseInputFormulaControl {
    a.Set("readOnly", value)
    return a
}

/**
 * validationErrors
 */
func (a *BaseInputFormulaControl) ValidationErrors(value interface{}) *BaseInputFormulaControl {
    a.Set("validationErrors", value)
    return a
}
