package renderers


/**
 * Datetime日期时间选择控件 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/datetime

 * @author wcz0
 * @version 6.2.2
 */
type DateTimeControl struct {
	*BaseRenderer
}

func NewDateTimeControl() *DateTimeControl {
    a := &DateTimeControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-datetime")
    return a
}


func (a *DateTimeControl) Set(name string, value interface{}) *DateTimeControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * descriptionClassName
 */
func (a *DateTimeControl) DescriptionClassName(value interface{}) *DateTimeControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * disabled
 */
func (a *DateTimeControl) Disabled(value interface{}) *DateTimeControl {
    a.Set("disabled", value)
    return a
}

/**
 * hidden
 */
func (a *DateTimeControl) Hidden(value interface{}) *DateTimeControl {
    a.Set("hidden", value)
    return a
}

/**
 * id
 */
func (a *DateTimeControl) Id(value interface{}) *DateTimeControl {
    a.Set("id", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *DateTimeControl) StaticInputClassName(value interface{}) *DateTimeControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * validationErrors
 */
func (a *DateTimeControl) ValidationErrors(value interface{}) *DateTimeControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *DateTimeControl) ReadOnlyOn(value interface{}) *DateTimeControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * inline
 */
func (a *DateTimeControl) Inline(value interface{}) *DateTimeControl {
    a.Set("inline", value)
    return a
}

/**
 * visibleOn
 */
func (a *DateTimeControl) VisibleOn(value interface{}) *DateTimeControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * size
 */
func (a *DateTimeControl) Size(value interface{}) *DateTimeControl {
    a.Set("size", value)
    return a
}

/**
 * labelRemark
 */
func (a *DateTimeControl) LabelRemark(value interface{}) *DateTimeControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * initAutoFill
 */
func (a *DateTimeControl) InitAutoFill(value interface{}) *DateTimeControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * clearable
 */
func (a *DateTimeControl) Clearable(value interface{}) *DateTimeControl {
    a.Set("clearable", value)
    return a
}

/**
 * displayFormat
 */
func (a *DateTimeControl) DisplayFormat(value interface{}) *DateTimeControl {
    a.Set("displayFormat", value)
    return a
}

/**
 * timeConstraints
 */
func (a *DateTimeControl) TimeConstraints(value interface{}) *DateTimeControl {
    a.Set("timeConstraints", value)
    return a
}

/**
 * labelOverflow
 */
func (a *DateTimeControl) LabelOverflow(value interface{}) *DateTimeControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * hint
 */
func (a *DateTimeControl) Hint(value interface{}) *DateTimeControl {
    a.Set("hint", value)
    return a
}

/**
 * static
 */
func (a *DateTimeControl) Static(value interface{}) *DateTimeControl {
    a.Set("static", value)
    return a
}

/**
 * maxDate
 */
func (a *DateTimeControl) MaxDate(value interface{}) *DateTimeControl {
    a.Set("maxDate", value)
    return a
}

/**
 * extraName
 */
func (a *DateTimeControl) ExtraName(value interface{}) *DateTimeControl {
    a.Set("extraName", value)
    return a
}

/**
 * staticClassName
 */
func (a *DateTimeControl) StaticClassName(value interface{}) *DateTimeControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * useMobileUI
 */
func (a *DateTimeControl) UseMobileUI(value interface{}) *DateTimeControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * labelWidth
 */
func (a *DateTimeControl) LabelWidth(value interface{}) *DateTimeControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * readOnly
 */
func (a *DateTimeControl) ReadOnly(value interface{}) *DateTimeControl {
    a.Set("readOnly", value)
    return a
}

/**
 * shortcuts
 */
func (a *DateTimeControl) Shortcuts(value interface{}) *DateTimeControl {
    a.Set("shortcuts", value)
    return a
}

/**
 * remark
 */
func (a *DateTimeControl) Remark(value interface{}) *DateTimeControl {
    a.Set("remark", value)
    return a
}

/**
 * validateOnChange
 */
func (a *DateTimeControl) ValidateOnChange(value interface{}) *DateTimeControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *DateTimeControl) StaticPlaceholder(value interface{}) *DateTimeControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 */
func (a *DateTimeControl) Type(value interface{}) *DateTimeControl {
    a.Set("type", value)
    return a
}

/**
 * staticOn
 */
func (a *DateTimeControl) StaticOn(value interface{}) *DateTimeControl {
    a.Set("staticOn", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *DateTimeControl) StaticLabelClassName(value interface{}) *DateTimeControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * submitOnChange
 */
func (a *DateTimeControl) SubmitOnChange(value interface{}) *DateTimeControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * visible
 */
func (a *DateTimeControl) Visible(value interface{}) *DateTimeControl {
    a.Set("visible", value)
    return a
}

/**
 * editorSetting
 */
func (a *DateTimeControl) EditorSetting(value interface{}) *DateTimeControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * inputFormat
 */
func (a *DateTimeControl) InputFormat(value interface{}) *DateTimeControl {
    a.Set("inputFormat", value)
    return a
}

/**
 * isEndDate
 */
func (a *DateTimeControl) IsEndDate(value interface{}) *DateTimeControl {
    a.Set("isEndDate", value)
    return a
}

/**
 * labelAlign
 */
func (a *DateTimeControl) LabelAlign(value interface{}) *DateTimeControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * hiddenOn
 */
func (a *DateTimeControl) HiddenOn(value interface{}) *DateTimeControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * onEvent
 */
func (a *DateTimeControl) OnEvent(value interface{}) *DateTimeControl {
    a.Set("onEvent", value)
    return a
}

/**
 * valueFormat
 */
func (a *DateTimeControl) ValueFormat(value interface{}) *DateTimeControl {
    a.Set("valueFormat", value)
    return a
}

/**
 * inputForbid
 */
func (a *DateTimeControl) InputForbid(value interface{}) *DateTimeControl {
    a.Set("inputForbid", value)
    return a
}

/**
 * row
 */
func (a *DateTimeControl) Row(value interface{}) *DateTimeControl {
    a.Set("row", value)
    return a
}

/**
 * staticSchema
 */
func (a *DateTimeControl) StaticSchema(value interface{}) *DateTimeControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *DateTimeControl) Width(value interface{}) *DateTimeControl {
    a.Set("width", value)
    return a
}

/**
 * inputClassName
 */
func (a *DateTimeControl) InputClassName(value interface{}) *DateTimeControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * placeholder
 */
func (a *DateTimeControl) Placeholder(value interface{}) *DateTimeControl {
    a.Set("placeholder", value)
    return a
}

/**
 * required
 */
func (a *DateTimeControl) Required(value interface{}) *DateTimeControl {
    a.Set("required", value)
    return a
}

/**
 * name
 */
func (a *DateTimeControl) Name(value interface{}) *DateTimeControl {
    a.Set("name", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *DateTimeControl) ClearValueOnHidden(value interface{}) *DateTimeControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * style
 */
func (a *DateTimeControl) Style(value interface{}) *DateTimeControl {
    a.Set("style", value)
    return a
}

/**
 * disabledDate
 */
func (a *DateTimeControl) DisabledDate(value interface{}) *DateTimeControl {
    a.Set("disabledDate", value)
    return a
}

/**
 * timeFormat
 */
func (a *DateTimeControl) TimeFormat(value interface{}) *DateTimeControl {
    a.Set("timeFormat", value)
    return a
}

/**
 * validations
 */
func (a *DateTimeControl) Validations(value interface{}) *DateTimeControl {
    a.Set("validations", value)
    return a
}

/**
 * validateApi
 */
func (a *DateTimeControl) ValidateApi(value interface{}) *DateTimeControl {
    a.Set("validateApi", value)
    return a
}

/**
 * description
 */
func (a *DateTimeControl) Description(value interface{}) *DateTimeControl {
    a.Set("description", value)
    return a
}

/**
 * desc
 */
func (a *DateTimeControl) Desc(value interface{}) *DateTimeControl {
    a.Set("desc", value)
    return a
}

/**
 * horizontal
 */
func (a *DateTimeControl) Horizontal(value interface{}) *DateTimeControl {
    a.Set("horizontal", value)
    return a
}

/**
 * label
 */
func (a *DateTimeControl) Label(value interface{}) *DateTimeControl {
    a.Set("label", value)
    return a
}

/**
 * autoFill
 */
func (a *DateTimeControl) AutoFill(value interface{}) *DateTimeControl {
    a.Set("autoFill", value)
    return a
}

/**
 * disabledOn
 */
func (a *DateTimeControl) DisabledOn(value interface{}) *DateTimeControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * format
 */
func (a *DateTimeControl) Format(value interface{}) *DateTimeControl {
    a.Set("format", value)
    return a
}

/**
 * emebed
 */
func (a *DateTimeControl) Emebed(value interface{}) *DateTimeControl {
    a.Set("emebed", value)
    return a
}

/**
 * borderMode
 */
func (a *DateTimeControl) BorderMode(value interface{}) *DateTimeControl {
    a.Set("borderMode", value)
    return a
}

/**
 * labelClassName
 */
func (a *DateTimeControl) LabelClassName(value interface{}) *DateTimeControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * mode
 */
func (a *DateTimeControl) Mode(value interface{}) *DateTimeControl {
    a.Set("mode", value)
    return a
}

/**
 * value
 */
func (a *DateTimeControl) Value(value interface{}) *DateTimeControl {
    a.Set("value", value)
    return a
}

/**
 * className
 */
func (a *DateTimeControl) ClassName(value interface{}) *DateTimeControl {
    a.Set("className", value)
    return a
}

/**
 * utc
 */
func (a *DateTimeControl) Utc(value interface{}) *DateTimeControl {
    a.Set("utc", value)
    return a
}

/**
 * minDate
 */
func (a *DateTimeControl) MinDate(value interface{}) *DateTimeControl {
    a.Set("minDate", value)
    return a
}
