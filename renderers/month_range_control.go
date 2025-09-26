package renderers


/**
 * MonthRange 月范围控件 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/month-range

 * @author wcz0
 * @version 6.2.2
 */
type MonthRangeControl struct {
	*BaseRenderer
}

func NewMonthRangeControl() *MonthRangeControl {
    a := &MonthRangeControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-month-range")
    return a
}


func (a *MonthRangeControl) Set(name string, value interface{}) *MonthRangeControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * labelWidth
 */
func (a *MonthRangeControl) LabelWidth(value interface{}) *MonthRangeControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * remark
 */
func (a *MonthRangeControl) Remark(value interface{}) *MonthRangeControl {
    a.Set("remark", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *MonthRangeControl) ReadOnlyOn(value interface{}) *MonthRangeControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *MonthRangeControl) StaticPlaceholder(value interface{}) *MonthRangeControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * format
 */
func (a *MonthRangeControl) Format(value interface{}) *MonthRangeControl {
    a.Set("format", value)
    return a
}

/**
 * inputFormat
 */
func (a *MonthRangeControl) InputFormat(value interface{}) *MonthRangeControl {
    a.Set("inputFormat", value)
    return a
}

/**
 * readOnly
 */
func (a *MonthRangeControl) ReadOnly(value interface{}) *MonthRangeControl {
    a.Set("readOnly", value)
    return a
}

/**
 * desc
 */
func (a *MonthRangeControl) Desc(value interface{}) *MonthRangeControl {
    a.Set("desc", value)
    return a
}

/**
 * displayFormat
 */
func (a *MonthRangeControl) DisplayFormat(value interface{}) *MonthRangeControl {
    a.Set("displayFormat", value)
    return a
}

/**
 * transform
 */
func (a *MonthRangeControl) Transform(value interface{}) *MonthRangeControl {
    a.Set("transform", value)
    return a
}

/**
 * row
 */
func (a *MonthRangeControl) Row(value interface{}) *MonthRangeControl {
    a.Set("row", value)
    return a
}

/**
 * joinValues
 */
func (a *MonthRangeControl) JoinValues(value interface{}) *MonthRangeControl {
    a.Set("joinValues", value)
    return a
}

/**
 * shortcuts
 */
func (a *MonthRangeControl) Shortcuts(value interface{}) *MonthRangeControl {
    a.Set("shortcuts", value)
    return a
}

/**
 * extraName
 */
func (a *MonthRangeControl) ExtraName(value interface{}) *MonthRangeControl {
    a.Set("extraName", value)
    return a
}

/**
 * validateApi
 */
func (a *MonthRangeControl) ValidateApi(value interface{}) *MonthRangeControl {
    a.Set("validateApi", value)
    return a
}

/**
 * autoFill
 */
func (a *MonthRangeControl) AutoFill(value interface{}) *MonthRangeControl {
    a.Set("autoFill", value)
    return a
}

/**
 * disabledOn
 */
func (a *MonthRangeControl) DisabledOn(value interface{}) *MonthRangeControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * hiddenOn
 */
func (a *MonthRangeControl) HiddenOn(value interface{}) *MonthRangeControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * size
 */
func (a *MonthRangeControl) Size(value interface{}) *MonthRangeControl {
    a.Set("size", value)
    return a
}

/**
 * animation
 */
func (a *MonthRangeControl) Animation(value interface{}) *MonthRangeControl {
    a.Set("animation", value)
    return a
}

/**
 * name
 */
func (a *MonthRangeControl) Name(value interface{}) *MonthRangeControl {
    a.Set("name", value)
    return a
}

/**
 * inputClassName
 */
func (a *MonthRangeControl) InputClassName(value interface{}) *MonthRangeControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * useMobileUI
 */
func (a *MonthRangeControl) UseMobileUI(value interface{}) *MonthRangeControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * valueFormat
 */
func (a *MonthRangeControl) ValueFormat(value interface{}) *MonthRangeControl {
    a.Set("valueFormat", value)
    return a
}

/**
 * inline
 */
func (a *MonthRangeControl) Inline(value interface{}) *MonthRangeControl {
    a.Set("inline", value)
    return a
}

/**
 * required
 */
func (a *MonthRangeControl) Required(value interface{}) *MonthRangeControl {
    a.Set("required", value)
    return a
}

/**
 * id
 */
func (a *MonthRangeControl) Id(value interface{}) *MonthRangeControl {
    a.Set("id", value)
    return a
}

/**
 * staticClassName
 */
func (a *MonthRangeControl) StaticClassName(value interface{}) *MonthRangeControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *MonthRangeControl) StaticSchema(value interface{}) *MonthRangeControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * ranges
 */
func (a *MonthRangeControl) Ranges(value interface{}) *MonthRangeControl {
    a.Set("ranges", value)
    return a
}

/**
 * labelRemark
 */
func (a *MonthRangeControl) LabelRemark(value interface{}) *MonthRangeControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * visibleOn
 */
func (a *MonthRangeControl) VisibleOn(value interface{}) *MonthRangeControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * hint
 */
func (a *MonthRangeControl) Hint(value interface{}) *MonthRangeControl {
    a.Set("hint", value)
    return a
}

/**
 * visible
 */
func (a *MonthRangeControl) Visible(value interface{}) *MonthRangeControl {
    a.Set("visible", value)
    return a
}

/**
 * minDuration
 */
func (a *MonthRangeControl) MinDuration(value interface{}) *MonthRangeControl {
    a.Set("minDuration", value)
    return a
}

/**
 * popOverContainerSelector
 */
func (a *MonthRangeControl) PopOverContainerSelector(value interface{}) *MonthRangeControl {
    a.Set("popOverContainerSelector", value)
    return a
}

/**
 * mode
 */
func (a *MonthRangeControl) Mode(value interface{}) *MonthRangeControl {
    a.Set("mode", value)
    return a
}

/**
 * editorSetting
 */
func (a *MonthRangeControl) EditorSetting(value interface{}) *MonthRangeControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * labelAlign
 */
func (a *MonthRangeControl) LabelAlign(value interface{}) *MonthRangeControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * validationErrors
 */
func (a *MonthRangeControl) ValidationErrors(value interface{}) *MonthRangeControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * style
 */
func (a *MonthRangeControl) Style(value interface{}) *MonthRangeControl {
    a.Set("style", value)
    return a
}

/**
 * delimiter
 */
func (a *MonthRangeControl) Delimiter(value interface{}) *MonthRangeControl {
    a.Set("delimiter", value)
    return a
}

/**
 * description
 */
func (a *MonthRangeControl) Description(value interface{}) *MonthRangeControl {
    a.Set("description", value)
    return a
}

/**
 * validations
 */
func (a *MonthRangeControl) Validations(value interface{}) *MonthRangeControl {
    a.Set("validations", value)
    return a
}

/**
 * startPlaceholder
 */
func (a *MonthRangeControl) StartPlaceholder(value interface{}) *MonthRangeControl {
    a.Set("startPlaceholder", value)
    return a
}

/**
 * submitOnChange
 */
func (a *MonthRangeControl) SubmitOnChange(value interface{}) *MonthRangeControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *MonthRangeControl) ClearValueOnHidden(value interface{}) *MonthRangeControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * static
 */
func (a *MonthRangeControl) Static(value interface{}) *MonthRangeControl {
    a.Set("static", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *MonthRangeControl) StaticLabelClassName(value interface{}) *MonthRangeControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *MonthRangeControl) StaticInputClassName(value interface{}) *MonthRangeControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * maxDate
 */
func (a *MonthRangeControl) MaxDate(value interface{}) *MonthRangeControl {
    a.Set("maxDate", value)
    return a
}

/**
 * validateOnChange
 */
func (a *MonthRangeControl) ValidateOnChange(value interface{}) *MonthRangeControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * hidden
 */
func (a *MonthRangeControl) Hidden(value interface{}) *MonthRangeControl {
    a.Set("hidden", value)
    return a
}

/**
 * minDate
 */
func (a *MonthRangeControl) MinDate(value interface{}) *MonthRangeControl {
    a.Set("minDate", value)
    return a
}

/**
 * borderMode
 */
func (a *MonthRangeControl) BorderMode(value interface{}) *MonthRangeControl {
    a.Set("borderMode", value)
    return a
}

/**
 * value
 */
func (a *MonthRangeControl) Value(value interface{}) *MonthRangeControl {
    a.Set("value", value)
    return a
}

/**
 * initAutoFill
 */
func (a *MonthRangeControl) InitAutoFill(value interface{}) *MonthRangeControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * className
 */
func (a *MonthRangeControl) ClassName(value interface{}) *MonthRangeControl {
    a.Set("className", value)
    return a
}

/**
 * disabled
 */
func (a *MonthRangeControl) Disabled(value interface{}) *MonthRangeControl {
    a.Set("disabled", value)
    return a
}

/**
 * onEvent
 */
func (a *MonthRangeControl) OnEvent(value interface{}) *MonthRangeControl {
    a.Set("onEvent", value)
    return a
}

/**
 * staticOn
 */
func (a *MonthRangeControl) StaticOn(value interface{}) *MonthRangeControl {
    a.Set("staticOn", value)
    return a
}

/**
 * labelClassName
 */
func (a *MonthRangeControl) LabelClassName(value interface{}) *MonthRangeControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *MonthRangeControl) DescriptionClassName(value interface{}) *MonthRangeControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * embed
 */
func (a *MonthRangeControl) Embed(value interface{}) *MonthRangeControl {
    a.Set("embed", value)
    return a
}

/**
 * endPlaceholder
 */
func (a *MonthRangeControl) EndPlaceholder(value interface{}) *MonthRangeControl {
    a.Set("endPlaceholder", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *MonthRangeControl) Width(value interface{}) *MonthRangeControl {
    a.Set("width", value)
    return a
}

/**
 * label
 */
func (a *MonthRangeControl) Label(value interface{}) *MonthRangeControl {
    a.Set("label", value)
    return a
}

/**
 * labelOverflow
 */
func (a *MonthRangeControl) LabelOverflow(value interface{}) *MonthRangeControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * horizontal
 */
func (a *MonthRangeControl) Horizontal(value interface{}) *MonthRangeControl {
    a.Set("horizontal", value)
    return a
}

/**
 * placeholder
 */
func (a *MonthRangeControl) Placeholder(value interface{}) *MonthRangeControl {
    a.Set("placeholder", value)
    return a
}

/**
 */
func (a *MonthRangeControl) Type(value interface{}) *MonthRangeControl {
    a.Set("type", value)
    return a
}

/**
 * maxDuration
 */
func (a *MonthRangeControl) MaxDuration(value interface{}) *MonthRangeControl {
    a.Set("maxDuration", value)
    return a
}
