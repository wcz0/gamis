package renderers


/**
 * Month 月份选择控件 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/Month

 * @author wcz0
 * @version 6.2.2
 */
type MonthControl struct {
	*BaseRenderer
}

func NewMonthControl() *MonthControl {
    a := &MonthControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-month")
    return a
}


func (a *MonthControl) Set(name string, value interface{}) *MonthControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * extraName
 */
func (a *MonthControl) ExtraName(value interface{}) *MonthControl {
    a.Set("extraName", value)
    return a
}

/**
 * readOnly
 */
func (a *MonthControl) ReadOnly(value interface{}) *MonthControl {
    a.Set("readOnly", value)
    return a
}

/**
 * visibleOn
 */
func (a *MonthControl) VisibleOn(value interface{}) *MonthControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticClassName
 */
func (a *MonthControl) StaticClassName(value interface{}) *MonthControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *MonthControl) StaticInputClassName(value interface{}) *MonthControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * size
 */
func (a *MonthControl) Size(value interface{}) *MonthControl {
    a.Set("size", value)
    return a
}

/**
 */
func (a *MonthControl) Type(value interface{}) *MonthControl {
    a.Set("type", value)
    return a
}

/**
 * utc
 */
func (a *MonthControl) Utc(value interface{}) *MonthControl {
    a.Set("utc", value)
    return a
}

/**
 * borderMode
 */
func (a *MonthControl) BorderMode(value interface{}) *MonthControl {
    a.Set("borderMode", value)
    return a
}

/**
 * hiddenOn
 */
func (a *MonthControl) HiddenOn(value interface{}) *MonthControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * id
 */
func (a *MonthControl) Id(value interface{}) *MonthControl {
    a.Set("id", value)
    return a
}

/**
 * onEvent
 */
func (a *MonthControl) OnEvent(value interface{}) *MonthControl {
    a.Set("onEvent", value)
    return a
}

/**
 * submitOnChange
 */
func (a *MonthControl) SubmitOnChange(value interface{}) *MonthControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * labelAlign
 */
func (a *MonthControl) LabelAlign(value interface{}) *MonthControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * labelClassName
 */
func (a *MonthControl) LabelClassName(value interface{}) *MonthControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * remark
 */
func (a *MonthControl) Remark(value interface{}) *MonthControl {
    a.Set("remark", value)
    return a
}

/**
 * hint
 */
func (a *MonthControl) Hint(value interface{}) *MonthControl {
    a.Set("hint", value)
    return a
}

/**
 * horizontal
 */
func (a *MonthControl) Horizontal(value interface{}) *MonthControl {
    a.Set("horizontal", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *MonthControl) StaticLabelClassName(value interface{}) *MonthControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * label
 */
func (a *MonthControl) Label(value interface{}) *MonthControl {
    a.Set("label", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *MonthControl) DescriptionClassName(value interface{}) *MonthControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * initAutoFill
 */
func (a *MonthControl) InitAutoFill(value interface{}) *MonthControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * row
 */
func (a *MonthControl) Row(value interface{}) *MonthControl {
    a.Set("row", value)
    return a
}

/**
 * staticSchema
 */
func (a *MonthControl) StaticSchema(value interface{}) *MonthControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * inputFormat
 */
func (a *MonthControl) InputFormat(value interface{}) *MonthControl {
    a.Set("inputFormat", value)
    return a
}

/**
 * emebed
 */
func (a *MonthControl) Emebed(value interface{}) *MonthControl {
    a.Set("emebed", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *MonthControl) Width(value interface{}) *MonthControl {
    a.Set("width", value)
    return a
}

/**
 * labelOverflow
 */
func (a *MonthControl) LabelOverflow(value interface{}) *MonthControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *MonthControl) ReadOnlyOn(value interface{}) *MonthControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * placeholder
 */
func (a *MonthControl) Placeholder(value interface{}) *MonthControl {
    a.Set("placeholder", value)
    return a
}

/**
 * validationErrors
 */
func (a *MonthControl) ValidationErrors(value interface{}) *MonthControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * className
 */
func (a *MonthControl) ClassName(value interface{}) *MonthControl {
    a.Set("className", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *MonthControl) StaticPlaceholder(value interface{}) *MonthControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * disabled
 */
func (a *MonthControl) Disabled(value interface{}) *MonthControl {
    a.Set("disabled", value)
    return a
}

/**
 * displayFormat
 */
func (a *MonthControl) DisplayFormat(value interface{}) *MonthControl {
    a.Set("displayFormat", value)
    return a
}

/**
 * labelWidth
 */
func (a *MonthControl) LabelWidth(value interface{}) *MonthControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * labelRemark
 */
func (a *MonthControl) LabelRemark(value interface{}) *MonthControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * mode
 */
func (a *MonthControl) Mode(value interface{}) *MonthControl {
    a.Set("mode", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *MonthControl) ClearValueOnHidden(value interface{}) *MonthControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * autoFill
 */
func (a *MonthControl) AutoFill(value interface{}) *MonthControl {
    a.Set("autoFill", value)
    return a
}

/**
 * disabledOn
 */
func (a *MonthControl) DisabledOn(value interface{}) *MonthControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * description
 */
func (a *MonthControl) Description(value interface{}) *MonthControl {
    a.Set("description", value)
    return a
}

/**
 * value
 */
func (a *MonthControl) Value(value interface{}) *MonthControl {
    a.Set("value", value)
    return a
}

/**
 * name
 */
func (a *MonthControl) Name(value interface{}) *MonthControl {
    a.Set("name", value)
    return a
}

/**
 * inputClassName
 */
func (a *MonthControl) InputClassName(value interface{}) *MonthControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * static
 */
func (a *MonthControl) Static(value interface{}) *MonthControl {
    a.Set("static", value)
    return a
}

/**
 * staticOn
 */
func (a *MonthControl) StaticOn(value interface{}) *MonthControl {
    a.Set("staticOn", value)
    return a
}

/**
 * editorSetting
 */
func (a *MonthControl) EditorSetting(value interface{}) *MonthControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * valueFormat
 */
func (a *MonthControl) ValueFormat(value interface{}) *MonthControl {
    a.Set("valueFormat", value)
    return a
}

/**
 * shortcuts
 */
func (a *MonthControl) Shortcuts(value interface{}) *MonthControl {
    a.Set("shortcuts", value)
    return a
}

/**
 * clearable
 */
func (a *MonthControl) Clearable(value interface{}) *MonthControl {
    a.Set("clearable", value)
    return a
}

/**
 * disabledDate
 */
func (a *MonthControl) DisabledDate(value interface{}) *MonthControl {
    a.Set("disabledDate", value)
    return a
}

/**
 * inputForbid
 */
func (a *MonthControl) InputForbid(value interface{}) *MonthControl {
    a.Set("inputForbid", value)
    return a
}

/**
 * validateOnChange
 */
func (a *MonthControl) ValidateOnChange(value interface{}) *MonthControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * inline
 */
func (a *MonthControl) Inline(value interface{}) *MonthControl {
    a.Set("inline", value)
    return a
}

/**
 * hidden
 */
func (a *MonthControl) Hidden(value interface{}) *MonthControl {
    a.Set("hidden", value)
    return a
}

/**
 * required
 */
func (a *MonthControl) Required(value interface{}) *MonthControl {
    a.Set("required", value)
    return a
}

/**
 * validateApi
 */
func (a *MonthControl) ValidateApi(value interface{}) *MonthControl {
    a.Set("validateApi", value)
    return a
}

/**
 * visible
 */
func (a *MonthControl) Visible(value interface{}) *MonthControl {
    a.Set("visible", value)
    return a
}

/**
 * style
 */
func (a *MonthControl) Style(value interface{}) *MonthControl {
    a.Set("style", value)
    return a
}

/**
 * useMobileUI
 */
func (a *MonthControl) UseMobileUI(value interface{}) *MonthControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * desc
 */
func (a *MonthControl) Desc(value interface{}) *MonthControl {
    a.Set("desc", value)
    return a
}

/**
 * validations
 */
func (a *MonthControl) Validations(value interface{}) *MonthControl {
    a.Set("validations", value)
    return a
}

/**
 * format
 */
func (a *MonthControl) Format(value interface{}) *MonthControl {
    a.Set("format", value)
    return a
}
