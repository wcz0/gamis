package renderers


/**
 * DateRange 日期范围控件 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/date-range

 * @author wcz0
 * @version 6.2.2
 */
type DateRangeControl struct {
	*BaseRenderer
}

func NewDateRangeControl() *DateRangeControl {
    a := &DateRangeControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-date-range")
    return a
}


func (a *DateRangeControl) Set(name string, value interface{}) *DateRangeControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * editorSetting
 */
func (a *DateRangeControl) EditorSetting(value interface{}) *DateRangeControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * validations
 */
func (a *DateRangeControl) Validations(value interface{}) *DateRangeControl {
    a.Set("validations", value)
    return a
}

/**
 * className
 */
func (a *DateRangeControl) ClassName(value interface{}) *DateRangeControl {
    a.Set("className", value)
    return a
}

/**
 * onEvent
 */
func (a *DateRangeControl) OnEvent(value interface{}) *DateRangeControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 可选值: input-date-range | input-datetime-range | input-time-range
 */
func (a *DateRangeControl) Type(value interface{}) *DateRangeControl {
    a.Set("type", value)
    return a
}

/**
 * shortcuts
 */
func (a *DateRangeControl) Shortcuts(value interface{}) *DateRangeControl {
    a.Set("shortcuts", value)
    return a
}

/**
 * remark
 */
func (a *DateRangeControl) Remark(value interface{}) *DateRangeControl {
    a.Set("remark", value)
    return a
}

/**
 * hint
 */
func (a *DateRangeControl) Hint(value interface{}) *DateRangeControl {
    a.Set("hint", value)
    return a
}

/**
 * initAutoFill
 */
func (a *DateRangeControl) InitAutoFill(value interface{}) *DateRangeControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * row
 */
func (a *DateRangeControl) Row(value interface{}) *DateRangeControl {
    a.Set("row", value)
    return a
}

/**
 * hiddenOn
 */
func (a *DateRangeControl) HiddenOn(value interface{}) *DateRangeControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * ranges
 */
func (a *DateRangeControl) Ranges(value interface{}) *DateRangeControl {
    a.Set("ranges", value)
    return a
}

/**
 * placeholder
 */
func (a *DateRangeControl) Placeholder(value interface{}) *DateRangeControl {
    a.Set("placeholder", value)
    return a
}

/**
 * useMobileUI
 */
func (a *DateRangeControl) UseMobileUI(value interface{}) *DateRangeControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * startPlaceholder
 */
func (a *DateRangeControl) StartPlaceholder(value interface{}) *DateRangeControl {
    a.Set("startPlaceholder", value)
    return a
}

/**
 * endPlaceholder
 */
func (a *DateRangeControl) EndPlaceholder(value interface{}) *DateRangeControl {
    a.Set("endPlaceholder", value)
    return a
}

/**
 * animation
 */
func (a *DateRangeControl) Animation(value interface{}) *DateRangeControl {
    a.Set("animation", value)
    return a
}

/**
 * transform
 */
func (a *DateRangeControl) Transform(value interface{}) *DateRangeControl {
    a.Set("transform", value)
    return a
}

/**
 * visible
 */
func (a *DateRangeControl) Visible(value interface{}) *DateRangeControl {
    a.Set("visible", value)
    return a
}

/**
 * maxDate
 */
func (a *DateRangeControl) MaxDate(value interface{}) *DateRangeControl {
    a.Set("maxDate", value)
    return a
}

/**
 * minDate
 */
func (a *DateRangeControl) MinDate(value interface{}) *DateRangeControl {
    a.Set("minDate", value)
    return a
}

/**
 * label
 */
func (a *DateRangeControl) Label(value interface{}) *DateRangeControl {
    a.Set("label", value)
    return a
}

/**
 * labelWidth
 */
func (a *DateRangeControl) LabelWidth(value interface{}) *DateRangeControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * labelClassName
 */
func (a *DateRangeControl) LabelClassName(value interface{}) *DateRangeControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * validateOnChange
 */
func (a *DateRangeControl) ValidateOnChange(value interface{}) *DateRangeControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *DateRangeControl) DescriptionClassName(value interface{}) *DateRangeControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * static
 */
func (a *DateRangeControl) Static(value interface{}) *DateRangeControl {
    a.Set("static", value)
    return a
}

/**
 * staticSchema
 */
func (a *DateRangeControl) StaticSchema(value interface{}) *DateRangeControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * displayFormat
 */
func (a *DateRangeControl) DisplayFormat(value interface{}) *DateRangeControl {
    a.Set("displayFormat", value)
    return a
}

/**
 * mode
 */
func (a *DateRangeControl) Mode(value interface{}) *DateRangeControl {
    a.Set("mode", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *DateRangeControl) ClearValueOnHidden(value interface{}) *DateRangeControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * disabledOn
 */
func (a *DateRangeControl) DisabledOn(value interface{}) *DateRangeControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * inputFormat
 */
func (a *DateRangeControl) InputFormat(value interface{}) *DateRangeControl {
    a.Set("inputFormat", value)
    return a
}

/**
 * embed
 */
func (a *DateRangeControl) Embed(value interface{}) *DateRangeControl {
    a.Set("embed", value)
    return a
}

/**
 * staticOn
 */
func (a *DateRangeControl) StaticOn(value interface{}) *DateRangeControl {
    a.Set("staticOn", value)
    return a
}

/**
 * style
 */
func (a *DateRangeControl) Style(value interface{}) *DateRangeControl {
    a.Set("style", value)
    return a
}

/**
 * maxDuration
 */
func (a *DateRangeControl) MaxDuration(value interface{}) *DateRangeControl {
    a.Set("maxDuration", value)
    return a
}

/**
 * labelAlign
 */
func (a *DateRangeControl) LabelAlign(value interface{}) *DateRangeControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * labelOverflow
 */
func (a *DateRangeControl) LabelOverflow(value interface{}) *DateRangeControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * inline
 */
func (a *DateRangeControl) Inline(value interface{}) *DateRangeControl {
    a.Set("inline", value)
    return a
}

/**
 * disabled
 */
func (a *DateRangeControl) Disabled(value interface{}) *DateRangeControl {
    a.Set("disabled", value)
    return a
}

/**
 * minDuration
 */
func (a *DateRangeControl) MinDuration(value interface{}) *DateRangeControl {
    a.Set("minDuration", value)
    return a
}

/**
 * popOverContainerSelector
 */
func (a *DateRangeControl) PopOverContainerSelector(value interface{}) *DateRangeControl {
    a.Set("popOverContainerSelector", value)
    return a
}

/**
 * labelRemark
 */
func (a *DateRangeControl) LabelRemark(value interface{}) *DateRangeControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * value
 */
func (a *DateRangeControl) Value(value interface{}) *DateRangeControl {
    a.Set("value", value)
    return a
}

/**
 * visibleOn
 */
func (a *DateRangeControl) VisibleOn(value interface{}) *DateRangeControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * id
 */
func (a *DateRangeControl) Id(value interface{}) *DateRangeControl {
    a.Set("id", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *DateRangeControl) StaticPlaceholder(value interface{}) *DateRangeControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticClassName
 */
func (a *DateRangeControl) StaticClassName(value interface{}) *DateRangeControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * extraName
 */
func (a *DateRangeControl) ExtraName(value interface{}) *DateRangeControl {
    a.Set("extraName", value)
    return a
}

/**
 * inputClassName
 */
func (a *DateRangeControl) InputClassName(value interface{}) *DateRangeControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * validateApi
 */
func (a *DateRangeControl) ValidateApi(value interface{}) *DateRangeControl {
    a.Set("validateApi", value)
    return a
}

/**
 * name
 */
func (a *DateRangeControl) Name(value interface{}) *DateRangeControl {
    a.Set("name", value)
    return a
}

/**
 * readOnly
 */
func (a *DateRangeControl) ReadOnly(value interface{}) *DateRangeControl {
    a.Set("readOnly", value)
    return a
}

/**
 * delimiter
 */
func (a *DateRangeControl) Delimiter(value interface{}) *DateRangeControl {
    a.Set("delimiter", value)
    return a
}

/**
 * format
 */
func (a *DateRangeControl) Format(value interface{}) *DateRangeControl {
    a.Set("format", value)
    return a
}

/**
 * borderMode
 */
func (a *DateRangeControl) BorderMode(value interface{}) *DateRangeControl {
    a.Set("borderMode", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *DateRangeControl) Width(value interface{}) *DateRangeControl {
    a.Set("width", value)
    return a
}

/**
 * submitOnChange
 */
func (a *DateRangeControl) SubmitOnChange(value interface{}) *DateRangeControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * required
 */
func (a *DateRangeControl) Required(value interface{}) *DateRangeControl {
    a.Set("required", value)
    return a
}

/**
 * autoFill
 */
func (a *DateRangeControl) AutoFill(value interface{}) *DateRangeControl {
    a.Set("autoFill", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *DateRangeControl) StaticLabelClassName(value interface{}) *DateRangeControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *DateRangeControl) StaticInputClassName(value interface{}) *DateRangeControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * valueFormat
 */
func (a *DateRangeControl) ValueFormat(value interface{}) *DateRangeControl {
    a.Set("valueFormat", value)
    return a
}

/**
 * joinValues
 */
func (a *DateRangeControl) JoinValues(value interface{}) *DateRangeControl {
    a.Set("joinValues", value)
    return a
}

/**
 * size
 */
func (a *DateRangeControl) Size(value interface{}) *DateRangeControl {
    a.Set("size", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *DateRangeControl) ReadOnlyOn(value interface{}) *DateRangeControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * description
 */
func (a *DateRangeControl) Description(value interface{}) *DateRangeControl {
    a.Set("description", value)
    return a
}

/**
 * desc
 */
func (a *DateRangeControl) Desc(value interface{}) *DateRangeControl {
    a.Set("desc", value)
    return a
}

/**
 * horizontal
 */
func (a *DateRangeControl) Horizontal(value interface{}) *DateRangeControl {
    a.Set("horizontal", value)
    return a
}

/**
 * validationErrors
 */
func (a *DateRangeControl) ValidationErrors(value interface{}) *DateRangeControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * hidden
 */
func (a *DateRangeControl) Hidden(value interface{}) *DateRangeControl {
    a.Set("hidden", value)
    return a
}
