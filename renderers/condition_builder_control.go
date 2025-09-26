package renderers


/**
 * 条件组合控件 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/condition-builder

 * @author wcz0
 * @version 6.2.2
 */
type ConditionBuilderControl struct {
	*BaseRenderer
}

func NewConditionBuilderControl() *ConditionBuilderControl {
    a := &ConditionBuilderControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "condition-builder")
    return a
}


func (a *ConditionBuilderControl) Set(name string, value interface{}) *ConditionBuilderControl {
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
func (a *ConditionBuilderControl) VisibleOn(value interface{}) *ConditionBuilderControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticOn
 */
func (a *ConditionBuilderControl) StaticOn(value interface{}) *ConditionBuilderControl {
    a.Set("staticOn", value)
    return a
}

/**
 * onEvent
 */
func (a *ConditionBuilderControl) OnEvent(value interface{}) *ConditionBuilderControl {
    a.Set("onEvent", value)
    return a
}

/**
 * remark
 */
func (a *ConditionBuilderControl) Remark(value interface{}) *ConditionBuilderControl {
    a.Set("remark", value)
    return a
}

/**
 * hint
 */
func (a *ConditionBuilderControl) Hint(value interface{}) *ConditionBuilderControl {
    a.Set("hint", value)
    return a
}

/**
 * value
 */
func (a *ConditionBuilderControl) Value(value interface{}) *ConditionBuilderControl {
    a.Set("value", value)
    return a
}

/**
 * editorSetting
 */
func (a *ConditionBuilderControl) EditorSetting(value interface{}) *ConditionBuilderControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * draggable
 */
func (a *ConditionBuilderControl) Draggable(value interface{}) *ConditionBuilderControl {
    a.Set("draggable", value)
    return a
}

/**
 * labelOverflow
 */
func (a *ConditionBuilderControl) LabelOverflow(value interface{}) *ConditionBuilderControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * staticClassName
 */
func (a *ConditionBuilderControl) StaticClassName(value interface{}) *ConditionBuilderControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * extraName
 */
func (a *ConditionBuilderControl) ExtraName(value interface{}) *ConditionBuilderControl {
    a.Set("extraName", value)
    return a
}

/**
 * placeholder
 */
func (a *ConditionBuilderControl) Placeholder(value interface{}) *ConditionBuilderControl {
    a.Set("placeholder", value)
    return a
}

/**
 * validations
 */
func (a *ConditionBuilderControl) Validations(value interface{}) *ConditionBuilderControl {
    a.Set("validations", value)
    return a
}

/**
 * initAutoFill
 */
func (a *ConditionBuilderControl) InitAutoFill(value interface{}) *ConditionBuilderControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * disabledOn
 */
func (a *ConditionBuilderControl) DisabledOn(value interface{}) *ConditionBuilderControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * hidden
 */
func (a *ConditionBuilderControl) Hidden(value interface{}) *ConditionBuilderControl {
    a.Set("hidden", value)
    return a
}

/**
 * disabled
 */
func (a *ConditionBuilderControl) Disabled(value interface{}) *ConditionBuilderControl {
    a.Set("disabled", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *ConditionBuilderControl) StaticPlaceholder(value interface{}) *ConditionBuilderControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * embed
 */
func (a *ConditionBuilderControl) Embed(value interface{}) *ConditionBuilderControl {
    a.Set("embed", value)
    return a
}

/**
 * builderMode
 */
func (a *ConditionBuilderControl) BuilderMode(value interface{}) *ConditionBuilderControl {
    a.Set("builderMode", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *ConditionBuilderControl) DescriptionClassName(value interface{}) *ConditionBuilderControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * required
 */
func (a *ConditionBuilderControl) Required(value interface{}) *ConditionBuilderControl {
    a.Set("required", value)
    return a
}

/**
 * visible
 */
func (a *ConditionBuilderControl) Visible(value interface{}) *ConditionBuilderControl {
    a.Set("visible", value)
    return a
}

/**
 * addBtnVisibleOn
 */
func (a *ConditionBuilderControl) AddBtnVisibleOn(value interface{}) *ConditionBuilderControl {
    a.Set("addBtnVisibleOn", value)
    return a
}

/**
 * staticSchema
 */
func (a *ConditionBuilderControl) StaticSchema(value interface{}) *ConditionBuilderControl {
    a.Set("staticSchema", value)
    return a
}

/**
 */
func (a *ConditionBuilderControl) Type(value interface{}) *ConditionBuilderControl {
    a.Set("type", value)
    return a
}

/**
 * labelWidth
 */
func (a *ConditionBuilderControl) LabelWidth(value interface{}) *ConditionBuilderControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * name
 */
func (a *ConditionBuilderControl) Name(value interface{}) *ConditionBuilderControl {
    a.Set("name", value)
    return a
}

/**
 * labelRemark
 */
func (a *ConditionBuilderControl) LabelRemark(value interface{}) *ConditionBuilderControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * description
 */
func (a *ConditionBuilderControl) Description(value interface{}) *ConditionBuilderControl {
    a.Set("description", value)
    return a
}

/**
 * className
 */
func (a *ConditionBuilderControl) ClassName(value interface{}) *ConditionBuilderControl {
    a.Set("className", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *ConditionBuilderControl) Width(value interface{}) *ConditionBuilderControl {
    a.Set("width", value)
    return a
}

/**
 * labelClassName
 */
func (a *ConditionBuilderControl) LabelClassName(value interface{}) *ConditionBuilderControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * inputClassName
 */
func (a *ConditionBuilderControl) InputClassName(value interface{}) *ConditionBuilderControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *ConditionBuilderControl) StaticLabelClassName(value interface{}) *ConditionBuilderControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * size
 */
func (a *ConditionBuilderControl) Size(value interface{}) *ConditionBuilderControl {
    a.Set("size", value)
    return a
}

/**
 * validateApi
 */
func (a *ConditionBuilderControl) ValidateApi(value interface{}) *ConditionBuilderControl {
    a.Set("validateApi", value)
    return a
}

/**
 * id
 */
func (a *ConditionBuilderControl) Id(value interface{}) *ConditionBuilderControl {
    a.Set("id", value)
    return a
}

/**
 * static
 */
func (a *ConditionBuilderControl) Static(value interface{}) *ConditionBuilderControl {
    a.Set("static", value)
    return a
}

/**
 * formulaForIf
 */
func (a *ConditionBuilderControl) FormulaForIf(value interface{}) *ConditionBuilderControl {
    a.Set("formulaForIf", value)
    return a
}

/**
 * labelAlign
 */
func (a *ConditionBuilderControl) LabelAlign(value interface{}) *ConditionBuilderControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * mode
 */
func (a *ConditionBuilderControl) Mode(value interface{}) *ConditionBuilderControl {
    a.Set("mode", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *ConditionBuilderControl) ClearValueOnHidden(value interface{}) *ConditionBuilderControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * hiddenOn
 */
func (a *ConditionBuilderControl) HiddenOn(value interface{}) *ConditionBuilderControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * useMobileUI
 */
func (a *ConditionBuilderControl) UseMobileUI(value interface{}) *ConditionBuilderControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * validateOnChange
 */
func (a *ConditionBuilderControl) ValidateOnChange(value interface{}) *ConditionBuilderControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * readOnly
 */
func (a *ConditionBuilderControl) ReadOnly(value interface{}) *ConditionBuilderControl {
    a.Set("readOnly", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *ConditionBuilderControl) StaticInputClassName(value interface{}) *ConditionBuilderControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * fields
 */
func (a *ConditionBuilderControl) Fields(value interface{}) *ConditionBuilderControl {
    a.Set("fields", value)
    return a
}

/**
 * formula
 */
func (a *ConditionBuilderControl) Formula(value interface{}) *ConditionBuilderControl {
    a.Set("formula", value)
    return a
}

/**
 * submitOnChange
 */
func (a *ConditionBuilderControl) SubmitOnChange(value interface{}) *ConditionBuilderControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *ConditionBuilderControl) ReadOnlyOn(value interface{}) *ConditionBuilderControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * inline
 */
func (a *ConditionBuilderControl) Inline(value interface{}) *ConditionBuilderControl {
    a.Set("inline", value)
    return a
}

/**
 * style
 */
func (a *ConditionBuilderControl) Style(value interface{}) *ConditionBuilderControl {
    a.Set("style", value)
    return a
}

/**
 * source
 */
func (a *ConditionBuilderControl) Source(value interface{}) *ConditionBuilderControl {
    a.Set("source", value)
    return a
}

/**
 * label
 */
func (a *ConditionBuilderControl) Label(value interface{}) *ConditionBuilderControl {
    a.Set("label", value)
    return a
}

/**
 * desc
 */
func (a *ConditionBuilderControl) Desc(value interface{}) *ConditionBuilderControl {
    a.Set("desc", value)
    return a
}

/**
 * validationErrors
 */
func (a *ConditionBuilderControl) ValidationErrors(value interface{}) *ConditionBuilderControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * funcs
 */
func (a *ConditionBuilderControl) Funcs(value interface{}) *ConditionBuilderControl {
    a.Set("funcs", value)
    return a
}

/**
 * config
 */
func (a *ConditionBuilderControl) Config(value interface{}) *ConditionBuilderControl {
    a.Set("config", value)
    return a
}

/**
 * showANDOR
 */
func (a *ConditionBuilderControl) ShowANDOR(value interface{}) *ConditionBuilderControl {
    a.Set("showANDOR", value)
    return a
}

/**
 * addGroupBtnVisibleOn
 */
func (a *ConditionBuilderControl) AddGroupBtnVisibleOn(value interface{}) *ConditionBuilderControl {
    a.Set("addGroupBtnVisibleOn", value)
    return a
}

/**
 * horizontal
 */
func (a *ConditionBuilderControl) Horizontal(value interface{}) *ConditionBuilderControl {
    a.Set("horizontal", value)
    return a
}

/**
 * autoFill
 */
func (a *ConditionBuilderControl) AutoFill(value interface{}) *ConditionBuilderControl {
    a.Set("autoFill", value)
    return a
}

/**
 * row
 */
func (a *ConditionBuilderControl) Row(value interface{}) *ConditionBuilderControl {
    a.Set("row", value)
    return a
}

/**
 * pickerIcon
 */
func (a *ConditionBuilderControl) PickerIcon(value interface{}) *ConditionBuilderControl {
    a.Set("pickerIcon", value)
    return a
}
