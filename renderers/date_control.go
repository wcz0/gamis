package renderers


/**
 * Date日期选择控件 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/date

 * @author wcz0
 * @version 6.2.2
 */
type DateControl struct {
	*BaseRenderer
}

func NewDateControl() *DateControl {
    a := &DateControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-date")
    return a
}


func (a *DateControl) Set(name string, value interface{}) *DateControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * validateOnChange
 */
func (a *DateControl) ValidateOnChange(value interface{}) *DateControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * placeholder
 */
func (a *DateControl) Placeholder(value interface{}) *DateControl {
    a.Set("placeholder", value)
    return a
}

/**
 * className
 */
func (a *DateControl) ClassName(value interface{}) *DateControl {
    a.Set("className", value)
    return a
}

/**
 * staticClassName
 */
func (a *DateControl) StaticClassName(value interface{}) *DateControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *DateControl) StaticLabelClassName(value interface{}) *DateControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * format
 */
func (a *DateControl) Format(value interface{}) *DateControl {
    a.Set("format", value)
    return a
}

/**
 * disabledDate
 */
func (a *DateControl) DisabledDate(value interface{}) *DateControl {
    a.Set("disabledDate", value)
    return a
}

/**
 * label
 */
func (a *DateControl) Label(value interface{}) *DateControl {
    a.Set("label", value)
    return a
}

/**
 * labelClassName
 */
func (a *DateControl) LabelClassName(value interface{}) *DateControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * desc
 */
func (a *DateControl) Desc(value interface{}) *DateControl {
    a.Set("desc", value)
    return a
}

/**
 * mode
 */
func (a *DateControl) Mode(value interface{}) *DateControl {
    a.Set("mode", value)
    return a
}

/**
 * validationErrors
 */
func (a *DateControl) ValidationErrors(value interface{}) *DateControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * useMobileUI
 */
func (a *DateControl) UseMobileUI(value interface{}) *DateControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * emebed
 */
func (a *DateControl) Emebed(value interface{}) *DateControl {
    a.Set("emebed", value)
    return a
}

/**
 * inputForbid
 */
func (a *DateControl) InputForbid(value interface{}) *DateControl {
    a.Set("inputForbid", value)
    return a
}

/**
 * popOverContainerSelector
 */
func (a *DateControl) PopOverContainerSelector(value interface{}) *DateControl {
    a.Set("popOverContainerSelector", value)
    return a
}

/**
 * labelRemark
 */
func (a *DateControl) LabelRemark(value interface{}) *DateControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * validations
 */
func (a *DateControl) Validations(value interface{}) *DateControl {
    a.Set("validations", value)
    return a
}

/**
 * static
 */
func (a *DateControl) Static(value interface{}) *DateControl {
    a.Set("static", value)
    return a
}

/**
 */
func (a *DateControl) Type(value interface{}) *DateControl {
    a.Set("type", value)
    return a
}

/**
 * labelWidth
 */
func (a *DateControl) LabelWidth(value interface{}) *DateControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * autoFill
 */
func (a *DateControl) AutoFill(value interface{}) *DateControl {
    a.Set("autoFill", value)
    return a
}

/**
 * staticOn
 */
func (a *DateControl) StaticOn(value interface{}) *DateControl {
    a.Set("staticOn", value)
    return a
}

/**
 * clearable
 */
func (a *DateControl) Clearable(value interface{}) *DateControl {
    a.Set("clearable", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *DateControl) ReadOnlyOn(value interface{}) *DateControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *DateControl) DescriptionClassName(value interface{}) *DateControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * row
 */
func (a *DateControl) Row(value interface{}) *DateControl {
    a.Set("row", value)
    return a
}

/**
 * hidden
 */
func (a *DateControl) Hidden(value interface{}) *DateControl {
    a.Set("hidden", value)
    return a
}

/**
 * id
 */
func (a *DateControl) Id(value interface{}) *DateControl {
    a.Set("id", value)
    return a
}

/**
 * staticSchema
 */
func (a *DateControl) StaticSchema(value interface{}) *DateControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * style
 */
func (a *DateControl) Style(value interface{}) *DateControl {
    a.Set("style", value)
    return a
}

/**
 * size
 */
func (a *DateControl) Size(value interface{}) *DateControl {
    a.Set("size", value)
    return a
}

/**
 * description
 */
func (a *DateControl) Description(value interface{}) *DateControl {
    a.Set("description", value)
    return a
}

/**
 * disabled
 */
func (a *DateControl) Disabled(value interface{}) *DateControl {
    a.Set("disabled", value)
    return a
}

/**
 * inputFormat
 */
func (a *DateControl) InputFormat(value interface{}) *DateControl {
    a.Set("inputFormat", value)
    return a
}

/**
 * displayFormat
 */
func (a *DateControl) DisplayFormat(value interface{}) *DateControl {
    a.Set("displayFormat", value)
    return a
}

/**
 * shortcuts
 */
func (a *DateControl) Shortcuts(value interface{}) *DateControl {
    a.Set("shortcuts", value)
    return a
}

/**
 * closeOnSelect
 */
func (a *DateControl) CloseOnSelect(value interface{}) *DateControl {
    a.Set("closeOnSelect", value)
    return a
}

/**
 * labelOverflow
 */
func (a *DateControl) LabelOverflow(value interface{}) *DateControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * name
 */
func (a *DateControl) Name(value interface{}) *DateControl {
    a.Set("name", value)
    return a
}

/**
 * extraName
 */
func (a *DateControl) ExtraName(value interface{}) *DateControl {
    a.Set("extraName", value)
    return a
}

/**
 * readOnly
 */
func (a *DateControl) ReadOnly(value interface{}) *DateControl {
    a.Set("readOnly", value)
    return a
}

/**
 * validateApi
 */
func (a *DateControl) ValidateApi(value interface{}) *DateControl {
    a.Set("validateApi", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *DateControl) StaticPlaceholder(value interface{}) *DateControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * editorSetting
 */
func (a *DateControl) EditorSetting(value interface{}) *DateControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * horizontal
 */
func (a *DateControl) Horizontal(value interface{}) *DateControl {
    a.Set("horizontal", value)
    return a
}

/**
 * utc
 */
func (a *DateControl) Utc(value interface{}) *DateControl {
    a.Set("utc", value)
    return a
}

/**
 * labelAlign
 */
func (a *DateControl) LabelAlign(value interface{}) *DateControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * submitOnChange
 */
func (a *DateControl) SubmitOnChange(value interface{}) *DateControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * inline
 */
func (a *DateControl) Inline(value interface{}) *DateControl {
    a.Set("inline", value)
    return a
}

/**
 * required
 */
func (a *DateControl) Required(value interface{}) *DateControl {
    a.Set("required", value)
    return a
}

/**
 * value
 */
func (a *DateControl) Value(value interface{}) *DateControl {
    a.Set("value", value)
    return a
}

/**
 * hiddenOn
 */
func (a *DateControl) HiddenOn(value interface{}) *DateControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * valueFormat
 */
func (a *DateControl) ValueFormat(value interface{}) *DateControl {
    a.Set("valueFormat", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *DateControl) Width(value interface{}) *DateControl {
    a.Set("width", value)
    return a
}

/**
 * visible
 */
func (a *DateControl) Visible(value interface{}) *DateControl {
    a.Set("visible", value)
    return a
}

/**
 * visibleOn
 */
func (a *DateControl) VisibleOn(value interface{}) *DateControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * borderMode
 */
func (a *DateControl) BorderMode(value interface{}) *DateControl {
    a.Set("borderMode", value)
    return a
}

/**
 * remark
 */
func (a *DateControl) Remark(value interface{}) *DateControl {
    a.Set("remark", value)
    return a
}

/**
 * hint
 */
func (a *DateControl) Hint(value interface{}) *DateControl {
    a.Set("hint", value)
    return a
}

/**
 * minDate
 */
func (a *DateControl) MinDate(value interface{}) *DateControl {
    a.Set("minDate", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *DateControl) ClearValueOnHidden(value interface{}) *DateControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * onEvent
 */
func (a *DateControl) OnEvent(value interface{}) *DateControl {
    a.Set("onEvent", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *DateControl) StaticInputClassName(value interface{}) *DateControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * maxDate
 */
func (a *DateControl) MaxDate(value interface{}) *DateControl {
    a.Set("maxDate", value)
    return a
}

/**
 * disabledOn
 */
func (a *DateControl) DisabledOn(value interface{}) *DateControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * inputClassName
 */
func (a *DateControl) InputClassName(value interface{}) *DateControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * initAutoFill
 */
func (a *DateControl) InitAutoFill(value interface{}) *DateControl {
    a.Set("initAutoFill", value)
    return a
}
