package renderers


/**
 * QuarterRange 季度范围控件 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/input-quarter-range

 * @author wcz0
 * @version 6.2.2
 */
type QuarterRangeControl struct {
	*BaseRenderer
}

func NewQuarterRangeControl() *QuarterRangeControl {
    a := &QuarterRangeControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-quarter-range")
    return a
}


func (a *QuarterRangeControl) Set(name string, value interface{}) *QuarterRangeControl {
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
func (a *QuarterRangeControl) LabelAlign(value interface{}) *QuarterRangeControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * validateOnChange
 */
func (a *QuarterRangeControl) ValidateOnChange(value interface{}) *QuarterRangeControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * embed
 */
func (a *QuarterRangeControl) Embed(value interface{}) *QuarterRangeControl {
    a.Set("embed", value)
    return a
}

/**
 * validateApi
 */
func (a *QuarterRangeControl) ValidateApi(value interface{}) *QuarterRangeControl {
    a.Set("validateApi", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *QuarterRangeControl) DescriptionClassName(value interface{}) *QuarterRangeControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * hiddenOn
 */
func (a *QuarterRangeControl) HiddenOn(value interface{}) *QuarterRangeControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * animation
 */
func (a *QuarterRangeControl) Animation(value interface{}) *QuarterRangeControl {
    a.Set("animation", value)
    return a
}

/**
 * label
 */
func (a *QuarterRangeControl) Label(value interface{}) *QuarterRangeControl {
    a.Set("label", value)
    return a
}

/**
 * extraName
 */
func (a *QuarterRangeControl) ExtraName(value interface{}) *QuarterRangeControl {
    a.Set("extraName", value)
    return a
}

/**
 * mode
 */
func (a *QuarterRangeControl) Mode(value interface{}) *QuarterRangeControl {
    a.Set("mode", value)
    return a
}

/**
 * inputFormat
 */
func (a *QuarterRangeControl) InputFormat(value interface{}) *QuarterRangeControl {
    a.Set("inputFormat", value)
    return a
}

/**
 * maxDate
 */
func (a *QuarterRangeControl) MaxDate(value interface{}) *QuarterRangeControl {
    a.Set("maxDate", value)
    return a
}

/**
 * validations
 */
func (a *QuarterRangeControl) Validations(value interface{}) *QuarterRangeControl {
    a.Set("validations", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *QuarterRangeControl) StaticLabelClassName(value interface{}) *QuarterRangeControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * labelOverflow
 */
func (a *QuarterRangeControl) LabelOverflow(value interface{}) *QuarterRangeControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * name
 */
func (a *QuarterRangeControl) Name(value interface{}) *QuarterRangeControl {
    a.Set("name", value)
    return a
}

/**
 * className
 */
func (a *QuarterRangeControl) ClassName(value interface{}) *QuarterRangeControl {
    a.Set("className", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *QuarterRangeControl) StaticPlaceholder(value interface{}) *QuarterRangeControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * labelRemark
 */
func (a *QuarterRangeControl) LabelRemark(value interface{}) *QuarterRangeControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * horizontal
 */
func (a *QuarterRangeControl) Horizontal(value interface{}) *QuarterRangeControl {
    a.Set("horizontal", value)
    return a
}

/**
 * placeholder
 */
func (a *QuarterRangeControl) Placeholder(value interface{}) *QuarterRangeControl {
    a.Set("placeholder", value)
    return a
}

/**
 */
func (a *QuarterRangeControl) Type(value interface{}) *QuarterRangeControl {
    a.Set("type", value)
    return a
}

/**
 * borderMode
 */
func (a *QuarterRangeControl) BorderMode(value interface{}) *QuarterRangeControl {
    a.Set("borderMode", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *QuarterRangeControl) ClearValueOnHidden(value interface{}) *QuarterRangeControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *QuarterRangeControl) StaticInputClassName(value interface{}) *QuarterRangeControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * style
 */
func (a *QuarterRangeControl) Style(value interface{}) *QuarterRangeControl {
    a.Set("style", value)
    return a
}

/**
 * useMobileUI
 */
func (a *QuarterRangeControl) UseMobileUI(value interface{}) *QuarterRangeControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * valueFormat
 */
func (a *QuarterRangeControl) ValueFormat(value interface{}) *QuarterRangeControl {
    a.Set("valueFormat", value)
    return a
}

/**
 * value
 */
func (a *QuarterRangeControl) Value(value interface{}) *QuarterRangeControl {
    a.Set("value", value)
    return a
}

/**
 * format
 */
func (a *QuarterRangeControl) Format(value interface{}) *QuarterRangeControl {
    a.Set("format", value)
    return a
}

/**
 * endPlaceholder
 */
func (a *QuarterRangeControl) EndPlaceholder(value interface{}) *QuarterRangeControl {
    a.Set("endPlaceholder", value)
    return a
}

/**
 * delimiter
 */
func (a *QuarterRangeControl) Delimiter(value interface{}) *QuarterRangeControl {
    a.Set("delimiter", value)
    return a
}

/**
 * hint
 */
func (a *QuarterRangeControl) Hint(value interface{}) *QuarterRangeControl {
    a.Set("hint", value)
    return a
}

/**
 * submitOnChange
 */
func (a *QuarterRangeControl) SubmitOnChange(value interface{}) *QuarterRangeControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * required
 */
func (a *QuarterRangeControl) Required(value interface{}) *QuarterRangeControl {
    a.Set("required", value)
    return a
}

/**
 * hidden
 */
func (a *QuarterRangeControl) Hidden(value interface{}) *QuarterRangeControl {
    a.Set("hidden", value)
    return a
}

/**
 * static
 */
func (a *QuarterRangeControl) Static(value interface{}) *QuarterRangeControl {
    a.Set("static", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *QuarterRangeControl) ReadOnlyOn(value interface{}) *QuarterRangeControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * initAutoFill
 */
func (a *QuarterRangeControl) InitAutoFill(value interface{}) *QuarterRangeControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * disabledOn
 */
func (a *QuarterRangeControl) DisabledOn(value interface{}) *QuarterRangeControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * staticClassName
 */
func (a *QuarterRangeControl) StaticClassName(value interface{}) *QuarterRangeControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * maxDuration
 */
func (a *QuarterRangeControl) MaxDuration(value interface{}) *QuarterRangeControl {
    a.Set("maxDuration", value)
    return a
}

/**
 * labelWidth
 */
func (a *QuarterRangeControl) LabelWidth(value interface{}) *QuarterRangeControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * desc
 */
func (a *QuarterRangeControl) Desc(value interface{}) *QuarterRangeControl {
    a.Set("desc", value)
    return a
}

/**
 * id
 */
func (a *QuarterRangeControl) Id(value interface{}) *QuarterRangeControl {
    a.Set("id", value)
    return a
}

/**
 * onEvent
 */
func (a *QuarterRangeControl) OnEvent(value interface{}) *QuarterRangeControl {
    a.Set("onEvent", value)
    return a
}

/**
 * size
 */
func (a *QuarterRangeControl) Size(value interface{}) *QuarterRangeControl {
    a.Set("size", value)
    return a
}

/**
 * joinValues
 */
func (a *QuarterRangeControl) JoinValues(value interface{}) *QuarterRangeControl {
    a.Set("joinValues", value)
    return a
}

/**
 * ranges
 */
func (a *QuarterRangeControl) Ranges(value interface{}) *QuarterRangeControl {
    a.Set("ranges", value)
    return a
}

/**
 * remark
 */
func (a *QuarterRangeControl) Remark(value interface{}) *QuarterRangeControl {
    a.Set("remark", value)
    return a
}

/**
 * validationErrors
 */
func (a *QuarterRangeControl) ValidationErrors(value interface{}) *QuarterRangeControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * autoFill
 */
func (a *QuarterRangeControl) AutoFill(value interface{}) *QuarterRangeControl {
    a.Set("autoFill", value)
    return a
}

/**
 * minDuration
 */
func (a *QuarterRangeControl) MinDuration(value interface{}) *QuarterRangeControl {
    a.Set("minDuration", value)
    return a
}

/**
 * popOverContainerSelector
 */
func (a *QuarterRangeControl) PopOverContainerSelector(value interface{}) *QuarterRangeControl {
    a.Set("popOverContainerSelector", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *QuarterRangeControl) Width(value interface{}) *QuarterRangeControl {
    a.Set("width", value)
    return a
}

/**
 * labelClassName
 */
func (a *QuarterRangeControl) LabelClassName(value interface{}) *QuarterRangeControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * description
 */
func (a *QuarterRangeControl) Description(value interface{}) *QuarterRangeControl {
    a.Set("description", value)
    return a
}

/**
 * inputClassName
 */
func (a *QuarterRangeControl) InputClassName(value interface{}) *QuarterRangeControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * row
 */
func (a *QuarterRangeControl) Row(value interface{}) *QuarterRangeControl {
    a.Set("row", value)
    return a
}

/**
 * visibleOn
 */
func (a *QuarterRangeControl) VisibleOn(value interface{}) *QuarterRangeControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * displayFormat
 */
func (a *QuarterRangeControl) DisplayFormat(value interface{}) *QuarterRangeControl {
    a.Set("displayFormat", value)
    return a
}

/**
 * minDate
 */
func (a *QuarterRangeControl) MinDate(value interface{}) *QuarterRangeControl {
    a.Set("minDate", value)
    return a
}

/**
 * transform
 */
func (a *QuarterRangeControl) Transform(value interface{}) *QuarterRangeControl {
    a.Set("transform", value)
    return a
}

/**
 * readOnly
 */
func (a *QuarterRangeControl) ReadOnly(value interface{}) *QuarterRangeControl {
    a.Set("readOnly", value)
    return a
}

/**
 * inline
 */
func (a *QuarterRangeControl) Inline(value interface{}) *QuarterRangeControl {
    a.Set("inline", value)
    return a
}

/**
 * visible
 */
func (a *QuarterRangeControl) Visible(value interface{}) *QuarterRangeControl {
    a.Set("visible", value)
    return a
}

/**
 * staticSchema
 */
func (a *QuarterRangeControl) StaticSchema(value interface{}) *QuarterRangeControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * shortcuts
 */
func (a *QuarterRangeControl) Shortcuts(value interface{}) *QuarterRangeControl {
    a.Set("shortcuts", value)
    return a
}

/**
 * startPlaceholder
 */
func (a *QuarterRangeControl) StartPlaceholder(value interface{}) *QuarterRangeControl {
    a.Set("startPlaceholder", value)
    return a
}

/**
 * disabled
 */
func (a *QuarterRangeControl) Disabled(value interface{}) *QuarterRangeControl {
    a.Set("disabled", value)
    return a
}

/**
 * staticOn
 */
func (a *QuarterRangeControl) StaticOn(value interface{}) *QuarterRangeControl {
    a.Set("staticOn", value)
    return a
}

/**
 * editorSetting
 */
func (a *QuarterRangeControl) EditorSetting(value interface{}) *QuarterRangeControl {
    a.Set("editorSetting", value)
    return a
}
