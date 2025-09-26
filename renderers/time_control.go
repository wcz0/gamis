package renderers


/**
 * Time 时间选择控件 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/time

 * @author wcz0
 * @version 6.2.2
 */
type TimeControl struct {
	*BaseRenderer
}

func NewTimeControl() *TimeControl {
    a := &TimeControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-time")
    return a
}


func (a *TimeControl) Set(name string, value interface{}) *TimeControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * onEvent
 */
func (a *TimeControl) OnEvent(value interface{}) *TimeControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *TimeControl) Width(value interface{}) *TimeControl {
    a.Set("width", value)
    return a
}

/**
 * useMobileUI
 */
func (a *TimeControl) UseMobileUI(value interface{}) *TimeControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * size
 */
func (a *TimeControl) Size(value interface{}) *TimeControl {
    a.Set("size", value)
    return a
}

/**
 */
func (a *TimeControl) Type(value interface{}) *TimeControl {
    a.Set("type", value)
    return a
}

/**
 * clearable
 */
func (a *TimeControl) Clearable(value interface{}) *TimeControl {
    a.Set("clearable", value)
    return a
}

/**
 * staticClassName
 */
func (a *TimeControl) StaticClassName(value interface{}) *TimeControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *TimeControl) StaticInputClassName(value interface{}) *TimeControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * labelAlign
 */
func (a *TimeControl) LabelAlign(value interface{}) *TimeControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * remark
 */
func (a *TimeControl) Remark(value interface{}) *TimeControl {
    a.Set("remark", value)
    return a
}

/**
 * validations
 */
func (a *TimeControl) Validations(value interface{}) *TimeControl {
    a.Set("validations", value)
    return a
}

/**
 * timeFormat
 */
func (a *TimeControl) TimeFormat(value interface{}) *TimeControl {
    a.Set("timeFormat", value)
    return a
}

/**
 * mode
 */
func (a *TimeControl) Mode(value interface{}) *TimeControl {
    a.Set("mode", value)
    return a
}

/**
 * autoFill
 */
func (a *TimeControl) AutoFill(value interface{}) *TimeControl {
    a.Set("autoFill", value)
    return a
}

/**
 * className
 */
func (a *TimeControl) ClassName(value interface{}) *TimeControl {
    a.Set("className", value)
    return a
}

/**
 * visible
 */
func (a *TimeControl) Visible(value interface{}) *TimeControl {
    a.Set("visible", value)
    return a
}

/**
 * static
 */
func (a *TimeControl) Static(value interface{}) *TimeControl {
    a.Set("static", value)
    return a
}

/**
 * staticSchema
 */
func (a *TimeControl) StaticSchema(value interface{}) *TimeControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * format
 */
func (a *TimeControl) Format(value interface{}) *TimeControl {
    a.Set("format", value)
    return a
}

/**
 * description
 */
func (a *TimeControl) Description(value interface{}) *TimeControl {
    a.Set("description", value)
    return a
}

/**
 * style
 */
func (a *TimeControl) Style(value interface{}) *TimeControl {
    a.Set("style", value)
    return a
}

/**
 * editorSetting
 */
func (a *TimeControl) EditorSetting(value interface{}) *TimeControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * inputFormat
 */
func (a *TimeControl) InputFormat(value interface{}) *TimeControl {
    a.Set("inputFormat", value)
    return a
}

/**
 * displayFormat
 */
func (a *TimeControl) DisplayFormat(value interface{}) *TimeControl {
    a.Set("displayFormat", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *TimeControl) ReadOnlyOn(value interface{}) *TimeControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * value
 */
func (a *TimeControl) Value(value interface{}) *TimeControl {
    a.Set("value", value)
    return a
}

/**
 * row
 */
func (a *TimeControl) Row(value interface{}) *TimeControl {
    a.Set("row", value)
    return a
}

/**
 * timeConstraints
 */
func (a *TimeControl) TimeConstraints(value interface{}) *TimeControl {
    a.Set("timeConstraints", value)
    return a
}

/**
 * labelWidth
 */
func (a *TimeControl) LabelWidth(value interface{}) *TimeControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * labelRemark
 */
func (a *TimeControl) LabelRemark(value interface{}) *TimeControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * validationErrors
 */
func (a *TimeControl) ValidationErrors(value interface{}) *TimeControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * hidden
 */
func (a *TimeControl) Hidden(value interface{}) *TimeControl {
    a.Set("hidden", value)
    return a
}

/**
 * id
 */
func (a *TimeControl) Id(value interface{}) *TimeControl {
    a.Set("id", value)
    return a
}

/**
 * valueFormat
 */
func (a *TimeControl) ValueFormat(value interface{}) *TimeControl {
    a.Set("valueFormat", value)
    return a
}

/**
 * utc
 */
func (a *TimeControl) Utc(value interface{}) *TimeControl {
    a.Set("utc", value)
    return a
}

/**
 * borderMode
 */
func (a *TimeControl) BorderMode(value interface{}) *TimeControl {
    a.Set("borderMode", value)
    return a
}

/**
 * labelClassName
 */
func (a *TimeControl) LabelClassName(value interface{}) *TimeControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * validateApi
 */
func (a *TimeControl) ValidateApi(value interface{}) *TimeControl {
    a.Set("validateApi", value)
    return a
}

/**
 * shortcuts
 */
func (a *TimeControl) Shortcuts(value interface{}) *TimeControl {
    a.Set("shortcuts", value)
    return a
}

/**
 * hint
 */
func (a *TimeControl) Hint(value interface{}) *TimeControl {
    a.Set("hint", value)
    return a
}

/**
 * visibleOn
 */
func (a *TimeControl) VisibleOn(value interface{}) *TimeControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * labelOverflow
 */
func (a *TimeControl) LabelOverflow(value interface{}) *TimeControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *TimeControl) DescriptionClassName(value interface{}) *TimeControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * horizontal
 */
func (a *TimeControl) Horizontal(value interface{}) *TimeControl {
    a.Set("horizontal", value)
    return a
}

/**
 * placeholder
 */
func (a *TimeControl) Placeholder(value interface{}) *TimeControl {
    a.Set("placeholder", value)
    return a
}

/**
 * staticOn
 */
func (a *TimeControl) StaticOn(value interface{}) *TimeControl {
    a.Set("staticOn", value)
    return a
}

/**
 * disabledDate
 */
func (a *TimeControl) DisabledDate(value interface{}) *TimeControl {
    a.Set("disabledDate", value)
    return a
}

/**
 * initAutoFill
 */
func (a *TimeControl) InitAutoFill(value interface{}) *TimeControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * emebed
 */
func (a *TimeControl) Emebed(value interface{}) *TimeControl {
    a.Set("emebed", value)
    return a
}

/**
 * label
 */
func (a *TimeControl) Label(value interface{}) *TimeControl {
    a.Set("label", value)
    return a
}

/**
 * readOnly
 */
func (a *TimeControl) ReadOnly(value interface{}) *TimeControl {
    a.Set("readOnly", value)
    return a
}

/**
 * validateOnChange
 */
func (a *TimeControl) ValidateOnChange(value interface{}) *TimeControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * inputClassName
 */
func (a *TimeControl) InputClassName(value interface{}) *TimeControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *TimeControl) StaticLabelClassName(value interface{}) *TimeControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * extraName
 */
func (a *TimeControl) ExtraName(value interface{}) *TimeControl {
    a.Set("extraName", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *TimeControl) ClearValueOnHidden(value interface{}) *TimeControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * disabled
 */
func (a *TimeControl) Disabled(value interface{}) *TimeControl {
    a.Set("disabled", value)
    return a
}

/**
 * inputForbid
 */
func (a *TimeControl) InputForbid(value interface{}) *TimeControl {
    a.Set("inputForbid", value)
    return a
}

/**
 * submitOnChange
 */
func (a *TimeControl) SubmitOnChange(value interface{}) *TimeControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * desc
 */
func (a *TimeControl) Desc(value interface{}) *TimeControl {
    a.Set("desc", value)
    return a
}

/**
 * inline
 */
func (a *TimeControl) Inline(value interface{}) *TimeControl {
    a.Set("inline", value)
    return a
}

/**
 * name
 */
func (a *TimeControl) Name(value interface{}) *TimeControl {
    a.Set("name", value)
    return a
}

/**
 * disabledOn
 */
func (a *TimeControl) DisabledOn(value interface{}) *TimeControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *TimeControl) StaticPlaceholder(value interface{}) *TimeControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * required
 */
func (a *TimeControl) Required(value interface{}) *TimeControl {
    a.Set("required", value)
    return a
}

/**
 * hiddenOn
 */
func (a *TimeControl) HiddenOn(value interface{}) *TimeControl {
    a.Set("hiddenOn", value)
    return a
}
